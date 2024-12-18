package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/tinkoff/invest-api-go-sdk/investgo"
)

type Config struct {
	Tinkoff  investgo.Config
	BotToken string `yaml:"TgBotToken"`
}

func New() (*Config, error) {
	cfg := Config{}

	err := cleanenv.ReadConfig("../../config.yaml", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
