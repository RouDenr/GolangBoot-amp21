package main

import (
	"errors"
	"flag"
	"fmt"

	// "fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GetDBReader(name string) (*DBReader, error) {
	var reader DBReader = nil
	if strings.HasSuffix(name, ".json") {
		reader = GetDBReaderJSON()
	} else if strings.HasSuffix(name, ".xml") {
		reader = GetDBReaderXML()
	} else {
		var err error
		if name == "" {
			err = errors.New("Enter file name")
		} else {
			err = errors.New("The file: " + name + " must have extension: json or xml")
		}
		return nil, err
	}
	return &reader, nil
}
func GetDiffDBReader(reader *DBReader) (*DBReader, error) {
	var diff_reader DBReader = nil

	if _, ok := (*reader).(*DBReaderJSON); ok {
		diff_reader = GetDBReaderXML()
	} else if _, ok = (*reader).(*DBReaderXML); ok {
		diff_reader = GetDBReaderJSON()
	}

	return &diff_reader, nil
}

func Handler(name string, reader *DBReader) (*Recipes, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("data: %s\n", data)
	recipes, err := (*reader).ConvertFileToMap(data)
	defer file.Close()
	return recipes, err
}

func main() {
	f_name := flag.String("f", "", "json or xml")
	var err error
	flag.Parse()
	reader, err := GetDBReader(*f_name)
	if err == nil {
		recipes, err := Handler(*f_name, reader)
		if err == nil {
			diff_reader, err := GetDiffDBReader(reader)
			if err == nil {
				v, err := (*diff_reader).ConvertMapToFile(*recipes)
				if err == nil {
					fmt.Printf("%s\n", v)

				}
			}
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}
