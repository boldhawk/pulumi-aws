// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an OpsWorks stack resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/opsworks_stack.html.markdown.
type Stack struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// If set to `"LATEST"`, OpsWorks will automatically install the latest version.
	AgentVersion pulumi.StringOutput `pulumi:"agentVersion"`

	Arn pulumi.StringOutput `pulumi:"arn"`

	// If `manageBerkshelf` is enabled, the version of Berkshelf to use.
	BerkshelfVersion pulumi.StringOutput `pulumi:"berkshelfVersion"`

	// Color to paint next to the stack's resources in the OpsWorks console.
	Color pulumi.StringOutput `pulumi:"color"`

	// Name of the configuration manager to use. Defaults to "Chef".
	ConfigurationManagerName pulumi.StringOutput `pulumi:"configurationManagerName"`

	// Version of the configuration manager to use. Defaults to "11.4".
	ConfigurationManagerVersion pulumi.StringOutput `pulumi:"configurationManagerVersion"`

	// When `useCustomCookbooks` is set, provide this sub-object as
	// described below.
	CustomCookbooksSources pulumi.ArrayOutput `pulumi:"customCookbooksSources"`

	// Custom JSON attributes to apply to the entire stack.
	CustomJson pulumi.StringOutput `pulumi:"customJson"`

	// Name of the availability zone where instances will be created
	// by default. This is required unless you set `vpcId`.
	DefaultAvailabilityZone pulumi.StringOutput `pulumi:"defaultAvailabilityZone"`

	// The ARN of an IAM Instance Profile that created instances
	// will have by default.
	DefaultInstanceProfileArn pulumi.StringOutput `pulumi:"defaultInstanceProfileArn"`

	// Name of OS that will be installed on instances by default.
	DefaultOs pulumi.StringOutput `pulumi:"defaultOs"`

	// Name of the type of root device instances will have by default.
	DefaultRootDeviceType pulumi.StringOutput `pulumi:"defaultRootDeviceType"`

	// Name of the SSH keypair that instances will have by default.
	DefaultSshKeyName pulumi.StringOutput `pulumi:"defaultSshKeyName"`

	// Id of the subnet in which instances will be created by default. Mandatory
	// if `vpcId` is set, and forbidden if it isn't.
	DefaultSubnetId pulumi.StringOutput `pulumi:"defaultSubnetId"`

	// Keyword representing the naming scheme that will be used for instance hostnames
	// within this stack.
	HostnameTheme pulumi.StringOutput `pulumi:"hostnameTheme"`

	// Boolean value controlling whether Opsworks will run Berkshelf for this stack.
	ManageBerkshelf pulumi.BoolOutput `pulumi:"manageBerkshelf"`

	// The name of the stack.
	Name pulumi.StringOutput `pulumi:"name"`

	// The name of the region where the stack will exist.
	Region pulumi.StringOutput `pulumi:"region"`

	// The ARN of an IAM role that the OpsWorks service will act as.
	ServiceRoleArn pulumi.StringOutput `pulumi:"serviceRoleArn"`

	StackEndpoint pulumi.StringOutput `pulumi:"stackEndpoint"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// Boolean value controlling whether the custom cookbook settings are
	// enabled.
	UseCustomCookbooks pulumi.BoolOutput `pulumi:"useCustomCookbooks"`

	// Boolean value controlling whether the standard OpsWorks
	// security groups apply to created instances.
	UseOpsworksSecurityGroups pulumi.BoolOutput `pulumi:"useOpsworksSecurityGroups"`

	// The id of the VPC that this stack belongs to.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewStack registers a new resource with the given unique name, arguments, and options.
func NewStack(ctx *pulumi.Context,
	name string, args *StackArgs, opts ...pulumi.ResourceOpt) (*Stack, error) {
	if args == nil || args.DefaultInstanceProfileArn == nil {
		return nil, errors.New("missing required argument 'DefaultInstanceProfileArn'")
	}
	if args == nil || args.Region == nil {
		return nil, errors.New("missing required argument 'Region'")
	}
	if args == nil || args.ServiceRoleArn == nil {
		return nil, errors.New("missing required argument 'ServiceRoleArn'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["agentVersion"] = args.AgentVersion
		inputs["berkshelfVersion"] = args.BerkshelfVersion
		inputs["color"] = args.Color
		inputs["configurationManagerName"] = args.ConfigurationManagerName
		inputs["configurationManagerVersion"] = args.ConfigurationManagerVersion
		inputs["customCookbooksSources"] = args.CustomCookbooksSources
		inputs["customJson"] = args.CustomJson
		inputs["defaultAvailabilityZone"] = args.DefaultAvailabilityZone
		inputs["defaultInstanceProfileArn"] = args.DefaultInstanceProfileArn
		inputs["defaultOs"] = args.DefaultOs
		inputs["defaultRootDeviceType"] = args.DefaultRootDeviceType
		inputs["defaultSshKeyName"] = args.DefaultSshKeyName
		inputs["defaultSubnetId"] = args.DefaultSubnetId
		inputs["hostnameTheme"] = args.HostnameTheme
		inputs["manageBerkshelf"] = args.ManageBerkshelf
		inputs["name"] = args.Name
		inputs["region"] = args.Region
		inputs["serviceRoleArn"] = args.ServiceRoleArn
		inputs["tags"] = args.Tags
		inputs["useCustomCookbooks"] = args.UseCustomCookbooks
		inputs["useOpsworksSecurityGroups"] = args.UseOpsworksSecurityGroups
		inputs["vpcId"] = args.VpcId
	}
	var resource Stack
	err := ctx.RegisterResource("aws:opsworks/stack:Stack", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStack gets an existing Stack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStack(ctx *pulumi.Context,
	name string, id pulumi.ID, state *StackState, opts ...pulumi.ResourceOpt) (*Stack, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["agentVersion"] = state.AgentVersion
		inputs["arn"] = state.Arn
		inputs["berkshelfVersion"] = state.BerkshelfVersion
		inputs["color"] = state.Color
		inputs["configurationManagerName"] = state.ConfigurationManagerName
		inputs["configurationManagerVersion"] = state.ConfigurationManagerVersion
		inputs["customCookbooksSources"] = state.CustomCookbooksSources
		inputs["customJson"] = state.CustomJson
		inputs["defaultAvailabilityZone"] = state.DefaultAvailabilityZone
		inputs["defaultInstanceProfileArn"] = state.DefaultInstanceProfileArn
		inputs["defaultOs"] = state.DefaultOs
		inputs["defaultRootDeviceType"] = state.DefaultRootDeviceType
		inputs["defaultSshKeyName"] = state.DefaultSshKeyName
		inputs["defaultSubnetId"] = state.DefaultSubnetId
		inputs["hostnameTheme"] = state.HostnameTheme
		inputs["manageBerkshelf"] = state.ManageBerkshelf
		inputs["name"] = state.Name
		inputs["region"] = state.Region
		inputs["serviceRoleArn"] = state.ServiceRoleArn
		inputs["stackEndpoint"] = state.StackEndpoint
		inputs["tags"] = state.Tags
		inputs["useCustomCookbooks"] = state.UseCustomCookbooks
		inputs["useOpsworksSecurityGroups"] = state.UseOpsworksSecurityGroups
		inputs["vpcId"] = state.VpcId
	}
	var resource Stack
	err := ctx.ReadResource("aws:opsworks/stack:Stack", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *Stack) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *Stack) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering Stack resources.
type StackState struct {
	// If set to `"LATEST"`, OpsWorks will automatically install the latest version.
	AgentVersion pulumi.StringInput `pulumi:"agentVersion"`
	Arn pulumi.StringInput `pulumi:"arn"`
	// If `manageBerkshelf` is enabled, the version of Berkshelf to use.
	BerkshelfVersion pulumi.StringInput `pulumi:"berkshelfVersion"`
	// Color to paint next to the stack's resources in the OpsWorks console.
	Color pulumi.StringInput `pulumi:"color"`
	// Name of the configuration manager to use. Defaults to "Chef".
	ConfigurationManagerName pulumi.StringInput `pulumi:"configurationManagerName"`
	// Version of the configuration manager to use. Defaults to "11.4".
	ConfigurationManagerVersion pulumi.StringInput `pulumi:"configurationManagerVersion"`
	// When `useCustomCookbooks` is set, provide this sub-object as
	// described below.
	CustomCookbooksSources pulumi.ArrayInput `pulumi:"customCookbooksSources"`
	// Custom JSON attributes to apply to the entire stack.
	CustomJson pulumi.StringInput `pulumi:"customJson"`
	// Name of the availability zone where instances will be created
	// by default. This is required unless you set `vpcId`.
	DefaultAvailabilityZone pulumi.StringInput `pulumi:"defaultAvailabilityZone"`
	// The ARN of an IAM Instance Profile that created instances
	// will have by default.
	DefaultInstanceProfileArn pulumi.StringInput `pulumi:"defaultInstanceProfileArn"`
	// Name of OS that will be installed on instances by default.
	DefaultOs pulumi.StringInput `pulumi:"defaultOs"`
	// Name of the type of root device instances will have by default.
	DefaultRootDeviceType pulumi.StringInput `pulumi:"defaultRootDeviceType"`
	// Name of the SSH keypair that instances will have by default.
	DefaultSshKeyName pulumi.StringInput `pulumi:"defaultSshKeyName"`
	// Id of the subnet in which instances will be created by default. Mandatory
	// if `vpcId` is set, and forbidden if it isn't.
	DefaultSubnetId pulumi.StringInput `pulumi:"defaultSubnetId"`
	// Keyword representing the naming scheme that will be used for instance hostnames
	// within this stack.
	HostnameTheme pulumi.StringInput `pulumi:"hostnameTheme"`
	// Boolean value controlling whether Opsworks will run Berkshelf for this stack.
	ManageBerkshelf pulumi.BoolInput `pulumi:"manageBerkshelf"`
	// The name of the stack.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the region where the stack will exist.
	Region pulumi.StringInput `pulumi:"region"`
	// The ARN of an IAM role that the OpsWorks service will act as.
	ServiceRoleArn pulumi.StringInput `pulumi:"serviceRoleArn"`
	StackEndpoint pulumi.StringInput `pulumi:"stackEndpoint"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Boolean value controlling whether the custom cookbook settings are
	// enabled.
	UseCustomCookbooks pulumi.BoolInput `pulumi:"useCustomCookbooks"`
	// Boolean value controlling whether the standard OpsWorks
	// security groups apply to created instances.
	UseOpsworksSecurityGroups pulumi.BoolInput `pulumi:"useOpsworksSecurityGroups"`
	// The id of the VPC that this stack belongs to.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

// The set of arguments for constructing a Stack resource.
type StackArgs struct {
	// If set to `"LATEST"`, OpsWorks will automatically install the latest version.
	AgentVersion pulumi.StringInput `pulumi:"agentVersion"`
	// If `manageBerkshelf` is enabled, the version of Berkshelf to use.
	BerkshelfVersion pulumi.StringInput `pulumi:"berkshelfVersion"`
	// Color to paint next to the stack's resources in the OpsWorks console.
	Color pulumi.StringInput `pulumi:"color"`
	// Name of the configuration manager to use. Defaults to "Chef".
	ConfigurationManagerName pulumi.StringInput `pulumi:"configurationManagerName"`
	// Version of the configuration manager to use. Defaults to "11.4".
	ConfigurationManagerVersion pulumi.StringInput `pulumi:"configurationManagerVersion"`
	// When `useCustomCookbooks` is set, provide this sub-object as
	// described below.
	CustomCookbooksSources pulumi.ArrayInput `pulumi:"customCookbooksSources"`
	// Custom JSON attributes to apply to the entire stack.
	CustomJson pulumi.StringInput `pulumi:"customJson"`
	// Name of the availability zone where instances will be created
	// by default. This is required unless you set `vpcId`.
	DefaultAvailabilityZone pulumi.StringInput `pulumi:"defaultAvailabilityZone"`
	// The ARN of an IAM Instance Profile that created instances
	// will have by default.
	DefaultInstanceProfileArn pulumi.StringInput `pulumi:"defaultInstanceProfileArn"`
	// Name of OS that will be installed on instances by default.
	DefaultOs pulumi.StringInput `pulumi:"defaultOs"`
	// Name of the type of root device instances will have by default.
	DefaultRootDeviceType pulumi.StringInput `pulumi:"defaultRootDeviceType"`
	// Name of the SSH keypair that instances will have by default.
	DefaultSshKeyName pulumi.StringInput `pulumi:"defaultSshKeyName"`
	// Id of the subnet in which instances will be created by default. Mandatory
	// if `vpcId` is set, and forbidden if it isn't.
	DefaultSubnetId pulumi.StringInput `pulumi:"defaultSubnetId"`
	// Keyword representing the naming scheme that will be used for instance hostnames
	// within this stack.
	HostnameTheme pulumi.StringInput `pulumi:"hostnameTheme"`
	// Boolean value controlling whether Opsworks will run Berkshelf for this stack.
	ManageBerkshelf pulumi.BoolInput `pulumi:"manageBerkshelf"`
	// The name of the stack.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the region where the stack will exist.
	Region pulumi.StringInput `pulumi:"region"`
	// The ARN of an IAM role that the OpsWorks service will act as.
	ServiceRoleArn pulumi.StringInput `pulumi:"serviceRoleArn"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Boolean value controlling whether the custom cookbook settings are
	// enabled.
	UseCustomCookbooks pulumi.BoolInput `pulumi:"useCustomCookbooks"`
	// Boolean value controlling whether the standard OpsWorks
	// security groups apply to created instances.
	UseOpsworksSecurityGroups pulumi.BoolInput `pulumi:"useOpsworksSecurityGroups"`
	// The id of the VPC that this stack belongs to.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}
