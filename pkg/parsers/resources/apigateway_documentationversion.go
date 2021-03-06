package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type ApiGatewayDocumentationVersion struct {
	Type       string                                   `yaml:"Type"`
	Properties ApiGatewayDocumentationVersionProperties `yaml:"Properties"`
	Condition  interface{}                              `yaml:"Condition,omitempty"`
	Metadata   interface{}                              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                              `yaml:"DependsOn,omitempty"`
}

type ApiGatewayDocumentationVersionProperties struct {
	Description          interface{} `yaml:"Description,omitempty"`
	DocumentationVersion interface{} `yaml:"DocumentationVersion"`
	RestApiId            interface{} `yaml:"RestApiId"`
}

func NewApiGatewayDocumentationVersion(properties ApiGatewayDocumentationVersionProperties, deps ...interface{}) ApiGatewayDocumentationVersion {
	return ApiGatewayDocumentationVersion{
		Type:       "AWS::ApiGateway::DocumentationVersion",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayDocumentationVersion(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource ApiGatewayDocumentationVersion
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayDocumentationVersion - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource ApiGatewayDocumentationVersion) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayDocumentationVersionProperties) Validate() []error {
	errs := []error{}
	if resource.DocumentationVersion == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DocumentationVersion'"))
	}
	if resource.RestApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errs
}
