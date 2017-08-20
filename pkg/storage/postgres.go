package storage

import (
	"database/sql"
	"fmt"
	"log"

	//Postgres
	_ "github.com/lib/pq"

	"github.com/sauercrowd/hpf-timetable/pkg/flags"
)

const dbname = "hpftimetable"

func newPostgres(flags *flags.Flags) (*sql.DB, error) {
	if err := postgresCreateDatabaseIfNotExist(flags, dbname); err != nil {
		return nil, err
	}
	dbstr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", flags.PostgresUser, flags.PostgresPass, flags.PostgresHost, flags.PostgresPort, dbname)
	return sql.Open("postgres", dbstr)
}

func setupPostgres(db *sql.DB) error {
	if err := createFestivalTable(db); err != nil {
		return err
	}
	if err := createLocationTable(db); err != nil {
		return err
	}
	if err := createBandTable(db); err != nil {
		return err
	}
	return nil
}

// Creates a temporary connection to check if the database exists and create it if not
func postgresCreateDatabaseIfNotExist(flags *flags.Flags, dbname string) error {
	dbstr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", flags.PostgresUser, flags.PostgresPass, flags.PostgresHost, flags.PostgresPort, flags.PostgresUser)
	db, err := sql.Open("postgres", dbstr)
	if err != nil {
		return err
	}
	var count int64
	err = db.QueryRow("SELECT COUNT(1) FROM pg_database WHERE datname = $1", dbname).Scan(&count)
	//return if database exists or error happend
	if err != nil || count == 1 {
		if err == nil {
			err = db.Close()
		}
		return err
	}
	err = db.QueryRow(fmt.Sprintf("CREATE DATABASE %s", dbname)).Scan()
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("Could not create database %s: %v", dbname, err)
		return err
	}
	if err := db.Close(); err != nil {
		log.Println("Could not close temporary dbconnection: ", err)
		return err
	}
	return nil
}

//create tables
func createFestivalTable(db *sql.DB) error {
	err := db.QueryRow("CREATE TABLE IF NOT EXISTS festivals(festivalid SERIAL PRIMARY KEY, name TEXT)").Scan()
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func createLocationTable(db *sql.DB) error {
	err := db.QueryRow("CREATE TABLE IF NOT EXISTS locations(locationid SERIAL PRIMARY KEY, festivalid INTEGER REFERENCES festivals (festivalid), name TEXT, description TEXT)").Scan()
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func createBandTable(db *sql.DB) error {
	err := db.QueryRow("CREATE TABLE IF NOT EXISTS bands(bandid SERIAL PRIMARY KEY, festivalid INTEGER REFERENCES festivals (festivalid), locationid INTEGER REFERENCES locations (locationid), name TEXT, start timestamp, stop timestamp)").Scan()
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}
