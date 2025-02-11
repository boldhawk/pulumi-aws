// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancing

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a load balancer SSL negotiation policy, which allows an ELB to control the ciphers and protocols that are supported during SSL negotiations between a client and a load balancer.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_ssl_negotiation_policy_legacy.html.markdown.
type SslNegotiationPolicy struct {
	s *pulumi.ResourceState
}

// NewSslNegotiationPolicy registers a new resource with the given unique name, arguments, and options.
func NewSslNegotiationPolicy(ctx *pulumi.Context,
	name string, args *SslNegotiationPolicyArgs, opts ...pulumi.ResourceOpt) (*SslNegotiationPolicy, error) {
	if args == nil || args.LbPort == nil {
		return nil, errors.New("missing required argument 'LbPort'")
	}
	if args == nil || args.LoadBalancer == nil {
		return nil, errors.New("missing required argument 'LoadBalancer'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["attributes"] = nil
		inputs["lbPort"] = nil
		inputs["loadBalancer"] = nil
		inputs["name"] = nil
	} else {
		inputs["attributes"] = args.Attributes
		inputs["lbPort"] = args.LbPort
		inputs["loadBalancer"] = args.LoadBalancer
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("aws:elasticloadbalancing/sslNegotiationPolicy:SslNegotiationPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SslNegotiationPolicy{s: s}, nil
}

// GetSslNegotiationPolicy gets an existing SslNegotiationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSslNegotiationPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SslNegotiationPolicyState, opts ...pulumi.ResourceOpt) (*SslNegotiationPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["attributes"] = state.Attributes
		inputs["lbPort"] = state.LbPort
		inputs["loadBalancer"] = state.LoadBalancer
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("aws:elasticloadbalancing/sslNegotiationPolicy:SslNegotiationPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SslNegotiationPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SslNegotiationPolicy) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SslNegotiationPolicy) ID() pulumi.IDOutput {
	return r.s.ID()
}

// An SSL Negotiation policy attribute. Each has two properties:
func (r *SslNegotiationPolicy) Attributes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["attributes"])
}

// The load balancer port to which the policy
// should be applied. This must be an active listener on the load
// balancer.
func (r *SslNegotiationPolicy) LbPort() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["lbPort"])
}

// The load balancer to which the policy
// should be attached.
func (r *SslNegotiationPolicy) LoadBalancer() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["loadBalancer"])
}

// The name of the attribute
func (r *SslNegotiationPolicy) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering SslNegotiationPolicy resources.
type SslNegotiationPolicyState struct {
	// An SSL Negotiation policy attribute. Each has two properties:
	Attributes interface{}
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort interface{}
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer interface{}
	// The name of the attribute
	Name interface{}
}

// The set of arguments for constructing a SslNegotiationPolicy resource.
type SslNegotiationPolicyArgs struct {
	// An SSL Negotiation policy attribute. Each has two properties:
	Attributes interface{}
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort interface{}
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer interface{}
	// The name of the attribute
	Name interface{}
}
