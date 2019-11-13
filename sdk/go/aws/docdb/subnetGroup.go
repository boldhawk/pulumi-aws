// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package docdb

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an DocumentDB subnet group resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/docdb_subnet_group.html.markdown.
type SubnetGroup struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN of the docDB subnet group.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The description of the docDB subnet group. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`

	// The name of the docDB subnet group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`

	// A list of VPC subnet IDs.
	SubnetIds pulumi.ArrayOutput `pulumi:"subnetIds"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewSubnetGroup registers a new resource with the given unique name, arguments, and options.
func NewSubnetGroup(ctx *pulumi.Context,
	name string, args *SubnetGroupArgs, opts ...pulumi.ResourceOpt) (*SubnetGroup, error) {
	if args == nil || args.SubnetIds == nil {
		return nil, errors.New("missing required argument 'SubnetIds'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["description"] = pulumi.Any("Managed by Pulumi")
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["subnetIds"] = args.SubnetIds
		inputs["tags"] = args.Tags
	}
	var resource SubnetGroup
	err := ctx.RegisterResource("aws:docdb/subnetGroup:SubnetGroup", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetGroup gets an existing SubnetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SubnetGroupState, opts ...pulumi.ResourceOpt) (*SubnetGroup, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["subnetIds"] = state.SubnetIds
		inputs["tags"] = state.Tags
	}
	var resource SubnetGroup
	err := ctx.ReadResource("aws:docdb/subnetGroup:SubnetGroup", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *SubnetGroup) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *SubnetGroup) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering SubnetGroup resources.
type SubnetGroupState struct {
	// The ARN of the docDB subnet group.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The description of the docDB subnet group. Defaults to "Managed by Pulumi".
	Description pulumi.StringInput `pulumi:"description"`
	// The name of the docDB subnet group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// A list of VPC subnet IDs.
	SubnetIds pulumi.ArrayInput `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a SubnetGroup resource.
type SubnetGroupArgs struct {
	// The description of the docDB subnet group. Defaults to "Managed by Pulumi".
	Description pulumi.StringInput `pulumi:"description"`
	// The name of the docDB subnet group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// A list of VPC subnet IDs.
	SubnetIds pulumi.ArrayInput `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
