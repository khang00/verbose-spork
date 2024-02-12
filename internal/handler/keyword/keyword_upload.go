package keyword

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/khang00/verbose-spork/internal/handler/auth"
	"github.com/khang00/verbose-spork/internal/model"
	"net/http"
)

type UploadKeywordsRequest struct {
	Keywords []string `json:"keywords"`
}

type UploadKeywordsResponse struct {
	Results []*Result `json:"results"`
}

type Result struct {
	ID      uint   `json:"id"`
	Keyword string `json:"keyword"`
}

type ResultDetails struct {
	Result
	ResultStats   int    `json:"result_stats"`
	NumberOfLinks int    `json:"number_of_links"`
	NumberOfAds   int    `json:"number_of_ads"`
	HTML          string `json:"html"`
}

func (s *KeywordHandler) UploadKeywords(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	req := &UploadKeywordsRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := s.uploadKeywords(r.Context(), req)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Printf("%s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (s *KeywordHandler) uploadKeywords(ctx context.Context, req *UploadKeywordsRequest) (*UploadKeywordsResponse, error) {
	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	results, err := s.querier.Search(req.Keywords)
	if err != nil {
		return nil, err
	}

	keywordModels := make([]*model.Keyword, 0)
	for _, result := range results {
		keywordModels = append(keywordModels, &model.Keyword{
			Keyword:       result.Keyword,
			ResultStats:   result.ResultStats,
			NumberOfLinks: result.NumberOfLinks,
			NumberOfAds:   result.NumberOfAds,
			HTML:          result.HTMLPage,
			UserID:        userID,
		})
	}

	keywords, err := s.keywordStore.CreateKeywords(keywordModels)
	if err != nil {
		return nil, err
	}

	resultsResp := make([]*Result, 0)
	for _, keyword := range keywords {
		resultsResp = append(resultsResp, &Result{
			ID:      keyword.ID,
			Keyword: keyword.Keyword,
		})
	}

	return &UploadKeywordsResponse{Results: resultsResp}, nil
}
