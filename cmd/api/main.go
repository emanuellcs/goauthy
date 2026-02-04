package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/emanuellcs/goauthy/internal/config"
)

func main() {
	// Initialize Logger (JSON format is better for production/Docker)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	logger.Info("Starting GoAuthy Service...")

	// Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	// Verify Configuration (Sanity Check)
	logger.Info("Configuration loaded successfully",
		"env", cfg.Server.Env,
		"port", cfg.Server.Port,
		"strategy_steps", len(cfg.Strategy.Steps),
		"twilio_account", maskString(cfg.Twilio.AccountSID), // Never log full secrets
	)

	// Keep the app alive for now
	fmt.Println("Server is ready to accept connections (mock)...")
}

// Helper to mask sensitive data in logs
func maskString(s string) string {
	if len(s) <= 4 {
		return "****"
	}
	return s[:4] + "****" + s[len(s)-4:]
}