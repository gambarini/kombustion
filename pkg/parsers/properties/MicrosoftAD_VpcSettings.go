package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type MicrosoftAD_VpcSettings struct {
	VpcId     interface{} `yaml:"VpcId"`
	SubnetIds interface{} `yaml:"SubnetIds"`
}

func (resource MicrosoftAD_VpcSettings) Validate() []error {
	errs := []error{}

	if resource.VpcId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcId'"))
	}
	if resource.SubnetIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetIds'"))
	}
	return errs
}
