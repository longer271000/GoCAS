package config

import (
	"encoding/json"
	"os"
)

type APPConfig struct {
	AppName    string "json:'app_name'"
	Port       string "json:'port'"
	StaticPath string "json:'static_path'"
	Mode       string "json:'mode'"
}
var ServiceConfig APPConfig

func InitConfig()*APPConfig  {
	file,err := os.Open("config.json")
	if err !=nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := APPConfig{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return &conf
}