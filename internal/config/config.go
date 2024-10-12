package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    Port              string `mapstructure:"PORT"`
    DatabaseURL       string `mapstructure:"DATABASE_URL"`
    ExchangeAPIKey    string `mapstructure:"EXCHANGE_API_KEY"`
    ExchangeSecretKey string `mapstructure:"EXCHANGE_SECRET_KEY"`
    LogLevel          string `mapstructure:"LOG_LEVEL"`
}

func Load() (*Config, error) {
    viper.SetConfigFile(".env")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, err
    }

    return &config, nil
}