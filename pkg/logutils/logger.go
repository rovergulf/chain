package logutils

import (
	"errors"
	"go.uber.org/zap"
)

var ErrInvalidLoggerInterface = errors.New("invalid logger interface")

type Logger interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)

	Debugf(message string, args ...any)
	Infof(message string, args ...any)
	Warnf(message string, args ...any)
	Errorf(message string, args ...any)
	Fatalf(message string, args ...any)

	Debugw(message string, keysAndValues ...any)
	Infow(message string, keysAndValues ...any)
	Warnw(message string, keysAndValues ...any)
	Errorw(message string, keysAndValues ...any)
	Fatalw(message string, keysAndValues ...any)
}

func Sugar(logger Logger) (*zap.SugaredLogger, error) {
	zlog, ok := logger.(*zap.SugaredLogger)
	if !ok {
		return nil, ErrInvalidLoggerInterface
	}

	return zlog, nil
}
