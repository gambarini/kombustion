package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type Filter_FindingCriteria struct {
	Criterion interface{}       `yaml:"Criterion,omitempty"`
	ItemType  *Filter_Condition `yaml:"ItemType,omitempty"`
}

func (resource Filter_FindingCriteria) Validate() []error {
	errs := []error{}

	return errs
}
