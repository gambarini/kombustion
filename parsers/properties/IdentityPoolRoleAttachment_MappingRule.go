package properties

	import "fmt"

type IdentityPoolRoleAttachment_MappingRule struct {
	
	
	
	
	Claim interface{} `yaml:"Claim"`
	MatchType interface{} `yaml:"MatchType"`
	RoleARN interface{} `yaml:"RoleARN"`
	Value interface{} `yaml:"Value"`
}

func (resource IdentityPoolRoleAttachment_MappingRule) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.Claim == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Claim'"))
	}
	if resource.MatchType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MatchType'"))
	}
	if resource.RoleARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleARN'"))
	}
	if resource.Value == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Value'"))
	}
	return errs
}
