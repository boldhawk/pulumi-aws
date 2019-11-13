// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Retrieve information about an EC2 DHCP Options configuration.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/vpc_dhcp_options.html.markdown.
func LookupVpcDhcpOptions(ctx *pulumi.Context, args *GetVpcDhcpOptionsArgs) (*GetVpcDhcpOptionsResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["dhcpOptionsId"] = args.DhcpOptionsId
		inputs["filters"] = args.Filters
		inputs["tags"] = args.Tags
	}
	outputs, err := ctx.Invoke("aws:ec2/getVpcDhcpOptions:getVpcDhcpOptions", inputs)
	if err != nil {
		return nil, err
	}
	return &GetVpcDhcpOptionsResult{
		DhcpOptionsId: outputs["dhcpOptionsId"],
		DomainName: outputs["domainName"],
		DomainNameServers: outputs["domainNameServers"],
		Filters: outputs["filters"],
		NetbiosNameServers: outputs["netbiosNameServers"],
		NetbiosNodeType: outputs["netbiosNodeType"],
		NtpServers: outputs["ntpServers"],
		OwnerId: outputs["ownerId"],
		Tags: outputs["tags"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getVpcDhcpOptions.
type GetVpcDhcpOptionsArgs struct {
	// The EC2 DHCP Options ID.
	DhcpOptionsId pulumi.StringInput `pulumi:"dhcpOptionsId"`
	// List of custom filters as described below.
	Filters pulumi.ArrayInput `pulumi:"filters"`
	Tags pulumi.MapInput `pulumi:"tags"`
}

// A collection of values returned by getVpcDhcpOptions.
type GetVpcDhcpOptionsResult struct {
	// EC2 DHCP Options ID
	DhcpOptionsId string `pulumi:"dhcpOptionsId"`
	// The suffix domain name to used when resolving non Fully Qualified Domain Names. e.g. the `search` value in the `/etc/resolv.conf` file.
	DomainName string `pulumi:"domainName"`
	// List of name servers.
	DomainNameServers []interface{} `pulumi:"domainNameServers"`
	Filters []interface{} `pulumi:"filters"`
	// List of NETBIOS name servers.
	NetbiosNameServers []interface{} `pulumi:"netbiosNameServers"`
	// The NetBIOS node type (1, 2, 4, or 8). For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType string `pulumi:"netbiosNodeType"`
	// List of NTP servers.
	NtpServers []interface{} `pulumi:"ntpServers"`
	// The ID of the AWS account that owns the DHCP options set.
	OwnerId string `pulumi:"ownerId"`
	// A mapping of tags assigned to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
