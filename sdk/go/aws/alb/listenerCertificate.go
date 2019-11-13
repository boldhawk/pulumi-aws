// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Load Balancer Listener Certificate resource.
// 
// This resource is for additional certificates and does not replace the default certificate on the listener.
// 
// > **Note:** `alb.ListenerCertificate` is known as `lb.ListenerCertificate`. The functionality is identical.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/alb_listener_certificate.html.markdown.
type ListenerCertificate struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN of the certificate to attach to the listener.
	CertificateArn pulumi.StringOutput `pulumi:"certificateArn"`

	// The ARN of the listener to which to attach the certificate.
	ListenerArn pulumi.StringOutput `pulumi:"listenerArn"`
}

// NewListenerCertificate registers a new resource with the given unique name, arguments, and options.
func NewListenerCertificate(ctx *pulumi.Context,
	name string, args *ListenerCertificateArgs, opts ...pulumi.ResourceOpt) (*ListenerCertificate, error) {
	if args == nil || args.CertificateArn == nil {
		return nil, errors.New("missing required argument 'CertificateArn'")
	}
	if args == nil || args.ListenerArn == nil {
		return nil, errors.New("missing required argument 'ListenerArn'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["certificateArn"] = args.CertificateArn
		inputs["listenerArn"] = args.ListenerArn
	}
	var resource ListenerCertificate
	err := ctx.RegisterResource("aws:alb/listenerCertificate:ListenerCertificate", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListenerCertificate gets an existing ListenerCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListenerCertificate(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ListenerCertificateState, opts ...pulumi.ResourceOpt) (*ListenerCertificate, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["certificateArn"] = state.CertificateArn
		inputs["listenerArn"] = state.ListenerArn
	}
	var resource ListenerCertificate
	err := ctx.ReadResource("aws:alb/listenerCertificate:ListenerCertificate", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *ListenerCertificate) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *ListenerCertificate) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering ListenerCertificate resources.
type ListenerCertificateState struct {
	// The ARN of the certificate to attach to the listener.
	CertificateArn pulumi.StringInput `pulumi:"certificateArn"`
	// The ARN of the listener to which to attach the certificate.
	ListenerArn pulumi.StringInput `pulumi:"listenerArn"`
}

// The set of arguments for constructing a ListenerCertificate resource.
type ListenerCertificateArgs struct {
	// The ARN of the certificate to attach to the listener.
	CertificateArn pulumi.StringInput `pulumi:"certificateArn"`
	// The ARN of the listener to which to attach the certificate.
	ListenerArn pulumi.StringInput `pulumi:"listenerArn"`
}
