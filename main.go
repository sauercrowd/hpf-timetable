package main

import (
	"log"

	"github.com/sauercrowd/hpf-timetable/pkg/flags"
	"github.com/sauercrowd/hpf-timetable/pkg/storage"
)

func main() {
	f := flags.ParseFlags()
	// tt, err := parser.Parse("times/2017.yml")
	// if err != nil {
	// 	panic(err)
	// }
	// s := search.Setup(&f)
	// s.AddFestistval(tt)
	s, err := storage.NewStorage(&f)
	if err != nil {
		log.Fatal("Could not do stuff", err)
	}
	_ = s

}
