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

type DynamoDBTable struct {
	Type       string                  `yaml:"Type"`
	Properties DynamoDBTableProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

type DynamoDBTableProperties struct {
	TableName               interface{}                               `yaml:"TableName,omitempty"`
	TimeToLiveSpecification *properties.Table_TimeToLiveSpecification `yaml:"TimeToLiveSpecification,omitempty"`
	StreamSpecification     *properties.Table_StreamSpecification     `yaml:"StreamSpecification,omitempty"`
	SSESpecification        *properties.Table_SSESpecification        `yaml:"SSESpecification,omitempty"`
	ProvisionedThroughput   *properties.Table_ProvisionedThroughput   `yaml:"ProvisionedThroughput"`
	GlobalSecondaryIndexes  interface{}                               `yaml:"GlobalSecondaryIndexes,omitempty"`
	LocalSecondaryIndexes   interface{}                               `yaml:"LocalSecondaryIndexes,omitempty"`
	AttributeDefinitions    interface{}                               `yaml:"AttributeDefinitions,omitempty"`
	Tags                    interface{}                               `yaml:"Tags,omitempty"`
	KeySchema               interface{}                               `yaml:"KeySchema"`
}

func NewDynamoDBTable(properties DynamoDBTableProperties, deps ...interface{}) DynamoDBTable {
	return DynamoDBTable{
		Type:       "AWS::DynamoDB::Table",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseDynamoDBTable(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource DynamoDBTable
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DynamoDBTable - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource DynamoDBTable) Validate() []error {
	return resource.Properties.Validate()
}

func (resource DynamoDBTableProperties) Validate() []error {
	errs := []error{}
	if resource.ProvisionedThroughput == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ProvisionedThroughput'"))
	} else {
		errs = append(errs, resource.ProvisionedThroughput.Validate()...)
	}
	if resource.KeySchema == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'KeySchema'"))
	}
	return errs
}
