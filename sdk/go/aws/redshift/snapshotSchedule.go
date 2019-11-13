// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type SnapshotSchedule struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	Arn pulumi.StringOutput `pulumi:"arn"`

	// The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
	Definitions pulumi.ArrayOutput `pulumi:"definitions"`

	// The description of the snapshot schedule.
	Description pulumi.StringOutput `pulumi:"description"`

	// Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
	ForceDestroy pulumi.BoolOutput `pulumi:"forceDestroy"`

	Identifier pulumi.StringOutput `pulumi:"identifier"`

	// Creates a unique
	// identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix pulumi.StringOutput `pulumi:"identifierPrefix"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewSnapshotSchedule registers a new resource with the given unique name, arguments, and options.
func NewSnapshotSchedule(ctx *pulumi.Context,
	name string, args *SnapshotScheduleArgs, opts ...pulumi.ResourceOpt) (*SnapshotSchedule, error) {
	if args == nil || args.Definitions == nil {
		return nil, errors.New("missing required argument 'Definitions'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["definitions"] = args.Definitions
		inputs["description"] = args.Description
		inputs["forceDestroy"] = args.ForceDestroy
		inputs["identifier"] = args.Identifier
		inputs["identifierPrefix"] = args.IdentifierPrefix
		inputs["tags"] = args.Tags
	}
	var resource SnapshotSchedule
	err := ctx.RegisterResource("aws:redshift/snapshotSchedule:SnapshotSchedule", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotSchedule gets an existing SnapshotSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotSchedule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SnapshotScheduleState, opts ...pulumi.ResourceOpt) (*SnapshotSchedule, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["definitions"] = state.Definitions
		inputs["description"] = state.Description
		inputs["forceDestroy"] = state.ForceDestroy
		inputs["identifier"] = state.Identifier
		inputs["identifierPrefix"] = state.IdentifierPrefix
		inputs["tags"] = state.Tags
	}
	var resource SnapshotSchedule
	err := ctx.ReadResource("aws:redshift/snapshotSchedule:SnapshotSchedule", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *SnapshotSchedule) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *SnapshotSchedule) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering SnapshotSchedule resources.
type SnapshotScheduleState struct {
	Arn pulumi.StringInput `pulumi:"arn"`
	// The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
	Definitions pulumi.ArrayInput `pulumi:"definitions"`
	// The description of the snapshot schedule.
	Description pulumi.StringInput `pulumi:"description"`
	// Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
	ForceDestroy pulumi.BoolInput `pulumi:"forceDestroy"`
	Identifier pulumi.StringInput `pulumi:"identifier"`
	// Creates a unique
	// identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix pulumi.StringInput `pulumi:"identifierPrefix"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a SnapshotSchedule resource.
type SnapshotScheduleArgs struct {
	// The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 *)` or `rate(12 hours)`.
	Definitions pulumi.ArrayInput `pulumi:"definitions"`
	// The description of the snapshot schedule.
	Description pulumi.StringInput `pulumi:"description"`
	// Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
	ForceDestroy pulumi.BoolInput `pulumi:"forceDestroy"`
	Identifier pulumi.StringInput `pulumi:"identifier"`
	// Creates a unique
	// identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix pulumi.StringInput `pulumi:"identifierPrefix"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
