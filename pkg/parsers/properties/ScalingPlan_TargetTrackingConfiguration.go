package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type ScalingPlan_TargetTrackingConfiguration struct {
	DisableScaleIn                       interface{}                                       `yaml:"DisableScaleIn,omitempty"`
	EstimatedInstanceWarmup              interface{}                                       `yaml:"EstimatedInstanceWarmup,omitempty"`
	ScaleInCooldown                      interface{}                                       `yaml:"ScaleInCooldown,omitempty"`
	ScaleOutCooldown                     interface{}                                       `yaml:"ScaleOutCooldown,omitempty"`
	TargetValue                          interface{}                                       `yaml:"TargetValue"`
	PredefinedScalingMetricSpecification *ScalingPlan_PredefinedScalingMetricSpecification `yaml:"PredefinedScalingMetricSpecification,omitempty"`
	CustomizedScalingMetricSpecification *ScalingPlan_CustomizedScalingMetricSpecification `yaml:"CustomizedScalingMetricSpecification,omitempty"`
}

func (resource ScalingPlan_TargetTrackingConfiguration) Validate() []error {
	errs := []error{}

	if resource.TargetValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetValue'"))
	}
	return errs
}
