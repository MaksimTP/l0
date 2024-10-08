package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("config file is not found", err)
	}
	log.Println("Using config file 'config.yaml'")
}

// TODO
func getDbInfo() []string {
	return viper.GetViper().AllKeys()
}

// TODO
func getKafkaInfo() []string {
	return viper.GetViper().AllKeys()
}

func main() {
	initConfig()
	fmt.Println(getDbInfo())
}
