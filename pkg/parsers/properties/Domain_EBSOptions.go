package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type Domain_EBSOptions struct {
	EBSEnabled interface{} `yaml:"EBSEnabled,omitempty"`
	Iops       interface{} `yaml:"Iops,omitempty"`
	VolumeSize interface{} `yaml:"VolumeSize,omitempty"`
	VolumeType interface{} `yaml:"VolumeType,omitempty"`
}

func (resource Domain_EBSOptions) Validate() []error {
	errs := []error{}

	return errs
}
