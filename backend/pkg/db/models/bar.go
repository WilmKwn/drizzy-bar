package models

import (
	"errors"

	"github.com/go-pg/pg/v10"
)

type Bar struct {
	ID       int64  `json:"id"`
	Lyrics   string `json:"lyrics"`
	Wins     int64  `json:"wins"`
	SongName string `json:"song_name"`
}

func GetBars(db *pg.DB) ([]*Bar, error) {
	if db == nil {
		return nil, errors.New("database connection is nil")
	}

	bars := make([]*Bar, 0)
	err := db.Model(&bars).Select()

	return bars, err
}

func GetTwoRandomBars(db *pg.DB) ([]*Bar, error) {
	if db == nil {
		return nil, errors.New("database connection is nil")
	}

	resultsChan := make(chan *Bar, 2)
	errorsChan := make(chan error, 2)

	fetchRandomBar := func() {
		var bar Bar
		query := `
            SELECT id, lyrics, wins, song_name 
            FROM bars 
            ORDER BY RANDOM() 
            LIMIT 1`
		_, err := db.QueryOne(&bar, query)
		if err != nil {
			errorsChan <- err
			return
		}
		resultsChan <- &bar
	}

	go fetchRandomBar()
	go fetchRandomBar()

	bars := make([]*Bar, 0, 2)
	for i := 0; i < 2; i++ {
		select {
		case bar := <-resultsChan:
			bars = append(bars, bar)
		case err := <-errorsChan:
			return nil, err
		}
	}

	return bars, nil
}

func GetLeaderboard(db *pg.DB) ([]*Bar, error) {
	if db == nil {
		return nil, errors.New("database connection is nil")
	}

	bars := make([]*Bar, 0)
	query := `
	SELECT id, lyrics, wins, song_name 
	FROM bars 
	ORDER BY wins DESC 
	LIMIT 10`
	_, err := db.Query(&bars, query)

	return bars, err
}

func UpdateBar(db *pg.DB, id int64) (bool, error) {
	if db == nil {
		return false, errors.New("database connection is nil")
	}

	_, err := db.Model((*Bar)(nil)).Set("wins = wins + 1").Where("id = ?", id).Update()
	if err != nil {
		return false, err
	}
	return true, err
}

func DeleteBar(db *pg.DB, id int64) (bool, error) {
	if db == nil {
		return false, errors.New("database connection is nil")
	}

	_, err := db.Model((*Bar)(nil)).Where("id = ?", id).Delete()
	if err != nil {
		return false, err
	}
	return true, err
}

func CreateBar(db *pg.DB, req *Bar) (bool, error) {
	if db == nil {
		return false, errors.New("database connection is nil")
	}

	_, err := db.Model(req).Insert()
	if err != nil {
		return false, err
	}
	return true, err
}
