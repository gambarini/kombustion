package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type Cluster_Configuration struct {
	Classification          interface{} `yaml:"Classification,omitempty"`
	ConfigurationProperties interface{} `yaml:"ConfigurationProperties,omitempty"`
	Configurations          interface{} `yaml:"Configurations,omitempty"`
}

func (resource Cluster_Configuration) Validate() []error {
	errs := []error{}

	return errs
}
