package config

import (
	"github.com/spf13/viper"
)

// TODO インフラレイヤーの切り替えを考慮して設定ファイルもインフラ関連を分けることを検討！

type Config struct {
	Listen string `mapstructure:"listen"`
	Log    *LogConfig
	Queue  *QueueConfig
}

type LogConfig struct {
	Level   string
	AppName string
}

type QueueConfig struct {
	CredentialsPath string
	ProjectID       string
	TopicName       string
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
