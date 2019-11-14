// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the Account ID of the [AWS Billing and Cost Management Service Account](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-getting-started.html#step-2) for the purpose of whitelisting in S3 bucket policy.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/billing_service_account.html.markdown.
func LookupBillingServiceAccount(ctx *pulumi.Context) (*GetBillingServiceAccountResult, error) {
	var rv GetBillingServiceAccountResult
	err := ctx.Invoke("aws:index/getBillingServiceAccount:getBillingServiceAccount", nil, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getBillingServiceAccount.
type GetBillingServiceAccountResult struct {
	// The ARN of the AWS billing service account.
	Arn string `pulumi:"arn"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
