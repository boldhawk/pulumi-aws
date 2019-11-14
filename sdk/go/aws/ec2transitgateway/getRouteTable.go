// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get information on an EC2 Transit Gateway Route Table.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ec2_transit_gateway_route_table.html.markdown.
func LookupRouteTable(ctx *pulumi.Context, args *GetRouteTableArgs) (*GetRouteTableResult, error) {
	var rv GetRouteTableResult
	err := ctx.Invoke("aws:ec2transitgateway/getRouteTable:getRouteTable", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouteTable.
type GetRouteTableArgs struct {
	// One or more configuration blocks containing name-values filters. Detailed below.
	Filters []interface{} `pulumi:"filters"`
	// Identifier of the EC2 Transit Gateway Route Table.
	Id string `pulumi:"id"`
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getRouteTable.
type GetRouteTableResult struct {
	// Boolean whether this is the default association route table for the EC2 Transit Gateway
	DefaultAssociationRouteTable bool `pulumi:"defaultAssociationRouteTable"`
	// Boolean whether this is the default propagation route table for the EC2 Transit Gateway
	DefaultPropagationRouteTable bool `pulumi:"defaultPropagationRouteTable"`
	Filters []interface{} `pulumi:"filters"`
	// EC2 Transit Gateway Route Table identifier
	Id string `pulumi:"id"`
	// Key-value tags for the EC2 Transit Gateway Route Table
	Tags map[string]string `pulumi:"tags"`
	// EC2 Transit Gateway identifier
	TransitGatewayId string `pulumi:"transitGatewayId"`
}
