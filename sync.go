package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
)

func (app *applets) Applet_sync(args []string) {
	var set = flag.NewFlagSet("sync [OPTION] FILE...", flag.ExitOnError)
	var dflag = set.Bool("d", false, "Avoid syncing metadata")
	set.Parse(args[1:])
	args = set.Args()

	var retval int
	if len(args) == 0 {
		syscall.Sync()
	} else {
		for i := range args {
			file, err := os.Open(args[i])
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
				retval = 1
			} else {
				if *dflag {
					err := syscall.Fdatasync(int(file.Fd()))
					if err != nil {
						fmt.Fprintf(os.Stderr, "%v", err)
						retval = 1
					}
				} else {
					err := syscall.Fsync(int(file.Fd()))
					if err != nil {
						fmt.Fprintf(os.Stderr, "%v", err)
						retval = 1
					}
				}
				file.Close()
			}
		}
	}
	os.Exit(retval)
}
