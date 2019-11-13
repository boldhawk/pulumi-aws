// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage the accepter's side of a VPC Peering Connection.
// 
// When a cross-account (requester's AWS account differs from the accepter's AWS account) or an inter-region
// VPC Peering Connection is created, a VPC Peering Connection resource is automatically created in the
// accepter's account.
// The requester can use the `ec2.VpcPeeringConnection` resource to manage its side of the connection
// and the accepter can use the `ec2.VpcPeeringConnectionAccepter` resource to "adopt" its side of the
// connection into management.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpc_peering_connection_accepter.html.markdown.
type VpcPeeringConnectionAccepter struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The status of the VPC Peering Connection request.
	AcceptStatus pulumi.StringOutput `pulumi:"acceptStatus"`

	// A configuration block that describes [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the accepter VPC.
	Accepter pulumi.AnyOutput `pulumi:"accepter"`

	// Whether or not to accept the peering request. Defaults to `false`.
	AutoAccept pulumi.BoolOutput `pulumi:"autoAccept"`

	// The AWS account ID of the owner of the requester VPC.
	PeerOwnerId pulumi.StringOutput `pulumi:"peerOwnerId"`

	// The region of the accepter VPC.
	PeerRegion pulumi.StringOutput `pulumi:"peerRegion"`

	// The ID of the requester VPC.
	PeerVpcId pulumi.StringOutput `pulumi:"peerVpcId"`

	// A configuration block that describes [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the requester VPC.
	Requester pulumi.AnyOutput `pulumi:"requester"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// The ID of the accepter VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`

	// The VPC Peering Connection ID to manage.
	VpcPeeringConnectionId pulumi.StringOutput `pulumi:"vpcPeeringConnectionId"`
}

// NewVpcPeeringConnectionAccepter registers a new resource with the given unique name, arguments, and options.
func NewVpcPeeringConnectionAccepter(ctx *pulumi.Context,
	name string, args *VpcPeeringConnectionAccepterArgs, opts ...pulumi.ResourceOpt) (*VpcPeeringConnectionAccepter, error) {
	if args == nil || args.VpcPeeringConnectionId == nil {
		return nil, errors.New("missing required argument 'VpcPeeringConnectionId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["accepter"] = args.Accepter
		inputs["autoAccept"] = args.AutoAccept
		inputs["requester"] = args.Requester
		inputs["tags"] = args.Tags
		inputs["vpcPeeringConnectionId"] = args.VpcPeeringConnectionId
	}
	var resource VpcPeeringConnectionAccepter
	err := ctx.RegisterResource("aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcPeeringConnectionAccepter gets an existing VpcPeeringConnectionAccepter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPeeringConnectionAccepter(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VpcPeeringConnectionAccepterState, opts ...pulumi.ResourceOpt) (*VpcPeeringConnectionAccepter, error) {
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
		inputs["vpcPeeringConnectionId"] = state.VpcPeeringConnectionId
	}
	var resource VpcPeeringConnectionAccepter
	err := ctx.ReadResource("aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *VpcPeeringConnectionAccepter) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *VpcPeeringConnectionAccepter) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering VpcPeeringConnectionAccepter resources.
type VpcPeeringConnectionAccepterState struct {
	// The status of the VPC Peering Connection request.
	AcceptStatus pulumi.StringInput `pulumi:"acceptStatus"`
	// A configuration block that describes [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the accepter VPC.
	Accepter pulumi.AnyInput `pulumi:"accepter"`
	// Whether or not to accept the peering request. Defaults to `false`.
	AutoAccept pulumi.BoolInput `pulumi:"autoAccept"`
	// The AWS account ID of the owner of the requester VPC.
	PeerOwnerId pulumi.StringInput `pulumi:"peerOwnerId"`
	// The region of the accepter VPC.
	PeerRegion pulumi.StringInput `pulumi:"peerRegion"`
	// The ID of the requester VPC.
	PeerVpcId pulumi.StringInput `pulumi:"peerVpcId"`
	// A configuration block that describes [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the requester VPC.
	Requester pulumi.AnyInput `pulumi:"requester"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The ID of the accepter VPC.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
	// The VPC Peering Connection ID to manage.
	VpcPeeringConnectionId pulumi.StringInput `pulumi:"vpcPeeringConnectionId"`
}

// The set of arguments for constructing a VpcPeeringConnectionAccepter resource.
type VpcPeeringConnectionAccepterArgs struct {
	// A configuration block that describes [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the accepter VPC.
	Accepter pulumi.AnyInput `pulumi:"accepter"`
	// Whether or not to accept the peering request. Defaults to `false`.
	AutoAccept pulumi.BoolInput `pulumi:"autoAccept"`
	// A configuration block that describes [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the requester VPC.
	Requester pulumi.AnyInput `pulumi:"requester"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The VPC Peering Connection ID to manage.
	VpcPeeringConnectionId pulumi.StringInput `pulumi:"vpcPeeringConnectionId"`
}
