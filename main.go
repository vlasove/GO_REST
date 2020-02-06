package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

//Config ...
type Config struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBname   string `toml:"dbname"`
	SSLmode  string `toml:"sslmode"`
	Port     string `toml:"bind_port"`
}

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "config/app.toml", "path to config .toml file")
}

func main() {
	flag.Parse()
	var conf Config
	_, err := toml.DecodeFile(configPath, &conf)
	if err != nil {
		log.Fatal(err)
	}

	app := App{}

	app.Initialize(conf.User, conf.Password, conf.DBname, conf.SSLmode)

	app.Run(conf.Port)

}
