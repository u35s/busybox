package main

import (
	"fmt"
	"os"
	"syscall"
)

func (app *applets) Applet_link(args []string) {
	var retval int
	if len(args) == 3 {
		e := syscall.Link(args[1], args[2])
		if e != nil {
			fmt.Fprintf(os.Stderr, "%v\n", e)
			retval = 1
		}
	} else {
		fmt.Fprintf(os.Stderr, "Usage: link FILE1 FILE2\n")
		retval = 1
	}
	os.Exit(retval)
}
