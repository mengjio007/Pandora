package logs

import "github.com/sirupsen/logrus"

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	logrus.Debugf(template, args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Infof(template string, args ...interface{}) {
	logrus.Infof(template, args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logrus.Warnf(template, args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logrus.Errorf(template, args...)
}

func Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	logrus.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	logrus.Fatalf(template, args...)
}
