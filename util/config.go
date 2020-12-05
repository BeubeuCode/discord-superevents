package util

type Confid struct {
	DiscordPublicKey string `mapstructure:"DISCORD_PUBLIC_KEY"`
	DiscordSecretKey string `mapstructure:"DISCORD_SECRET_KEY"`
}
