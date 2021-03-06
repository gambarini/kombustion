package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Project_VpcConfig struct {
	VpcId            interface{} `yaml:"VpcId"`
	SecurityGroupIds interface{} `yaml:"SecurityGroupIds"`
	Subnets          interface{} `yaml:"Subnets"`
}

func (resource Project_VpcConfig) Validate() []error {
	errs := []error{}

	if resource.VpcId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcId'"))
	}
	if resource.SecurityGroupIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SecurityGroupIds'"))
	}
	if resource.Subnets == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Subnets'"))
	}
	return errs
}
