package logging

import (
	"encoding/json"
	"io"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	log *logrus.Entry
}

func (l *Logger) BaseLog(fields map[string]interface{}, data interface{}) (log *logrus.Entry) {
	file, function, line := GetCaller()
	log = l.log.WithFields(logrus.Fields{
		"file":     file,
		"line":     line,
		"function": function,
	}).WithFields(fields)
	if data != nil {
		payload, err := json.Marshal(data)
		if err != nil {
			log = log.WithField("data_error", err.Error())
			return
		}
		log = log.WithField("data", string(payload))
	}
	return
}

func (l *Logger) Info(fields map[string]interface{}, data interface{}, message string) {
	l.BaseLog(fields, data).Info(message)
}
func (l *Logger) Warn(fields map[string]interface{}, data interface{}, message string) {
	l.BaseLog(fields, data).Warn(message)
}
func (l *Logger) Error(fields map[string]interface{}, data interface{}, message string) {
	l.BaseLog(fields, data).Error(message)
}
func (l *Logger) Fatal(fields map[string]interface{}, data interface{}, message string) {
	l.BaseLog(fields, data).Fatal(message)
}
func (l *Logger) Panic(fields map[string]interface{}, data interface{}, message string) {
	l.BaseLog(fields, data).Panic(message)
}

func (l *Logger) SetFileOutput(file io.Writer) {
	l.log.Logger.Out = file
}

func NewLogger(name string) (log *Logger) {
	l := logrus.New()
	log = &Logger{
		log: l.WithField("service", name),
	}
	return
}
