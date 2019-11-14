// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sfn

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Step Function State Machine resource
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/sfn_state_machine.html.markdown.
type StateMachine struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The date the state machine was created.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`

	// The Amazon States Language definition of the state machine.
	Definition pulumi.StringOutput `pulumi:"definition"`

	// The name of the state machine.
	Name pulumi.StringOutput `pulumi:"name"`

	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`

	// The current status of the state machine. Either "ACTIVE" or "DELETING".
	Status pulumi.StringOutput `pulumi:"status"`

	// Key-value mapping of resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewStateMachine registers a new resource with the given unique name, arguments, and options.
func NewStateMachine(ctx *pulumi.Context,
	name string, args *StateMachineArgs, opts ...pulumi.ResourceOpt) (*StateMachine, error) {
	if args == nil || args.Definition == nil {
		return nil, errors.New("missing required argument 'Definition'")
	}
	if args == nil || args.RoleArn == nil {
		return nil, errors.New("missing required argument 'RoleArn'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["definition"] = args.Definition
		inputs["name"] = args.Name
		inputs["roleArn"] = args.RoleArn
		inputs["tags"] = args.Tags
	}
	var resource StateMachine
	err := ctx.RegisterResource("aws:sfn/stateMachine:StateMachine", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStateMachine gets an existing StateMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStateMachine(ctx *pulumi.Context,
	name string, id pulumi.ID, state *StateMachineState, opts ...pulumi.ResourceOpt) (*StateMachine, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["creationDate"] = state.CreationDate
		inputs["definition"] = state.Definition
		inputs["name"] = state.Name
		inputs["roleArn"] = state.RoleArn
		inputs["status"] = state.Status
		inputs["tags"] = state.Tags
	}
	var resource StateMachine
	err := ctx.ReadResource("aws:sfn/stateMachine:StateMachine", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *StateMachine) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *StateMachine) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering StateMachine resources.
type StateMachineState struct {
	// The date the state machine was created.
	CreationDate pulumi.StringInput `pulumi:"creationDate"`
	// The Amazon States Language definition of the state machine.
	Definition pulumi.StringInput `pulumi:"definition"`
	// The name of the state machine.
	Name pulumi.StringInput `pulumi:"name"`
	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn pulumi.StringInput `pulumi:"roleArn"`
	// The current status of the state machine. Either "ACTIVE" or "DELETING".
	Status pulumi.StringInput `pulumi:"status"`
	// Key-value mapping of resource tags
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a StateMachine resource.
type StateMachineArgs struct {
	// The Amazon States Language definition of the state machine.
	Definition pulumi.StringInput `pulumi:"definition"`
	// The name of the state machine.
	Name pulumi.StringInput `pulumi:"name"`
	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn pulumi.StringInput `pulumi:"roleArn"`
	// Key-value mapping of resource tags
	Tags pulumi.MapInput `pulumi:"tags"`
}
