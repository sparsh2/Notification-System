package config

import "gopkg.in/yaml.v3"

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type Config struct {
	DBConfig *DBConfig `yaml:"db"`
}

var Configs *Config

func LoadConfig() {
	// read config from file
	Configs = &Config{}
	Configs.DBConfig = &DBConfig{}
	configFilePath := "/etc/config/config.yaml"
	yaml.Unmarshal([]byte(configFilePath), Configs)
}
