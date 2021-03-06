package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type MetricFilter_MetricTransformation struct {
	MetricName      interface{} `yaml:"MetricName"`
	MetricNamespace interface{} `yaml:"MetricNamespace"`
	MetricValue     interface{} `yaml:"MetricValue"`
}

func (resource MetricFilter_MetricTransformation) Validate() []error {
	errs := []error{}

	if resource.MetricName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricName'"))
	}
	if resource.MetricNamespace == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricNamespace'"))
	}
	if resource.MetricValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricValue'"))
	}
	return errs
}
