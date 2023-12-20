package server

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/amosehiguese/restaurant-api/routes"
)



var (
	socketAddr = flag.String("socketaddr", "127.0.0.1:8080", "socket address to listen on")
	cert = flag.String("cert", "", "TLS certificate")
	pkey = flag.String("pkey", "", "TLS private key")
)

func Run() error {
	flag.Parse()

	mux := http.NewServeMux()
	l := NewLog()

	srv := &http.Server{
		Addr: *socketAddr,
		Handler: mux,
		IdleTimeout: 5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}


	mux.HandleFunc("/menu/", routes.HandleMenuRequest)

	
	done := make(chan struct{})

	go func ()  {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)

		for {
			if <-sigChan == os.Interrupt{
				if err := srv.Shutdown(context.Background()); err != nil {
					l.log.Sugar().Infof("shutting down... -->%v",err )
				}
				close(done)
				return
			}
		}
	}()

	l.log.Sugar().Infof("Serving request over %s\n", srv.Addr)
	
	var err error 

	if *cert != "" && *pkey != "" {
		log.Println("TLS enabled")
		err = srv.ListenAndServeTLS(*cert, *pkey)
	} else {
		err = srv.ListenAndServe()
	}

	if err == http.ErrServerClosed {
		err = nil
	}

	<-done
	return err
}


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
	log *zap.Logger
}

func NewLog() *logger  {
	file, err := os.OpenFile("restaurant.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}

	l := logger{}
	l.log = zap.New(
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

	l.log = l.log.WithOptions(
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
		_ = l.log.Sync()
	}()

	return &l
}