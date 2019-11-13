// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get information about an EBS volume for use in other
// resources.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ebs_volume.html.markdown.
func LookupVolume(ctx *pulumi.Context, args *GetVolumeArgs) (*GetVolumeResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["filters"] = args.Filters
		inputs["mostRecent"] = args.MostRecent
		inputs["tags"] = args.Tags
	}
	outputs, err := ctx.Invoke("aws:ebs/getVolume:getVolume", inputs)
	if err != nil {
		return nil, err
	}
	return &GetVolumeResult{
		Arn: outputs["arn"],
		AvailabilityZone: outputs["availabilityZone"],
		Encrypted: outputs["encrypted"],
		Filters: outputs["filters"],
		Iops: outputs["iops"],
		KmsKeyId: outputs["kmsKeyId"],
		MostRecent: outputs["mostRecent"],
		Size: outputs["size"],
		SnapshotId: outputs["snapshotId"],
		Tags: outputs["tags"],
		VolumeId: outputs["volumeId"],
		VolumeType: outputs["volumeType"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getVolume.
type GetVolumeArgs struct {
	// One or more name/value pairs to filter off of. There are
	// several valid keys, for a full reference, check out
	// [describe-volumes in the AWS CLI reference][1].
	Filters pulumi.ArrayInput `pulumi:"filters"`
	// If more than one result is returned, use the most
	// recent Volume.
	MostRecent pulumi.BoolInput `pulumi:"mostRecent"`
	Tags pulumi.MapInput `pulumi:"tags"`
}

// A collection of values returned by getVolume.
type GetVolumeResult struct {
	// The volume ARN (e.g. arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).
	Arn string `pulumi:"arn"`
	// The AZ where the EBS volume exists.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Whether the disk is encrypted.
	Encrypted bool `pulumi:"encrypted"`
	Filters []interface{} `pulumi:"filters"`
	// The amount of IOPS for the disk.
	Iops int `pulumi:"iops"`
	// The ARN for the KMS encryption key.
	KmsKeyId string `pulumi:"kmsKeyId"`
	MostRecent bool `pulumi:"mostRecent"`
	// The size of the drive in GiBs.
	Size int `pulumi:"size"`
	// The snapshotId the EBS volume is based off.
	SnapshotId string `pulumi:"snapshotId"`
	// A mapping of tags for the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The volume ID (e.g. vol-59fcb34e).
	VolumeId string `pulumi:"volumeId"`
	// The type of EBS volume.
	VolumeType string `pulumi:"volumeType"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
