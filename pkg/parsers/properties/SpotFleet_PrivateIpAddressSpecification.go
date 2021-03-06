package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type SpotFleet_PrivateIpAddressSpecification struct {
	Primary          interface{} `yaml:"Primary,omitempty"`
	PrivateIpAddress interface{} `yaml:"PrivateIpAddress"`
}

func (resource SpotFleet_PrivateIpAddressSpecification) Validate() []error {
	errs := []error{}

	if resource.PrivateIpAddress == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PrivateIpAddress'"))
	}
	return errs
}
