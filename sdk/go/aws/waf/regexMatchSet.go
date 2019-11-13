// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a WAF Regex Match Set Resource
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/waf_regex_match_set.html.markdown.
type RegexMatchSet struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The name or description of the Regex Match Set.
	Name pulumi.StringOutput `pulumi:"name"`

	// The regular expression pattern that you want AWS WAF to search for in web requests,
	// the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples pulumi.ArrayOutput `pulumi:"regexMatchTuples"`
}

// NewRegexMatchSet registers a new resource with the given unique name, arguments, and options.
func NewRegexMatchSet(ctx *pulumi.Context,
	name string, args *RegexMatchSetArgs, opts ...pulumi.ResourceOpt) (*RegexMatchSet, error) {
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["name"] = args.Name
		inputs["regexMatchTuples"] = args.RegexMatchTuples
	}
	var resource RegexMatchSet
	err := ctx.RegisterResource("aws:waf/regexMatchSet:RegexMatchSet", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegexMatchSet gets an existing RegexMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegexMatchSet(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RegexMatchSetState, opts ...pulumi.ResourceOpt) (*RegexMatchSet, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["name"] = state.Name
		inputs["regexMatchTuples"] = state.RegexMatchTuples
	}
	var resource RegexMatchSet
	err := ctx.ReadResource("aws:waf/regexMatchSet:RegexMatchSet", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *RegexMatchSet) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *RegexMatchSet) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering RegexMatchSet resources.
type RegexMatchSetState struct {
	// The name or description of the Regex Match Set.
	Name pulumi.StringInput `pulumi:"name"`
	// The regular expression pattern that you want AWS WAF to search for in web requests,
	// the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples pulumi.ArrayInput `pulumi:"regexMatchTuples"`
}

// The set of arguments for constructing a RegexMatchSet resource.
type RegexMatchSetArgs struct {
	// The name or description of the Regex Match Set.
	Name pulumi.StringInput `pulumi:"name"`
	// The regular expression pattern that you want AWS WAF to search for in web requests,
	// the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples pulumi.ArrayInput `pulumi:"regexMatchTuples"`
}
