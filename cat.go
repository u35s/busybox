package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func (app *applets) Applet_cat(args []string) {
	var set = flag.NewFlagSet(args[0], flag.ExitOnError)
	var bflag = set.Bool("b", false, "number nonblank")
	var nflag = set.Bool("n", false, "number all output lines")
	set.Parse(args[1:])
	args = set.Args()

	var curLine int
	var retval int

	read := func(file *os.File) {
		reader := bufio.NewReader(file)
		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			} else if err != nil {
				retval = 1
				break
			}
			if *nflag && (!*bflag || len(line) > 0) {
				curLine++
			}

			if *bflag && len(line) == 0 {
				// number nonblank
				fmt.Println()
			} else if *nflag {
				// number all output lines
				fmt.Printf("%6d  %s\n", curLine, line)
			} else {
				fmt.Printf("%s\n", line)
			}
		}
	}

	if len(args) == 0 {
		read(os.Stdin)
	} else {
		for i := range args {
			if args[i] == "-" {
				read(os.Stdin)
			} else {
				f, err := os.Open(args[i])
				if err != nil {
					fmt.Println(err)
					retval = 1
				} else {
					read(f)
					f.Close()
				}
			}
		}
	}
	os.Exit(retval)
}
