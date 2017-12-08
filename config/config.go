package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Listen string `mapstructure:"listen"`
	Log    *LogConfig
}

type LogConfig struct {
	Level   string
	AppName string
}

func New(path string) (*Config, error) {
	var conf Config

	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
