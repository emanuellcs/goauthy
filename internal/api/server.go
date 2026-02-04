package api

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/emanuellcs/goauthy/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo   *echo.Echo
	config *config.Config
}

// NewServer initializes the Echo server with middleware and configuration
func NewServer(cfg *config.Config) *Server {
	e := echo.New()

	// Hide the Echo banner on startup (we use our own logs)
	e.HideBanner = true
	e.HidePort = true

	// Standard Middleware
	e.Use(middleware.Recover()) // Recovers from panics
	e.Use(middleware.Logger())  // Logs HTTP requests
	e.Use(middleware.CORS())    // Essential for web clients

	s := &Server{
		echo:   e,
		config: cfg,
	}

	s.setupRoutes()

	return s
}

// Start runs the HTTP server on the configured port
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.config.Server.Port)
	slog.Info("HTTP Server listening", "addr", addr, "env", s.config.Server.Env)
	return s.echo.Start(addr)
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	slog.Info("Shutting down HTTP server...")
	return s.echo.Shutdown(ctx)
}
