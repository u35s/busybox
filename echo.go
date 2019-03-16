package main

import (
	"flag"
	"fmt"
	"strings"
)

func (app *applets) Applet_echo(args []string) {
	var set = flag.NewFlagSet(args[0], flag.ExitOnError)
	var eflag = set.Bool("e", false, "disable escape sequences")
	var nflag = set.Bool("n", false, `disable print \n`)
	set.Parse(args)
	args = set.Args()

	for i := range args {
		if *eflag {
			chars1 := []string{`\a`, `\b`, `\e`, `\f`, `\n`, `\r`, `\t`, `\v`, `\\`, `\0`}
			chars2 := []byte{'\a', '\b', 27, '\f', '\n', '\r', '\t', '\v', '\\', '\\'}
			for j := range chars1 {
				args[i] = strings.ReplaceAll(args[i], string(chars1[j]), string(chars2[j]))
			}
		}
		fmt.Print(args[i])
		if i+1 != len(args) {
			fmt.Print(" ")
		}
	}
	if !(*nflag) {
		fmt.Println()
	}
}
