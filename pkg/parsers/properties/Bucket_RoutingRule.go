package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Bucket_RoutingRule struct {
	RoutingRuleCondition *Bucket_RoutingRuleCondition `yaml:"RoutingRuleCondition,omitempty"`
	RedirectRule         *Bucket_RedirectRule         `yaml:"RedirectRule"`
}

func (resource Bucket_RoutingRule) Validate() []error {
	errs := []error{}

	if resource.RedirectRule == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RedirectRule'"))
	} else {
		errs = append(errs, resource.RedirectRule.Validate()...)
	}
	return errs
}
