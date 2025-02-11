// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get information about a Network Interface.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/network_interface.html.markdown.
func LookupNetworkInterface(ctx *pulumi.Context, args *GetNetworkInterfaceArgs) (*GetNetworkInterfaceResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["filters"] = args.Filters
		inputs["id"] = args.Id
		inputs["tags"] = args.Tags
	}
	outputs, err := ctx.Invoke("aws:ec2/getNetworkInterface:getNetworkInterface", inputs)
	if err != nil {
		return nil, err
	}
	return &GetNetworkInterfaceResult{
		Associations: outputs["associations"],
		Attachments: outputs["attachments"],
		AvailabilityZone: outputs["availabilityZone"],
		Description: outputs["description"],
		Filters: outputs["filters"],
		Id: outputs["id"],
		InterfaceType: outputs["interfaceType"],
		Ipv6Addresses: outputs["ipv6Addresses"],
		MacAddress: outputs["macAddress"],
		OwnerId: outputs["ownerId"],
		PrivateDnsName: outputs["privateDnsName"],
		PrivateIp: outputs["privateIp"],
		PrivateIps: outputs["privateIps"],
		RequesterId: outputs["requesterId"],
		SecurityGroups: outputs["securityGroups"],
		SubnetId: outputs["subnetId"],
		Tags: outputs["tags"],
		VpcId: outputs["vpcId"],
	}, nil
}

// A collection of arguments for invoking getNetworkInterface.
type GetNetworkInterfaceArgs struct {
	// One or more name/value pairs to filter off of. There are several valid keys, for a full reference, check out [describe-network-interfaces](https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-network-interfaces.html) in the AWS CLI reference.
	Filters interface{}
	// The identifier for the network interface.
	Id interface{}
	Tags interface{}
}

// A collection of values returned by getNetworkInterface.
type GetNetworkInterfaceResult struct {
	// The association information for an Elastic IP address (IPv4) associated with the network interface. See supported fields below.
	Associations interface{}
	Attachments interface{}
	// The Availability Zone.
	AvailabilityZone interface{}
	// Description of the network interface.
	Description interface{}
	Filters interface{}
	Id interface{}
	// The type of interface.
	InterfaceType interface{}
	// List of IPv6 addresses to assign to the ENI.
	Ipv6Addresses interface{}
	// The MAC address.
	MacAddress interface{}
	// The AWS account ID of the owner of the network interface.
	OwnerId interface{}
	// The private DNS name.
	PrivateDnsName interface{}
	// The private IPv4 address of the network interface within the subnet.
	PrivateIp interface{}
	// The private IPv4 addresses associated with the network interface.
	PrivateIps interface{}
	// The ID of the entity that launched the instance on your behalf.
	RequesterId interface{}
	// The list of security groups for the network interface.
	SecurityGroups interface{}
	// The ID of the subnet.
	SubnetId interface{}
	// Any tags assigned to the network interface.
	Tags interface{}
	// The ID of the VPC.
	VpcId interface{}
}
