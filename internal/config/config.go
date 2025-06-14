package config

import (
	"encoding/json"
	"os"
)

const (
	configDirName  = "./internal/config/"
	configFileName = ".gatorconfig.json"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	err := cfg.WriteCfgFile()
	if err != nil {
		return err
	}

	return nil
}

func ReadConfig() *Config {
	cfgFile, err := os.Open(configDirName + configFileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := cfgFile.Close()
		if err != nil {
			panic(err)
		}
	}()

	var cfgStruct Config
	decoder := json.NewDecoder(cfgFile)
	err = decoder.Decode(&cfgStruct)
	if err != nil {
		panic(err)
	}
	return &cfgStruct
}

func (cfg *Config) WriteCfgFile() error {
	cfgFile, err := os.OpenFile(configDirName+configFileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer func() {
		err := cfgFile.Close()
		if err != nil {
			panic(err)
		}
	}()

	encoder := json.NewEncoder(cfgFile)
	err = encoder.Encode(cfg)
	if err != nil {
		panic(err)
	}
	return nil
}
