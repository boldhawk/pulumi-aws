// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an ElastiCache Security Group to control access to one or more cache
// clusters.
// 
// > **NOTE:** ElastiCache Security Groups are for use only when working with an
// ElastiCache cluster **outside** of a VPC. If you are using a VPC, see the
// ElastiCache Subnet Group resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/elasticache_security_group.html.markdown.
type SecurityGroup struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// description for the cache security group. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`

	// Name for the cache security group. This value is stored as a lowercase string.
	Name pulumi.StringOutput `pulumi:"name"`

	// List of EC2 security group names to be
	// authorized for ingress to the cache security group
	SecurityGroupNames pulumi.ArrayOutput `pulumi:"securityGroupNames"`
}

// NewSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroup(ctx *pulumi.Context,
	name string, args *SecurityGroupArgs, opts ...pulumi.ResourceOpt) (*SecurityGroup, error) {
	if args == nil || args.SecurityGroupNames == nil {
		return nil, errors.New("missing required argument 'SecurityGroupNames'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["description"] = pulumi.Any("Managed by Pulumi")
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["securityGroupNames"] = args.SecurityGroupNames
	}
	var resource SecurityGroup
	err := ctx.RegisterResource("aws:elasticache/securityGroup:SecurityGroup", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityGroup gets an existing SecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SecurityGroupState, opts ...pulumi.ResourceOpt) (*SecurityGroup, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["securityGroupNames"] = state.SecurityGroupNames
	}
	var resource SecurityGroup
	err := ctx.ReadResource("aws:elasticache/securityGroup:SecurityGroup", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *SecurityGroup) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *SecurityGroup) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering SecurityGroup resources.
type SecurityGroupState struct {
	// description for the cache security group. Defaults to "Managed by Pulumi".
	Description pulumi.StringInput `pulumi:"description"`
	// Name for the cache security group. This value is stored as a lowercase string.
	Name pulumi.StringInput `pulumi:"name"`
	// List of EC2 security group names to be
	// authorized for ingress to the cache security group
	SecurityGroupNames pulumi.ArrayInput `pulumi:"securityGroupNames"`
}

// The set of arguments for constructing a SecurityGroup resource.
type SecurityGroupArgs struct {
	// description for the cache security group. Defaults to "Managed by Pulumi".
	Description pulumi.StringInput `pulumi:"description"`
	// Name for the cache security group. This value is stored as a lowercase string.
	Name pulumi.StringInput `pulumi:"name"`
	// List of EC2 security group names to be
	// authorized for ingress to the cache security group
	SecurityGroupNames pulumi.ArrayInput `pulumi:"securityGroupNames"`
}
