package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Alias_AliasRoutingConfiguration struct {
	AdditionalVersionWeights interface{} `yaml:"AdditionalVersionWeights"`
}

func (resource Alias_AliasRoutingConfiguration) Validate() []error {
	errs := []error{}

	if resource.AdditionalVersionWeights == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AdditionalVersionWeights'"))
	}
	return errs
}
