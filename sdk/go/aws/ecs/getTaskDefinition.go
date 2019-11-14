// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ECS task definition data source allows access to details of
// a specific AWS ECS task definition.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ecs_task_definition.html.markdown.
func LookupTaskDefinition(ctx *pulumi.Context, args *GetTaskDefinitionArgs) (*GetTaskDefinitionResult, error) {
	var rv GetTaskDefinitionResult
	err := ctx.Invoke("aws:ecs/getTaskDefinition:getTaskDefinition", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTaskDefinition.
type GetTaskDefinitionArgs struct {
	// The family for the latest ACTIVE revision, family and revision (family:revision) for a specific revision in the family, the ARN of the task definition to access to.
	TaskDefinition string `pulumi:"taskDefinition"`
}

// A collection of values returned by getTaskDefinition.
type GetTaskDefinitionResult struct {
	// The family of this task definition
	Family string `pulumi:"family"`
	// The Docker networking mode to use for the containers in this task.
	NetworkMode string `pulumi:"networkMode"`
	// The revision of this task definition
	Revision int `pulumi:"revision"`
	// The status of this task definition
	Status string `pulumi:"status"`
	TaskDefinition string `pulumi:"taskDefinition"`
	// The ARN of the IAM role that containers in this task can assume
	TaskRoleArn string `pulumi:"taskRoleArn"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
