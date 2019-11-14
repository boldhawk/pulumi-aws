// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kinesis

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Kinesis Analytics Application resource. Kinesis Analytics is a managed service that
// allows processing and analyzing streaming data using standard SQL.
// 
// For more details, see the [Amazon Kinesis Analytics Documentation][1].
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/kinesis_analytics_application.html.markdown.
type AnalyticsApplication struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN of the Kinesis Analytics Appliation.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The CloudWatch log stream options to monitor application errors.
	// See CloudWatch Logging Options below for more details.
	CloudwatchLoggingOptions pulumi.AnyOutput `pulumi:"cloudwatchLoggingOptions"`

	// SQL Code to transform input data, and generate output.
	Code pulumi.StringOutput `pulumi:"code"`

	// The Timestamp when the application version was created.
	CreateTimestamp pulumi.StringOutput `pulumi:"createTimestamp"`

	// Description of the application.
	Description pulumi.StringOutput `pulumi:"description"`

	// Input configuration of the application. See Inputs below for more details.
	Inputs pulumi.AnyOutput `pulumi:"inputs"`

	// The Timestamp when the application was last updated.
	LastUpdateTimestamp pulumi.StringOutput `pulumi:"lastUpdateTimestamp"`

	// Name of the Kinesis Analytics Application.
	Name pulumi.StringOutput `pulumi:"name"`

	// Output destination configuration of the application. See Outputs below for more details.
	Outputs pulumi.ArrayOutput `pulumi:"outputs"`

	// An S3 Reference Data Source for the application.
	// See Reference Data Sources below for more details.
	ReferenceDataSources pulumi.AnyOutput `pulumi:"referenceDataSources"`

	// The Status of the application.
	Status pulumi.StringOutput `pulumi:"status"`

	// Key-value mapping of tags for the Kinesis Analytics Application.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// The Version of the application.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewAnalyticsApplication registers a new resource with the given unique name, arguments, and options.
func NewAnalyticsApplication(ctx *pulumi.Context,
	name string, args *AnalyticsApplicationArgs, opts ...pulumi.ResourceOpt) (*AnalyticsApplication, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["cloudwatchLoggingOptions"] = args.CloudwatchLoggingOptions
		inputs["code"] = args.Code
		inputs["description"] = args.Description
		inputs["inputs"] = args.Inputs
		inputs["name"] = args.Name
		inputs["outputs"] = args.Outputs
		inputs["referenceDataSources"] = args.ReferenceDataSources
		inputs["tags"] = args.Tags
	}
	var resource AnalyticsApplication
	err := ctx.RegisterResource("aws:kinesis/analyticsApplication:AnalyticsApplication", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalyticsApplication gets an existing AnalyticsApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyticsApplication(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AnalyticsApplicationState, opts ...pulumi.ResourceOpt) (*AnalyticsApplication, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["cloudwatchLoggingOptions"] = state.CloudwatchLoggingOptions
		inputs["code"] = state.Code
		inputs["createTimestamp"] = state.CreateTimestamp
		inputs["description"] = state.Description
		inputs["inputs"] = state.Inputs
		inputs["lastUpdateTimestamp"] = state.LastUpdateTimestamp
		inputs["name"] = state.Name
		inputs["outputs"] = state.Outputs
		inputs["referenceDataSources"] = state.ReferenceDataSources
		inputs["status"] = state.Status
		inputs["tags"] = state.Tags
		inputs["version"] = state.Version
	}
	var resource AnalyticsApplication
	err := ctx.ReadResource("aws:kinesis/analyticsApplication:AnalyticsApplication", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *AnalyticsApplication) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *AnalyticsApplication) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering AnalyticsApplication resources.
type AnalyticsApplicationState struct {
	// The ARN of the Kinesis Analytics Appliation.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The CloudWatch log stream options to monitor application errors.
	// See CloudWatch Logging Options below for more details.
	CloudwatchLoggingOptions pulumi.AnyInput `pulumi:"cloudwatchLoggingOptions"`
	// SQL Code to transform input data, and generate output.
	Code pulumi.StringInput `pulumi:"code"`
	// The Timestamp when the application version was created.
	CreateTimestamp pulumi.StringInput `pulumi:"createTimestamp"`
	// Description of the application.
	Description pulumi.StringInput `pulumi:"description"`
	// Input configuration of the application. See Inputs below for more details.
	Inputs pulumi.AnyInput `pulumi:"inputs"`
	// The Timestamp when the application was last updated.
	LastUpdateTimestamp pulumi.StringInput `pulumi:"lastUpdateTimestamp"`
	// Name of the Kinesis Analytics Application.
	Name pulumi.StringInput `pulumi:"name"`
	// Output destination configuration of the application. See Outputs below for more details.
	Outputs pulumi.ArrayInput `pulumi:"outputs"`
	// An S3 Reference Data Source for the application.
	// See Reference Data Sources below for more details.
	ReferenceDataSources pulumi.AnyInput `pulumi:"referenceDataSources"`
	// The Status of the application.
	Status pulumi.StringInput `pulumi:"status"`
	// Key-value mapping of tags for the Kinesis Analytics Application.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The Version of the application.
	Version pulumi.IntInput `pulumi:"version"`
}

// The set of arguments for constructing a AnalyticsApplication resource.
type AnalyticsApplicationArgs struct {
	// The CloudWatch log stream options to monitor application errors.
	// See CloudWatch Logging Options below for more details.
	CloudwatchLoggingOptions pulumi.AnyInput `pulumi:"cloudwatchLoggingOptions"`
	// SQL Code to transform input data, and generate output.
	Code pulumi.StringInput `pulumi:"code"`
	// Description of the application.
	Description pulumi.StringInput `pulumi:"description"`
	// Input configuration of the application. See Inputs below for more details.
	Inputs pulumi.AnyInput `pulumi:"inputs"`
	// Name of the Kinesis Analytics Application.
	Name pulumi.StringInput `pulumi:"name"`
	// Output destination configuration of the application. See Outputs below for more details.
	Outputs pulumi.ArrayInput `pulumi:"outputs"`
	// An S3 Reference Data Source for the application.
	// See Reference Data Sources below for more details.
	ReferenceDataSources pulumi.AnyInput `pulumi:"referenceDataSources"`
	// Key-value mapping of tags for the Kinesis Analytics Application.
	Tags pulumi.MapInput `pulumi:"tags"`
}
