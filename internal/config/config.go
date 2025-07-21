package config

import (
	"os"
	"github.com/charmbracelet/log"
	"github.com/tidwall/gjson"
	flags "github.com/jessevdk/go-flags"
)

var sc *ServerConfig

type ServerConfig struct {
	ConfigFile	string  `short:"c" long:"config" description:"Config file"`
	PreSharedKey    string   `short:"p" long:"key" default:"" description:"Preshared key"`
	Expire	int   `short:"e" long:"expire" default:"0" description:"Credential expire time"`
	ListenAddress	string	`long:"listen" default:"0.0.0.0:3000" description:"Listen address" `
	LogLevel	string	`long:"log-level" default:"info" description:"Log level"`
	DummyUrl	string	`long:"dummy" default:"https://www.baidu.com" description:"Dummy url"`
	BaseUrl			string `long:"url" description:"Base url"`
	WhitelistPath	string 	`long:"whitelist" default:"/etc/haproxy/whitelist" description:"Location of whitelist file"`
}

func LoadConfig() {
	sc = ParseFromCmd()
	if sc.ConfigFile != "" {
		log.Infof("Reading config file: %v\n", sc.ConfigFile)
		file, err := os.ReadFile(sc.ConfigFile)
		if err != nil {
			log.Error(err)
		}
		ParseFromJson(string(file), sc)
	}
	log.Infof("Loglevel=%v, Listenaddr=%v, Baseurl=%s, PreSharedKey=%s, DummyUrl=%s, WhitelistPath=%s, QR expire=%v" , 
	sc.LogLevel, sc.ListenAddress, sc.BaseUrl, sc.PreSharedKey, sc.DummyUrl, sc.WhitelistPath, sc.Expire)
}

func ParseFromCmd() *ServerConfig {
    sc := &ServerConfig{
		Expire: 0,
	}
	_, err := flags.Parse(sc)
	if err != nil {
		log.Fatal(err)
	}
	if sc.PreSharedKey == ""  {
		log.Fatal("You must specify a PreShared Key!")
	}
	if sc.BaseUrl == ""  {
		log.Fatal("You must specify a base url!")
	}
	
	return sc
}

func ParseFromJson(s string, sc *ServerConfig) {
	if item := gjson.Get(s, "loglevel");item.Exists() {
		sc.LogLevel = item.String()
	}	
	if item := gjson.Get(s, "key");item.Exists() {
		sc.PreSharedKey = item.String()
	}
	if item := gjson.Get(s, "listen");item.Exists() {
		sc.ListenAddress = item.String()
	}
	if item := gjson.Get(s, "url");item.Exists() {
		sc.BaseUrl = item.String()
	}		
	if item := gjson.Get(s, "expire");item.Exists() {
		sc.Expire  = int(item.Int())
	}	
	if item := gjson.Get(s, "dummy");item.Exists() {
		sc.DummyUrl = item.String()
	}	
	if item := gjson.Get(s, "whitelist");item.Exists() {
		sc.WhitelistPath = item.String()
	}	
}

func GetGlobalConfig() *ServerConfig {
	return sc
}

func GetPresharedKey() string {
	return sc.PreSharedKey
}

func GetLoglevel() string {
	return sc.LogLevel
}

func GetListenAddress() string {
	return sc.ListenAddress
}

func GetExpire() int {
	return sc.Expire
}

func GetDummyUrl() string {
	return sc.DummyUrl
}

func GetBaseUrl() string {
	return sc.BaseUrl
}

func GetWhitelistPath() string {
	return sc.WhitelistPath
}
