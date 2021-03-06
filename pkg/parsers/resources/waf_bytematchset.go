package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type WAFByteMatchSet struct {
	Type       string                    `yaml:"Type"`
	Properties WAFByteMatchSetProperties `yaml:"Properties"`
	Condition  interface{}               `yaml:"Condition,omitempty"`
	Metadata   interface{}               `yaml:"Metadata,omitempty"`
	DependsOn  interface{}               `yaml:"DependsOn,omitempty"`
}

type WAFByteMatchSetProperties struct {
	Name            interface{} `yaml:"Name"`
	ByteMatchTuples interface{} `yaml:"ByteMatchTuples,omitempty"`
}

func NewWAFByteMatchSet(properties WAFByteMatchSetProperties, deps ...interface{}) WAFByteMatchSet {
	return WAFByteMatchSet{
		Type:       "AWS::WAF::ByteMatchSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseWAFByteMatchSet(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource WAFByteMatchSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFByteMatchSet - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource WAFByteMatchSet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource WAFByteMatchSetProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
