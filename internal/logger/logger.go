package logger

import "go.uber.org/zap"

type Logger interface {
	Infof(msg string, fields ...any)
	Errorf(msg string, fields ...any)
	Fatalf(msg string, fields ...any)
}

type zapLogger struct {
	logger *zap.Logger
}

func NewLogger() (Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return &zapLogger{}, err
	}

	return &zapLogger{logger: logger}, nil
}

func (l *zapLogger) Infof(msg string, fields ...any) {
	l.logger.Sugar().Infof(msg, fields...)
}

func (l *zapLogger) Errorf(msg string, fields ...any) {
	l.logger.Sugar().Infof(msg, fields...)
}

func (l *zapLogger) Fatalf(msg string, fields ...any) {
	l.logger.Sugar().Infof(msg, fields...)
}
