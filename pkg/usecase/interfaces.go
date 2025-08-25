package usecase

import (
	"context"
	"time"
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
