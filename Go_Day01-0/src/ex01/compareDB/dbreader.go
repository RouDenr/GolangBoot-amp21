package main


import (
	"encoding/xml"
)

type DBReader interface {
	ConvertFileToMap(data []byte) (*Recipes, error)
	ConvertMapToFile(recipes Recipes) ([]byte, error)
}


type Item struct {
	Itemname  string `xml:"itemname" json:"ingredient_name"`
	Itemcount string `xml:"itemcount" json:"ingredient_count"`
	Itemunit  string `xml:"itemunit" json:"ingredient_unit,omitempty"`
}

type Cake struct {
	Name       string `xml:"name" json:"name"`
	Stovetime  string `xml:"stovetime" json:"time"`
	Ingredient []Item `xml:"ingredients>item" json:"ingredients"`
}

type Recipes struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Recipes []Cake   `xml:"cake" json:"cake"`
}
