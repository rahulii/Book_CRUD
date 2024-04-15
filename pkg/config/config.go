package config

import "github.com/spf13/viper"

type Config struct {
	DBUrl string `mapstructure:"DB_URL"`
	Port string `mapstructure:"PORT"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("app")
	viper.AddConfigPath("./envs")

	viper.SetConfigType("env")

	viper.AutomaticEnv()

	var cfg Config

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}