package log

import (
	"ha/pkg/config"
	"ha/pkg/service"
)

type logManager struct {
	service.MangerConfig
	Level string
}

func NewLogManager(config config.Configuration) service.LogManager {
	logManager := &logManager{}
	if config.Debug {
		logManager.Level = Debug
	}
	return logManager
}

func (l *logManager) SetUp() {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()
	if l.Initialized {
		return
	}
	_ = SetFormat(TextFormat)
	if l.Level == Debug {
		_ = SetLevel(Debug)
	}
	l.Initialized = true
}

func (l *logManager) Shutdown() {
	//TODO implement me
	panic("implement me")
}
