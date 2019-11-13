// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an API Gateway Authorizer.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_authorizer.html.markdown.
type Authorizer struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The credentials required for the authorizer.
	// To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.
	AuthorizerCredentials pulumi.StringOutput `pulumi:"authorizerCredentials"`

	// The TTL of cached authorizer results in seconds.
	// Defaults to `300`.
	AuthorizerResultTtlInSeconds pulumi.IntOutput `pulumi:"authorizerResultTtlInSeconds"`

	// The authorizer's Uniform Resource Identifier (URI).
	// This must be a well-formed Lambda function URI in the form of `arn:aws:apigateway:{region}:lambda:path/{service_api}`,
	// e.g. `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:012345678912:function:my-function/invocations`
	AuthorizerUri pulumi.StringOutput `pulumi:"authorizerUri"`

	// The source of the identity in an incoming request.
	// Defaults to `method.request.header.Authorization`. For `REQUEST` type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g. `"method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName"`
	IdentitySource pulumi.StringOutput `pulumi:"identitySource"`

	// A validation expression for the incoming identity.
	// For `TOKEN` type, this value should be a regular expression. The incoming token from the client is matched
	// against this expression, and will proceed if the token matches. If the token doesn't match,
	// the client receives a 401 Unauthorized response.
	IdentityValidationExpression pulumi.StringOutput `pulumi:"identityValidationExpression"`

	// The name of the authorizer
	Name pulumi.StringOutput `pulumi:"name"`

	// A list of the Amazon Cognito user pool ARNs.
	// Each element is of this format: `arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}`.
	ProviderArns pulumi.ArrayOutput `pulumi:"providerArns"`

	// The ID of the associated REST API
	RestApi pulumi.StringOutput `pulumi:"restApi"`

	// The type of the authorizer. Possible values are `TOKEN` for a Lambda function using a single authorization token submitted in a custom header, `REQUEST` for a Lambda function using incoming request parameters, or `COGNITO_USER_POOLS` for using an Amazon Cognito user pool.
	// Defaults to `TOKEN`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAuthorizer registers a new resource with the given unique name, arguments, and options.
func NewAuthorizer(ctx *pulumi.Context,
	name string, args *AuthorizerArgs, opts ...pulumi.ResourceOpt) (*Authorizer, error) {
	if args == nil || args.RestApi == nil {
		return nil, errors.New("missing required argument 'RestApi'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["authorizerCredentials"] = args.AuthorizerCredentials
		inputs["authorizerResultTtlInSeconds"] = args.AuthorizerResultTtlInSeconds
		inputs["authorizerUri"] = args.AuthorizerUri
		inputs["identitySource"] = args.IdentitySource
		inputs["identityValidationExpression"] = args.IdentityValidationExpression
		inputs["name"] = args.Name
		inputs["providerArns"] = args.ProviderArns
		inputs["restApi"] = args.RestApi
		inputs["type"] = args.Type
	}
	var resource Authorizer
	err := ctx.RegisterResource("aws:apigateway/authorizer:Authorizer", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthorizer gets an existing Authorizer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthorizer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AuthorizerState, opts ...pulumi.ResourceOpt) (*Authorizer, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["authorizerCredentials"] = state.AuthorizerCredentials
		inputs["authorizerResultTtlInSeconds"] = state.AuthorizerResultTtlInSeconds
		inputs["authorizerUri"] = state.AuthorizerUri
		inputs["identitySource"] = state.IdentitySource
		inputs["identityValidationExpression"] = state.IdentityValidationExpression
		inputs["name"] = state.Name
		inputs["providerArns"] = state.ProviderArns
		inputs["restApi"] = state.RestApi
		inputs["type"] = state.Type
	}
	var resource Authorizer
	err := ctx.ReadResource("aws:apigateway/authorizer:Authorizer", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *Authorizer) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *Authorizer) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering Authorizer resources.
type AuthorizerState struct {
	// The credentials required for the authorizer.
	// To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.
	AuthorizerCredentials pulumi.StringInput `pulumi:"authorizerCredentials"`
	// The TTL of cached authorizer results in seconds.
	// Defaults to `300`.
	AuthorizerResultTtlInSeconds pulumi.IntInput `pulumi:"authorizerResultTtlInSeconds"`
	// The authorizer's Uniform Resource Identifier (URI).
	// This must be a well-formed Lambda function URI in the form of `arn:aws:apigateway:{region}:lambda:path/{service_api}`,
	// e.g. `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:012345678912:function:my-function/invocations`
	AuthorizerUri pulumi.StringInput `pulumi:"authorizerUri"`
	// The source of the identity in an incoming request.
	// Defaults to `method.request.header.Authorization`. For `REQUEST` type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g. `"method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName"`
	IdentitySource pulumi.StringInput `pulumi:"identitySource"`
	// A validation expression for the incoming identity.
	// For `TOKEN` type, this value should be a regular expression. The incoming token from the client is matched
	// against this expression, and will proceed if the token matches. If the token doesn't match,
	// the client receives a 401 Unauthorized response.
	IdentityValidationExpression pulumi.StringInput `pulumi:"identityValidationExpression"`
	// The name of the authorizer
	Name pulumi.StringInput `pulumi:"name"`
	// A list of the Amazon Cognito user pool ARNs.
	// Each element is of this format: `arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}`.
	ProviderArns pulumi.ArrayInput `pulumi:"providerArns"`
	// The ID of the associated REST API
	RestApi pulumi.StringInput `pulumi:"restApi"`
	// The type of the authorizer. Possible values are `TOKEN` for a Lambda function using a single authorization token submitted in a custom header, `REQUEST` for a Lambda function using incoming request parameters, or `COGNITO_USER_POOLS` for using an Amazon Cognito user pool.
	// Defaults to `TOKEN`.
	Type pulumi.StringInput `pulumi:"type"`
}

// The set of arguments for constructing a Authorizer resource.
type AuthorizerArgs struct {
	// The credentials required for the authorizer.
	// To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.
	AuthorizerCredentials pulumi.StringInput `pulumi:"authorizerCredentials"`
	// The TTL of cached authorizer results in seconds.
	// Defaults to `300`.
	AuthorizerResultTtlInSeconds pulumi.IntInput `pulumi:"authorizerResultTtlInSeconds"`
	// The authorizer's Uniform Resource Identifier (URI).
	// This must be a well-formed Lambda function URI in the form of `arn:aws:apigateway:{region}:lambda:path/{service_api}`,
	// e.g. `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:012345678912:function:my-function/invocations`
	AuthorizerUri pulumi.StringInput `pulumi:"authorizerUri"`
	// The source of the identity in an incoming request.
	// Defaults to `method.request.header.Authorization`. For `REQUEST` type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g. `"method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName"`
	IdentitySource pulumi.StringInput `pulumi:"identitySource"`
	// A validation expression for the incoming identity.
	// For `TOKEN` type, this value should be a regular expression. The incoming token from the client is matched
	// against this expression, and will proceed if the token matches. If the token doesn't match,
	// the client receives a 401 Unauthorized response.
	IdentityValidationExpression pulumi.StringInput `pulumi:"identityValidationExpression"`
	// The name of the authorizer
	Name pulumi.StringInput `pulumi:"name"`
	// A list of the Amazon Cognito user pool ARNs.
	// Each element is of this format: `arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}`.
	ProviderArns pulumi.ArrayInput `pulumi:"providerArns"`
	// The ID of the associated REST API
	RestApi pulumi.StringInput `pulumi:"restApi"`
	// The type of the authorizer. Possible values are `TOKEN` for a Lambda function using a single authorization token submitted in a custom header, `REQUEST` for a Lambda function using incoming request parameters, or `COGNITO_USER_POOLS` for using an Amazon Cognito user pool.
	// Defaults to `TOKEN`.
	Type pulumi.StringInput `pulumi:"type"`
}
