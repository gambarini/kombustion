package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type WorkSpacesWorkspace struct {
	Type       string                        `yaml:"Type"`
	Properties WorkSpacesWorkspaceProperties `yaml:"Properties"`
	Condition  interface{}                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                   `yaml:"DependsOn,omitempty"`
}

type WorkSpacesWorkspaceProperties struct {
	BundleId                    interface{} `yaml:"BundleId"`
	DirectoryId                 interface{} `yaml:"DirectoryId"`
	RootVolumeEncryptionEnabled interface{} `yaml:"RootVolumeEncryptionEnabled,omitempty"`
	UserName                    interface{} `yaml:"UserName"`
	UserVolumeEncryptionEnabled interface{} `yaml:"UserVolumeEncryptionEnabled,omitempty"`
	VolumeEncryptionKey         interface{} `yaml:"VolumeEncryptionKey,omitempty"`
}

func NewWorkSpacesWorkspace(properties WorkSpacesWorkspaceProperties, deps ...interface{}) WorkSpacesWorkspace {
	return WorkSpacesWorkspace{
		Type:       "AWS::WorkSpaces::Workspace",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseWorkSpacesWorkspace(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource WorkSpacesWorkspace
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WorkSpacesWorkspace - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource WorkSpacesWorkspace) Validate() []error {
	return resource.Properties.Validate()
}

func (resource WorkSpacesWorkspaceProperties) Validate() []error {
	errs := []error{}
	if resource.BundleId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BundleId'"))
	}
	if resource.DirectoryId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DirectoryId'"))
	}
	if resource.UserName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'UserName'"))
	}
	return errs
}
