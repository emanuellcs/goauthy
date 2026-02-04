package ports

import (
	"context"

	"github.com/emanuellcs/goauthy/internal/core/domain"
)

// OTPService defines the input port (Use Cases).
// The HTTP Handler will call these methods.
type OTPService interface {
	SendOTP(ctx context.Context, to string) (*domain.OTP, error)
	VerifyOTP(ctx context.Context, otpID string, code string) (bool, error)
}

// CommunicationProvider defines the output port (Adapter).
// Twilio, WhatsApp, or ConsoleMock must implement this.
type CommunicationProvider interface {
	Send(ctx context.Context, to string, message string) error
	Name() string // e.g., "twilio-sms", "mock"
}

// OTPRepository defines how we store state (Redis/Postgres).
// I'll implement this later, but defining it now is good practice.
type OTPRepository interface {
	Save(ctx context.Context, otp *domain.OTP) error
	Get(ctx context.Context, otpID string) (*domain.OTP, error)
}
