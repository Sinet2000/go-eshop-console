package pagination

import (
	"errors"
)

// PageQuery is a generic struct for paginated queries.
// - PageIndex must be >= 1
// - PageSize must be > 0
// - Filter can be anything (SQL string, bson.M, or any other interface)
type PageQuery struct {
	PageIndex int64       `json:"pageIndex"`
	PageSize  int64       `json:"pageSize"`
	Filter    interface{} `json:"filter"`
}

func NewPageQuery(pageIndex, pageSize int64, filter interface{}) (*PageQuery, error) {
	if pageIndex < 1 {
		return nil, errors.New("pageIndex must be >= 1")
	}
	if pageSize <= 0 {
		return nil, errors.New("pageSize must be > 0")
	}
	return &PageQuery{
		PageIndex: pageIndex,
		PageSize:  pageSize,
		Filter:    filter,
	}, nil
}
