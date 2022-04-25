package log

import "go.uber.org/zap"

var Logger *zap.Logger
var Sugar *zap.SugaredLogger

func New(debug bool) (*zap.Logger, error) {
	var err error
	if debug {
		Logger, err = zap.NewDevelopment()
	} else {
		Logger, err = zap.NewProduction()
	}
	Sugar = Logger.WithOptions(zap.AddCallerSkip(1)).Sugar()
	if err != nil {
		return nil, err
	}
	return Logger, nil
}

func GetLogger() *zap.Logger {
	if Logger == nil {
		panic("logger is not initialized")
	}
	return Logger
}

func Debug(args ...interface{}) {
	Sugar.Debug(args)
}

func Debugf(template string, args ...interface{}) {
	Sugar.Debugf(template, args...)
}

func Info(args ...interface{}) {
	Sugar.Info(args)
}

func Infof(template string, args ...interface{}) {
	Sugar.Infof(template, args...)
}

func Warn(args ...interface{}) {
	Sugar.Warn(args)
}

func Warnf(template string, args ...interface{}) {
	Sugar.Warnf(template, args...)
}

func Error(args ...interface{}) {
	Sugar.Error(args)
}

func Errorf(template string, args ...interface{}) {
	Sugar.Errorf(template, args...)
}

func Panic(args ...interface{}) {
	Sugar.Panic(args)
}

func Panicf(template string, args ...interface{}) {
	Sugar.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	Sugar.Fatal(args)
}

func Fatalf(template string, args ...interface{}) {
	Sugar.Fatalf(template, args...)
}
