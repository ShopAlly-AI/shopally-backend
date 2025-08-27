package usecase

import (
	"context"
)

// SearchProductsUseCase contains the business logic for searching products.
// It orchestrates calls to external gateways (LLM, Alibaba, Cache).
type SearchProductsUseCase struct {
	alibabaGateway AlibabaGateway
	llmGateway     LLMGateway
	cacheGateway   CacheGateway
}

// NewSearchProductsUseCase creates a new SearchProductsUseCase.
func NewSearchProductsUseCase(ag AlibabaGateway, lg LLMGateway, cg CacheGateway) *SearchProductsUseCase {
	return &SearchProductsUseCase{
		alibabaGateway: ag,
		llmGateway:     lg,
		cacheGateway:   cg,
	}
}

// Search runs the mocked search pipeline: Parse -> Fetch (using intent as filters).
func (uc *SearchProductsUseCase) Search(ctx context.Context, query string) (interface{}, error) {
	// Parse intent via LLM
	intent, err := uc.llmGateway.ParseIntent(ctx, query)
	if err != nil {
		// For V1 mock, fail soft by using empty filters
		intent = map[string]interface{}{}
	}

	// Fetch products from the gateway
	products, err := uc.alibabaGateway.FetchProducts(ctx, query, intent)
	if err != nil {
		return nil, err
	}
	// Return the envelope-compatible data payload
	return map[string]interface{}{"products": products}, nil
}
