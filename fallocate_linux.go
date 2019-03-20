package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"

	"github.com/u35s/busybox/libbb"
)

func (app *applets) Applet_fallocate(args []string) {
	var set = flag.NewFlagSet("fallocate [options] <filename>", flag.ExitOnError)
	var lflag = set.String("l", "", "length of the allocation, in bytes")
	var oflag = set.String("o", "", "offset of the allocation, in bytes")
	set.Parse(args[1:])
	args = set.Args()

	var retval int
	if *lflag == "" {
		fmt.Fprintf(os.Stderr, "%v\n", "fallocate: no length argument specified")
		retval = 1
	} else if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "%v\n", "fallocate: no filename specified")
		retval = 1
	} else {
		var (
			offset int64
			length int64
			file   *os.File
			err    error
		)
		length = libbb.Atoi_sfx(*lflag, libbb.Suffixes_kmg_i)
		offset = libbb.Atoi_sfx(*oflag, libbb.Suffixes_kmg_i)

		file, err = os.OpenFile(args[0], os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			retval = 1
		} else {
			err = syscall.Fallocate(int(file.Fd()), 0, offset, length)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				retval = 1
			}
			file.Close()
		}
	}
	os.Exit(retval)
}
