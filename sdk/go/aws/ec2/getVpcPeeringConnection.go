// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The VPC Peering Connection data source provides details about
// a specific VPC peering connection.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/vpc_peering_connection.html.markdown.
func LookupVpcPeeringConnection(ctx *pulumi.Context, args *GetVpcPeeringConnectionArgs) (*GetVpcPeeringConnectionResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["cidrBlock"] = args.CidrBlock
		inputs["filters"] = args.Filters
		inputs["id"] = args.Id
		inputs["ownerId"] = args.OwnerId
		inputs["peerCidrBlock"] = args.PeerCidrBlock
		inputs["peerOwnerId"] = args.PeerOwnerId
		inputs["peerRegion"] = args.PeerRegion
		inputs["peerVpcId"] = args.PeerVpcId
		inputs["region"] = args.Region
		inputs["status"] = args.Status
		inputs["tags"] = args.Tags
		inputs["vpcId"] = args.VpcId
	}
	outputs, err := ctx.Invoke("aws:ec2/getVpcPeeringConnection:getVpcPeeringConnection", inputs)
	if err != nil {
		return nil, err
	}
	return &GetVpcPeeringConnectionResult{
		Accepter: outputs["accepter"],
		CidrBlock: outputs["cidrBlock"],
		Filters: outputs["filters"],
		Id: outputs["id"],
		OwnerId: outputs["ownerId"],
		PeerCidrBlock: outputs["peerCidrBlock"],
		PeerOwnerId: outputs["peerOwnerId"],
		PeerRegion: outputs["peerRegion"],
		PeerVpcId: outputs["peerVpcId"],
		Region: outputs["region"],
		Requester: outputs["requester"],
		Status: outputs["status"],
		Tags: outputs["tags"],
		VpcId: outputs["vpcId"],
	}, nil
}

// A collection of arguments for invoking getVpcPeeringConnection.
type GetVpcPeeringConnectionArgs struct {
	// The CIDR block of the requester VPC of the specific VPC Peering Connection to retrieve.
	CidrBlock pulumi.StringInput `pulumi:"cidrBlock"`
	// Custom filter block as described below.
	Filters pulumi.ArrayInput `pulumi:"filters"`
	// The ID of the specific VPC Peering Connection to retrieve.
	Id pulumi.StringInput `pulumi:"id"`
	// The AWS account ID of the owner of the requester VPC of the specific VPC Peering Connection to retrieve.
	OwnerId pulumi.StringInput `pulumi:"ownerId"`
	// The CIDR block of the accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerCidrBlock pulumi.StringInput `pulumi:"peerCidrBlock"`
	// The AWS account ID of the owner of the accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerOwnerId pulumi.StringInput `pulumi:"peerOwnerId"`
	// The region of the accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerRegion pulumi.StringInput `pulumi:"peerRegion"`
	// The ID of the accepter VPC of the specific VPC Peering Connection to retrieve.
	PeerVpcId pulumi.StringInput `pulumi:"peerVpcId"`
	// The region of the requester VPC of the specific VPC Peering Connection to retrieve.
	Region pulumi.StringInput `pulumi:"region"`
	// The status of the specific VPC Peering Connection to retrieve.
	Status pulumi.StringInput `pulumi:"status"`
	// A mapping of tags, each pair of which must exactly match
	// a pair on the desired VPC Peering Connection.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The ID of the requester VPC of the specific VPC Peering Connection to retrieve.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

// A collection of values returned by getVpcPeeringConnection.
type GetVpcPeeringConnectionResult struct {
	// A configuration block that describes [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the accepter VPC.
	Accepter map[string]interface{} `pulumi:"accepter"`
	CidrBlock string `pulumi:"cidrBlock"`
	Filters []interface{} `pulumi:"filters"`
	Id string `pulumi:"id"`
	OwnerId string `pulumi:"ownerId"`
	PeerCidrBlock string `pulumi:"peerCidrBlock"`
	PeerOwnerId string `pulumi:"peerOwnerId"`
	PeerRegion string `pulumi:"peerRegion"`
	PeerVpcId string `pulumi:"peerVpcId"`
	Region string `pulumi:"region"`
	// A configuration block that describes [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the requester VPC.
	Requester map[string]interface{} `pulumi:"requester"`
	Status string `pulumi:"status"`
	Tags map[string]interface{} `pulumi:"tags"`
	VpcId string `pulumi:"vpcId"`
}
