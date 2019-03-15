package main

import (
	"fmt"
	"os"
)

func (app *applets) Applet_busybox(args []string) {
	if len(args) > 0 {
		app.Applet_main(args)
	} else {
		fmt.Fprintf(os.Stderr, `%v

Usage: busybox [function] [arguments]...
   or: [function] [arguments]...

BusyBox is a multi-call binary that combines many common Unix
utilities into a single executable.  Most people will create a
link to busybox for each function they wish to use, and BusyBox
will act like whatever it was invoked as. 
Currently defined functions:%v`, FullVersion, "\n")
	}
}
