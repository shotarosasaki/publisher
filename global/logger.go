package global

import (
	"fmt"

	"github.com/shotarosasaki/publisher/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TODO 共通リポジトリにロガー作られたら差し替え、ないし、薄いラッパーとして残し、内部で扱うロガーの方を差し替え
var Logger *WrappedLogger

type WrappedLogger struct {
	logger *zap.Logger
	cfg    *config.LogConfig
}

func InitLogger(cfg *config.LogConfig) {
	var logger *zap.Logger
	var err error
	if cfg.Level == "debug" {
		fmt.Println("debug") // TODO 消す！
		logger, err = zap.NewDevelopment(zap.Fields(zap.String("appName", cfg.AppName)))
	} else {
		logger, err = zap.NewProduction(zap.Fields(zap.String("appName", cfg.AppName)))
	}
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	Logger = &WrappedLogger{
		logger: logger,
		cfg:    cfg,
	}
}

// TODO ひとまずzapcoreは隠蔽なしで使ってしまう
func (l *WrappedLogger) Debug(msg string, fields ...zapcore.Field) {
	l.logger.Debug(msg, fields...)
}

// TODO ひとまずzapcoreは隠蔽なしで使ってしまう
func (l *WrappedLogger) Info(msg string, fields ...zapcore.Field) {
	l.logger.Info(msg, fields...)
}

// TODO ひとまずzapcoreは隠蔽なしで使ってしまう
func (l *WrappedLogger) Warn(msg string, fields ...zapcore.Field) {
	l.logger.Warn(msg, fields...)
}

// TODO ひとまずzapcoreは隠蔽なしで使ってしまう
func (l *WrappedLogger) Error(msg string, fields ...zapcore.Field) {
	l.logger.Error(msg, fields...)
}

// TODO ひとまずzapcoreは隠蔽なしで使ってしまう
func (l *WrappedLogger) Fatal(msg string, fields ...zapcore.Field) {
	l.logger.Fatal(msg, fields...)
}
