package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func (app *applets) Applet_pwd(args []string) {
	var set = flag.NewFlagSet(args[0], flag.ExitOnError)
	var pflag = set.Bool("P", false, "print the current directory, and resolve all symlinks")
	set.Parse(args[1:])
	args = set.Args()

	var retval int
	defer func() { os.Exit(retval) }()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		retval = 1
	} else {
		if *pflag {
			tmp, err2 := filepath.EvalSymlinks(dir)
			if err2 == nil {
				dir = tmp
			}
		}
	}

	fmt.Println(dir)
}
