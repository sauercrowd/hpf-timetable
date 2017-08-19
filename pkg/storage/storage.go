package storage

import (
	"database/sql"
	"fmt"

	"github.com/sauercrowd/hpf-timetable/pkg/flags"
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
	return &ret, nil
}

func newPostgres(flags *flags.Flags) (*sql.DB, error) {
	dbstr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", flags.PostgresUser, flags.PostgresPass, flags.PostgresHost, flags.PostgresPort, "hpf-timetable")
	return sql.Open("postgres", dbstr)
}
