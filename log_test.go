// Package logger
// file    log_test.go
// @author
//  ___  _  _  ____
// / __)( \/ )(_   )
// \__ \ \  /  / /_
// (___/  \/  (____)
// (903943711@qq.com)
// @date    2022/2/24
// @desc

package logger_test

import (
	"testing"

	"github.com/SVz777/logger"
)

func TestNew(t *testing.T) {
	a := logger.Config{
		{
			Writer:    "console",
			Formatter: "console",
			Level:     "info",
			LevelOp:   ">=",
		},
	}
	l := logger.New(a).WithField("k1", "v1").WithField("k2", "v2")
	l.Tracef("trace")
	l.Debugf("debug")
	l.Infof("info")
	l.Warnf("warn")
	l.Errorf("error")
}
