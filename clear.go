package main

import (
	"fmt"
)

func (app *applets) Applet_clear(args []string) {
	fmt.Print("\033[H\033[J")
}
