package main

import (
	// "flag"
	"errors"
	"fmt"
	// "io/fs"
	"io/ioutil"
	"log"
	"os"
	// "fs"
)

// func print_ln(path string, file fs.FileInfo) error {
// 	strlink, err := os.Readlink(file.Name())
// 	if err == nil {
// 		fmt.Printf("%s -> %s \n", strlink)
// 	}
// 	return err
// }

func recusive_ls(path string) error {
	files, err := ioutil.ReadDir(path)
	if err == nil {

		for _, file := range files {
			new_path := path
			if path[len(path)-1] != '/' {
				new_path += "/"
			}
			new_path += file.Name()
			if file.Mode()&os.ModeSymlink != 0 {
				// strlink, err := os.Lstat(file.Name())
				strlink, err := os.Readlink(file.Name())
				if err == nil {
					fmt.Printf("%s -> %s \n", new_path, strlink)
				}
			} else if file.IsDir() {
				fmt.Printf("%s\n", new_path)
				recusive_ls(new_path)
			} else {
				fmt.Printf("%s\n", new_path)
			}
		}
	}
	return err
}

func main() {
	len_args := len(os.Args)
	if len_args >= 2 {

		path := os.Args[len_args-1]
		if err := recusive_ls(path); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(errors.New("no path argument"))
	}
}
