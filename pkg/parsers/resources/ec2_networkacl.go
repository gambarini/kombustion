package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type EC2NetworkAcl struct {
	Type       string                  `yaml:"Type"`
	Properties EC2NetworkAclProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

type EC2NetworkAclProperties struct {
	VpcId interface{} `yaml:"VpcId"`
	Tags  interface{} `yaml:"Tags,omitempty"`
}

func NewEC2NetworkAcl(properties EC2NetworkAclProperties, deps ...interface{}) EC2NetworkAcl {
	return EC2NetworkAcl{
		Type:       "AWS::EC2::NetworkAcl",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2NetworkAcl(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource EC2NetworkAcl
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2NetworkAcl - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource EC2NetworkAcl) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2NetworkAclProperties) Validate() []error {
	errs := []error{}
	if resource.VpcId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcId'"))
	}
	return errs
}
