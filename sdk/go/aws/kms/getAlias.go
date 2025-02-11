// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the ARN of a KMS key alias.
// By using this data source, you can reference key alias
// without having to hard code the ARN as input.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/kms_alias.html.markdown.
func LookupAlias(ctx *pulumi.Context, args *GetAliasArgs) (*GetAliasResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("aws:kms/getAlias:getAlias", inputs)
	if err != nil {
		return nil, err
	}
	return &GetAliasResult{
		Arn: outputs["arn"],
		Name: outputs["name"],
		TargetKeyArn: outputs["targetKeyArn"],
		TargetKeyId: outputs["targetKeyId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getAlias.
type GetAliasArgs struct {
	// The display name of the alias. The name must start with the word "alias" followed by a forward slash (alias/)
	Name interface{}
}

// A collection of values returned by getAlias.
type GetAliasResult struct {
	// The Amazon Resource Name(ARN) of the key alias.
	Arn interface{}
	Name interface{}
	// ARN pointed to by the alias.
	TargetKeyArn interface{}
	// Key identifier pointed to by the alias.
	TargetKeyId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
