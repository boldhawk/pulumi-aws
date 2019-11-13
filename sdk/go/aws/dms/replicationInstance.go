// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a DMS (Data Migration Service) replication instance resource. DMS replication instances can be created, updated, deleted, and imported.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dms_replication_instance.html.markdown.
type ReplicationInstance struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	AllocatedStorage pulumi.IntOutput `pulumi:"allocatedStorage"`

	// Indicates whether the changes should be applied immediately or during the next maintenance window. Only used when updating an existing resource.
	ApplyImmediately pulumi.BoolOutput `pulumi:"applyImmediately"`

	// Indicates that minor engine upgrades will be applied automatically to the replication instance during the maintenance window.
	AutoMinorVersionUpgrade pulumi.BoolOutput `pulumi:"autoMinorVersionUpgrade"`

	// The EC2 Availability Zone that the replication instance will be created in.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`

	// The engine version number of the replication instance.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`

	// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kmsKeyArn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn pulumi.StringOutput `pulumi:"kmsKeyArn"`

	// Specifies if the replication instance is a multi-az deployment. You cannot set the `availabilityZone` parameter if the `multiAz` parameter is set to `true`.
	MultiAz pulumi.BoolOutput `pulumi:"multiAz"`

	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	PreferredMaintenanceWindow pulumi.StringOutput `pulumi:"preferredMaintenanceWindow"`

	// Specifies the accessibility options for the replication instance. A value of true represents an instance with a public IP address. A value of false represents an instance with a private IP address.
	PubliclyAccessible pulumi.BoolOutput `pulumi:"publiclyAccessible"`

	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn pulumi.StringOutput `pulumi:"replicationInstanceArn"`

	// The compute and memory capacity of the replication instance as specified by the replication instance class. Can be one of `dms.t2.micro | dms.t2.small | dms.t2.medium | dms.t2.large | dms.c4.large | dms.c4.xlarge | dms.c4.2xlarge | dms.c4.4xlarge`
	ReplicationInstanceClass pulumi.StringOutput `pulumi:"replicationInstanceClass"`

	// The replication instance identifier. This parameter is stored as a lowercase string.
	ReplicationInstanceId pulumi.StringOutput `pulumi:"replicationInstanceId"`

	// A list of the private IP addresses of the replication instance.
	ReplicationInstancePrivateIps pulumi.ArrayOutput `pulumi:"replicationInstancePrivateIps"`

	// A list of the public IP addresses of the replication instance.
	ReplicationInstancePublicIps pulumi.ArrayOutput `pulumi:"replicationInstancePublicIps"`

	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupId pulumi.StringOutput `pulumi:"replicationSubnetGroupId"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// A list of VPC security group IDs to be used with the replication instance. The VPC security groups must work with the VPC containing the replication instance.
	VpcSecurityGroupIds pulumi.ArrayOutput `pulumi:"vpcSecurityGroupIds"`
}

// NewReplicationInstance registers a new resource with the given unique name, arguments, and options.
func NewReplicationInstance(ctx *pulumi.Context,
	name string, args *ReplicationInstanceArgs, opts ...pulumi.ResourceOpt) (*ReplicationInstance, error) {
	if args == nil || args.ReplicationInstanceClass == nil {
		return nil, errors.New("missing required argument 'ReplicationInstanceClass'")
	}
	if args == nil || args.ReplicationInstanceId == nil {
		return nil, errors.New("missing required argument 'ReplicationInstanceId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["allocatedStorage"] = args.AllocatedStorage
		inputs["applyImmediately"] = args.ApplyImmediately
		inputs["autoMinorVersionUpgrade"] = args.AutoMinorVersionUpgrade
		inputs["availabilityZone"] = args.AvailabilityZone
		inputs["engineVersion"] = args.EngineVersion
		inputs["kmsKeyArn"] = args.KmsKeyArn
		inputs["multiAz"] = args.MultiAz
		inputs["preferredMaintenanceWindow"] = args.PreferredMaintenanceWindow
		inputs["publiclyAccessible"] = args.PubliclyAccessible
		inputs["replicationInstanceClass"] = args.ReplicationInstanceClass
		inputs["replicationInstanceId"] = args.ReplicationInstanceId
		inputs["replicationSubnetGroupId"] = args.ReplicationSubnetGroupId
		inputs["tags"] = args.Tags
		inputs["vpcSecurityGroupIds"] = args.VpcSecurityGroupIds
	}
	var resource ReplicationInstance
	err := ctx.RegisterResource("aws:dms/replicationInstance:ReplicationInstance", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationInstance gets an existing ReplicationInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationInstance(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ReplicationInstanceState, opts ...pulumi.ResourceOpt) (*ReplicationInstance, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["allocatedStorage"] = state.AllocatedStorage
		inputs["applyImmediately"] = state.ApplyImmediately
		inputs["autoMinorVersionUpgrade"] = state.AutoMinorVersionUpgrade
		inputs["availabilityZone"] = state.AvailabilityZone
		inputs["engineVersion"] = state.EngineVersion
		inputs["kmsKeyArn"] = state.KmsKeyArn
		inputs["multiAz"] = state.MultiAz
		inputs["preferredMaintenanceWindow"] = state.PreferredMaintenanceWindow
		inputs["publiclyAccessible"] = state.PubliclyAccessible
		inputs["replicationInstanceArn"] = state.ReplicationInstanceArn
		inputs["replicationInstanceClass"] = state.ReplicationInstanceClass
		inputs["replicationInstanceId"] = state.ReplicationInstanceId
		inputs["replicationInstancePrivateIps"] = state.ReplicationInstancePrivateIps
		inputs["replicationInstancePublicIps"] = state.ReplicationInstancePublicIps
		inputs["replicationSubnetGroupId"] = state.ReplicationSubnetGroupId
		inputs["tags"] = state.Tags
		inputs["vpcSecurityGroupIds"] = state.VpcSecurityGroupIds
	}
	var resource ReplicationInstance
	err := ctx.ReadResource("aws:dms/replicationInstance:ReplicationInstance", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *ReplicationInstance) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *ReplicationInstance) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering ReplicationInstance resources.
type ReplicationInstanceState struct {
	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	AllocatedStorage pulumi.IntInput `pulumi:"allocatedStorage"`
	// Indicates whether the changes should be applied immediately or during the next maintenance window. Only used when updating an existing resource.
	ApplyImmediately pulumi.BoolInput `pulumi:"applyImmediately"`
	// Indicates that minor engine upgrades will be applied automatically to the replication instance during the maintenance window.
	AutoMinorVersionUpgrade pulumi.BoolInput `pulumi:"autoMinorVersionUpgrade"`
	// The EC2 Availability Zone that the replication instance will be created in.
	AvailabilityZone pulumi.StringInput `pulumi:"availabilityZone"`
	// The engine version number of the replication instance.
	EngineVersion pulumi.StringInput `pulumi:"engineVersion"`
	// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kmsKeyArn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn pulumi.StringInput `pulumi:"kmsKeyArn"`
	// Specifies if the replication instance is a multi-az deployment. You cannot set the `availabilityZone` parameter if the `multiAz` parameter is set to `true`.
	MultiAz pulumi.BoolInput `pulumi:"multiAz"`
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	PreferredMaintenanceWindow pulumi.StringInput `pulumi:"preferredMaintenanceWindow"`
	// Specifies the accessibility options for the replication instance. A value of true represents an instance with a public IP address. A value of false represents an instance with a private IP address.
	PubliclyAccessible pulumi.BoolInput `pulumi:"publiclyAccessible"`
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn pulumi.StringInput `pulumi:"replicationInstanceArn"`
	// The compute and memory capacity of the replication instance as specified by the replication instance class. Can be one of `dms.t2.micro | dms.t2.small | dms.t2.medium | dms.t2.large | dms.c4.large | dms.c4.xlarge | dms.c4.2xlarge | dms.c4.4xlarge`
	ReplicationInstanceClass pulumi.StringInput `pulumi:"replicationInstanceClass"`
	// The replication instance identifier. This parameter is stored as a lowercase string.
	ReplicationInstanceId pulumi.StringInput `pulumi:"replicationInstanceId"`
	// A list of the private IP addresses of the replication instance.
	ReplicationInstancePrivateIps pulumi.ArrayInput `pulumi:"replicationInstancePrivateIps"`
	// A list of the public IP addresses of the replication instance.
	ReplicationInstancePublicIps pulumi.ArrayInput `pulumi:"replicationInstancePublicIps"`
	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupId pulumi.StringInput `pulumi:"replicationSubnetGroupId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// A list of VPC security group IDs to be used with the replication instance. The VPC security groups must work with the VPC containing the replication instance.
	VpcSecurityGroupIds pulumi.ArrayInput `pulumi:"vpcSecurityGroupIds"`
}

// The set of arguments for constructing a ReplicationInstance resource.
type ReplicationInstanceArgs struct {
	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	AllocatedStorage pulumi.IntInput `pulumi:"allocatedStorage"`
	// Indicates whether the changes should be applied immediately or during the next maintenance window. Only used when updating an existing resource.
	ApplyImmediately pulumi.BoolInput `pulumi:"applyImmediately"`
	// Indicates that minor engine upgrades will be applied automatically to the replication instance during the maintenance window.
	AutoMinorVersionUpgrade pulumi.BoolInput `pulumi:"autoMinorVersionUpgrade"`
	// The EC2 Availability Zone that the replication instance will be created in.
	AvailabilityZone pulumi.StringInput `pulumi:"availabilityZone"`
	// The engine version number of the replication instance.
	EngineVersion pulumi.StringInput `pulumi:"engineVersion"`
	// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kmsKeyArn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn pulumi.StringInput `pulumi:"kmsKeyArn"`
	// Specifies if the replication instance is a multi-az deployment. You cannot set the `availabilityZone` parameter if the `multiAz` parameter is set to `true`.
	MultiAz pulumi.BoolInput `pulumi:"multiAz"`
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	PreferredMaintenanceWindow pulumi.StringInput `pulumi:"preferredMaintenanceWindow"`
	// Specifies the accessibility options for the replication instance. A value of true represents an instance with a public IP address. A value of false represents an instance with a private IP address.
	PubliclyAccessible pulumi.BoolInput `pulumi:"publiclyAccessible"`
	// The compute and memory capacity of the replication instance as specified by the replication instance class. Can be one of `dms.t2.micro | dms.t2.small | dms.t2.medium | dms.t2.large | dms.c4.large | dms.c4.xlarge | dms.c4.2xlarge | dms.c4.4xlarge`
	ReplicationInstanceClass pulumi.StringInput `pulumi:"replicationInstanceClass"`
	// The replication instance identifier. This parameter is stored as a lowercase string.
	ReplicationInstanceId pulumi.StringInput `pulumi:"replicationInstanceId"`
	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupId pulumi.StringInput `pulumi:"replicationSubnetGroupId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// A list of VPC security group IDs to be used with the replication instance. The VPC security groups must work with the VPC containing the replication instance.
	VpcSecurityGroupIds pulumi.ArrayInput `pulumi:"vpcSecurityGroupIds"`
}
