package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var Cfg *ini.File

func init() {
	var err error
	Cfg, err = ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	Cfg.BlockMode = false
}
