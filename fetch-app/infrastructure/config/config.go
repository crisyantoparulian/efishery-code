package config

import (
	"encoding/json"
	"os"
)

type (
	Config struct {
		Host              string            `json:"host"`
		Database          Database          `json:"database"`
		EfisheryApi       EfisheryAPI       `json:"efishery_api"`
		CurrencyConverter CurrencyConverter `json:"currency_converter"`
	}

	Database struct {
		Mysql MysqlDB `json:"mysql"`
	}

	MysqlDB struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
		SSLMode  bool   `json:"ssl_mode"`
		LogMode  bool   `json:"log_mode"`
	}
	EfisheryAPI struct {
		FetchResources string `json:"fetch_resources"`
	}

	CurrencyConverter struct {
		Url    string `json:"url"`
		ApiKey string `json:"api_key"`
	}
)

func ReadConfig() (*Config, error) {

	bytes, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var cfg Config

	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
