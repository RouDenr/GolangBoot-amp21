package main

import (
	"errors"
	"flag"
	"fmt"

	// "fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/r3labs/diff"
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
	new_f := flag.String("new", "", "json or xml")
	old_f := flag.String("old", "", "json or xml")
	flag.Parse()
	var err error
	if *new_f != "" && *old_f != "" {
		new_reader, err := GetDBReader(*new_f)
		if err != nil {
			log.Fatal(err)
		}
		old_reader, err := GetDBReader(*old_f)
		if err != nil {
			log.Fatal(err)
		}
		old_recipes, err := Handler(*old_f, old_reader)
		if err != nil {
			log.Fatal(err)
		}
		new_recipes, err := Handler(*new_f, new_reader)
		if err != nil {
			log.Fatal(err)
		}
		if reflect.DeepEqual(old_recipes, new_recipes) {
			fmt.Println("Equal")
		} else {
			compare(old_recipes, new_recipes)
		}
	} else {
		err = errors.New("invalid flags")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func pairs_change(p []string, recipies *Recipes) string {
	pairs := make([]string, len(p)/2+len(p)%2)
	var a, b int
	var index_rec, index_ing int
	for a = len(pairs) - 1; b < len(p)&^1; b, a = b+2, a-1 {
		var arg string
		if p[b] == "Recipes" {
			index_rec, _ = strconv.Atoi(p[b+1])
			arg = (*recipies).Recipes[index_rec].Name
		} else if p[b] == "Ingredient" {
			index_ing, _ = strconv.Atoi(p[b+1])
			arg = (*recipies).Recipes[index_rec].Ingredient[index_ing].Itemname
		} else {
			arg = p[b+1]
		}
		pairs[a] = fmt.Sprintf("%s %s", p[b], arg)
	}
	if a == 0 {
		pairs[a] = p[b]
	}
	return strings.Join(pairs, " for ")
}

func compare(old_recipes, new_recipes *Recipes) {

	differ, err := diff.NewDiffer(diff.DisableStructValues())
	if err == nil {
		diflog, err := differ.Diff(old_recipes, new_recipes)
		if err == nil {
			for _, v := range diflog {
				if v.Path[0] == "XMLName" {
					continue
				}
				path := v.Path
				if v.Type == diff.CREATE || v.From == "" {
					fmt.Printf("ADDED %s\n", pairs_change(path, new_recipes))
				} else if v.Type == diff.DELETE || v.To == "" {
					switch n := len(path) - 1; path[n] {
					case "unit":
						path = append(path, v.From.(string))
					case "ingredient":
						path = path[:n]
					}
					fmt.Printf("REMOVED %s\n", pairs_change(path, old_recipes))
				} else if v.Type == diff.UPDATE {
					fmt.Printf("CHANGED %s - %s instead of %s\n", pairs_change(path, old_recipes), v.To, v.From)
				}
			}
		}
	}
	if err != nil {
		log.Fatal(err)
	}

}
