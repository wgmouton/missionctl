package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type About struct {
	Version string `yaml:"version"`
	Test    string `yaml:"test"`
}

func (about *About) Parse(file string) *About {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	yamlFile = []byte(os.ExpandEnv(string(yamlFile)))
	err = yaml.Unmarshal(yamlFile, about)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return about
}
