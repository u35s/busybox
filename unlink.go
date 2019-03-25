package main

import (
	"fmt"
	"os"
	"syscall"
)

func (app *applets) Applet_unlink(args []string) {
	var retval int
	if len(args) == 2 {
		e := syscall.Unlink(args[1])
		if e != nil {
			fmt.Fprintf(os.Stderr, "%v\n", e)
			retval = 1
		}
	} else {
		fmt.Fprintf(os.Stderr, "Usage: unlink FILE\n")
		retval = 1
	}
	os.Exit(retval)
}
