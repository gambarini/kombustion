package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type UserPool_PasswordPolicy struct {
	MinimumLength    interface{} `yaml:"MinimumLength,omitempty"`
	RequireLowercase interface{} `yaml:"RequireLowercase,omitempty"`
	RequireNumbers   interface{} `yaml:"RequireNumbers,omitempty"`
	RequireSymbols   interface{} `yaml:"RequireSymbols,omitempty"`
	RequireUppercase interface{} `yaml:"RequireUppercase,omitempty"`
}

func (resource UserPool_PasswordPolicy) Validate() []error {
	errs := []error{}

	return errs
}
