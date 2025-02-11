// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an HTTP Method Integration Response for an API Gateway Resource.
// 
// > **Note:** Depends on having `apigateway.Integration` inside your rest api. To ensure this
// you might need to add an explicit `dependsOn` for clean runs.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_integration_response.html.markdown.
type IntegrationResponse struct {
	s *pulumi.ResourceState
}

// NewIntegrationResponse registers a new resource with the given unique name, arguments, and options.
func NewIntegrationResponse(ctx *pulumi.Context,
	name string, args *IntegrationResponseArgs, opts ...pulumi.ResourceOpt) (*IntegrationResponse, error) {
	if args == nil || args.HttpMethod == nil {
		return nil, errors.New("missing required argument 'HttpMethod'")
	}
	if args == nil || args.ResourceId == nil {
		return nil, errors.New("missing required argument 'ResourceId'")
	}
	if args == nil || args.RestApi == nil {
		return nil, errors.New("missing required argument 'RestApi'")
	}
	if args == nil || args.StatusCode == nil {
		return nil, errors.New("missing required argument 'StatusCode'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["contentHandling"] = nil
		inputs["httpMethod"] = nil
		inputs["resourceId"] = nil
		inputs["responseParameters"] = nil
		inputs["responseTemplates"] = nil
		inputs["restApi"] = nil
		inputs["selectionPattern"] = nil
		inputs["statusCode"] = nil
	} else {
		inputs["contentHandling"] = args.ContentHandling
		inputs["httpMethod"] = args.HttpMethod
		inputs["resourceId"] = args.ResourceId
		inputs["responseParameters"] = args.ResponseParameters
		inputs["responseTemplates"] = args.ResponseTemplates
		inputs["restApi"] = args.RestApi
		inputs["selectionPattern"] = args.SelectionPattern
		inputs["statusCode"] = args.StatusCode
	}
	s, err := ctx.RegisterResource("aws:apigateway/integrationResponse:IntegrationResponse", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IntegrationResponse{s: s}, nil
}

// GetIntegrationResponse gets an existing IntegrationResponse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationResponse(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IntegrationResponseState, opts ...pulumi.ResourceOpt) (*IntegrationResponse, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["contentHandling"] = state.ContentHandling
		inputs["httpMethod"] = state.HttpMethod
		inputs["resourceId"] = state.ResourceId
		inputs["responseParameters"] = state.ResponseParameters
		inputs["responseTemplates"] = state.ResponseTemplates
		inputs["restApi"] = state.RestApi
		inputs["selectionPattern"] = state.SelectionPattern
		inputs["statusCode"] = state.StatusCode
	}
	s, err := ctx.ReadResource("aws:apigateway/integrationResponse:IntegrationResponse", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IntegrationResponse{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IntegrationResponse) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IntegrationResponse) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
func (r *IntegrationResponse) ContentHandling() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["contentHandling"])
}

// The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
func (r *IntegrationResponse) HttpMethod() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["httpMethod"])
}

// The API resource ID
func (r *IntegrationResponse) ResourceId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceId"])
}

// A map of response parameters that can be read from the backend response.
// For example: `responseParameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`
func (r *IntegrationResponse) ResponseParameters() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["responseParameters"])
}

// A map specifying the templates used to transform the integration response body
func (r *IntegrationResponse) ResponseTemplates() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["responseTemplates"])
}

// The ID of the associated REST API
func (r *IntegrationResponse) RestApi() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["restApi"])
}

// Specifies the regular expression pattern used to choose
// an integration response based on the response from the backend. Setting this to `-` makes the integration the default one.
// If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
// For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
func (r *IntegrationResponse) SelectionPattern() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["selectionPattern"])
}

// The HTTP status code
func (r *IntegrationResponse) StatusCode() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["statusCode"])
}

// Input properties used for looking up and filtering IntegrationResponse resources.
type IntegrationResponseState struct {
	// Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
	ContentHandling interface{}
	// The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod interface{}
	// The API resource ID
	ResourceId interface{}
	// A map of response parameters that can be read from the backend response.
	// For example: `responseParameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`
	ResponseParameters interface{}
	// A map specifying the templates used to transform the integration response body
	ResponseTemplates interface{}
	// The ID of the associated REST API
	RestApi interface{}
	// Specifies the regular expression pattern used to choose
	// an integration response based on the response from the backend. Setting this to `-` makes the integration the default one.
	// If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
	// For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
	SelectionPattern interface{}
	// The HTTP status code
	StatusCode interface{}
}

// The set of arguments for constructing a IntegrationResponse resource.
type IntegrationResponseArgs struct {
	// Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
	ContentHandling interface{}
	// The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod interface{}
	// The API resource ID
	ResourceId interface{}
	// A map of response parameters that can be read from the backend response.
	// For example: `responseParameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`
	ResponseParameters interface{}
	// A map specifying the templates used to transform the integration response body
	ResponseTemplates interface{}
	// The ID of the associated REST API
	RestApi interface{}
	// Specifies the regular expression pattern used to choose
	// an integration response based on the response from the backend. Setting this to `-` makes the integration the default one.
	// If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
	// For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
	SelectionPattern interface{}
	// The HTTP status code
	StatusCode interface{}
}
