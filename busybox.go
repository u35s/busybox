package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func (app *applets) Applet_busybox(args []string) {
	if len(args) > 1 {
		app.main(args[1:])
	} else {
		var buf bytes.Buffer
		obj := reflect.TypeOf(app)
		var col int
		var n int
		for i := 0; i < obj.NumMethod(); i++ {
			var fnname string
			fnname = obj.Method(i).Name
			if strings.HasPrefix(fnname, "Applet_") {
				fnname = strings.TrimPrefix(fnname, "Applet_")
			} else {
				continue
			}
			if col == 0 {
				buf.WriteString("\t")
			} else {
				buf.WriteString(",")
			}
			col++
			n, _ = buf.WriteString(fnname)
			col += n
			if col > 60 {
				buf.WriteString(",\n")
				col = 0
			}
		}
		buf.WriteString("\n")

		fmt.Fprintf(os.Stderr, "%v\n\n"+
			"Usage: busybox [function] [arguments]...\n"+
			"    or: [function] [arguments]...\n\n"+
			"\tBusyBox is a multi-call binary that combines many common Unix\n"+
			"\tutilities into a single executable.  Most people will create a\n"+
			"\tlink to busybox for each function they wish to use, and BusyBox\n"+
			"\twill act like whatever it was invoked as.\n"+
			"\nCurrently defined functions:\n%v", FullVersion, buf.String())
	}
}
