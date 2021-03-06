package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type GuardDutyDetector struct {
	Type       string                      `yaml:"Type"`
	Properties GuardDutyDetectorProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GuardDutyDetectorProperties struct {
	Enable interface{} `yaml:"Enable"`
}

func NewGuardDutyDetector(properties GuardDutyDetectorProperties, deps ...interface{}) GuardDutyDetector {
	return GuardDutyDetector{
		Type:       "AWS::GuardDuty::Detector",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGuardDutyDetector(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource GuardDutyDetector
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GuardDutyDetector - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource GuardDutyDetector) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GuardDutyDetectorProperties) Validate() []error {
	errs := []error{}
	if resource.Enable == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Enable'"))
	}
	return errs
}
