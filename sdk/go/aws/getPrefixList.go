// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `.getPrefixList` provides details about a specific prefix list (PL)
// in the current region.
// 
// This can be used both to validate a prefix list given in a variable
// and to obtain the CIDR blocks (IP address ranges) for the associated
// AWS service. The latter may be useful e.g. for adding network ACL
// rules.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/prefix_list.html.markdown.
func LookupPrefixList(ctx *pulumi.Context, args *GetPrefixListArgs) (*GetPrefixListResult, error) {
	var rv GetPrefixListResult
	err := ctx.Invoke("aws:index/getPrefixList:getPrefixList", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrefixList.
type GetPrefixListArgs struct {
	// The name of the prefix list to select.
	Name string `pulumi:"name"`
	// The ID of the prefix list to select.
	PrefixListId string `pulumi:"prefixListId"`
}

// A collection of values returned by getPrefixList.
type GetPrefixListResult struct {
	// The list of CIDR blocks for the AWS service associated
	// with the prefix list.
	CidrBlocks []string `pulumi:"cidrBlocks"`
	// The name of the selected prefix list.
	Name string `pulumi:"name"`
	PrefixListId string `pulumi:"prefixListId"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
