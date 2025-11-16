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
)

func main() {
	server := &http.Server{
		Addr: ":3000",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 4)
		w.Write([]byte("Hello world"))
	})

	go func() {
		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			log.Fatalf("could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	fmt.Println("shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("could not shutdown server: %v\n", err)
	}
	fmt.Println("server stopped")
}
