package server

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/amosehiguese/restaurant-api/handlers"
	"github.com/amosehiguese/restaurant-api/log"
	"github.com/amosehiguese/restaurant-api/store"
	_ "github.com/amosehiguese/restaurant-api/store"
	_ "github.com/joho/godotenv/autoload"
)



var (
	socketAddr = flag.String("socketaddr", "127.0.0.1:8080", "socket address to listen on")
	cert = flag.String("cert", "", "TLS certificate")
	pkey = flag.String("pkey", "", "TLS private key")
)

func Run() error {
	flag.Parse()

	l := log.NewLog()

	srv := &http.Server{
		Addr: *socketAddr,
		Handler: http.HandlerFunc(handlers.Serve),
		IdleTimeout: 5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}
	
	done := make(chan struct{})

	go func ()  {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)

		for {
			if <-sigChan == os.Interrupt{
				if err := srv.Shutdown(context.Background()); err != nil {
					l.Infof("shutting down... -->%v",err )
				}
				close(done)
				return
			}
		}
	}()

	l.Infof("Serving request over %s\n", srv.Addr)
	store.SetUpDB()
	
	var err error 

	if *cert != "" && *pkey != "" {
		l.Info("TLS enabled")
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


