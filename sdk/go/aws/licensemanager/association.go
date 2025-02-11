// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package licensemanager

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a License Manager association.
// 
// > **Note:** License configurations can also be associated with launch templates by specifying the `licenseSpecifications` block for an `ec2.LaunchTemplate`.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/licensemanager_association.html.markdown.
type Association struct {
	s *pulumi.ResourceState
}

// NewAssociation registers a new resource with the given unique name, arguments, and options.
func NewAssociation(ctx *pulumi.Context,
	name string, args *AssociationArgs, opts ...pulumi.ResourceOpt) (*Association, error) {
	if args == nil || args.LicenseConfigurationArn == nil {
		return nil, errors.New("missing required argument 'LicenseConfigurationArn'")
	}
	if args == nil || args.ResourceArn == nil {
		return nil, errors.New("missing required argument 'ResourceArn'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["licenseConfigurationArn"] = nil
		inputs["resourceArn"] = nil
	} else {
		inputs["licenseConfigurationArn"] = args.LicenseConfigurationArn
		inputs["resourceArn"] = args.ResourceArn
	}
	s, err := ctx.RegisterResource("aws:licensemanager/association:Association", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Association{s: s}, nil
}

// GetAssociation gets an existing Association resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssociation(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AssociationState, opts ...pulumi.ResourceOpt) (*Association, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["licenseConfigurationArn"] = state.LicenseConfigurationArn
		inputs["resourceArn"] = state.ResourceArn
	}
	s, err := ctx.ReadResource("aws:licensemanager/association:Association", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Association{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Association) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Association) ID() pulumi.IDOutput {
	return r.s.ID()
}

// ARN of the license configuration.
func (r *Association) LicenseConfigurationArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["licenseConfigurationArn"])
}

// ARN of the resource associated with the license configuration.
func (r *Association) ResourceArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceArn"])
}

// Input properties used for looking up and filtering Association resources.
type AssociationState struct {
	// ARN of the license configuration.
	LicenseConfigurationArn interface{}
	// ARN of the resource associated with the license configuration.
	ResourceArn interface{}
}

// The set of arguments for constructing a Association resource.
type AssociationArgs struct {
	// ARN of the license configuration.
	LicenseConfigurationArn interface{}
	// ARN of the resource associated with the license configuration.
	ResourceArn interface{}
}
