package util

import "github.com/spf13/viper"

// Config data from ENV file
type Config struct {
	DiscordPublicKey string `mapstructure:"DISCORD_PUBLIC_KEY"`
	DiscordSecretKey string `mapstructure:"DISCORD_SECRET_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
}
