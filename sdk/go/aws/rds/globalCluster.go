// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a RDS Global Cluster, which is an Aurora global database spread across multiple regions. The global database contains a single primary cluster with read-write capability, and a read-only secondary cluster that receives data from the primary cluster through high-speed replication performed by the Aurora storage subsystem.
// 
// More information about Aurora global databases can be found in the [Aurora User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database-creating).
// 
// > **NOTE:** RDS only supports the `aurora` engine (MySQL 5.6 compatible) for Global Clusters at this time.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/rds_global_cluster.html.markdown.
type GlobalCluster struct {
	s *pulumi.ResourceState
}

// NewGlobalCluster registers a new resource with the given unique name, arguments, and options.
func NewGlobalCluster(ctx *pulumi.Context,
	name string, args *GlobalClusterArgs, opts ...pulumi.ResourceOpt) (*GlobalCluster, error) {
	if args == nil || args.GlobalClusterIdentifier == nil {
		return nil, errors.New("missing required argument 'GlobalClusterIdentifier'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["databaseName"] = nil
		inputs["deletionProtection"] = nil
		inputs["engine"] = nil
		inputs["engineVersion"] = nil
		inputs["globalClusterIdentifier"] = nil
		inputs["storageEncrypted"] = nil
	} else {
		inputs["databaseName"] = args.DatabaseName
		inputs["deletionProtection"] = args.DeletionProtection
		inputs["engine"] = args.Engine
		inputs["engineVersion"] = args.EngineVersion
		inputs["globalClusterIdentifier"] = args.GlobalClusterIdentifier
		inputs["storageEncrypted"] = args.StorageEncrypted
	}
	inputs["arn"] = nil
	inputs["globalClusterResourceId"] = nil
	s, err := ctx.RegisterResource("aws:rds/globalCluster:GlobalCluster", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &GlobalCluster{s: s}, nil
}

// GetGlobalCluster gets an existing GlobalCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalCluster(ctx *pulumi.Context,
	name string, id pulumi.ID, state *GlobalClusterState, opts ...pulumi.ResourceOpt) (*GlobalCluster, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["databaseName"] = state.DatabaseName
		inputs["deletionProtection"] = state.DeletionProtection
		inputs["engine"] = state.Engine
		inputs["engineVersion"] = state.EngineVersion
		inputs["globalClusterIdentifier"] = state.GlobalClusterIdentifier
		inputs["globalClusterResourceId"] = state.GlobalClusterResourceId
		inputs["storageEncrypted"] = state.StorageEncrypted
	}
	s, err := ctx.ReadResource("aws:rds/globalCluster:GlobalCluster", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &GlobalCluster{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *GlobalCluster) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *GlobalCluster) ID() pulumi.IDOutput {
	return r.s.ID()
}

// RDS Global Cluster Amazon Resource Name (ARN)
func (r *GlobalCluster) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// Name for an automatically created database on cluster creation.
func (r *GlobalCluster) DatabaseName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["databaseName"])
}

// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
func (r *GlobalCluster) DeletionProtection() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["deletionProtection"])
}

// Name of the database engine to be used for this DB cluster. Valid values: `aurora`. Defaults to `aurora`.
func (r *GlobalCluster) Engine() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["engine"])
}

// Engine version of the Aurora global database.
func (r *GlobalCluster) EngineVersion() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["engineVersion"])
}

// The global cluster identifier.
func (r *GlobalCluster) GlobalClusterIdentifier() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["globalClusterIdentifier"])
}

// AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
func (r *GlobalCluster) GlobalClusterResourceId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["globalClusterResourceId"])
}

// Specifies whether the DB cluster is encrypted. The default is `false`.
func (r *GlobalCluster) StorageEncrypted() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["storageEncrypted"])
}

// Input properties used for looking up and filtering GlobalCluster resources.
type GlobalClusterState struct {
	// RDS Global Cluster Amazon Resource Name (ARN)
	Arn interface{}
	// Name for an automatically created database on cluster creation.
	DatabaseName interface{}
	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection interface{}
	// Name of the database engine to be used for this DB cluster. Valid values: `aurora`. Defaults to `aurora`.
	Engine interface{}
	// Engine version of the Aurora global database.
	EngineVersion interface{}
	// The global cluster identifier.
	GlobalClusterIdentifier interface{}
	// AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
	GlobalClusterResourceId interface{}
	// Specifies whether the DB cluster is encrypted. The default is `false`.
	StorageEncrypted interface{}
}

// The set of arguments for constructing a GlobalCluster resource.
type GlobalClusterArgs struct {
	// Name for an automatically created database on cluster creation.
	DatabaseName interface{}
	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection interface{}
	// Name of the database engine to be used for this DB cluster. Valid values: `aurora`. Defaults to `aurora`.
	Engine interface{}
	// Engine version of the Aurora global database.
	EngineVersion interface{}
	// The global cluster identifier.
	GlobalClusterIdentifier interface{}
	// Specifies whether the DB cluster is encrypted. The default is `false`.
	StorageEncrypted interface{}
}
