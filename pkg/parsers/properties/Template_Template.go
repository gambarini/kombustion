package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type Template_Template struct {
	HtmlPart     interface{} `yaml:"HtmlPart,omitempty"`
	SubjectPart  interface{} `yaml:"SubjectPart,omitempty"`
	TemplateName interface{} `yaml:"TemplateName,omitempty"`
	TextPart     interface{} `yaml:"TextPart,omitempty"`
}

func (resource Template_Template) Validate() []error {
	errs := []error{}

	return errs
}
