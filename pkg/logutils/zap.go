package logutils

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() (*zap.SugaredLogger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.Development = viper.GetString("env") != "main"
	cfg.DisableStacktrace = !viper.GetBool("log_stacktrace")

	if viper.GetBool("log_json") {
		cfg.Encoding = "json"
	} else {
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	l, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return l.Sugar(), nil
}
