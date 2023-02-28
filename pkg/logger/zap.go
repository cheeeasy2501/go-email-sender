package logger

import "go.uber.org/zap"

type ZapInstance struct {
	i *zap.SugaredLogger
}

func newZapInstance() (ILoggerInstance, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	defer l.Sync()
	sugar := l.Sugar()

	return &ZapInstance{
		i: sugar,
	}, nil
}

func (l *ZapInstance) Error(args ...interface{}) {
	l.i.Errorln(args)
}

func (l *ZapInstance) Warning(args ...interface{}) {
	l.i.Warnln(args)
}

func (l *ZapInstance) Info(args ...interface{}) {
	l.i.Info(args)
}
