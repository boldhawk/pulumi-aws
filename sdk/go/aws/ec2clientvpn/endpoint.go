// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2clientvpn

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS Client VPN endpoint for OpenVPN clients. For more information on usage, please see the
// [AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_client_vpn_endpoint.html.markdown.
type Endpoint struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// Information about the authentication method to be used to authenticate clients.
	AuthenticationOptions pulumi.AnyOutput `pulumi:"authenticationOptions"`

	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
	ClientCidrBlock pulumi.StringOutput `pulumi:"clientCidrBlock"`

	// Information about the client connection logging options.
	ConnectionLogOptions pulumi.AnyOutput `pulumi:"connectionLogOptions"`

	// Name of the repository.
	Description pulumi.StringOutput `pulumi:"description"`

	// The DNS name to be used by clients when establishing their VPN session.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`

	// Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
	DnsServers pulumi.ArrayOutput `pulumi:"dnsServers"`

	// The ARN of the ACM server certificate.
	ServerCertificateArn pulumi.StringOutput `pulumi:"serverCertificateArn"`

	// Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
	SplitTunnel pulumi.BoolOutput `pulumi:"splitTunnel"`

	// The current state of the Client VPN endpoint.
	Status pulumi.StringOutput `pulumi:"status"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// The transport protocol to be used by the VPN session. Default value is `udp`.
	TransportProtocol pulumi.StringOutput `pulumi:"transportProtocol"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOpt) (*Endpoint, error) {
	if args == nil || args.AuthenticationOptions == nil {
		return nil, errors.New("missing required argument 'AuthenticationOptions'")
	}
	if args == nil || args.ClientCidrBlock == nil {
		return nil, errors.New("missing required argument 'ClientCidrBlock'")
	}
	if args == nil || args.ConnectionLogOptions == nil {
		return nil, errors.New("missing required argument 'ConnectionLogOptions'")
	}
	if args == nil || args.ServerCertificateArn == nil {
		return nil, errors.New("missing required argument 'ServerCertificateArn'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["authenticationOptions"] = args.AuthenticationOptions
		inputs["clientCidrBlock"] = args.ClientCidrBlock
		inputs["connectionLogOptions"] = args.ConnectionLogOptions
		inputs["description"] = args.Description
		inputs["dnsServers"] = args.DnsServers
		inputs["serverCertificateArn"] = args.ServerCertificateArn
		inputs["splitTunnel"] = args.SplitTunnel
		inputs["tags"] = args.Tags
		inputs["transportProtocol"] = args.TransportProtocol
	}
	var resource Endpoint
	err := ctx.RegisterResource("aws:ec2clientvpn/endpoint:Endpoint", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.ID, state *EndpointState, opts ...pulumi.ResourceOpt) (*Endpoint, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["authenticationOptions"] = state.AuthenticationOptions
		inputs["clientCidrBlock"] = state.ClientCidrBlock
		inputs["connectionLogOptions"] = state.ConnectionLogOptions
		inputs["description"] = state.Description
		inputs["dnsName"] = state.DnsName
		inputs["dnsServers"] = state.DnsServers
		inputs["serverCertificateArn"] = state.ServerCertificateArn
		inputs["splitTunnel"] = state.SplitTunnel
		inputs["status"] = state.Status
		inputs["tags"] = state.Tags
		inputs["transportProtocol"] = state.TransportProtocol
	}
	var resource Endpoint
	err := ctx.ReadResource("aws:ec2clientvpn/endpoint:Endpoint", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *Endpoint) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *Endpoint) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering Endpoint resources.
type EndpointState struct {
	// Information about the authentication method to be used to authenticate clients.
	AuthenticationOptions pulumi.AnyInput `pulumi:"authenticationOptions"`
	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
	ClientCidrBlock pulumi.StringInput `pulumi:"clientCidrBlock"`
	// Information about the client connection logging options.
	ConnectionLogOptions pulumi.AnyInput `pulumi:"connectionLogOptions"`
	// Name of the repository.
	Description pulumi.StringInput `pulumi:"description"`
	// The DNS name to be used by clients when establishing their VPN session.
	DnsName pulumi.StringInput `pulumi:"dnsName"`
	// Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
	DnsServers pulumi.ArrayInput `pulumi:"dnsServers"`
	// The ARN of the ACM server certificate.
	ServerCertificateArn pulumi.StringInput `pulumi:"serverCertificateArn"`
	// Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
	SplitTunnel pulumi.BoolInput `pulumi:"splitTunnel"`
	// The current state of the Client VPN endpoint.
	Status pulumi.StringInput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The transport protocol to be used by the VPN session. Default value is `udp`.
	TransportProtocol pulumi.StringInput `pulumi:"transportProtocol"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// Information about the authentication method to be used to authenticate clients.
	AuthenticationOptions pulumi.AnyInput `pulumi:"authenticationOptions"`
	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
	ClientCidrBlock pulumi.StringInput `pulumi:"clientCidrBlock"`
	// Information about the client connection logging options.
	ConnectionLogOptions pulumi.AnyInput `pulumi:"connectionLogOptions"`
	// Name of the repository.
	Description pulumi.StringInput `pulumi:"description"`
	// Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
	DnsServers pulumi.ArrayInput `pulumi:"dnsServers"`
	// The ARN of the ACM server certificate.
	ServerCertificateArn pulumi.StringInput `pulumi:"serverCertificateArn"`
	// Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
	SplitTunnel pulumi.BoolInput `pulumi:"splitTunnel"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The transport protocol to be used by the VPN session. Default value is `udp`.
	TransportProtocol pulumi.StringInput `pulumi:"transportProtocol"`
}
