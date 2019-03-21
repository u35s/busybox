package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
)

func (app *applets) Applet_whoami(args []string) {
	var set = flag.NewFlagSet("whoami [OPTION]... ", flag.ExitOnError)
	var hflag = set.Bool("h", false, "display this help and exit")
	set.Parse(args[1:])
	args = set.Args()

	if *hflag {
		set.Usage()
		fmt.Fprintf(os.Stderr, "%v\n",
			"Print the user name associated with the current effective user ID")
	} else {
		u, err := user.Current()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		} else {
			fmt.Println(u.Username)
		}
	}
}
