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
			want: "search result is not empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := r.Search(tt.args.keyword)
			if got.ResultStats == 0 || got.NumberOfLinks == 0 || got.HTMLPage == "" {
				t.Errorf("Search() got = %v, want %s", got, tt.want)
			}
		})
	}
}
