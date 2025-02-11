// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Provides a resource to manage EC2 Fleets.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_fleet.html.markdown.
    /// </summary>
    public partial class Fleet : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`.
        /// </summary>
        [Output("excessCapacityTerminationPolicy")]
        public Output<string?> ExcessCapacityTerminationPolicy { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing EC2 Launch Template configurations. Defined below.
        /// </summary>
        [Output("launchTemplateConfig")]
        public Output<Outputs.FleetLaunchTemplateConfig> LaunchTemplateConfig { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing On-Demand configurations. Defined below.
        /// </summary>
        [Output("onDemandOptions")]
        public Output<Outputs.FleetOnDemandOptions?> OnDemandOptions { get; private set; } = null!;

        /// <summary>
        /// Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`.
        /// </summary>
        [Output("replaceUnhealthyInstances")]
        public Output<bool?> ReplaceUnhealthyInstances { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing Spot configurations. Defined below.
        /// </summary>
        [Output("spotOptions")]
        public Output<Outputs.FleetSpotOptions?> SpotOptions { get; private set; } = null!;

        /// <summary>
        /// Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing target capacity configurations. Defined below.
        /// </summary>
        [Output("targetCapacitySpecification")]
        public Output<Outputs.FleetTargetCapacitySpecification> TargetCapacitySpecification { get; private set; } = null!;

        /// <summary>
        /// Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
        /// </summary>
        [Output("terminateInstances")]
        public Output<bool?> TerminateInstances { get; private set; } = null!;

        /// <summary>
        /// Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
        /// </summary>
        [Output("terminateInstancesWithExpiration")]
        public Output<bool?> TerminateInstancesWithExpiration { get; private set; } = null!;

        /// <summary>
        /// The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`. Defaults to `maintain`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Fleet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Fleet(string name, FleetArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/fleet:Fleet", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Fleet(string name, Input<string> id, FleetState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/fleet:Fleet", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Fleet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Fleet Get(string name, Input<string> id, FleetState? state = null, CustomResourceOptions? options = null)
        {
            return new Fleet(name, id, state, options);
        }
    }

    public sealed class FleetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`.
        /// </summary>
        [Input("excessCapacityTerminationPolicy")]
        public Input<string>? ExcessCapacityTerminationPolicy { get; set; }

        /// <summary>
        /// Nested argument containing EC2 Launch Template configurations. Defined below.
        /// </summary>
        [Input("launchTemplateConfig", required: true)]
        public Input<Inputs.FleetLaunchTemplateConfigArgs> LaunchTemplateConfig { get; set; } = null!;

        /// <summary>
        /// Nested argument containing On-Demand configurations. Defined below.
        /// </summary>
        [Input("onDemandOptions")]
        public Input<Inputs.FleetOnDemandOptionsArgs>? OnDemandOptions { get; set; }

        /// <summary>
        /// Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`.
        /// </summary>
        [Input("replaceUnhealthyInstances")]
        public Input<bool>? ReplaceUnhealthyInstances { get; set; }

        /// <summary>
        /// Nested argument containing Spot configurations. Defined below.
        /// </summary>
        [Input("spotOptions")]
        public Input<Inputs.FleetSpotOptionsArgs>? SpotOptions { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Nested argument containing target capacity configurations. Defined below.
        /// </summary>
        [Input("targetCapacitySpecification", required: true)]
        public Input<Inputs.FleetTargetCapacitySpecificationArgs> TargetCapacitySpecification { get; set; } = null!;

        /// <summary>
        /// Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
        /// </summary>
        [Input("terminateInstances")]
        public Input<bool>? TerminateInstances { get; set; }

        /// <summary>
        /// Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
        /// </summary>
        [Input("terminateInstancesWithExpiration")]
        public Input<bool>? TerminateInstancesWithExpiration { get; set; }

        /// <summary>
        /// The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`. Defaults to `maintain`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public FleetArgs()
        {
        }
    }

    public sealed class FleetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`.
        /// </summary>
        [Input("excessCapacityTerminationPolicy")]
        public Input<string>? ExcessCapacityTerminationPolicy { get; set; }

        /// <summary>
        /// Nested argument containing EC2 Launch Template configurations. Defined below.
        /// </summary>
        [Input("launchTemplateConfig")]
        public Input<Inputs.FleetLaunchTemplateConfigGetArgs>? LaunchTemplateConfig { get; set; }

        /// <summary>
        /// Nested argument containing On-Demand configurations. Defined below.
        /// </summary>
        [Input("onDemandOptions")]
        public Input<Inputs.FleetOnDemandOptionsGetArgs>? OnDemandOptions { get; set; }

        /// <summary>
        /// Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`.
        /// </summary>
        [Input("replaceUnhealthyInstances")]
        public Input<bool>? ReplaceUnhealthyInstances { get; set; }

        /// <summary>
        /// Nested argument containing Spot configurations. Defined below.
        /// </summary>
        [Input("spotOptions")]
        public Input<Inputs.FleetSpotOptionsGetArgs>? SpotOptions { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Nested argument containing target capacity configurations. Defined below.
        /// </summary>
        [Input("targetCapacitySpecification")]
        public Input<Inputs.FleetTargetCapacitySpecificationGetArgs>? TargetCapacitySpecification { get; set; }

        /// <summary>
        /// Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
        /// </summary>
        [Input("terminateInstances")]
        public Input<bool>? TerminateInstances { get; set; }

        /// <summary>
        /// Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
        /// </summary>
        [Input("terminateInstancesWithExpiration")]
        public Input<bool>? TerminateInstancesWithExpiration { get; set; }

        /// <summary>
        /// The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`. Defaults to `maintain`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public FleetState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class FleetLaunchTemplateConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Nested argument containing EC2 Launch Template to use. Defined below.
        /// </summary>
        [Input("launchTemplateSpecification", required: true)]
        public Input<FleetLaunchTemplateConfigLaunchTemplateSpecificationArgs> LaunchTemplateSpecification { get; set; } = null!;

        [Input("overrides")]
        private InputList<FleetLaunchTemplateConfigOverridesArgs>? _overrides;

        /// <summary>
        /// Nested argument(s) containing parameters to override the same parameters in the Launch Template. Defined below.
        /// </summary>
        public InputList<FleetLaunchTemplateConfigOverridesArgs> Overrides
        {
            get => _overrides ?? (_overrides = new InputList<FleetLaunchTemplateConfigOverridesArgs>());
            set => _overrides = value;
        }

        public FleetLaunchTemplateConfigArgs()
        {
        }
    }

    public sealed class FleetLaunchTemplateConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Nested argument containing EC2 Launch Template to use. Defined below.
        /// </summary>
        [Input("launchTemplateSpecification", required: true)]
        public Input<FleetLaunchTemplateConfigLaunchTemplateSpecificationGetArgs> LaunchTemplateSpecification { get; set; } = null!;

        [Input("overrides")]
        private InputList<FleetLaunchTemplateConfigOverridesGetArgs>? _overrides;

        /// <summary>
        /// Nested argument(s) containing parameters to override the same parameters in the Launch Template. Defined below.
        /// </summary>
        public InputList<FleetLaunchTemplateConfigOverridesGetArgs> Overrides
        {
            get => _overrides ?? (_overrides = new InputList<FleetLaunchTemplateConfigOverridesGetArgs>());
            set => _overrides = value;
        }

        public FleetLaunchTemplateConfigGetArgs()
        {
        }
    }

    public sealed class FleetLaunchTemplateConfigLaunchTemplateSpecificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the launch template.
        /// </summary>
        [Input("launchTemplateId")]
        public Input<string>? LaunchTemplateId { get; set; }

        /// <summary>
        /// Name of the launch template.
        /// </summary>
        [Input("launchTemplateName")]
        public Input<string>? LaunchTemplateName { get; set; }

        /// <summary>
        /// Version number of the launch template.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public FleetLaunchTemplateConfigLaunchTemplateSpecificationArgs()
        {
        }
    }

    public sealed class FleetLaunchTemplateConfigLaunchTemplateSpecificationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the launch template.
        /// </summary>
        [Input("launchTemplateId")]
        public Input<string>? LaunchTemplateId { get; set; }

        /// <summary>
        /// Name of the launch template.
        /// </summary>
        [Input("launchTemplateName")]
        public Input<string>? LaunchTemplateName { get; set; }

        /// <summary>
        /// Version number of the launch template.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public FleetLaunchTemplateConfigLaunchTemplateSpecificationGetArgs()
        {
        }
    }

    public sealed class FleetLaunchTemplateConfigOverridesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Availability Zone in which to launch the instances.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Instance type.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// Maximum price per unit hour that you are willing to pay for a Spot Instance.
        /// </summary>
        [Input("maxPrice")]
        public Input<string>? MaxPrice { get; set; }

        /// <summary>
        /// Priority for the launch template override. If `on_demand_options` `allocation_strategy` is set to `prioritized`, EC2 Fleet uses priority to determine which launch template override to use first in fulfilling On-Demand capacity. The highest priority is launched first. The lower the number, the higher the priority. If no number is set, the launch template override has the lowest priority. Valid values are whole numbers starting at 0.
        /// </summary>
        [Input("priority")]
        public Input<double>? Priority { get; set; }

        /// <summary>
        /// ID of the subnet in which to launch the instances.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// Number of units provided by the specified instance type.
        /// </summary>
        [Input("weightedCapacity")]
        public Input<double>? WeightedCapacity { get; set; }

        public FleetLaunchTemplateConfigOverridesArgs()
        {
        }
    }

    public sealed class FleetLaunchTemplateConfigOverridesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Availability Zone in which to launch the instances.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Instance type.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// Maximum price per unit hour that you are willing to pay for a Spot Instance.
        /// </summary>
        [Input("maxPrice")]
        public Input<string>? MaxPrice { get; set; }

        /// <summary>
        /// Priority for the launch template override. If `on_demand_options` `allocation_strategy` is set to `prioritized`, EC2 Fleet uses priority to determine which launch template override to use first in fulfilling On-Demand capacity. The highest priority is launched first. The lower the number, the higher the priority. If no number is set, the launch template override has the lowest priority. Valid values are whole numbers starting at 0.
        /// </summary>
        [Input("priority")]
        public Input<double>? Priority { get; set; }

        /// <summary>
        /// ID of the subnet in which to launch the instances.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// Number of units provided by the specified instance type.
        /// </summary>
        [Input("weightedCapacity")]
        public Input<double>? WeightedCapacity { get; set; }

        public FleetLaunchTemplateConfigOverridesGetArgs()
        {
        }
    }

    public sealed class FleetOnDemandOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How to allocate the target capacity across the Spot pools. Valid values: `diversified`, `lowestPrice`. Default: `lowestPrice`.
        /// </summary>
        [Input("allocationStrategy")]
        public Input<string>? AllocationStrategy { get; set; }

        public FleetOnDemandOptionsArgs()
        {
        }
    }

    public sealed class FleetOnDemandOptionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How to allocate the target capacity across the Spot pools. Valid values: `diversified`, `lowestPrice`. Default: `lowestPrice`.
        /// </summary>
        [Input("allocationStrategy")]
        public Input<string>? AllocationStrategy { get; set; }

        public FleetOnDemandOptionsGetArgs()
        {
        }
    }

    public sealed class FleetSpotOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How to allocate the target capacity across the Spot pools. Valid values: `diversified`, `lowestPrice`. Default: `lowestPrice`.
        /// </summary>
        [Input("allocationStrategy")]
        public Input<string>? AllocationStrategy { get; set; }

        /// <summary>
        /// Behavior when a Spot Instance is interrupted. Valid values: `hibernate`, `stop`, `terminate`. Default: `terminate`.
        /// </summary>
        [Input("instanceInterruptionBehavior")]
        public Input<string>? InstanceInterruptionBehavior { get; set; }

        /// <summary>
        /// Number of Spot pools across which to allocate your target Spot capacity. Valid only when Spot `allocation_strategy` is set to `lowestPrice`. Default: `1`.
        /// </summary>
        [Input("instancePoolsToUseCount")]
        public Input<int>? InstancePoolsToUseCount { get; set; }

        public FleetSpotOptionsArgs()
        {
        }
    }

    public sealed class FleetSpotOptionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How to allocate the target capacity across the Spot pools. Valid values: `diversified`, `lowestPrice`. Default: `lowestPrice`.
        /// </summary>
        [Input("allocationStrategy")]
        public Input<string>? AllocationStrategy { get; set; }

        /// <summary>
        /// Behavior when a Spot Instance is interrupted. Valid values: `hibernate`, `stop`, `terminate`. Default: `terminate`.
        /// </summary>
        [Input("instanceInterruptionBehavior")]
        public Input<string>? InstanceInterruptionBehavior { get; set; }

        /// <summary>
        /// Number of Spot pools across which to allocate your target Spot capacity. Valid only when Spot `allocation_strategy` is set to `lowestPrice`. Default: `1`.
        /// </summary>
        [Input("instancePoolsToUseCount")]
        public Input<int>? InstancePoolsToUseCount { get; set; }

        public FleetSpotOptionsGetArgs()
        {
        }
    }

    public sealed class FleetTargetCapacitySpecificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default target capacity type. Valid values: `on-demand`, `spot`.
        /// </summary>
        [Input("defaultTargetCapacityType", required: true)]
        public Input<string> DefaultTargetCapacityType { get; set; } = null!;

        /// <summary>
        /// The number of On-Demand units to request.
        /// </summary>
        [Input("onDemandTargetCapacity")]
        public Input<int>? OnDemandTargetCapacity { get; set; }

        /// <summary>
        /// The number of Spot units to request.
        /// </summary>
        [Input("spotTargetCapacity")]
        public Input<int>? SpotTargetCapacity { get; set; }

        /// <summary>
        /// The number of units to request, filled using `default_target_capacity_type`.
        /// </summary>
        [Input("totalTargetCapacity", required: true)]
        public Input<int> TotalTargetCapacity { get; set; } = null!;

        public FleetTargetCapacitySpecificationArgs()
        {
        }
    }

    public sealed class FleetTargetCapacitySpecificationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default target capacity type. Valid values: `on-demand`, `spot`.
        /// </summary>
        [Input("defaultTargetCapacityType", required: true)]
        public Input<string> DefaultTargetCapacityType { get; set; } = null!;

        /// <summary>
        /// The number of On-Demand units to request.
        /// </summary>
        [Input("onDemandTargetCapacity")]
        public Input<int>? OnDemandTargetCapacity { get; set; }

        /// <summary>
        /// The number of Spot units to request.
        /// </summary>
        [Input("spotTargetCapacity")]
        public Input<int>? SpotTargetCapacity { get; set; }

        /// <summary>
        /// The number of units to request, filled using `default_target_capacity_type`.
        /// </summary>
        [Input("totalTargetCapacity", required: true)]
        public Input<int> TotalTargetCapacity { get; set; } = null!;

        public FleetTargetCapacitySpecificationGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class FleetLaunchTemplateConfig
    {
        /// <summary>
        /// Nested argument containing EC2 Launch Template to use. Defined below.
        /// </summary>
        public readonly FleetLaunchTemplateConfigLaunchTemplateSpecification LaunchTemplateSpecification;
        /// <summary>
        /// Nested argument(s) containing parameters to override the same parameters in the Launch Template. Defined below.
        /// </summary>
        public readonly ImmutableArray<FleetLaunchTemplateConfigOverrides> Overrides;

        [OutputConstructor]
        private FleetLaunchTemplateConfig(
            FleetLaunchTemplateConfigLaunchTemplateSpecification launchTemplateSpecification,
            ImmutableArray<FleetLaunchTemplateConfigOverrides> overrides)
        {
            LaunchTemplateSpecification = launchTemplateSpecification;
            Overrides = overrides;
        }
    }

    [OutputType]
    public sealed class FleetLaunchTemplateConfigLaunchTemplateSpecification
    {
        /// <summary>
        /// ID of the launch template.
        /// </summary>
        public readonly string? LaunchTemplateId;
        /// <summary>
        /// Name of the launch template.
        /// </summary>
        public readonly string? LaunchTemplateName;
        /// <summary>
        /// Version number of the launch template.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private FleetLaunchTemplateConfigLaunchTemplateSpecification(
            string? launchTemplateId,
            string? launchTemplateName,
            string version)
        {
            LaunchTemplateId = launchTemplateId;
            LaunchTemplateName = launchTemplateName;
            Version = version;
        }
    }

    [OutputType]
    public sealed class FleetLaunchTemplateConfigOverrides
    {
        /// <summary>
        /// Availability Zone in which to launch the instances.
        /// </summary>
        public readonly string? AvailabilityZone;
        /// <summary>
        /// Instance type.
        /// </summary>
        public readonly string? InstanceType;
        /// <summary>
        /// Maximum price per unit hour that you are willing to pay for a Spot Instance.
        /// </summary>
        public readonly string? MaxPrice;
        /// <summary>
        /// Priority for the launch template override. If `on_demand_options` `allocation_strategy` is set to `prioritized`, EC2 Fleet uses priority to determine which launch template override to use first in fulfilling On-Demand capacity. The highest priority is launched first. The lower the number, the higher the priority. If no number is set, the launch template override has the lowest priority. Valid values are whole numbers starting at 0.
        /// </summary>
        public readonly double? Priority;
        /// <summary>
        /// ID of the subnet in which to launch the instances.
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// Number of units provided by the specified instance type.
        /// </summary>
        public readonly double? WeightedCapacity;

        [OutputConstructor]
        private FleetLaunchTemplateConfigOverrides(
            string? availabilityZone,
            string? instanceType,
            string? maxPrice,
            double? priority,
            string? subnetId,
            double? weightedCapacity)
        {
            AvailabilityZone = availabilityZone;
            InstanceType = instanceType;
            MaxPrice = maxPrice;
            Priority = priority;
            SubnetId = subnetId;
            WeightedCapacity = weightedCapacity;
        }
    }

    [OutputType]
    public sealed class FleetOnDemandOptions
    {
        /// <summary>
        /// How to allocate the target capacity across the Spot pools. Valid values: `diversified`, `lowestPrice`. Default: `lowestPrice`.
        /// </summary>
        public readonly string? AllocationStrategy;

        [OutputConstructor]
        private FleetOnDemandOptions(string? allocationStrategy)
        {
            AllocationStrategy = allocationStrategy;
        }
    }

    [OutputType]
    public sealed class FleetSpotOptions
    {
        /// <summary>
        /// How to allocate the target capacity across the Spot pools. Valid values: `diversified`, `lowestPrice`. Default: `lowestPrice`.
        /// </summary>
        public readonly string? AllocationStrategy;
        /// <summary>
        /// Behavior when a Spot Instance is interrupted. Valid values: `hibernate`, `stop`, `terminate`. Default: `terminate`.
        /// </summary>
        public readonly string? InstanceInterruptionBehavior;
        /// <summary>
        /// Number of Spot pools across which to allocate your target Spot capacity. Valid only when Spot `allocation_strategy` is set to `lowestPrice`. Default: `1`.
        /// </summary>
        public readonly int? InstancePoolsToUseCount;

        [OutputConstructor]
        private FleetSpotOptions(
            string? allocationStrategy,
            string? instanceInterruptionBehavior,
            int? instancePoolsToUseCount)
        {
            AllocationStrategy = allocationStrategy;
            InstanceInterruptionBehavior = instanceInterruptionBehavior;
            InstancePoolsToUseCount = instancePoolsToUseCount;
        }
    }

    [OutputType]
    public sealed class FleetTargetCapacitySpecification
    {
        /// <summary>
        /// Default target capacity type. Valid values: `on-demand`, `spot`.
        /// </summary>
        public readonly string DefaultTargetCapacityType;
        /// <summary>
        /// The number of On-Demand units to request.
        /// </summary>
        public readonly int? OnDemandTargetCapacity;
        /// <summary>
        /// The number of Spot units to request.
        /// </summary>
        public readonly int? SpotTargetCapacity;
        /// <summary>
        /// The number of units to request, filled using `default_target_capacity_type`.
        /// </summary>
        public readonly int TotalTargetCapacity;

        [OutputConstructor]
        private FleetTargetCapacitySpecification(
            string defaultTargetCapacityType,
            int? onDemandTargetCapacity,
            int? spotTargetCapacity,
            int totalTargetCapacity)
        {
            DefaultTargetCapacityType = defaultTargetCapacityType;
            OnDemandTargetCapacity = onDemandTargetCapacity;
            SpotTargetCapacity = spotTargetCapacity;
            TotalTargetCapacity = totalTargetCapacity;
        }
    }
    }
}
