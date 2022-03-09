package yourlogger

import (
	"log"
	"strings"
)

type Logger struct {
	fields []string
}

func New() *Logger {
	return &Logger{
		fields: make([]string, 0),
	}
}

func (m *Logger) WithField(k string, v string) *Logger {
	m.fields = append(m.fields, k+"~~"+v)
	return m
}

func (m *Logger) Info(msg string) {
	log.Printf("%s : %s", strings.Join(m.fields, ","), msg)
}
