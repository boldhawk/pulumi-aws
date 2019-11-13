// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a WAF IPSet Resource
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/waf_ipset.html.markdown.
type IpSet struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN of the WAF IPSet.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR format) from which web requests originate.
	IpSetDescriptors pulumi.ArrayOutput `pulumi:"ipSetDescriptors"`

	// The name or description of the IPSet.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewIpSet registers a new resource with the given unique name, arguments, and options.
func NewIpSet(ctx *pulumi.Context,
	name string, args *IpSetArgs, opts ...pulumi.ResourceOpt) (*IpSet, error) {
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["ipSetDescriptors"] = args.IpSetDescriptors
		inputs["name"] = args.Name
	}
	var resource IpSet
	err := ctx.RegisterResource("aws:waf/ipSet:IpSet", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpSet gets an existing IpSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpSet(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IpSetState, opts ...pulumi.ResourceOpt) (*IpSet, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["ipSetDescriptors"] = state.IpSetDescriptors
		inputs["name"] = state.Name
	}
	var resource IpSet
	err := ctx.ReadResource("aws:waf/ipSet:IpSet", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *IpSet) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *IpSet) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering IpSet resources.
type IpSetState struct {
	// The ARN of the WAF IPSet.
	Arn pulumi.StringInput `pulumi:"arn"`
	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR format) from which web requests originate.
	IpSetDescriptors pulumi.ArrayInput `pulumi:"ipSetDescriptors"`
	// The name or description of the IPSet.
	Name pulumi.StringInput `pulumi:"name"`
}

// The set of arguments for constructing a IpSet resource.
type IpSetArgs struct {
	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR format) from which web requests originate.
	IpSetDescriptors pulumi.ArrayInput `pulumi:"ipSetDescriptors"`
	// The name or description of the IPSet.
	Name pulumi.StringInput `pulumi:"name"`
}
