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
	s *pulumi.ResourceState
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
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["exportPath"] = nil
		inputs["importPath"] = nil
		inputs["importedFileChunkSize"] = nil
		inputs["securityGroupIds"] = nil
		inputs["storageCapacity"] = nil
		inputs["subnetIds"] = nil
		inputs["tags"] = nil
		inputs["weeklyMaintenanceStartTime"] = nil
	} else {
		inputs["exportPath"] = args.ExportPath
		inputs["importPath"] = args.ImportPath
		inputs["importedFileChunkSize"] = args.ImportedFileChunkSize
		inputs["securityGroupIds"] = args.SecurityGroupIds
		inputs["storageCapacity"] = args.StorageCapacity
		inputs["subnetIds"] = args.SubnetIds
		inputs["tags"] = args.Tags
		inputs["weeklyMaintenanceStartTime"] = args.WeeklyMaintenanceStartTime
	}
	inputs["arn"] = nil
	inputs["dnsName"] = nil
	inputs["networkInterfaceIds"] = nil
	inputs["ownerId"] = nil
	inputs["vpcId"] = nil
	s, err := ctx.RegisterResource("aws:fsx/lustreFileSystem:LustreFileSystem", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LustreFileSystem{s: s}, nil
}

// GetLustreFileSystem gets an existing LustreFileSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLustreFileSystem(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LustreFileSystemState, opts ...pulumi.ResourceOpt) (*LustreFileSystem, error) {
	inputs := make(map[string]interface{})
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
	s, err := ctx.ReadResource("aws:fsx/lustreFileSystem:LustreFileSystem", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LustreFileSystem{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *LustreFileSystem) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *LustreFileSystem) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Amazon Resource Name of the file system.
func (r *LustreFileSystem) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// DNS name for the file system, e.g. `fs-12345678.fsx.us-west-2.amazonaws.com`
func (r *LustreFileSystem) DnsName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dnsName"])
}

// S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `importPath` argument and the path must use the same Amazon S3 bucket as specified in `importPath`. Set equal to `importPath` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
func (r *LustreFileSystem) ExportPath() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["exportPath"])
}

// S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
func (r *LustreFileSystem) ImportPath() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["importPath"])
}

// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `importPath` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
func (r *LustreFileSystem) ImportedFileChunkSize() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["importedFileChunkSize"])
}

// Set of Elastic Network Interface identifiers from which the file system is accessible.
func (r *LustreFileSystem) NetworkInterfaceIds() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["networkInterfaceIds"])
}

// AWS account identifier that created the file system.
func (r *LustreFileSystem) OwnerId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ownerId"])
}

// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
func (r *LustreFileSystem) SecurityGroupIds() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["securityGroupIds"])
}

// The storage capacity (GiB) of the file system. Minimum of `3600`. Storage capacity is provisioned in increments of 3,600 GiB.
func (r *LustreFileSystem) StorageCapacity() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["storageCapacity"])
}

// A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
func (r *LustreFileSystem) SubnetIds() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["subnetIds"])
}

// A mapping of tags to assign to the file system.
func (r *LustreFileSystem) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Identifier of the Virtual Private Cloud for the file system.
func (r *LustreFileSystem) VpcId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["vpcId"])
}

// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
func (r *LustreFileSystem) WeeklyMaintenanceStartTime() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["weeklyMaintenanceStartTime"])
}

// Input properties used for looking up and filtering LustreFileSystem resources.
type LustreFileSystemState struct {
	// Amazon Resource Name of the file system.
	Arn interface{}
	// DNS name for the file system, e.g. `fs-12345678.fsx.us-west-2.amazonaws.com`
	DnsName interface{}
	// S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `importPath` argument and the path must use the same Amazon S3 bucket as specified in `importPath`. Set equal to `importPath` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
	ExportPath interface{}
	// S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
	ImportPath interface{}
	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `importPath` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
	ImportedFileChunkSize interface{}
	// Set of Elastic Network Interface identifiers from which the file system is accessible.
	NetworkInterfaceIds interface{}
	// AWS account identifier that created the file system.
	OwnerId interface{}
	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds interface{}
	// The storage capacity (GiB) of the file system. Minimum of `3600`. Storage capacity is provisioned in increments of 3,600 GiB.
	StorageCapacity interface{}
	// A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
	SubnetIds interface{}
	// A mapping of tags to assign to the file system.
	Tags interface{}
	// Identifier of the Virtual Private Cloud for the file system.
	VpcId interface{}
	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime interface{}
}

// The set of arguments for constructing a LustreFileSystem resource.
type LustreFileSystemArgs struct {
	// S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `importPath` argument and the path must use the same Amazon S3 bucket as specified in `importPath`. Set equal to `importPath` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
	ExportPath interface{}
	// S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
	ImportPath interface{}
	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `importPath` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
	ImportedFileChunkSize interface{}
	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds interface{}
	// The storage capacity (GiB) of the file system. Minimum of `3600`. Storage capacity is provisioned in increments of 3,600 GiB.
	StorageCapacity interface{}
	// A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
	SubnetIds interface{}
	// A mapping of tags to assign to the file system.
	Tags interface{}
	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime interface{}
}
