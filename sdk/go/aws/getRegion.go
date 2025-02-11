// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `.getRegion` provides details about a specific AWS region.
// 
// As well as validating a given region name this resource can be used to
// discover the name of the region configured within the provider. The latter
// can be useful in a child module which is inheriting an AWS provider
// configuration from its parent module.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/region.html.markdown.
func LookupRegion(ctx *pulumi.Context, args *GetRegionArgs) (*GetRegionResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["endpoint"] = args.Endpoint
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("aws:index/getRegion:getRegion", inputs)
	if err != nil {
		return nil, err
	}
	return &GetRegionResult{
		Description: outputs["description"],
		Endpoint: outputs["endpoint"],
		Name: outputs["name"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getRegion.
type GetRegionArgs struct {
	// The EC2 endpoint of the region to select.
	Endpoint interface{}
	// The full name of the region to select.
	Name interface{}
}

// A collection of values returned by getRegion.
type GetRegionResult struct {
	// The region's description in this format: "Location (Region name)".
	Description interface{}
	// The EC2 endpoint for the selected region.
	Endpoint interface{}
	// The name of the selected region.
	Name interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
