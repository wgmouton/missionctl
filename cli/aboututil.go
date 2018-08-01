package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type About struct {
	Version string `yaml:"version"`
}

func (about *About) Parse(file string) *About {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, about)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return about
}
