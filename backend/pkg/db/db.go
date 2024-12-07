package db

import (
	"log"
	"time"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
)

func StartDB() (*pg.DB, error) {
	var (
		opts       *pg.Options
		err        error
		retryDelay = 2 * time.Second
	)

	opts = &pg.Options{
		Addr:     "db:5432",
		User:     "postgres",
		Password: "admin",
	}

	db := pg.Connect(opts)
	for i := 0; i < 5; i++ {
		db = pg.Connect(opts)

		_, err = db.Exec("SELECT 1")
		if err == nil {
			break
		}

		log.Printf("Failed to connect to database: %v. Retrying in %s...", err, retryDelay)
		time.Sleep(retryDelay)
	}

	collection := migrations.NewCollection()
	err = collection.DiscoverSQLMigrations("migrations")
	if err != nil {
		return nil, err
	}

	_, _, err = collection.Run(db, "init")
	if err != nil {
		return nil, err
	}

	oldVersion, newVersion, err := collection.Run(db, "up")
	if err != nil {
		return nil, err
	}
	if newVersion != oldVersion {
		log.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		log.Printf("version is %d\n", oldVersion)
	}

	return db, err
}
