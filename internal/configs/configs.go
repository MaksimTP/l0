package configs

import (
	"log"

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
		log.Fatalln("config file is not found", err)
	}
	log.Println("Using config file 'config.yaml'")
}

func GetConfig() Config {
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	return config
}
