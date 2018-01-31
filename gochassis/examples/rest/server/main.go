package main

import (
	"github.com/ServiceComb/go-chassis"
	"github.com/ServiceComb/go-chassis/core/lager"
	"github.com/feixiao/learning/gochassis/examples/schemas"
	serverOption "github.com/ServiceComb/go-chassis/third_party/forked/go-micro/server"
	"os"
)

//	if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/path/to/conf/folder

func main() {
	os.Setenv("CHASSIS_HOME","rest/server")
	chassis.RegisterSchema("rest", &schemas.RestFulHello{}, serverOption.WithSchemaID("RestHelloService"))
	if err := chassis.Init(); err != nil {
		lager.Logger.Error("Init failed.", err)
		return
	}
	chassis.Run()
}
