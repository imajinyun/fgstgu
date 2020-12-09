package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

// MessageHandler struct.
type MessageHandler struct {
	ctx  context.Context
	name string
}

func (m *MessageHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Write([]byte(fmt.Sprintf("🎉 Message from %s", m.name)))
}

func newMessageServer(
	ctx context.Context,
	name string,
	port int,
) *http.Server {

	mux := http.NewServeMux()
	handler := &MessageHandler{ctx: ctx, name: name}
	mux.Handle("/", handler)
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	return httpServer
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(quit)
	g, ctx := errgroup.WithContext(ctx)

	server1 := newMessageServer(ctx, "server1", 8080)
	g.Go(func() error {
		log.Println("💜 Server 1 listening on port 8080")
		if err := server1.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	server2 := newMessageServer(ctx, "server2", 8081)
	g.Go(func() error {
		log.Println("💚 Server 2 listening on port 8081")
		if err := server2.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	select {
	case <-quit:
		break
	case <-ctx.Done():
		break
	}
	cancel()

	timeoutCtx, timeoutCancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer timeoutCancel()

	log.Println("🔥 Shutting down servers, please wait...")

	server1.Shutdown(timeoutCtx)
	server2.Shutdown(timeoutCtx)

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
	log.Println("😂 Bye")
}
