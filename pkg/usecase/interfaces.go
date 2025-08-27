package usecase

import (
	"context"
	"time"

	"github.com/shopally-ai/pkg/domain"
)

type IFXClient interface {
	GetRate(ctx context.Context, from, to string) (float64, error)
}

type ICachePort interface {
	// Get returns the value, whether it was found, and any error.
	Get(ctx context.Context, key string) (string, bool, error)
	// Set stores the value with a TTL; use 0 for no expiration.
	Set(ctx context.Context, key, val string, ttl time.Duration) error
}

type AlertRepository interface {
	CreateAlert(alert *domain.Alert) error
	GetAlert(alertID string) (*domain.Alert, error)
	DeleteAlert(alertID string) error
}
