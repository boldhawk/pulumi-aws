// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a settings of an API Gateway Account. Settings is applied region-wide per `provider` block.
// 
// > **Note:** As there is no API method for deleting account settings or resetting it to defaults, destroying this resource will keep your account settings intact
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_account.html.markdown.
type Account struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN of an IAM role for CloudWatch (to allow logging & monitoring).
	// See more [in AWS Docs](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-stage-settings.html#how-to-stage-settings-console).
	// Logging & monitoring can be enabled/disabled and otherwise tuned on the API Gateway Stage level.
	CloudwatchRoleArn pulumi.StringOutput `pulumi:"cloudwatchRoleArn"`

	// Account-Level throttle settings. See exported fields below.
	ThrottleSettings pulumi.AnyOutput `pulumi:"throttleSettings"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOpt) (*Account, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["cloudwatchRoleArn"] = args.CloudwatchRoleArn
	}
	var resource Account
	err := ctx.RegisterResource("aws:apigateway/account:Account", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AccountState, opts ...pulumi.ResourceOpt) (*Account, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["cloudwatchRoleArn"] = state.CloudwatchRoleArn
		inputs["throttleSettings"] = state.ThrottleSettings
	}
	var resource Account
	err := ctx.ReadResource("aws:apigateway/account:Account", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *Account) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *Account) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering Account resources.
type AccountState struct {
	// The ARN of an IAM role for CloudWatch (to allow logging & monitoring).
	// See more [in AWS Docs](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-stage-settings.html#how-to-stage-settings-console).
	// Logging & monitoring can be enabled/disabled and otherwise tuned on the API Gateway Stage level.
	CloudwatchRoleArn pulumi.StringInput `pulumi:"cloudwatchRoleArn"`
	// Account-Level throttle settings. See exported fields below.
	ThrottleSettings pulumi.AnyInput `pulumi:"throttleSettings"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The ARN of an IAM role for CloudWatch (to allow logging & monitoring).
	// See more [in AWS Docs](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-stage-settings.html#how-to-stage-settings-console).
	// Logging & monitoring can be enabled/disabled and otherwise tuned on the API Gateway Stage level.
	CloudwatchRoleArn pulumi.StringInput `pulumi:"cloudwatchRoleArn"`
}
