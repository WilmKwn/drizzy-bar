package models

import (
	"errors"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type Song struct {
	Name    string `pg:"title" json:"title"`
	AlbumID string `pg:"album_name" json:"album_name"`
}

func GetSongs(db *pg.DB) ([]*Song, error) {
	songs := make([]*Song, 0)
	err := db.Model(&songs).Select()

	return songs, err
}

func GetSongReport(db *pg.DB, direction string) (*Song, error) {
	if db == nil {
		return nil, errors.New("database connection is nil")
	}

	result := make([]*Song, 0)
	query := fmt.Sprintf(`
	SELECT 
		s.title,
		s.album_name 
	FROM 
		songs s
	JOIN 
		bars b ON s.title = b.song_name
	GROUP BY 
		s.title
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
