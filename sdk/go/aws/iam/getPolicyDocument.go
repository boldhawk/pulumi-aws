// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Generates an IAM policy document in JSON format.
// 
// This is a data source which can be used to construct a JSON representation of
// an IAM policy document, for use with resources which expect policy documents,
// such as the `iam.Policy` resource.
// 
// 
// Using this data source to generate policy documents is *optional*. It is also
// valid to use literal JSON strings within your configuration, or to use the
// `file` interpolation function to read a raw JSON policy document from a file.
// 
// ## Context Variable Interpolation
// 
// The IAM policy document format allows context variables to be interpolated
// into various strings within a statement. The native IAM policy document format
// uses `${...}`-style syntax that is in conflict with interpolation
// syntax, so this data source instead uses `&{...}` syntax for interpolations that
// should be processed by AWS rather than by this provider.
// 
// ## Wildcard Principal
// 
// In order to define wildcard principal (a.k.a. anonymous user) use `type = "*"` and
// `identifiers = ["*"]`. In that case the rendered json will contain `"Principal": "*"`.
// Note, that even though the [IAM Documentation](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html)
// states that `"Principal": "*"` and `"Principal": {"AWS": "*"}` are equivalent,
// those principals have different behavior for IAM Role Trust Policy. Therefore
// this provider will normalize the principal field only in above-mentioned case and principals
// like `type = "AWS"` and `identifiers = ["*"]` will be rendered as `"Principal": {"AWS": "*"}`.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/iam_policy_document.html.markdown.
func LookupPolicyDocument(ctx *pulumi.Context, args *GetPolicyDocumentArgs) (*GetPolicyDocumentResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["overrideJson"] = args.OverrideJson
		inputs["policyId"] = args.PolicyId
		inputs["sourceJson"] = args.SourceJson
		inputs["statements"] = args.Statements
		inputs["version"] = args.Version
	}
	outputs, err := ctx.Invoke("aws:iam/getPolicyDocument:getPolicyDocument", inputs)
	if err != nil {
		return nil, err
	}
	return &GetPolicyDocumentResult{
		Json: outputs["json"],
		OverrideJson: outputs["overrideJson"],
		PolicyId: outputs["policyId"],
		SourceJson: outputs["sourceJson"],
		Statements: outputs["statements"],
		Version: outputs["version"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getPolicyDocument.
type GetPolicyDocumentArgs struct {
	// An IAM policy document to import and override the
	// current policy document.  Statements with non-blank `sid`s in the override
	// document will overwrite statements with the same `sid` in the current document.
	// Statements without an `sid` cannot be overwritten.
	OverrideJson pulumi.StringInput `pulumi:"overrideJson"`
	// An ID for the policy document.
	PolicyId pulumi.StringInput `pulumi:"policyId"`
	// An IAM policy document to import as a base for the
	// current policy document.  Statements with non-blank `sid`s in the current
	// policy document will overwrite statements with the same `sid` in the source
	// json.  Statements without an `sid` cannot be overwritten.
	SourceJson pulumi.StringInput `pulumi:"sourceJson"`
	// A nested configuration block (described below)
	// configuring one *statement* to be included in the policy document.
	Statements pulumi.ArrayInput `pulumi:"statements"`
	// IAM policy document version. Valid values: `2008-10-17`, `2012-10-17`. Defaults to `2012-10-17`. For more information, see the [AWS IAM User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_version.html).
	Version pulumi.StringInput `pulumi:"version"`
}

// A collection of values returned by getPolicyDocument.
type GetPolicyDocumentResult struct {
	// The above arguments serialized as a standard JSON policy document.
	Json string `pulumi:"json"`
	OverrideJson string `pulumi:"overrideJson"`
	PolicyId string `pulumi:"policyId"`
	SourceJson string `pulumi:"sourceJson"`
	Statements []interface{} `pulumi:"statements"`
	Version string `pulumi:"version"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
