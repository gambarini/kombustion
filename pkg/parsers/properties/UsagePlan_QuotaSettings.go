package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type UsagePlan_QuotaSettings struct {
	Limit  interface{} `yaml:"Limit,omitempty"`
	Offset interface{} `yaml:"Offset,omitempty"`
	Period interface{} `yaml:"Period,omitempty"`
}

func (resource UsagePlan_QuotaSettings) Validate() []error {
	errs := []error{}

	return errs
}
