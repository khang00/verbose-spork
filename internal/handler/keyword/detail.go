package keyword

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type KeywordDetailsRequest struct {
	ID uint `json:"id"`
}

type KeywordDetailsResponse struct {
	Detail
}

type Detail struct {
	Result
	ResultStats   int    `json:"result_stats"`
	NumberOfLinks int    `json:"number_of_links"`
	NumberOfAds   int    `json:"number_of_ads"`
	HTML          string `json:"html"`
}

func (s *KeywordHandler) GetKeywordsDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	req := &KeywordDetailsRequest{}
	err := decoder.Decode(req, r.URL.Query())
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := s.getKeywordsDetail(r.Context(), req)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Printf("%s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (s *KeywordHandler) getKeywordsDetail(ctx context.Context, req *KeywordDetailsRequest) (*KeywordDetailsResponse, error) {
	keyword, err := s.keywordStore.GetKeywordByID(req.ID)
	if err != nil {
		return nil, err
	}

	return &KeywordDetailsResponse{Detail{
		Result: Result{
			ID:      keyword.ID,
			Keyword: keyword.Keyword,
		},
		ResultStats:   keyword.ResultStats,
		NumberOfLinks: keyword.NumberOfLinks,
		NumberOfAds:   keyword.NumberOfAds,
		HTML:          keyword.HTML,
	}}, nil
}
