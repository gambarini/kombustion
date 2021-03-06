package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type EC2VPCGatewayAttachment struct {
	Type       string                            `yaml:"Type"`
	Properties EC2VPCGatewayAttachmentProperties `yaml:"Properties"`
	Condition  interface{}                       `yaml:"Condition,omitempty"`
	Metadata   interface{}                       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                       `yaml:"DependsOn,omitempty"`
}

type EC2VPCGatewayAttachmentProperties struct {
	InternetGatewayId interface{} `yaml:"InternetGatewayId,omitempty"`
	VpcId             interface{} `yaml:"VpcId"`
	VpnGatewayId      interface{} `yaml:"VpnGatewayId,omitempty"`
}

func NewEC2VPCGatewayAttachment(properties EC2VPCGatewayAttachmentProperties, deps ...interface{}) EC2VPCGatewayAttachment {
	return EC2VPCGatewayAttachment{
		Type:       "AWS::EC2::VPCGatewayAttachment",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2VPCGatewayAttachment(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource EC2VPCGatewayAttachment
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2VPCGatewayAttachment - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource EC2VPCGatewayAttachment) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2VPCGatewayAttachmentProperties) Validate() []error {
	errs := []error{}
	if resource.VpcId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcId'"))
	}
	return errs
}
