// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an EC2 Transit Gateway Route Table association.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route_table_association.html.markdown.
type RouteTableAssociation struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// Identifier of the resource
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`

	// Type of the resource
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`

	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId pulumi.StringOutput `pulumi:"transitGatewayAttachmentId"`

	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId pulumi.StringOutput `pulumi:"transitGatewayRouteTableId"`
}

// NewRouteTableAssociation registers a new resource with the given unique name, arguments, and options.
func NewRouteTableAssociation(ctx *pulumi.Context,
	name string, args *RouteTableAssociationArgs, opts ...pulumi.ResourceOpt) (*RouteTableAssociation, error) {
	if args == nil || args.TransitGatewayAttachmentId == nil {
		return nil, errors.New("missing required argument 'TransitGatewayAttachmentId'")
	}
	if args == nil || args.TransitGatewayRouteTableId == nil {
		return nil, errors.New("missing required argument 'TransitGatewayRouteTableId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["transitGatewayAttachmentId"] = args.TransitGatewayAttachmentId
		inputs["transitGatewayRouteTableId"] = args.TransitGatewayRouteTableId
	}
	var resource RouteTableAssociation
	err := ctx.RegisterResource("aws:ec2transitgateway/routeTableAssociation:RouteTableAssociation", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteTableAssociation gets an existing RouteTableAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteTableAssociation(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RouteTableAssociationState, opts ...pulumi.ResourceOpt) (*RouteTableAssociation, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["resourceId"] = state.ResourceId
		inputs["resourceType"] = state.ResourceType
		inputs["transitGatewayAttachmentId"] = state.TransitGatewayAttachmentId
		inputs["transitGatewayRouteTableId"] = state.TransitGatewayRouteTableId
	}
	var resource RouteTableAssociation
	err := ctx.ReadResource("aws:ec2transitgateway/routeTableAssociation:RouteTableAssociation", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *RouteTableAssociation) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *RouteTableAssociation) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering RouteTableAssociation resources.
type RouteTableAssociationState struct {
	// Identifier of the resource
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// Type of the resource
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId pulumi.StringInput `pulumi:"transitGatewayAttachmentId"`
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId pulumi.StringInput `pulumi:"transitGatewayRouteTableId"`
}

// The set of arguments for constructing a RouteTableAssociation resource.
type RouteTableAssociationArgs struct {
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId pulumi.StringInput `pulumi:"transitGatewayAttachmentId"`
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId pulumi.StringInput `pulumi:"transitGatewayRouteTableId"`
}
