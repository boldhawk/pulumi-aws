// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package applicationloadbalancing

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Load Balancer Listener Rule resource.
// 
// > **Note:** `alb.ListenerRule` is known as `lb.ListenerRule`. The functionality is identical.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/alb_listener_rule_legacy.html.markdown.
type ListenerRule struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// An Action block. Action blocks are documented below.
	Actions pulumi.ArrayOutput `pulumi:"actions"`

	// The ARN of the rule (matches `id`)
	Arn pulumi.StringOutput `pulumi:"arn"`

	// A Condition block. Condition blocks are documented below.
	Conditions pulumi.ArrayOutput `pulumi:"conditions"`

	// The ARN of the listener to which to attach the rule.
	ListenerArn pulumi.StringOutput `pulumi:"listenerArn"`

	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority pulumi.IntOutput `pulumi:"priority"`
}

// NewListenerRule registers a new resource with the given unique name, arguments, and options.
func NewListenerRule(ctx *pulumi.Context,
	name string, args *ListenerRuleArgs, opts ...pulumi.ResourceOpt) (*ListenerRule, error) {
	if args == nil || args.Actions == nil {
		return nil, errors.New("missing required argument 'Actions'")
	}
	if args == nil || args.Conditions == nil {
		return nil, errors.New("missing required argument 'Conditions'")
	}
	if args == nil || args.ListenerArn == nil {
		return nil, errors.New("missing required argument 'ListenerArn'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["actions"] = args.Actions
		inputs["conditions"] = args.Conditions
		inputs["listenerArn"] = args.ListenerArn
		inputs["priority"] = args.Priority
	}
	var resource ListenerRule
	err := ctx.RegisterResource("aws:applicationloadbalancing/listenerRule:ListenerRule", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListenerRule gets an existing ListenerRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListenerRule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ListenerRuleState, opts ...pulumi.ResourceOpt) (*ListenerRule, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["actions"] = state.Actions
		inputs["arn"] = state.Arn
		inputs["conditions"] = state.Conditions
		inputs["listenerArn"] = state.ListenerArn
		inputs["priority"] = state.Priority
	}
	var resource ListenerRule
	err := ctx.ReadResource("aws:applicationloadbalancing/listenerRule:ListenerRule", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *ListenerRule) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *ListenerRule) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering ListenerRule resources.
type ListenerRuleState struct {
	// An Action block. Action blocks are documented below.
	Actions pulumi.ArrayInput `pulumi:"actions"`
	// The ARN of the rule (matches `id`)
	Arn pulumi.StringInput `pulumi:"arn"`
	// A Condition block. Condition blocks are documented below.
	Conditions pulumi.ArrayInput `pulumi:"conditions"`
	// The ARN of the listener to which to attach the rule.
	ListenerArn pulumi.StringInput `pulumi:"listenerArn"`
	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority pulumi.IntInput `pulumi:"priority"`
}

// The set of arguments for constructing a ListenerRule resource.
type ListenerRuleArgs struct {
	// An Action block. Action blocks are documented below.
	Actions pulumi.ArrayInput `pulumi:"actions"`
	// A Condition block. Condition blocks are documented below.
	Conditions pulumi.ArrayInput `pulumi:"conditions"`
	// The ARN of the listener to which to attach the rule.
	ListenerArn pulumi.StringInput `pulumi:"listenerArn"`
	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority pulumi.IntInput `pulumi:"priority"`
}
