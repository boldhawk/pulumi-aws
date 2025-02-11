// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Direct Connect public virtual interface resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dx_public_virtual_interface.html.markdown.
type PublicVirtualInterface struct {
	s *pulumi.ResourceState
}

// NewPublicVirtualInterface registers a new resource with the given unique name, arguments, and options.
func NewPublicVirtualInterface(ctx *pulumi.Context,
	name string, args *PublicVirtualInterfaceArgs, opts ...pulumi.ResourceOpt) (*PublicVirtualInterface, error) {
	if args == nil || args.AddressFamily == nil {
		return nil, errors.New("missing required argument 'AddressFamily'")
	}
	if args == nil || args.BgpAsn == nil {
		return nil, errors.New("missing required argument 'BgpAsn'")
	}
	if args == nil || args.ConnectionId == nil {
		return nil, errors.New("missing required argument 'ConnectionId'")
	}
	if args == nil || args.RouteFilterPrefixes == nil {
		return nil, errors.New("missing required argument 'RouteFilterPrefixes'")
	}
	if args == nil || args.Vlan == nil {
		return nil, errors.New("missing required argument 'Vlan'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["addressFamily"] = nil
		inputs["amazonAddress"] = nil
		inputs["bgpAsn"] = nil
		inputs["bgpAuthKey"] = nil
		inputs["connectionId"] = nil
		inputs["customerAddress"] = nil
		inputs["name"] = nil
		inputs["routeFilterPrefixes"] = nil
		inputs["tags"] = nil
		inputs["vlan"] = nil
	} else {
		inputs["addressFamily"] = args.AddressFamily
		inputs["amazonAddress"] = args.AmazonAddress
		inputs["bgpAsn"] = args.BgpAsn
		inputs["bgpAuthKey"] = args.BgpAuthKey
		inputs["connectionId"] = args.ConnectionId
		inputs["customerAddress"] = args.CustomerAddress
		inputs["name"] = args.Name
		inputs["routeFilterPrefixes"] = args.RouteFilterPrefixes
		inputs["tags"] = args.Tags
		inputs["vlan"] = args.Vlan
	}
	inputs["arn"] = nil
	inputs["awsDevice"] = nil
	s, err := ctx.RegisterResource("aws:directconnect/publicVirtualInterface:PublicVirtualInterface", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PublicVirtualInterface{s: s}, nil
}

// GetPublicVirtualInterface gets an existing PublicVirtualInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicVirtualInterface(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PublicVirtualInterfaceState, opts ...pulumi.ResourceOpt) (*PublicVirtualInterface, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["addressFamily"] = state.AddressFamily
		inputs["amazonAddress"] = state.AmazonAddress
		inputs["arn"] = state.Arn
		inputs["awsDevice"] = state.AwsDevice
		inputs["bgpAsn"] = state.BgpAsn
		inputs["bgpAuthKey"] = state.BgpAuthKey
		inputs["connectionId"] = state.ConnectionId
		inputs["customerAddress"] = state.CustomerAddress
		inputs["name"] = state.Name
		inputs["routeFilterPrefixes"] = state.RouteFilterPrefixes
		inputs["tags"] = state.Tags
		inputs["vlan"] = state.Vlan
	}
	s, err := ctx.ReadResource("aws:directconnect/publicVirtualInterface:PublicVirtualInterface", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PublicVirtualInterface{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *PublicVirtualInterface) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *PublicVirtualInterface) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The address family for the BGP peer. `ipv4 ` or `ipv6`.
func (r *PublicVirtualInterface) AddressFamily() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["addressFamily"])
}

// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
func (r *PublicVirtualInterface) AmazonAddress() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["amazonAddress"])
}

// The ARN of the virtual interface.
func (r *PublicVirtualInterface) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// The Direct Connect endpoint on which the virtual interface terminates.
func (r *PublicVirtualInterface) AwsDevice() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["awsDevice"])
}

// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
func (r *PublicVirtualInterface) BgpAsn() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["bgpAsn"])
}

// The authentication key for BGP configuration.
func (r *PublicVirtualInterface) BgpAuthKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["bgpAuthKey"])
}

// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
func (r *PublicVirtualInterface) ConnectionId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["connectionId"])
}

// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
func (r *PublicVirtualInterface) CustomerAddress() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["customerAddress"])
}

// The name for the virtual interface.
func (r *PublicVirtualInterface) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// A list of routes to be advertised to the AWS network in this region.
func (r *PublicVirtualInterface) RouteFilterPrefixes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["routeFilterPrefixes"])
}

// A mapping of tags to assign to the resource.
func (r *PublicVirtualInterface) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// The VLAN ID.
func (r *PublicVirtualInterface) Vlan() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["vlan"])
}

// Input properties used for looking up and filtering PublicVirtualInterface resources.
type PublicVirtualInterfaceState struct {
	// The address family for the BGP peer. `ipv4 ` or `ipv6`.
	AddressFamily interface{}
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress interface{}
	// The ARN of the virtual interface.
	Arn interface{}
	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDevice interface{}
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn interface{}
	// The authentication key for BGP configuration.
	BgpAuthKey interface{}
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId interface{}
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress interface{}
	// The name for the virtual interface.
	Name interface{}
	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The VLAN ID.
	Vlan interface{}
}

// The set of arguments for constructing a PublicVirtualInterface resource.
type PublicVirtualInterfaceArgs struct {
	// The address family for the BGP peer. `ipv4 ` or `ipv6`.
	AddressFamily interface{}
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress interface{}
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn interface{}
	// The authentication key for BGP configuration.
	BgpAuthKey interface{}
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId interface{}
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress interface{}
	// The name for the virtual interface.
	Name interface{}
	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The VLAN ID.
	Vlan interface{}
}
