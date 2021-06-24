package pipefilter

import (
	"errors"
	"fmt"
	"strings"
)

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter: delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, errors.New("SplitFilter wrong format error")
	}
	parts := strings.Split(str, sf.delimiter)
	fmt.Println(len(parts))
	return parts, nil
}
