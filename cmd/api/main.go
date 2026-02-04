package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/emanuellcs/goauthy/internal/adapter/provider"
	"github.com/emanuellcs/goauthy/internal/api"
	"github.com/emanuellcs/goauthy/internal/config"
	"github.com/emanuellcs/goauthy/internal/core/ports"
	"github.com/emanuellcs/goauthy/internal/core/service"
)

func main() {
	// Setup Logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		os.Exit(1)
	}

	// Initialize Adapters (Infrastructure)
	var smsProvider ports.CommunicationProvider = provider.NewMockProvider()

	// Initialize Core Service (Domain Logic)
	otpService := service.NewOTPService(cfg, smsProvider)

	// Initialize HTTP Server (Primary Adapter)
	srv := api.NewServer(cfg, otpService)

	// Start Server in a Goroutine
	// We do this so it doesn't block the main thread, allowing us to listen for signals below.
	go func() {
		if err := srv.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("HTTP server error", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for Interrupt Signal (Graceful Shutdown)
	quit := make(chan os.Signal, 1)
	// SIGINT = Ctrl+C, SIGTERM = Docker stop
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Block here until signal is received
	sig := <-quit
	slog.Info("Shutdown signal received", "signal", sig.String())

	// Execute Shutdown with Timeout
	// We give the server 10 seconds to finish ongoing requests.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
	}

	slog.Info("Server exited properly")
}
