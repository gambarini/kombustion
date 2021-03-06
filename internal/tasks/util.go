package tasks

import (
	"strings"

	"github.com/KablamoOSS/kombustion/internal/cloudformation"
	"github.com/KablamoOSS/kombustion/types"
	"github.com/aws/aws-sdk-go/aws"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/urfave/cli"
)

func getParamMap(c *cli.Context) map[string]string {
	paramMap := make(map[string]string)
	params := c.StringSlice("param")
	for _, param := range params {
		parts := strings.Split(param, "=")
		if len(parts) > 1 {
			paramMap[parts[0]] = strings.Join(parts[1:], "=")
		}
	}
	return paramMap
}

func resolveParameters(c *cli.Context, cfYaml cloudformation.YamlCloudformation) []*awsCF.Parameter {
	results := []*awsCF.Parameter{}

	// Get params from the envFile
	env := cloudformation.ResolveEnvironment(c.String("env-file"), c.String("env"))

	// override envFile values with optional --param values
	params := getParamMap(c)
	for k, v := range params {
		env[k] = v
	}

	// convert to aws Parameter list
	for paramK := range cfYaml.Parameters {
		for k, v := range env {
			if paramK == k {
				if s, ok := v.(string); ok {
					// Filter to params in the stack
					results = append(results, &awsCF.Parameter{
						ParameterKey:   aws.String(k),
						ParameterValue: aws.String(s),
					})
				}
			}
		}
	}

	return results
}

func resolveParametersS3(c *cli.Context) []*awsCF.Parameter {
	results := []*awsCF.Parameter{}

	var params types.TemplateObject

	// override envFile values with optional --param values
	paramMap := getParamMap(c)
	for k, v := range paramMap {
		params[k] = v
	}

	// convert to aws Parameter list
	for k, v := range params {
		if s, ok := v.(string); ok {
			// Filter to params in the stack
			results = append(results, &awsCF.Parameter{
				ParameterKey:   aws.String(k),
				ParameterValue: aws.String(s),
			})
		}
	}

	return results
}
