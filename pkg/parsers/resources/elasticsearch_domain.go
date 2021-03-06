package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type ElasticsearchDomain struct {
	Type       string                        `yaml:"Type"`
	Properties ElasticsearchDomainProperties `yaml:"Properties"`
	Condition  interface{}                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                   `yaml:"DependsOn,omitempty"`
}

type ElasticsearchDomainProperties struct {
	AccessPolicies             interface{}                                   `yaml:"AccessPolicies,omitempty"`
	DomainName                 interface{}                                   `yaml:"DomainName,omitempty"`
	ElasticsearchVersion       interface{}                                   `yaml:"ElasticsearchVersion,omitempty"`
	VPCOptions                 *properties.Domain_VPCOptions                 `yaml:"VPCOptions,omitempty"`
	SnapshotOptions            *properties.Domain_SnapshotOptions            `yaml:"SnapshotOptions,omitempty"`
	AdvancedOptions            interface{}                                   `yaml:"AdvancedOptions,omitempty"`
	Tags                       interface{}                                   `yaml:"Tags,omitempty"`
	ElasticsearchClusterConfig *properties.Domain_ElasticsearchClusterConfig `yaml:"ElasticsearchClusterConfig,omitempty"`
	EBSOptions                 *properties.Domain_EBSOptions                 `yaml:"EBSOptions,omitempty"`
}

func NewElasticsearchDomain(properties ElasticsearchDomainProperties, deps ...interface{}) ElasticsearchDomain {
	return ElasticsearchDomain{
		Type:       "AWS::Elasticsearch::Domain",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElasticsearchDomain(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource ElasticsearchDomain
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticsearchDomain - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource ElasticsearchDomain) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElasticsearchDomainProperties) Validate() []error {
	errs := []error{}
	return errs
}
