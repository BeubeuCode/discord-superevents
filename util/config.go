package util

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config data from ENV file
type Config struct {
	DiscordPublicKey string `mapstructure:"DISCORD_PUBLIC_KEY"`
	DiscordSecretKey string `mapstructure:"DISCORD_SECRET_KEY"`
	FirebaseKeyPath  string `mapstructure:"FIREBASE_KEY_PATH"`
}

//LoadConfig loads config file and returns it
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read config  %s", err))
	}

	err = viper.Unmarshal(&config)
	return config, err
}
