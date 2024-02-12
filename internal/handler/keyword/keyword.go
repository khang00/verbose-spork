package keyword

import (
	"github.com/khang00/verbose-spork/internal/handler"
	"github.com/khang00/verbose-spork/internal/pkg/search"
)

type KeywordHandler struct {
	querier      search.BatchQuerier
	keywordStore handler.KeywordStore
}

func NewKeywordHandler(keywordStore handler.KeywordStore) *KeywordHandler {
	return &KeywordHandler{
		querier:      search.NewRateLimitQuerier(search.DefaultReqPerSec, search.DefaultBurst),
		keywordStore: keywordStore,
	}
}
