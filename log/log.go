package log

import (
	"log"
	"os"
	"runtime"

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


type logger struct {
	Log *zap.Logger
}

func NewLog() *logger  {
	file, err := os.OpenFile("restaurant.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}

	l := logger{}
	l.Log = zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encodeCfg),
			zapcore.Lock(os.Stdout),
			zapcore.DebugLevel,
		),
		zap.AddCaller(),
		zap.Fields(
			zap.String("version", runtime.Version()),
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