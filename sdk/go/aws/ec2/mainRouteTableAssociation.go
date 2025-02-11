// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource for managing the main routing table of a VPC.
// 
// ## Notes
// 
// On VPC creation, the AWS API always creates an initial Main Route Table. This
// resource records the ID of that Route Table under `originalRouteTableId`.
// The "Delete" action for a `mainRouteTableAssociation` consists of resetting
// this original table as the Main Route Table for the VPC. You'll see this
// additional Route Table in the AWS console; it must remain intact in order for
// the `mainRouteTableAssociation` delete to work properly.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/main_route_table_association.html.markdown.
type MainRouteTableAssociation struct {
	s *pulumi.ResourceState
}

// NewMainRouteTableAssociation registers a new resource with the given unique name, arguments, and options.
func NewMainRouteTableAssociation(ctx *pulumi.Context,
	name string, args *MainRouteTableAssociationArgs, opts ...pulumi.ResourceOpt) (*MainRouteTableAssociation, error) {
	if args == nil || args.RouteTableId == nil {
		return nil, errors.New("missing required argument 'RouteTableId'")
	}
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["routeTableId"] = nil
		inputs["vpcId"] = nil
	} else {
		inputs["routeTableId"] = args.RouteTableId
		inputs["vpcId"] = args.VpcId
	}
	inputs["originalRouteTableId"] = nil
	s, err := ctx.RegisterResource("aws:ec2/mainRouteTableAssociation:MainRouteTableAssociation", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MainRouteTableAssociation{s: s}, nil
}

// GetMainRouteTableAssociation gets an existing MainRouteTableAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMainRouteTableAssociation(ctx *pulumi.Context,
	name string, id pulumi.ID, state *MainRouteTableAssociationState, opts ...pulumi.ResourceOpt) (*MainRouteTableAssociation, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["originalRouteTableId"] = state.OriginalRouteTableId
		inputs["routeTableId"] = state.RouteTableId
		inputs["vpcId"] = state.VpcId
	}
	s, err := ctx.ReadResource("aws:ec2/mainRouteTableAssociation:MainRouteTableAssociation", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MainRouteTableAssociation{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *MainRouteTableAssociation) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *MainRouteTableAssociation) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Used internally, see __Notes__ below
func (r *MainRouteTableAssociation) OriginalRouteTableId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["originalRouteTableId"])
}

// The ID of the Route Table to set as the new
// main route table for the target VPC
func (r *MainRouteTableAssociation) RouteTableId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["routeTableId"])
}

// The ID of the VPC whose main route table should be set
func (r *MainRouteTableAssociation) VpcId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["vpcId"])
}

// Input properties used for looking up and filtering MainRouteTableAssociation resources.
type MainRouteTableAssociationState struct {
	// Used internally, see __Notes__ below
	OriginalRouteTableId interface{}
	// The ID of the Route Table to set as the new
	// main route table for the target VPC
	RouteTableId interface{}
	// The ID of the VPC whose main route table should be set
	VpcId interface{}
}

// The set of arguments for constructing a MainRouteTableAssociation resource.
type MainRouteTableAssociationArgs struct {
	// The ID of the Route Table to set as the new
	// main route table for the target VPC
	RouteTableId interface{}
	// The ID of the VPC whose main route table should be set
	VpcId interface{}
}
