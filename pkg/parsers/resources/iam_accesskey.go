package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type IAMAccessKey struct {
	Type       string                 `yaml:"Type"`
	Properties IAMAccessKeyProperties `yaml:"Properties"`
	Condition  interface{}            `yaml:"Condition,omitempty"`
	Metadata   interface{}            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}            `yaml:"DependsOn,omitempty"`
}

type IAMAccessKeyProperties struct {
	Serial   interface{} `yaml:"Serial,omitempty"`
	Status   interface{} `yaml:"Status,omitempty"`
	UserName interface{} `yaml:"UserName"`
}

func NewIAMAccessKey(properties IAMAccessKeyProperties, deps ...interface{}) IAMAccessKey {
	return IAMAccessKey{
		Type:       "AWS::IAM::AccessKey",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseIAMAccessKey(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource IAMAccessKey
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IAMAccessKey - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource IAMAccessKey) Validate() []error {
	return resource.Properties.Validate()
}

func (resource IAMAccessKeyProperties) Validate() []error {
	errs := []error{}
	if resource.UserName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'UserName'"))
	}
	return errs
}
