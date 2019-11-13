// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an EC2 Transit Gateway Route Table.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route_table.html.markdown.
type RouteTable struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// Boolean whether this is the default association route table for the EC2 Transit Gateway.
	DefaultAssociationRouteTable pulumi.BoolOutput `pulumi:"defaultAssociationRouteTable"`

	// Boolean whether this is the default propagation route table for the EC2 Transit Gateway.
	DefaultPropagationRouteTable pulumi.BoolOutput `pulumi:"defaultPropagationRouteTable"`

	// Key-value tags for the EC2 Transit Gateway Route Table.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// Identifier of EC2 Transit Gateway.
	TransitGatewayId pulumi.StringOutput `pulumi:"transitGatewayId"`
}

// NewRouteTable registers a new resource with the given unique name, arguments, and options.
func NewRouteTable(ctx *pulumi.Context,
	name string, args *RouteTableArgs, opts ...pulumi.ResourceOpt) (*RouteTable, error) {
	if args == nil || args.TransitGatewayId == nil {
		return nil, errors.New("missing required argument 'TransitGatewayId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["tags"] = args.Tags
		inputs["transitGatewayId"] = args.TransitGatewayId
	}
	var resource RouteTable
	err := ctx.RegisterResource("aws:ec2transitgateway/routeTable:RouteTable", name, inputs, &resource, opts...)
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
		inputs["defaultAssociationRouteTable"] = state.DefaultAssociationRouteTable
		inputs["defaultPropagationRouteTable"] = state.DefaultPropagationRouteTable
		inputs["tags"] = state.Tags
		inputs["transitGatewayId"] = state.TransitGatewayId
	}
	var resource RouteTable
	err := ctx.ReadResource("aws:ec2transitgateway/routeTable:RouteTable", name, id, inputs, &resource, opts...)
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
	// Boolean whether this is the default association route table for the EC2 Transit Gateway.
	DefaultAssociationRouteTable pulumi.BoolInput `pulumi:"defaultAssociationRouteTable"`
	// Boolean whether this is the default propagation route table for the EC2 Transit Gateway.
	DefaultPropagationRouteTable pulumi.BoolInput `pulumi:"defaultPropagationRouteTable"`
	// Key-value tags for the EC2 Transit Gateway Route Table.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId pulumi.StringInput `pulumi:"transitGatewayId"`
}

// The set of arguments for constructing a RouteTable resource.
type RouteTableArgs struct {
	// Key-value tags for the EC2 Transit Gateway Route Table.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId pulumi.StringInput `pulumi:"transitGatewayId"`
}
