package main

import (
	"log"
	"net/http"

	"github.com/sauercrowd/timetable/pkg/flags"
	"github.com/sauercrowd/timetable/pkg/storage"
	"github.com/sauercrowd/timetable/pkg/web"
)

func main() {
	f := flags.ParseFlags()
	s, err := storage.NewStorage(&f)
	if err != nil {
		log.Fatal("Could not do stuff", err)
	}
	_ = s
	web.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))

}
