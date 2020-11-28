package b01

// NewStraightPipeline func.
func NewStraightPipeline(name string, filters ...Filter) *StraightPipeline {
	return &StraightPipeline{
		Name:    name,
		Filters: &filters,
	}
}

// StraightPipeline struct.
type StraightPipeline struct {
	Name    string
	Filters *[]Filter
}

// Process func.
func (sp *StraightPipeline) Process(req Request) (Response, error) {
	var ret interface{}
	var err error
	for _, filter := range *sp.Filters {
		ret, err = filter.Process(req)
		if err != nil {
			return ret, err
		}
		req = ret
	}
	return ret, err
}
