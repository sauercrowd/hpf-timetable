package main

import (
	"log"

	"github.com/sauercrowd/timetable/pkg/flags"
	"github.com/sauercrowd/timetable/pkg/storage"
)

func main() {
	f := flags.ParseFlags()
	s, err := storage.NewStorage(&f)
	if err != nil {
		log.Fatal("Could not do stuff", err)
	}
	_ = s

}
