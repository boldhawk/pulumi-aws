// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Decrypt multiple secrets from data encrypted with the AWS KMS service.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/kms_secrets.html.markdown.
func LookupSecrets(ctx *pulumi.Context, args *GetSecretsArgs) (*GetSecretsResult, error) {
	var rv GetSecretsResult
	err := ctx.Invoke("aws:kms/getSecrets:getSecrets", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecrets.
type GetSecretsArgs struct {
	// One or more encrypted payload definitions from the KMS service. See the Secret Definitions below.
	Secrets []interface{} `pulumi:"secrets"`
}

// A collection of values returned by getSecrets.
type GetSecretsResult struct {
	// Map containing each `secret` `name` as the key with its decrypted plaintext value
	Plaintext map[string]string `pulumi:"plaintext"`
	Secrets []interface{} `pulumi:"secrets"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
