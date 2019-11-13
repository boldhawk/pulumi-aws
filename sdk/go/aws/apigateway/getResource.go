// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the id of a Resource in API Gateway. 
// To fetch the Resource, you must provide the REST API id as well as the full path.  
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/api_gateway_resource.html.markdown.
func LookupResource(ctx *pulumi.Context, args *GetResourceArgs) (*GetResourceResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["path"] = args.Path
		inputs["restApiId"] = args.RestApiId
	}
	outputs, err := ctx.Invoke("aws:apigateway/getResource:getResource", inputs)
	if err != nil {
		return nil, err
	}
	return &GetResourceResult{
		ParentId: outputs["parentId"],
		Path: outputs["path"],
		PathPart: outputs["pathPart"],
		RestApiId: outputs["restApiId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getResource.
type GetResourceArgs struct {
	// The full path of the resource.  If no path is found, an error will be returned.
	Path pulumi.StringInput `pulumi:"path"`
	// The REST API id that owns the resource. If no REST API is found, an error will be returned.
	RestApiId pulumi.StringInput `pulumi:"restApiId"`
}

// A collection of values returned by getResource.
type GetResourceResult struct {
	// Set to the ID of the parent Resource.
	ParentId string `pulumi:"parentId"`
	Path string `pulumi:"path"`
	// Set to the path relative to the parent Resource.
	PathPart string `pulumi:"pathPart"`
	RestApiId string `pulumi:"restApiId"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
