package tasks

import (
	"github.com/KablamoOSS/kombustion/internal/cloudformation"
	"github.com/KablamoOSS/kombustion/internal/cloudformation/tasks"
	"github.com/aws/aws-sdk-go/aws"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/urfave/cli"
)

// UpsertFlags for kombustion upsert
var UpsertFlags = []cli.Flag{
	cli.StringSliceFlag{
		Name:  "param, p",
		Usage: "cloudformation parameters. eg. [ --param Env=dev --param BucketName=test ]",
	},
	cli.BoolFlag{
		Name:  "no-base-outputs, b",
		Usage: "disable generation of outputs for Base AWS types",
	},
	cli.BoolFlag{
		Name:  "iam, i",
		Usage: "gives the capability to perform upserts of IAM resources",
	},
	cli.StringSliceFlag{
		Name:  "capability",
		Usage: "set capabilities for the upsert, eg [ --capability CAPABILITY_IAM ]",
	},
}

func init() {
	UpsertFlags = append(CloudFormationStackFlags)
}

// Upsert a stack
func Upsert(c *cli.Context) {
	paramMap := getParamMap(c)

	// Template generation parameters
	generateParams := cloudformation.GenerateParams{
		Filename:           c.Args().Get(0),
		EnvFile:            c.String("env-file"),
		Env:                c.String("env"),
		DisableBaseOutputs: c.Bool("no-base-outputs"),
		ParamMap:           paramMap,
	}

	capabilities := getCapabilities(c)

	// Cloudformation Stack parameters
	var parameters []*awsCF.Parameter

	stackName := c.Args().Get(0)
	if len(c.String("stack-name")) > 0 {
		stackName = c.String("stack-name")
	}
	if len(c.String("url")) > 0 {
		parameters = resolveParametersS3(c)

		templateURL := c.String("url")

		tasks.UpsertStackViaS3(
			templateURL,
			parameters,
			capabilities,
			stackName,
			c.GlobalString("profile"),
			c.String("region"),
		)
	} else {

		templateBody, cfYaml := tasks.GenerateYamlTemplate(generateParams)
		parameters = resolveParameters(c, cfYaml)

		tasks.UpsertStack(
			templateBody,
			parameters,
			capabilities,
			stackName,
			c.GlobalString("profile"),
			c.String("region"),
		)
	}
}

// Extract capabilities from the cli call
func getCapabilities(c *cli.Context) []*string {
	capabilities := aws.StringSlice([]string{})

	capabilitySlices := c.StringSlice("capability")

	iamCapability := "CAPABILITY_NAMED_IAM"
	haveAddedIam := false

	for _, capability := range capabilitySlices {
		if capability == iamCapability {
			haveAddedIam = true
		}
		capabilities = append(capabilities, &capability)
	}

	if c.Bool("iam") {
		if haveAddedIam == false {
			capabilities = append(capabilities, &iamCapability)
		}
	}

	return capabilities
}
