// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `ec2.Vpc` provides details about a specific VPC.
// 
// This resource can prove useful when a module accepts a vpc id as
// an input variable and needs to, for example, determine the CIDR block of that
// VPC.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/vpc.html.markdown.
func LookupVpc(ctx *pulumi.Context, args *GetVpcArgs) (*GetVpcResult, error) {
var rv GetVpcResult
	err := ctx.Invoke("aws:ec2/getVpc:getVpc", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpc.
type GetVpcArgs struct {
	// The cidr block of the desired VPC.
	CidrBlock string `pulumi:"cidrBlock"`
	// Boolean constraint on whether the desired VPC is
	// the default VPC for the region.
	Default bool `pulumi:"default"`
	// The DHCP options id of the desired VPC.
	DhcpOptionsId string `pulumi:"dhcpOptionsId"`
	// Custom filter block as described below.
	Filters []interface{} `pulumi:"filters"`
	// The id of the specific VPC to retrieve.
	Id string `pulumi:"id"`
	// The current state of the desired VPC.
	// Can be either `"pending"` or `"available"`.
	State string `pulumi:"state"`
	// A mapping of tags, each pair of which must exactly match
	// a pair on the desired VPC.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getVpc.
type GetVpcResult struct {
	// Amazon Resource Name (ARN) of VPC
	Arn string `pulumi:"arn"`
	// The CIDR block for the association.
	CidrBlock string `pulumi:"cidrBlock"`
	CidrBlockAssociations []interface{} `pulumi:"cidrBlockAssociations"`
	Default bool `pulumi:"default"`
	DhcpOptionsId string `pulumi:"dhcpOptionsId"`
	// Whether or not the VPC has DNS hostname support
	EnableDnsHostnames bool `pulumi:"enableDnsHostnames"`
	// Whether or not the VPC has DNS support
	EnableDnsSupport bool `pulumi:"enableDnsSupport"`
	Filters []interface{} `pulumi:"filters"`
	Id string `pulumi:"id"`
	// The allowed tenancy of instances launched into the
	// selected VPC. May be any of `"default"`, `"dedicated"`, or `"host"`.
	InstanceTenancy string `pulumi:"instanceTenancy"`
	// The association ID for the IPv6 CIDR block.
	Ipv6AssociationId string `pulumi:"ipv6AssociationId"`
	// The IPv6 CIDR block.
	Ipv6CidrBlock string `pulumi:"ipv6CidrBlock"`
	// The ID of the main route table associated with this VPC.
	MainRouteTableId string `pulumi:"mainRouteTableId"`
	// The ID of the AWS account that owns the VPC.
	OwnerId string `pulumi:"ownerId"`
	// The State of the association.
	State string `pulumi:"state"`
	Tags map[string]interface{} `pulumi:"tags"`
}
