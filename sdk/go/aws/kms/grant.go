// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource-based access control mechanism for a KMS customer master key.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/kms_grant.html.markdown.
type Grant struct {
	s *pulumi.ResourceState
}

// NewGrant registers a new resource with the given unique name, arguments, and options.
func NewGrant(ctx *pulumi.Context,
	name string, args *GrantArgs, opts ...pulumi.ResourceOpt) (*Grant, error) {
	if args == nil || args.GranteePrincipal == nil {
		return nil, errors.New("missing required argument 'GranteePrincipal'")
	}
	if args == nil || args.KeyId == nil {
		return nil, errors.New("missing required argument 'KeyId'")
	}
	if args == nil || args.Operations == nil {
		return nil, errors.New("missing required argument 'Operations'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["constraints"] = nil
		inputs["grantCreationTokens"] = nil
		inputs["granteePrincipal"] = nil
		inputs["keyId"] = nil
		inputs["name"] = nil
		inputs["operations"] = nil
		inputs["retireOnDelete"] = nil
		inputs["retiringPrincipal"] = nil
	} else {
		inputs["constraints"] = args.Constraints
		inputs["grantCreationTokens"] = args.GrantCreationTokens
		inputs["granteePrincipal"] = args.GranteePrincipal
		inputs["keyId"] = args.KeyId
		inputs["name"] = args.Name
		inputs["operations"] = args.Operations
		inputs["retireOnDelete"] = args.RetireOnDelete
		inputs["retiringPrincipal"] = args.RetiringPrincipal
	}
	inputs["grantId"] = nil
	inputs["grantToken"] = nil
	s, err := ctx.RegisterResource("aws:kms/grant:Grant", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Grant{s: s}, nil
}

// GetGrant gets an existing Grant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGrant(ctx *pulumi.Context,
	name string, id pulumi.ID, state *GrantState, opts ...pulumi.ResourceOpt) (*Grant, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["constraints"] = state.Constraints
		inputs["grantCreationTokens"] = state.GrantCreationTokens
		inputs["grantId"] = state.GrantId
		inputs["grantToken"] = state.GrantToken
		inputs["granteePrincipal"] = state.GranteePrincipal
		inputs["keyId"] = state.KeyId
		inputs["name"] = state.Name
		inputs["operations"] = state.Operations
		inputs["retireOnDelete"] = state.RetireOnDelete
		inputs["retiringPrincipal"] = state.RetiringPrincipal
	}
	s, err := ctx.ReadResource("aws:kms/grant:Grant", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Grant{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Grant) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Grant) ID() pulumi.IDOutput {
	return r.s.ID()
}

// A structure that you can use to allow certain operations in the grant only when the desired encryption context is present. For more information about encryption context, see [Encryption Context](http://docs.aws.amazon.com/kms/latest/developerguide/encryption-context.html).
func (r *Grant) Constraints() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["constraints"])
}

// A list of grant tokens to be used when creating the grant. See [Grant Tokens](http://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token) for more information about grant tokens.
// * `retireOnDelete` -(Defaults to false, Forces new resources) If set to false (the default) the grants will be revoked upon deletion, and if set to true the grants will try to be retired upon deletion. Note that retiring grants requires special permissions, hence why we default to revoking grants.
// See [RetireGrant](https://docs.aws.amazon.com/kms/latest/APIReference/API_RetireGrant.html) for more information.
func (r *Grant) GrantCreationTokens() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["grantCreationTokens"])
}

// The unique identifier for the grant.
func (r *Grant) GrantId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["grantId"])
}

// The grant token for the created grant. For more information, see [Grant Tokens](http://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token).
func (r *Grant) GrantToken() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["grantToken"])
}

// The principal that is given permission to perform the operations that the grant permits in ARN format. Note that due to eventual consistency issues around IAM principals, the state may not always be refreshed to reflect what is true in AWS.
func (r *Grant) GranteePrincipal() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["granteePrincipal"])
}

// The unique identifier for the customer master key (CMK) that the grant applies to. Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To specify a CMK in a different AWS account, you must use the key ARN.
func (r *Grant) KeyId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["keyId"])
}

// A friendly name for identifying the grant.
func (r *Grant) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// A list of operations that the grant permits. The permitted values are: `Decrypt, Encrypt, GenerateDataKey, GenerateDataKeyWithoutPlaintext, ReEncryptFrom, ReEncryptTo, CreateGrant, RetireGrant, DescribeKey`
func (r *Grant) Operations() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["operations"])
}

func (r *Grant) RetireOnDelete() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["retireOnDelete"])
}

// The principal that is given permission to retire the grant by using RetireGrant operation in ARN format. Note that due to eventual consistency issues around IAM principals, the state may not always be refreshed to reflect what is true in AWS.
func (r *Grant) RetiringPrincipal() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["retiringPrincipal"])
}

// Input properties used for looking up and filtering Grant resources.
type GrantState struct {
	// A structure that you can use to allow certain operations in the grant only when the desired encryption context is present. For more information about encryption context, see [Encryption Context](http://docs.aws.amazon.com/kms/latest/developerguide/encryption-context.html).
	Constraints interface{}
	// A list of grant tokens to be used when creating the grant. See [Grant Tokens](http://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token) for more information about grant tokens.
	// * `retireOnDelete` -(Defaults to false, Forces new resources) If set to false (the default) the grants will be revoked upon deletion, and if set to true the grants will try to be retired upon deletion. Note that retiring grants requires special permissions, hence why we default to revoking grants.
	// See [RetireGrant](https://docs.aws.amazon.com/kms/latest/APIReference/API_RetireGrant.html) for more information.
	GrantCreationTokens interface{}
	// The unique identifier for the grant.
	GrantId interface{}
	// The grant token for the created grant. For more information, see [Grant Tokens](http://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token).
	GrantToken interface{}
	// The principal that is given permission to perform the operations that the grant permits in ARN format. Note that due to eventual consistency issues around IAM principals, the state may not always be refreshed to reflect what is true in AWS.
	GranteePrincipal interface{}
	// The unique identifier for the customer master key (CMK) that the grant applies to. Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To specify a CMK in a different AWS account, you must use the key ARN.
	KeyId interface{}
	// A friendly name for identifying the grant.
	Name interface{}
	// A list of operations that the grant permits. The permitted values are: `Decrypt, Encrypt, GenerateDataKey, GenerateDataKeyWithoutPlaintext, ReEncryptFrom, ReEncryptTo, CreateGrant, RetireGrant, DescribeKey`
	Operations interface{}
	RetireOnDelete interface{}
	// The principal that is given permission to retire the grant by using RetireGrant operation in ARN format. Note that due to eventual consistency issues around IAM principals, the state may not always be refreshed to reflect what is true in AWS.
	RetiringPrincipal interface{}
}

// The set of arguments for constructing a Grant resource.
type GrantArgs struct {
	// A structure that you can use to allow certain operations in the grant only when the desired encryption context is present. For more information about encryption context, see [Encryption Context](http://docs.aws.amazon.com/kms/latest/developerguide/encryption-context.html).
	Constraints interface{}
	// A list of grant tokens to be used when creating the grant. See [Grant Tokens](http://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token) for more information about grant tokens.
	// * `retireOnDelete` -(Defaults to false, Forces new resources) If set to false (the default) the grants will be revoked upon deletion, and if set to true the grants will try to be retired upon deletion. Note that retiring grants requires special permissions, hence why we default to revoking grants.
	// See [RetireGrant](https://docs.aws.amazon.com/kms/latest/APIReference/API_RetireGrant.html) for more information.
	GrantCreationTokens interface{}
	// The principal that is given permission to perform the operations that the grant permits in ARN format. Note that due to eventual consistency issues around IAM principals, the state may not always be refreshed to reflect what is true in AWS.
	GranteePrincipal interface{}
	// The unique identifier for the customer master key (CMK) that the grant applies to. Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To specify a CMK in a different AWS account, you must use the key ARN.
	KeyId interface{}
	// A friendly name for identifying the grant.
	Name interface{}
	// A list of operations that the grant permits. The permitted values are: `Decrypt, Encrypt, GenerateDataKey, GenerateDataKeyWithoutPlaintext, ReEncryptFrom, ReEncryptTo, CreateGrant, RetireGrant, DescribeKey`
	Operations interface{}
	RetireOnDelete interface{}
	// The principal that is given permission to retire the grant by using RetireGrant operation in ARN format. Note that due to eventual consistency issues around IAM principals, the state may not always be refreshed to reflect what is true in AWS.
	RetiringPrincipal interface{}
}
