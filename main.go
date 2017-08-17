package main

import (
	"github.com/sauercrowd/hpf-timetable/pkg/flags"
	"github.com/sauercrowd/hpf-timetable/pkg/parser"
	"github.com/sauercrowd/hpf-timetable/pkg/search"
)

func main() {
	f := flags.ParseFlags()
	tt, err := parser.Parse("times/2017.yml")
	if err != nil {
		panic(err)
	}
	s := search.Setup(&f)
	s.AddFestistval(tt)
}
