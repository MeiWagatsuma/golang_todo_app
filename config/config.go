package config

import (
	"log"

	"../utils"
	"gopkg.in/go-ini/ini.v1"
)

type configList struct {
	Port string
	SQLDriver string
	DbName string
	LogFile string
}

var Config configList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal(err)
	}
	Config = configList{
		Port: cfg.Section("Web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName: cfg.Section("db").Key("name").String(),
		LogFile: cfg.Section("web").Key("logfile").String(),
	}
}