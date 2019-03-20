package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

const FullVersion = "GoBusyBox 0.0.0"

type applets struct{}

func main() {
	app := &applets{}
	app.main(os.Args)
}

func (app *applets) main(args []string) {
	if !app.exec(args) {
		fmt.Fprintln(os.Stderr, "applet not found\n")
	}
}

func (app *applets) exec(args []string) bool {
	_, appletName := filepath.Split(args[0])
	obj := reflect.ValueOf(app)
	functionName := fmt.Sprintf("Applet_%v", appletName)
	function := obj.MethodByName(functionName)
	if function != (reflect.Value{}) {
		values := []reflect.Value{reflect.ValueOf(args)}
		function.Call(values)
		return true
	} else {
		return false
	}
}
