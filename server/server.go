package server

import (
	"context"
	"fmt"
	"go-api-test/server/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	ServiceName = "go-api-test-server"
)

var server *http.Server
var stop chan os.Signal

func init() {
	stop = make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
}

func Start() {
	log.Printf("Starting %s ...", ServiceName)
	defer func() {
		log.Printf("\n****\n%s is down!\n****", ServiceName)
	}()

	initDependencyInjection()
	startHttpServer()

	log.Printf("\n****\n%s started!\n****", ServiceName)

	waitForShutdown()
	handleGracefulShutdown()
}

func startHttpServer() error {

	port := getPort()
	server = &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: routes.GetServerRouter()}
	go func() {
		log.Printf("Starting http server on port %s", port)
		if err := server.ListenAndServe(); err != nil {
			log.Print(fmt.Sprintf("Http server stopped: %s", err))
			forceShutdown()
		}
	}()
	return nil
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	return port

}

func waitForShutdown() {
	<-stop
	log.Print("Interrupt signal received!")
}

func forceShutdown() {
	stop <- os.Interrupt
}

func handleGracefulShutdown() {
	log.Printf("Shutting down %s...", ServiceName)
	server.Shutdown(context.Background())
}
