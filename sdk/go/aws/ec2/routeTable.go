// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create a VPC routing table.
// 
// > **NOTE on Route Tables and Routes:** This provider currently
// provides both a standalone Route resource and a Route Table resource with routes
// defined in-line. At this time you cannot use a Route Table with in-line routes
// in conjunction with any Route resources. Doing so will cause
// a conflict of rule settings and will overwrite rules.
// 
// > **NOTE on `gatewayId` and `natGatewayId`:** The AWS API is very forgiving with these two
// attributes and the `ec2.RouteTable` resource can be created with a NAT ID specified as a Gateway ID attribute.
// This _will_ lead to a permanent diff between your configuration and statefile, as the API returns the correct
// parameters in the returned route table. If you're experiencing constant diffs in your `ec2.RouteTable` resources,
// the first thing to check is whether or not you're specifying a NAT ID instead of a Gateway ID, or vice-versa.
// 
// > **NOTE on `propagatingVgws` and the `ec2.VpnGatewayRoutePropagation` resource:**
// If the `propagatingVgws` argument is present, it's not supported to _also_
// define route propagations using `ec2.VpnGatewayRoutePropagation`, since
// this resource will delete any propagating gateways not explicitly listed in
// `propagatingVgws`. Omit this argument when defining route propagation using
// the separate resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/route_table.html.markdown.
type RouteTable struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ID of the AWS account that owns the route table.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`

	// A list of virtual gateways for propagation.
	PropagatingVgws pulumi.ArrayOutput `pulumi:"propagatingVgws"`

	// A list of route objects. Their keys are documented below. This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
	Routes pulumi.ArrayOutput `pulumi:"routes"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// The VPC ID.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewRouteTable registers a new resource with the given unique name, arguments, and options.
func NewRouteTable(ctx *pulumi.Context,
	name string, args *RouteTableArgs, opts ...pulumi.ResourceOpt) (*RouteTable, error) {
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["propagatingVgws"] = args.PropagatingVgws
		inputs["routes"] = args.Routes
		inputs["tags"] = args.Tags
		inputs["vpcId"] = args.VpcId
	}
	var resource RouteTable
	err := ctx.RegisterResource("aws:ec2/routeTable:RouteTable", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteTable gets an existing RouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteTable(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RouteTableState, opts ...pulumi.ResourceOpt) (*RouteTable, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["ownerId"] = state.OwnerId
		inputs["propagatingVgws"] = state.PropagatingVgws
		inputs["routes"] = state.Routes
		inputs["tags"] = state.Tags
		inputs["vpcId"] = state.VpcId
	}
	var resource RouteTable
	err := ctx.ReadResource("aws:ec2/routeTable:RouteTable", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *RouteTable) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *RouteTable) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering RouteTable resources.
type RouteTableState struct {
	// The ID of the AWS account that owns the route table.
	OwnerId pulumi.StringInput `pulumi:"ownerId"`
	// A list of virtual gateways for propagation.
	PropagatingVgws pulumi.ArrayInput `pulumi:"propagatingVgws"`
	// A list of route objects. Their keys are documented below. This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
	Routes pulumi.ArrayInput `pulumi:"routes"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The VPC ID.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

// The set of arguments for constructing a RouteTable resource.
type RouteTableArgs struct {
	// A list of virtual gateways for propagation.
	PropagatingVgws pulumi.ArrayInput `pulumi:"propagatingVgws"`
	// A list of route objects. Their keys are documented below. This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
	Routes pulumi.ArrayInput `pulumi:"routes"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The VPC ID.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}
