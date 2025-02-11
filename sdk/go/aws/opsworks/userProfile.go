// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an OpsWorks User Profile resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/opsworks_user_profile.html.markdown.
type UserProfile struct {
	s *pulumi.ResourceState
}

// NewUserProfile registers a new resource with the given unique name, arguments, and options.
func NewUserProfile(ctx *pulumi.Context,
	name string, args *UserProfileArgs, opts ...pulumi.ResourceOpt) (*UserProfile, error) {
	if args == nil || args.SshUsername == nil {
		return nil, errors.New("missing required argument 'SshUsername'")
	}
	if args == nil || args.UserArn == nil {
		return nil, errors.New("missing required argument 'UserArn'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["allowSelfManagement"] = nil
		inputs["sshPublicKey"] = nil
		inputs["sshUsername"] = nil
		inputs["userArn"] = nil
	} else {
		inputs["allowSelfManagement"] = args.AllowSelfManagement
		inputs["sshPublicKey"] = args.SshPublicKey
		inputs["sshUsername"] = args.SshUsername
		inputs["userArn"] = args.UserArn
	}
	s, err := ctx.RegisterResource("aws:opsworks/userProfile:UserProfile", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &UserProfile{s: s}, nil
}

// GetUserProfile gets an existing UserProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserProfile(ctx *pulumi.Context,
	name string, id pulumi.ID, state *UserProfileState, opts ...pulumi.ResourceOpt) (*UserProfile, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["allowSelfManagement"] = state.AllowSelfManagement
		inputs["sshPublicKey"] = state.SshPublicKey
		inputs["sshUsername"] = state.SshUsername
		inputs["userArn"] = state.UserArn
	}
	s, err := ctx.ReadResource("aws:opsworks/userProfile:UserProfile", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &UserProfile{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *UserProfile) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *UserProfile) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Whether users can specify their own SSH public key through the My Settings page
func (r *UserProfile) AllowSelfManagement() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["allowSelfManagement"])
}

// The users public key
func (r *UserProfile) SshPublicKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sshPublicKey"])
}

// The ssh username, with witch this user wants to log in
func (r *UserProfile) SshUsername() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sshUsername"])
}

// The user's IAM ARN
func (r *UserProfile) UserArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["userArn"])
}

// Input properties used for looking up and filtering UserProfile resources.
type UserProfileState struct {
	// Whether users can specify their own SSH public key through the My Settings page
	AllowSelfManagement interface{}
	// The users public key
	SshPublicKey interface{}
	// The ssh username, with witch this user wants to log in
	SshUsername interface{}
	// The user's IAM ARN
	UserArn interface{}
}

// The set of arguments for constructing a UserProfile resource.
type UserProfileArgs struct {
	// Whether users can specify their own SSH public key through the My Settings page
	AllowSelfManagement interface{}
	// The users public key
	SshPublicKey interface{}
	// The ssh username, with witch this user wants to log in
	SshUsername interface{}
	// The user's IAM ARN
	UserArn interface{}
}
