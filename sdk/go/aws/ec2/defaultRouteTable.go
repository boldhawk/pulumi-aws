// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage a Default VPC Routing Table.
// 
// Each VPC created in AWS comes with a Default Route Table that can be managed, but not
// destroyed. **This is an advanced resource**, and has special caveats to be aware
// of when using it. Please read this document in its entirety before using this
// resource. It is recommended you **do not** use both `ec2.DefaultRouteTable` to
// manage the default route table **and** use the `ec2.MainRouteTableAssociation`,
// due to possible conflict in routes.
// 
// The `ec2.DefaultRouteTable` behaves differently from normal resources, in that
// this provider does not _create_ this resource, but instead attempts to "adopt" it
// into management. We can do this because each VPC created has a Default Route
// Table that cannot be destroyed, and is created with a single route.
// 
// When this provider first adopts the Default Route Table, it **immediately removes all
// defined routes**. It then proceeds to create any routes specified in the
// configuration. This step is required so that only the routes specified in the
// configuration present in the Default Route Table.
// 
// For more information about Route Tables, see the AWS Documentation on
// [Route Tables][aws-route-tables].
// 
// For more information about managing normal Route Tables in this provider, see our
// documentation on [ec2.RouteTable][tf-route-tables].
// 
// > **NOTE on Route Tables and Routes:** This provider currently
// provides both a standalone Route resource and a Route Table resource with routes
// defined in-line. At this time you cannot use a Route Table with in-line routes
// in conjunction with any Route resources. Doing so will cause
// a conflict of rule settings and will overwrite routes.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/default_route_table.html.markdown.
type DefaultRouteTable struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ID of the Default Routing Table.
	DefaultRouteTableId pulumi.StringOutput `pulumi:"defaultRouteTableId"`

	// The ID of the AWS account that owns the route table
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`

	// A list of virtual gateways for propagation.
	PropagatingVgws pulumi.ArrayOutput `pulumi:"propagatingVgws"`

	// A list of route objects. Their keys are documented below.
	// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
	Routes pulumi.ArrayOutput `pulumi:"routes"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewDefaultRouteTable registers a new resource with the given unique name, arguments, and options.
func NewDefaultRouteTable(ctx *pulumi.Context,
	name string, args *DefaultRouteTableArgs, opts ...pulumi.ResourceOpt) (*DefaultRouteTable, error) {
	if args == nil || args.DefaultRouteTableId == nil {
		return nil, errors.New("missing required argument 'DefaultRouteTableId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["defaultRouteTableId"] = args.DefaultRouteTableId
		inputs["propagatingVgws"] = args.PropagatingVgws
		inputs["routes"] = args.Routes
		inputs["tags"] = args.Tags
	}
	var resource DefaultRouteTable
	err := ctx.RegisterResource("aws:ec2/defaultRouteTable:DefaultRouteTable", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultRouteTable gets an existing DefaultRouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultRouteTable(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DefaultRouteTableState, opts ...pulumi.ResourceOpt) (*DefaultRouteTable, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["defaultRouteTableId"] = state.DefaultRouteTableId
		inputs["ownerId"] = state.OwnerId
		inputs["propagatingVgws"] = state.PropagatingVgws
		inputs["routes"] = state.Routes
		inputs["tags"] = state.Tags
		inputs["vpcId"] = state.VpcId
	}
	var resource DefaultRouteTable
	err := ctx.ReadResource("aws:ec2/defaultRouteTable:DefaultRouteTable", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *DefaultRouteTable) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *DefaultRouteTable) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering DefaultRouteTable resources.
type DefaultRouteTableState struct {
	// The ID of the Default Routing Table.
	DefaultRouteTableId pulumi.StringInput `pulumi:"defaultRouteTableId"`
	// The ID of the AWS account that owns the route table
	OwnerId pulumi.StringInput `pulumi:"ownerId"`
	// A list of virtual gateways for propagation.
	PropagatingVgws pulumi.ArrayInput `pulumi:"propagatingVgws"`
	// A list of route objects. Their keys are documented below.
	// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
	Routes pulumi.ArrayInput `pulumi:"routes"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

// The set of arguments for constructing a DefaultRouteTable resource.
type DefaultRouteTableArgs struct {
	// The ID of the Default Routing Table.
	DefaultRouteTableId pulumi.StringInput `pulumi:"defaultRouteTableId"`
	// A list of virtual gateways for propagation.
	PropagatingVgws pulumi.ArrayInput `pulumi:"propagatingVgws"`
	// A list of route objects. Their keys are documented below.
	// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
	Routes pulumi.ArrayInput `pulumi:"routes"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
