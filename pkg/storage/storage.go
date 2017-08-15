package	storage

import (
	"database/sql"
	"github.com/sauercrowd/hpf-timetable/pkg/flags"
)

type Storage struct{
	db *sql.DB
}

func NewStorage(flags *)