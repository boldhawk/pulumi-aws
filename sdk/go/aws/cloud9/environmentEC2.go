// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloud9

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Cloud9 EC2 Development Environment.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloud9_environment_ec2.html.markdown.
type EnvironmentEC2 struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN of the environment.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The number of minutes until the running instance is shut down after the environment has last been used.
	AutomaticStopTimeMinutes pulumi.IntOutput `pulumi:"automaticStopTimeMinutes"`

	// The description of the environment.
	Description pulumi.StringOutput `pulumi:"description"`

	// The type of instance to connect to the environment, e.g. `t2.micro`.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`

	// The name of the environment.
	Name pulumi.StringOutput `pulumi:"name"`

	// The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
	OwnerArn pulumi.StringOutput `pulumi:"ownerArn"`

	// The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`

	// The type of the environment (e.g. `ssh` or `ec2`)
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEnvironmentEC2 registers a new resource with the given unique name, arguments, and options.
func NewEnvironmentEC2(ctx *pulumi.Context,
	name string, args *EnvironmentEC2Args, opts ...pulumi.ResourceOpt) (*EnvironmentEC2, error) {
	if args == nil || args.InstanceType == nil {
		return nil, errors.New("missing required argument 'InstanceType'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["automaticStopTimeMinutes"] = args.AutomaticStopTimeMinutes
		inputs["description"] = args.Description
		inputs["instanceType"] = args.InstanceType
		inputs["name"] = args.Name
		inputs["ownerArn"] = args.OwnerArn
		inputs["subnetId"] = args.SubnetId
	}
	var resource EnvironmentEC2
	err := ctx.RegisterResource("aws:cloud9/environmentEC2:EnvironmentEC2", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironmentEC2 gets an existing EnvironmentEC2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironmentEC2(ctx *pulumi.Context,
	name string, id pulumi.ID, state *EnvironmentEC2State, opts ...pulumi.ResourceOpt) (*EnvironmentEC2, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["automaticStopTimeMinutes"] = state.AutomaticStopTimeMinutes
		inputs["description"] = state.Description
		inputs["instanceType"] = state.InstanceType
		inputs["name"] = state.Name
		inputs["ownerArn"] = state.OwnerArn
		inputs["subnetId"] = state.SubnetId
		inputs["type"] = state.Type
	}
	var resource EnvironmentEC2
	err := ctx.ReadResource("aws:cloud9/environmentEC2:EnvironmentEC2", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *EnvironmentEC2) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *EnvironmentEC2) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering EnvironmentEC2 resources.
type EnvironmentEC2State struct {
	// The ARN of the environment.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The number of minutes until the running instance is shut down after the environment has last been used.
	AutomaticStopTimeMinutes pulumi.IntInput `pulumi:"automaticStopTimeMinutes"`
	// The description of the environment.
	Description pulumi.StringInput `pulumi:"description"`
	// The type of instance to connect to the environment, e.g. `t2.micro`.
	InstanceType pulumi.StringInput `pulumi:"instanceType"`
	// The name of the environment.
	Name pulumi.StringInput `pulumi:"name"`
	// The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
	OwnerArn pulumi.StringInput `pulumi:"ownerArn"`
	// The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
	SubnetId pulumi.StringInput `pulumi:"subnetId"`
	// The type of the environment (e.g. `ssh` or `ec2`)
	Type pulumi.StringInput `pulumi:"type"`
}

// The set of arguments for constructing a EnvironmentEC2 resource.
type EnvironmentEC2Args struct {
	// The number of minutes until the running instance is shut down after the environment has last been used.
	AutomaticStopTimeMinutes pulumi.IntInput `pulumi:"automaticStopTimeMinutes"`
	// The description of the environment.
	Description pulumi.StringInput `pulumi:"description"`
	// The type of instance to connect to the environment, e.g. `t2.micro`.
	InstanceType pulumi.StringInput `pulumi:"instanceType"`
	// The name of the environment.
	Name pulumi.StringInput `pulumi:"name"`
	// The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
	OwnerArn pulumi.StringInput `pulumi:"ownerArn"`
	// The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
	SubnetId pulumi.StringInput `pulumi:"subnetId"`
}
