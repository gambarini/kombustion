package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type InstanceFleetConfig_SpotProvisioningSpecification struct {
	BlockDurationMinutes   interface{} `yaml:"BlockDurationMinutes,omitempty"`
	TimeoutAction          interface{} `yaml:"TimeoutAction"`
	TimeoutDurationMinutes interface{} `yaml:"TimeoutDurationMinutes"`
}

func (resource InstanceFleetConfig_SpotProvisioningSpecification) Validate() []error {
	errs := []error{}

	if resource.TimeoutAction == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TimeoutAction'"))
	}
	if resource.TimeoutDurationMinutes == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TimeoutDurationMinutes'"))
	}
	return errs
}
