package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) setupRoutes() {
	// Health Check
	s.echo.GET("/health", s.handleHealthCheck)

	// API Group v1
	v1 := s.echo.Group("/v1")

	// Placeholder for future OTP routes
	otp := v1.Group("/otp")
	otp.POST("/send", func(c echo.Context) error {
		return c.JSON(http.StatusNotImplemented, map[string]string{"message": "TODO: Implement Send"})
	})
}

func (s *Server) handleHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"version": "1.0.0",
	})
}
