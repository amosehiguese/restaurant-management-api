package log

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var encodeCfg = zapcore.EncoderConfig{
	MessageKey: "msg", 
	NameKey: "name",

	LevelKey: "level",
	EncodeLevel: zapcore.LowercaseLevelEncoder,

	CallerKey: "caller",
	EncodeCaller: zapcore.ShortCallerEncoder,

	TimeKey: "time",
	EncodeTime: zapcore.ISO8601TimeEncoder,
}


type Logger struct {
	Log *zap.Logger
}

func NewLog() *Logger  {
	file, err := os.OpenFile("log/restaurant.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}

	l := Logger{}
	l.Log = zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encodeCfg),
			zapcore.Lock(os.Stdout),
			zapcore.DebugLevel,
		),

	)

	l.Log = l.Log.WithOptions(
		zap.WrapCore(
			func(c zapcore.Core) zapcore.Core {
				return zapcore.NewTee(
					c,
					zapcore.NewCore(
						zapcore.NewJSONEncoder(encodeCfg),
						zapcore.Lock(zapcore.AddSync(file)),
						zapcore.InfoLevel,
					),
				)
			},
		),
	)
	defer func() {
		_ = l.Log.Sync()
	}()

	return &l
}

func (l *Logger) Info(msg string) {
	l.Log.Info(msg)
}
func (l *Logger) Infof(msg string, args ...any) {
	l.Log.Sugar().Infof(msg, args...)
}
func (l *Logger) Infoln(msg ...any) {
	l.Log.Sugar().Infoln(msg...)
}
func (l *Logger) Error(msg string) {
	l.Log.Error(msg)
}
func (l *Logger) Errorf(msg string, args ...any) {
	l.Log.Sugar().Errorf(msg, args...)
}
func (l *Logger) Errorln(msg ...any) {
	l.Log.Sugar().Errorln(msg...)
}
func (l *Logger) Debug(msg string) {
	l.Log.Debug(msg)
}
func (l *Logger) Debugf(msg string, args ...any ) {
	l.Log.Sugar().Debugf(msg, args...)
}
func (l *Logger) Debugln(msg ...any) {
	l.Log.Sugar().Debugln(msg...)
}

