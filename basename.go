package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (app *applets) Applet_basename(args []string) {
	var set = flag.NewFlagSet(args[0], flag.ExitOnError)
	var aflag = set.Bool("a", false, "support multiple arguments and treat each as a NAME")
	var sflag = set.String("s", "", "remove a trailing SUFFIX")
	var zflag = set.Bool("z", false, "separate output with NUL rather than newline")
	set.Parse(args[1:])
	args = set.Args()

	output := func(s string) {
		base := filepath.Base(s)
		base = strings.TrimSuffix(base, *sflag)
		if !*zflag {
			fmt.Println(base)
		} else {
			fmt.Print(base)
		}
	}
	if !*aflag && len(args) == 2 {
		*sflag = args[len(args)-1]
		args = args[:len(args)-1]
	} else if len(args) == 0 || (!*aflag && len(args) > 2) {
		set.Usage()
		os.Exit(1)
	}

	for i := range args {
		output(args[i])
	}
}
