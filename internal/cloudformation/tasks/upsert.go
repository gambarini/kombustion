package tasks

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
)

// Upsert a stack
func UpsertStack(
	templateBody []byte,
	parameters []*awsCF.Parameter,
	capabilities []*string,
	stackName, profile, region string,
) {
	cf := GetCloudformationClient(profile, region)

	var err error
	var status *awsCF.DescribeStacksOutput

	// use template from file
	_, err = cf.DescribeStacks(&awsCF.DescribeStacksInput{StackName: aws.String(stackName)})
	if err == nil { //update
		_, err = cf.UpdateStack(&awsCF.UpdateStackInput{
			StackName:    aws.String(stackName),
			TemplateBody: aws.String(string(templateBody)),
			Parameters:   parameters,
			Capabilities: capabilities,
		})
	} else {
		_, err = cf.CreateStack(&awsCF.CreateStackInput{
			StackName:    aws.String(stackName),
			TemplateBody: aws.String(string(templateBody)),
			Parameters:   parameters,
			Capabilities: capabilities,
		})
	}
	checkError(err)

	// Make sure upsert works
	for {
		time.Sleep(2 * time.Second)
		status, err = cf.DescribeStacks(&awsCF.DescribeStacksInput{StackName: aws.String(stackName)})
		checkError(err)
		if len(status.Stacks) > 0 {
			stack := status.Stacks[0]
			stackStatus := *stack.StackStatus
			fmt.Println(stackStatus)
			if stackStatus != awsCF.StackStatusCreateInProgress &&
				stackStatus != awsCF.StackStatusUpdateInProgress &&
				stackStatus != awsCF.StackStatusUpdateCompleteCleanupInProgress {
				if stackStatus == awsCF.StackStatusCreateComplete ||
					stackStatus == awsCF.StackStatusUpdateComplete {
					os.Exit(0)
				} else {
					log.Error("Upsert Failed: ")
					time.Sleep(2 * time.Second)

					PrintStackEvents(cf, stackName)
					os.Exit(1)
				}
			}
		}
	}
}

// Upsert a stack
func UpsertStackViaS3(
	templateURL string,
	parameters []*awsCF.Parameter,
	capabilities []*string,
	stackName, profile, region string,
) {
	cf := GetCloudformationClient(profile, region)

	var err error
	var status *awsCF.DescribeStacksOutput

	// use cf template url
	_, err = cf.DescribeStacks(&awsCF.DescribeStacksInput{StackName: aws.String(stackName)})
	if err == nil { //update
		_, err = cf.UpdateStack(&awsCF.UpdateStackInput{
			StackName:    aws.String(stackName),
			TemplateURL:  aws.String(templateURL),
			Parameters:   parameters,
			Capabilities: capabilities,
		})
	} else { //create
		_, err = cf.CreateStack(&awsCF.CreateStackInput{
			StackName:    aws.String(stackName),
			TemplateURL:  aws.String(templateURL),
			Parameters:   parameters,
			Capabilities: capabilities,
		})
	}
	checkError(err)

	// Make sure upsert works
	for {
		time.Sleep(2 * time.Second)
		status, err = cf.DescribeStacks(&awsCF.DescribeStacksInput{StackName: aws.String(stackName)})
		checkError(err)
		if len(status.Stacks) > 0 {
			stack := status.Stacks[0]
			stackStatus := *stack.StackStatus
			fmt.Println(stackStatus)
			if stackStatus != awsCF.StackStatusCreateInProgress &&
				stackStatus != awsCF.StackStatusUpdateInProgress &&
				stackStatus != awsCF.StackStatusUpdateCompleteCleanupInProgress {
				if stackStatus == awsCF.StackStatusCreateComplete ||
					stackStatus == awsCF.StackStatusUpdateComplete {
					os.Exit(0)
				} else {
					log.Error("Upsert Failed: ")
					time.Sleep(2 * time.Second)

					PrintStackEvents(cf, stackName)
					os.Exit(1)
				}
			}
		}
	}
}
