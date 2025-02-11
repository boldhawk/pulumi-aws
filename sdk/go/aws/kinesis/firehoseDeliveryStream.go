// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kinesis

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Kinesis Firehose Delivery Stream resource. Amazon Kinesis Firehose is a fully managed, elastic service to easily deliver real-time data streams to destinations such as Amazon S3 and Amazon Redshift.
// 
// For more details, see the [Amazon Kinesis Firehose Documentation][1].
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/kinesis_firehose_delivery_stream.html.markdown.
type FirehoseDeliveryStream struct {
	s *pulumi.ResourceState
}

// NewFirehoseDeliveryStream registers a new resource with the given unique name, arguments, and options.
func NewFirehoseDeliveryStream(ctx *pulumi.Context,
	name string, args *FirehoseDeliveryStreamArgs, opts ...pulumi.ResourceOpt) (*FirehoseDeliveryStream, error) {
	if args == nil || args.Destination == nil {
		return nil, errors.New("missing required argument 'Destination'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["arn"] = nil
		inputs["destination"] = nil
		inputs["destinationId"] = nil
		inputs["elasticsearchConfiguration"] = nil
		inputs["extendedS3Configuration"] = nil
		inputs["kinesisSourceConfiguration"] = nil
		inputs["name"] = nil
		inputs["redshiftConfiguration"] = nil
		inputs["s3Configuration"] = nil
		inputs["serverSideEncryption"] = nil
		inputs["splunkConfiguration"] = nil
		inputs["tags"] = nil
		inputs["versionId"] = nil
	} else {
		inputs["arn"] = args.Arn
		inputs["destination"] = args.Destination
		inputs["destinationId"] = args.DestinationId
		inputs["elasticsearchConfiguration"] = args.ElasticsearchConfiguration
		inputs["extendedS3Configuration"] = args.ExtendedS3Configuration
		inputs["kinesisSourceConfiguration"] = args.KinesisSourceConfiguration
		inputs["name"] = args.Name
		inputs["redshiftConfiguration"] = args.RedshiftConfiguration
		inputs["s3Configuration"] = args.S3Configuration
		inputs["serverSideEncryption"] = args.ServerSideEncryption
		inputs["splunkConfiguration"] = args.SplunkConfiguration
		inputs["tags"] = args.Tags
		inputs["versionId"] = args.VersionId
	}
	s, err := ctx.RegisterResource("aws:kinesis/firehoseDeliveryStream:FirehoseDeliveryStream", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FirehoseDeliveryStream{s: s}, nil
}

// GetFirehoseDeliveryStream gets an existing FirehoseDeliveryStream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirehoseDeliveryStream(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FirehoseDeliveryStreamState, opts ...pulumi.ResourceOpt) (*FirehoseDeliveryStream, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["destination"] = state.Destination
		inputs["destinationId"] = state.DestinationId
		inputs["elasticsearchConfiguration"] = state.ElasticsearchConfiguration
		inputs["extendedS3Configuration"] = state.ExtendedS3Configuration
		inputs["kinesisSourceConfiguration"] = state.KinesisSourceConfiguration
		inputs["name"] = state.Name
		inputs["redshiftConfiguration"] = state.RedshiftConfiguration
		inputs["s3Configuration"] = state.S3Configuration
		inputs["serverSideEncryption"] = state.ServerSideEncryption
		inputs["splunkConfiguration"] = state.SplunkConfiguration
		inputs["tags"] = state.Tags
		inputs["versionId"] = state.VersionId
	}
	s, err := ctx.ReadResource("aws:kinesis/firehoseDeliveryStream:FirehoseDeliveryStream", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FirehoseDeliveryStream{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *FirehoseDeliveryStream) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *FirehoseDeliveryStream) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The Amazon Resource Name (ARN) specifying the Stream
func (r *FirehoseDeliveryStream) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// This is the destination to where the data is delivered. The only options are `s3` (Deprecated, use `extendedS3` instead), `extendedS3`, `redshift`, `elasticsearch`, and `splunk`.
func (r *FirehoseDeliveryStream) Destination() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["destination"])
}

func (r *FirehoseDeliveryStream) DestinationId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["destinationId"])
}

func (r *FirehoseDeliveryStream) ElasticsearchConfiguration() pulumi.Output {
	return r.s.State["elasticsearchConfiguration"]
}

// Enhanced configuration options for the s3 destination. More details are given below.
func (r *FirehoseDeliveryStream) ExtendedS3Configuration() pulumi.Output {
	return r.s.State["extendedS3Configuration"]
}

// Allows the ability to specify the kinesis stream that is used as the source of the firehose delivery stream.
func (r *FirehoseDeliveryStream) KinesisSourceConfiguration() pulumi.Output {
	return r.s.State["kinesisSourceConfiguration"]
}

// A name to identify the stream. This is unique to the
// AWS account and region the Stream is created in.
func (r *FirehoseDeliveryStream) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Configuration options if redshift is the destination.
// Using `redshiftConfiguration` requires the user to also specify a
// `s3Configuration` block. More details are given below.
func (r *FirehoseDeliveryStream) RedshiftConfiguration() pulumi.Output {
	return r.s.State["redshiftConfiguration"]
}

// Required for non-S3 destinations. For S3 destination, use `extendedS3Configuration` instead. Configuration options for the s3 destination (or the intermediate bucket if the destination
// is redshift). More details are given below.
func (r *FirehoseDeliveryStream) S3Configuration() pulumi.Output {
	return r.s.State["s3Configuration"]
}

// Encrypt at rest options.
// Server-side encryption should not be enabled when a kinesis stream is configured as the source of the firehose delivery stream.
func (r *FirehoseDeliveryStream) ServerSideEncryption() pulumi.Output {
	return r.s.State["serverSideEncryption"]
}

func (r *FirehoseDeliveryStream) SplunkConfiguration() pulumi.Output {
	return r.s.State["splunkConfiguration"]
}

// A mapping of tags to assign to the resource.
func (r *FirehoseDeliveryStream) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Specifies the table version for the output data schema. Defaults to `LATEST`.
func (r *FirehoseDeliveryStream) VersionId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["versionId"])
}

// Input properties used for looking up and filtering FirehoseDeliveryStream resources.
type FirehoseDeliveryStreamState struct {
	// The Amazon Resource Name (ARN) specifying the Stream
	Arn interface{}
	// This is the destination to where the data is delivered. The only options are `s3` (Deprecated, use `extendedS3` instead), `extendedS3`, `redshift`, `elasticsearch`, and `splunk`.
	Destination interface{}
	DestinationId interface{}
	ElasticsearchConfiguration interface{}
	// Enhanced configuration options for the s3 destination. More details are given below.
	ExtendedS3Configuration interface{}
	// Allows the ability to specify the kinesis stream that is used as the source of the firehose delivery stream.
	KinesisSourceConfiguration interface{}
	// A name to identify the stream. This is unique to the
	// AWS account and region the Stream is created in.
	Name interface{}
	// Configuration options if redshift is the destination.
	// Using `redshiftConfiguration` requires the user to also specify a
	// `s3Configuration` block. More details are given below.
	RedshiftConfiguration interface{}
	// Required for non-S3 destinations. For S3 destination, use `extendedS3Configuration` instead. Configuration options for the s3 destination (or the intermediate bucket if the destination
	// is redshift). More details are given below.
	S3Configuration interface{}
	// Encrypt at rest options.
	// Server-side encryption should not be enabled when a kinesis stream is configured as the source of the firehose delivery stream.
	ServerSideEncryption interface{}
	SplunkConfiguration interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Specifies the table version for the output data schema. Defaults to `LATEST`.
	VersionId interface{}
}

// The set of arguments for constructing a FirehoseDeliveryStream resource.
type FirehoseDeliveryStreamArgs struct {
	// The Amazon Resource Name (ARN) specifying the Stream
	Arn interface{}
	// This is the destination to where the data is delivered. The only options are `s3` (Deprecated, use `extendedS3` instead), `extendedS3`, `redshift`, `elasticsearch`, and `splunk`.
	Destination interface{}
	DestinationId interface{}
	ElasticsearchConfiguration interface{}
	// Enhanced configuration options for the s3 destination. More details are given below.
	ExtendedS3Configuration interface{}
	// Allows the ability to specify the kinesis stream that is used as the source of the firehose delivery stream.
	KinesisSourceConfiguration interface{}
	// A name to identify the stream. This is unique to the
	// AWS account and region the Stream is created in.
	Name interface{}
	// Configuration options if redshift is the destination.
	// Using `redshiftConfiguration` requires the user to also specify a
	// `s3Configuration` block. More details are given below.
	RedshiftConfiguration interface{}
	// Required for non-S3 destinations. For S3 destination, use `extendedS3Configuration` instead. Configuration options for the s3 destination (or the intermediate bucket if the destination
	// is redshift). More details are given below.
	S3Configuration interface{}
	// Encrypt at rest options.
	// Server-side encryption should not be enabled when a kinesis stream is configured as the source of the firehose delivery stream.
	ServerSideEncryption interface{}
	SplunkConfiguration interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Specifies the table version for the output data schema. Defaults to `LATEST`.
	VersionId interface{}
}
