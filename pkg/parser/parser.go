package parser

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Band struct {
	Name  string `yaml:"name"`
	Start string `yaml:"start"`
	Stop  string `yaml:"stop"`
}

type Day struct {
	Date  string `yaml:"day"`
	Bands []Band `yaml:"bands"`
}

type Location struct {
	Name string `yaml:"name"`
	Days []Day  `yaml:"days"`
}

type TimeTable struct {
	Locations []Location `yaml:"locations"`
}

func Parse(filepath string) (*TimeTable, error) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Println("Could not open file: ", err)
		return nil, err
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("Could not read file: ", err)
		return nil, err
	}
	var ttable TimeTable
	if err := yaml.Unmarshal(content, &ttable); err != nil {
		return nil, err
	}
	return &ttable, err
}
