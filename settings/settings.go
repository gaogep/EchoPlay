package settings

import (
	"gopkg.in/ini.v1"
	"log"
)

var DBConf = make(map[string]string)
var GlobalConf = make(map[string]interface{})

const TimeFormt = "2006-01-02 15:04:05"

func init() {
	cfg, err := ini.Load("conf/config.ini")
	if err != nil {
		log.Fatal("config file error:", err)
	}

	LoadDBCOnf(cfg)
	LoadGlobalConf(cfg)
}

func LoadDBCOnf(cfg *ini.File) {
	DBConf["DBTYPE"] = cfg.Section("database").Key("DBTYPE").Value()
	DBConf["USER"] = cfg.Section("database").Key("USER").Value()
	DBConf["PASSWORD"] = cfg.Section("database").Key("PASSWORD").Value()
	DBConf["DBNAME"] = cfg.Section("database").Key("DBNAME").Value()
	DBConf["HOST"] = cfg.Section("database").Key("HOST").Value()
}

func LoadGlobalConf(cfg *ini.File) {
	GlobalConf["PAGESIZE"] = cfg.Section("glb").Key("PSIZE").MustInt()
	GlobalConf["JWTSECRET"] = cfg.Section("glb").Key("JWTSECRET").Value()
}
