// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appautoscaling

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Application AutoScaling Policy resource.
// 
// ## Nested fields
// 
// ### `targetTrackingScalingPolicyConfiguration`
// 
// * `targetValue` - (Required) The target value for the metric.
// * `disableScaleIn` - (Optional) Indicates whether scale in by the target tracking policy is disabled. If the value is true, scale in is disabled and the target tracking policy won't remove capacity from the scalable resource. Otherwise, scale in is enabled and the target tracking policy can remove capacity from the scalable resource. The default value is `false`.
// * `scaleInCooldown` - (Optional) The amount of time, in seconds, after a scale in activity completes before another scale in activity can start.
// * `scaleOutCooldown` - (Optional) The amount of time, in seconds, after a scale out activity completes before another scale out activity can start.
// * `customizedMetricSpecification` - (Optional) A custom CloudWatch metric. Documentation can be found  at: [AWS Customized Metric Specification](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CustomizedMetricSpecification.html). See supported fields below.
// * `predefinedMetricSpecification` - (Optional) A predefined metric. See supported fields below.
// 
// ### `customizedMetricSpecification`
// 
// * `dimensions` - (Optional) The dimensions of the metric.
// * `metricName` - (Required) The name of the metric.
// * `namespace` - (Required) The namespace of the metric.
// * `statistic` - (Required) The statistic of the metric.
// * `unit` - (Optional) The unit of the metric.
// 
// ### `predefinedMetricSpecification`
// 
// * `predefinedMetricType` - (Required) The metric type.
// * `resourceLabel` - (Optional) Reserved for future use.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appautoscaling_policy.html.markdown.
type Policy struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	Alarms pulumi.ArrayOutput `pulumi:"alarms"`

	// The ARN assigned by AWS to the scaling policy.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The name of the policy.
	Name pulumi.StringOutput `pulumi:"name"`

	// For DynamoDB, only `TargetTrackingScaling` is supported. For Amazon ECS, Spot Fleet, and Amazon RDS, both `StepScaling` and `TargetTrackingScaling` are supported. For any other service, only `StepScaling` is supported. Defaults to `StepScaling`.
	PolicyType pulumi.StringOutput `pulumi:"policyType"`

	// The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`

	// The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ScalableDimension pulumi.StringOutput `pulumi:"scalableDimension"`

	// The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ServiceNamespace pulumi.StringOutput `pulumi:"serviceNamespace"`

	// Step scaling policy configuration, requires `policyType = "StepScaling"` (default). See supported fields below.
	StepScalingPolicyConfiguration pulumi.AnyOutput `pulumi:"stepScalingPolicyConfiguration"`

	// A target tracking policy, requires `policyType = "TargetTrackingScaling"`. See supported fields below.
	TargetTrackingScalingPolicyConfiguration pulumi.AnyOutput `pulumi:"targetTrackingScalingPolicyConfiguration"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOpt) (*Policy, error) {
	if args == nil || args.ResourceId == nil {
		return nil, errors.New("missing required argument 'ResourceId'")
	}
	if args == nil || args.ScalableDimension == nil {
		return nil, errors.New("missing required argument 'ScalableDimension'")
	}
	if args == nil || args.ServiceNamespace == nil {
		return nil, errors.New("missing required argument 'ServiceNamespace'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["alarms"] = args.Alarms
		inputs["name"] = args.Name
		inputs["policyType"] = args.PolicyType
		inputs["resourceId"] = args.ResourceId
		inputs["scalableDimension"] = args.ScalableDimension
		inputs["serviceNamespace"] = args.ServiceNamespace
		inputs["stepScalingPolicyConfiguration"] = args.StepScalingPolicyConfiguration
		inputs["targetTrackingScalingPolicyConfiguration"] = args.TargetTrackingScalingPolicyConfiguration
	}
	var resource Policy
	err := ctx.RegisterResource("aws:appautoscaling/policy:Policy", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PolicyState, opts ...pulumi.ResourceOpt) (*Policy, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["alarms"] = state.Alarms
		inputs["arn"] = state.Arn
		inputs["name"] = state.Name
		inputs["policyType"] = state.PolicyType
		inputs["resourceId"] = state.ResourceId
		inputs["scalableDimension"] = state.ScalableDimension
		inputs["serviceNamespace"] = state.ServiceNamespace
		inputs["stepScalingPolicyConfiguration"] = state.StepScalingPolicyConfiguration
		inputs["targetTrackingScalingPolicyConfiguration"] = state.TargetTrackingScalingPolicyConfiguration
	}
	var resource Policy
	err := ctx.ReadResource("aws:appautoscaling/policy:Policy", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *Policy) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *Policy) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering Policy resources.
type PolicyState struct {
	Alarms pulumi.ArrayInput `pulumi:"alarms"`
	// The ARN assigned by AWS to the scaling policy.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The name of the policy.
	Name pulumi.StringInput `pulumi:"name"`
	// For DynamoDB, only `TargetTrackingScaling` is supported. For Amazon ECS, Spot Fleet, and Amazon RDS, both `StepScaling` and `TargetTrackingScaling` are supported. For any other service, only `StepScaling` is supported. Defaults to `StepScaling`.
	PolicyType pulumi.StringInput `pulumi:"policyType"`
	// The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ScalableDimension pulumi.StringInput `pulumi:"scalableDimension"`
	// The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ServiceNamespace pulumi.StringInput `pulumi:"serviceNamespace"`
	// Step scaling policy configuration, requires `policyType = "StepScaling"` (default). See supported fields below.
	StepScalingPolicyConfiguration pulumi.AnyInput `pulumi:"stepScalingPolicyConfiguration"`
	// A target tracking policy, requires `policyType = "TargetTrackingScaling"`. See supported fields below.
	TargetTrackingScalingPolicyConfiguration pulumi.AnyInput `pulumi:"targetTrackingScalingPolicyConfiguration"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	Alarms pulumi.ArrayInput `pulumi:"alarms"`
	// The name of the policy.
	Name pulumi.StringInput `pulumi:"name"`
	// For DynamoDB, only `TargetTrackingScaling` is supported. For Amazon ECS, Spot Fleet, and Amazon RDS, both `StepScaling` and `TargetTrackingScaling` are supported. For any other service, only `StepScaling` is supported. Defaults to `StepScaling`.
	PolicyType pulumi.StringInput `pulumi:"policyType"`
	// The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ScalableDimension pulumi.StringInput `pulumi:"scalableDimension"`
	// The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ServiceNamespace pulumi.StringInput `pulumi:"serviceNamespace"`
	// Step scaling policy configuration, requires `policyType = "StepScaling"` (default). See supported fields below.
	StepScalingPolicyConfiguration pulumi.AnyInput `pulumi:"stepScalingPolicyConfiguration"`
	// A target tracking policy, requires `policyType = "TargetTrackingScaling"`. See supported fields below.
	TargetTrackingScalingPolicyConfiguration pulumi.AnyInput `pulumi:"targetTrackingScalingPolicyConfiguration"`
}
