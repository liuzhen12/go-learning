package pipefilter

import (
	"errors"
	"fmt"
	"strconv"
)

type ToIntFilter struct {
}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (tf *ToIntFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]string)
	if !ok {
		return nil, errors.New("ToIntFilter wrong format error")
	}
	fmt.Printf("%T\n", parts)
	ret := []int{}
	for _, part := range parts {
		fmt.Println(part)
		s, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}
	return ret, nil
}
