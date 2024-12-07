package models

import (
	"errors"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type Album struct {
	Name string `pg:"album_name" json:"album_name"`
	Year string `pg:"album_year" json:"album_year"`
}

func GetAlbums(db *pg.DB) ([]*Album, error) {
	albums := make([]*Album, 0)
	err := db.Model(&albums).Select()

	return albums, err
}

// REPORT
func GetYearReport(db *pg.DB, direction string) (int64, error) {
	if db == nil {
		return -1, errors.New("database connection is nil")
	}

	type Result struct {
		Year      int64 `pg:"year"`
		TotalWins int64 `pg:"total_wins"`
	}
	var result Result

	query := fmt.Sprintf(`
	SELECT 
		a.album_year AS year,
		SUM(b.wins) AS total_wins
	FROM 
		albums a
	JOIN 
		songs s ON a.album_name = s.album_name
	JOIN 
		bars b ON s.title = b.song_name
	GROUP BY 
		a.album_year
	ORDER BY 
		total_wins %s
	LIMIT 1;
	`, direction)
	_, err := db.Query(&result, query)

	return result.Year, err
}

func GetAlbumReport(db *pg.DB, direction string) (*Album, error) {
	if db == nil {
		return nil, errors.New("database connection is nil")
	}

	result := make([]*Album, 0)
	query := fmt.Sprintf(`
	SELECT 
		a.album_name,
		a.album_year
	FROM 
		albums a
	JOIN 
		songs s ON a.album_name = s.album_name
	JOIN 
		bars b ON s.title = b.song_name
	GROUP BY 
		a.album_name,
		a.album_year
	ORDER BY 
		SUM(b.wins) %s
	LIMIT 1;
	`, direction)
	_, err := db.Query(&result, query)

	if len(result) == 0 {
		return nil, errors.New("none found")
	}
	return result[0], err
}
