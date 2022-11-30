package main

import (
	"encoding/json"
	// "fmt"
)

type DBReaderJSON struct {
}

func GetDBReaderJSON() *DBReaderJSON {
	return &DBReaderJSON{}
}

func (this *DBReaderJSON) ConvertFileToMap(data []byte) (*Recipes, error) {
	var recipes Recipes

	// fmt.Printf("data: %s\n", data)
	err := json.Unmarshal(data, &recipes)
	// fmt.Println(recipes)
	if err != nil {
		return nil, err
	}
	return &recipes, nil
}

func (this *DBReaderJSON) ConvertMapToFile(recipes Recipes) ([]byte, error) {
	if data, err := json.MarshalIndent(recipes, "", "  "); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

// func third_party() {
// 	x := map[string]interface{}{"a": 1, "b": 2}
// 	b, err := json.MarshalIndent(x, "", "    ")
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	fmt.Print(string(b))
// }
