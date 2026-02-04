package api

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) setupRoutes() {
	// Health Check
	s.echo.GET("/health", s.handleHealthCheck)

	// API Group v1
	v1 := s.echo.Group("/v1")

	// OTP Routes
	otp := v1.Group("/otp")
	otp.POST("/send", s.handleSendOTP)
}

func (s *Server) handleSendOTP(c echo.Context) error {
	// Define Request DTO
	type Request struct {
		To string `json:"to"`
	}

	var req Request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	// Call Service
	// Pass the request context specifically
	otp, err := s.otpService.SendOTP(c.Request().Context(), req.To)
	if err != nil {
		slog.Error("Failed to send OTP", "error", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal error"})
	}

	// Response
	// WARNING: returning 'debug_code' is for development only.
	return c.JSON(http.StatusOK, map[string]string{
		"message":    "OTP Sent",
		"otp_id":     otp.ID,
		"debug_code": otp.Code,
	})
}

func (s *Server) handleHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"version": "1.0.0",
	})
}
