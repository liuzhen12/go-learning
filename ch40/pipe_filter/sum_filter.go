package pipefilter

import "errors"

type SumFilter struct{}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Response, error) {

	parts, ok := data.([]int)
	if !ok {
		return nil, errors.New("SumFilter wrong format error")
	}
	var ret int
	for _, part := range parts {
		ret += part
	}
	return ret, nil
}
