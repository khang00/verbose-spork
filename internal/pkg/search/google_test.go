package search

import (
	"testing"
)

func TestGoogleSearchQuerier_Search(t *testing.T) {
	r := NewGoogleSearchQuerier()
	type args struct {
		keyword string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "search for Aquafina",
			args: args{keyword: "Airpods"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.Search(tt.args.keyword)
			if err != nil {
				t.Errorf("Search() got error: %s", err)
			}

			if !isValid(got) {
				t.Errorf("Search() result is empty: %v", got)
			}
		})
	}
}

func isValid(result *Result) bool {
	return result.Keyword != "" && result.ResultStats > 0 && result.NumberOfLinks > 0 && result.HTMLPage != ""
}
