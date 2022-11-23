package entities

import (
	"fmt"
)

type indexPatternAttributes struct {
	Title         string `json:"title"`
	TimeFieldName string `json:"timeFieldName"`
	Fields        string `json:"fields,omitempty"`
}

type indexPattern struct {
	Attributes indexPatternAttributes `json:"attributes"`
}

func NewIndexPattern(indexName string) *indexPattern {
	pattern := indexPattern{}
	pattern.Attributes.Title = fmt.Sprintf("%s-*", indexName)
	pattern.Attributes.TimeFieldName = "@timestamp"
	return &pattern
}
