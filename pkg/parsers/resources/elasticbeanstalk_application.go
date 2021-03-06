package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type ElasticBeanstalkApplication struct {
	Type       string                                `yaml:"Type"`
	Properties ElasticBeanstalkApplicationProperties `yaml:"Properties"`
	Condition  interface{}                           `yaml:"Condition,omitempty"`
	Metadata   interface{}                           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                           `yaml:"DependsOn,omitempty"`
}

type ElasticBeanstalkApplicationProperties struct {
	ApplicationName         interface{}                                                `yaml:"ApplicationName,omitempty"`
	Description             interface{}                                                `yaml:"Description,omitempty"`
	ResourceLifecycleConfig *properties.Application_ApplicationResourceLifecycleConfig `yaml:"ResourceLifecycleConfig,omitempty"`
}

func NewElasticBeanstalkApplication(properties ElasticBeanstalkApplicationProperties, deps ...interface{}) ElasticBeanstalkApplication {
	return ElasticBeanstalkApplication{
		Type:       "AWS::ElasticBeanstalk::Application",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElasticBeanstalkApplication(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource ElasticBeanstalkApplication
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticBeanstalkApplication - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource ElasticBeanstalkApplication) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElasticBeanstalkApplicationProperties) Validate() []error {
	errs := []error{}
	return errs
}
