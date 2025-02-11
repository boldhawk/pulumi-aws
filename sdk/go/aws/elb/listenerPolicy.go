// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elb

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Attaches a load balancer policy to an ELB Listener.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/load_balancer_listener_policy.html.markdown.
type ListenerPolicy struct {
	s *pulumi.ResourceState
}

// NewListenerPolicy registers a new resource with the given unique name, arguments, and options.
func NewListenerPolicy(ctx *pulumi.Context,
	name string, args *ListenerPolicyArgs, opts ...pulumi.ResourceOpt) (*ListenerPolicy, error) {
	if args == nil || args.LoadBalancerName == nil {
		return nil, errors.New("missing required argument 'LoadBalancerName'")
	}
	if args == nil || args.LoadBalancerPort == nil {
		return nil, errors.New("missing required argument 'LoadBalancerPort'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["loadBalancerName"] = nil
		inputs["loadBalancerPort"] = nil
		inputs["policyNames"] = nil
	} else {
		inputs["loadBalancerName"] = args.LoadBalancerName
		inputs["loadBalancerPort"] = args.LoadBalancerPort
		inputs["policyNames"] = args.PolicyNames
	}
	s, err := ctx.RegisterResource("aws:elb/listenerPolicy:ListenerPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ListenerPolicy{s: s}, nil
}

// GetListenerPolicy gets an existing ListenerPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListenerPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ListenerPolicyState, opts ...pulumi.ResourceOpt) (*ListenerPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["loadBalancerName"] = state.LoadBalancerName
		inputs["loadBalancerPort"] = state.LoadBalancerPort
		inputs["policyNames"] = state.PolicyNames
	}
	s, err := ctx.ReadResource("aws:elb/listenerPolicy:ListenerPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ListenerPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ListenerPolicy) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ListenerPolicy) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The load balancer to attach the policy to.
func (r *ListenerPolicy) LoadBalancerName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["loadBalancerName"])
}

// The load balancer listener port to apply the policy to.
func (r *ListenerPolicy) LoadBalancerPort() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["loadBalancerPort"])
}

// List of Policy Names to apply to the backend server.
func (r *ListenerPolicy) PolicyNames() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["policyNames"])
}

// Input properties used for looking up and filtering ListenerPolicy resources.
type ListenerPolicyState struct {
	// The load balancer to attach the policy to.
	LoadBalancerName interface{}
	// The load balancer listener port to apply the policy to.
	LoadBalancerPort interface{}
	// List of Policy Names to apply to the backend server.
	PolicyNames interface{}
}

// The set of arguments for constructing a ListenerPolicy resource.
type ListenerPolicyArgs struct {
	// The load balancer to attach the policy to.
	LoadBalancerName interface{}
	// The load balancer listener port to apply the policy to.
	LoadBalancerPort interface{}
	// List of Policy Names to apply to the backend server.
	PolicyNames interface{}
}
