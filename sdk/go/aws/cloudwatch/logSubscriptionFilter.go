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
	s *pulumi.ResourceState
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
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["destinationArn"] = nil
		inputs["distribution"] = nil
		inputs["filterPattern"] = nil
		inputs["logGroup"] = nil
		inputs["name"] = nil
		inputs["roleArn"] = nil
	} else {
		inputs["destinationArn"] = args.DestinationArn
		inputs["distribution"] = args.Distribution
		inputs["filterPattern"] = args.FilterPattern
		inputs["logGroup"] = args.LogGroup
		inputs["name"] = args.Name
		inputs["roleArn"] = args.RoleArn
	}
	s, err := ctx.RegisterResource("aws:cloudwatch/logSubscriptionFilter:LogSubscriptionFilter", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LogSubscriptionFilter{s: s}, nil
}

// GetLogSubscriptionFilter gets an existing LogSubscriptionFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSubscriptionFilter(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LogSubscriptionFilterState, opts ...pulumi.ResourceOpt) (*LogSubscriptionFilter, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["destinationArn"] = state.DestinationArn
		inputs["distribution"] = state.Distribution
		inputs["filterPattern"] = state.FilterPattern
		inputs["logGroup"] = state.LogGroup
		inputs["name"] = state.Name
		inputs["roleArn"] = state.RoleArn
	}
	s, err := ctx.ReadResource("aws:cloudwatch/logSubscriptionFilter:LogSubscriptionFilter", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LogSubscriptionFilter{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *LogSubscriptionFilter) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *LogSubscriptionFilter) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
func (r *LogSubscriptionFilter) DestinationArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["destinationArn"])
}

// The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
func (r *LogSubscriptionFilter) Distribution() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["distribution"])
}

// A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
func (r *LogSubscriptionFilter) FilterPattern() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["filterPattern"])
}

// The name of the log group to associate the subscription filter with
func (r *LogSubscriptionFilter) LogGroup() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["logGroup"])
}

// A name for the subscription filter
func (r *LogSubscriptionFilter) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function. 
func (r *LogSubscriptionFilter) RoleArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["roleArn"])
}

// Input properties used for looking up and filtering LogSubscriptionFilter resources.
type LogSubscriptionFilterState struct {
	// The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
	DestinationArn interface{}
	// The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
	Distribution interface{}
	// A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
	FilterPattern interface{}
	// The name of the log group to associate the subscription filter with
	LogGroup interface{}
	// A name for the subscription filter
	Name interface{}
	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function. 
	RoleArn interface{}
}

// The set of arguments for constructing a LogSubscriptionFilter resource.
type LogSubscriptionFilterArgs struct {
	// The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
	DestinationArn interface{}
	// The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
	Distribution interface{}
	// A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
	FilterPattern interface{}
	// The name of the log group to associate the subscription filter with
	LogGroup interface{}
	// A name for the subscription filter
	Name interface{}
	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function. 
	RoleArn interface{}
}
