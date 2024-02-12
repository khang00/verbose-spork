package search

import (
	"testing"
)

func TestRateLimitQuerier_Search(t *testing.T) {
	type args struct {
		keywords []string
	}
	tests := []struct {
		name string
		args args
		want []*Result
	}{
		{
			name: "search a list of keywords",
			args: args{keywords: []string{
				"airpods",
				"aquafina",
				"levi",
			}},
		},
	}

	testRPS := 1
	testBurst := 1
	r := NewRateLimitQuerier(testRPS, testBurst)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, err := r.Search(tt.args.keywords)
			if err != nil {
				t.Errorf("Search() got error: %s", err)
				return
			}

			if len(results) != len(tt.args.keywords) {
				t.Errorf("Search() size of results not equal to size of keywords: %d want %d",
					len(results), len(tt.args.keywords))
			}

			for _, result := range results {
				if !isValid(result) {
					t.Errorf("Search() result is empty: %v", result)
				}
			}
		})
	}
}
