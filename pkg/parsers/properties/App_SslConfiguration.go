package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type App_SslConfiguration struct {
	Certificate interface{} `yaml:"Certificate,omitempty"`
	Chain       interface{} `yaml:"Chain,omitempty"`
	PrivateKey  interface{} `yaml:"PrivateKey,omitempty"`
}

func (resource App_SslConfiguration) Validate() []error {
	errs := []error{}

	return errs
}
