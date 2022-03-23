package config

import (
	"log"

	"github.com/spf13/viper"
)

func Init(path string) {
	viper.SetConfigFile(path)
	viper.AddConfigPath(".")
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
}
func ServicePort() int {
	return viper.GetInt("service.port")
}

func TokenExpired() int {
	return viper.GetInt("jwt.tokenExpired")
}
