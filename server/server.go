package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	socketAddr = flag.String("socketaddr", "127.0.0.1:8080", "socket address to listen on")
	cert = flag.String("cert", "", "TLS certificate")
	pkey = flag.String("pkey", "", "TLS private key")
)

func Run() error {
	flag.Parse()

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr: *socketAddr,
		Handler: mux,
		IdleTimeout: 5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}

	mux.HandleFunc("/test", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintf(w, "Hello world")
	})

	done := make(chan struct{})

	go func ()  {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)

		for {
			if <-sigChan == os.Interrupt{
				if err := srv.Shutdown(context.Background()); err != nil {
					log.Printf("shutting down... -->%v",err )
				}
				close(done)
				return
			}
		}
	}()

	log.Printf("Serving request over %s\n", srv.Addr)
	
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
