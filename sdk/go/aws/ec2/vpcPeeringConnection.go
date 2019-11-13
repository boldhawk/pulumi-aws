// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage a VPC peering connection.
// 
// > **NOTE on VPC Peering Connections and VPC Peering Connection Options:** This provider provides
// both a standalone VPC Peering Connection Options and a VPC Peering Connection
// resource with `accepter` and `requester` attributes. Do not manage options for the same VPC peering
// connection in both a VPC Peering Connection resource and a VPC Peering Connection Options resource.
// Doing so will cause a conflict of options and will overwrite the options.
// Using a VPC Peering Connection Options resource decouples management of the connection options from
// management of the VPC Peering Connection and allows options to be set correctly in cross-account scenarios.
// 
// > **Note:** For cross-account (requester's AWS account differs from the accepter's AWS account) or inter-region
// VPC Peering Connections use the `ec2.VpcPeeringConnection` resource to manage the requester's side of the
// connection and use the `ec2.VpcPeeringConnectionAccepter` resource to manage the accepter's side of the connection.
// 
// ## Notes
// 
// If both VPCs are not in the same AWS account do not enable the `autoAccept` attribute.
// The accepter can manage its side of the connection using the `ec2.VpcPeeringConnectionAccepter` resource
// or accept the connection manually using the AWS Management Console, AWS CLI, through SDKs, etc.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpc_peering_connection.html.markdown.
type VpcPeeringConnection struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The status of the VPC Peering Connection request.
	AcceptStatus pulumi.StringOutput `pulumi:"acceptStatus"`

	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter pulumi.AnyOutput `pulumi:"accepter"`

	// Accept the peering (both VPCs need to be in the same AWS account).
	AutoAccept pulumi.BoolOutput `pulumi:"autoAccept"`

	// The AWS account ID of the owner of the peer VPC.
	// Defaults to the account ID the [AWS provider][1] is currently connected to.
	PeerOwnerId pulumi.StringOutput `pulumi:"peerOwnerId"`

	// The region of the accepter VPC of the [VPC Peering Connection]. `autoAccept` must be `false`,
	// and use the `ec2.VpcPeeringConnectionAccepter` to manage the accepter side.
	PeerRegion pulumi.StringOutput `pulumi:"peerRegion"`

	// The ID of the VPC with which you are creating the VPC Peering Connection.
	PeerVpcId pulumi.StringOutput `pulumi:"peerVpcId"`

	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester pulumi.AnyOutput `pulumi:"requester"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// The ID of the requester VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewVpcPeeringConnection registers a new resource with the given unique name, arguments, and options.
func NewVpcPeeringConnection(ctx *pulumi.Context,
	name string, args *VpcPeeringConnectionArgs, opts ...pulumi.ResourceOpt) (*VpcPeeringConnection, error) {
	if args == nil || args.PeerVpcId == nil {
		return nil, errors.New("missing required argument 'PeerVpcId'")
	}
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["accepter"] = args.Accepter
		inputs["autoAccept"] = args.AutoAccept
		inputs["peerOwnerId"] = args.PeerOwnerId
		inputs["peerRegion"] = args.PeerRegion
		inputs["peerVpcId"] = args.PeerVpcId
		inputs["requester"] = args.Requester
		inputs["tags"] = args.Tags
		inputs["vpcId"] = args.VpcId
	}
	var resource VpcPeeringConnection
	err := ctx.RegisterResource("aws:ec2/vpcPeeringConnection:VpcPeeringConnection", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcPeeringConnection gets an existing VpcPeeringConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPeeringConnection(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VpcPeeringConnectionState, opts ...pulumi.ResourceOpt) (*VpcPeeringConnection, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["acceptStatus"] = state.AcceptStatus
		inputs["accepter"] = state.Accepter
		inputs["autoAccept"] = state.AutoAccept
		inputs["peerOwnerId"] = state.PeerOwnerId
		inputs["peerRegion"] = state.PeerRegion
		inputs["peerVpcId"] = state.PeerVpcId
		inputs["requester"] = state.Requester
		inputs["tags"] = state.Tags
		inputs["vpcId"] = state.VpcId
	}
	var resource VpcPeeringConnection
	err := ctx.ReadResource("aws:ec2/vpcPeeringConnection:VpcPeeringConnection", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *VpcPeeringConnection) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *VpcPeeringConnection) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering VpcPeeringConnection resources.
type VpcPeeringConnectionState struct {
	// The status of the VPC Peering Connection request.
	AcceptStatus pulumi.StringInput `pulumi:"acceptStatus"`
	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter pulumi.AnyInput `pulumi:"accepter"`
	// Accept the peering (both VPCs need to be in the same AWS account).
	AutoAccept pulumi.BoolInput `pulumi:"autoAccept"`
	// The AWS account ID of the owner of the peer VPC.
	// Defaults to the account ID the [AWS provider][1] is currently connected to.
	PeerOwnerId pulumi.StringInput `pulumi:"peerOwnerId"`
	// The region of the accepter VPC of the [VPC Peering Connection]. `autoAccept` must be `false`,
	// and use the `ec2.VpcPeeringConnectionAccepter` to manage the accepter side.
	PeerRegion pulumi.StringInput `pulumi:"peerRegion"`
	// The ID of the VPC with which you are creating the VPC Peering Connection.
	PeerVpcId pulumi.StringInput `pulumi:"peerVpcId"`
	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester pulumi.AnyInput `pulumi:"requester"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The ID of the requester VPC.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpcPeeringConnection resource.
type VpcPeeringConnectionArgs struct {
	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter pulumi.AnyInput `pulumi:"accepter"`
	// Accept the peering (both VPCs need to be in the same AWS account).
	AutoAccept pulumi.BoolInput `pulumi:"autoAccept"`
	// The AWS account ID of the owner of the peer VPC.
	// Defaults to the account ID the [AWS provider][1] is currently connected to.
	PeerOwnerId pulumi.StringInput `pulumi:"peerOwnerId"`
	// The region of the accepter VPC of the [VPC Peering Connection]. `autoAccept` must be `false`,
	// and use the `ec2.VpcPeeringConnectionAccepter` to manage the accepter side.
	PeerRegion pulumi.StringInput `pulumi:"peerRegion"`
	// The ID of the VPC with which you are creating the VPC Peering Connection.
	PeerVpcId pulumi.StringInput `pulumi:"peerVpcId"`
	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester pulumi.AnyInput `pulumi:"requester"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The ID of the requester VPC.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}
