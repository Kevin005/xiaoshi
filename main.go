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

var cfgYaml = &conf.ConfigYaml{} //全局配置文件

func init() {
	conf.LoadCfgYaml(cfgYaml)
}

func main() {
	logcf := flag.String("lf", "log4go.xml", "log xml file path")
	flag.Parse()
	if len(*logcf) > 0 {
		log.LoadConfiguration(*logcf)
	}
	log.Info("======= main start =======")
	app := &handler.App{}
	app.Initialize(&conf.Config{
		&conf.DBConfig{
			Dialect:  cfgYaml.DB.Dialect,
			Username: cfgYaml.DB.Username,
			Password: cfgYaml.DB.Password,
			DBName:   cfgYaml.DB.DBName,
			Charset:  cfgYaml.DB.Charset,
		},
	})
	app.Run(":" + cfgYaml.Server.Port)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-c
	log.Info("======= main end =======")
}
