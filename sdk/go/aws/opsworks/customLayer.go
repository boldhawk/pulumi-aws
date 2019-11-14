// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an OpsWorks custom layer resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/opsworks_custom_layer.html.markdown.
type CustomLayer struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolOutput `pulumi:"autoAssignElasticIps"`

	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolOutput `pulumi:"autoAssignPublicIps"`

	// Whether to enable auto-healing for the layer.
	AutoHealing pulumi.BoolOutput `pulumi:"autoHealing"`

	CustomConfigureRecipes pulumi.ArrayOutput `pulumi:"customConfigureRecipes"`

	CustomDeployRecipes pulumi.ArrayOutput `pulumi:"customDeployRecipes"`

	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn pulumi.StringOutput `pulumi:"customInstanceProfileArn"`

	// Custom JSON attributes to apply to the layer.
	CustomJson pulumi.StringOutput `pulumi:"customJson"`

	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.ArrayOutput `pulumi:"customSecurityGroupIds"`

	CustomSetupRecipes pulumi.ArrayOutput `pulumi:"customSetupRecipes"`

	CustomShutdownRecipes pulumi.ArrayOutput `pulumi:"customShutdownRecipes"`

	CustomUndeployRecipes pulumi.ArrayOutput `pulumi:"customUndeployRecipes"`

	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolOutput `pulumi:"drainElbOnShutdown"`

	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes pulumi.ArrayOutput `pulumi:"ebsVolumes"`

	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringOutput `pulumi:"elasticLoadBalancer"`

	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolOutput `pulumi:"installUpdatesOnBoot"`

	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntOutput `pulumi:"instanceShutdownTimeout"`

	// A human-readable name for the layer.
	Name pulumi.StringOutput `pulumi:"name"`

	// A short, machine-readable name for the layer, which will be used to identify it in the Chef node JSON.
	ShortName pulumi.StringOutput `pulumi:"shortName"`

	// The id of the stack the layer will belong to.
	StackId pulumi.StringOutput `pulumi:"stackId"`

	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.ArrayOutput `pulumi:"systemPackages"`

	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolOutput `pulumi:"useEbsOptimizedInstances"`
}

// NewCustomLayer registers a new resource with the given unique name, arguments, and options.
func NewCustomLayer(ctx *pulumi.Context,
	name string, args *CustomLayerArgs, opts ...pulumi.ResourceOpt) (*CustomLayer, error) {
	if args == nil || args.ShortName == nil {
		return nil, errors.New("missing required argument 'ShortName'")
	}
	if args == nil || args.StackId == nil {
		return nil, errors.New("missing required argument 'StackId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
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
		inputs["shortName"] = args.ShortName
		inputs["stackId"] = args.StackId
		inputs["systemPackages"] = args.SystemPackages
		inputs["useEbsOptimizedInstances"] = args.UseEbsOptimizedInstances
	}
	var resource CustomLayer
	err := ctx.RegisterResource("aws:opsworks/customLayer:CustomLayer", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomLayer gets an existing CustomLayer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomLayer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *CustomLayerState, opts ...pulumi.ResourceOpt) (*CustomLayer, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
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
		inputs["shortName"] = state.ShortName
		inputs["stackId"] = state.StackId
		inputs["systemPackages"] = state.SystemPackages
		inputs["useEbsOptimizedInstances"] = state.UseEbsOptimizedInstances
	}
	var resource CustomLayer
	err := ctx.ReadResource("aws:opsworks/customLayer:CustomLayer", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *CustomLayer) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *CustomLayer) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering CustomLayer resources.
type CustomLayerState struct {
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolInput `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolInput `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing pulumi.BoolInput `pulumi:"autoHealing"`
	CustomConfigureRecipes pulumi.ArrayInput `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes pulumi.ArrayInput `pulumi:"customDeployRecipes"`
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn pulumi.StringInput `pulumi:"customInstanceProfileArn"`
	// Custom JSON attributes to apply to the layer.
	CustomJson pulumi.StringInput `pulumi:"customJson"`
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.ArrayInput `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes pulumi.ArrayInput `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes pulumi.ArrayInput `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes pulumi.ArrayInput `pulumi:"customUndeployRecipes"`
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolInput `pulumi:"drainElbOnShutdown"`
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes pulumi.ArrayInput `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringInput `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolInput `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntInput `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name pulumi.StringInput `pulumi:"name"`
	// A short, machine-readable name for the layer, which will be used to identify it in the Chef node JSON.
	ShortName pulumi.StringInput `pulumi:"shortName"`
	// The id of the stack the layer will belong to.
	StackId pulumi.StringInput `pulumi:"stackId"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.ArrayInput `pulumi:"systemPackages"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolInput `pulumi:"useEbsOptimizedInstances"`
}

// The set of arguments for constructing a CustomLayer resource.
type CustomLayerArgs struct {
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolInput `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolInput `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing pulumi.BoolInput `pulumi:"autoHealing"`
	CustomConfigureRecipes pulumi.ArrayInput `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes pulumi.ArrayInput `pulumi:"customDeployRecipes"`
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn pulumi.StringInput `pulumi:"customInstanceProfileArn"`
	// Custom JSON attributes to apply to the layer.
	CustomJson pulumi.StringInput `pulumi:"customJson"`
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.ArrayInput `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes pulumi.ArrayInput `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes pulumi.ArrayInput `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes pulumi.ArrayInput `pulumi:"customUndeployRecipes"`
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolInput `pulumi:"drainElbOnShutdown"`
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes pulumi.ArrayInput `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringInput `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolInput `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntInput `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name pulumi.StringInput `pulumi:"name"`
	// A short, machine-readable name for the layer, which will be used to identify it in the Chef node JSON.
	ShortName pulumi.StringInput `pulumi:"shortName"`
	// The id of the stack the layer will belong to.
	StackId pulumi.StringInput `pulumi:"stackId"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.ArrayInput `pulumi:"systemPackages"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolInput `pulumi:"useEbsOptimizedInstances"`
}
