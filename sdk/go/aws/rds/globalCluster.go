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
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// RDS Global Cluster Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`

	// Name for an automatically created database on cluster creation.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`

	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`

	// Name of the database engine to be used for this DB cluster. Valid values: `aurora`. Defaults to `aurora`.
	Engine pulumi.StringOutput `pulumi:"engine"`

	// Engine version of the Aurora global database.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`

	// The global cluster identifier.
	GlobalClusterIdentifier pulumi.StringOutput `pulumi:"globalClusterIdentifier"`

	// AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
	GlobalClusterResourceId pulumi.StringOutput `pulumi:"globalClusterResourceId"`

	// Specifies whether the DB cluster is encrypted. The default is `false`.
	StorageEncrypted pulumi.BoolOutput `pulumi:"storageEncrypted"`
}

// NewGlobalCluster registers a new resource with the given unique name, arguments, and options.
func NewGlobalCluster(ctx *pulumi.Context,
	name string, args *GlobalClusterArgs, opts ...pulumi.ResourceOpt) (*GlobalCluster, error) {
	if args == nil || args.GlobalClusterIdentifier == nil {
		return nil, errors.New("missing required argument 'GlobalClusterIdentifier'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["databaseName"] = args.DatabaseName
		inputs["deletionProtection"] = args.DeletionProtection
		inputs["engine"] = args.Engine
		inputs["engineVersion"] = args.EngineVersion
		inputs["globalClusterIdentifier"] = args.GlobalClusterIdentifier
		inputs["storageEncrypted"] = args.StorageEncrypted
	}
	var resource GlobalCluster
	err := ctx.RegisterResource("aws:rds/globalCluster:GlobalCluster", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalCluster gets an existing GlobalCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalCluster(ctx *pulumi.Context,
	name string, id pulumi.ID, state *GlobalClusterState, opts ...pulumi.ResourceOpt) (*GlobalCluster, error) {
	inputs := map[string]pulumi.Input{}
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
	var resource GlobalCluster
	err := ctx.ReadResource("aws:rds/globalCluster:GlobalCluster", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *GlobalCluster) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *GlobalCluster) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering GlobalCluster resources.
type GlobalClusterState struct {
	// RDS Global Cluster Amazon Resource Name (ARN)
	Arn pulumi.StringInput `pulumi:"arn"`
	// Name for an automatically created database on cluster creation.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection pulumi.BoolInput `pulumi:"deletionProtection"`
	// Name of the database engine to be used for this DB cluster. Valid values: `aurora`. Defaults to `aurora`.
	Engine pulumi.StringInput `pulumi:"engine"`
	// Engine version of the Aurora global database.
	EngineVersion pulumi.StringInput `pulumi:"engineVersion"`
	// The global cluster identifier.
	GlobalClusterIdentifier pulumi.StringInput `pulumi:"globalClusterIdentifier"`
	// AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
	GlobalClusterResourceId pulumi.StringInput `pulumi:"globalClusterResourceId"`
	// Specifies whether the DB cluster is encrypted. The default is `false`.
	StorageEncrypted pulumi.BoolInput `pulumi:"storageEncrypted"`
}

// The set of arguments for constructing a GlobalCluster resource.
type GlobalClusterArgs struct {
	// Name for an automatically created database on cluster creation.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection pulumi.BoolInput `pulumi:"deletionProtection"`
	// Name of the database engine to be used for this DB cluster. Valid values: `aurora`. Defaults to `aurora`.
	Engine pulumi.StringInput `pulumi:"engine"`
	// Engine version of the Aurora global database.
	EngineVersion pulumi.StringInput `pulumi:"engineVersion"`
	// The global cluster identifier.
	GlobalClusterIdentifier pulumi.StringInput `pulumi:"globalClusterIdentifier"`
	// Specifies whether the DB cluster is encrypted. The default is `false`.
	StorageEncrypted pulumi.BoolInput `pulumi:"storageEncrypted"`
}
