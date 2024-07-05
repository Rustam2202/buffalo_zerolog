package logger

import (
	"os"

	"github.com/gobuffalo/logger"
	"github.com/rs/zerolog"
)

type ZeroLogger struct {
	zerolog.Logger
}

func NewZeroLogger() logger.FieldLogger {
	// l := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &ZeroLogger{zerolog.New(os.Stdout).With().Timestamp().Logger()}
}

func (c *ZeroLogger) WithField(key string, value interface{}) logger.FieldLogger {
	return &ZeroLogger{}
}

func (c *ZeroLogger) WithFields(fields map[string]interface{}) logger.FieldLogger {
	logCtx := c.Logger.With()
	for k, v := range fields {
		logCtx = logCtx.Interface(k, v)
	}
	return &ZeroLogger{}
}

func (l ZeroLogger) Debugf(string, ...interface{}) {
}
func (l ZeroLogger) Infof(string, ...interface{})  {}
func (l ZeroLogger) Printf(string, ...interface{}) {}
func (l ZeroLogger) Warnf(string, ...interface{})  {}
func (l ZeroLogger) Errorf(string, ...interface{}) {}
func (l ZeroLogger) Fatalf(string, ...interface{}) {}
func (l ZeroLogger) Debug(args ...interface{}) {
	l.Logger.Debug().Msg(args[0].(string))
}
func (l ZeroLogger) Info(...interface{})  {}
func (l ZeroLogger) Warn(...interface{})  {}
func (l ZeroLogger) Error(...interface{}) {}
func (l ZeroLogger) Fatal(...interface{}) {}
func (l ZeroLogger) Panic(...interface{}) {}
