package pagination

import "math"

type PagedResult[T any] struct {
	Data           []T   `json:"data"`
	HasPrevPage    bool  `json:"hasPrevPage"`
	HasNextPage    bool  `json:"hasNextPage"`
	TotalCount     int64 `json:"totalCount"`
	Count          int   `json:"count"`
	Page           int64 `json:"page"`
	TotalPageCount int64 `json:"totalPageCount"`
}

func CreatePagedResult[T any](data []T, totalCount, page, pageSize int64) PagedResult[T] {
	totalPageCount := int64(math.Ceil(float64(totalCount) / float64(pageSize)))
	hasPrevPage := page > 1
	hasNextPage := page < totalPageCount

	return PagedResult[T]{
		Data:           data,
		HasPrevPage:    hasPrevPage,
		HasNextPage:    hasNextPage,
		TotalCount:     totalCount,
		Count:          len(data),
		Page:           page,
		TotalPageCount: totalPageCount,
	}
}
