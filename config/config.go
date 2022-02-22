package config

import "github.com/spf13/viper"

var config *viper.Viper

func InitConfig()  {
	config = viper.New()
	config.SetConfigName("config")
	config.AddConfigPath("./config")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		return
	}
	config.WatchConfig()
	return
}

func GetConfig() *viper.Viper {
	return config
}

