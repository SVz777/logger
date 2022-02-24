// Package log
// file    zap.go
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
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ConsoleZapCore = "console"
	FileZapCore    = "file"

	ConsoleFormatter = "console"
	JsonFormatter    = "json"

	LevelOpLt  = "<"
	LevelOpLte = "<="
	LevelOpGt  = ">"
	LevelOpGte = ">="
)

var Levels = map[string]zapcore.Level{
	"":      zapcore.DebugLevel,
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

func newEncoder(c *OutputConfig) zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "M",
		LevelKey:       "L",
		TimeKey:        "T",
		NameKey:        "N",
		CallerKey:      "C",
		FunctionKey:    "F",
		StacktraceKey:  "S",
		SkipLineEnding: false,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	var encoder zapcore.Encoder
	switch c.Formatter {
	case ConsoleFormatter:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	case JsonFormatter:
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	return encoder
}

func newConsoleCore(c *OutputConfig, levelEnabler zapcore.LevelEnabler) zapcore.Core {
	return zapcore.NewCore(
		newEncoder(c),
		zapcore.Lock(os.Stdout),
		levelEnabler,
	)
}

func newFileCore(c *OutputConfig, levelEnabler zapcore.LevelEnabler) zapcore.Core {
	writer := lumberjack.Logger{
		Filename:   c.WriteConfig.Filename,
		MaxSize:    c.WriteConfig.MaxSize,
		MaxBackups: c.WriteConfig.MaxBackups,
		MaxAge:     c.WriteConfig.MaxAge,
		Compress:   c.WriteConfig.Compress,
		LocalTime:  true,
	}

	// 日志级别
	return zapcore.NewCore(
		newEncoder(c),
		zapcore.AddSync(&writer),
		levelEnabler,
	)
}

func newCore(c *OutputConfig) zapcore.Core {
	var (
		core         zapcore.Core
		levelEnabler zapcore.LevelEnabler
		lvl          = Levels[c.Level]
	)
	switch c.LevelOp {
	case LevelOpLt:
		levelEnabler = zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level < lvl
		})
	case LevelOpLte:
		levelEnabler = zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level <= lvl
		})
	case LevelOpGt:
		levelEnabler = zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level > lvl
		})
	case LevelOpGte:
		levelEnabler = zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level >= lvl
		})
	default:
		levelEnabler = zap.NewAtomicLevelAt(lvl)
	}

	switch c.Writer {
	case ConsoleZapCore:
		core = newConsoleCore(c, levelEnabler)
	case FileZapCore:
		core = newFileCore(c, levelEnabler)
	default:
		core = newConsoleCore(c, levelEnabler)
	}
	return core
}
