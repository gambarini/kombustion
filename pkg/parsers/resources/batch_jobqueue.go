package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type BatchJobQueue struct {
	Type       string                  `yaml:"Type"`
	Properties BatchJobQueueProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

type BatchJobQueueProperties struct {
	JobQueueName            interface{} `yaml:"JobQueueName,omitempty"`
	Priority                interface{} `yaml:"Priority"`
	State                   interface{} `yaml:"State,omitempty"`
	ComputeEnvironmentOrder interface{} `yaml:"ComputeEnvironmentOrder"`
}

func NewBatchJobQueue(properties BatchJobQueueProperties, deps ...interface{}) BatchJobQueue {
	return BatchJobQueue{
		Type:       "AWS::Batch::JobQueue",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseBatchJobQueue(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource BatchJobQueue
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: BatchJobQueue - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource BatchJobQueue) Validate() []error {
	return resource.Properties.Validate()
}

func (resource BatchJobQueueProperties) Validate() []error {
	errs := []error{}
	if resource.Priority == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Priority'"))
	}
	if resource.ComputeEnvironmentOrder == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ComputeEnvironmentOrder'"))
	}
	return errs
}
