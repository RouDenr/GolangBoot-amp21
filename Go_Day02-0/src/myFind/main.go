package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"io/ioutil"
	"log"
	"os"
)

func recusive_ls(path string, fl _flags) error {
	files, err := ioutil.ReadDir(path)
	if err == nil {

		for _, file := range files {
			new_path := path
			if path[len(path)-1] != '/' {
				new_path += "/"
			}
			new_path += file.Name()
			if file.Mode()&os.ModeSymlink != 0 {
				if fl.sl {
					//! check broken sl
					strlink, err := os.Readlink(file.Name())
					if err == nil {
						fmt.Printf("%s -> %s \n", new_path, strlink)
					}
				}
			} else if file.IsDir() {
				if fl.d {
					fmt.Printf("%s\n", new_path)
				}
				err = recusive_ls(new_path, fl)
			} else {
				if fl.f && strings.HasSuffix(new_path, fl.ext_str) {
					fmt.Printf("%s\n", new_path)
				}
			}
		}
	}
	return err
}

// -sl, -d or -f
// -ext (works ONLY when -f is specified)

// # Finding only *.go files ignoring all the rest.
// ~$ ./myFind -f -ext 'go' /go
// /go/src/github.com/mycoolproject/main.go
// /go/src/github.com/mycoolproject/magic.go

type _flags struct {
	sl, d, f bool
	ext_str  string
}

func set_flags(flags *_flags) {
	flag.BoolVar(&flags.d, "d", false, "find dirs")
	flag.BoolVar(&flags.f, "f", false, "find files")
	flag.BoolVar(&flags.sl, "sl", false, "find symlinks")
	flag.StringVar(&flags.ext_str, "ext", "", "file's ext")

	flag.Parse()

	if !(flags.d || flags.f || flags.sl) {
		flags.d, flags.f, flags.sl = true, true, true
	}

}

func main() {
	len_args := len(os.Args)
	if len_args >= 2 {

		var flags _flags
		set_flags(&flags)

		path := os.Args[len_args-1]
		if err := recusive_ls(path, flags); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(errors.New("no path argument"))
	}
}
