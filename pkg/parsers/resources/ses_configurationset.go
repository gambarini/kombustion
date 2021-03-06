package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type SESConfigurationSet struct {
	Type       string                        `yaml:"Type"`
	Properties SESConfigurationSetProperties `yaml:"Properties"`
	Condition  interface{}                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                   `yaml:"DependsOn,omitempty"`
}

type SESConfigurationSetProperties struct {
	Name interface{} `yaml:"Name,omitempty"`
}

func NewSESConfigurationSet(properties SESConfigurationSetProperties, deps ...interface{}) SESConfigurationSet {
	return SESConfigurationSet{
		Type:       "AWS::SES::ConfigurationSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSESConfigurationSet(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource SESConfigurationSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SESConfigurationSet - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource SESConfigurationSet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SESConfigurationSetProperties) Validate() []error {
	errs := []error{}
	return errs
}
