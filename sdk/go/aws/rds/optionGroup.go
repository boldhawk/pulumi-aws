// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an RDS DB option group resource. Documentation of the available options for various RDS engines can be found at:
// * [MariaDB Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MariaDB.Options.html)
// * [Microsoft SQL Server Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.Options.html)
// * [MySQL Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MySQL.Options.html)
// * [Oracle Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.Options.html)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/db_option_group.html.markdown.
type OptionGroup struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN of the db option group.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// Specifies the name of the engine that this option group should be associated with.
	EngineName pulumi.StringOutput `pulumi:"engineName"`

	// Specifies the major version of the engine that this option group should be associated with.
	MajorEngineVersion pulumi.StringOutput `pulumi:"majorEngineVersion"`

	// The Name of the setting.
	Name pulumi.StringOutput `pulumi:"name"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`

	// A list of Options to apply.
	Options pulumi.ArrayOutput `pulumi:"options"`

	// The description of the option group. Defaults to "Managed by Pulumi".
	OptionGroupDescription pulumi.StringOutput `pulumi:"optionGroupDescription"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewOptionGroup registers a new resource with the given unique name, arguments, and options.
func NewOptionGroup(ctx *pulumi.Context,
	name string, args *OptionGroupArgs, opts ...pulumi.ResourceOpt) (*OptionGroup, error) {
	if args == nil || args.EngineName == nil {
		return nil, errors.New("missing required argument 'EngineName'")
	}
	if args == nil || args.MajorEngineVersion == nil {
		return nil, errors.New("missing required argument 'MajorEngineVersion'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	inputs["optionGroupDescription"] = pulumi.Any("Managed by Pulumi")
	if args != nil {
		inputs["engineName"] = args.EngineName
		inputs["majorEngineVersion"] = args.MajorEngineVersion
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["options"] = args.Options
		inputs["optionGroupDescription"] = args.OptionGroupDescription
		inputs["tags"] = args.Tags
	}
	var resource OptionGroup
	err := ctx.RegisterResource("aws:rds/optionGroup:OptionGroup", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOptionGroup gets an existing OptionGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOptionGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *OptionGroupState, opts ...pulumi.ResourceOpt) (*OptionGroup, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["engineName"] = state.EngineName
		inputs["majorEngineVersion"] = state.MajorEngineVersion
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["options"] = state.Options
		inputs["optionGroupDescription"] = state.OptionGroupDescription
		inputs["tags"] = state.Tags
	}
	var resource OptionGroup
	err := ctx.ReadResource("aws:rds/optionGroup:OptionGroup", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *OptionGroup) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *OptionGroup) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering OptionGroup resources.
type OptionGroupState struct {
	// The ARN of the db option group.
	Arn pulumi.StringInput `pulumi:"arn"`
	// Specifies the name of the engine that this option group should be associated with.
	EngineName pulumi.StringInput `pulumi:"engineName"`
	// Specifies the major version of the engine that this option group should be associated with.
	MajorEngineVersion pulumi.StringInput `pulumi:"majorEngineVersion"`
	// The Name of the setting.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// A list of Options to apply.
	Options pulumi.ArrayInput `pulumi:"options"`
	// The description of the option group. Defaults to "Managed by Pulumi".
	OptionGroupDescription pulumi.StringInput `pulumi:"optionGroupDescription"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a OptionGroup resource.
type OptionGroupArgs struct {
	// Specifies the name of the engine that this option group should be associated with.
	EngineName pulumi.StringInput `pulumi:"engineName"`
	// Specifies the major version of the engine that this option group should be associated with.
	MajorEngineVersion pulumi.StringInput `pulumi:"majorEngineVersion"`
	// The Name of the setting.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// A list of Options to apply.
	Options pulumi.ArrayInput `pulumi:"options"`
	// The description of the option group. Defaults to "Managed by Pulumi".
	OptionGroupDescription pulumi.StringInput `pulumi:"optionGroupDescription"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
