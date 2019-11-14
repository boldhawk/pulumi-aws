// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a Lambda function alias. Creates an alias that points to the specified Lambda function version.
// 
// For information about Lambda and how to use it, see [What is AWS Lambda?][1]
// For information about function aliases, see [CreateAlias][2] and [AliasRoutingConfiguration][3] in the API docs.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lambda_alias.html.markdown.
type Alias struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The Amazon Resource Name (ARN) identifying your Lambda function alias.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// Description of the alias.
	Description pulumi.StringOutput `pulumi:"description"`

	// The function ARN of the Lambda function for which you want to create an alias.
	FunctionName pulumi.StringOutput `pulumi:"functionName"`

	// Lambda function version for which you are creating the alias. Pattern: `(\$LATEST|[0-9]+)`.
	FunctionVersion pulumi.StringOutput `pulumi:"functionVersion"`

	// The ARN to be used for invoking Lambda Function from API Gateway - to be used in [`apigateway.Integration`](https://www.terraform.io/docs/providers/aws/r/api_gateway_integration.html)'s `uri`
	InvokeArn pulumi.StringOutput `pulumi:"invokeArn"`

	// Name for the alias you are creating. Pattern: `(?!^[0-9]+$)([a-zA-Z0-9-_]+)`
	Name pulumi.StringOutput `pulumi:"name"`

	// The Lambda alias' route configuration settings. Fields documented below
	RoutingConfig pulumi.AnyOutput `pulumi:"routingConfig"`
}

// NewAlias registers a new resource with the given unique name, arguments, and options.
func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOpt) (*Alias, error) {
	if args == nil || args.FunctionName == nil {
		return nil, errors.New("missing required argument 'FunctionName'")
	}
	if args == nil || args.FunctionVersion == nil {
		return nil, errors.New("missing required argument 'FunctionVersion'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["description"] = args.Description
		inputs["functionName"] = args.FunctionName
		inputs["functionVersion"] = args.FunctionVersion
		inputs["name"] = args.Name
		inputs["routingConfig"] = args.RoutingConfig
	}
	var resource Alias
	err := ctx.RegisterResource("aws:lambda/alias:Alias", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlias gets an existing Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AliasState, opts ...pulumi.ResourceOpt) (*Alias, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["description"] = state.Description
		inputs["functionName"] = state.FunctionName
		inputs["functionVersion"] = state.FunctionVersion
		inputs["invokeArn"] = state.InvokeArn
		inputs["name"] = state.Name
		inputs["routingConfig"] = state.RoutingConfig
	}
	var resource Alias
	err := ctx.ReadResource("aws:lambda/alias:Alias", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *Alias) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *Alias) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering Alias resources.
type AliasState struct {
	// The Amazon Resource Name (ARN) identifying your Lambda function alias.
	Arn pulumi.StringInput `pulumi:"arn"`
	// Description of the alias.
	Description pulumi.StringInput `pulumi:"description"`
	// The function ARN of the Lambda function for which you want to create an alias.
	FunctionName pulumi.StringInput `pulumi:"functionName"`
	// Lambda function version for which you are creating the alias. Pattern: `(\$LATEST|[0-9]+)`.
	FunctionVersion pulumi.StringInput `pulumi:"functionVersion"`
	// The ARN to be used for invoking Lambda Function from API Gateway - to be used in [`apigateway.Integration`](https://www.terraform.io/docs/providers/aws/r/api_gateway_integration.html)'s `uri`
	InvokeArn pulumi.StringInput `pulumi:"invokeArn"`
	// Name for the alias you are creating. Pattern: `(?!^[0-9]+$)([a-zA-Z0-9-_]+)`
	Name pulumi.StringInput `pulumi:"name"`
	// The Lambda alias' route configuration settings. Fields documented below
	RoutingConfig pulumi.AnyInput `pulumi:"routingConfig"`
}

// The set of arguments for constructing a Alias resource.
type AliasArgs struct {
	// Description of the alias.
	Description pulumi.StringInput `pulumi:"description"`
	// The function ARN of the Lambda function for which you want to create an alias.
	FunctionName pulumi.StringInput `pulumi:"functionName"`
	// Lambda function version for which you are creating the alias. Pattern: `(\$LATEST|[0-9]+)`.
	FunctionVersion pulumi.StringInput `pulumi:"functionVersion"`
	// Name for the alias you are creating. Pattern: `(?!^[0-9]+$)([a-zA-Z0-9-_]+)`
	Name pulumi.StringInput `pulumi:"name"`
	// The Lambda alias' route configuration settings. Fields documented below
	RoutingConfig pulumi.AnyInput `pulumi:"routingConfig"`
}
