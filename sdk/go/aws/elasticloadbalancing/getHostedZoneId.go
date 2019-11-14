// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancing

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the HostedZoneId of the AWS Elastic Load Balancing HostedZoneId
// in a given region for the purpose of using in an AWS Route53 Alias.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elb_hosted_zone_id_legacy.html.markdown.
func LookupHostedZoneId(ctx *pulumi.Context, args *GetHostedZoneIdArgs) (*GetHostedZoneIdResult, error) {
	var rv GetHostedZoneIdResult
	err := ctx.Invoke("aws:elasticloadbalancing/getHostedZoneId:getHostedZoneId", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHostedZoneId.
type GetHostedZoneIdArgs struct {
	// Name of the region whose AWS ELB HostedZoneId is desired.
	// Defaults to the region from the AWS provider configuration.
	Region string `pulumi:"region"`
}

// A collection of values returned by getHostedZoneId.
type GetHostedZoneIdResult struct {
	Region string `pulumi:"region"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
