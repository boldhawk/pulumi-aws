// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an EC2 VPN connection. These objects can be connected to customer gateways, and allow you to establish tunnels between your network and Amazon.
// 
// > **Note:** All arguments including `tunnel1PresharedKey` and `tunnel2PresharedKey` will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
// 
// > **Note:** The CIDR blocks in the arguments `tunnel1InsideCidr` and `tunnel2InsideCidr` must have a prefix of /30 and be a part of a specific range.
// [Read more about this in the AWS documentation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpnTunnelOptionsSpecification.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpn_connection.html.markdown.
type VpnConnection struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The configuration information for the VPN connection's customer gateway (in the native XML format).
	CustomerGatewayConfiguration pulumi.StringOutput `pulumi:"customerGatewayConfiguration"`

	// The ID of the customer gateway.
	CustomerGatewayId pulumi.StringOutput `pulumi:"customerGatewayId"`

	Routes pulumi.ArrayOutput `pulumi:"routes"`

	// Whether the VPN connection uses static routes exclusively. Static routes must be used for devices that don't support BGP.
	StaticRoutesOnly pulumi.BoolOutput `pulumi:"staticRoutesOnly"`

	// Tags to apply to the connection.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// When associated with an EC2 Transit Gateway (`transitGatewayId` argument), the attachment ID.
	TransitGatewayAttachmentId pulumi.StringOutput `pulumi:"transitGatewayAttachmentId"`

	// The ID of the EC2 Transit Gateway.
	TransitGatewayId pulumi.StringOutput `pulumi:"transitGatewayId"`

	// The public IP address of the first VPN tunnel.
	Tunnel1Address pulumi.StringOutput `pulumi:"tunnel1Address"`

	// The bgp asn number of the first VPN tunnel.
	Tunnel1BgpAsn pulumi.StringOutput `pulumi:"tunnel1BgpAsn"`

	// The bgp holdtime of the first VPN tunnel.
	Tunnel1BgpHoldtime pulumi.IntOutput `pulumi:"tunnel1BgpHoldtime"`

	// The RFC 6890 link-local address of the first VPN tunnel (Customer Gateway Side).
	Tunnel1CgwInsideAddress pulumi.StringOutput `pulumi:"tunnel1CgwInsideAddress"`

	// The CIDR block of the inside IP addresses for the first VPN tunnel.
	Tunnel1InsideCidr pulumi.StringOutput `pulumi:"tunnel1InsideCidr"`

	// The preshared key of the first VPN tunnel.
	Tunnel1PresharedKey pulumi.StringOutput `pulumi:"tunnel1PresharedKey"`

	// The RFC 6890 link-local address of the first VPN tunnel (VPN Gateway Side).
	Tunnel1VgwInsideAddress pulumi.StringOutput `pulumi:"tunnel1VgwInsideAddress"`

	// The public IP address of the second VPN tunnel.
	Tunnel2Address pulumi.StringOutput `pulumi:"tunnel2Address"`

	// The bgp asn number of the second VPN tunnel.
	Tunnel2BgpAsn pulumi.StringOutput `pulumi:"tunnel2BgpAsn"`

	// The bgp holdtime of the second VPN tunnel.
	Tunnel2BgpHoldtime pulumi.IntOutput `pulumi:"tunnel2BgpHoldtime"`

	// The RFC 6890 link-local address of the second VPN tunnel (Customer Gateway Side).
	Tunnel2CgwInsideAddress pulumi.StringOutput `pulumi:"tunnel2CgwInsideAddress"`

	// The CIDR block of the inside IP addresses for the second VPN tunnel.
	Tunnel2InsideCidr pulumi.StringOutput `pulumi:"tunnel2InsideCidr"`

	// The preshared key of the second VPN tunnel.
	Tunnel2PresharedKey pulumi.StringOutput `pulumi:"tunnel2PresharedKey"`

	// The RFC 6890 link-local address of the second VPN tunnel (VPN Gateway Side).
	Tunnel2VgwInsideAddress pulumi.StringOutput `pulumi:"tunnel2VgwInsideAddress"`

	// The type of VPN connection. The only type AWS supports at this time is "ipsec.1".
	Type pulumi.StringOutput `pulumi:"type"`

	VgwTelemetries pulumi.ArrayOutput `pulumi:"vgwTelemetries"`

	// The ID of the Virtual Private Gateway.
	VpnGatewayId pulumi.StringOutput `pulumi:"vpnGatewayId"`
}

// NewVpnConnection registers a new resource with the given unique name, arguments, and options.
func NewVpnConnection(ctx *pulumi.Context,
	name string, args *VpnConnectionArgs, opts ...pulumi.ResourceOpt) (*VpnConnection, error) {
	if args == nil || args.CustomerGatewayId == nil {
		return nil, errors.New("missing required argument 'CustomerGatewayId'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["customerGatewayId"] = args.CustomerGatewayId
		inputs["staticRoutesOnly"] = args.StaticRoutesOnly
		inputs["tags"] = args.Tags
		inputs["transitGatewayId"] = args.TransitGatewayId
		inputs["tunnel1InsideCidr"] = args.Tunnel1InsideCidr
		inputs["tunnel1PresharedKey"] = args.Tunnel1PresharedKey
		inputs["tunnel2InsideCidr"] = args.Tunnel2InsideCidr
		inputs["tunnel2PresharedKey"] = args.Tunnel2PresharedKey
		inputs["type"] = args.Type
		inputs["vpnGatewayId"] = args.VpnGatewayId
	}
	var resource VpnConnection
	err := ctx.RegisterResource("aws:ec2/vpnConnection:VpnConnection", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnConnection gets an existing VpnConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnConnection(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VpnConnectionState, opts ...pulumi.ResourceOpt) (*VpnConnection, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["customerGatewayConfiguration"] = state.CustomerGatewayConfiguration
		inputs["customerGatewayId"] = state.CustomerGatewayId
		inputs["routes"] = state.Routes
		inputs["staticRoutesOnly"] = state.StaticRoutesOnly
		inputs["tags"] = state.Tags
		inputs["transitGatewayAttachmentId"] = state.TransitGatewayAttachmentId
		inputs["transitGatewayId"] = state.TransitGatewayId
		inputs["tunnel1Address"] = state.Tunnel1Address
		inputs["tunnel1BgpAsn"] = state.Tunnel1BgpAsn
		inputs["tunnel1BgpHoldtime"] = state.Tunnel1BgpHoldtime
		inputs["tunnel1CgwInsideAddress"] = state.Tunnel1CgwInsideAddress
		inputs["tunnel1InsideCidr"] = state.Tunnel1InsideCidr
		inputs["tunnel1PresharedKey"] = state.Tunnel1PresharedKey
		inputs["tunnel1VgwInsideAddress"] = state.Tunnel1VgwInsideAddress
		inputs["tunnel2Address"] = state.Tunnel2Address
		inputs["tunnel2BgpAsn"] = state.Tunnel2BgpAsn
		inputs["tunnel2BgpHoldtime"] = state.Tunnel2BgpHoldtime
		inputs["tunnel2CgwInsideAddress"] = state.Tunnel2CgwInsideAddress
		inputs["tunnel2InsideCidr"] = state.Tunnel2InsideCidr
		inputs["tunnel2PresharedKey"] = state.Tunnel2PresharedKey
		inputs["tunnel2VgwInsideAddress"] = state.Tunnel2VgwInsideAddress
		inputs["type"] = state.Type
		inputs["vgwTelemetries"] = state.VgwTelemetries
		inputs["vpnGatewayId"] = state.VpnGatewayId
	}
	var resource VpnConnection
	err := ctx.ReadResource("aws:ec2/vpnConnection:VpnConnection", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *VpnConnection) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *VpnConnection) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering VpnConnection resources.
type VpnConnectionState struct {
	// The configuration information for the VPN connection's customer gateway (in the native XML format).
	CustomerGatewayConfiguration pulumi.StringInput `pulumi:"customerGatewayConfiguration"`
	// The ID of the customer gateway.
	CustomerGatewayId pulumi.StringInput `pulumi:"customerGatewayId"`
	Routes pulumi.ArrayInput `pulumi:"routes"`
	// Whether the VPN connection uses static routes exclusively. Static routes must be used for devices that don't support BGP.
	StaticRoutesOnly pulumi.BoolInput `pulumi:"staticRoutesOnly"`
	// Tags to apply to the connection.
	Tags pulumi.MapInput `pulumi:"tags"`
	// When associated with an EC2 Transit Gateway (`transitGatewayId` argument), the attachment ID.
	TransitGatewayAttachmentId pulumi.StringInput `pulumi:"transitGatewayAttachmentId"`
	// The ID of the EC2 Transit Gateway.
	TransitGatewayId pulumi.StringInput `pulumi:"transitGatewayId"`
	// The public IP address of the first VPN tunnel.
	Tunnel1Address pulumi.StringInput `pulumi:"tunnel1Address"`
	// The bgp asn number of the first VPN tunnel.
	Tunnel1BgpAsn pulumi.StringInput `pulumi:"tunnel1BgpAsn"`
	// The bgp holdtime of the first VPN tunnel.
	Tunnel1BgpHoldtime pulumi.IntInput `pulumi:"tunnel1BgpHoldtime"`
	// The RFC 6890 link-local address of the first VPN tunnel (Customer Gateway Side).
	Tunnel1CgwInsideAddress pulumi.StringInput `pulumi:"tunnel1CgwInsideAddress"`
	// The CIDR block of the inside IP addresses for the first VPN tunnel.
	Tunnel1InsideCidr pulumi.StringInput `pulumi:"tunnel1InsideCidr"`
	// The preshared key of the first VPN tunnel.
	Tunnel1PresharedKey pulumi.StringInput `pulumi:"tunnel1PresharedKey"`
	// The RFC 6890 link-local address of the first VPN tunnel (VPN Gateway Side).
	Tunnel1VgwInsideAddress pulumi.StringInput `pulumi:"tunnel1VgwInsideAddress"`
	// The public IP address of the second VPN tunnel.
	Tunnel2Address pulumi.StringInput `pulumi:"tunnel2Address"`
	// The bgp asn number of the second VPN tunnel.
	Tunnel2BgpAsn pulumi.StringInput `pulumi:"tunnel2BgpAsn"`
	// The bgp holdtime of the second VPN tunnel.
	Tunnel2BgpHoldtime pulumi.IntInput `pulumi:"tunnel2BgpHoldtime"`
	// The RFC 6890 link-local address of the second VPN tunnel (Customer Gateway Side).
	Tunnel2CgwInsideAddress pulumi.StringInput `pulumi:"tunnel2CgwInsideAddress"`
	// The CIDR block of the inside IP addresses for the second VPN tunnel.
	Tunnel2InsideCidr pulumi.StringInput `pulumi:"tunnel2InsideCidr"`
	// The preshared key of the second VPN tunnel.
	Tunnel2PresharedKey pulumi.StringInput `pulumi:"tunnel2PresharedKey"`
	// The RFC 6890 link-local address of the second VPN tunnel (VPN Gateway Side).
	Tunnel2VgwInsideAddress pulumi.StringInput `pulumi:"tunnel2VgwInsideAddress"`
	// The type of VPN connection. The only type AWS supports at this time is "ipsec.1".
	Type pulumi.StringInput `pulumi:"type"`
	VgwTelemetries pulumi.ArrayInput `pulumi:"vgwTelemetries"`
	// The ID of the Virtual Private Gateway.
	VpnGatewayId pulumi.StringInput `pulumi:"vpnGatewayId"`
}

// The set of arguments for constructing a VpnConnection resource.
type VpnConnectionArgs struct {
	// The ID of the customer gateway.
	CustomerGatewayId pulumi.StringInput `pulumi:"customerGatewayId"`
	// Whether the VPN connection uses static routes exclusively. Static routes must be used for devices that don't support BGP.
	StaticRoutesOnly pulumi.BoolInput `pulumi:"staticRoutesOnly"`
	// Tags to apply to the connection.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The ID of the EC2 Transit Gateway.
	TransitGatewayId pulumi.StringInput `pulumi:"transitGatewayId"`
	// The CIDR block of the inside IP addresses for the first VPN tunnel.
	Tunnel1InsideCidr pulumi.StringInput `pulumi:"tunnel1InsideCidr"`
	// The preshared key of the first VPN tunnel.
	Tunnel1PresharedKey pulumi.StringInput `pulumi:"tunnel1PresharedKey"`
	// The CIDR block of the inside IP addresses for the second VPN tunnel.
	Tunnel2InsideCidr pulumi.StringInput `pulumi:"tunnel2InsideCidr"`
	// The preshared key of the second VPN tunnel.
	Tunnel2PresharedKey pulumi.StringInput `pulumi:"tunnel2PresharedKey"`
	// The type of VPN connection. The only type AWS supports at this time is "ipsec.1".
	Type pulumi.StringInput `pulumi:"type"`
	// The ID of the Virtual Private Gateway.
	VpnGatewayId pulumi.StringInput `pulumi:"vpnGatewayId"`
}
