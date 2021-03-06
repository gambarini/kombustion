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

type CloudFrontStreamingDistribution struct {
	Type       string                                    `yaml:"Type"`
	Properties CloudFrontStreamingDistributionProperties `yaml:"Properties"`
	Condition  interface{}                               `yaml:"Condition,omitempty"`
	Metadata   interface{}                               `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                               `yaml:"DependsOn,omitempty"`
}

type CloudFrontStreamingDistributionProperties struct {
	StreamingDistributionConfig *properties.StreamingDistribution_StreamingDistributionConfig `yaml:"StreamingDistributionConfig"`
	Tags                        interface{}                                                   `yaml:"Tags"`
}

func NewCloudFrontStreamingDistribution(properties CloudFrontStreamingDistributionProperties, deps ...interface{}) CloudFrontStreamingDistribution {
	return CloudFrontStreamingDistribution{
		Type:       "AWS::CloudFront::StreamingDistribution",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCloudFrontStreamingDistribution(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource CloudFrontStreamingDistribution
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CloudFrontStreamingDistribution - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource CloudFrontStreamingDistribution) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CloudFrontStreamingDistributionProperties) Validate() []error {
	errs := []error{}
	if resource.StreamingDistributionConfig == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StreamingDistributionConfig'"))
	} else {
		errs = append(errs, resource.StreamingDistributionConfig.Validate()...)
	}
	if resource.Tags == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Tags'"))
	}
	return errs
}
