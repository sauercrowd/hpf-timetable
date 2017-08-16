package storage

import (
	"database/sql"

	"github.com/sauercrowd/hpf-timetable/pkg/flags"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(flags *flags.Flags) (*Storage, error) {
	var ret Storage

}

func newPostgres(flags *flags.Flags) (*sql.DB, error) {

}
