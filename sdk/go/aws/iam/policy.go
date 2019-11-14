// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an IAM policy.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_policy.html.markdown.
type Policy struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN assigned by AWS to this policy.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// Description of the IAM policy.
	Description pulumi.StringOutput `pulumi:"description"`

	// The name of the policy. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`

	// Path in which to create the policy.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path pulumi.StringOutput `pulumi:"path"`

	// The policy document. This is a JSON formatted string.
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOpt) (*Policy, error) {
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["path"] = args.Path
		inputs["policy"] = args.Policy
	}
	var resource Policy
	err := ctx.RegisterResource("aws:iam/policy:Policy", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PolicyState, opts ...pulumi.ResourceOpt) (*Policy, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["path"] = state.Path
		inputs["policy"] = state.Policy
	}
	var resource Policy
	err := ctx.ReadResource("aws:iam/policy:Policy", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *Policy) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *Policy) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering Policy resources.
type PolicyState struct {
	// The ARN assigned by AWS to this policy.
	Arn pulumi.StringInput `pulumi:"arn"`
	// Description of the IAM policy.
	Description pulumi.StringInput `pulumi:"description"`
	// The name of the policy. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// Path in which to create the policy.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path pulumi.StringInput `pulumi:"path"`
	// The policy document. This is a JSON formatted string.
	Policy pulumi.StringInput `pulumi:"policy"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Description of the IAM policy.
	Description pulumi.StringInput `pulumi:"description"`
	// The name of the policy. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// Path in which to create the policy.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path pulumi.StringInput `pulumi:"path"`
	// The policy document. This is a JSON formatted string.
	Policy pulumi.StringInput `pulumi:"policy"`
}
