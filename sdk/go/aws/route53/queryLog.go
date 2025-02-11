// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Route53 query logging configuration resource.
// 
// > **NOTE:** There are restrictions on the configuration of query logging. Notably,
// the CloudWatch log group must be in the `us-east-1` region,
// a permissive CloudWatch log resource policy must be in place, and
// the Route53 hosted zone must be public.
// See [Configuring Logging for DNS Queries](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/query-logs.html?console_help=true#query-logs-configuring) for additional details.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/route53_query_log.html.markdown.
type QueryLog struct {
	s *pulumi.ResourceState
}

// NewQueryLog registers a new resource with the given unique name, arguments, and options.
func NewQueryLog(ctx *pulumi.Context,
	name string, args *QueryLogArgs, opts ...pulumi.ResourceOpt) (*QueryLog, error) {
	if args == nil || args.CloudwatchLogGroupArn == nil {
		return nil, errors.New("missing required argument 'CloudwatchLogGroupArn'")
	}
	if args == nil || args.ZoneId == nil {
		return nil, errors.New("missing required argument 'ZoneId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cloudwatchLogGroupArn"] = nil
		inputs["zoneId"] = nil
	} else {
		inputs["cloudwatchLogGroupArn"] = args.CloudwatchLogGroupArn
		inputs["zoneId"] = args.ZoneId
	}
	s, err := ctx.RegisterResource("aws:route53/queryLog:QueryLog", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &QueryLog{s: s}, nil
}

// GetQueryLog gets an existing QueryLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueryLog(ctx *pulumi.Context,
	name string, id pulumi.ID, state *QueryLogState, opts ...pulumi.ResourceOpt) (*QueryLog, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["cloudwatchLogGroupArn"] = state.CloudwatchLogGroupArn
		inputs["zoneId"] = state.ZoneId
	}
	s, err := ctx.ReadResource("aws:route53/queryLog:QueryLog", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &QueryLog{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *QueryLog) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *QueryLog) ID() pulumi.IDOutput {
	return r.s.ID()
}

// CloudWatch log group ARN to send query logs.
func (r *QueryLog) CloudwatchLogGroupArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["cloudwatchLogGroupArn"])
}

// Route53 hosted zone ID to enable query logs.
func (r *QueryLog) ZoneId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["zoneId"])
}

// Input properties used for looking up and filtering QueryLog resources.
type QueryLogState struct {
	// CloudWatch log group ARN to send query logs.
	CloudwatchLogGroupArn interface{}
	// Route53 hosted zone ID to enable query logs.
	ZoneId interface{}
}

// The set of arguments for constructing a QueryLog resource.
type QueryLogArgs struct {
	// CloudWatch log group ARN to send query logs.
	CloudwatchLogGroupArn interface{}
	// Route53 hosted zone ID to enable query logs.
	ZoneId interface{}
}
