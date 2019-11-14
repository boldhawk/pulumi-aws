// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an API Gateway REST API.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_rest_api.html.markdown.
type RestApi struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The source of the API key for requests. Valid values are HEADER (default) and AUTHORIZER.
	ApiKeySource pulumi.StringOutput `pulumi:"apiKeySource"`

	// The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.
	BinaryMediaTypes pulumi.ArrayOutput `pulumi:"binaryMediaTypes"`

	// An OpenAPI specification that defines the set of routes and integrations to create as part of the REST API.
	Body pulumi.StringOutput `pulumi:"body"`

	// The creation date of the REST API
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`

	// The description of the REST API
	Description pulumi.StringOutput `pulumi:"description"`

	// Nested argument defining API endpoint configuration including endpoint type. Defined below.
	EndpointConfiguration pulumi.AnyOutput `pulumi:"endpointConfiguration"`

	// The execution ARN part to be used in [`lambdaPermission`](https://www.terraform.io/docs/providers/aws/r/lambda_permission.html)'s `sourceArn`
	// when allowing API Gateway to invoke a Lambda function,
	// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j`, which can be concatenated with allowed stage, method and resource path.
	ExecutionArn pulumi.StringOutput `pulumi:"executionArn"`

	// Minimum response size to compress for the REST API. Integer between -1 and 10485760 (10MB). Setting a value greater than -1 will enable compression, -1 disables compression (default).
	MinimumCompressionSize pulumi.IntOutput `pulumi:"minimumCompressionSize"`

	// The name of the REST API
	Name pulumi.StringOutput `pulumi:"name"`

	// JSON formatted policy document that controls access to the API Gateway.
	Policy pulumi.StringOutput `pulumi:"policy"`

	// The resource ID of the REST API's root
	RootResourceId pulumi.StringOutput `pulumi:"rootResourceId"`
}

// NewRestApi registers a new resource with the given unique name, arguments, and options.
func NewRestApi(ctx *pulumi.Context,
	name string, args *RestApiArgs, opts ...pulumi.ResourceOpt) (*RestApi, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["apiKeySource"] = args.ApiKeySource
		inputs["binaryMediaTypes"] = args.BinaryMediaTypes
		inputs["body"] = args.Body
		inputs["description"] = args.Description
		inputs["endpointConfiguration"] = args.EndpointConfiguration
		inputs["minimumCompressionSize"] = args.MinimumCompressionSize
		inputs["name"] = args.Name
		inputs["policy"] = args.Policy
	}
	var resource RestApi
	err := ctx.RegisterResource("aws:apigateway/restApi:RestApi", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRestApi gets an existing RestApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRestApi(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RestApiState, opts ...pulumi.ResourceOpt) (*RestApi, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["apiKeySource"] = state.ApiKeySource
		inputs["binaryMediaTypes"] = state.BinaryMediaTypes
		inputs["body"] = state.Body
		inputs["createdDate"] = state.CreatedDate
		inputs["description"] = state.Description
		inputs["endpointConfiguration"] = state.EndpointConfiguration
		inputs["executionArn"] = state.ExecutionArn
		inputs["minimumCompressionSize"] = state.MinimumCompressionSize
		inputs["name"] = state.Name
		inputs["policy"] = state.Policy
		inputs["rootResourceId"] = state.RootResourceId
	}
	var resource RestApi
	err := ctx.ReadResource("aws:apigateway/restApi:RestApi", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *RestApi) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *RestApi) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering RestApi resources.
type RestApiState struct {
	// The source of the API key for requests. Valid values are HEADER (default) and AUTHORIZER.
	ApiKeySource pulumi.StringInput `pulumi:"apiKeySource"`
	// The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.
	BinaryMediaTypes pulumi.ArrayInput `pulumi:"binaryMediaTypes"`
	// An OpenAPI specification that defines the set of routes and integrations to create as part of the REST API.
	Body pulumi.StringInput `pulumi:"body"`
	// The creation date of the REST API
	CreatedDate pulumi.StringInput `pulumi:"createdDate"`
	// The description of the REST API
	Description pulumi.StringInput `pulumi:"description"`
	// Nested argument defining API endpoint configuration including endpoint type. Defined below.
	EndpointConfiguration pulumi.AnyInput `pulumi:"endpointConfiguration"`
	// The execution ARN part to be used in [`lambdaPermission`](https://www.terraform.io/docs/providers/aws/r/lambda_permission.html)'s `sourceArn`
	// when allowing API Gateway to invoke a Lambda function,
	// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j`, which can be concatenated with allowed stage, method and resource path.
	ExecutionArn pulumi.StringInput `pulumi:"executionArn"`
	// Minimum response size to compress for the REST API. Integer between -1 and 10485760 (10MB). Setting a value greater than -1 will enable compression, -1 disables compression (default).
	MinimumCompressionSize pulumi.IntInput `pulumi:"minimumCompressionSize"`
	// The name of the REST API
	Name pulumi.StringInput `pulumi:"name"`
	// JSON formatted policy document that controls access to the API Gateway.
	Policy pulumi.StringInput `pulumi:"policy"`
	// The resource ID of the REST API's root
	RootResourceId pulumi.StringInput `pulumi:"rootResourceId"`
}

// The set of arguments for constructing a RestApi resource.
type RestApiArgs struct {
	// The source of the API key for requests. Valid values are HEADER (default) and AUTHORIZER.
	ApiKeySource pulumi.StringInput `pulumi:"apiKeySource"`
	// The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.
	BinaryMediaTypes pulumi.ArrayInput `pulumi:"binaryMediaTypes"`
	// An OpenAPI specification that defines the set of routes and integrations to create as part of the REST API.
	Body pulumi.StringInput `pulumi:"body"`
	// The description of the REST API
	Description pulumi.StringInput `pulumi:"description"`
	// Nested argument defining API endpoint configuration including endpoint type. Defined below.
	EndpointConfiguration pulumi.AnyInput `pulumi:"endpointConfiguration"`
	// Minimum response size to compress for the REST API. Integer between -1 and 10485760 (10MB). Setting a value greater than -1 will enable compression, -1 disables compression (default).
	MinimumCompressionSize pulumi.IntInput `pulumi:"minimumCompressionSize"`
	// The name of the REST API
	Name pulumi.StringInput `pulumi:"name"`
	// JSON formatted policy document that controls access to the API Gateway.
	Policy pulumi.StringInput `pulumi:"policy"`
}
