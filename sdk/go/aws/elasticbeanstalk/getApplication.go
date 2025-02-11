// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Retrieve information about an Elastic Beanstalk Application.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elastic_beanstalk_application.html.markdown.
func LookupApplication(ctx *pulumi.Context, args *GetApplicationArgs) (*GetApplicationResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("aws:elasticbeanstalk/getApplication:getApplication", inputs)
	if err != nil {
		return nil, err
	}
	return &GetApplicationResult{
		AppversionLifecycle: outputs["appversionLifecycle"],
		Arn: outputs["arn"],
		Description: outputs["description"],
		Name: outputs["name"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getApplication.
type GetApplicationArgs struct {
	// The name of the application
	Name interface{}
}

// A collection of values returned by getApplication.
type GetApplicationResult struct {
	AppversionLifecycle interface{}
	// The Amazon Resource Name (ARN) of the application.
	Arn interface{}
	// Short description of the application
	Description interface{}
	Name interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
