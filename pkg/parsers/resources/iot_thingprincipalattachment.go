package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type IoTThingPrincipalAttachment struct {
	Type       string                                `yaml:"Type"`
	Properties IoTThingPrincipalAttachmentProperties `yaml:"Properties"`
	Condition  interface{}                           `yaml:"Condition,omitempty"`
	Metadata   interface{}                           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                           `yaml:"DependsOn,omitempty"`
}

type IoTThingPrincipalAttachmentProperties struct {
	Principal interface{} `yaml:"Principal"`
	ThingName interface{} `yaml:"ThingName"`
}

func NewIoTThingPrincipalAttachment(properties IoTThingPrincipalAttachmentProperties, deps ...interface{}) IoTThingPrincipalAttachment {
	return IoTThingPrincipalAttachment{
		Type:       "AWS::IoT::ThingPrincipalAttachment",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseIoTThingPrincipalAttachment(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource IoTThingPrincipalAttachment
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IoTThingPrincipalAttachment - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource IoTThingPrincipalAttachment) Validate() []error {
	return resource.Properties.Validate()
}

func (resource IoTThingPrincipalAttachmentProperties) Validate() []error {
	errs := []error{}
	if resource.Principal == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Principal'"))
	}
	if resource.ThingName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ThingName'"))
	}
	return errs
}