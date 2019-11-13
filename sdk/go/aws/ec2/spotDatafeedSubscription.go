// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > **Note:** There is only a single subscription allowed per account.
// 
// To help you understand the charges for your Spot instances, Amazon EC2 provides a data feed that describes your Spot instance usage and pricing.
// This data feed is sent to an Amazon S3 bucket that you specify when you subscribe to the data feed.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/spot_datafeed_subscription.html.markdown.
type SpotDatafeedSubscription struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket pulumi.StringOutput `pulumi:"bucket"`

	// Path of folder inside bucket to place spot pricing data.
	Prefix pulumi.StringOutput `pulumi:"prefix"`
}

// NewSpotDatafeedSubscription registers a new resource with the given unique name, arguments, and options.
func NewSpotDatafeedSubscription(ctx *pulumi.Context,
	name string, args *SpotDatafeedSubscriptionArgs, opts ...pulumi.ResourceOpt) (*SpotDatafeedSubscription, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["bucket"] = args.Bucket
		inputs["prefix"] = args.Prefix
	}
	var resource SpotDatafeedSubscription
	err := ctx.RegisterResource("aws:ec2/spotDatafeedSubscription:SpotDatafeedSubscription", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpotDatafeedSubscription gets an existing SpotDatafeedSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpotDatafeedSubscription(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SpotDatafeedSubscriptionState, opts ...pulumi.ResourceOpt) (*SpotDatafeedSubscription, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["bucket"] = state.Bucket
		inputs["prefix"] = state.Prefix
	}
	var resource SpotDatafeedSubscription
	err := ctx.ReadResource("aws:ec2/spotDatafeedSubscription:SpotDatafeedSubscription", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *SpotDatafeedSubscription) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *SpotDatafeedSubscription) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering SpotDatafeedSubscription resources.
type SpotDatafeedSubscriptionState struct {
	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket pulumi.StringInput `pulumi:"bucket"`
	// Path of folder inside bucket to place spot pricing data.
	Prefix pulumi.StringInput `pulumi:"prefix"`
}

// The set of arguments for constructing a SpotDatafeedSubscription resource.
type SpotDatafeedSubscriptionArgs struct {
	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket pulumi.StringInput `pulumi:"bucket"`
	// Path of folder inside bucket to place spot pricing data.
	Prefix pulumi.StringInput `pulumi:"prefix"`
}
