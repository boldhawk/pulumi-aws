// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CloudWatch Logs subscription filter resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_subscription_filter.html.markdown.
type LogSubscriptionFilter struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
	DestinationArn pulumi.StringOutput `pulumi:"destinationArn"`

	// The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
	Distribution pulumi.StringOutput `pulumi:"distribution"`

	// A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
	FilterPattern pulumi.StringOutput `pulumi:"filterPattern"`

	// The name of the log group to associate the subscription filter with
	LogGroup pulumi.StringOutput `pulumi:"logGroup"`

	// A name for the subscription filter
	Name pulumi.StringOutput `pulumi:"name"`

	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function. 
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
}

// NewLogSubscriptionFilter registers a new resource with the given unique name, arguments, and options.
func NewLogSubscriptionFilter(ctx *pulumi.Context,
	name string, args *LogSubscriptionFilterArgs, opts ...pulumi.ResourceOpt) (*LogSubscriptionFilter, error) {
	if args == nil || args.DestinationArn == nil {
		return nil, errors.New("missing required argument 'DestinationArn'")
	}
	if args == nil || args.FilterPattern == nil {
		return nil, errors.New("missing required argument 'FilterPattern'")
	}
	if args == nil || args.LogGroup == nil {
		return nil, errors.New("missing required argument 'LogGroup'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["destinationArn"] = args.DestinationArn
		inputs["distribution"] = args.Distribution
		inputs["filterPattern"] = args.FilterPattern
		inputs["logGroup"] = args.LogGroup
		inputs["name"] = args.Name
		inputs["roleArn"] = args.RoleArn
	}
	var resource LogSubscriptionFilter
	err := ctx.RegisterResource("aws:cloudwatch/logSubscriptionFilter:LogSubscriptionFilter", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSubscriptionFilter gets an existing LogSubscriptionFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSubscriptionFilter(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LogSubscriptionFilterState, opts ...pulumi.ResourceOpt) (*LogSubscriptionFilter, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["destinationArn"] = state.DestinationArn
		inputs["distribution"] = state.Distribution
		inputs["filterPattern"] = state.FilterPattern
		inputs["logGroup"] = state.LogGroup
		inputs["name"] = state.Name
		inputs["roleArn"] = state.RoleArn
	}
	var resource LogSubscriptionFilter
	err := ctx.ReadResource("aws:cloudwatch/logSubscriptionFilter:LogSubscriptionFilter", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *LogSubscriptionFilter) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *LogSubscriptionFilter) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering LogSubscriptionFilter resources.
type LogSubscriptionFilterState struct {
	// The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
	DestinationArn pulumi.StringInput `pulumi:"destinationArn"`
	// The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
	Distribution pulumi.StringInput `pulumi:"distribution"`
	// A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
	FilterPattern pulumi.StringInput `pulumi:"filterPattern"`
	// The name of the log group to associate the subscription filter with
	LogGroup pulumi.StringInput `pulumi:"logGroup"`
	// A name for the subscription filter
	Name pulumi.StringInput `pulumi:"name"`
	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function. 
	RoleArn pulumi.StringInput `pulumi:"roleArn"`
}

// The set of arguments for constructing a LogSubscriptionFilter resource.
type LogSubscriptionFilterArgs struct {
	// The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
	DestinationArn pulumi.StringInput `pulumi:"destinationArn"`
	// The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
	Distribution pulumi.StringInput `pulumi:"distribution"`
	// A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
	FilterPattern pulumi.StringInput `pulumi:"filterPattern"`
	// The name of the log group to associate the subscription filter with
	LogGroup pulumi.StringInput `pulumi:"logGroup"`
	// A name for the subscription filter
	Name pulumi.StringInput `pulumi:"name"`
	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function. 
	RoleArn pulumi.StringInput `pulumi:"roleArn"`
}
