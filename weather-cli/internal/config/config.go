package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

type Config struct {
	Env    string       `yaml:"env" env-default:"dev"`
	Client ConfigClient `yaml:"client"`
}

type ConfigClient struct {
	Geo     ConfigGeo     `yaml:"geo"`
	Weather ConfigWeather `yaml:"weather"`
	Timeout time.Duration `yaml:"timeout" env-default:"5s"`
}

type ConfigGeo struct {
	URL      string `yaml:"geo_url"`
	Language string `yaml:"language"`
}

type ConfigWeather struct {
	URL          string `yaml:"weather_url"`
	Hourly       string `yaml:"hourly"`
	ForecastDays int    `yaml:"forecast_days"`
}

func MustLoad() *Config {

	configPath := getConfigPath()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		zap.L().Fatal("config file not found", zap.Error(err))
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		zap.L().Fatal("failed to read config", zap.Error(err))
	}

	return &cfg
}

func getConfigPath() string {
	envType := os.Getenv("WEATHER_CLI_ENV")
	switch envType {
	case "dev", "prod":
		return "config/" + envType + "/" + envType + ".yaml"
	default:
		return "config/prod/prod.yaml"
	}
}
