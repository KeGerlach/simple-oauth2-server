package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	router "simple-oauth2-server/internal/api"
	"syscall"
)

func main() {
	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8080"
	}

	r, err := router.New()
	if err != nil {
		fmt.Printf("Failed to create router. Error: %s", err)
		return
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: r,
	}

	closed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		fmt.Printf("Shutting down server on %v\n", server.Addr)

		ctx := context.TODO()
		if err := server.Shutdown(ctx); err != nil {
			fmt.Printf("Failed to shutdown server. Error: %s", err)
		}

		close(closed)
	}()

	fmt.Printf("Starting server on %v\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Failed to start server. Error: %s", err)
	}

	<-closed
	fmt.Println("Server shutdown successfully")
}
