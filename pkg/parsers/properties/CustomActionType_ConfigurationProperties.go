package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type CustomActionType_ConfigurationProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	Key         interface{} `yaml:"Key"`
	Name        interface{} `yaml:"Name"`
	Queryable   interface{} `yaml:"Queryable,omitempty"`
	Required    interface{} `yaml:"Required"`
	Secret      interface{} `yaml:"Secret"`
	Type        interface{} `yaml:"Type,omitempty"`
}

func (resource CustomActionType_ConfigurationProperties) Validate() []error {
	errs := []error{}

	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Required == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Required'"))
	}
	if resource.Secret == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Secret'"))
	}
	return errs
}
