package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Twilio   TwilioConfig   `mapstructure:"twilio"`
	Strategy StrategyConfig `mapstructure:"strategy"`
	Redis    RedisConfig    `mapstructure:"redis"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type TwilioConfig struct {
	AccountSID   string `mapstructure:"account_sid"`
	AuthToken    string `mapstructure:"auth_token"`
	SenderNumber string `mapstructure:"sender_number"`
}

type StrategyConfig struct {
	OTPLength     int           `mapstructure:"otp_length"`
	OTPExpiration string        `mapstructure:"otp_expiration"` // Parsed as time.Duration later
	Steps         []Step        `mapstructure:"steps"`
}

type Step struct {
	Method  string `mapstructure:"method"`
	Timeout string `mapstructure:"timeout"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// LoadConfig reads configuration from file and environment variables.
func LoadConfig() (*Config, error) {
	v := viper.New()

	// Set default values
	v.SetDefault("server.port", 8080)
	v.SetDefault("server.env", "development")

	// Setup config file (policy.yaml)
	v.SetConfigName("policy")
	v.SetConfigType("yaml")
	v.AddConfigPath(".") // Look in root
	v.AddConfigPath("./config") // Look in config folder

	// Setup Environment Variables
	// Result: TWILIO_AUTH_TOKEN maps to Twilio.AuthToken
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// Read Config
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if we rely only on Env Vars
			fmt.Println("No config file found. Using defaults and environment variables.")
		} else {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &cfg, nil
}