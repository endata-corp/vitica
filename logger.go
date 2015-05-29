package vitica

import (
	"vitica/_vendor/src/github.com/chakrit/go-bunyan"
)

var log = bunyan.NewStdLogger("vitica", bunyan.StdoutSink())

func Include(info bunyan.Info) bunyan.Log {
	return log.Include(info)
}

func Record(key string, value interface{}) bunyan.Log {
	return log.Record(key, value)
}

func Recordf(key, value string, args ...interface{}) bunyan.Log {
	return log.Recordf(key, value, args...)
}

func Child() bunyan.Log {
	return log.Child()
}

func Trace(msg string, args ...interface{}) {
	log.Tracef(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	log.Debugf(msg, args...)
}

func Info(msg string, args ...interface{}) {
	log.Infof(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	log.Warnf(msg, args...)
}

func Error(msg string, args ...interface{}) {
	log.Errorf(msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	log.Fatalf(msg, args)
}
