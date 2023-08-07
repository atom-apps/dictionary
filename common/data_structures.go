package common

import (
	"strings"

	"golang.org/x/exp/constraints"
)

type SortQueryFilter struct {
	Asc  *string `json:"asc" form:"asc"`
	Desc *string `json:"desc" form:"desc"`
}

func NewSortQueryFilter() *SortQueryFilter {
	return &SortQueryFilter{}
}

func (s *SortQueryFilter) AscFields() []string {
	if s.Asc == nil {
		return nil
	}
	return strings.Split(*s.Asc, ",")
}

func (s *SortQueryFilter) DescFields() []string {
	if s.Desc == nil {
		return nil
	}
	return strings.Split(*s.Desc, ",")
}

type PageDataResponse struct {
	PageQueryFilter `json:",inline"`
	Total           int64       `json:"total"`
	Items           interface{} `json:"items"`
}

type PageQueryFilter struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

func NewPageQueryFilter[T constraints.Integer](page, limit T) *PageQueryFilter {
	return &PageQueryFilter{
		Page:  int(page),
		Limit: int(limit),
	}
}

func (filter *PageQueryFilter) Offset() int {
	return (filter.Page - 1) * filter.Limit
}

func (filter *PageQueryFilter) Format() *PageQueryFilter {
	if filter.Page <= 0 {
		filter.Page = 1
	}

	if filter.Limit <= 0 {
		filter.Limit = 10
	}

	if filter.Limit > 50 {
		filter.Limit = 50
	}
	return filter
}
