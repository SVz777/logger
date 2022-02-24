// Package log
// file    log.go
// @author
//  ___  _  _  ____
// / __)( \/ )(_   )
// \__ \ \  /  / /_
// (___/  \/  (____)
// (903943711@qq.com)
// @date    2022/2/24
// @desc
package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	// Trace logs to TRACE log. Arguments are handled in the manner of fmt.Print.
	Trace(args ...interface{})
	// Tracef logs to TRACE log. Arguments are handled in the manner of fmt.Printf.
	Tracef(format string, args ...interface{})
	// Debug logs to DEBUG log. Arguments are handled in the manner of fmt.Print.
	Debug(args ...interface{})
	// Debugf logs to DEBUG log. Arguments are handled in the manner of fmt.Printf.
	Debugf(format string, args ...interface{})
	// Info logs to INFO log. Arguments are handled in the manner of fmt.Print.
	Info(args ...interface{})
	// Infof logs to INFO log. Arguments are handled in the manner of fmt.Printf.
	Infof(format string, args ...interface{})
	// Warn logs to WARNING log. Arguments are handled in the manner of fmt.Print.
	Warn(args ...interface{})
	// Warnf logs to WARNING log. Arguments are handled in the manner of fmt.Printf.
	Warnf(format string, args ...interface{})
	// Error logs to ERROR log. Arguments are handled in the manner of fmt.Print.
	Error(args ...interface{})
	// Errorf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
	Errorf(format string, args ...interface{})
	// Fatal logs to ERROR log. Arguments are handled in the manner of fmt.Print.
	// that all Fatal logs will exit with os.Exit(1).
	// Implementations may also call os.Exit() with a non-zero exit code.
	Fatal(args ...interface{})
	// Fatalf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
	Fatalf(format string, args ...interface{})
	// WithField 自定义嵌入字段
	WithField(field ...string) Logger

	Sync() error
}

var defaultConfig = []OutputConfig{
	{
		Writer:    "console",
		Level:     "debug",
		Formatter: "console",
	},
}
var DefaultLogger Logger

func init() {
	DefaultLogger = New(defaultConfig)
}

type log struct {
	l *zap.Logger
}

func (l *log) Trace(args ...interface{}) {
	if l.l.Core().Enabled(zapcore.DebugLevel) {
		l.l.Debug(fmt.Sprint(args...))
	}
}

func (l *log) Tracef(format string, args ...interface{}) {
	if l.l.Core().Enabled(zapcore.DebugLevel) {
		l.l.Debug(fmt.Sprintf(format, args...))
	}
}

func (l *log) Debug(args ...interface{}) {
	if l.l.Core().Enabled(zapcore.DebugLevel) {
		l.l.Debug(fmt.Sprint(args...))
	}
}

func (l *log) Debugf(format string, args ...interface{}) {
	if l.l.Core().Enabled(zapcore.DebugLevel) {
		l.l.Debug(fmt.Sprintf(format, args...))
	}
}

func (l *log) Info(args ...interface{}) {
	if l.l.Core().Enabled(zapcore.InfoLevel) {
		l.l.Info(fmt.Sprint(args...))
	}
}

func (l *log) Infof(format string, args ...interface{}) {
	if l.l.Core().Enabled(zapcore.InfoLevel) {
		l.l.Info(fmt.Sprintf(format, args...))
	}
}

func (l *log) Warn(args ...interface{}) {
	if l.l.Core().Enabled(zapcore.WarnLevel) {
		l.l.Warn(fmt.Sprint(args...))
	}
}
func (l *log) Warnf(format string, args ...interface{}) {
	if l.l.Core().Enabled(zapcore.WarnLevel) {
		l.l.Warn(fmt.Sprintf(format, args...))
	}
}

func (l *log) Error(args ...interface{}) {
	if l.l.Core().Enabled(zapcore.ErrorLevel) {
		l.l.Error(fmt.Sprint(args...))
	}
}

func (l *log) Errorf(format string, args ...interface{}) {
	if l.l.Core().Enabled(zapcore.ErrorLevel) {
		l.l.Error(fmt.Sprintf(format, args...))
	}
}

func (l *log) Fatal(args ...interface{}) {
	if l.l.Core().Enabled(zapcore.FatalLevel) {
		l.l.Fatal(fmt.Sprint(args...))
	}
}

func (l *log) Fatalf(format string, args ...interface{}) {
	if l.l.Core().Enabled(zapcore.FatalLevel) {
		l.l.Fatal(fmt.Sprintf(format, args...))
	}
}

func (l *log) WithField(fields ...string) Logger {
	zapFields := make([]zap.Field, len(fields)/2)
	for index := range zapFields {
		zapFields[index] = zap.String(fields[2*index], fields[2*index+1])
	}

	return &log{l: l.l.With(zapFields...)}
}

func (l *log) Sync() error {
	return l.l.Sync()
}

func New(c Config) Logger {
	cores := make([]zapcore.Core, 0, len(c))
	for _, o := range c {
		cores = append(cores, newCore(&o))
	}
	l := &log{
		l: zap.New(
			zapcore.NewTee(cores...),
			zap.AddCallerSkip(2),
			zap.AddCaller(),
		),
	}

	return l
}
