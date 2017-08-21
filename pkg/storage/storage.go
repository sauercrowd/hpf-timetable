package storage

import (
	"database/sql"
	"log"

	"github.com/sauercrowd/timetable/pkg/flags"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(flags *flags.Flags) (*Storage, error) {
	var ret Storage
	db, err := newPostgres(flags)
	if err != nil {
		return nil, err
	}
	ret.db = db
	if err := setupPostgres(db); err != nil {
		log.Fatal(err)
	}
	return &ret, nil
}
