package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/jaisheel15/daos/internal/config"
	"github.com/jaisheel15/daos/internal/logger"
)


func main() {
	cfg , err := config.Load()
	if err != nil {
		fmt.Println("Error loading config:", err)
		panic(err)
	}
	
	log := logger.NewLogger(cfg.LogLevel)

	router := chi.NewRouter()

	router.Get("/health/live" , func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("alive"))
	})

	server := &http.Server{
		Addr:  cfg.HTTPAddr,
		Handler:  router,
		ReadHeaderTimeout: 5 * time.Second,
	}
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	serverErrors := make(chan error, 1)

	go func(){
		log.Info("API server starting", "address", cfg.HTTPAddr)

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	select {
	case <-ctx.Done() :
		log.Info("API server shutting down")
	

	case err := <-serverErrors:
		log.Error("API server error", "error", err)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Error("API server shutdown error", "error", err)
	}

}