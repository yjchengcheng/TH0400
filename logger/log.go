package logger

import (
	"strings"

	"go.uber.org/zap"
)

// logger application gloable logger instance
var logger *zap.SugaredLogger

func init() {
	// todo it should be as config setting
	// with dev mode or release mode
	cfg := zap.NewDevelopmentConfig()

	if logger == nil {
		if log, err := cfg.Build(); err != nil {
			log.Fatal(err.Error())
		} else {
			logger = log.Sugar()
		}
	}

	// that's end logger init
	// todo
	// maybe log and config init should be together

}

// Debug ...
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf ...
func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

// Debugw ...
func Debugw(msg string, keysAndValues ...interface{}) {
	logger.Debugw(msg, keysAndValues...)
}

// Info ...
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof ...
func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

// Infow ...
func Infow(msg string, keysAndValues ...interface{}) {
	logger.Infow(msg, keysAndValues...)
}

// Error ...
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf ...
func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

// Errorw ...
func Errorw(msg string, keysAndValues ...interface{}) {
	logger.Errorw(msg, keysAndValues...)
}

// XInfof 按 elems = [modName, content, logType, message] 格式传入
func XInfof(elems []string, args ...interface{}) {
	logger.Infof(strings.Join(elems, "; "), args...)
}

// XErrorf 按 elems = [modName, content, logType, message] 格式传入
func XErrorf(elems []string, args ...interface{}) {
	logger.Errorf(strings.Join(elems, ";"), args...)
}

// XDebugf 按 elems = [modName, content, logType, message] 格式传入
func XDebugf(elems []string, args ...interface{}) {
	logger.Debugf(strings.Join(elems, ";"), args...)
}
