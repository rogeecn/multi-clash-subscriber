package main

import (
	"log"
	"multi-clash-subscriber/config"
	"multi-clash-subscriber/internal/http"
	"os"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("toml")
	viper.SetConfigName("config.local")

	wd, _ := os.Getwd()
	viper.AddConfigPath(wd)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("1.", err)
	}

	var c config.Config
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatal("2.", err)
	}

	if err := http.Serve(&c); err != nil {
		log.Fatal("3.", err)
	}
}
