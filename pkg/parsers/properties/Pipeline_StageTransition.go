package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Pipeline_StageTransition struct {
	Reason    interface{} `yaml:"Reason"`
	StageName interface{} `yaml:"StageName"`
}

func (resource Pipeline_StageTransition) Validate() []error {
	errs := []error{}

	if resource.Reason == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Reason'"))
	}
	if resource.StageName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StageName'"))
	}
	return errs
}
