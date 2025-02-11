// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an OpsWorks memcached layer resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/opsworks_memcached_layer.html.markdown.
type MemcachedLayer struct {
	s *pulumi.ResourceState
}

// NewMemcachedLayer registers a new resource with the given unique name, arguments, and options.
func NewMemcachedLayer(ctx *pulumi.Context,
	name string, args *MemcachedLayerArgs, opts ...pulumi.ResourceOpt) (*MemcachedLayer, error) {
	if args == nil || args.StackId == nil {
		return nil, errors.New("missing required argument 'StackId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["allocatedMemory"] = nil
		inputs["autoAssignElasticIps"] = nil
		inputs["autoAssignPublicIps"] = nil
		inputs["autoHealing"] = nil
		inputs["customConfigureRecipes"] = nil
		inputs["customDeployRecipes"] = nil
		inputs["customInstanceProfileArn"] = nil
		inputs["customJson"] = nil
		inputs["customSecurityGroupIds"] = nil
		inputs["customSetupRecipes"] = nil
		inputs["customShutdownRecipes"] = nil
		inputs["customUndeployRecipes"] = nil
		inputs["drainElbOnShutdown"] = nil
		inputs["ebsVolumes"] = nil
		inputs["elasticLoadBalancer"] = nil
		inputs["installUpdatesOnBoot"] = nil
		inputs["instanceShutdownTimeout"] = nil
		inputs["name"] = nil
		inputs["stackId"] = nil
		inputs["systemPackages"] = nil
		inputs["useEbsOptimizedInstances"] = nil
	} else {
		inputs["allocatedMemory"] = args.AllocatedMemory
		inputs["autoAssignElasticIps"] = args.AutoAssignElasticIps
		inputs["autoAssignPublicIps"] = args.AutoAssignPublicIps
		inputs["autoHealing"] = args.AutoHealing
		inputs["customConfigureRecipes"] = args.CustomConfigureRecipes
		inputs["customDeployRecipes"] = args.CustomDeployRecipes
		inputs["customInstanceProfileArn"] = args.CustomInstanceProfileArn
		inputs["customJson"] = args.CustomJson
		inputs["customSecurityGroupIds"] = args.CustomSecurityGroupIds
		inputs["customSetupRecipes"] = args.CustomSetupRecipes
		inputs["customShutdownRecipes"] = args.CustomShutdownRecipes
		inputs["customUndeployRecipes"] = args.CustomUndeployRecipes
		inputs["drainElbOnShutdown"] = args.DrainElbOnShutdown
		inputs["ebsVolumes"] = args.EbsVolumes
		inputs["elasticLoadBalancer"] = args.ElasticLoadBalancer
		inputs["installUpdatesOnBoot"] = args.InstallUpdatesOnBoot
		inputs["instanceShutdownTimeout"] = args.InstanceShutdownTimeout
		inputs["name"] = args.Name
		inputs["stackId"] = args.StackId
		inputs["systemPackages"] = args.SystemPackages
		inputs["useEbsOptimizedInstances"] = args.UseEbsOptimizedInstances
	}
	s, err := ctx.RegisterResource("aws:opsworks/memcachedLayer:MemcachedLayer", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MemcachedLayer{s: s}, nil
}

// GetMemcachedLayer gets an existing MemcachedLayer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMemcachedLayer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *MemcachedLayerState, opts ...pulumi.ResourceOpt) (*MemcachedLayer, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["allocatedMemory"] = state.AllocatedMemory
		inputs["autoAssignElasticIps"] = state.AutoAssignElasticIps
		inputs["autoAssignPublicIps"] = state.AutoAssignPublicIps
		inputs["autoHealing"] = state.AutoHealing
		inputs["customConfigureRecipes"] = state.CustomConfigureRecipes
		inputs["customDeployRecipes"] = state.CustomDeployRecipes
		inputs["customInstanceProfileArn"] = state.CustomInstanceProfileArn
		inputs["customJson"] = state.CustomJson
		inputs["customSecurityGroupIds"] = state.CustomSecurityGroupIds
		inputs["customSetupRecipes"] = state.CustomSetupRecipes
		inputs["customShutdownRecipes"] = state.CustomShutdownRecipes
		inputs["customUndeployRecipes"] = state.CustomUndeployRecipes
		inputs["drainElbOnShutdown"] = state.DrainElbOnShutdown
		inputs["ebsVolumes"] = state.EbsVolumes
		inputs["elasticLoadBalancer"] = state.ElasticLoadBalancer
		inputs["installUpdatesOnBoot"] = state.InstallUpdatesOnBoot
		inputs["instanceShutdownTimeout"] = state.InstanceShutdownTimeout
		inputs["name"] = state.Name
		inputs["stackId"] = state.StackId
		inputs["systemPackages"] = state.SystemPackages
		inputs["useEbsOptimizedInstances"] = state.UseEbsOptimizedInstances
	}
	s, err := ctx.ReadResource("aws:opsworks/memcachedLayer:MemcachedLayer", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MemcachedLayer{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *MemcachedLayer) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *MemcachedLayer) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Amount of memory to allocate for the cache on each instance, in megabytes. Defaults to 512MB.
func (r *MemcachedLayer) AllocatedMemory() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["allocatedMemory"])
}

// Whether to automatically assign an elastic IP address to the layer's instances.
func (r *MemcachedLayer) AutoAssignElasticIps() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["autoAssignElasticIps"])
}

// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
func (r *MemcachedLayer) AutoAssignPublicIps() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["autoAssignPublicIps"])
}

// Whether to enable auto-healing for the layer.
func (r *MemcachedLayer) AutoHealing() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["autoHealing"])
}

func (r *MemcachedLayer) CustomConfigureRecipes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["customConfigureRecipes"])
}

func (r *MemcachedLayer) CustomDeployRecipes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["customDeployRecipes"])
}

// The ARN of an IAM profile that will be used for the layer's instances.
func (r *MemcachedLayer) CustomInstanceProfileArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["customInstanceProfileArn"])
}

// Custom JSON attributes to apply to the layer.
func (r *MemcachedLayer) CustomJson() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["customJson"])
}

// Ids for a set of security groups to apply to the layer's instances.
func (r *MemcachedLayer) CustomSecurityGroupIds() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["customSecurityGroupIds"])
}

func (r *MemcachedLayer) CustomSetupRecipes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["customSetupRecipes"])
}

func (r *MemcachedLayer) CustomShutdownRecipes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["customShutdownRecipes"])
}

func (r *MemcachedLayer) CustomUndeployRecipes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["customUndeployRecipes"])
}

// Whether to enable Elastic Load Balancing connection draining.
func (r *MemcachedLayer) DrainElbOnShutdown() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["drainElbOnShutdown"])
}

// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
func (r *MemcachedLayer) EbsVolumes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["ebsVolumes"])
}

// Name of an Elastic Load Balancer to attach to this layer
func (r *MemcachedLayer) ElasticLoadBalancer() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["elasticLoadBalancer"])
}

// Whether to install OS and package updates on each instance when it boots.
func (r *MemcachedLayer) InstallUpdatesOnBoot() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["installUpdatesOnBoot"])
}

// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
func (r *MemcachedLayer) InstanceShutdownTimeout() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["instanceShutdownTimeout"])
}

// A human-readable name for the layer.
func (r *MemcachedLayer) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The id of the stack the layer will belong to.
func (r *MemcachedLayer) StackId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["stackId"])
}

// Names of a set of system packages to install on the layer's instances.
func (r *MemcachedLayer) SystemPackages() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["systemPackages"])
}

// Whether to use EBS-optimized instances.
func (r *MemcachedLayer) UseEbsOptimizedInstances() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["useEbsOptimizedInstances"])
}

// Input properties used for looking up and filtering MemcachedLayer resources.
type MemcachedLayerState struct {
	// Amount of memory to allocate for the cache on each instance, in megabytes. Defaults to 512MB.
	AllocatedMemory interface{}
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps interface{}
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps interface{}
	// Whether to enable auto-healing for the layer.
	AutoHealing interface{}
	CustomConfigureRecipes interface{}
	CustomDeployRecipes interface{}
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn interface{}
	// Custom JSON attributes to apply to the layer.
	CustomJson interface{}
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds interface{}
	CustomSetupRecipes interface{}
	CustomShutdownRecipes interface{}
	CustomUndeployRecipes interface{}
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown interface{}
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes interface{}
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer interface{}
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot interface{}
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout interface{}
	// A human-readable name for the layer.
	Name interface{}
	// The id of the stack the layer will belong to.
	StackId interface{}
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages interface{}
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances interface{}
}

// The set of arguments for constructing a MemcachedLayer resource.
type MemcachedLayerArgs struct {
	// Amount of memory to allocate for the cache on each instance, in megabytes. Defaults to 512MB.
	AllocatedMemory interface{}
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps interface{}
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps interface{}
	// Whether to enable auto-healing for the layer.
	AutoHealing interface{}
	CustomConfigureRecipes interface{}
	CustomDeployRecipes interface{}
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn interface{}
	// Custom JSON attributes to apply to the layer.
	CustomJson interface{}
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds interface{}
	CustomSetupRecipes interface{}
	CustomShutdownRecipes interface{}
	CustomUndeployRecipes interface{}
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown interface{}
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes interface{}
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer interface{}
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot interface{}
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout interface{}
	// A human-readable name for the layer.
	Name interface{}
	// The id of the stack the layer will belong to.
	StackId interface{}
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages interface{}
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances interface{}
}
