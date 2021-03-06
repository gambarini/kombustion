package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type LogsLogGroup struct {
	Type       string                 `yaml:"Type"`
	Properties LogsLogGroupProperties `yaml:"Properties"`
	Condition  interface{}            `yaml:"Condition,omitempty"`
	Metadata   interface{}            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}            `yaml:"DependsOn,omitempty"`
}

type LogsLogGroupProperties struct {
	LogGroupName    interface{} `yaml:"LogGroupName,omitempty"`
	RetentionInDays interface{} `yaml:"RetentionInDays,omitempty"`
}

func NewLogsLogGroup(properties LogsLogGroupProperties, deps ...interface{}) LogsLogGroup {
	return LogsLogGroup{
		Type:       "AWS::Logs::LogGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseLogsLogGroup(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource LogsLogGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: LogsLogGroup - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource LogsLogGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource LogsLogGroupProperties) Validate() []error {
	errs := []error{}
	return errs
}
