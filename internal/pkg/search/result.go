package search

type Link = string
type Result struct {
	Keyword       string
	ResultStats   int
	NumberOfLinks int
	NumberOfAds   int
	HTMLPage      string
}
