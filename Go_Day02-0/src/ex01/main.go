package main

import (
	"bytes"
	"errors"
	"io"
	"sync"
	// "unicode"

	// "flag"
	"fmt"
	"log"
	"os"
)

func isspace(c byte) bool {
	return c == '\t' || c == '\n' || c == '\v' || c == '\f' || c == '\r' || c == ' '
}

func count_words_in_buf(buf []byte) int {
	var new_word bool = true
	var count int

	for _, v := range buf {
		if isspace(v) {
			new_word = true
			continue
		}
		if new_word {
			new_word = false
			count++
		}
	}
	return count
}

func count_words(r os.File) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0

	for {
		c, err := r.Read(buf)
		count += count_words_in_buf(buf[:c])

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
func count_lines(r os.File) (int, error) {

	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
func count_char(r os.File) (int, error) {
	return 0, nil
}

func myWc(func_count func(os.File) (int, error), arg string) error {
	r, err := os.Open(arg)
	if err == nil {

		result, err := func_count(*r)
		if err == nil {
			fmt.Printf("%d\t%s\n", result, arg)
		}
		defer r.Close()
	}
	return err
}

//  -l for counting lines, -m for counting characters and -w for counting words.
func switch_count_func(flag_count string) func(os.File) (int, error) {

	switch flag_count {
	case "-w":
		return count_words
	case "-m":
		return count_char
	case "-l":
		return count_lines
	default:
		return nil
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

	func_count := switch_count_func(flag_count)

	var wg sync.WaitGroup
	for i, arg := range os.Args {
		if i == 0 || (i == 1 && arg == flag_count) {
			continue
		}
		wg.Add(1)
		go func(arg string) {
			if err := myWc(func_count, arg); err != nil {
				fmt.Println(err)
			}
			defer wg.Done()
		}(arg)
	}
	wg.Wait()
}
