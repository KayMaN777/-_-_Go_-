package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/version", VersionHandler)
	mux.HandleFunc("/decode", DecodeHandler)
	mux.HandleFunc("/hard-op", HardOpHandler)

	srv := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	go func() {
		log.Println("Starting server on " + port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", port, err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
