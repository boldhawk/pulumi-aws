// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `wafregional.IpSet` Retrieves a WAF Regional IP Set Resource Id.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/wafregional_ipset.html.markdown.
func LookupIpset(ctx *pulumi.Context, args *GetIpsetArgs) (*GetIpsetResult, error) {
	var rv GetIpsetResult
	err := ctx.Invoke("aws:wafregional/getIpset:getIpset", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpset.
type GetIpsetArgs struct {
	// The name of the WAF Regional IP set.
	Name string `pulumi:"name"`
}

// A collection of values returned by getIpset.
type GetIpsetResult struct {
	Name string `pulumi:"name"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
