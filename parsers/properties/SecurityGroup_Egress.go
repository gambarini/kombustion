package properties

	import "fmt"

type SecurityGroup_Egress struct {
	
	
	
	
	
	
	
	
	CidrIp interface{} `yaml:"CidrIp,omitempty"`
	CidrIpv6 interface{} `yaml:"CidrIpv6,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	DestinationPrefixListId interface{} `yaml:"DestinationPrefixListId,omitempty"`
	DestinationSecurityGroupId interface{} `yaml:"DestinationSecurityGroupId,omitempty"`
	FromPort interface{} `yaml:"FromPort,omitempty"`
	IpProtocol interface{} `yaml:"IpProtocol"`
	ToPort interface{} `yaml:"ToPort,omitempty"`
}

func (resource SecurityGroup_Egress) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	if resource.IpProtocol == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IpProtocol'"))
	}
	return errs
}
