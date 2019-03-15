package main

import "fmt"

func (app *applets) Applet_echo(args []string) {
	for i := range args {
		fmt.Print(args[i], " ")
	}
	fmt.Println()
}
