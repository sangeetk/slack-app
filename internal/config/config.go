package config

import "os"

type Config struct {
	SlackBotToken      string
	SlackSigningSecret string
	SlackAppToken      string
	DatabaseURL        string
}

func LoadConfig() (*Config, error) {
	return &Config{
		SlackBotToken:      os.Getenv("SLACK_BOT_TOKEN"),
		SlackSigningSecret: os.Getenv("SLACK_SIGNING_SECRET"),
		SlackAppToken:      os.Getenv("SLACK_APP_TOKEN"),
		DatabaseURL:        os.Getenv("DATABASE_URL"),
	}, nil
}
