package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)


func execXargs(cmd_name string, cmd_args []string) {

	cmd := exec.Command(cmd_name, cmd_args...)

	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s %v: ", cmd_name, cmd_args)
		fmt.Println(err)
	}
	fmt.Printf("%s", stdout)
}

func main() {

	if len(os.Args) < 2 {
		log.Fatalln("cmd is empty")
	}

	cmd_name, cmd_flags := os.Args[1], os.Args[2:]
	var wg sync.WaitGroup
	for {
		getnextline := bufio.NewReader(os.Stdin)
		read_args, errio := getnextline.ReadString('\n')

		switch {
		case errio == io.EOF:
			wg.Wait()
			return
		case errio != nil:
			log.Fatalln(errio)
		}
		read_args = strings.Trim(read_args, "\n")
		read_args_split := strings.Split(read_args, " ")
		cmd_args := append(cmd_flags, read_args_split...)

		wg.Add(1)
		go func(string, []string) {
			execXargs(cmd_name, cmd_args)
			defer wg.Done()
		} (cmd_name, cmd_args)

	}
}
