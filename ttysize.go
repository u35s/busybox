package main

import (
	"flag"
	"fmt"
	"syscall"
	"unsafe"
)

func (app *applets) Applet_ttysize(args []string) {
	var set = flag.NewFlagSet("ttysize [w] [h]", flag.ExitOnError)
	var hflag = set.Bool("h", false, "display this help and exit")
	set.Parse(args[1:])
	args = set.Args()

	if *hflag {
		set.Usage()
	} else {
		var winsize = &struct {
			Row    uint16
			Col    uint16
			Xpixel uint16
			Ypixel uint16
		}{}
		syscall.Syscall(syscall.SYS_IOCTL,
			uintptr(syscall.Stdin),
			uintptr(syscall.TIOCGWINSZ),
			uintptr(unsafe.Pointer(winsize)))

		height := int(winsize.Row)
		width := int(winsize.Col)

		if len(args) == 0 {
			fmt.Println(width, height)
		} else {
			for i := range args {
				if i != 0 {
					fmt.Print(" ")
				}

				if args[i] == "w" {
					fmt.Print(width)
				}

				if args[i] == "h" {
					fmt.Print(height)
				}
			}
			fmt.Print("\n")
		}
	}
}
