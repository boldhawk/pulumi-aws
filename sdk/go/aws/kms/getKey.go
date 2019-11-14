// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get detailed information about 
// the specified KMS Key with flexible key id input. 
// This can be useful to reference key alias 
// without having to hard code the ARN as input.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/kms_key.html.markdown.
func LookupKey(ctx *pulumi.Context, args *GetKeyArgs) (*GetKeyResult, error) {
var rv GetKeyResult
	err := ctx.Invoke("aws:kms/getKey:getKey", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKey.
type GetKeyArgs struct {
	// List of grant tokens
	GrantTokens []interface{} `pulumi:"grantTokens"`
	// Key identifier which can be one of the following format:
	// * Key ID. E.g: `1234abcd-12ab-34cd-56ef-1234567890ab`
	// * Key ARN. E.g.: `arn:aws:kms:us-east-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`
	// * Alias name. E.g.: `alias/my-key`
	// * Alias ARN: E.g.: `arn:aws:kms:us-east-1:111122223333:alias/my-key`
	KeyId string `pulumi:"keyId"`
}

// A collection of values returned by getKey.
type GetKeyResult struct {
	Arn string `pulumi:"arn"`
	AwsAccountId string `pulumi:"awsAccountId"`
	CreationDate string `pulumi:"creationDate"`
	DeletionDate string `pulumi:"deletionDate"`
	Description string `pulumi:"description"`
	Enabled bool `pulumi:"enabled"`
	ExpirationModel string `pulumi:"expirationModel"`
	GrantTokens []interface{} `pulumi:"grantTokens"`
	KeyId string `pulumi:"keyId"`
	KeyManager string `pulumi:"keyManager"`
	KeyState string `pulumi:"keyState"`
	KeyUsage string `pulumi:"keyUsage"`
	Origin string `pulumi:"origin"`
	ValidTo string `pulumi:"validTo"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
