package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type InstanceFleetConfig_EbsConfiguration struct {
	EbsOptimized          interface{} `yaml:"EbsOptimized,omitempty"`
	EbsBlockDeviceConfigs interface{} `yaml:"EbsBlockDeviceConfigs,omitempty"`
}

func (resource InstanceFleetConfig_EbsConfiguration) Validate() []error {
	errs := []error{}

	return errs
}
