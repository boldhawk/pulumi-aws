// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appautoscaling

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Application AutoScaling ScheduledAction resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appautoscaling_scheduled_action.html.markdown.
type ScheduledAction struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The Amazon Resource Name (ARN) of the scheduled action.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The date and time for the scheduled action to end. Specify the following format: 2006-01-02T15:04:05Z
	EndTime pulumi.StringOutput `pulumi:"endTime"`

	// The name of the scheduled action.
	Name pulumi.StringOutput `pulumi:"name"`

	// The identifier of the resource associated with the scheduled action. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ResourceId)
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`

	// The scalable dimension. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ScalableDimension) Example: ecs:service:DesiredCount
	ScalableDimension pulumi.StringOutput `pulumi:"scalableDimension"`

	// The new minimum and maximum capacity. You can set both values or just one. See below
	ScalableTargetAction pulumi.AnyOutput `pulumi:"scalableTargetAction"`

	// The schedule for this action. The following formats are supported: At expressions - at(yyyy-mm-ddThh:mm:ss), Rate expressions - rate(valueunit), Cron expressions - cron(fields). In UTC. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-Schedule)
	Schedule pulumi.StringOutput `pulumi:"schedule"`

	// The namespace of the AWS service. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ServiceNamespace) Example: ecs
	ServiceNamespace pulumi.StringOutput `pulumi:"serviceNamespace"`

	// The date and time for the scheduled action to start. Specify the following format: 2006-01-02T15:04:05Z
	StartTime pulumi.StringOutput `pulumi:"startTime"`
}

// NewScheduledAction registers a new resource with the given unique name, arguments, and options.
func NewScheduledAction(ctx *pulumi.Context,
	name string, args *ScheduledActionArgs, opts ...pulumi.ResourceOpt) (*ScheduledAction, error) {
	if args == nil || args.ResourceId == nil {
		return nil, errors.New("missing required argument 'ResourceId'")
	}
	if args == nil || args.ServiceNamespace == nil {
		return nil, errors.New("missing required argument 'ServiceNamespace'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["endTime"] = args.EndTime
		inputs["name"] = args.Name
		inputs["resourceId"] = args.ResourceId
		inputs["scalableDimension"] = args.ScalableDimension
		inputs["scalableTargetAction"] = args.ScalableTargetAction
		inputs["schedule"] = args.Schedule
		inputs["serviceNamespace"] = args.ServiceNamespace
		inputs["startTime"] = args.StartTime
	}
	var resource ScheduledAction
	err := ctx.RegisterResource("aws:appautoscaling/scheduledAction:ScheduledAction", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduledAction gets an existing ScheduledAction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduledAction(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ScheduledActionState, opts ...pulumi.ResourceOpt) (*ScheduledAction, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["endTime"] = state.EndTime
		inputs["name"] = state.Name
		inputs["resourceId"] = state.ResourceId
		inputs["scalableDimension"] = state.ScalableDimension
		inputs["scalableTargetAction"] = state.ScalableTargetAction
		inputs["schedule"] = state.Schedule
		inputs["serviceNamespace"] = state.ServiceNamespace
		inputs["startTime"] = state.StartTime
	}
	var resource ScheduledAction
	err := ctx.ReadResource("aws:appautoscaling/scheduledAction:ScheduledAction", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *ScheduledAction) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *ScheduledAction) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering ScheduledAction resources.
type ScheduledActionState struct {
	// The Amazon Resource Name (ARN) of the scheduled action.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The date and time for the scheduled action to end. Specify the following format: 2006-01-02T15:04:05Z
	EndTime pulumi.StringInput `pulumi:"endTime"`
	// The name of the scheduled action.
	Name pulumi.StringInput `pulumi:"name"`
	// The identifier of the resource associated with the scheduled action. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ResourceId)
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// The scalable dimension. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ScalableDimension) Example: ecs:service:DesiredCount
	ScalableDimension pulumi.StringInput `pulumi:"scalableDimension"`
	// The new minimum and maximum capacity. You can set both values or just one. See below
	ScalableTargetAction pulumi.AnyInput `pulumi:"scalableTargetAction"`
	// The schedule for this action. The following formats are supported: At expressions - at(yyyy-mm-ddThh:mm:ss), Rate expressions - rate(valueunit), Cron expressions - cron(fields). In UTC. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-Schedule)
	Schedule pulumi.StringInput `pulumi:"schedule"`
	// The namespace of the AWS service. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ServiceNamespace) Example: ecs
	ServiceNamespace pulumi.StringInput `pulumi:"serviceNamespace"`
	// The date and time for the scheduled action to start. Specify the following format: 2006-01-02T15:04:05Z
	StartTime pulumi.StringInput `pulumi:"startTime"`
}

// The set of arguments for constructing a ScheduledAction resource.
type ScheduledActionArgs struct {
	// The date and time for the scheduled action to end. Specify the following format: 2006-01-02T15:04:05Z
	EndTime pulumi.StringInput `pulumi:"endTime"`
	// The name of the scheduled action.
	Name pulumi.StringInput `pulumi:"name"`
	// The identifier of the resource associated with the scheduled action. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ResourceId)
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// The scalable dimension. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ScalableDimension) Example: ecs:service:DesiredCount
	ScalableDimension pulumi.StringInput `pulumi:"scalableDimension"`
	// The new minimum and maximum capacity. You can set both values or just one. See below
	ScalableTargetAction pulumi.AnyInput `pulumi:"scalableTargetAction"`
	// The schedule for this action. The following formats are supported: At expressions - at(yyyy-mm-ddThh:mm:ss), Rate expressions - rate(valueunit), Cron expressions - cron(fields). In UTC. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-Schedule)
	Schedule pulumi.StringInput `pulumi:"schedule"`
	// The namespace of the AWS service. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ServiceNamespace) Example: ecs
	ServiceNamespace pulumi.StringInput `pulumi:"serviceNamespace"`
	// The date and time for the scheduled action to start. Specify the following format: 2006-01-02T15:04:05Z
	StartTime pulumi.StringInput `pulumi:"startTime"`
}
