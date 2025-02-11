// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Represents a successful verification of an SES domain identity.
// 
// Most commonly, this resource is used together with `route53.Record` and
// `ses.DomainIdentity` to request an SES domain identity,
// deploy the required DNS verification records, and wait for verification to complete.
// 
// > **WARNING:** This resource implements a part of the verification workflow. It does not represent a real-world entity in AWS, therefore changing or deleting this resource on its own has no immediate effect.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_domain_identity_verification.html.markdown.
type DomainIdentityVerification struct {
	s *pulumi.ResourceState
}

// NewDomainIdentityVerification registers a new resource with the given unique name, arguments, and options.
func NewDomainIdentityVerification(ctx *pulumi.Context,
	name string, args *DomainIdentityVerificationArgs, opts ...pulumi.ResourceOpt) (*DomainIdentityVerification, error) {
	if args == nil || args.Domain == nil {
		return nil, errors.New("missing required argument 'Domain'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["domain"] = nil
	} else {
		inputs["domain"] = args.Domain
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:ses/domainIdentityVerification:DomainIdentityVerification", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DomainIdentityVerification{s: s}, nil
}

// GetDomainIdentityVerification gets an existing DomainIdentityVerification resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainIdentityVerification(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DomainIdentityVerificationState, opts ...pulumi.ResourceOpt) (*DomainIdentityVerification, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["domain"] = state.Domain
	}
	s, err := ctx.ReadResource("aws:ses/domainIdentityVerification:DomainIdentityVerification", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DomainIdentityVerification{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DomainIdentityVerification) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DomainIdentityVerification) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the domain identity.
func (r *DomainIdentityVerification) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// The domain name of the SES domain identity to verify.
func (r *DomainIdentityVerification) Domain() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["domain"])
}

// Input properties used for looking up and filtering DomainIdentityVerification resources.
type DomainIdentityVerificationState struct {
	// The ARN of the domain identity.
	Arn interface{}
	// The domain name of the SES domain identity to verify.
	Domain interface{}
}

// The set of arguments for constructing a DomainIdentityVerification resource.
type DomainIdentityVerificationArgs struct {
	// The domain name of the SES domain identity to verify.
	Domain interface{}
}
