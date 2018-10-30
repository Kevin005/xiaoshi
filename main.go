package main

import (
	"xiaoshi/conf"
	"xiaoshi/handler"
	log "github.com/alecthomas/log4go"
	"flag"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logcf := flag.String("lf", "log4go.xml", "log xml file path")
	flag.Parse()
	if len(*logcf) > 0 {
		log.LoadConfiguration(*logcf)
	}
	log.Info("======= main start =======")
	config := conf.GetConfig()
	app := &handler.App{}
	app.Initialize(config)
	app.Run(":9090")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-c
	log.Info("======= main end =======")
}
