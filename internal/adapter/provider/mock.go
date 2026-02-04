package provider

import (
	"context"
	"log/slog"
)

// MockProvider is a fake sender for local development.
type MockProvider struct{}

func NewMockProvider() *MockProvider {
	return &MockProvider{}
}

func (m *MockProvider) Send(ctx context.Context, to string, message string) error {
	// Simulate latency
	// time.Sleep(500 * time.Millisecond)

	slog.Info("MOCK PROVIDER: Sending message",
		"to", to,
		"body", message,
	)
	return nil
}

func (m *MockProvider) Name() string {
	return "mock-console"
}
