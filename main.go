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
	app.Applet_main(os.Args)
}

func (app *applets) Applet_main(args []string) {
	_, appletName := filepath.Split(args[0])
	obj := reflect.ValueOf(app)
	functionName := fmt.Sprintf("Applet_%v", appletName)
	function := obj.MethodByName(functionName)
	if function != (reflect.Value{}) {
		values := []reflect.Value{reflect.ValueOf(args[1:])}
		function.Call(values)
	} else {
		fmt.Print("applet not found\n")
	}
}
