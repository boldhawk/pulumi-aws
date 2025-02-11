// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The "AMI copy" resource allows duplication of an Amazon Machine Image (AMI),
// including cross-region copies.
// 
// If the source AMI has associated EBS snapshots, those will also be duplicated
// along with the AMI.
// 
// This is useful for taking a single AMI provisioned in one region and making
// it available in another for a multi-region deployment.
// 
// Copying an AMI can take several minutes. The creation of this resource will
// block until the new AMI is available for use on new instances.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ami_copy.html.markdown.
type AmiCopy struct {
	s *pulumi.ResourceState
}

// NewAmiCopy registers a new resource with the given unique name, arguments, and options.
func NewAmiCopy(ctx *pulumi.Context,
	name string, args *AmiCopyArgs, opts ...pulumi.ResourceOpt) (*AmiCopy, error) {
	if args == nil || args.SourceAmiId == nil {
		return nil, errors.New("missing required argument 'SourceAmiId'")
	}
	if args == nil || args.SourceAmiRegion == nil {
		return nil, errors.New("missing required argument 'SourceAmiRegion'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["ebsBlockDevices"] = nil
		inputs["encrypted"] = nil
		inputs["ephemeralBlockDevices"] = nil
		inputs["kmsKeyId"] = nil
		inputs["name"] = nil
		inputs["sourceAmiId"] = nil
		inputs["sourceAmiRegion"] = nil
		inputs["tags"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["ebsBlockDevices"] = args.EbsBlockDevices
		inputs["encrypted"] = args.Encrypted
		inputs["ephemeralBlockDevices"] = args.EphemeralBlockDevices
		inputs["kmsKeyId"] = args.KmsKeyId
		inputs["name"] = args.Name
		inputs["sourceAmiId"] = args.SourceAmiId
		inputs["sourceAmiRegion"] = args.SourceAmiRegion
		inputs["tags"] = args.Tags
	}
	inputs["architecture"] = nil
	inputs["enaSupport"] = nil
	inputs["imageLocation"] = nil
	inputs["kernelId"] = nil
	inputs["manageEbsSnapshots"] = nil
	inputs["ramdiskId"] = nil
	inputs["rootDeviceName"] = nil
	inputs["rootSnapshotId"] = nil
	inputs["sriovNetSupport"] = nil
	inputs["virtualizationType"] = nil
	s, err := ctx.RegisterResource("aws:ec2/amiCopy:AmiCopy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AmiCopy{s: s}, nil
}

// GetAmiCopy gets an existing AmiCopy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAmiCopy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AmiCopyState, opts ...pulumi.ResourceOpt) (*AmiCopy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["architecture"] = state.Architecture
		inputs["description"] = state.Description
		inputs["ebsBlockDevices"] = state.EbsBlockDevices
		inputs["enaSupport"] = state.EnaSupport
		inputs["encrypted"] = state.Encrypted
		inputs["ephemeralBlockDevices"] = state.EphemeralBlockDevices
		inputs["imageLocation"] = state.ImageLocation
		inputs["kernelId"] = state.KernelId
		inputs["kmsKeyId"] = state.KmsKeyId
		inputs["manageEbsSnapshots"] = state.ManageEbsSnapshots
		inputs["name"] = state.Name
		inputs["ramdiskId"] = state.RamdiskId
		inputs["rootDeviceName"] = state.RootDeviceName
		inputs["rootSnapshotId"] = state.RootSnapshotId
		inputs["sourceAmiId"] = state.SourceAmiId
		inputs["sourceAmiRegion"] = state.SourceAmiRegion
		inputs["sriovNetSupport"] = state.SriovNetSupport
		inputs["tags"] = state.Tags
		inputs["virtualizationType"] = state.VirtualizationType
	}
	s, err := ctx.ReadResource("aws:ec2/amiCopy:AmiCopy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AmiCopy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AmiCopy) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AmiCopy) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Machine architecture for created instances. Defaults to "x8664".
func (r *AmiCopy) Architecture() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["architecture"])
}

// A longer, human-readable description for the AMI.
func (r *AmiCopy) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// Nested block describing an EBS block device that should be
// attached to created instances. The structure of this block is described below.
func (r *AmiCopy) EbsBlockDevices() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["ebsBlockDevices"])
}

// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
func (r *AmiCopy) EnaSupport() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enaSupport"])
}

// Specifies whether the destination snapshots of the copied image should be encrypted. Defaults to `false`
func (r *AmiCopy) Encrypted() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["encrypted"])
}

// Nested block describing an ephemeral block device that
// should be attached to created instances. The structure of this block is described below.
func (r *AmiCopy) EphemeralBlockDevices() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["ephemeralBlockDevices"])
}

// Path to an S3 object containing an image manifest, e.g. created
// by the `ec2-upload-bundle` command in the EC2 command line tools.
func (r *AmiCopy) ImageLocation() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["imageLocation"])
}

// The id of the kernel image (AKI) that will be used as the paravirtual
// kernel in created instances.
func (r *AmiCopy) KernelId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["kernelId"])
}

// The full ARN of the KMS Key to use when encrypting the snapshots of an image during a copy operation. If not specified, then the default AWS KMS Key will be used
func (r *AmiCopy) KmsKeyId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["kmsKeyId"])
}

func (r *AmiCopy) ManageEbsSnapshots() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["manageEbsSnapshots"])
}

// A region-unique name for the AMI.
func (r *AmiCopy) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The id of an initrd image (ARI) that will be used when booting the
// created instances.
func (r *AmiCopy) RamdiskId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ramdiskId"])
}

// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
func (r *AmiCopy) RootDeviceName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["rootDeviceName"])
}

func (r *AmiCopy) RootSnapshotId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["rootSnapshotId"])
}

// The id of the AMI to copy. This id must be valid in the region
// given by `sourceAmiRegion`.
func (r *AmiCopy) SourceAmiId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sourceAmiId"])
}

// The region from which the AMI will be copied. This may be the
// same as the AWS provider region in order to create a copy within the same region.
func (r *AmiCopy) SourceAmiRegion() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sourceAmiRegion"])
}

// When set to "simple" (the default), enables enhanced networking
// for created instances. No other value is supported at this time.
func (r *AmiCopy) SriovNetSupport() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sriovNetSupport"])
}

// A mapping of tags to assign to the resource.
func (r *AmiCopy) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Keyword to choose what virtualization mode created instances
// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
// changes the set of further arguments that are required, as described below.
func (r *AmiCopy) VirtualizationType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["virtualizationType"])
}

// Input properties used for looking up and filtering AmiCopy resources.
type AmiCopyState struct {
	// Machine architecture for created instances. Defaults to "x8664".
	Architecture interface{}
	// A longer, human-readable description for the AMI.
	Description interface{}
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices interface{}
	// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport interface{}
	// Specifies whether the destination snapshots of the copied image should be encrypted. Defaults to `false`
	Encrypted interface{}
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices interface{}
	// Path to an S3 object containing an image manifest, e.g. created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation interface{}
	// The id of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId interface{}
	// The full ARN of the KMS Key to use when encrypting the snapshots of an image during a copy operation. If not specified, then the default AWS KMS Key will be used
	KmsKeyId interface{}
	ManageEbsSnapshots interface{}
	// A region-unique name for the AMI.
	Name interface{}
	// The id of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId interface{}
	// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName interface{}
	RootSnapshotId interface{}
	// The id of the AMI to copy. This id must be valid in the region
	// given by `sourceAmiRegion`.
	SourceAmiId interface{}
	// The region from which the AMI will be copied. This may be the
	// same as the AWS provider region in order to create a copy within the same region.
	SourceAmiRegion interface{}
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType interface{}
}

// The set of arguments for constructing a AmiCopy resource.
type AmiCopyArgs struct {
	// A longer, human-readable description for the AMI.
	Description interface{}
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices interface{}
	// Specifies whether the destination snapshots of the copied image should be encrypted. Defaults to `false`
	Encrypted interface{}
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices interface{}
	// The full ARN of the KMS Key to use when encrypting the snapshots of an image during a copy operation. If not specified, then the default AWS KMS Key will be used
	KmsKeyId interface{}
	// A region-unique name for the AMI.
	Name interface{}
	// The id of the AMI to copy. This id must be valid in the region
	// given by `sourceAmiRegion`.
	SourceAmiId interface{}
	// The region from which the AMI will be copied. This may be the
	// same as the AWS provider region in order to create a copy within the same region.
	SourceAmiRegion interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
