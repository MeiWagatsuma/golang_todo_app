package config

import (
	"log"

	"../utils"
	"gopkg.in/go-ini/ini.v1"
)

type configList struct {
	Port string
	LogFile string
	SQLDriver string
	UserName string
	Password string
	Host string
	DbName string
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
		LogFile: cfg.Section("web").Key("logfile").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		UserName: cfg.Section("db").Key("username").String(),
		Password: cfg.Section("db").Key("password").String(),
		Host: cfg.Section("db").Key("host").String(),
		DbName: cfg.Section("db").Key("dbname").String(),
	}
}