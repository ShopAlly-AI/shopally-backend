package usecase

import (
	"context"
	"time"

	"github.com/shopally-ai/pkg/domain"
)

// AlibabaGateway defines the contract for fetching products from an external source.
type AlibabaGateway interface {
	FetchProducts(ctx context.Context, query string, filters map[string]interface{}) ([]*domain.Product, error)
}

// LLMGateway defines the contract for a Large Language Model service
// to parse user intent from a search query.
type LLMGateway interface {
	ParseIntent(ctx context.Context, query string) (map[string]interface{}, error)
}

// CacheGateway defines the contract for a caching service.
type CacheGateway interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}
