package main

import (
	"github.com/ServiceComb/go-chassis"
	_ "github.com/feixiao/learning/gochassis/examples/plugin/handler"
	_ "github.com/feixiao/learning/gochassis/examples/schemas"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/path/to/conf/folder
func main() {
	chassis.Init()
	chassis.Run()
}
