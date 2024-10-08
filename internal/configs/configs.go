package configs

import (
	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

type Config struct {
	Postgres struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
	} `yaml:"postgres"`
	Kafka struct {
		Server  string `yaml:"server"`
		Topic   string `yaml:"topic"`
		GroupID string `yaml:"groupId"`
	} `yaml:"kafka"`
}

func InitConfig() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Err(err)
	}
	log.Info().Msg("Using config file 'config.yaml'")
}

func GetConfig() Config {
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Err(err)
	}
	return config
}
