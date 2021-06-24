package pipefilter

import "fmt"

type StraightPipeline struct {
	Name    string
	Filters *[]Filter
}

func NewStraightPipeline(name string, filters ...Filter) *StraightPipeline {
	return &StraightPipeline{
		Name:    name,
		Filters: &filters,
	}
}

func (f *StraightPipeline) Process(data Request) (Response, error) {
	var (
		ret Response
		err error
	)

	for _, filter := range *f.Filters {
		fmt.Printf("%T\n", filter)
		if ret, err = filter.Process(data); err != nil {
			return ret, err
		}
		fmt.Println(ret)
		data = ret
	}
	return ret, nil
}
