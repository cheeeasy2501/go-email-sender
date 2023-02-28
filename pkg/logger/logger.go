package logger

import (
	"errors"
	"fmt"
)

type LoggerType int

// реализации логгеров
const (
	// ZapDevelopment LoggerType = "ZapDevelopment"
	// ZapProduction  LoggerType = "ZapProduction"
	ZapDevelopment LoggerType = iota
	ZapProduction
)

func (s LoggerType) String() string {
	switch s {
	case ZapDevelopment:
		return "ZapDevelopment"
	case ZapProduction:
		return "ZapProduction"
	}

	return "unknown"
}

// имплементить интерфейс для разных пакетов
type ILoggerInstance interface {
	Error(args ...interface{})
	Warning(args ...interface{})
}

// интерфейс для обертки вокруг логгеров
type ILogger interface {
	Instance() ILoggerInstance
}

// фабрика логгеров
func NewLoggerInstance(t LoggerType) (ILoggerInstance, error) {
	switch t {
	case ZapDevelopment:
		fmt.Println("ZapDev logger initializing ...")
		l, err := newZapInstance()
		if err != nil {
			return nil, err
		}
		fmt.Println("ZapDev logger initialized")

		return l, nil
	case ZapProduction:
		fmt.Println("ZapProd logger initializing ...")
		l, err := newZapInstance()
		if err != nil {
			return nil, err
		}
		fmt.Println("ZapProd logger initialized")

		return l, nil
	default:
		return nil, errors.New("Logger not implemented!")
	}
}

type Logger struct {
	instance ILoggerInstance
}

func (l *Logger) Instance() ILoggerInstance {
	return l.instance
}

func NewLogger(l ILoggerInstance) ILogger {
	return &Logger{
		instance: l,
	}
}
