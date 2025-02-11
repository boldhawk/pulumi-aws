// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancing

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an application cookie stickiness policy, which allows an ELB to wed its sticky cookie's expiration to a cookie generated by your application.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/app_cookie_stickiness_policy_legacy.html.markdown.
type AppCookieStickinessPolicy struct {
	s *pulumi.ResourceState
}

// NewAppCookieStickinessPolicy registers a new resource with the given unique name, arguments, and options.
func NewAppCookieStickinessPolicy(ctx *pulumi.Context,
	name string, args *AppCookieStickinessPolicyArgs, opts ...pulumi.ResourceOpt) (*AppCookieStickinessPolicy, error) {
	if args == nil || args.CookieName == nil {
		return nil, errors.New("missing required argument 'CookieName'")
	}
	if args == nil || args.LbPort == nil {
		return nil, errors.New("missing required argument 'LbPort'")
	}
	if args == nil || args.LoadBalancer == nil {
		return nil, errors.New("missing required argument 'LoadBalancer'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cookieName"] = nil
		inputs["lbPort"] = nil
		inputs["loadBalancer"] = nil
		inputs["name"] = nil
	} else {
		inputs["cookieName"] = args.CookieName
		inputs["lbPort"] = args.LbPort
		inputs["loadBalancer"] = args.LoadBalancer
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("aws:elasticloadbalancing/appCookieStickinessPolicy:AppCookieStickinessPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AppCookieStickinessPolicy{s: s}, nil
}

// GetAppCookieStickinessPolicy gets an existing AppCookieStickinessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppCookieStickinessPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AppCookieStickinessPolicyState, opts ...pulumi.ResourceOpt) (*AppCookieStickinessPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["cookieName"] = state.CookieName
		inputs["lbPort"] = state.LbPort
		inputs["loadBalancer"] = state.LoadBalancer
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("aws:elasticloadbalancing/appCookieStickinessPolicy:AppCookieStickinessPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AppCookieStickinessPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AppCookieStickinessPolicy) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AppCookieStickinessPolicy) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The application cookie whose lifetime the ELB's cookie should follow.
func (r *AppCookieStickinessPolicy) CookieName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["cookieName"])
}

// The load balancer port to which the policy
// should be applied. This must be an active listener on the load
// balancer.
func (r *AppCookieStickinessPolicy) LbPort() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["lbPort"])
}

// The name of load balancer to which the policy
// should be attached.
func (r *AppCookieStickinessPolicy) LoadBalancer() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["loadBalancer"])
}

// The name of the stickiness policy.
func (r *AppCookieStickinessPolicy) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering AppCookieStickinessPolicy resources.
type AppCookieStickinessPolicyState struct {
	// The application cookie whose lifetime the ELB's cookie should follow.
	CookieName interface{}
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort interface{}
	// The name of load balancer to which the policy
	// should be attached.
	LoadBalancer interface{}
	// The name of the stickiness policy.
	Name interface{}
}

// The set of arguments for constructing a AppCookieStickinessPolicy resource.
type AppCookieStickinessPolicyArgs struct {
	// The application cookie whose lifetime the ELB's cookie should follow.
	CookieName interface{}
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort interface{}
	// The name of load balancer to which the policy
	// should be attached.
	LoadBalancer interface{}
	// The name of the stickiness policy.
	Name interface{}
}
