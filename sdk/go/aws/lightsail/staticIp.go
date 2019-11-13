// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allocates a static IP address.
// 
// > **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lightsail_static_ip.html.markdown.
type StaticIp struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN of the Lightsail static IP
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The allocated static IP address
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`

	// The name for the allocated static IP
	Name pulumi.StringOutput `pulumi:"name"`

	// The support code.
	SupportCode pulumi.StringOutput `pulumi:"supportCode"`
}

// NewStaticIp registers a new resource with the given unique name, arguments, and options.
func NewStaticIp(ctx *pulumi.Context,
	name string, args *StaticIpArgs, opts ...pulumi.ResourceOpt) (*StaticIp, error) {
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["name"] = args.Name
	}
	var resource StaticIp
	err := ctx.RegisterResource("aws:lightsail/staticIp:StaticIp", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStaticIp gets an existing StaticIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStaticIp(ctx *pulumi.Context,
	name string, id pulumi.ID, state *StaticIpState, opts ...pulumi.ResourceOpt) (*StaticIp, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["ipAddress"] = state.IpAddress
		inputs["name"] = state.Name
		inputs["supportCode"] = state.SupportCode
	}
	var resource StaticIp
	err := ctx.ReadResource("aws:lightsail/staticIp:StaticIp", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *StaticIp) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *StaticIp) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering StaticIp resources.
type StaticIpState struct {
	// The ARN of the Lightsail static IP
	Arn pulumi.StringInput `pulumi:"arn"`
	// The allocated static IP address
	IpAddress pulumi.StringInput `pulumi:"ipAddress"`
	// The name for the allocated static IP
	Name pulumi.StringInput `pulumi:"name"`
	// The support code.
	SupportCode pulumi.StringInput `pulumi:"supportCode"`
}

// The set of arguments for constructing a StaticIp resource.
type StaticIpArgs struct {
	// The name for the allocated static IP
	Name pulumi.StringInput `pulumi:"name"`
}
