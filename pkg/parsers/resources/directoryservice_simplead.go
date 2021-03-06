package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type DirectoryServiceSimpleAD struct {
	Type       string                             `yaml:"Type"`
	Properties DirectoryServiceSimpleADProperties `yaml:"Properties"`
	Condition  interface{}                        `yaml:"Condition,omitempty"`
	Metadata   interface{}                        `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                        `yaml:"DependsOn,omitempty"`
}

type DirectoryServiceSimpleADProperties struct {
	CreateAlias interface{}                      `yaml:"CreateAlias,omitempty"`
	Description interface{}                      `yaml:"Description,omitempty"`
	EnableSso   interface{}                      `yaml:"EnableSso,omitempty"`
	Name        interface{}                      `yaml:"Name"`
	Password    interface{}                      `yaml:"Password"`
	ShortName   interface{}                      `yaml:"ShortName,omitempty"`
	Size        interface{}                      `yaml:"Size"`
	VpcSettings *properties.SimpleAD_VpcSettings `yaml:"VpcSettings"`
}

func NewDirectoryServiceSimpleAD(properties DirectoryServiceSimpleADProperties, deps ...interface{}) DirectoryServiceSimpleAD {
	return DirectoryServiceSimpleAD{
		Type:       "AWS::DirectoryService::SimpleAD",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseDirectoryServiceSimpleAD(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource DirectoryServiceSimpleAD
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DirectoryServiceSimpleAD - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource DirectoryServiceSimpleAD) Validate() []error {
	return resource.Properties.Validate()
}

func (resource DirectoryServiceSimpleADProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Password == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Password'"))
	}
	if resource.Size == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Size'"))
	}
	if resource.VpcSettings == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcSettings'"))
	} else {
		errs = append(errs, resource.VpcSettings.Validate()...)
	}
	return errs
}
