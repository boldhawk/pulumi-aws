// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CloudWatch Log Stream resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_stream.html.markdown.
type LogStream struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The Amazon Resource Name (ARN) specifying the log stream.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The name of the log group under which the log stream is to be created.
	LogGroupName pulumi.StringOutput `pulumi:"logGroupName"`

	// The name of the log stream. Must not be longer than 512 characters and must not contain `:`
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewLogStream registers a new resource with the given unique name, arguments, and options.
func NewLogStream(ctx *pulumi.Context,
	name string, args *LogStreamArgs, opts ...pulumi.ResourceOpt) (*LogStream, error) {
	if args == nil || args.LogGroupName == nil {
		return nil, errors.New("missing required argument 'LogGroupName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["logGroupName"] = args.LogGroupName
		inputs["name"] = args.Name
	}
	var resource LogStream
	err := ctx.RegisterResource("aws:cloudwatch/logStream:LogStream", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogStream gets an existing LogStream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogStream(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LogStreamState, opts ...pulumi.ResourceOpt) (*LogStream, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["logGroupName"] = state.LogGroupName
		inputs["name"] = state.Name
	}
	var resource LogStream
	err := ctx.ReadResource("aws:cloudwatch/logStream:LogStream", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *LogStream) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *LogStream) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering LogStream resources.
type LogStreamState struct {
	// The Amazon Resource Name (ARN) specifying the log stream.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The name of the log group under which the log stream is to be created.
	LogGroupName pulumi.StringInput `pulumi:"logGroupName"`
	// The name of the log stream. Must not be longer than 512 characters and must not contain `:`
	Name pulumi.StringInput `pulumi:"name"`
}

// The set of arguments for constructing a LogStream resource.
type LogStreamArgs struct {
	// The name of the log group under which the log stream is to be created.
	LogGroupName pulumi.StringInput `pulumi:"logGroupName"`
	// The name of the log stream. Must not be longer than 512 characters and must not contain `:`
	Name pulumi.StringInput `pulumi:"name"`
}
