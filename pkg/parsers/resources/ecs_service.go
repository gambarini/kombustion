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

type ECSService struct {
	Type       string               `yaml:"Type"`
	Properties ECSServiceProperties `yaml:"Properties"`
	Condition  interface{}          `yaml:"Condition,omitempty"`
	Metadata   interface{}          `yaml:"Metadata,omitempty"`
	DependsOn  interface{}          `yaml:"DependsOn,omitempty"`
}

type ECSServiceProperties struct {
	Cluster                       interface{}                                 `yaml:"Cluster,omitempty"`
	DesiredCount                  interface{}                                 `yaml:"DesiredCount,omitempty"`
	HealthCheckGracePeriodSeconds interface{}                                 `yaml:"HealthCheckGracePeriodSeconds,omitempty"`
	LaunchType                    interface{}                                 `yaml:"LaunchType,omitempty"`
	PlatformVersion               interface{}                                 `yaml:"PlatformVersion,omitempty"`
	Role                          interface{}                                 `yaml:"Role,omitempty"`
	ServiceName                   interface{}                                 `yaml:"ServiceName,omitempty"`
	TaskDefinition                interface{}                                 `yaml:"TaskDefinition"`
	NetworkConfiguration          *properties.Service_NetworkConfiguration    `yaml:"NetworkConfiguration,omitempty"`
	LoadBalancers                 interface{}                                 `yaml:"LoadBalancers,omitempty"`
	PlacementConstraints          interface{}                                 `yaml:"PlacementConstraints,omitempty"`
	PlacementStrategies           interface{}                                 `yaml:"PlacementStrategies,omitempty"`
	DeploymentConfiguration       *properties.Service_DeploymentConfiguration `yaml:"DeploymentConfiguration,omitempty"`
}

func NewECSService(properties ECSServiceProperties, deps ...interface{}) ECSService {
	return ECSService{
		Type:       "AWS::ECS::Service",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseECSService(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource ECSService
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ECSService - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource ECSService) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ECSServiceProperties) Validate() []error {
	errs := []error{}
	if resource.TaskDefinition == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TaskDefinition'"))
	}
	return errs
}
