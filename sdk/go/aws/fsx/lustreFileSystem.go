// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fsx

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a FSx Lustre File System. See the [FSx Lustre Guide](https://docs.aws.amazon.com/fsx/latest/LustreGuide/what-is.html) for more information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/fsx_lustre_file_system.html.markdown.
type LustreFileSystem struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// Amazon Resource Name of the file system.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// DNS name for the file system, e.g. `fs-12345678.fsx.us-west-2.amazonaws.com`
	DnsName pulumi.StringOutput `pulumi:"dnsName"`

	// S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `importPath` argument and the path must use the same Amazon S3 bucket as specified in `importPath`. Set equal to `importPath` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
	ExportPath pulumi.StringOutput `pulumi:"exportPath"`

	// S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
	ImportPath pulumi.StringOutput `pulumi:"importPath"`

	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `importPath` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
	ImportedFileChunkSize pulumi.IntOutput `pulumi:"importedFileChunkSize"`

	// Set of Elastic Network Interface identifiers from which the file system is accessible.
	NetworkInterfaceIds pulumi.ArrayOutput `pulumi:"networkInterfaceIds"`

	// AWS account identifier that created the file system.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`

	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds pulumi.ArrayOutput `pulumi:"securityGroupIds"`

	// The storage capacity (GiB) of the file system. Minimum of `3600`. Storage capacity is provisioned in increments of 3,600 GiB.
	StorageCapacity pulumi.IntOutput `pulumi:"storageCapacity"`

	// A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
	SubnetIds pulumi.StringOutput `pulumi:"subnetIds"`

	// A mapping of tags to assign to the file system.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// Identifier of the Virtual Private Cloud for the file system.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`

	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime pulumi.StringOutput `pulumi:"weeklyMaintenanceStartTime"`
}

// NewLustreFileSystem registers a new resource with the given unique name, arguments, and options.
func NewLustreFileSystem(ctx *pulumi.Context,
	name string, args *LustreFileSystemArgs, opts ...pulumi.ResourceOpt) (*LustreFileSystem, error) {
	if args == nil || args.StorageCapacity == nil {
		return nil, errors.New("missing required argument 'StorageCapacity'")
	}
	if args == nil || args.SubnetIds == nil {
		return nil, errors.New("missing required argument 'SubnetIds'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["exportPath"] = args.ExportPath
		inputs["importPath"] = args.ImportPath
		inputs["importedFileChunkSize"] = args.ImportedFileChunkSize
		inputs["securityGroupIds"] = args.SecurityGroupIds
		inputs["storageCapacity"] = args.StorageCapacity
		inputs["subnetIds"] = args.SubnetIds
		inputs["tags"] = args.Tags
		inputs["weeklyMaintenanceStartTime"] = args.WeeklyMaintenanceStartTime
	}
	var resource LustreFileSystem
	err := ctx.RegisterResource("aws:fsx/lustreFileSystem:LustreFileSystem", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLustreFileSystem gets an existing LustreFileSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLustreFileSystem(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LustreFileSystemState, opts ...pulumi.ResourceOpt) (*LustreFileSystem, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["dnsName"] = state.DnsName
		inputs["exportPath"] = state.ExportPath
		inputs["importPath"] = state.ImportPath
		inputs["importedFileChunkSize"] = state.ImportedFileChunkSize
		inputs["networkInterfaceIds"] = state.NetworkInterfaceIds
		inputs["ownerId"] = state.OwnerId
		inputs["securityGroupIds"] = state.SecurityGroupIds
		inputs["storageCapacity"] = state.StorageCapacity
		inputs["subnetIds"] = state.SubnetIds
		inputs["tags"] = state.Tags
		inputs["vpcId"] = state.VpcId
		inputs["weeklyMaintenanceStartTime"] = state.WeeklyMaintenanceStartTime
	}
	var resource LustreFileSystem
	err := ctx.ReadResource("aws:fsx/lustreFileSystem:LustreFileSystem", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *LustreFileSystem) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *LustreFileSystem) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering LustreFileSystem resources.
type LustreFileSystemState struct {
	// Amazon Resource Name of the file system.
	Arn pulumi.StringInput `pulumi:"arn"`
	// DNS name for the file system, e.g. `fs-12345678.fsx.us-west-2.amazonaws.com`
	DnsName pulumi.StringInput `pulumi:"dnsName"`
	// S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `importPath` argument and the path must use the same Amazon S3 bucket as specified in `importPath`. Set equal to `importPath` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
	ExportPath pulumi.StringInput `pulumi:"exportPath"`
	// S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
	ImportPath pulumi.StringInput `pulumi:"importPath"`
	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `importPath` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
	ImportedFileChunkSize pulumi.IntInput `pulumi:"importedFileChunkSize"`
	// Set of Elastic Network Interface identifiers from which the file system is accessible.
	NetworkInterfaceIds pulumi.ArrayInput `pulumi:"networkInterfaceIds"`
	// AWS account identifier that created the file system.
	OwnerId pulumi.StringInput `pulumi:"ownerId"`
	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds pulumi.ArrayInput `pulumi:"securityGroupIds"`
	// The storage capacity (GiB) of the file system. Minimum of `3600`. Storage capacity is provisioned in increments of 3,600 GiB.
	StorageCapacity pulumi.IntInput `pulumi:"storageCapacity"`
	// A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
	SubnetIds pulumi.StringInput `pulumi:"subnetIds"`
	// A mapping of tags to assign to the file system.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Identifier of the Virtual Private Cloud for the file system.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime pulumi.StringInput `pulumi:"weeklyMaintenanceStartTime"`
}

// The set of arguments for constructing a LustreFileSystem resource.
type LustreFileSystemArgs struct {
	// S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `importPath` argument and the path must use the same Amazon S3 bucket as specified in `importPath`. Set equal to `importPath` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
	ExportPath pulumi.StringInput `pulumi:"exportPath"`
	// S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
	ImportPath pulumi.StringInput `pulumi:"importPath"`
	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `importPath` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
	ImportedFileChunkSize pulumi.IntInput `pulumi:"importedFileChunkSize"`
	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds pulumi.ArrayInput `pulumi:"securityGroupIds"`
	// The storage capacity (GiB) of the file system. Minimum of `3600`. Storage capacity is provisioned in increments of 3,600 GiB.
	StorageCapacity pulumi.IntInput `pulumi:"storageCapacity"`
	// A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
	SubnetIds pulumi.StringInput `pulumi:"subnetIds"`
	// A mapping of tags to assign to the file system.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime pulumi.StringInput `pulumi:"weeklyMaintenanceStartTime"`
}
