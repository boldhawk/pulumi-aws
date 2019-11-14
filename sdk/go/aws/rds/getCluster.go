// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides information about a RDS cluster.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/rds_cluster.html.markdown.
func LookupCluster(ctx *pulumi.Context, args *GetClusterArgs) (*GetClusterResult, error) {
	var rv GetClusterResult
	err := ctx.Invoke("aws:rds/getCluster:getCluster", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type GetClusterArgs struct {
	// The cluster identifier of the RDS cluster.
	ClusterIdentifier string `pulumi:"clusterIdentifier"`
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getCluster.
type GetClusterResult struct {
	Arn string `pulumi:"arn"`
	AvailabilityZones []string `pulumi:"availabilityZones"`
	BackupRetentionPeriod int `pulumi:"backupRetentionPeriod"`
	ClusterIdentifier string `pulumi:"clusterIdentifier"`
	ClusterMembers []string `pulumi:"clusterMembers"`
	ClusterResourceId string `pulumi:"clusterResourceId"`
	DatabaseName string `pulumi:"databaseName"`
	DbClusterParameterGroupName string `pulumi:"dbClusterParameterGroupName"`
	DbSubnetGroupName string `pulumi:"dbSubnetGroupName"`
	EnabledCloudwatchLogsExports []string `pulumi:"enabledCloudwatchLogsExports"`
	Endpoint string `pulumi:"endpoint"`
	Engine string `pulumi:"engine"`
	EngineVersion string `pulumi:"engineVersion"`
	FinalSnapshotIdentifier string `pulumi:"finalSnapshotIdentifier"`
	HostedZoneId string `pulumi:"hostedZoneId"`
	IamDatabaseAuthenticationEnabled bool `pulumi:"iamDatabaseAuthenticationEnabled"`
	IamRoles []string `pulumi:"iamRoles"`
	KmsKeyId string `pulumi:"kmsKeyId"`
	MasterUsername string `pulumi:"masterUsername"`
	Port int `pulumi:"port"`
	PreferredBackupWindow string `pulumi:"preferredBackupWindow"`
	PreferredMaintenanceWindow string `pulumi:"preferredMaintenanceWindow"`
	ReaderEndpoint string `pulumi:"readerEndpoint"`
	ReplicationSourceIdentifier string `pulumi:"replicationSourceIdentifier"`
	StorageEncrypted bool `pulumi:"storageEncrypted"`
	Tags map[string]string `pulumi:"tags"`
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
