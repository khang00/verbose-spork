package handler

import (
	"context"
	"encoding/json"
	"github.com/khang00/verbose-spork/internal/handler/auth"
	"github.com/khang00/verbose-spork/internal/model"
	"github.com/khang00/verbose-spork/internal/pkg/search"
	"net/http"
)

type KeywordStore interface {
	CreateKeywords(keywords []*model.Keyword) ([]*model.Keyword, error)
}

type KeywordHandler struct {
	querier search.BatchQuerier
}

type UploadKeywordsRequest struct {
	Keywords []string `json:"keywords"`
}

type UploadKeywordsResponse struct {
	Results []*Result `json:"results"`
}

type Result struct {
	ID      string `json:"id"`
	Keyword string `json:"keyword"`
}

type ResultDetails struct {
	Result
	ResultStats   int    `json:"result_stats"`
	NumberOfLinks int    `json:"number_of_links"`
	NumberOfAds   int    `json:"number_of_ads"`
	HTML          string `json:"html"`
}

func NewKeywordHandler() *KeywordHandler {
	return &KeywordHandler{
		querier: search.NewRateLimitQuerier(search.DefaultReqPerSec, search.DefaultBurst),
	}
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
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (s *KeywordHandler) uploadKeywords(ctx context.Context, req *UploadKeywordsRequest) (*UploadKeywordsResponse, error) {
	results, err := s.querier.Search(req.Keywords)
	if err != nil {
		return nil, err
	}

	resultsResp := make([]*Result, 0)
	for _, result := range results {
		resultsResp = append(resultsResp, &Result{
			Keyword: result.Keyword,
		})
	}

	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	keywords := make([]*model.Keyword, 0)
	for _, result := range results {
		keywords = append(keywords, &model.Keyword{
			Keyword:       result.Keyword,
			ResultStats:   result.ResultStats,
			NumberOfLinks: result.NumberOfLinks,
			NumberOfAds:   result.NumberOfAds,
			HTML:          result.HTMLPage,
			UserID:        userID,
		})
	}

	return &UploadKeywordsResponse{Results: resultsResp}, nil
}

/*&Result{
Keyword:       result.Keyword,
ResultStats:   result.ResultStats,
NumberOfLinks: result.NumberOfLinks,
NumberOfAds:   result.NumberOfAds,
HTML:          result.HTMLPage,
})*/
