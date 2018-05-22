package properties

	import "fmt"

type Bucket_Rule struct {
	
	
	
	
	
	
	
	
	
	
	
	
	ExpirationDate interface{} `yaml:"ExpirationDate,omitempty"`
	ExpirationInDays interface{} `yaml:"ExpirationInDays,omitempty"`
	Id interface{} `yaml:"Id,omitempty"`
	NoncurrentVersionExpirationInDays interface{} `yaml:"NoncurrentVersionExpirationInDays,omitempty"`
	Prefix interface{} `yaml:"Prefix,omitempty"`
	Status interface{} `yaml:"Status"`
	Transition *Bucket_Transition `yaml:"Transition,omitempty"`
	NoncurrentVersionTransition *Bucket_NoncurrentVersionTransition `yaml:"NoncurrentVersionTransition,omitempty"`
	NoncurrentVersionTransitions interface{} `yaml:"NoncurrentVersionTransitions,omitempty"`
	TagFilters interface{} `yaml:"TagFilters,omitempty"`
	Transitions interface{} `yaml:"Transitions,omitempty"`
	AbortIncompleteMultipartUpload *Bucket_AbortIncompleteMultipartUpload `yaml:"AbortIncompleteMultipartUpload,omitempty"`
}

func (resource Bucket_Rule) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	
	
	if resource.Status == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Status'"))
	}
	return errs
}
