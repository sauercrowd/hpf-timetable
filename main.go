package main

import (
	"fmt"

	"github.com/sauercrowd/hpf-timetable/pkg/parser"
)

func main() {
	tt, err := parser.Parse("times/2017.yml")
	if err != nil {
		panic(err)
	}
	fmt.Println(tt)
}
