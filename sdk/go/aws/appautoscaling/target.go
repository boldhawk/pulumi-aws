// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appautoscaling

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Application AutoScaling ScalableTarget resource. To manage policies which get attached to the target, see the [`appautoscaling.Policy` resource](https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appautoscaling_target.html.markdown.
type Target struct {
	s *pulumi.ResourceState
}

// NewTarget registers a new resource with the given unique name, arguments, and options.
func NewTarget(ctx *pulumi.Context,
	name string, args *TargetArgs, opts ...pulumi.ResourceOpt) (*Target, error) {
	if args == nil || args.MaxCapacity == nil {
		return nil, errors.New("missing required argument 'MaxCapacity'")
	}
	if args == nil || args.MinCapacity == nil {
		return nil, errors.New("missing required argument 'MinCapacity'")
	}
	if args == nil || args.ResourceId == nil {
		return nil, errors.New("missing required argument 'ResourceId'")
	}
	if args == nil || args.ScalableDimension == nil {
		return nil, errors.New("missing required argument 'ScalableDimension'")
	}
	if args == nil || args.ServiceNamespace == nil {
		return nil, errors.New("missing required argument 'ServiceNamespace'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["maxCapacity"] = nil
		inputs["minCapacity"] = nil
		inputs["resourceId"] = nil
		inputs["roleArn"] = nil
		inputs["scalableDimension"] = nil
		inputs["serviceNamespace"] = nil
	} else {
		inputs["maxCapacity"] = args.MaxCapacity
		inputs["minCapacity"] = args.MinCapacity
		inputs["resourceId"] = args.ResourceId
		inputs["roleArn"] = args.RoleArn
		inputs["scalableDimension"] = args.ScalableDimension
		inputs["serviceNamespace"] = args.ServiceNamespace
	}
	s, err := ctx.RegisterResource("aws:appautoscaling/target:Target", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Target{s: s}, nil
}

// GetTarget gets an existing Target resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTarget(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TargetState, opts ...pulumi.ResourceOpt) (*Target, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["maxCapacity"] = state.MaxCapacity
		inputs["minCapacity"] = state.MinCapacity
		inputs["resourceId"] = state.ResourceId
		inputs["roleArn"] = state.RoleArn
		inputs["scalableDimension"] = state.ScalableDimension
		inputs["serviceNamespace"] = state.ServiceNamespace
	}
	s, err := ctx.ReadResource("aws:appautoscaling/target:Target", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Target{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Target) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Target) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The max capacity of the scalable target.
func (r *Target) MaxCapacity() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["maxCapacity"])
}

// The min capacity of the scalable target.
func (r *Target) MinCapacity() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["minCapacity"])
}

// The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
func (r *Target) ResourceId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceId"])
}

// The ARN of the IAM role that allows Application
// AutoScaling to modify your scalable target on your behalf.
func (r *Target) RoleArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["roleArn"])
}

// The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
func (r *Target) ScalableDimension() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["scalableDimension"])
}

// The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
func (r *Target) ServiceNamespace() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["serviceNamespace"])
}

// Input properties used for looking up and filtering Target resources.
type TargetState struct {
	// The max capacity of the scalable target.
	MaxCapacity interface{}
	// The min capacity of the scalable target.
	MinCapacity interface{}
	// The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ResourceId interface{}
	// The ARN of the IAM role that allows Application
	// AutoScaling to modify your scalable target on your behalf.
	RoleArn interface{}
	// The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ScalableDimension interface{}
	// The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ServiceNamespace interface{}
}

// The set of arguments for constructing a Target resource.
type TargetArgs struct {
	// The max capacity of the scalable target.
	MaxCapacity interface{}
	// The min capacity of the scalable target.
	MinCapacity interface{}
	// The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ResourceId interface{}
	// The ARN of the IAM role that allows Application
	// AutoScaling to modify your scalable target on your behalf.
	RoleArn interface{}
	// The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ScalableDimension interface{}
	// The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ServiceNamespace interface{}
}
