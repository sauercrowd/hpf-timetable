package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/sauercrowd/timetable/pkg/festival"
)

func addBandsTransaction(ctx context.Context, tx *sql.Tx, errorCh chan error, festivalID int64, locationID int64, day festival.Day) {
	bandsCh := make(chan error)
	ctx, cancel := context.WithCancel(ctx)
	for _, band := range day.Bands {
		go addBandConcurrentTransaction(ctx, tx, bandsCh, festivalID, locationID, day.Date, band)
	}
	count := len(day.Bands)
	for n := range bandsCh {
		count--
		if n != nil {
			cancel()
			errorCh <- n
			return
		}
		if count <= 0 {
			break
		}
	}
	close(bandsCh)
	//free resources from context
	cancel()
	errorCh <- nil
}

func addBandConcurrentTransaction(ctx context.Context, tx *sql.Tx, errorCh chan error, festivalID int64, locationID int64, date string, band festival.Band) {
	start, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", date, band.Start))
	if err != nil {
		errorCh <- err
		return
	}
	stop, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", date, band.Stop))
	if err != nil {
		errorCh <- err
		return
	}
	//don't care about result
	_, err = tx.ExecContext(ctx, "INSERT INTO bands(festivalid, locationid, name, start, stop, imageurl, infourl) VALUES($1, $2, $3, $4, $5, $6, $7)",
		festivalID, locationID, band.Name, start, stop, "", "")
	errorCh <- err
}
