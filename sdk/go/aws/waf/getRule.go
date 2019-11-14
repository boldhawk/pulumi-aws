// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `waf.Rule` Retrieves a WAF Rule Resource Id.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/waf_rule.html.markdown.
func LookupRule(ctx *pulumi.Context, args *GetRuleArgs) (*GetRuleResult, error) {
var rv GetRuleResult
	err := ctx.Invoke("aws:waf/getRule:getRule", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRule.
type GetRuleArgs struct {
	// The name of the WAF rule.
	Name string `pulumi:"name"`
}

// A collection of values returned by getRule.
type GetRuleResult struct {
	Name string `pulumi:"name"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
