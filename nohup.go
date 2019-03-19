package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/mattn/go-isatty"
	"github.com/u35s/busybox/libbb"
)

func (app *applets) Applet_nohup(args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: nohup [--] utility [arguments]\n")
		os.Exit(1)
	}

	if isatty.IsTerminal(os.Stdin.Fd()) {
		os.Stdin.Close()
		os.OpenFile(os.DevNull, os.O_RDONLY, 0000)
	}

	nohupout := "nohup.out"
	if isatty.IsTerminal(os.Stdout.Fd()) {
		os.Stdout.Close()
		var (
			err error
			fd  int
		)
		fd, err = libbb.Open(nohupout, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			home := os.Getenv("HOME")
			if len(home) > 0 {
				nohupout = fmt.Sprintf("%v/%v", strings.TrimSuffix(home, "/"), nohupout)
				fd, _ = libbb.Open(nohupout, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
			} else {
				os.OpenFile("/dev/null", os.O_RDONLY, 0000)
			}
		}
		os.Stdout = os.NewFile(uintptr(fd), "/dev/stdout")
		fmt.Fprintf(os.Stderr, "appending output to %s\n", nohupout)
	}

	if isatty.IsTerminal(os.Stderr.Fd()) {
		syscall.Dup2(1, int(os.Stderr.Fd()))
	}

	signal.Ignore(syscall.SIGHUP)
	if !app.exec(args[1:]) {
		syscall.Exec(args[1], args[1:], os.Environ())
	}
}
