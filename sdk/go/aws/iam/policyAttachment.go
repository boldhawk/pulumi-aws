// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Attaches a Managed IAM Policy to user(s), role(s), and/or group(s)
// 
// !> **WARNING:** The iam.PolicyAttachment resource creates **exclusive** attachments of IAM policies. Across the entire AWS account, all of the users/roles/groups to which a single policy is attached must be declared by a single iam.PolicyAttachment resource. This means that even any users/roles/groups that have the attached policy via any other mechanism (including other resources managed by this provider) will have that attached policy revoked by this resource. Consider `iam.RolePolicyAttachment`, `iam.UserPolicyAttachment`, or `iam.GroupPolicyAttachment` instead. These resources do not enforce exclusive attachment of an IAM policy.
// 
// > **NOTE:** The usage of this resource conflicts with the `iam.GroupPolicyAttachment`, `iam.RolePolicyAttachment`, and `iam.UserPolicyAttachment` resources and will permanently show a difference if both are defined.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_policy_attachment.html.markdown.
type PolicyAttachment struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The group(s) the policy should be applied to
	Groups pulumi.ArrayOutput `pulumi:"groups"`

	// The name of the attachment. This cannot be an empty string.
	Name pulumi.StringOutput `pulumi:"name"`

	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringOutput `pulumi:"policyArn"`

	// The role(s) the policy should be applied to
	Roles pulumi.ArrayOutput `pulumi:"roles"`

	// The user(s) the policy should be applied to
	Users pulumi.ArrayOutput `pulumi:"users"`
}

// NewPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewPolicyAttachment(ctx *pulumi.Context,
	name string, args *PolicyAttachmentArgs, opts ...pulumi.ResourceOpt) (*PolicyAttachment, error) {
	if args == nil || args.PolicyArn == nil {
		return nil, errors.New("missing required argument 'PolicyArn'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["groups"] = args.Groups
		inputs["name"] = args.Name
		inputs["policyArn"] = args.PolicyArn
		inputs["roles"] = args.Roles
		inputs["users"] = args.Users
	}
	var resource PolicyAttachment
	err := ctx.RegisterResource("aws:iam/policyAttachment:PolicyAttachment", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyAttachment gets an existing PolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PolicyAttachmentState, opts ...pulumi.ResourceOpt) (*PolicyAttachment, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["groups"] = state.Groups
		inputs["name"] = state.Name
		inputs["policyArn"] = state.PolicyArn
		inputs["roles"] = state.Roles
		inputs["users"] = state.Users
	}
	var resource PolicyAttachment
	err := ctx.ReadResource("aws:iam/policyAttachment:PolicyAttachment", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *PolicyAttachment) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *PolicyAttachment) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering PolicyAttachment resources.
type PolicyAttachmentState struct {
	// The group(s) the policy should be applied to
	Groups pulumi.ArrayInput `pulumi:"groups"`
	// The name of the attachment. This cannot be an empty string.
	Name pulumi.StringInput `pulumi:"name"`
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringInput `pulumi:"policyArn"`
	// The role(s) the policy should be applied to
	Roles pulumi.ArrayInput `pulumi:"roles"`
	// The user(s) the policy should be applied to
	Users pulumi.ArrayInput `pulumi:"users"`
}

// The set of arguments for constructing a PolicyAttachment resource.
type PolicyAttachmentArgs struct {
	// The group(s) the policy should be applied to
	Groups pulumi.ArrayInput `pulumi:"groups"`
	// The name of the attachment. This cannot be an empty string.
	Name pulumi.StringInput `pulumi:"name"`
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringInput `pulumi:"policyArn"`
	// The role(s) the policy should be applied to
	Roles pulumi.ArrayInput `pulumi:"roles"`
	// The user(s) the policy should be applied to
	Users pulumi.ArrayInput `pulumi:"users"`
}
