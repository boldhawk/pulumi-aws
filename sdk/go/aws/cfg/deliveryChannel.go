// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS Config Delivery Channel.
// 
// > **Note:** Delivery Channel requires a [Configuration Recorder](https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html) to be present. Use of `dependsOn` (as shown below) is recommended to avoid race conditions.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/config_delivery_channel.html.markdown.
type DeliveryChannel struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
	Name pulumi.StringOutput `pulumi:"name"`

	// The name of the S3 bucket used to store the configuration history.
	S3BucketName pulumi.StringOutput `pulumi:"s3BucketName"`

	// The prefix for the specified S3 bucket.
	S3KeyPrefix pulumi.StringOutput `pulumi:"s3KeyPrefix"`

	// Options for how AWS Config delivers configuration snapshots. See below
	SnapshotDeliveryProperties pulumi.AnyOutput `pulumi:"snapshotDeliveryProperties"`

	// The ARN of the SNS topic that AWS Config delivers notifications to.
	SnsTopicArn pulumi.StringOutput `pulumi:"snsTopicArn"`
}

// NewDeliveryChannel registers a new resource with the given unique name, arguments, and options.
func NewDeliveryChannel(ctx *pulumi.Context,
	name string, args *DeliveryChannelArgs, opts ...pulumi.ResourceOpt) (*DeliveryChannel, error) {
	if args == nil || args.S3BucketName == nil {
		return nil, errors.New("missing required argument 'S3BucketName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["name"] = args.Name
		inputs["s3BucketName"] = args.S3BucketName
		inputs["s3KeyPrefix"] = args.S3KeyPrefix
		inputs["snapshotDeliveryProperties"] = args.SnapshotDeliveryProperties
		inputs["snsTopicArn"] = args.SnsTopicArn
	}
	var resource DeliveryChannel
	err := ctx.RegisterResource("aws:cfg/deliveryChannel:DeliveryChannel", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeliveryChannel gets an existing DeliveryChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeliveryChannel(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DeliveryChannelState, opts ...pulumi.ResourceOpt) (*DeliveryChannel, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["name"] = state.Name
		inputs["s3BucketName"] = state.S3BucketName
		inputs["s3KeyPrefix"] = state.S3KeyPrefix
		inputs["snapshotDeliveryProperties"] = state.SnapshotDeliveryProperties
		inputs["snsTopicArn"] = state.SnsTopicArn
	}
	var resource DeliveryChannel
	err := ctx.ReadResource("aws:cfg/deliveryChannel:DeliveryChannel", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *DeliveryChannel) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *DeliveryChannel) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering DeliveryChannel resources.
type DeliveryChannelState struct {
	// The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the S3 bucket used to store the configuration history.
	S3BucketName pulumi.StringInput `pulumi:"s3BucketName"`
	// The prefix for the specified S3 bucket.
	S3KeyPrefix pulumi.StringInput `pulumi:"s3KeyPrefix"`
	// Options for how AWS Config delivers configuration snapshots. See below
	SnapshotDeliveryProperties pulumi.AnyInput `pulumi:"snapshotDeliveryProperties"`
	// The ARN of the SNS topic that AWS Config delivers notifications to.
	SnsTopicArn pulumi.StringInput `pulumi:"snsTopicArn"`
}

// The set of arguments for constructing a DeliveryChannel resource.
type DeliveryChannelArgs struct {
	// The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the S3 bucket used to store the configuration history.
	S3BucketName pulumi.StringInput `pulumi:"s3BucketName"`
	// The prefix for the specified S3 bucket.
	S3KeyPrefix pulumi.StringInput `pulumi:"s3KeyPrefix"`
	// Options for how AWS Config delivers configuration snapshots. See below
	SnapshotDeliveryProperties pulumi.AnyInput `pulumi:"snapshotDeliveryProperties"`
	// The ARN of the SNS topic that AWS Config delivers notifications to.
	SnsTopicArn pulumi.StringInput `pulumi:"snsTopicArn"`
}
