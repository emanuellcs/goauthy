package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/emanuellcs/goauthy/internal/config"
	"github.com/emanuellcs/goauthy/internal/core/domain"
	"github.com/emanuellcs/goauthy/internal/core/ports"
	"github.com/google/uuid"
)

type OTPService struct {
	cfg      *config.Config
	provider ports.CommunicationProvider
	// repo  ports.OTPRepository (Coming soon with Redis)
}

func NewOTPService(cfg *config.Config, provider ports.CommunicationProvider) *OTPService {
	return &OTPService{
		cfg:      cfg,
		provider: provider,
	}
}

func (s *OTPService) SendOTP(ctx context.Context, to string) (*domain.OTP, error) {
	// Generate Code (Simple random for now)
	code := fmt.Sprintf("%06d", rand.Intn(999999))

	// Create Domain Entity
	// Note: I need to parse duration from config properly later. Using static 5m for now.
	otp := &domain.OTP{
		ID:        uuid.New().String(),
		To:        to,
		Code:      code,
		Method:    s.provider.Name(),
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	// Construct Message
	msg := fmt.Sprintf("Your security code is: %s", code)

	// Use the Provider (Port) to send
	if err := s.provider.Send(ctx, to, msg); err != nil {
		return nil, fmt.Errorf("failed to send otp via %s: %w", s.provider.Name(), err)
	}

	// TODO: Save to Redis (Repo.Save(otp))

	return otp, nil
}

func (s *OTPService) VerifyOTP(ctx context.Context, otpID string, code string) (bool, error) {
	return false, fmt.Errorf("not implemented yet")
}
