package storage

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sauercrowd/timetable/pkg/festival"
)

func (s *Storage) AddBands(ctx context.Context, festivalID int, locationID int, day festival.Day) error {
	ch := make(chan error)
	var mtx sync.Mutex
	count := len(day.Bands)
	ctx, cancel := context.WithCancel(ctx)
	for _, band := range day.Bands {
		go s.AddBandConcurrent(ctx, &mtx, &count, ch, festivalID, locationID, day.Date, band)
	}
	for n := range ch {
		if n != nil {
			cancel()
			return n
		}
	}
	//free resources
	cancel()
	return nil
}

func (s *Storage) AddBandConcurrent(ctx context.Context, mtx *sync.Mutex, count *int, ch chan error, festivalID int, locationID int, date string, band festival.Band) {
	start, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", date, band.Start))
	if err != nil {
		ch <- err
	}
	stop, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", date, band.Stop))
	if err != nil {
		ch <- err
	}
	err = s.db.QueryRowContext(ctx, "INSERT INTO bands(festivalid, locationid, name, start, stop, imageurl, infourl) VALUES($1, $2, $3, $4, $5, $6, $7)",
		festivalID, locationID, band.Name, start, stop, "", "").Scan()
	if err != nil {
		ch <- err
	}
	//reduced count and close channel if equal to zero
	mtx.Lock()
	newC := (*count) - 1
	count = &newC
	mtx.Unlock()
	if newC <= 0 {
		close(ch)
	}
}
