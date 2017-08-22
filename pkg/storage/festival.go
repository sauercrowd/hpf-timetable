package storage

import (
	"context"
	"database/sql"
	"log"

	"github.com/sauercrowd/timetable/pkg/festival"
)

func (s *Storage) AddWholeFestival(ctx context.Context, festival *festival.Festival) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Could not begin transaction:", err)
		return err
	}

	res, err := tx.ExecContext(ctx, "INSERT INTO festivals(name) VALUES($1)", festival.Name)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	festivalID, err := res.LastInsertId()
	if err != nil {
		log.Println("Could not get festivalid:", err)
		return err
	}
	//create channel for errors, how many results are expected and create a cancel function in case a error happens
	locationChan := make(chan error)
	count := len(festival.Locations)
	ctx, cancel := context.WithCancel(ctx)
	//start goroutines
	for _, location := range festival.Locations {
		go addLocationTransaction(ctx, tx, locationChan, festivalID, &location)
	}
	//wait for results
	for n := range locationChan {
		count--
		if n != nil {
			cancel()
			if err := tx.Rollback(); err != nil {
				log.Println("Could not rollback transaction after error", err)
				return err
			}
		}
		if count == 0 {
			break
		}
	}
	close(locationChan)
	tx.Commit()
	//free context ressources
	cancel()
	return nil
}
