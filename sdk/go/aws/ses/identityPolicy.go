// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a SES Identity Policy. More information about SES Sending Authorization Policies can be found in the [SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-policies.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_identity_policy.html.markdown.
type IdentityPolicy struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// Name or Amazon Resource Name (ARN) of the SES Identity.
	Identity pulumi.StringOutput `pulumi:"identity"`

	// Name of the policy.
	Name pulumi.StringOutput `pulumi:"name"`

	// JSON string of the policy.
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewIdentityPolicy registers a new resource with the given unique name, arguments, and options.
func NewIdentityPolicy(ctx *pulumi.Context,
	name string, args *IdentityPolicyArgs, opts ...pulumi.ResourceOpt) (*IdentityPolicy, error) {
	if args == nil || args.Identity == nil {
		return nil, errors.New("missing required argument 'Identity'")
	}
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["identity"] = args.Identity
		inputs["name"] = args.Name
		inputs["policy"] = args.Policy
	}
	var resource IdentityPolicy
	err := ctx.RegisterResource("aws:ses/identityPolicy:IdentityPolicy", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityPolicy gets an existing IdentityPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IdentityPolicyState, opts ...pulumi.ResourceOpt) (*IdentityPolicy, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["identity"] = state.Identity
		inputs["name"] = state.Name
		inputs["policy"] = state.Policy
	}
	var resource IdentityPolicy
	err := ctx.ReadResource("aws:ses/identityPolicy:IdentityPolicy", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *IdentityPolicy) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *IdentityPolicy) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering IdentityPolicy resources.
type IdentityPolicyState struct {
	// Name or Amazon Resource Name (ARN) of the SES Identity.
	Identity pulumi.StringInput `pulumi:"identity"`
	// Name of the policy.
	Name pulumi.StringInput `pulumi:"name"`
	// JSON string of the policy.
	Policy pulumi.StringInput `pulumi:"policy"`
}

// The set of arguments for constructing a IdentityPolicy resource.
type IdentityPolicyArgs struct {
	// Name or Amazon Resource Name (ARN) of the SES Identity.
	Identity pulumi.StringInput `pulumi:"identity"`
	// Name of the policy.
	Name pulumi.StringInput `pulumi:"name"`
	// JSON string of the policy.
	Policy pulumi.StringInput `pulumi:"policy"`
}
