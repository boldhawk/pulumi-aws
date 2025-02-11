// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get IDs and VPC membership of Security Groups that are created
// outside of this provider.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/security_groups.html.markdown.
func LookupSecurityGroups(ctx *pulumi.Context, args *GetSecurityGroupsArgs) (*GetSecurityGroupsResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["filters"] = args.Filters
		inputs["tags"] = args.Tags
	}
	outputs, err := ctx.Invoke("aws:ec2/getSecurityGroups:getSecurityGroups", inputs)
	if err != nil {
		return nil, err
	}
	return &GetSecurityGroupsResult{
		Filters: outputs["filters"],
		Ids: outputs["ids"],
		Tags: outputs["tags"],
		VpcIds: outputs["vpcIds"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getSecurityGroups.
type GetSecurityGroupsArgs struct {
	// One or more name/value pairs to use as filters. There are
	// several valid keys, for a full reference, check out
	// [describe-security-groups in the AWS CLI reference][1].
	Filters interface{}
	// A mapping of tags, each pair of which must exactly match for
	// desired security groups.
	Tags interface{}
}

// A collection of values returned by getSecurityGroups.
type GetSecurityGroupsResult struct {
	Filters interface{}
	// IDs of the matches security groups.
	Ids interface{}
	Tags interface{}
	// The VPC IDs of the matched security groups. The data source's tag or filter *will span VPCs*
	// unless the `vpc-id` filter is also used.
	VpcIds interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
