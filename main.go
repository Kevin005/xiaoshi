package main

import (
	"xiaoshi/util"
	"xiaoshi/app"
)

func main() {
	config := util.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":9090")
}
