package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type ApplicationReferenceDataSource_RecordFormat struct {
	RecordFormatType  interface{}                                       `yaml:"RecordFormatType"`
	MappingParameters *ApplicationReferenceDataSource_MappingParameters `yaml:"MappingParameters,omitempty"`
}

func (resource ApplicationReferenceDataSource_RecordFormat) Validate() []error {
	errs := []error{}

	if resource.RecordFormatType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RecordFormatType'"))
	}
	return errs
}
