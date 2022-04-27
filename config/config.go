package config

import (
	"encoding/json"
	"os"
)

type Db struct {
	Address  string
	DbName   string
	User     string
	Password string
	Port     int
}

type MiniAPP struct {
	AppId     string
	AppSecret string
}

type Configuration struct {
	DB      *Db
	MiniAPP *MiniAPP
}

var ConfAll *Configuration

func LoadConfig() error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	ConfAll = &Configuration{}
	err = decoder.Decode(ConfAll)
	if err != nil {
		return err
	}
	return nil
}
