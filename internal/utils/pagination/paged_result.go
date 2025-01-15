package pagination

import "math"

type PagedResult[T any] struct {
	Data           []T  `json:"data"`
	HasPrevPage    bool `json:"hasPrevPage"`
	HasNextPage    bool `json:"hasNextPage"`
	TotalCount     int  `json:"totalCount"`
	Count          int  `json:"count"`
	Page           int  `json:"page"`
	TotalPageCount int  `json:"totalPageCount"`
}

func CreatePagedResult[T any](data []T, totalCount, page, pageSize int) PagedResult[T] {
	totalPageCount := int(math.Ceil(float64(totalCount) / float64(pageSize)))
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
