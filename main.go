package main

import (
	"xiaoshi/conf"
	"xiaoshi/handler"
)

func main() {
	config := conf.GetConfig()
	app := &handler.App{}
	app.Initialize(config)
	app.Run(":9090")
}
