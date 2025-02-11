// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a Snapshot of a snapshot.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ebs_snapshot_copy.html.markdown.
type SnapshotCopy struct {
	s *pulumi.ResourceState
}

// NewSnapshotCopy registers a new resource with the given unique name, arguments, and options.
func NewSnapshotCopy(ctx *pulumi.Context,
	name string, args *SnapshotCopyArgs, opts ...pulumi.ResourceOpt) (*SnapshotCopy, error) {
	if args == nil || args.SourceRegion == nil {
		return nil, errors.New("missing required argument 'SourceRegion'")
	}
	if args == nil || args.SourceSnapshotId == nil {
		return nil, errors.New("missing required argument 'SourceSnapshotId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["encrypted"] = nil
		inputs["kmsKeyId"] = nil
		inputs["sourceRegion"] = nil
		inputs["sourceSnapshotId"] = nil
		inputs["tags"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["encrypted"] = args.Encrypted
		inputs["kmsKeyId"] = args.KmsKeyId
		inputs["sourceRegion"] = args.SourceRegion
		inputs["sourceSnapshotId"] = args.SourceSnapshotId
		inputs["tags"] = args.Tags
	}
	inputs["dataEncryptionKeyId"] = nil
	inputs["ownerAlias"] = nil
	inputs["ownerId"] = nil
	inputs["volumeId"] = nil
	inputs["volumeSize"] = nil
	s, err := ctx.RegisterResource("aws:ebs/snapshotCopy:SnapshotCopy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SnapshotCopy{s: s}, nil
}

// GetSnapshotCopy gets an existing SnapshotCopy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotCopy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SnapshotCopyState, opts ...pulumi.ResourceOpt) (*SnapshotCopy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["dataEncryptionKeyId"] = state.DataEncryptionKeyId
		inputs["description"] = state.Description
		inputs["encrypted"] = state.Encrypted
		inputs["kmsKeyId"] = state.KmsKeyId
		inputs["ownerAlias"] = state.OwnerAlias
		inputs["ownerId"] = state.OwnerId
		inputs["sourceRegion"] = state.SourceRegion
		inputs["sourceSnapshotId"] = state.SourceSnapshotId
		inputs["tags"] = state.Tags
		inputs["volumeId"] = state.VolumeId
		inputs["volumeSize"] = state.VolumeSize
	}
	s, err := ctx.ReadResource("aws:ebs/snapshotCopy:SnapshotCopy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SnapshotCopy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SnapshotCopy) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SnapshotCopy) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The data encryption key identifier for the snapshot.
// * `sourceSnapshotId` The ARN of the copied snapshot.
// * `sourceRegion` The region of the source snapshot.
func (r *SnapshotCopy) DataEncryptionKeyId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dataEncryptionKeyId"])
}

// A description of what the snapshot is.
func (r *SnapshotCopy) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// Whether the snapshot is encrypted.
func (r *SnapshotCopy) Encrypted() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["encrypted"])
}

// The ARN for the KMS encryption key.
// * `sourceSnapshotId` The ARN for the snapshot to be copied.
// * `sourceRegion` The region of the source snapshot.
func (r *SnapshotCopy) KmsKeyId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["kmsKeyId"])
}

// Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
func (r *SnapshotCopy) OwnerAlias() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ownerAlias"])
}

// The AWS account ID of the snapshot owner.
func (r *SnapshotCopy) OwnerId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ownerId"])
}

func (r *SnapshotCopy) SourceRegion() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sourceRegion"])
}

func (r *SnapshotCopy) SourceSnapshotId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sourceSnapshotId"])
}

// A mapping of tags for the snapshot.
func (r *SnapshotCopy) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

func (r *SnapshotCopy) VolumeId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["volumeId"])
}

// The size of the drive in GiBs.
func (r *SnapshotCopy) VolumeSize() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["volumeSize"])
}

// Input properties used for looking up and filtering SnapshotCopy resources.
type SnapshotCopyState struct {
	// The data encryption key identifier for the snapshot.
	// * `sourceSnapshotId` The ARN of the copied snapshot.
	// * `sourceRegion` The region of the source snapshot.
	DataEncryptionKeyId interface{}
	// A description of what the snapshot is.
	Description interface{}
	// Whether the snapshot is encrypted.
	Encrypted interface{}
	// The ARN for the KMS encryption key.
	// * `sourceSnapshotId` The ARN for the snapshot to be copied.
	// * `sourceRegion` The region of the source snapshot.
	KmsKeyId interface{}
	// Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
	OwnerAlias interface{}
	// The AWS account ID of the snapshot owner.
	OwnerId interface{}
	SourceRegion interface{}
	SourceSnapshotId interface{}
	// A mapping of tags for the snapshot.
	Tags interface{}
	VolumeId interface{}
	// The size of the drive in GiBs.
	VolumeSize interface{}
}

// The set of arguments for constructing a SnapshotCopy resource.
type SnapshotCopyArgs struct {
	// A description of what the snapshot is.
	Description interface{}
	// Whether the snapshot is encrypted.
	Encrypted interface{}
	// The ARN for the KMS encryption key.
	// * `sourceSnapshotId` The ARN for the snapshot to be copied.
	// * `sourceRegion` The region of the source snapshot.
	KmsKeyId interface{}
	SourceRegion interface{}
	SourceSnapshotId interface{}
	// A mapping of tags for the snapshot.
	Tags interface{}
}
