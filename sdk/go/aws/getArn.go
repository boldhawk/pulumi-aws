// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Parses an Amazon Resource Name (ARN) into its constituent parts.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/arn.html.markdown.
func LookupArn(ctx *pulumi.Context, args *GetArnArgs) (*GetArnResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["arn"] = args.Arn
	}
	outputs, err := ctx.Invoke("aws:index/getArn:getArn", inputs)
	if err != nil {
		return nil, err
	}
	return &GetArnResult{
		Account: outputs["account"],
		Arn: outputs["arn"],
		Partition: outputs["partition"],
		Region: outputs["region"],
		Resource: outputs["resource"],
		Service: outputs["service"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getArn.
type GetArnArgs struct {
	// The ARN to parse.
	Arn pulumi.StringInput `pulumi:"arn"`
}

// A collection of values returned by getArn.
type GetArnResult struct {
	// The [ID](https://docs.aws.amazon.com/general/latest/gr/acct-identifiers.html) of the AWS account that owns the resource, without the hyphens.
	Account string `pulumi:"account"`
	Arn string `pulumi:"arn"`
	// The partition that the resource is in.
	Partition string `pulumi:"partition"`
	// The region the resource resides in.
	// Note that the ARNs for some resources do not require a region, so this component might be omitted.
	Region string `pulumi:"region"`
	// The content of this part of the ARN varies by service.
	// It often includes an indicator of the type of resource—for example, an IAM user or Amazon RDS database —followed by a slash (/) or a colon (:), followed by the resource name itself.
	Resource string `pulumi:"resource"`
	// The [service namespace](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces) that identifies the AWS product.
	Service string `pulumi:"service"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
