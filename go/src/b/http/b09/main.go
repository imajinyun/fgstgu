package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func serve(ctx context.Context) (err error) {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "OK")
		},
	))

	srv := &http.Server{Addr: ":8080", Handler: mux}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %+s\n", err)
		}
	}()
	log.Printf("Server started")
	<-ctx.Done()
	log.Printf("Server stopped")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("Server shutdown failed: %+s", err)
	}
	log.Printf("Server exited properly")
	if err == http.ErrServerClosed {
		err = nil
	}
	return
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("System call: %+v\n", oscall)
		cancel()
	}()

	if err := serve(ctx); err != nil {
		log.Printf("Failed to serve: %v\n", err)
	}
}
