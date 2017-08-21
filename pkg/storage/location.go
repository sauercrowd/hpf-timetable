package storage

import (
	"context"

	"github.com/sauercrowd/timetable/pkg/festival"
)

func (s *Storage) AddLocation(ctx context.Context, festivalID int, location *festival.Location) error {
	var locationID int
	err := s.db.QueryRowContext(ctx, "INSERT INTO locations(festivalid, name, description) VALUES($1, $2, $3) RETURNING locationid", festivalID, location.Name, location.Description).Scan(&locationID)
	if err != nil {
		return err
	}
	return nil
}
