package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type AutoScalingGroup_MetricsCollection struct {
	Granularity interface{} `yaml:"Granularity"`
	Metrics     interface{} `yaml:"Metrics,omitempty"`
}

func (resource AutoScalingGroup_MetricsCollection) Validate() []error {
	errs := []error{}

	if resource.Granularity == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Granularity'"))
	}
	return errs
}
