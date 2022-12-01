package main

import (
	"errors"
	"sync"
	// "flag"
	"fmt"
	"log"
	"os"
)

//  -l for counting lines, -m for counting characters and -w for counting words.

func count_words(arg string) (int, error) {
	return 0, nil
}
func count_lines(arg string) (int, error) {
	return 0, nil
}
func count_char(arg string) (int, error) {
	return 0, nil
}

func myWc(func_count func(string) (int, error), arg string) {
	result, err := func_count(arg)
	// fmt.Println("Hi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d\t%s\n", result, arg)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln(errors.New("./myWc -[w, m, l] input.txt"))
	}
	flag_count := os.Args[1]
	if flag_count != "-w" && flag_count != "-m" && flag_count != "-l" {
		flag_count = "-w"
	}

	var func_count func(arg string) (int, error)

	switch flag_count {
	case "-w":
		func_count = count_words
	case "-m":
		func_count = count_char
	case "-l":
		func_count = count_lines
	default:

	}

	var wg sync.WaitGroup
	for i, arg := range os.Args {
		if i == 0 || (i == 1 && arg == flag_count) {
			continue
		}
		wg.Add(1)
		go func (arg string)  {
			myWc(func_count, arg)
			defer wg.Done()
		}(arg)
	}
	wg.Wait()
}
