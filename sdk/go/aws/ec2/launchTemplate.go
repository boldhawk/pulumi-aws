// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an EC2 launch template resource. Can be used to create instances or auto scaling groups.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/launch_template.html.markdown.
type LaunchTemplate struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// Amazon Resource Name (ARN) of the launch template.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// Specify volumes to attach to the instance besides the volumes specified by the AMI.
	// See Block Devices below for details.
	BlockDeviceMappings pulumi.ArrayOutput `pulumi:"blockDeviceMappings"`

	// Targeting for EC2 capacity reservations. See Capacity Reservation Specification below for more details.
	CapacityReservationSpecification pulumi.AnyOutput `pulumi:"capacityReservationSpecification"`

	// Customize the credit specification of the instance. See Credit
	// Specification below for more details.
	CreditSpecification pulumi.AnyOutput `pulumi:"creditSpecification"`

	// The default version of the launch template.
	DefaultVersion pulumi.IntOutput `pulumi:"defaultVersion"`

	// Description of the launch template.
	Description pulumi.StringOutput `pulumi:"description"`

	// If `true`, enables [EC2 Instance
	// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
	DisableApiTermination pulumi.BoolOutput `pulumi:"disableApiTermination"`

	// If `true`, the launched EC2 instance will be EBS-optimized.
	EbsOptimized pulumi.StringOutput `pulumi:"ebsOptimized"`

	// The elastic GPU to attach to the instance. See Elastic GPU
	// below for more details.
	ElasticGpuSpecifications pulumi.ArrayOutput `pulumi:"elasticGpuSpecifications"`

	// Configuration block containing an Elastic Inference Accelerator to attach to the instance. See Elastic Inference Accelerator below for more details.
	ElasticInferenceAccelerator pulumi.AnyOutput `pulumi:"elasticInferenceAccelerator"`

	// The IAM Instance Profile to launch the instance with. See Instance Profile
	// below for more details.
	IamInstanceProfile pulumi.AnyOutput `pulumi:"iamInstanceProfile"`

	// The AMI from which to launch the instance.
	ImageId pulumi.StringOutput `pulumi:"imageId"`

	// Shutdown behavior for the instance. Can be `stop` or `terminate`.
	// (Default: `stop`).
	InstanceInitiatedShutdownBehavior pulumi.StringOutput `pulumi:"instanceInitiatedShutdownBehavior"`

	// The market (purchasing) option for the instance. See Market Options
	// below for details.
	InstanceMarketOptions pulumi.AnyOutput `pulumi:"instanceMarketOptions"`

	// The type of the instance.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`

	// The kernel ID.
	KernelId pulumi.StringOutput `pulumi:"kernelId"`

	// The key name to use for the instance.
	KeyName pulumi.StringOutput `pulumi:"keyName"`

	// The latest version of the launch template.
	LatestVersion pulumi.IntOutput `pulumi:"latestVersion"`

	// A list of license specifications to associate with. See License Specification below for more details.
	LicenseSpecifications pulumi.ArrayOutput `pulumi:"licenseSpecifications"`

	// The monitoring option for the instance. See Monitoring below for more details.
	Monitoring pulumi.AnyOutput `pulumi:"monitoring"`

	// The name of the launch template. If you leave this blank, this provider will auto-generate a unique name.
	Name pulumi.StringOutput `pulumi:"name"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`

	// Customize network interfaces to be attached at instance boot time. See Network
	// Interfaces below for more details.
	NetworkInterfaces pulumi.ArrayOutput `pulumi:"networkInterfaces"`

	// The placement of the instance. See Placement below for more details.
	Placement pulumi.AnyOutput `pulumi:"placement"`

	// The ID of the RAM disk.
	RamDiskId pulumi.StringOutput `pulumi:"ramDiskId"`

	// A list of security group names to associate with. If you are creating Instances in a VPC, use
	// `vpcSecurityGroupIds` instead.
	SecurityGroupNames pulumi.ArrayOutput `pulumi:"securityGroupNames"`

	// The tags to apply to the resources during launch. See Tag Specifications below for more details.
	TagSpecifications pulumi.ArrayOutput `pulumi:"tagSpecifications"`

	// A mapping of tags to assign to the launch template.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// The Base64-encoded user data to provide when launching the instance.
	UserData pulumi.StringOutput `pulumi:"userData"`

	// A list of security group IDs to associate with.
	VpcSecurityGroupIds pulumi.ArrayOutput `pulumi:"vpcSecurityGroupIds"`
}

// NewLaunchTemplate registers a new resource with the given unique name, arguments, and options.
func NewLaunchTemplate(ctx *pulumi.Context,
	name string, args *LaunchTemplateArgs, opts ...pulumi.ResourceOpt) (*LaunchTemplate, error) {
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["blockDeviceMappings"] = args.BlockDeviceMappings
		inputs["capacityReservationSpecification"] = args.CapacityReservationSpecification
		inputs["creditSpecification"] = args.CreditSpecification
		inputs["description"] = args.Description
		inputs["disableApiTermination"] = args.DisableApiTermination
		inputs["ebsOptimized"] = args.EbsOptimized
		inputs["elasticGpuSpecifications"] = args.ElasticGpuSpecifications
		inputs["elasticInferenceAccelerator"] = args.ElasticInferenceAccelerator
		inputs["iamInstanceProfile"] = args.IamInstanceProfile
		inputs["imageId"] = args.ImageId
		inputs["instanceInitiatedShutdownBehavior"] = args.InstanceInitiatedShutdownBehavior
		inputs["instanceMarketOptions"] = args.InstanceMarketOptions
		inputs["instanceType"] = args.InstanceType
		inputs["kernelId"] = args.KernelId
		inputs["keyName"] = args.KeyName
		inputs["licenseSpecifications"] = args.LicenseSpecifications
		inputs["monitoring"] = args.Monitoring
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["networkInterfaces"] = args.NetworkInterfaces
		inputs["placement"] = args.Placement
		inputs["ramDiskId"] = args.RamDiskId
		inputs["securityGroupNames"] = args.SecurityGroupNames
		inputs["tagSpecifications"] = args.TagSpecifications
		inputs["tags"] = args.Tags
		inputs["userData"] = args.UserData
		inputs["vpcSecurityGroupIds"] = args.VpcSecurityGroupIds
	}
	var resource LaunchTemplate
	err := ctx.RegisterResource("aws:ec2/launchTemplate:LaunchTemplate", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLaunchTemplate gets an existing LaunchTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLaunchTemplate(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LaunchTemplateState, opts ...pulumi.ResourceOpt) (*LaunchTemplate, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["blockDeviceMappings"] = state.BlockDeviceMappings
		inputs["capacityReservationSpecification"] = state.CapacityReservationSpecification
		inputs["creditSpecification"] = state.CreditSpecification
		inputs["defaultVersion"] = state.DefaultVersion
		inputs["description"] = state.Description
		inputs["disableApiTermination"] = state.DisableApiTermination
		inputs["ebsOptimized"] = state.EbsOptimized
		inputs["elasticGpuSpecifications"] = state.ElasticGpuSpecifications
		inputs["elasticInferenceAccelerator"] = state.ElasticInferenceAccelerator
		inputs["iamInstanceProfile"] = state.IamInstanceProfile
		inputs["imageId"] = state.ImageId
		inputs["instanceInitiatedShutdownBehavior"] = state.InstanceInitiatedShutdownBehavior
		inputs["instanceMarketOptions"] = state.InstanceMarketOptions
		inputs["instanceType"] = state.InstanceType
		inputs["kernelId"] = state.KernelId
		inputs["keyName"] = state.KeyName
		inputs["latestVersion"] = state.LatestVersion
		inputs["licenseSpecifications"] = state.LicenseSpecifications
		inputs["monitoring"] = state.Monitoring
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["networkInterfaces"] = state.NetworkInterfaces
		inputs["placement"] = state.Placement
		inputs["ramDiskId"] = state.RamDiskId
		inputs["securityGroupNames"] = state.SecurityGroupNames
		inputs["tagSpecifications"] = state.TagSpecifications
		inputs["tags"] = state.Tags
		inputs["userData"] = state.UserData
		inputs["vpcSecurityGroupIds"] = state.VpcSecurityGroupIds
	}
	var resource LaunchTemplate
	err := ctx.ReadResource("aws:ec2/launchTemplate:LaunchTemplate", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *LaunchTemplate) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *LaunchTemplate) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering LaunchTemplate resources.
type LaunchTemplateState struct {
	// Amazon Resource Name (ARN) of the launch template.
	Arn pulumi.StringInput `pulumi:"arn"`
	// Specify volumes to attach to the instance besides the volumes specified by the AMI.
	// See Block Devices below for details.
	BlockDeviceMappings pulumi.ArrayInput `pulumi:"blockDeviceMappings"`
	// Targeting for EC2 capacity reservations. See Capacity Reservation Specification below for more details.
	CapacityReservationSpecification pulumi.AnyInput `pulumi:"capacityReservationSpecification"`
	// Customize the credit specification of the instance. See Credit
	// Specification below for more details.
	CreditSpecification pulumi.AnyInput `pulumi:"creditSpecification"`
	// The default version of the launch template.
	DefaultVersion pulumi.IntInput `pulumi:"defaultVersion"`
	// Description of the launch template.
	Description pulumi.StringInput `pulumi:"description"`
	// If `true`, enables [EC2 Instance
	// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
	DisableApiTermination pulumi.BoolInput `pulumi:"disableApiTermination"`
	// If `true`, the launched EC2 instance will be EBS-optimized.
	EbsOptimized pulumi.StringInput `pulumi:"ebsOptimized"`
	// The elastic GPU to attach to the instance. See Elastic GPU
	// below for more details.
	ElasticGpuSpecifications pulumi.ArrayInput `pulumi:"elasticGpuSpecifications"`
	// Configuration block containing an Elastic Inference Accelerator to attach to the instance. See Elastic Inference Accelerator below for more details.
	ElasticInferenceAccelerator pulumi.AnyInput `pulumi:"elasticInferenceAccelerator"`
	// The IAM Instance Profile to launch the instance with. See Instance Profile
	// below for more details.
	IamInstanceProfile pulumi.AnyInput `pulumi:"iamInstanceProfile"`
	// The AMI from which to launch the instance.
	ImageId pulumi.StringInput `pulumi:"imageId"`
	// Shutdown behavior for the instance. Can be `stop` or `terminate`.
	// (Default: `stop`).
	InstanceInitiatedShutdownBehavior pulumi.StringInput `pulumi:"instanceInitiatedShutdownBehavior"`
	// The market (purchasing) option for the instance. See Market Options
	// below for details.
	InstanceMarketOptions pulumi.AnyInput `pulumi:"instanceMarketOptions"`
	// The type of the instance.
	InstanceType pulumi.StringInput `pulumi:"instanceType"`
	// The kernel ID.
	KernelId pulumi.StringInput `pulumi:"kernelId"`
	// The key name to use for the instance.
	KeyName pulumi.StringInput `pulumi:"keyName"`
	// The latest version of the launch template.
	LatestVersion pulumi.IntInput `pulumi:"latestVersion"`
	// A list of license specifications to associate with. See License Specification below for more details.
	LicenseSpecifications pulumi.ArrayInput `pulumi:"licenseSpecifications"`
	// The monitoring option for the instance. See Monitoring below for more details.
	Monitoring pulumi.AnyInput `pulumi:"monitoring"`
	// The name of the launch template. If you leave this blank, this provider will auto-generate a unique name.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// Customize network interfaces to be attached at instance boot time. See Network
	// Interfaces below for more details.
	NetworkInterfaces pulumi.ArrayInput `pulumi:"networkInterfaces"`
	// The placement of the instance. See Placement below for more details.
	Placement pulumi.AnyInput `pulumi:"placement"`
	// The ID of the RAM disk.
	RamDiskId pulumi.StringInput `pulumi:"ramDiskId"`
	// A list of security group names to associate with. If you are creating Instances in a VPC, use
	// `vpcSecurityGroupIds` instead.
	SecurityGroupNames pulumi.ArrayInput `pulumi:"securityGroupNames"`
	// The tags to apply to the resources during launch. See Tag Specifications below for more details.
	TagSpecifications pulumi.ArrayInput `pulumi:"tagSpecifications"`
	// A mapping of tags to assign to the launch template.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The Base64-encoded user data to provide when launching the instance.
	UserData pulumi.StringInput `pulumi:"userData"`
	// A list of security group IDs to associate with.
	VpcSecurityGroupIds pulumi.ArrayInput `pulumi:"vpcSecurityGroupIds"`
}

// The set of arguments for constructing a LaunchTemplate resource.
type LaunchTemplateArgs struct {
	// Specify volumes to attach to the instance besides the volumes specified by the AMI.
	// See Block Devices below for details.
	BlockDeviceMappings pulumi.ArrayInput `pulumi:"blockDeviceMappings"`
	// Targeting for EC2 capacity reservations. See Capacity Reservation Specification below for more details.
	CapacityReservationSpecification pulumi.AnyInput `pulumi:"capacityReservationSpecification"`
	// Customize the credit specification of the instance. See Credit
	// Specification below for more details.
	CreditSpecification pulumi.AnyInput `pulumi:"creditSpecification"`
	// Description of the launch template.
	Description pulumi.StringInput `pulumi:"description"`
	// If `true`, enables [EC2 Instance
	// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
	DisableApiTermination pulumi.BoolInput `pulumi:"disableApiTermination"`
	// If `true`, the launched EC2 instance will be EBS-optimized.
	EbsOptimized pulumi.StringInput `pulumi:"ebsOptimized"`
	// The elastic GPU to attach to the instance. See Elastic GPU
	// below for more details.
	ElasticGpuSpecifications pulumi.ArrayInput `pulumi:"elasticGpuSpecifications"`
	// Configuration block containing an Elastic Inference Accelerator to attach to the instance. See Elastic Inference Accelerator below for more details.
	ElasticInferenceAccelerator pulumi.AnyInput `pulumi:"elasticInferenceAccelerator"`
	// The IAM Instance Profile to launch the instance with. See Instance Profile
	// below for more details.
	IamInstanceProfile pulumi.AnyInput `pulumi:"iamInstanceProfile"`
	// The AMI from which to launch the instance.
	ImageId pulumi.StringInput `pulumi:"imageId"`
	// Shutdown behavior for the instance. Can be `stop` or `terminate`.
	// (Default: `stop`).
	InstanceInitiatedShutdownBehavior pulumi.StringInput `pulumi:"instanceInitiatedShutdownBehavior"`
	// The market (purchasing) option for the instance. See Market Options
	// below for details.
	InstanceMarketOptions pulumi.AnyInput `pulumi:"instanceMarketOptions"`
	// The type of the instance.
	InstanceType pulumi.StringInput `pulumi:"instanceType"`
	// The kernel ID.
	KernelId pulumi.StringInput `pulumi:"kernelId"`
	// The key name to use for the instance.
	KeyName pulumi.StringInput `pulumi:"keyName"`
	// A list of license specifications to associate with. See License Specification below for more details.
	LicenseSpecifications pulumi.ArrayInput `pulumi:"licenseSpecifications"`
	// The monitoring option for the instance. See Monitoring below for more details.
	Monitoring pulumi.AnyInput `pulumi:"monitoring"`
	// The name of the launch template. If you leave this blank, this provider will auto-generate a unique name.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// Customize network interfaces to be attached at instance boot time. See Network
	// Interfaces below for more details.
	NetworkInterfaces pulumi.ArrayInput `pulumi:"networkInterfaces"`
	// The placement of the instance. See Placement below for more details.
	Placement pulumi.AnyInput `pulumi:"placement"`
	// The ID of the RAM disk.
	RamDiskId pulumi.StringInput `pulumi:"ramDiskId"`
	// A list of security group names to associate with. If you are creating Instances in a VPC, use
	// `vpcSecurityGroupIds` instead.
	SecurityGroupNames pulumi.ArrayInput `pulumi:"securityGroupNames"`
	// The tags to apply to the resources during launch. See Tag Specifications below for more details.
	TagSpecifications pulumi.ArrayInput `pulumi:"tagSpecifications"`
	// A mapping of tags to assign to the launch template.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The Base64-encoded user data to provide when launching the instance.
	UserData pulumi.StringInput `pulumi:"userData"`
	// A list of security group IDs to associate with.
	VpcSecurityGroupIds pulumi.ArrayInput `pulumi:"vpcSecurityGroupIds"`
}
