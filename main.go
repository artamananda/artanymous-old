package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	logger := zerolog.New(os.Stdout)
	logger.Level(zerolog.InfoLevel)
}
