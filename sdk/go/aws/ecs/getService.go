// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ECS Service data source allows access to details of a specific
// Service within a AWS ECS Cluster.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ecs_service.html.markdown.
func LookupService(ctx *pulumi.Context, args *GetServiceArgs) (*GetServiceResult, error) {
	var rv GetServiceResult
	err := ctx.Invoke("aws:ecs/getService:getService", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getService.
type GetServiceArgs struct {
	// The arn of the ECS Cluster
	ClusterArn string `pulumi:"clusterArn"`
	// The name of the ECS Service
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getService.
type GetServiceResult struct {
	// The ARN of the ECS Service
	Arn string `pulumi:"arn"`
	ClusterArn string `pulumi:"clusterArn"`
	// The number of tasks for the ECS Service
	DesiredCount int `pulumi:"desiredCount"`
	// The launch type for the ECS Service
	LaunchType string `pulumi:"launchType"`
	// The scheduling strategy for the ECS Service
	SchedulingStrategy string `pulumi:"schedulingStrategy"`
	ServiceName string `pulumi:"serviceName"`
	// The family for the latest ACTIVE revision
	TaskDefinition string `pulumi:"taskDefinition"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
