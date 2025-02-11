// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get information about an RDS instance
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/db_instance.html.markdown.
func LookupInstance(ctx *pulumi.Context, args *GetInstanceArgs) (*GetInstanceResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["dbInstanceIdentifier"] = args.DbInstanceIdentifier
		inputs["tags"] = args.Tags
	}
	outputs, err := ctx.Invoke("aws:rds/getInstance:getInstance", inputs)
	if err != nil {
		return nil, err
	}
	return &GetInstanceResult{
		Address: outputs["address"],
		AllocatedStorage: outputs["allocatedStorage"],
		AutoMinorVersionUpgrade: outputs["autoMinorVersionUpgrade"],
		AvailabilityZone: outputs["availabilityZone"],
		BackupRetentionPeriod: outputs["backupRetentionPeriod"],
		CaCertIdentifier: outputs["caCertIdentifier"],
		DbClusterIdentifier: outputs["dbClusterIdentifier"],
		DbInstanceArn: outputs["dbInstanceArn"],
		DbInstanceClass: outputs["dbInstanceClass"],
		DbInstanceIdentifier: outputs["dbInstanceIdentifier"],
		DbInstancePort: outputs["dbInstancePort"],
		DbName: outputs["dbName"],
		DbParameterGroups: outputs["dbParameterGroups"],
		DbSecurityGroups: outputs["dbSecurityGroups"],
		DbSubnetGroup: outputs["dbSubnetGroup"],
		EnabledCloudwatchLogsExports: outputs["enabledCloudwatchLogsExports"],
		Endpoint: outputs["endpoint"],
		Engine: outputs["engine"],
		EngineVersion: outputs["engineVersion"],
		HostedZoneId: outputs["hostedZoneId"],
		Iops: outputs["iops"],
		KmsKeyId: outputs["kmsKeyId"],
		LicenseModel: outputs["licenseModel"],
		MasterUsername: outputs["masterUsername"],
		MonitoringInterval: outputs["monitoringInterval"],
		MonitoringRoleArn: outputs["monitoringRoleArn"],
		MultiAz: outputs["multiAz"],
		OptionGroupMemberships: outputs["optionGroupMemberships"],
		Port: outputs["port"],
		PreferredBackupWindow: outputs["preferredBackupWindow"],
		PreferredMaintenanceWindow: outputs["preferredMaintenanceWindow"],
		PubliclyAccessible: outputs["publiclyAccessible"],
		ReplicateSourceDb: outputs["replicateSourceDb"],
		ResourceId: outputs["resourceId"],
		StorageEncrypted: outputs["storageEncrypted"],
		StorageType: outputs["storageType"],
		Tags: outputs["tags"],
		Timezone: outputs["timezone"],
		VpcSecurityGroups: outputs["vpcSecurityGroups"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getInstance.
type GetInstanceArgs struct {
	// The name of the RDS instance
	DbInstanceIdentifier interface{}
	Tags interface{}
}

// A collection of values returned by getInstance.
type GetInstanceResult struct {
	// The hostname of the RDS instance. See also `endpoint` and `port`.
	Address interface{}
	// Specifies the allocated storage size specified in gigabytes.
	AllocatedStorage interface{}
	// Indicates that minor version patches are applied automatically.
	AutoMinorVersionUpgrade interface{}
	// Specifies the name of the Availability Zone the DB instance is located in.
	AvailabilityZone interface{}
	// Specifies the number of days for which automatic DB snapshots are retained.
	BackupRetentionPeriod interface{}
	// Specifies the identifier of the CA certificate for the DB instance.
	CaCertIdentifier interface{}
	// If the DB instance is a member of a DB cluster, contains the name of the DB cluster that the DB instance is a member of.
	DbClusterIdentifier interface{}
	// The Amazon Resource Name (ARN) for the DB instance.
	DbInstanceArn interface{}
	// Contains the name of the compute and memory capacity class of the DB instance.
	DbInstanceClass interface{}
	DbInstanceIdentifier interface{}
	// Specifies the port that the DB instance listens on.
	DbInstancePort interface{}
	// Contains the name of the initial database of this instance that was provided at create time, if one was specified when the DB instance was created. This same name is returned for the life of the DB instance.
	DbName interface{}
	// Provides the list of DB parameter groups applied to this DB instance.
	DbParameterGroups interface{}
	// Provides List of DB security groups associated to this DB instance.
	DbSecurityGroups interface{}
	// Specifies the name of the subnet group associated with the DB instance.
	DbSubnetGroup interface{}
	// List of log types to export to cloudwatch.
	EnabledCloudwatchLogsExports interface{}
	// The connection endpoint in `address:port` format.
	Endpoint interface{}
	// Provides the name of the database engine to be used for this DB instance.
	Engine interface{}
	// Indicates the database engine version.
	EngineVersion interface{}
	// The canonical hosted zone ID of the DB instance (to be used in a Route 53 Alias record).
	HostedZoneId interface{}
	// Specifies the Provisioned IOPS (I/O operations per second) value.
	Iops interface{}
	// If StorageEncrypted is true, the KMS key identifier for the encrypted DB instance.
	KmsKeyId interface{}
	// License model information for this DB instance.
	LicenseModel interface{}
	// Contains the master username for the DB instance.
	MasterUsername interface{}
	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.
	MonitoringInterval interface{}
	// The ARN for the IAM role that permits RDS to send Enhanced Monitoring metrics to CloudWatch Logs.
	MonitoringRoleArn interface{}
	// Specifies if the DB instance is a Multi-AZ deployment.
	MultiAz interface{}
	// Provides the list of option group memberships for this DB instance.
	OptionGroupMemberships interface{}
	// The database port.
	Port interface{}
	// Specifies the daily time range during which automated backups are created.
	PreferredBackupWindow interface{}
	// Specifies the weekly time range during which system maintenance can occur in UTC.
	PreferredMaintenanceWindow interface{}
	// Specifies the accessibility options for the DB instance.
	PubliclyAccessible interface{}
	// The identifier of the source DB that this is a replica of.
	ReplicateSourceDb interface{}
	// The RDS Resource ID of this instance.
	ResourceId interface{}
	// Specifies whether the DB instance is encrypted.
	StorageEncrypted interface{}
	// Specifies the storage type associated with DB instance.
	StorageType interface{}
	Tags interface{}
	// The time zone of the DB instance.
	Timezone interface{}
	// Provides a list of VPC security group elements that the DB instance belongs to.
	VpcSecurityGroups interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
