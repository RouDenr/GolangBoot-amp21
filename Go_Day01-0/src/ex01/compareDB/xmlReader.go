package main

import (
	"encoding/xml"
)

type DBReaderXML struct {
}

func GetDBReaderXML() *DBReaderXML {
	return &DBReaderXML{
	}
}
func (this *DBReaderXML) ConvertFileToMap(data []byte) (*Recipes, error) {
	var recipes Recipes
	err := xml.Unmarshal(data, &recipes)
	if err != nil {
		return nil, err
	}
	return &recipes, nil
}
func (this *DBReaderXML) ConvertMapToFile(recipes Recipes) ([]byte, error) {
	if data, err := xml.MarshalIndent(recipes, "", "    "); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}
