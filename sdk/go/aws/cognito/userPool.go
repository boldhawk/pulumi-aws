// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Cognito User Pool resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cognito_user_pool.html.markdown.
type UserPool struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The configuration for AdminCreateUser requests.
	AdminCreateUserConfig pulumi.AnyOutput `pulumi:"adminCreateUserConfig"`

	// Attributes supported as an alias for this user pool. Possible values: phone_number, email, or preferred_username. Conflicts with `usernameAttributes`.
	AliasAttributes pulumi.ArrayOutput `pulumi:"aliasAttributes"`

	// The ARN of the user pool.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The attributes to be auto-verified. Possible values: email, phone_number.
	AutoVerifiedAttributes pulumi.ArrayOutput `pulumi:"autoVerifiedAttributes"`

	// The date the user pool was created.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`

	// The configuration for the user pool's device tracking.
	DeviceConfiguration pulumi.AnyOutput `pulumi:"deviceConfiguration"`

	// The Email Configuration.
	EmailConfiguration pulumi.AnyOutput `pulumi:"emailConfiguration"`

	// A string representing the email verification message. Conflicts with `verificationMessageTemplate` configuration block `emailMessage` argument.
	EmailVerificationMessage pulumi.StringOutput `pulumi:"emailVerificationMessage"`

	// A string representing the email verification subject. Conflicts with `verificationMessageTemplate` configuration block `emailSubject` argument.
	EmailVerificationSubject pulumi.StringOutput `pulumi:"emailVerificationSubject"`

	// The endpoint name of the user pool. Example format: cognito-idp.REGION.amazonaws.com/xxxx_yyyyy
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`

	// A container for the AWS Lambda triggers associated with the user pool.
	LambdaConfig pulumi.AnyOutput `pulumi:"lambdaConfig"`

	// The date the user pool was last modified.
	LastModifiedDate pulumi.StringOutput `pulumi:"lastModifiedDate"`

	// Set to enable multi-factor authentication. Must be one of the following values (ON, OFF, OPTIONAL)
	MfaConfiguration pulumi.StringOutput `pulumi:"mfaConfiguration"`

	// The name of the attribute.
	Name pulumi.StringOutput `pulumi:"name"`

	// A container for information about the user pool password policy.
	PasswordPolicy pulumi.AnyOutput `pulumi:"passwordPolicy"`

	// A container with the schema attributes of a user pool. Maximum of 50 attributes.
	Schemas pulumi.ArrayOutput `pulumi:"schemas"`

	// A string representing the SMS authentication message.
	SmsAuthenticationMessage pulumi.StringOutput `pulumi:"smsAuthenticationMessage"`

	// The SMS Configuration.
	SmsConfiguration pulumi.AnyOutput `pulumi:"smsConfiguration"`

	// A string representing the SMS verification message. Conflicts with `verificationMessageTemplate` configuration block `smsMessage` argument.
	SmsVerificationMessage pulumi.StringOutput `pulumi:"smsVerificationMessage"`

	// A mapping of tags to assign to the User Pool.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// Configuration block for user pool add-ons to enable user pool advanced security mode features.
	UserPoolAddOns pulumi.AnyOutput `pulumi:"userPoolAddOns"`

	// Specifies whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `aliasAttributes`.
	UsernameAttributes pulumi.ArrayOutput `pulumi:"usernameAttributes"`

	// The verification message templates configuration.
	VerificationMessageTemplate pulumi.AnyOutput `pulumi:"verificationMessageTemplate"`
}

// NewUserPool registers a new resource with the given unique name, arguments, and options.
func NewUserPool(ctx *pulumi.Context,
	name string, args *UserPoolArgs, opts ...pulumi.ResourceOpt) (*UserPool, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["adminCreateUserConfig"] = args.AdminCreateUserConfig
		inputs["aliasAttributes"] = args.AliasAttributes
		inputs["autoVerifiedAttributes"] = args.AutoVerifiedAttributes
		inputs["deviceConfiguration"] = args.DeviceConfiguration
		inputs["emailConfiguration"] = args.EmailConfiguration
		inputs["emailVerificationMessage"] = args.EmailVerificationMessage
		inputs["emailVerificationSubject"] = args.EmailVerificationSubject
		inputs["lambdaConfig"] = args.LambdaConfig
		inputs["mfaConfiguration"] = args.MfaConfiguration
		inputs["name"] = args.Name
		inputs["passwordPolicy"] = args.PasswordPolicy
		inputs["schemas"] = args.Schemas
		inputs["smsAuthenticationMessage"] = args.SmsAuthenticationMessage
		inputs["smsConfiguration"] = args.SmsConfiguration
		inputs["smsVerificationMessage"] = args.SmsVerificationMessage
		inputs["tags"] = args.Tags
		inputs["userPoolAddOns"] = args.UserPoolAddOns
		inputs["usernameAttributes"] = args.UsernameAttributes
		inputs["verificationMessageTemplate"] = args.VerificationMessageTemplate
	}
	var resource UserPool
	err := ctx.RegisterResource("aws:cognito/userPool:UserPool", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPool gets an existing UserPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPool(ctx *pulumi.Context,
	name string, id pulumi.ID, state *UserPoolState, opts ...pulumi.ResourceOpt) (*UserPool, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["adminCreateUserConfig"] = state.AdminCreateUserConfig
		inputs["aliasAttributes"] = state.AliasAttributes
		inputs["arn"] = state.Arn
		inputs["autoVerifiedAttributes"] = state.AutoVerifiedAttributes
		inputs["creationDate"] = state.CreationDate
		inputs["deviceConfiguration"] = state.DeviceConfiguration
		inputs["emailConfiguration"] = state.EmailConfiguration
		inputs["emailVerificationMessage"] = state.EmailVerificationMessage
		inputs["emailVerificationSubject"] = state.EmailVerificationSubject
		inputs["endpoint"] = state.Endpoint
		inputs["lambdaConfig"] = state.LambdaConfig
		inputs["lastModifiedDate"] = state.LastModifiedDate
		inputs["mfaConfiguration"] = state.MfaConfiguration
		inputs["name"] = state.Name
		inputs["passwordPolicy"] = state.PasswordPolicy
		inputs["schemas"] = state.Schemas
		inputs["smsAuthenticationMessage"] = state.SmsAuthenticationMessage
		inputs["smsConfiguration"] = state.SmsConfiguration
		inputs["smsVerificationMessage"] = state.SmsVerificationMessage
		inputs["tags"] = state.Tags
		inputs["userPoolAddOns"] = state.UserPoolAddOns
		inputs["usernameAttributes"] = state.UsernameAttributes
		inputs["verificationMessageTemplate"] = state.VerificationMessageTemplate
	}
	var resource UserPool
	err := ctx.ReadResource("aws:cognito/userPool:UserPool", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *UserPool) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *UserPool) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering UserPool resources.
type UserPoolState struct {
	// The configuration for AdminCreateUser requests.
	AdminCreateUserConfig pulumi.AnyInput `pulumi:"adminCreateUserConfig"`
	// Attributes supported as an alias for this user pool. Possible values: phone_number, email, or preferred_username. Conflicts with `usernameAttributes`.
	AliasAttributes pulumi.ArrayInput `pulumi:"aliasAttributes"`
	// The ARN of the user pool.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The attributes to be auto-verified. Possible values: email, phone_number.
	AutoVerifiedAttributes pulumi.ArrayInput `pulumi:"autoVerifiedAttributes"`
	// The date the user pool was created.
	CreationDate pulumi.StringInput `pulumi:"creationDate"`
	// The configuration for the user pool's device tracking.
	DeviceConfiguration pulumi.AnyInput `pulumi:"deviceConfiguration"`
	// The Email Configuration.
	EmailConfiguration pulumi.AnyInput `pulumi:"emailConfiguration"`
	// A string representing the email verification message. Conflicts with `verificationMessageTemplate` configuration block `emailMessage` argument.
	EmailVerificationMessage pulumi.StringInput `pulumi:"emailVerificationMessage"`
	// A string representing the email verification subject. Conflicts with `verificationMessageTemplate` configuration block `emailSubject` argument.
	EmailVerificationSubject pulumi.StringInput `pulumi:"emailVerificationSubject"`
	// The endpoint name of the user pool. Example format: cognito-idp.REGION.amazonaws.com/xxxx_yyyyy
	Endpoint pulumi.StringInput `pulumi:"endpoint"`
	// A container for the AWS Lambda triggers associated with the user pool.
	LambdaConfig pulumi.AnyInput `pulumi:"lambdaConfig"`
	// The date the user pool was last modified.
	LastModifiedDate pulumi.StringInput `pulumi:"lastModifiedDate"`
	// Set to enable multi-factor authentication. Must be one of the following values (ON, OFF, OPTIONAL)
	MfaConfiguration pulumi.StringInput `pulumi:"mfaConfiguration"`
	// The name of the attribute.
	Name pulumi.StringInput `pulumi:"name"`
	// A container for information about the user pool password policy.
	PasswordPolicy pulumi.AnyInput `pulumi:"passwordPolicy"`
	// A container with the schema attributes of a user pool. Maximum of 50 attributes.
	Schemas pulumi.ArrayInput `pulumi:"schemas"`
	// A string representing the SMS authentication message.
	SmsAuthenticationMessage pulumi.StringInput `pulumi:"smsAuthenticationMessage"`
	// The SMS Configuration.
	SmsConfiguration pulumi.AnyInput `pulumi:"smsConfiguration"`
	// A string representing the SMS verification message. Conflicts with `verificationMessageTemplate` configuration block `smsMessage` argument.
	SmsVerificationMessage pulumi.StringInput `pulumi:"smsVerificationMessage"`
	// A mapping of tags to assign to the User Pool.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Configuration block for user pool add-ons to enable user pool advanced security mode features.
	UserPoolAddOns pulumi.AnyInput `pulumi:"userPoolAddOns"`
	// Specifies whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `aliasAttributes`.
	UsernameAttributes pulumi.ArrayInput `pulumi:"usernameAttributes"`
	// The verification message templates configuration.
	VerificationMessageTemplate pulumi.AnyInput `pulumi:"verificationMessageTemplate"`
}

// The set of arguments for constructing a UserPool resource.
type UserPoolArgs struct {
	// The configuration for AdminCreateUser requests.
	AdminCreateUserConfig pulumi.AnyInput `pulumi:"adminCreateUserConfig"`
	// Attributes supported as an alias for this user pool. Possible values: phone_number, email, or preferred_username. Conflicts with `usernameAttributes`.
	AliasAttributes pulumi.ArrayInput `pulumi:"aliasAttributes"`
	// The attributes to be auto-verified. Possible values: email, phone_number.
	AutoVerifiedAttributes pulumi.ArrayInput `pulumi:"autoVerifiedAttributes"`
	// The configuration for the user pool's device tracking.
	DeviceConfiguration pulumi.AnyInput `pulumi:"deviceConfiguration"`
	// The Email Configuration.
	EmailConfiguration pulumi.AnyInput `pulumi:"emailConfiguration"`
	// A string representing the email verification message. Conflicts with `verificationMessageTemplate` configuration block `emailMessage` argument.
	EmailVerificationMessage pulumi.StringInput `pulumi:"emailVerificationMessage"`
	// A string representing the email verification subject. Conflicts with `verificationMessageTemplate` configuration block `emailSubject` argument.
	EmailVerificationSubject pulumi.StringInput `pulumi:"emailVerificationSubject"`
	// A container for the AWS Lambda triggers associated with the user pool.
	LambdaConfig pulumi.AnyInput `pulumi:"lambdaConfig"`
	// Set to enable multi-factor authentication. Must be one of the following values (ON, OFF, OPTIONAL)
	MfaConfiguration pulumi.StringInput `pulumi:"mfaConfiguration"`
	// The name of the attribute.
	Name pulumi.StringInput `pulumi:"name"`
	// A container for information about the user pool password policy.
	PasswordPolicy pulumi.AnyInput `pulumi:"passwordPolicy"`
	// A container with the schema attributes of a user pool. Maximum of 50 attributes.
	Schemas pulumi.ArrayInput `pulumi:"schemas"`
	// A string representing the SMS authentication message.
	SmsAuthenticationMessage pulumi.StringInput `pulumi:"smsAuthenticationMessage"`
	// The SMS Configuration.
	SmsConfiguration pulumi.AnyInput `pulumi:"smsConfiguration"`
	// A string representing the SMS verification message. Conflicts with `verificationMessageTemplate` configuration block `smsMessage` argument.
	SmsVerificationMessage pulumi.StringInput `pulumi:"smsVerificationMessage"`
	// A mapping of tags to assign to the User Pool.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Configuration block for user pool add-ons to enable user pool advanced security mode features.
	UserPoolAddOns pulumi.AnyInput `pulumi:"userPoolAddOns"`
	// Specifies whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `aliasAttributes`.
	UsernameAttributes pulumi.ArrayInput `pulumi:"usernameAttributes"`
	// The verification message templates configuration.
	VerificationMessageTemplate pulumi.AnyInput `pulumi:"verificationMessageTemplate"`
}
