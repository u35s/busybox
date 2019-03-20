package main

import (
	"flag"
	"os"

	"github.com/u35s/busybox/libbb"
)

func (app *applets) Applet_truncate(args []string) {
	var set = flag.NewFlagSet("truncate [options] <filename>", flag.ExitOnError)
	var sflag = set.String("s", "", "set the file size by SIZE bytes")
	set.Parse(args[1:])
	args = set.Args()

	var (
		retval int
		size   int64
	)

	if *sflag == "" {
		set.Usage()
		retval = 1
	} else {
		size = libbb.Atoi_sfx(*sflag, libbb.Suffixes_cwbkMG)
		for i := range args {
			os.Truncate(args[i], size)
		}
	}
	os.Exit(retval)
}
