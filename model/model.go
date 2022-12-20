package model

import (
	"gopkg.in/yaml.v3"
	"os"
	"task_rest/middleware"
)

type Body struct {
	Encrypt string `json:"encrypt"`
	Decrypt string `json:"decrypt"`
}

type Config struct {
	Api struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"api"`

	Sql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
		SslMode  string `yaml:"sslmode"`
	} `yaml:"sql"`
}

var (
	// ConfigFile - values read from "config.yml"
	ConfigFile Config
)

func init() {
	middleware.Logs.Debug().Msgf("[model] init started")
	b, err := os.ReadFile("config.yml")
	if err != nil {
		middleware.Logs.Err(err).Msgf("error reading config file")
	} else {
		middleware.Logs.Info().Msgf("reading config file is success")
	}
	err = yaml.Unmarshal(b, &ConfigFile)
	if err != nil {
		middleware.Logs.Err(err).Msgf("error unmarshal config file")
	} else {
		middleware.Logs.Info().Msgf("unmarshal config file is success")
	}
	middleware.Logs.Debug().Msgf("[model] init finished")
}
