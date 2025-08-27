package usecase

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
