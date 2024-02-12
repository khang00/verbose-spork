package search

type Querier interface {
	Search(keyword string) (*Result, error)
}

type BatchQuerier interface {
	Search(keyword []string) ([]*Result, error)
	GetStatus() ProcessStatus
}
