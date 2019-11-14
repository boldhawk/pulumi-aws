// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an ElastiCache Cluster resource, which manages a Memcached cluster or Redis instance.
// For working with Redis (Cluster Mode Enabled) replication groups, see the
// [`elasticache.ReplicationGroup` resource](https://www.terraform.io/docs/providers/aws/r/elasticache_replication_group.html).
// 
// > **Note:** When you change an attribute, such as `nodeType`, by default
// it is applied in the next maintenance window. Because of this, this provider may report
// a difference in its planning phase because the actual modification has not yet taken
// place. You can use the `applyImmediately` flag to instruct the service to apply the
// change immediately. Using `applyImmediately` can result in a brief downtime as the server reboots.
// See the AWS Docs on [Modifying an ElastiCache Cache Cluster][2] for more information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/elasticache_cluster.html.markdown.
type Cluster struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is
	// `false`. See [Amazon ElastiCache Documentation for more information.][1]
	// (Available since v0.6.0)
	ApplyImmediately pulumi.BoolOutput `pulumi:"applyImmediately"`

	// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferredAvailabilityZones` instead. Default: System chosen Availability Zone.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`

	// Specifies whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `numCacheNodes` must be greater than `1`
	AzMode pulumi.StringOutput `pulumi:"azMode"`

	// List of node objects including `id`, `address`, `port` and `availabilityZone`.
	// Referenceable e.g. as `${aws_elasticache_cluster.bar.cache_nodes.0.address}`
	CacheNodes pulumi.ArrayOutput `pulumi:"cacheNodes"`

	// (Memcached only) The DNS name of the cache cluster without the port appended.
	ClusterAddress pulumi.StringOutput `pulumi:"clusterAddress"`

	// Group identifier. ElastiCache converts
	// this name to lowercase
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`

	// (Memcached only) The configuration endpoint to allow host discovery.
	ConfigurationEndpoint pulumi.StringOutput `pulumi:"configurationEndpoint"`

	// Name of the cache engine to be used for this cache cluster.
	// Valid values for this parameter are `memcached` or `redis`
	Engine pulumi.StringOutput `pulumi:"engine"`

	// Version number of the cache engine to be used.
	// See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html)
	// in the AWS Documentation center for supported versions
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`

	// Specifies the weekly time range for when maintenance
	// on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
	// The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
	MaintenanceWindow pulumi.StringOutput `pulumi:"maintenanceWindow"`

	// The compute and memory capacity of the nodes. See
	// [Available Cache Node Types](https://aws.amazon.com/elasticache/details#Available_Cache_Node_Types) for
	// supported node types
	NodeType pulumi.StringOutput `pulumi:"nodeType"`

	// An Amazon Resource Name (ARN) of an
	// SNS topic to send ElastiCache notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn pulumi.StringOutput `pulumi:"notificationTopicArn"`

	// The initial number of cache nodes that the
	// cache cluster will have. For Redis, this value must be 1. For Memcache, this
	// value must be between 1 and 20. If this number is reduced on subsequent runs,
	// the highest numbered nodes will be removed.
	NumCacheNodes pulumi.IntOutput `pulumi:"numCacheNodes"`

	// Name of the parameter group to associate
	// with this cache cluster
	ParameterGroupName pulumi.StringOutput `pulumi:"parameterGroupName"`

	// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replicationGroupId`.
	Port pulumi.IntOutput `pulumi:"port"`

	// A list of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of `numCacheNodes`. If you want all the nodes in the same Availability Zone, use `availabilityZone` instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
	PreferredAvailabilityZones pulumi.ArrayOutput `pulumi:"preferredAvailabilityZones"`

	// The ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
	ReplicationGroupId pulumi.StringOutput `pulumi:"replicationGroupId"`

	// One or more VPC security groups associated
	// with the cache cluster
	SecurityGroupIds pulumi.ArrayOutput `pulumi:"securityGroupIds"`

	// List of security group
	// names to associate with this cache cluster
	SecurityGroupNames pulumi.ArrayOutput `pulumi:"securityGroupNames"`

	// A single-element string list containing an
	// Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
	// Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
	SnapshotArns pulumi.ArrayOutput `pulumi:"snapshotArns"`

	// The name of a snapshot from which to restore data into the new node group.  Changing the `snapshotName` forces a new resource.
	SnapshotName pulumi.StringOutput `pulumi:"snapshotName"`

	// The number of days for which ElastiCache will
	// retain automatic cache cluster snapshots before deleting them. For example, if you set
	// SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
	// before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
	// Please note that setting a `snapshotRetentionLimit` is not supported on cache.t1.micro or cache.t2.* cache nodes
	SnapshotRetentionLimit pulumi.IntOutput `pulumi:"snapshotRetentionLimit"`

	// The daily time range (in UTC) during which ElastiCache will
	// begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
	SnapshotWindow pulumi.StringOutput `pulumi:"snapshotWindow"`

	// Name of the subnet group to be used
	// for the cache cluster.
	SubnetGroupName pulumi.StringOutput `pulumi:"subnetGroupName"`

	// A mapping of tags to assign to the resource
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOpt) (*Cluster, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["applyImmediately"] = args.ApplyImmediately
		inputs["availabilityZone"] = args.AvailabilityZone
		inputs["azMode"] = args.AzMode
		inputs["clusterId"] = args.ClusterId
		inputs["engine"] = args.Engine
		inputs["engineVersion"] = args.EngineVersion
		inputs["maintenanceWindow"] = args.MaintenanceWindow
		inputs["nodeType"] = args.NodeType
		inputs["notificationTopicArn"] = args.NotificationTopicArn
		inputs["numCacheNodes"] = args.NumCacheNodes
		inputs["parameterGroupName"] = args.ParameterGroupName
		inputs["port"] = args.Port
		inputs["preferredAvailabilityZones"] = args.PreferredAvailabilityZones
		inputs["replicationGroupId"] = args.ReplicationGroupId
		inputs["securityGroupIds"] = args.SecurityGroupIds
		inputs["securityGroupNames"] = args.SecurityGroupNames
		inputs["snapshotArns"] = args.SnapshotArns
		inputs["snapshotName"] = args.SnapshotName
		inputs["snapshotRetentionLimit"] = args.SnapshotRetentionLimit
		inputs["snapshotWindow"] = args.SnapshotWindow
		inputs["subnetGroupName"] = args.SubnetGroupName
		inputs["tags"] = args.Tags
	}
	var resource Cluster
	err := ctx.RegisterResource("aws:elasticache/cluster:Cluster", name, inputs, &resource, opts...)
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
		inputs["availabilityZone"] = state.AvailabilityZone
		inputs["azMode"] = state.AzMode
		inputs["cacheNodes"] = state.CacheNodes
		inputs["clusterAddress"] = state.ClusterAddress
		inputs["clusterId"] = state.ClusterId
		inputs["configurationEndpoint"] = state.ConfigurationEndpoint
		inputs["engine"] = state.Engine
		inputs["engineVersion"] = state.EngineVersion
		inputs["maintenanceWindow"] = state.MaintenanceWindow
		inputs["nodeType"] = state.NodeType
		inputs["notificationTopicArn"] = state.NotificationTopicArn
		inputs["numCacheNodes"] = state.NumCacheNodes
		inputs["parameterGroupName"] = state.ParameterGroupName
		inputs["port"] = state.Port
		inputs["preferredAvailabilityZones"] = state.PreferredAvailabilityZones
		inputs["replicationGroupId"] = state.ReplicationGroupId
		inputs["securityGroupIds"] = state.SecurityGroupIds
		inputs["securityGroupNames"] = state.SecurityGroupNames
		inputs["snapshotArns"] = state.SnapshotArns
		inputs["snapshotName"] = state.SnapshotName
		inputs["snapshotRetentionLimit"] = state.SnapshotRetentionLimit
		inputs["snapshotWindow"] = state.SnapshotWindow
		inputs["subnetGroupName"] = state.SubnetGroupName
		inputs["tags"] = state.Tags
	}
	var resource Cluster
	err := ctx.ReadResource("aws:elasticache/cluster:Cluster", name, id, inputs, &resource, opts...)
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
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is
	// `false`. See [Amazon ElastiCache Documentation for more information.][1]
	// (Available since v0.6.0)
	ApplyImmediately pulumi.BoolInput `pulumi:"applyImmediately"`
	// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferredAvailabilityZones` instead. Default: System chosen Availability Zone.
	AvailabilityZone pulumi.StringInput `pulumi:"availabilityZone"`
	// Specifies whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `numCacheNodes` must be greater than `1`
	AzMode pulumi.StringInput `pulumi:"azMode"`
	// List of node objects including `id`, `address`, `port` and `availabilityZone`.
	// Referenceable e.g. as `${aws_elasticache_cluster.bar.cache_nodes.0.address}`
	CacheNodes pulumi.ArrayInput `pulumi:"cacheNodes"`
	// (Memcached only) The DNS name of the cache cluster without the port appended.
	ClusterAddress pulumi.StringInput `pulumi:"clusterAddress"`
	// Group identifier. ElastiCache converts
	// this name to lowercase
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// (Memcached only) The configuration endpoint to allow host discovery.
	ConfigurationEndpoint pulumi.StringInput `pulumi:"configurationEndpoint"`
	// Name of the cache engine to be used for this cache cluster.
	// Valid values for this parameter are `memcached` or `redis`
	Engine pulumi.StringInput `pulumi:"engine"`
	// Version number of the cache engine to be used.
	// See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html)
	// in the AWS Documentation center for supported versions
	EngineVersion pulumi.StringInput `pulumi:"engineVersion"`
	// Specifies the weekly time range for when maintenance
	// on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
	// The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
	MaintenanceWindow pulumi.StringInput `pulumi:"maintenanceWindow"`
	// The compute and memory capacity of the nodes. See
	// [Available Cache Node Types](https://aws.amazon.com/elasticache/details#Available_Cache_Node_Types) for
	// supported node types
	NodeType pulumi.StringInput `pulumi:"nodeType"`
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send ElastiCache notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn pulumi.StringInput `pulumi:"notificationTopicArn"`
	// The initial number of cache nodes that the
	// cache cluster will have. For Redis, this value must be 1. For Memcache, this
	// value must be between 1 and 20. If this number is reduced on subsequent runs,
	// the highest numbered nodes will be removed.
	NumCacheNodes pulumi.IntInput `pulumi:"numCacheNodes"`
	// Name of the parameter group to associate
	// with this cache cluster
	ParameterGroupName pulumi.StringInput `pulumi:"parameterGroupName"`
	// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replicationGroupId`.
	Port pulumi.IntInput `pulumi:"port"`
	// A list of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of `numCacheNodes`. If you want all the nodes in the same Availability Zone, use `availabilityZone` instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
	PreferredAvailabilityZones pulumi.ArrayInput `pulumi:"preferredAvailabilityZones"`
	// The ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
	ReplicationGroupId pulumi.StringInput `pulumi:"replicationGroupId"`
	// One or more VPC security groups associated
	// with the cache cluster
	SecurityGroupIds pulumi.ArrayInput `pulumi:"securityGroupIds"`
	// List of security group
	// names to associate with this cache cluster
	SecurityGroupNames pulumi.ArrayInput `pulumi:"securityGroupNames"`
	// A single-element string list containing an
	// Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
	// Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
	SnapshotArns pulumi.ArrayInput `pulumi:"snapshotArns"`
	// The name of a snapshot from which to restore data into the new node group.  Changing the `snapshotName` forces a new resource.
	SnapshotName pulumi.StringInput `pulumi:"snapshotName"`
	// The number of days for which ElastiCache will
	// retain automatic cache cluster snapshots before deleting them. For example, if you set
	// SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
	// before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
	// Please note that setting a `snapshotRetentionLimit` is not supported on cache.t1.micro or cache.t2.* cache nodes
	SnapshotRetentionLimit pulumi.IntInput `pulumi:"snapshotRetentionLimit"`
	// The daily time range (in UTC) during which ElastiCache will
	// begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
	SnapshotWindow pulumi.StringInput `pulumi:"snapshotWindow"`
	// Name of the subnet group to be used
	// for the cache cluster.
	SubnetGroupName pulumi.StringInput `pulumi:"subnetGroupName"`
	// A mapping of tags to assign to the resource
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is
	// `false`. See [Amazon ElastiCache Documentation for more information.][1]
	// (Available since v0.6.0)
	ApplyImmediately pulumi.BoolInput `pulumi:"applyImmediately"`
	// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferredAvailabilityZones` instead. Default: System chosen Availability Zone.
	AvailabilityZone pulumi.StringInput `pulumi:"availabilityZone"`
	// Specifies whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `numCacheNodes` must be greater than `1`
	AzMode pulumi.StringInput `pulumi:"azMode"`
	// Group identifier. ElastiCache converts
	// this name to lowercase
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// Name of the cache engine to be used for this cache cluster.
	// Valid values for this parameter are `memcached` or `redis`
	Engine pulumi.StringInput `pulumi:"engine"`
	// Version number of the cache engine to be used.
	// See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html)
	// in the AWS Documentation center for supported versions
	EngineVersion pulumi.StringInput `pulumi:"engineVersion"`
	// Specifies the weekly time range for when maintenance
	// on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
	// The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
	MaintenanceWindow pulumi.StringInput `pulumi:"maintenanceWindow"`
	// The compute and memory capacity of the nodes. See
	// [Available Cache Node Types](https://aws.amazon.com/elasticache/details#Available_Cache_Node_Types) for
	// supported node types
	NodeType pulumi.StringInput `pulumi:"nodeType"`
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send ElastiCache notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn pulumi.StringInput `pulumi:"notificationTopicArn"`
	// The initial number of cache nodes that the
	// cache cluster will have. For Redis, this value must be 1. For Memcache, this
	// value must be between 1 and 20. If this number is reduced on subsequent runs,
	// the highest numbered nodes will be removed.
	NumCacheNodes pulumi.IntInput `pulumi:"numCacheNodes"`
	// Name of the parameter group to associate
	// with this cache cluster
	ParameterGroupName pulumi.StringInput `pulumi:"parameterGroupName"`
	// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replicationGroupId`.
	Port pulumi.IntInput `pulumi:"port"`
	// A list of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of `numCacheNodes`. If you want all the nodes in the same Availability Zone, use `availabilityZone` instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
	PreferredAvailabilityZones pulumi.ArrayInput `pulumi:"preferredAvailabilityZones"`
	// The ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
	ReplicationGroupId pulumi.StringInput `pulumi:"replicationGroupId"`
	// One or more VPC security groups associated
	// with the cache cluster
	SecurityGroupIds pulumi.ArrayInput `pulumi:"securityGroupIds"`
	// List of security group
	// names to associate with this cache cluster
	SecurityGroupNames pulumi.ArrayInput `pulumi:"securityGroupNames"`
	// A single-element string list containing an
	// Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
	// Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
	SnapshotArns pulumi.ArrayInput `pulumi:"snapshotArns"`
	// The name of a snapshot from which to restore data into the new node group.  Changing the `snapshotName` forces a new resource.
	SnapshotName pulumi.StringInput `pulumi:"snapshotName"`
	// The number of days for which ElastiCache will
	// retain automatic cache cluster snapshots before deleting them. For example, if you set
	// SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
	// before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
	// Please note that setting a `snapshotRetentionLimit` is not supported on cache.t1.micro or cache.t2.* cache nodes
	SnapshotRetentionLimit pulumi.IntInput `pulumi:"snapshotRetentionLimit"`
	// The daily time range (in UTC) during which ElastiCache will
	// begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
	SnapshotWindow pulumi.StringInput `pulumi:"snapshotWindow"`
	// Name of the subnet group to be used
	// for the cache cluster.
	SubnetGroupName pulumi.StringInput `pulumi:"subnetGroupName"`
	// A mapping of tags to assign to the resource
	Tags pulumi.MapInput `pulumi:"tags"`
}
