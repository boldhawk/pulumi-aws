// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package worklink

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Fleet struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN of the created WorkLink Fleet.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The ARN of the Amazon Kinesis data stream that receives the audit events.
	AuditStreamArn pulumi.StringOutput `pulumi:"auditStreamArn"`

	// The identifier used by users to sign in to the Amazon WorkLink app.
	CompanyCode pulumi.StringOutput `pulumi:"companyCode"`

	// The time that the fleet was created.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`

	// The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
	DeviceCaCertificate pulumi.StringOutput `pulumi:"deviceCaCertificate"`

	// The name of the fleet.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`

	// Provide this to allow manage the identity provider configuration for the fleet. Fields documented below.
	IdentityProvider pulumi.AnyOutput `pulumi:"identityProvider"`

	// The time that the fleet was last updated.
	LastUpdatedTime pulumi.StringOutput `pulumi:"lastUpdatedTime"`

	// A region-unique name for the AMI.
	Name pulumi.StringOutput `pulumi:"name"`

	// Provide this to allow manage the company network configuration for the fleet. Fields documented below.
	Network pulumi.AnyOutput `pulumi:"network"`

	// The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
	OptimizeForEndUserLocation pulumi.BoolOutput `pulumi:"optimizeForEndUserLocation"`
}

// NewFleet registers a new resource with the given unique name, arguments, and options.
func NewFleet(ctx *pulumi.Context,
	name string, args *FleetArgs, opts ...pulumi.ResourceOpt) (*Fleet, error) {
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["auditStreamArn"] = args.AuditStreamArn
		inputs["deviceCaCertificate"] = args.DeviceCaCertificate
		inputs["displayName"] = args.DisplayName
		inputs["identityProvider"] = args.IdentityProvider
		inputs["name"] = args.Name
		inputs["network"] = args.Network
		inputs["optimizeForEndUserLocation"] = args.OptimizeForEndUserLocation
	}
	var resource Fleet
	err := ctx.RegisterResource("aws:worklink/fleet:Fleet", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFleet gets an existing Fleet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFleet(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FleetState, opts ...pulumi.ResourceOpt) (*Fleet, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["auditStreamArn"] = state.AuditStreamArn
		inputs["companyCode"] = state.CompanyCode
		inputs["createdTime"] = state.CreatedTime
		inputs["deviceCaCertificate"] = state.DeviceCaCertificate
		inputs["displayName"] = state.DisplayName
		inputs["identityProvider"] = state.IdentityProvider
		inputs["lastUpdatedTime"] = state.LastUpdatedTime
		inputs["name"] = state.Name
		inputs["network"] = state.Network
		inputs["optimizeForEndUserLocation"] = state.OptimizeForEndUserLocation
	}
	var resource Fleet
	err := ctx.ReadResource("aws:worklink/fleet:Fleet", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *Fleet) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *Fleet) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering Fleet resources.
type FleetState struct {
	// The ARN of the created WorkLink Fleet.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The ARN of the Amazon Kinesis data stream that receives the audit events.
	AuditStreamArn pulumi.StringInput `pulumi:"auditStreamArn"`
	// The identifier used by users to sign in to the Amazon WorkLink app.
	CompanyCode pulumi.StringInput `pulumi:"companyCode"`
	// The time that the fleet was created.
	CreatedTime pulumi.StringInput `pulumi:"createdTime"`
	// The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
	DeviceCaCertificate pulumi.StringInput `pulumi:"deviceCaCertificate"`
	// The name of the fleet.
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	// Provide this to allow manage the identity provider configuration for the fleet. Fields documented below.
	IdentityProvider pulumi.AnyInput `pulumi:"identityProvider"`
	// The time that the fleet was last updated.
	LastUpdatedTime pulumi.StringInput `pulumi:"lastUpdatedTime"`
	// A region-unique name for the AMI.
	Name pulumi.StringInput `pulumi:"name"`
	// Provide this to allow manage the company network configuration for the fleet. Fields documented below.
	Network pulumi.AnyInput `pulumi:"network"`
	// The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
	OptimizeForEndUserLocation pulumi.BoolInput `pulumi:"optimizeForEndUserLocation"`
}

// The set of arguments for constructing a Fleet resource.
type FleetArgs struct {
	// The ARN of the Amazon Kinesis data stream that receives the audit events.
	AuditStreamArn pulumi.StringInput `pulumi:"auditStreamArn"`
	// The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
	DeviceCaCertificate pulumi.StringInput `pulumi:"deviceCaCertificate"`
	// The name of the fleet.
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	// Provide this to allow manage the identity provider configuration for the fleet. Fields documented below.
	IdentityProvider pulumi.AnyInput `pulumi:"identityProvider"`
	// A region-unique name for the AMI.
	Name pulumi.StringInput `pulumi:"name"`
	// Provide this to allow manage the company network configuration for the fleet. Fields documented below.
	Network pulumi.AnyInput `pulumi:"network"`
	// The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
	OptimizeForEndUserLocation pulumi.BoolInput `pulumi:"optimizeForEndUserLocation"`
}
