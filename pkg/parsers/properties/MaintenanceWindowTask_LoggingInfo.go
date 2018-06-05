package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type MaintenanceWindowTask_LoggingInfo struct {
	Region   interface{} `yaml:"Region"`
	S3Bucket interface{} `yaml:"S3Bucket"`
	S3Prefix interface{} `yaml:"S3Prefix,omitempty"`
}

func (resource MaintenanceWindowTask_LoggingInfo) Validate() []error {
	errs := []error{}

	if resource.Region == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Region'"))
	}
	if resource.S3Bucket == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'S3Bucket'"))
	}
	return errs
}