// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage the default AWS Network ACL. VPC Only.
// 
// Each VPC created in AWS comes with a Default Network ACL that can be managed, but not
// destroyed. **This is an advanced resource**, and has special caveats to be aware
// of when using it. Please read this document in its entirety before using this
// resource.
// 
// The `ec2.DefaultNetworkAcl` behaves differently from normal resources, in that
// this provider does not _create_ this resource, but instead attempts to "adopt" it
// into management. We can do this because each VPC created has a Default Network
// ACL that cannot be destroyed, and is created with a known set of default rules.
// 
// When this provider first adopts the Default Network ACL, it **immediately removes all
// rules in the ACL**. It then proceeds to create any rules specified in the
// configuration. This step is required so that only the rules specified in the
// configuration are created.
// 
// This resource treats its inline rules as absolute; only the rules defined
// inline are created, and any additions/removals external to this resource will
// result in diffs being shown. For these reasons, this resource is incompatible with the
// `ec2.NetworkAclRule` resource.
// 
// For more information about Network ACLs, see the AWS Documentation on
// [Network ACLs][aws-network-acls].
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/default_network_acl.html.markdown.
type DefaultNetworkAcl struct {
	s *pulumi.ResourceState
}

// NewDefaultNetworkAcl registers a new resource with the given unique name, arguments, and options.
func NewDefaultNetworkAcl(ctx *pulumi.Context,
	name string, args *DefaultNetworkAclArgs, opts ...pulumi.ResourceOpt) (*DefaultNetworkAcl, error) {
	if args == nil || args.DefaultNetworkAclId == nil {
		return nil, errors.New("missing required argument 'DefaultNetworkAclId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["defaultNetworkAclId"] = nil
		inputs["egress"] = nil
		inputs["ingress"] = nil
		inputs["subnetIds"] = nil
		inputs["tags"] = nil
	} else {
		inputs["defaultNetworkAclId"] = args.DefaultNetworkAclId
		inputs["egress"] = args.Egress
		inputs["ingress"] = args.Ingress
		inputs["subnetIds"] = args.SubnetIds
		inputs["tags"] = args.Tags
	}
	inputs["ownerId"] = nil
	inputs["vpcId"] = nil
	s, err := ctx.RegisterResource("aws:ec2/defaultNetworkAcl:DefaultNetworkAcl", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DefaultNetworkAcl{s: s}, nil
}

// GetDefaultNetworkAcl gets an existing DefaultNetworkAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultNetworkAcl(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DefaultNetworkAclState, opts ...pulumi.ResourceOpt) (*DefaultNetworkAcl, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["defaultNetworkAclId"] = state.DefaultNetworkAclId
		inputs["egress"] = state.Egress
		inputs["ingress"] = state.Ingress
		inputs["ownerId"] = state.OwnerId
		inputs["subnetIds"] = state.SubnetIds
		inputs["tags"] = state.Tags
		inputs["vpcId"] = state.VpcId
	}
	s, err := ctx.ReadResource("aws:ec2/defaultNetworkAcl:DefaultNetworkAcl", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DefaultNetworkAcl{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DefaultNetworkAcl) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DefaultNetworkAcl) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The Network ACL ID to manage. This
// attribute is exported from `ec2.Vpc`, or manually found via the AWS Console.
func (r *DefaultNetworkAcl) DefaultNetworkAclId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["defaultNetworkAclId"])
}

// Specifies an egress rule. Parameters defined below.
func (r *DefaultNetworkAcl) Egress() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["egress"])
}

// Specifies an ingress rule. Parameters defined below.
func (r *DefaultNetworkAcl) Ingress() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["ingress"])
}

// The ID of the AWS account that owns the Default Network ACL
func (r *DefaultNetworkAcl) OwnerId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ownerId"])
}

// A list of Subnet IDs to apply the ACL to. See the
// notes below on managing Subnets in the Default Network ACL
func (r *DefaultNetworkAcl) SubnetIds() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["subnetIds"])
}

// A mapping of tags to assign to the resource.
func (r *DefaultNetworkAcl) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// The ID of the associated VPC
func (r *DefaultNetworkAcl) VpcId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["vpcId"])
}

// Input properties used for looking up and filtering DefaultNetworkAcl resources.
type DefaultNetworkAclState struct {
	// The Network ACL ID to manage. This
	// attribute is exported from `ec2.Vpc`, or manually found via the AWS Console.
	DefaultNetworkAclId interface{}
	// Specifies an egress rule. Parameters defined below.
	Egress interface{}
	// Specifies an ingress rule. Parameters defined below.
	Ingress interface{}
	// The ID of the AWS account that owns the Default Network ACL
	OwnerId interface{}
	// A list of Subnet IDs to apply the ACL to. See the
	// notes below on managing Subnets in the Default Network ACL
	SubnetIds interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The ID of the associated VPC
	VpcId interface{}
}

// The set of arguments for constructing a DefaultNetworkAcl resource.
type DefaultNetworkAclArgs struct {
	// The Network ACL ID to manage. This
	// attribute is exported from `ec2.Vpc`, or manually found via the AWS Console.
	DefaultNetworkAclId interface{}
	// Specifies an egress rule. Parameters defined below.
	Egress interface{}
	// Specifies an ingress rule. Parameters defined below.
	Ingress interface{}
	// A list of Subnet IDs to apply the ACL to. See the
	// notes below on managing Subnets in the Default Network ACL
	SubnetIds interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
