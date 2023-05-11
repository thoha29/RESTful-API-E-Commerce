package config

import "github.com/tkanos/gonfig"

type Configurations struct {
	DB_HOST string
	DB_PORT string
	DB_NAME string
	DB_USER string
	DB_PASS string
}

func GetConfig() Configurations {
	conf := Configurations{}
	gonfig.GetConf("config/config.json", &conf)
	return conf
}
