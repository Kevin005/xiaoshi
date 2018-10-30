package conf

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	log "github.com/alecthomas/log4go"
)

var cfgPar []byte

type ConfigYaml struct {
	Server serverInfo
	Test      string
}

type serverInfo struct {
	Ip string
}

func init() {
	var err error;
	if cfgPar, err = ioutil.ReadFile("config.yaml"); err != nil {
		log.Error(err.Error())
		return
	}
}

func GetCfgYaml(cfg *ConfigYaml) {
	if err := yaml.Unmarshal(cfgPar, &cfg); err != nil {
		log.Error(err.Error())
	}
}
