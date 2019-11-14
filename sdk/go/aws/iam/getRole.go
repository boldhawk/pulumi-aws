// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source can be used to fetch information about a specific
// IAM role. By using this data source, you can reference IAM role
// properties without having to hard code ARNs as input.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/iam_role.html.markdown.
func LookupRole(ctx *pulumi.Context, args *GetRoleArgs) (*GetRoleResult, error) {
	var rv GetRoleResult
	err := ctx.Invoke("aws:iam/getRole:getRole", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRole.
type GetRoleArgs struct {
	// The friendly IAM role name to match.
	Name string `pulumi:"name"`
}

// A collection of values returned by getRole.
type GetRoleResult struct {
	// The Amazon Resource Name (ARN) specifying the role.
	Arn string `pulumi:"arn"`
	// The policy document associated with the role.
	AssumeRolePolicy string `pulumi:"assumeRolePolicy"`
	// Creation date of the role in RFC 3339 format.
	CreateDate string `pulumi:"createDate"`
	// Description for the role.
	Description string `pulumi:"description"`
	// Maximum session duration.
	MaxSessionDuration int `pulumi:"maxSessionDuration"`
	Name string `pulumi:"name"`
	// The path to the role.
	Path string `pulumi:"path"`
	// The ARN of the policy that is used to set the permissions boundary for the role.
	PermissionsBoundary string `pulumi:"permissionsBoundary"`
	// The stable and unique string identifying the role.
	UniqueId string `pulumi:"uniqueId"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
