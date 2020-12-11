package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🎉 Hello World!")
		arr := [...]string{"😂", "🐶", "😇", "🍎", "🧚‍", "🔥", "🏜", "💜", "📛", "👌", "👹"}
		idx := rand.Intn(len(arr))
		log.Printf("%s Received a request...", arr[idx])
	})

	srv := http.Server{
		Addr:        ":8080",
		Handler:     mux,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("http server: %v", err)
		}
	}()

	sch := make(chan os.Signal, 1)
	signal.Notify(
		sch,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)
	<-sch
	log.Println("os.Interrupt: shutting down...")

	go func() {
		<-sch
		log.Fatalln("os.Kill: terminating...")
	}()

	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := srv.Shutdown(gracefullCtx); err != nil {
		log.Printf("shutdown error: %v\n", err)
		defer os.Exit(1)
		return
	}
	log.Println("gracefully stopped")

	cancel()
	defer os.Exit(0)
	return
}
