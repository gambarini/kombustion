package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type EC2VPCDHCPOptionsAssociation struct {
	Type       string                                 `yaml:"Type"`
	Properties EC2VPCDHCPOptionsAssociationProperties `yaml:"Properties"`
	Condition  interface{}                            `yaml:"Condition,omitempty"`
	Metadata   interface{}                            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                            `yaml:"DependsOn,omitempty"`
}

type EC2VPCDHCPOptionsAssociationProperties struct {
	DhcpOptionsId interface{} `yaml:"DhcpOptionsId"`
	VpcId         interface{} `yaml:"VpcId"`
}

func NewEC2VPCDHCPOptionsAssociation(properties EC2VPCDHCPOptionsAssociationProperties, deps ...interface{}) EC2VPCDHCPOptionsAssociation {
	return EC2VPCDHCPOptionsAssociation{
		Type:       "AWS::EC2::VPCDHCPOptionsAssociation",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2VPCDHCPOptionsAssociation(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource EC2VPCDHCPOptionsAssociation
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2VPCDHCPOptionsAssociation - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource EC2VPCDHCPOptionsAssociation) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2VPCDHCPOptionsAssociationProperties) Validate() []error {
	errs := []error{}
	if resource.DhcpOptionsId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DhcpOptionsId'"))
	}
	if resource.VpcId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcId'"))
	}
	return errs
}
