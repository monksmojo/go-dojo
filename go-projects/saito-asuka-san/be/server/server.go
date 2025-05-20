package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/monksmojo/go-dojo/go-projects/saito-asuka-san/router"
)

func Run() {
	server := newServer()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		println()
		log.Println("Shutting down server...")

		err := gracefulShutdown(server, 25*time.Second)

		if err != nil {
			log.Printf("Server stopped: %s", err.Error())
		}

		os.Exit(0)
	}()
}

func newServer() *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%v", os.Getenv("PORT")),
		Handler:      router.RegisterRouter(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func gracefulShutdown(server *http.Server, maximumTime time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), maximumTime)
	defer cancel()
	return server.Shutdown(ctx)
}
