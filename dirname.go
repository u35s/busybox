package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (app *applets) Applet_dirname(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: dirname path")
		os.Exit(1)
	}
	dir := args[1]
	if len(dir) > 1 {
		dir = strings.TrimSuffix(args[1], "/")
	}
	fmt.Println(filepath.Dir(dir))
}
