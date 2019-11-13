// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a [RDS Aurora Cluster][2]. To manage cluster instances that inherit configuration from the cluster (when not running the cluster in `serverless` engine mode), see the [`rds.ClusterInstance` resource](https://www.terraform.io/docs/providers/aws/r/rds_cluster_instance.html). To manage non-Aurora databases (e.g. MySQL, PostgreSQL, SQL Server, etc.), see the [`rds.Instance` resource](https://www.terraform.io/docs/providers/aws/r/db_instance.html).
// 
// For information on the difference between the available Aurora MySQL engines
// see [Comparison between Aurora MySQL 1 and Aurora MySQL 2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Updates.20180206.html)
// in the Amazon RDS User Guide.
// 
// Changes to a RDS Cluster can occur when you manually change a
// parameter, such as `port`, and are reflected in the next maintenance
// window. Because of this, this provider may report a difference in its planning
// phase because a modification has not yet taken place. You can use the
// `applyImmediately` flag to instruct the service to apply the change immediately
// (see documentation below).
// 
// > **Note:** using `applyImmediately` can result in a
// brief downtime as the server reboots. See the AWS Docs on [RDS Maintenance][4]
// for more information.
// 
// > **Note:** All arguments including the username and password will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/rds_cluster.html.markdown.
type Cluster struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// Specifies whether any cluster modifications
	// are applied immediately, or during the next maintenance window. Default is
	// `false`. See [Amazon RDS Documentation for more information.](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.DBInstance.Modifying.html)
	ApplyImmediately pulumi.BoolOutput `pulumi:"applyImmediately"`

	// Amazon Resource Name (ARN) of cluster
	Arn pulumi.StringOutput `pulumi:"arn"`

	// A list of EC2 Availability Zones for the DB cluster storage where DB cluster instances can be created. RDS automatically assigns 3 AZs if less than 3 AZs are configured, which will show as a difference requiring resource recreation next deployment. It is recommended to specify 3 AZs or use `ignoreChanges` if necessary.
	AvailabilityZones pulumi.ArrayOutput `pulumi:"availabilityZones"`

	// The target backtrack window, in seconds. Only available for `aurora` engine currently. To disable backtracking, set this value to `0`. Defaults to `0`. Must be between `0` and `259200` (72 hours)
	BacktrackWindow pulumi.IntOutput `pulumi:"backtrackWindow"`

	// The days to retain backups for. Default `1`
	BackupRetentionPeriod pulumi.IntOutput `pulumi:"backupRetentionPeriod"`

	// The cluster identifier. If omitted, this provider will assign a random, unique identifier.
	ClusterIdentifier pulumi.StringOutput `pulumi:"clusterIdentifier"`

	// Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `clusterIdentifier`.
	ClusterIdentifierPrefix pulumi.StringOutput `pulumi:"clusterIdentifierPrefix"`

	// List of RDS Instances that are a part of this cluster
	ClusterMembers pulumi.ArrayOutput `pulumi:"clusterMembers"`

	// The RDS Cluster Resource ID
	ClusterResourceId pulumi.StringOutput `pulumi:"clusterResourceId"`

	// Copy all Cluster `tags` to snapshots. Default is `false`.
	CopyTagsToSnapshot pulumi.BoolOutput `pulumi:"copyTagsToSnapshot"`

	// Name for an automatically created database on cluster creation. There are different naming restrictions per database engine: [RDS Naming Constraints][5]
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`

	// A cluster parameter group to associate with the cluster.
	DbClusterParameterGroupName pulumi.StringOutput `pulumi:"dbClusterParameterGroupName"`

	// A DB subnet group to associate with this DB instance. **NOTE:** This must match the `dbSubnetGroupName` specified on every [`rds.ClusterInstance`](https://www.terraform.io/docs/providers/aws/r/rds_cluster_instance.html) in the cluster.
	DbSubnetGroupName pulumi.StringOutput `pulumi:"dbSubnetGroupName"`

	// If the DB instance should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`

	// List of log types to export to cloudwatch. If omitted, no logs will be exported.
	// The following log types are supported: `audit`, `error`, `general`, `slowquery`, `postgresql` (PostgreSQL).
	EnabledCloudwatchLogsExports pulumi.ArrayOutput `pulumi:"enabledCloudwatchLogsExports"`

	// The DNS address of the RDS instance
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`

	// The name of the database engine to be used for this DB cluster. Defaults to `aurora`. Valid Values: `aurora`, `aurora-mysql`, `aurora-postgresql`
	Engine pulumi.StringOutput `pulumi:"engine"`

	// The database engine mode. Valid values: `global`, `multimaster`, `parallelquery`, `provisioned`, `serverless`. Defaults to: `provisioned`. See the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/aurora-serverless.html) for limitations when using `serverless`.
	EngineMode pulumi.StringOutput `pulumi:"engineMode"`

	// The database engine version. Updating this argument results in an outage. See the [Aurora MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Updates.html) and [Aurora Postgres](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Updates.html) documentation for your configured engine to determine this value. For example with Aurora MySQL 2, a potential value for this argument is `5.7.mysql_aurora.2.03.2`.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`

	// The name of your final DB snapshot
	// when this DB cluster is deleted. If omitted, no final snapshot will be
	// made.
	FinalSnapshotIdentifier pulumi.StringOutput `pulumi:"finalSnapshotIdentifier"`

	// The global cluster identifier specified on [`rds.GlobalCluster`](https://www.terraform.io/docs/providers/aws/r/rds_global_cluster.html).
	GlobalClusterIdentifier pulumi.StringOutput `pulumi:"globalClusterIdentifier"`

	// The Route53 Hosted Zone ID of the endpoint
	HostedZoneId pulumi.StringOutput `pulumi:"hostedZoneId"`

	// Specifies whether or mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled. Please see [AWS Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html) for availability and limitations.
	IamDatabaseAuthenticationEnabled pulumi.BoolOutput `pulumi:"iamDatabaseAuthenticationEnabled"`

	// A List of ARNs for the IAM roles to associate to the RDS Cluster.
	IamRoles pulumi.ArrayOutput `pulumi:"iamRoles"`

	// The ARN for the KMS encryption key. When specifying `kmsKeyId`, `storageEncrypted` needs to be set to true.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`

	// Password for the master DB user. Note that this may
	// show up in logs, and it will be stored in the state file. Please refer to the [RDS Naming Constraints][5]
	MasterPassword pulumi.StringOutput `pulumi:"masterPassword"`

	// Username for the master DB user. Please refer to the [RDS Naming Constraints][5]. This argument does not support in-place updates and cannot be changed during a restore from snapshot.
	MasterUsername pulumi.StringOutput `pulumi:"masterUsername"`

	// The port on which the DB accepts connections
	Port pulumi.IntOutput `pulumi:"port"`

	// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter.Time in UTC
	// Default: A 30-minute window selected at random from an 8-hour block of time per region. e.g. 04:00-09:00
	PreferredBackupWindow pulumi.StringOutput `pulumi:"preferredBackupWindow"`

	// The weekly time range during which system maintenance can occur, in (UTC) e.g. wed:04:00-wed:04:30
	PreferredMaintenanceWindow pulumi.StringOutput `pulumi:"preferredMaintenanceWindow"`

	// A read-only endpoint for the Aurora cluster, automatically
	// load-balanced across replicas
	ReaderEndpoint pulumi.StringOutput `pulumi:"readerEndpoint"`

	// ARN of a source DB cluster or DB instance if this DB cluster is to be created as a Read Replica.
	ReplicationSourceIdentifier pulumi.StringOutput `pulumi:"replicationSourceIdentifier"`

	S3Import pulumi.AnyOutput `pulumi:"s3Import"`

	// Nested attribute with scaling properties. Only valid when `engineMode` is set to `serverless`. More details below.
	ScalingConfiguration pulumi.AnyOutput `pulumi:"scalingConfiguration"`

	// Determines whether a final DB snapshot is created before the DB cluster is deleted. If true is specified, no DB snapshot is created. If false is specified, a DB snapshot is created before the DB cluster is deleted, using the value from `finalSnapshotIdentifier`. Default is `false`.
	SkipFinalSnapshot pulumi.BoolOutput `pulumi:"skipFinalSnapshot"`

	// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a DB cluster snapshot, or the ARN when specifying a DB snapshot.
	SnapshotIdentifier pulumi.StringOutput `pulumi:"snapshotIdentifier"`

	// The source region for an encrypted replica DB cluster.
	SourceRegion pulumi.StringOutput `pulumi:"sourceRegion"`

	// Specifies whether the DB cluster is encrypted. The default is `false` for `provisioned` `engineMode` and `true` for `serverless` `engineMode`.
	StorageEncrypted pulumi.BoolOutput `pulumi:"storageEncrypted"`

	// A mapping of tags to assign to the DB cluster.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// List of VPC security groups to associate
	// with the Cluster
	VpcSecurityGroupIds pulumi.ArrayOutput `pulumi:"vpcSecurityGroupIds"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOpt) (*Cluster, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["applyImmediately"] = args.ApplyImmediately
		inputs["availabilityZones"] = args.AvailabilityZones
		inputs["backtrackWindow"] = args.BacktrackWindow
		inputs["backupRetentionPeriod"] = args.BackupRetentionPeriod
		inputs["clusterIdentifier"] = args.ClusterIdentifier
		inputs["clusterIdentifierPrefix"] = args.ClusterIdentifierPrefix
		inputs["clusterMembers"] = args.ClusterMembers
		inputs["copyTagsToSnapshot"] = args.CopyTagsToSnapshot
		inputs["databaseName"] = args.DatabaseName
		inputs["dbClusterParameterGroupName"] = args.DbClusterParameterGroupName
		inputs["dbSubnetGroupName"] = args.DbSubnetGroupName
		inputs["deletionProtection"] = args.DeletionProtection
		inputs["enabledCloudwatchLogsExports"] = args.EnabledCloudwatchLogsExports
		inputs["engine"] = args.Engine
		inputs["engineMode"] = args.EngineMode
		inputs["engineVersion"] = args.EngineVersion
		inputs["finalSnapshotIdentifier"] = args.FinalSnapshotIdentifier
		inputs["globalClusterIdentifier"] = args.GlobalClusterIdentifier
		inputs["iamDatabaseAuthenticationEnabled"] = args.IamDatabaseAuthenticationEnabled
		inputs["iamRoles"] = args.IamRoles
		inputs["kmsKeyId"] = args.KmsKeyId
		inputs["masterPassword"] = args.MasterPassword
		inputs["masterUsername"] = args.MasterUsername
		inputs["port"] = args.Port
		inputs["preferredBackupWindow"] = args.PreferredBackupWindow
		inputs["preferredMaintenanceWindow"] = args.PreferredMaintenanceWindow
		inputs["replicationSourceIdentifier"] = args.ReplicationSourceIdentifier
		inputs["s3Import"] = args.S3Import
		inputs["scalingConfiguration"] = args.ScalingConfiguration
		inputs["skipFinalSnapshot"] = args.SkipFinalSnapshot
		inputs["snapshotIdentifier"] = args.SnapshotIdentifier
		inputs["sourceRegion"] = args.SourceRegion
		inputs["storageEncrypted"] = args.StorageEncrypted
		inputs["tags"] = args.Tags
		inputs["vpcSecurityGroupIds"] = args.VpcSecurityGroupIds
	}
	var resource Cluster
	err := ctx.RegisterResource("aws:rds/cluster:Cluster", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ClusterState, opts ...pulumi.ResourceOpt) (*Cluster, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["applyImmediately"] = state.ApplyImmediately
		inputs["arn"] = state.Arn
		inputs["availabilityZones"] = state.AvailabilityZones
		inputs["backtrackWindow"] = state.BacktrackWindow
		inputs["backupRetentionPeriod"] = state.BackupRetentionPeriod
		inputs["clusterIdentifier"] = state.ClusterIdentifier
		inputs["clusterIdentifierPrefix"] = state.ClusterIdentifierPrefix
		inputs["clusterMembers"] = state.ClusterMembers
		inputs["clusterResourceId"] = state.ClusterResourceId
		inputs["copyTagsToSnapshot"] = state.CopyTagsToSnapshot
		inputs["databaseName"] = state.DatabaseName
		inputs["dbClusterParameterGroupName"] = state.DbClusterParameterGroupName
		inputs["dbSubnetGroupName"] = state.DbSubnetGroupName
		inputs["deletionProtection"] = state.DeletionProtection
		inputs["enabledCloudwatchLogsExports"] = state.EnabledCloudwatchLogsExports
		inputs["endpoint"] = state.Endpoint
		inputs["engine"] = state.Engine
		inputs["engineMode"] = state.EngineMode
		inputs["engineVersion"] = state.EngineVersion
		inputs["finalSnapshotIdentifier"] = state.FinalSnapshotIdentifier
		inputs["globalClusterIdentifier"] = state.GlobalClusterIdentifier
		inputs["hostedZoneId"] = state.HostedZoneId
		inputs["iamDatabaseAuthenticationEnabled"] = state.IamDatabaseAuthenticationEnabled
		inputs["iamRoles"] = state.IamRoles
		inputs["kmsKeyId"] = state.KmsKeyId
		inputs["masterPassword"] = state.MasterPassword
		inputs["masterUsername"] = state.MasterUsername
		inputs["port"] = state.Port
		inputs["preferredBackupWindow"] = state.PreferredBackupWindow
		inputs["preferredMaintenanceWindow"] = state.PreferredMaintenanceWindow
		inputs["readerEndpoint"] = state.ReaderEndpoint
		inputs["replicationSourceIdentifier"] = state.ReplicationSourceIdentifier
		inputs["s3Import"] = state.S3Import
		inputs["scalingConfiguration"] = state.ScalingConfiguration
		inputs["skipFinalSnapshot"] = state.SkipFinalSnapshot
		inputs["snapshotIdentifier"] = state.SnapshotIdentifier
		inputs["sourceRegion"] = state.SourceRegion
		inputs["storageEncrypted"] = state.StorageEncrypted
		inputs["tags"] = state.Tags
		inputs["vpcSecurityGroupIds"] = state.VpcSecurityGroupIds
	}
	var resource Cluster
	err := ctx.ReadResource("aws:rds/cluster:Cluster", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *Cluster) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *Cluster) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering Cluster resources.
type ClusterState struct {
	// Specifies whether any cluster modifications
	// are applied immediately, or during the next maintenance window. Default is
	// `false`. See [Amazon RDS Documentation for more information.](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.DBInstance.Modifying.html)
	ApplyImmediately pulumi.BoolInput `pulumi:"applyImmediately"`
	// Amazon Resource Name (ARN) of cluster
	Arn pulumi.StringInput `pulumi:"arn"`
	// A list of EC2 Availability Zones for the DB cluster storage where DB cluster instances can be created. RDS automatically assigns 3 AZs if less than 3 AZs are configured, which will show as a difference requiring resource recreation next deployment. It is recommended to specify 3 AZs or use `ignoreChanges` if necessary.
	AvailabilityZones pulumi.ArrayInput `pulumi:"availabilityZones"`
	// The target backtrack window, in seconds. Only available for `aurora` engine currently. To disable backtracking, set this value to `0`. Defaults to `0`. Must be between `0` and `259200` (72 hours)
	BacktrackWindow pulumi.IntInput `pulumi:"backtrackWindow"`
	// The days to retain backups for. Default `1`
	BackupRetentionPeriod pulumi.IntInput `pulumi:"backupRetentionPeriod"`
	// The cluster identifier. If omitted, this provider will assign a random, unique identifier.
	ClusterIdentifier pulumi.StringInput `pulumi:"clusterIdentifier"`
	// Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `clusterIdentifier`.
	ClusterIdentifierPrefix pulumi.StringInput `pulumi:"clusterIdentifierPrefix"`
	// List of RDS Instances that are a part of this cluster
	ClusterMembers pulumi.ArrayInput `pulumi:"clusterMembers"`
	// The RDS Cluster Resource ID
	ClusterResourceId pulumi.StringInput `pulumi:"clusterResourceId"`
	// Copy all Cluster `tags` to snapshots. Default is `false`.
	CopyTagsToSnapshot pulumi.BoolInput `pulumi:"copyTagsToSnapshot"`
	// Name for an automatically created database on cluster creation. There are different naming restrictions per database engine: [RDS Naming Constraints][5]
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// A cluster parameter group to associate with the cluster.
	DbClusterParameterGroupName pulumi.StringInput `pulumi:"dbClusterParameterGroupName"`
	// A DB subnet group to associate with this DB instance. **NOTE:** This must match the `dbSubnetGroupName` specified on every [`rds.ClusterInstance`](https://www.terraform.io/docs/providers/aws/r/rds_cluster_instance.html) in the cluster.
	DbSubnetGroupName pulumi.StringInput `pulumi:"dbSubnetGroupName"`
	// If the DB instance should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection pulumi.BoolInput `pulumi:"deletionProtection"`
	// List of log types to export to cloudwatch. If omitted, no logs will be exported.
	// The following log types are supported: `audit`, `error`, `general`, `slowquery`, `postgresql` (PostgreSQL).
	EnabledCloudwatchLogsExports pulumi.ArrayInput `pulumi:"enabledCloudwatchLogsExports"`
	// The DNS address of the RDS instance
	Endpoint pulumi.StringInput `pulumi:"endpoint"`
	// The name of the database engine to be used for this DB cluster. Defaults to `aurora`. Valid Values: `aurora`, `aurora-mysql`, `aurora-postgresql`
	Engine pulumi.StringInput `pulumi:"engine"`
	// The database engine mode. Valid values: `global`, `multimaster`, `parallelquery`, `provisioned`, `serverless`. Defaults to: `provisioned`. See the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/aurora-serverless.html) for limitations when using `serverless`.
	EngineMode pulumi.StringInput `pulumi:"engineMode"`
	// The database engine version. Updating this argument results in an outage. See the [Aurora MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Updates.html) and [Aurora Postgres](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Updates.html) documentation for your configured engine to determine this value. For example with Aurora MySQL 2, a potential value for this argument is `5.7.mysql_aurora.2.03.2`.
	EngineVersion pulumi.StringInput `pulumi:"engineVersion"`
	// The name of your final DB snapshot
	// when this DB cluster is deleted. If omitted, no final snapshot will be
	// made.
	FinalSnapshotIdentifier pulumi.StringInput `pulumi:"finalSnapshotIdentifier"`
	// The global cluster identifier specified on [`rds.GlobalCluster`](https://www.terraform.io/docs/providers/aws/r/rds_global_cluster.html).
	GlobalClusterIdentifier pulumi.StringInput `pulumi:"globalClusterIdentifier"`
	// The Route53 Hosted Zone ID of the endpoint
	HostedZoneId pulumi.StringInput `pulumi:"hostedZoneId"`
	// Specifies whether or mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled. Please see [AWS Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html) for availability and limitations.
	IamDatabaseAuthenticationEnabled pulumi.BoolInput `pulumi:"iamDatabaseAuthenticationEnabled"`
	// A List of ARNs for the IAM roles to associate to the RDS Cluster.
	IamRoles pulumi.ArrayInput `pulumi:"iamRoles"`
	// The ARN for the KMS encryption key. When specifying `kmsKeyId`, `storageEncrypted` needs to be set to true.
	KmsKeyId pulumi.StringInput `pulumi:"kmsKeyId"`
	// Password for the master DB user. Note that this may
	// show up in logs, and it will be stored in the state file. Please refer to the [RDS Naming Constraints][5]
	MasterPassword pulumi.StringInput `pulumi:"masterPassword"`
	// Username for the master DB user. Please refer to the [RDS Naming Constraints][5]. This argument does not support in-place updates and cannot be changed during a restore from snapshot.
	MasterUsername pulumi.StringInput `pulumi:"masterUsername"`
	// The port on which the DB accepts connections
	Port pulumi.IntInput `pulumi:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter.Time in UTC
	// Default: A 30-minute window selected at random from an 8-hour block of time per region. e.g. 04:00-09:00
	PreferredBackupWindow pulumi.StringInput `pulumi:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur, in (UTC) e.g. wed:04:00-wed:04:30
	PreferredMaintenanceWindow pulumi.StringInput `pulumi:"preferredMaintenanceWindow"`
	// A read-only endpoint for the Aurora cluster, automatically
	// load-balanced across replicas
	ReaderEndpoint pulumi.StringInput `pulumi:"readerEndpoint"`
	// ARN of a source DB cluster or DB instance if this DB cluster is to be created as a Read Replica.
	ReplicationSourceIdentifier pulumi.StringInput `pulumi:"replicationSourceIdentifier"`
	S3Import pulumi.AnyInput `pulumi:"s3Import"`
	// Nested attribute with scaling properties. Only valid when `engineMode` is set to `serverless`. More details below.
	ScalingConfiguration pulumi.AnyInput `pulumi:"scalingConfiguration"`
	// Determines whether a final DB snapshot is created before the DB cluster is deleted. If true is specified, no DB snapshot is created. If false is specified, a DB snapshot is created before the DB cluster is deleted, using the value from `finalSnapshotIdentifier`. Default is `false`.
	SkipFinalSnapshot pulumi.BoolInput `pulumi:"skipFinalSnapshot"`
	// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a DB cluster snapshot, or the ARN when specifying a DB snapshot.
	SnapshotIdentifier pulumi.StringInput `pulumi:"snapshotIdentifier"`
	// The source region for an encrypted replica DB cluster.
	SourceRegion pulumi.StringInput `pulumi:"sourceRegion"`
	// Specifies whether the DB cluster is encrypted. The default is `false` for `provisioned` `engineMode` and `true` for `serverless` `engineMode`.
	StorageEncrypted pulumi.BoolInput `pulumi:"storageEncrypted"`
	// A mapping of tags to assign to the DB cluster.
	Tags pulumi.MapInput `pulumi:"tags"`
	// List of VPC security groups to associate
	// with the Cluster
	VpcSecurityGroupIds pulumi.ArrayInput `pulumi:"vpcSecurityGroupIds"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Specifies whether any cluster modifications
	// are applied immediately, or during the next maintenance window. Default is
	// `false`. See [Amazon RDS Documentation for more information.](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.DBInstance.Modifying.html)
	ApplyImmediately pulumi.BoolInput `pulumi:"applyImmediately"`
	// A list of EC2 Availability Zones for the DB cluster storage where DB cluster instances can be created. RDS automatically assigns 3 AZs if less than 3 AZs are configured, which will show as a difference requiring resource recreation next deployment. It is recommended to specify 3 AZs or use `ignoreChanges` if necessary.
	AvailabilityZones pulumi.ArrayInput `pulumi:"availabilityZones"`
	// The target backtrack window, in seconds. Only available for `aurora` engine currently. To disable backtracking, set this value to `0`. Defaults to `0`. Must be between `0` and `259200` (72 hours)
	BacktrackWindow pulumi.IntInput `pulumi:"backtrackWindow"`
	// The days to retain backups for. Default `1`
	BackupRetentionPeriod pulumi.IntInput `pulumi:"backupRetentionPeriod"`
	// The cluster identifier. If omitted, this provider will assign a random, unique identifier.
	ClusterIdentifier pulumi.StringInput `pulumi:"clusterIdentifier"`
	// Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `clusterIdentifier`.
	ClusterIdentifierPrefix pulumi.StringInput `pulumi:"clusterIdentifierPrefix"`
	// List of RDS Instances that are a part of this cluster
	ClusterMembers pulumi.ArrayInput `pulumi:"clusterMembers"`
	// Copy all Cluster `tags` to snapshots. Default is `false`.
	CopyTagsToSnapshot pulumi.BoolInput `pulumi:"copyTagsToSnapshot"`
	// Name for an automatically created database on cluster creation. There are different naming restrictions per database engine: [RDS Naming Constraints][5]
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// A cluster parameter group to associate with the cluster.
	DbClusterParameterGroupName pulumi.StringInput `pulumi:"dbClusterParameterGroupName"`
	// A DB subnet group to associate with this DB instance. **NOTE:** This must match the `dbSubnetGroupName` specified on every [`rds.ClusterInstance`](https://www.terraform.io/docs/providers/aws/r/rds_cluster_instance.html) in the cluster.
	DbSubnetGroupName pulumi.StringInput `pulumi:"dbSubnetGroupName"`
	// If the DB instance should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection pulumi.BoolInput `pulumi:"deletionProtection"`
	// List of log types to export to cloudwatch. If omitted, no logs will be exported.
	// The following log types are supported: `audit`, `error`, `general`, `slowquery`, `postgresql` (PostgreSQL).
	EnabledCloudwatchLogsExports pulumi.ArrayInput `pulumi:"enabledCloudwatchLogsExports"`
	// The name of the database engine to be used for this DB cluster. Defaults to `aurora`. Valid Values: `aurora`, `aurora-mysql`, `aurora-postgresql`
	Engine pulumi.StringInput `pulumi:"engine"`
	// The database engine mode. Valid values: `global`, `multimaster`, `parallelquery`, `provisioned`, `serverless`. Defaults to: `provisioned`. See the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/aurora-serverless.html) for limitations when using `serverless`.
	EngineMode pulumi.StringInput `pulumi:"engineMode"`
	// The database engine version. Updating this argument results in an outage. See the [Aurora MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Updates.html) and [Aurora Postgres](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Updates.html) documentation for your configured engine to determine this value. For example with Aurora MySQL 2, a potential value for this argument is `5.7.mysql_aurora.2.03.2`.
	EngineVersion pulumi.StringInput `pulumi:"engineVersion"`
	// The name of your final DB snapshot
	// when this DB cluster is deleted. If omitted, no final snapshot will be
	// made.
	FinalSnapshotIdentifier pulumi.StringInput `pulumi:"finalSnapshotIdentifier"`
	// The global cluster identifier specified on [`rds.GlobalCluster`](https://www.terraform.io/docs/providers/aws/r/rds_global_cluster.html).
	GlobalClusterIdentifier pulumi.StringInput `pulumi:"globalClusterIdentifier"`
	// Specifies whether or mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled. Please see [AWS Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html) for availability and limitations.
	IamDatabaseAuthenticationEnabled pulumi.BoolInput `pulumi:"iamDatabaseAuthenticationEnabled"`
	// A List of ARNs for the IAM roles to associate to the RDS Cluster.
	IamRoles pulumi.ArrayInput `pulumi:"iamRoles"`
	// The ARN for the KMS encryption key. When specifying `kmsKeyId`, `storageEncrypted` needs to be set to true.
	KmsKeyId pulumi.StringInput `pulumi:"kmsKeyId"`
	// Password for the master DB user. Note that this may
	// show up in logs, and it will be stored in the state file. Please refer to the [RDS Naming Constraints][5]
	MasterPassword pulumi.StringInput `pulumi:"masterPassword"`
	// Username for the master DB user. Please refer to the [RDS Naming Constraints][5]. This argument does not support in-place updates and cannot be changed during a restore from snapshot.
	MasterUsername pulumi.StringInput `pulumi:"masterUsername"`
	// The port on which the DB accepts connections
	Port pulumi.IntInput `pulumi:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter.Time in UTC
	// Default: A 30-minute window selected at random from an 8-hour block of time per region. e.g. 04:00-09:00
	PreferredBackupWindow pulumi.StringInput `pulumi:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur, in (UTC) e.g. wed:04:00-wed:04:30
	PreferredMaintenanceWindow pulumi.StringInput `pulumi:"preferredMaintenanceWindow"`
	// ARN of a source DB cluster or DB instance if this DB cluster is to be created as a Read Replica.
	ReplicationSourceIdentifier pulumi.StringInput `pulumi:"replicationSourceIdentifier"`
	S3Import pulumi.AnyInput `pulumi:"s3Import"`
	// Nested attribute with scaling properties. Only valid when `engineMode` is set to `serverless`. More details below.
	ScalingConfiguration pulumi.AnyInput `pulumi:"scalingConfiguration"`
	// Determines whether a final DB snapshot is created before the DB cluster is deleted. If true is specified, no DB snapshot is created. If false is specified, a DB snapshot is created before the DB cluster is deleted, using the value from `finalSnapshotIdentifier`. Default is `false`.
	SkipFinalSnapshot pulumi.BoolInput `pulumi:"skipFinalSnapshot"`
	// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a DB cluster snapshot, or the ARN when specifying a DB snapshot.
	SnapshotIdentifier pulumi.StringInput `pulumi:"snapshotIdentifier"`
	// The source region for an encrypted replica DB cluster.
	SourceRegion pulumi.StringInput `pulumi:"sourceRegion"`
	// Specifies whether the DB cluster is encrypted. The default is `false` for `provisioned` `engineMode` and `true` for `serverless` `engineMode`.
	StorageEncrypted pulumi.BoolInput `pulumi:"storageEncrypted"`
	// A mapping of tags to assign to the DB cluster.
	Tags pulumi.MapInput `pulumi:"tags"`
	// List of VPC security groups to associate
	// with the Cluster
	VpcSecurityGroupIds pulumi.ArrayInput `pulumi:"vpcSecurityGroupIds"`
}
