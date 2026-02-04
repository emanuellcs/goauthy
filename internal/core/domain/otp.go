package domain

import "time"

// OTP represents the core entity of a One-Time Password.
type OTP struct {
	ID        string
	To        string
	Code      string
	Method    string // sms, whatsapp, voice
	CreatedAt time.Time
	ExpiresAt time.Time
}

// IsExpired checks if the OTP is no longer valid.
func (o *OTP) IsExpired() bool {
	return time.Now().After(o.ExpiresAt)
}
