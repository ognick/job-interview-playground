package logger

type Config struct {
	Development   bool   `env:"DEVELOPMENT" envDefault:"false"`
	DisableCaller bool   `env:"DISABLE_CALLER" envDefault:"true"`
	DisableJson   bool   `env:"ENCODING" envDefault:"true"`
	Level         string `env:"LEVEL" envDefault:"debug"`
}

type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

type NoopLogger struct{}

func NewNopLogger() *NoopLogger {
	return &NoopLogger{}
}

func (l *NoopLogger) Debug(args ...interface{})                    {}
func (l *NoopLogger) Debugf(template string, args ...interface{})  {}
func (l *NoopLogger) Info(args ...interface{})                     {}
func (l *NoopLogger) Infof(template string, args ...interface{})   {}
func (l *NoopLogger) Warn(args ...interface{})                     {}
func (l *NoopLogger) Warnf(template string, args ...interface{})   {}
func (l *NoopLogger) Error(args ...interface{})                    {}
func (l *NoopLogger) Errorf(template string, args ...interface{})  {}
func (l *NoopLogger) DPanic(args ...interface{})                   {}
func (l *NoopLogger) DPanicf(template string, args ...interface{}) {}
func (l *NoopLogger) Panic(args ...interface{})                    {}
func (l *NoopLogger) Panicf(template string, args ...interface{})  {}
func (l *NoopLogger) Fatal(args ...interface{})                    {}
func (l *NoopLogger) Fatalf(template string, args ...interface{})  {}
