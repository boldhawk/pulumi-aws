// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a SageMaker model resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/sagemaker_model.html.markdown.
type Model struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The Amazon Resource Name (ARN) assigned by AWS to this model.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// Specifies containers in the inference pipeline. If not specified, the `primaryContainer` argument is required. Fields are documented below.
	Containers pulumi.ArrayOutput `pulumi:"containers"`

	// Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
	EnableNetworkIsolation pulumi.BoolOutput `pulumi:"enableNetworkIsolation"`

	// A role that SageMaker can assume to access model artifacts and docker images for deployment.
	ExecutionRoleArn pulumi.StringOutput `pulumi:"executionRoleArn"`

	// The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`

	// The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
	PrimaryContainer pulumi.AnyOutput `pulumi:"primaryContainer"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
	VpcConfig pulumi.AnyOutput `pulumi:"vpcConfig"`
}

// NewModel registers a new resource with the given unique name, arguments, and options.
func NewModel(ctx *pulumi.Context,
	name string, args *ModelArgs, opts ...pulumi.ResourceOpt) (*Model, error) {
	if args == nil || args.ExecutionRoleArn == nil {
		return nil, errors.New("missing required argument 'ExecutionRoleArn'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["containers"] = args.Containers
		inputs["enableNetworkIsolation"] = args.EnableNetworkIsolation
		inputs["executionRoleArn"] = args.ExecutionRoleArn
		inputs["name"] = args.Name
		inputs["primaryContainer"] = args.PrimaryContainer
		inputs["tags"] = args.Tags
		inputs["vpcConfig"] = args.VpcConfig
	}
	var resource Model
	err := ctx.RegisterResource("aws:sagemaker/model:Model", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModel gets an existing Model resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModel(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ModelState, opts ...pulumi.ResourceOpt) (*Model, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["containers"] = state.Containers
		inputs["enableNetworkIsolation"] = state.EnableNetworkIsolation
		inputs["executionRoleArn"] = state.ExecutionRoleArn
		inputs["name"] = state.Name
		inputs["primaryContainer"] = state.PrimaryContainer
		inputs["tags"] = state.Tags
		inputs["vpcConfig"] = state.VpcConfig
	}
	var resource Model
	err := ctx.ReadResource("aws:sagemaker/model:Model", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *Model) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *Model) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering Model resources.
type ModelState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this model.
	Arn pulumi.StringInput `pulumi:"arn"`
	// Specifies containers in the inference pipeline. If not specified, the `primaryContainer` argument is required. Fields are documented below.
	Containers pulumi.ArrayInput `pulumi:"containers"`
	// Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
	EnableNetworkIsolation pulumi.BoolInput `pulumi:"enableNetworkIsolation"`
	// A role that SageMaker can assume to access model artifacts and docker images for deployment.
	ExecutionRoleArn pulumi.StringInput `pulumi:"executionRoleArn"`
	// The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
	Name pulumi.StringInput `pulumi:"name"`
	// The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
	PrimaryContainer pulumi.AnyInput `pulumi:"primaryContainer"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
	VpcConfig pulumi.AnyInput `pulumi:"vpcConfig"`
}

// The set of arguments for constructing a Model resource.
type ModelArgs struct {
	// Specifies containers in the inference pipeline. If not specified, the `primaryContainer` argument is required. Fields are documented below.
	Containers pulumi.ArrayInput `pulumi:"containers"`
	// Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
	EnableNetworkIsolation pulumi.BoolInput `pulumi:"enableNetworkIsolation"`
	// A role that SageMaker can assume to access model artifacts and docker images for deployment.
	ExecutionRoleArn pulumi.StringInput `pulumi:"executionRoleArn"`
	// The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
	Name pulumi.StringInput `pulumi:"name"`
	// The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
	PrimaryContainer pulumi.AnyInput `pulumi:"primaryContainer"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
	VpcConfig pulumi.AnyInput `pulumi:"vpcConfig"`
}
