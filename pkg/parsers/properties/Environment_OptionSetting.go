package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Environment_OptionSetting struct {
	Namespace    interface{} `yaml:"Namespace"`
	OptionName   interface{} `yaml:"OptionName"`
	ResourceName interface{} `yaml:"ResourceName,omitempty"`
	Value        interface{} `yaml:"Value,omitempty"`
}

func (resource Environment_OptionSetting) Validate() []error {
	errs := []error{}

	if resource.Namespace == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Namespace'"))
	}
	if resource.OptionName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'OptionName'"))
	}
	return errs
}
