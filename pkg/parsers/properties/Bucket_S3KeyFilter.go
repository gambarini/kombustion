package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Bucket_S3KeyFilter struct {
	Rules interface{} `yaml:"Rules"`
}

func (resource Bucket_S3KeyFilter) Validate() []error {
	errs := []error{}

	if resource.Rules == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Rules'"))
	}
	return errs
}
