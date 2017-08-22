package storage

import (
	"context"
	"database/sql"
	"log"

	"github.com/sauercrowd/timetable/pkg/festival"
)

func addLocationTransaction(ctx context.Context, tx *sql.Tx, errorCh chan error, festivalID int64, location *festival.Location) {
	res, err := tx.ExecContext(ctx, "INSERT INTO locations(festivalid, name, description) VALUES($1, $2, $3)", festivalID, location.Name, location.Description)
	if err != nil {
		log.Println("Could not execute query:", err)
		errorCh <- err
		return
	}
	locationID, err := res.LastInsertId()
	if err != nil {
		log.Println("Could not get locationid:", err)
		errorCh <- err
		return
	}
	dayChan := make(chan error)
	count := len(location.Days)
	ctx, cancel := context.WithCancel(ctx)
	//start goroutines
	for _, day := range location.Days {
		go addBandsTransaction(ctx, tx, dayChan, festivalID, locationID, day)
	}
	//wait for results
	for n := range dayChan {
		count--
		if n != nil {
			cancel()
			errorCh <- err
			return
		}
		if count == 0 {
			break
		}
	}
	close(dayChan)
	//free context ressources
	cancel()
	errorCh <- nil
}
