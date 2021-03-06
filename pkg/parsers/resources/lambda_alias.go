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

type LambdaAlias struct {
	Type       string                `yaml:"Type"`
	Properties LambdaAliasProperties `yaml:"Properties"`
	Condition  interface{}           `yaml:"Condition,omitempty"`
	Metadata   interface{}           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}           `yaml:"DependsOn,omitempty"`
}

type LambdaAliasProperties struct {
	Description     interface{}                                 `yaml:"Description,omitempty"`
	FunctionName    interface{}                                 `yaml:"FunctionName"`
	FunctionVersion interface{}                                 `yaml:"FunctionVersion"`
	Name            interface{}                                 `yaml:"Name"`
	RoutingConfig   *properties.Alias_AliasRoutingConfiguration `yaml:"RoutingConfig,omitempty"`
}

func NewLambdaAlias(properties LambdaAliasProperties, deps ...interface{}) LambdaAlias {
	return LambdaAlias{
		Type:       "AWS::Lambda::Alias",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseLambdaAlias(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource LambdaAlias
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: LambdaAlias - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource LambdaAlias) Validate() []error {
	return resource.Properties.Validate()
}

func (resource LambdaAliasProperties) Validate() []error {
	errs := []error{}
	if resource.FunctionName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FunctionName'"))
	}
	if resource.FunctionVersion == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FunctionVersion'"))
	}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
