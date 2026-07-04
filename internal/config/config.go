package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	DatabaseURL    string        `env:"DATABASE_URL,required"`
	RedisURL       string        `env:"REDIS_URL,required"`
	JWTAccessSecret  string        `env:"JWT_ACCESS_SECRET,required"`
	JWTRefreshSecret string        `env:"JWT_REFRESH_SECRET,required"`
	JWTAccessTTL     time.Duration `env:"JWT_ACCESS_TTL" envDefault:"15m"`
	JWTRefreshTTL    time.Duration `env:"JWT_REFRESH_TTL" envDefault:"168h"`
	RateLimitAuth    int           `env:"RATE_LIMIT_AUTH" envDefault:"10"`
	InviteTTL        time.Duration `env:"INVITE_TTL" envDefault:"72h"`
	HTTPPort         string        `env:"HTTP_PORT" envDefault:"8080"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	return cfg, nil
}
