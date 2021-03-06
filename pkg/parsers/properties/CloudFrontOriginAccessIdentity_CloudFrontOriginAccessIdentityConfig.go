package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type CloudFrontOriginAccessIdentity_CloudFrontOriginAccessIdentityConfig struct {
	Comment interface{} `yaml:"Comment"`
}

func (resource CloudFrontOriginAccessIdentity_CloudFrontOriginAccessIdentityConfig) Validate() []error {
	errs := []error{}

	if resource.Comment == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Comment'"))
	}
	return errs
}
