package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Service_HealthCheckConfig struct {
	FailureThreshold interface{} `yaml:"FailureThreshold,omitempty"`
	ResourcePath     interface{} `yaml:"ResourcePath,omitempty"`
	Type             interface{} `yaml:"Type"`
}

func (resource Service_HealthCheckConfig) Validate() []error {
	errs := []error{}

	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
