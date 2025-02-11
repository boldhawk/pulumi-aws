// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Gamelift Fleet resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/gamelift_fleet.html.markdown.
type Fleet struct {
	s *pulumi.ResourceState
}

// NewFleet registers a new resource with the given unique name, arguments, and options.
func NewFleet(ctx *pulumi.Context,
	name string, args *FleetArgs, opts ...pulumi.ResourceOpt) (*Fleet, error) {
	if args == nil || args.BuildId == nil {
		return nil, errors.New("missing required argument 'BuildId'")
	}
	if args == nil || args.Ec2InstanceType == nil {
		return nil, errors.New("missing required argument 'Ec2InstanceType'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["buildId"] = nil
		inputs["description"] = nil
		inputs["ec2InboundPermissions"] = nil
		inputs["ec2InstanceType"] = nil
		inputs["metricGroups"] = nil
		inputs["name"] = nil
		inputs["newGameSessionProtectionPolicy"] = nil
		inputs["resourceCreationLimitPolicy"] = nil
		inputs["runtimeConfiguration"] = nil
	} else {
		inputs["buildId"] = args.BuildId
		inputs["description"] = args.Description
		inputs["ec2InboundPermissions"] = args.Ec2InboundPermissions
		inputs["ec2InstanceType"] = args.Ec2InstanceType
		inputs["metricGroups"] = args.MetricGroups
		inputs["name"] = args.Name
		inputs["newGameSessionProtectionPolicy"] = args.NewGameSessionProtectionPolicy
		inputs["resourceCreationLimitPolicy"] = args.ResourceCreationLimitPolicy
		inputs["runtimeConfiguration"] = args.RuntimeConfiguration
	}
	inputs["arn"] = nil
	inputs["logPaths"] = nil
	inputs["operatingSystem"] = nil
	s, err := ctx.RegisterResource("aws:gamelift/fleet:Fleet", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Fleet{s: s}, nil
}

// GetFleet gets an existing Fleet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFleet(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FleetState, opts ...pulumi.ResourceOpt) (*Fleet, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["buildId"] = state.BuildId
		inputs["description"] = state.Description
		inputs["ec2InboundPermissions"] = state.Ec2InboundPermissions
		inputs["ec2InstanceType"] = state.Ec2InstanceType
		inputs["logPaths"] = state.LogPaths
		inputs["metricGroups"] = state.MetricGroups
		inputs["name"] = state.Name
		inputs["newGameSessionProtectionPolicy"] = state.NewGameSessionProtectionPolicy
		inputs["operatingSystem"] = state.OperatingSystem
		inputs["resourceCreationLimitPolicy"] = state.ResourceCreationLimitPolicy
		inputs["runtimeConfiguration"] = state.RuntimeConfiguration
	}
	s, err := ctx.ReadResource("aws:gamelift/fleet:Fleet", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Fleet{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Fleet) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Fleet) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Fleet ARN.
func (r *Fleet) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// ID of the Gamelift Build to be deployed on the fleet.
func (r *Fleet) BuildId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["buildId"])
}

// Human-readable description of the fleet.
func (r *Fleet) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
func (r *Fleet) Ec2InboundPermissions() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["ec2InboundPermissions"])
}

// Name of an EC2 instance type. e.g. `t2.micro`
func (r *Fleet) Ec2InstanceType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ec2InstanceType"])
}

func (r *Fleet) LogPaths() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["logPaths"])
}

// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
func (r *Fleet) MetricGroups() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["metricGroups"])
}

// The name of the fleet.
func (r *Fleet) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
func (r *Fleet) NewGameSessionProtectionPolicy() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["newGameSessionProtectionPolicy"])
}

// Operating system of the fleet's computing resources.
func (r *Fleet) OperatingSystem() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["operatingSystem"])
}

// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
func (r *Fleet) ResourceCreationLimitPolicy() pulumi.Output {
	return r.s.State["resourceCreationLimitPolicy"]
}

// Instructions for launching server processes on each instance in the fleet. See below.
func (r *Fleet) RuntimeConfiguration() pulumi.Output {
	return r.s.State["runtimeConfiguration"]
}

// Input properties used for looking up and filtering Fleet resources.
type FleetState struct {
	// Fleet ARN.
	Arn interface{}
	// ID of the Gamelift Build to be deployed on the fleet.
	BuildId interface{}
	// Human-readable description of the fleet.
	Description interface{}
	// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
	Ec2InboundPermissions interface{}
	// Name of an EC2 instance type. e.g. `t2.micro`
	Ec2InstanceType interface{}
	LogPaths interface{}
	// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
	MetricGroups interface{}
	// The name of the fleet.
	Name interface{}
	// Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
	NewGameSessionProtectionPolicy interface{}
	// Operating system of the fleet's computing resources.
	OperatingSystem interface{}
	// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
	ResourceCreationLimitPolicy interface{}
	// Instructions for launching server processes on each instance in the fleet. See below.
	RuntimeConfiguration interface{}
}

// The set of arguments for constructing a Fleet resource.
type FleetArgs struct {
	// ID of the Gamelift Build to be deployed on the fleet.
	BuildId interface{}
	// Human-readable description of the fleet.
	Description interface{}
	// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
	Ec2InboundPermissions interface{}
	// Name of an EC2 instance type. e.g. `t2.micro`
	Ec2InstanceType interface{}
	// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
	MetricGroups interface{}
	// The name of the fleet.
	Name interface{}
	// Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
	NewGameSessionProtectionPolicy interface{}
	// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
	ResourceCreationLimitPolicy interface{}
	// Instructions for launching server processes on each instance in the fleet. See below.
	RuntimeConfiguration interface{}
}
