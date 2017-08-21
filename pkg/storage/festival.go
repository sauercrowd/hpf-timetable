package storage

import (
	"context"
	"database/sql"

	"github.com/sauercrowd/timetable/pkg/festival"
)

func (s *Storage) AddFestival(ctx context.Context, festival *festival.Festival) error {
	var festivalID int
	err := s.db.QueryRowContext(ctx, "INSERT INTO festivals(name) VALUES($1) RETURNING festivalid", festival.Name).Scan(&festivalID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}
