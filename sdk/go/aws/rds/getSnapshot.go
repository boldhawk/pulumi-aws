// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get information about a DB Snapshot for use when provisioning DB instances
// 
// > **NOTE:** This data source does not apply to snapshots created on Aurora DB clusters.
// See the [`rds.ClusterSnapshot` data source](https://www.terraform.io/docs/providers/aws/d/db_cluster_snapshot.html) for DB Cluster snapshots.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/db_snapshot.html.markdown.
func LookupSnapshot(ctx *pulumi.Context, args *GetSnapshotArgs) (*GetSnapshotResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["dbInstanceIdentifier"] = args.DbInstanceIdentifier
		inputs["dbSnapshotIdentifier"] = args.DbSnapshotIdentifier
		inputs["includePublic"] = args.IncludePublic
		inputs["includeShared"] = args.IncludeShared
		inputs["mostRecent"] = args.MostRecent
		inputs["snapshotType"] = args.SnapshotType
	}
	outputs, err := ctx.Invoke("aws:rds/getSnapshot:getSnapshot", inputs)
	if err != nil {
		return nil, err
	}
	return &GetSnapshotResult{
		AllocatedStorage: outputs["allocatedStorage"],
		AvailabilityZone: outputs["availabilityZone"],
		DbInstanceIdentifier: outputs["dbInstanceIdentifier"],
		DbSnapshotArn: outputs["dbSnapshotArn"],
		DbSnapshotIdentifier: outputs["dbSnapshotIdentifier"],
		Encrypted: outputs["encrypted"],
		Engine: outputs["engine"],
		EngineVersion: outputs["engineVersion"],
		IncludePublic: outputs["includePublic"],
		IncludeShared: outputs["includeShared"],
		Iops: outputs["iops"],
		KmsKeyId: outputs["kmsKeyId"],
		LicenseModel: outputs["licenseModel"],
		MostRecent: outputs["mostRecent"],
		OptionGroupName: outputs["optionGroupName"],
		Port: outputs["port"],
		SnapshotCreateTime: outputs["snapshotCreateTime"],
		SnapshotType: outputs["snapshotType"],
		SourceDbSnapshotIdentifier: outputs["sourceDbSnapshotIdentifier"],
		SourceRegion: outputs["sourceRegion"],
		Status: outputs["status"],
		StorageType: outputs["storageType"],
		VpcId: outputs["vpcId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getSnapshot.
type GetSnapshotArgs struct {
	// Returns the list of snapshots created by the specific db_instance
	DbInstanceIdentifier pulumi.StringInput `pulumi:"dbInstanceIdentifier"`
	// Returns information on a specific snapshot_id.
	DbSnapshotIdentifier pulumi.StringInput `pulumi:"dbSnapshotIdentifier"`
	// Set this value to true to include manual DB snapshots that are public and can be
	// copied or restored by any AWS account, otherwise set this value to false. The default is `false`.
	IncludePublic pulumi.BoolInput `pulumi:"includePublic"`
	// Set this value to true to include shared manual DB snapshots from other
	// AWS accounts that this AWS account has been given permission to copy or restore, otherwise set this value to false.
	// The default is `false`.
	IncludeShared pulumi.BoolInput `pulumi:"includeShared"`
	// If more than one result is returned, use the most
	// recent Snapshot.
	MostRecent pulumi.BoolInput `pulumi:"mostRecent"`
	// The type of snapshots to be returned. If you don't specify a SnapshotType
	// value, then both automated and manual snapshots are returned. Shared and public DB snapshots are not
	// included in the returned results by default. Possible values are, `automated`, `manual`, `shared` and `public`.
	SnapshotType pulumi.StringInput `pulumi:"snapshotType"`
}

// A collection of values returned by getSnapshot.
type GetSnapshotResult struct {
	// Specifies the allocated storage size in gigabytes (GB).
	AllocatedStorage int `pulumi:"allocatedStorage"`
	// Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
	AvailabilityZone string `pulumi:"availabilityZone"`
	DbInstanceIdentifier string `pulumi:"dbInstanceIdentifier"`
	// The Amazon Resource Name (ARN) for the DB snapshot.
	DbSnapshotArn string `pulumi:"dbSnapshotArn"`
	DbSnapshotIdentifier string `pulumi:"dbSnapshotIdentifier"`
	// Specifies whether the DB snapshot is encrypted.
	Encrypted bool `pulumi:"encrypted"`
	// Specifies the name of the database engine.
	Engine string `pulumi:"engine"`
	// Specifies the version of the database engine.
	EngineVersion string `pulumi:"engineVersion"`
	IncludePublic bool `pulumi:"includePublic"`
	IncludeShared bool `pulumi:"includeShared"`
	// Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
	Iops int `pulumi:"iops"`
	// The ARN for the KMS encryption key.
	KmsKeyId string `pulumi:"kmsKeyId"`
	// License model information for the restored DB instance.
	LicenseModel string `pulumi:"licenseModel"`
	MostRecent bool `pulumi:"mostRecent"`
	// Provides the option group name for the DB snapshot.
	OptionGroupName string `pulumi:"optionGroupName"`
	Port int `pulumi:"port"`
	// Provides the time when the snapshot was taken, in Universal Coordinated Time (UTC).
	SnapshotCreateTime string `pulumi:"snapshotCreateTime"`
	SnapshotType string `pulumi:"snapshotType"`
	// The DB snapshot Arn that the DB snapshot was copied from. It only has value in case of cross customer or cross region copy.
	SourceDbSnapshotIdentifier string `pulumi:"sourceDbSnapshotIdentifier"`
	// The region that the DB snapshot was created in or copied from.
	SourceRegion string `pulumi:"sourceRegion"`
	// Specifies the status of this DB snapshot.
	Status string `pulumi:"status"`
	// Specifies the storage type associated with DB snapshot.
	StorageType string `pulumi:"storageType"`
	// Specifies the ID of the VPC associated with the DB snapshot.
	VpcId string `pulumi:"vpcId"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
