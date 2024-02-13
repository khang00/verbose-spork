package search

import (
	"context"
	"golang.org/x/time/rate"
	"sync"
)

type ProcessStatus = string

const (
	Idle       ProcessStatus = "Idle"
	Done       ProcessStatus = "Done"
	Processing ProcessStatus = "Processing"
)

const (
	DefaultReqPerSec = 5
	DefaultBurst     = 10
)

type RateLimitQuerier struct {
	limiter *rate.Limiter
	querier Querier
	status  ProcessStatus
}

func NewRateLimitQuerier(reqPerSec int, burst int) *RateLimitQuerier {
	limit := rate.Limit(reqPerSec)
	defaultQuerier := NewGoogleSearchQuerier()
	return &RateLimitQuerier{
		limiter: rate.NewLimiter(limit, burst),
		querier: defaultQuerier,
		status:  Idle,
	}
}

func (r *RateLimitQuerier) Search(keywords []string) ([]*Result, error) {
	ctx := context.TODO()
	results := make([]*Result, 0)
	r.status = Processing
	defer func() {
		r.status = Done
	}()

	var wg sync.WaitGroup
	var err error
	for _, keyword := range keywords {
		err = r.limiter.Wait(ctx)

		wg.Add(1)
		go func() {
			defer wg.Done()

			result, err2 := r.querier.Search(keyword)
			if err2 != nil {
				err = err2
			}

			results = append(results, result)
		}()
	}

	wg.Wait()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *RateLimitQuerier) GetStatus() ProcessStatus {
	return r.status
}
