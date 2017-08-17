package search

import (
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/sauercrowd/hpf-timetable/pkg/flags"
	"github.com/sauercrowd/hpf-timetable/pkg/parser"
)

type Search struct {
	algolia        *algoliasearch.Client
	timeTableIndex *algoliasearch.Index
}

func Setup(f *flags.Flags) *Search {
	var s Search
	ac := algoliasearch.NewClient(f.AlgoliaID, f.AlgoliaKey)
	algIndex := ac.InitIndex("haldern pop")
	s.algolia = &ac
	s.timeTableIndex = &algIndex
	return &s
}

func (s *Search) AddFestistval(tt *parser.TimeTable) {
	entries := make([]algoliasearch.Object, 0)
	for _, location := range tt.Locations {
		for _, day := range location.Days {
			for _, band := range day.Bands {
				e := algoliasearch.Object{
					"location": location.Name,
					"day":      day.Date,
					"band":     band.Name,
					"start":    band.Start,
					"end":      band.Stop,
				}
				entries = append(entries, e)
			}
		}
	}
	(*s.timeTableIndex).AddObjects(entries)
}
