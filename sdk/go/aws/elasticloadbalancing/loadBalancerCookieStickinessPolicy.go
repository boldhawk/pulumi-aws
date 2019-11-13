// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancing

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a load balancer cookie stickiness policy, which allows an ELB to control the sticky session lifetime of the browser.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_cookie_stickiness_policy_legacy.html.markdown.
type LoadBalancerCookieStickinessPolicy struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The time period after which
	// the session cookie should be considered stale, expressed in seconds.
	CookieExpirationPeriod pulumi.IntOutput `pulumi:"cookieExpirationPeriod"`

	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntOutput `pulumi:"lbPort"`

	// The load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringOutput `pulumi:"loadBalancer"`

	// The name of the stickiness policy.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewLoadBalancerCookieStickinessPolicy registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancerCookieStickinessPolicy(ctx *pulumi.Context,
	name string, args *LoadBalancerCookieStickinessPolicyArgs, opts ...pulumi.ResourceOpt) (*LoadBalancerCookieStickinessPolicy, error) {
	if args == nil || args.LbPort == nil {
		return nil, errors.New("missing required argument 'LbPort'")
	}
	if args == nil || args.LoadBalancer == nil {
		return nil, errors.New("missing required argument 'LoadBalancer'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["cookieExpirationPeriod"] = args.CookieExpirationPeriod
		inputs["lbPort"] = args.LbPort
		inputs["loadBalancer"] = args.LoadBalancer
		inputs["name"] = args.Name
	}
	var resource LoadBalancerCookieStickinessPolicy
	err := ctx.RegisterResource("aws:elasticloadbalancing/loadBalancerCookieStickinessPolicy:LoadBalancerCookieStickinessPolicy", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancerCookieStickinessPolicy gets an existing LoadBalancerCookieStickinessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancerCookieStickinessPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LoadBalancerCookieStickinessPolicyState, opts ...pulumi.ResourceOpt) (*LoadBalancerCookieStickinessPolicy, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["cookieExpirationPeriod"] = state.CookieExpirationPeriod
		inputs["lbPort"] = state.LbPort
		inputs["loadBalancer"] = state.LoadBalancer
		inputs["name"] = state.Name
	}
	var resource LoadBalancerCookieStickinessPolicy
	err := ctx.ReadResource("aws:elasticloadbalancing/loadBalancerCookieStickinessPolicy:LoadBalancerCookieStickinessPolicy", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *LoadBalancerCookieStickinessPolicy) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *LoadBalancerCookieStickinessPolicy) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering LoadBalancerCookieStickinessPolicy resources.
type LoadBalancerCookieStickinessPolicyState struct {
	// The time period after which
	// the session cookie should be considered stale, expressed in seconds.
	CookieExpirationPeriod pulumi.IntInput `pulumi:"cookieExpirationPeriod"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntInput `pulumi:"lbPort"`
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringInput `pulumi:"loadBalancer"`
	// The name of the stickiness policy.
	Name pulumi.StringInput `pulumi:"name"`
}

// The set of arguments for constructing a LoadBalancerCookieStickinessPolicy resource.
type LoadBalancerCookieStickinessPolicyArgs struct {
	// The time period after which
	// the session cookie should be considered stale, expressed in seconds.
	CookieExpirationPeriod pulumi.IntInput `pulumi:"cookieExpirationPeriod"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntInput `pulumi:"lbPort"`
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringInput `pulumi:"loadBalancer"`
	// The name of the stickiness policy.
	Name pulumi.StringInput `pulumi:"name"`
}
