package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type ReceiptRule_AddHeaderAction struct {
	HeaderName  interface{} `yaml:"HeaderName"`
	HeaderValue interface{} `yaml:"HeaderValue"`
}

func (resource ReceiptRule_AddHeaderAction) Validate() []error {
	errs := []error{}

	if resource.HeaderName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HeaderName'"))
	}
	if resource.HeaderValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HeaderValue'"))
	}
	return errs
}