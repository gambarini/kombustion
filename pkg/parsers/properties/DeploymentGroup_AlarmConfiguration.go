package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type DeploymentGroup_AlarmConfiguration struct {
	Enabled                interface{} `yaml:"Enabled,omitempty"`
	IgnorePollAlarmFailure interface{} `yaml:"IgnorePollAlarmFailure,omitempty"`
	Alarms                 interface{} `yaml:"Alarms,omitempty"`
}

func (resource DeploymentGroup_AlarmConfiguration) Validate() []error {
	errs := []error{}

	return errs
}
