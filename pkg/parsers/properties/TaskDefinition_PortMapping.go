package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type TaskDefinition_PortMapping struct {
	ContainerPort interface{} `yaml:"ContainerPort,omitempty"`
	HostPort      interface{} `yaml:"HostPort,omitempty"`
	Protocol      interface{} `yaml:"Protocol,omitempty"`
}

func (resource TaskDefinition_PortMapping) Validate() []error {
	errs := []error{}

	return errs
}
