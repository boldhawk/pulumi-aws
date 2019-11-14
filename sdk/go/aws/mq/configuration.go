// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mq

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an MQ Configuration Resource. 
// 
// For more information on Amazon MQ, see [Amazon MQ documentation](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/welcome.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/mq_configuration.html.markdown.
type Configuration struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN of the configuration.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The broker configuration in XML format.
	// See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html)
	// for supported parameters and format of the XML.
	Data pulumi.StringOutput `pulumi:"data"`

	// The description of the configuration.
	Description pulumi.StringOutput `pulumi:"description"`

	// The type of broker engine.
	EngineType pulumi.StringOutput `pulumi:"engineType"`

	// The version of the broker engine.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`

	// The latest revision of the configuration.
	LatestRevision pulumi.IntOutput `pulumi:"latestRevision"`

	// The name of the configuration
	Name pulumi.StringOutput `pulumi:"name"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewConfiguration registers a new resource with the given unique name, arguments, and options.
func NewConfiguration(ctx *pulumi.Context,
	name string, args *ConfigurationArgs, opts ...pulumi.ResourceOpt) (*Configuration, error) {
	if args == nil || args.Data == nil {
		return nil, errors.New("missing required argument 'Data'")
	}
	if args == nil || args.EngineType == nil {
		return nil, errors.New("missing required argument 'EngineType'")
	}
	if args == nil || args.EngineVersion == nil {
		return nil, errors.New("missing required argument 'EngineVersion'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["data"] = args.Data
		inputs["description"] = args.Description
		inputs["engineType"] = args.EngineType
		inputs["engineVersion"] = args.EngineVersion
		inputs["name"] = args.Name
		inputs["tags"] = args.Tags
	}
	var resource Configuration
	err := ctx.RegisterResource("aws:mq/configuration:Configuration", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfiguration gets an existing Configuration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfiguration(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ConfigurationState, opts ...pulumi.ResourceOpt) (*Configuration, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["data"] = state.Data
		inputs["description"] = state.Description
		inputs["engineType"] = state.EngineType
		inputs["engineVersion"] = state.EngineVersion
		inputs["latestRevision"] = state.LatestRevision
		inputs["name"] = state.Name
		inputs["tags"] = state.Tags
	}
	var resource Configuration
	err := ctx.ReadResource("aws:mq/configuration:Configuration", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *Configuration) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *Configuration) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering Configuration resources.
type ConfigurationState struct {
	// The ARN of the configuration.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The broker configuration in XML format.
	// See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html)
	// for supported parameters and format of the XML.
	Data pulumi.StringInput `pulumi:"data"`
	// The description of the configuration.
	Description pulumi.StringInput `pulumi:"description"`
	// The type of broker engine.
	EngineType pulumi.StringInput `pulumi:"engineType"`
	// The version of the broker engine.
	EngineVersion pulumi.StringInput `pulumi:"engineVersion"`
	// The latest revision of the configuration.
	LatestRevision pulumi.IntInput `pulumi:"latestRevision"`
	// The name of the configuration
	Name pulumi.StringInput `pulumi:"name"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a Configuration resource.
type ConfigurationArgs struct {
	// The broker configuration in XML format.
	// See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html)
	// for supported parameters and format of the XML.
	Data pulumi.StringInput `pulumi:"data"`
	// The description of the configuration.
	Description pulumi.StringInput `pulumi:"description"`
	// The type of broker engine.
	EngineType pulumi.StringInput `pulumi:"engineType"`
	// The version of the broker engine.
	EngineVersion pulumi.StringInput `pulumi:"engineVersion"`
	// The name of the configuration
	Name pulumi.StringInput `pulumi:"name"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
