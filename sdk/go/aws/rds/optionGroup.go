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
	s *pulumi.ResourceState
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
	inputs := make(map[string]interface{})
	inputs["optionGroupDescription"] = "Managed by Pulumi"
	if args == nil {
		inputs["engineName"] = nil
		inputs["majorEngineVersion"] = nil
		inputs["name"] = nil
		inputs["namePrefix"] = nil
		inputs["options"] = nil
		inputs["tags"] = nil
	} else {
		inputs["engineName"] = args.EngineName
		inputs["majorEngineVersion"] = args.MajorEngineVersion
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["options"] = args.Options
		inputs["optionGroupDescription"] = args.OptionGroupDescription
		inputs["tags"] = args.Tags
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:rds/optionGroup:OptionGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OptionGroup{s: s}, nil
}

// GetOptionGroup gets an existing OptionGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOptionGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *OptionGroupState, opts ...pulumi.ResourceOpt) (*OptionGroup, error) {
	inputs := make(map[string]interface{})
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
	s, err := ctx.ReadResource("aws:rds/optionGroup:OptionGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OptionGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *OptionGroup) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *OptionGroup) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the db option group.
func (r *OptionGroup) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// Specifies the name of the engine that this option group should be associated with.
func (r *OptionGroup) EngineName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["engineName"])
}

// Specifies the major version of the engine that this option group should be associated with.
func (r *OptionGroup) MajorEngineVersion() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["majorEngineVersion"])
}

// The Name of the setting.
func (r *OptionGroup) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
func (r *OptionGroup) NamePrefix() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["namePrefix"])
}

// A list of Options to apply.
func (r *OptionGroup) Options() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["options"])
}

// The description of the option group. Defaults to "Managed by Pulumi".
func (r *OptionGroup) OptionGroupDescription() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["optionGroupDescription"])
}

// A mapping of tags to assign to the resource.
func (r *OptionGroup) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering OptionGroup resources.
type OptionGroupState struct {
	// The ARN of the db option group.
	Arn interface{}
	// Specifies the name of the engine that this option group should be associated with.
	EngineName interface{}
	// Specifies the major version of the engine that this option group should be associated with.
	MajorEngineVersion interface{}
	// The Name of the setting.
	Name interface{}
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
	NamePrefix interface{}
	// A list of Options to apply.
	Options interface{}
	// The description of the option group. Defaults to "Managed by Pulumi".
	OptionGroupDescription interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a OptionGroup resource.
type OptionGroupArgs struct {
	// Specifies the name of the engine that this option group should be associated with.
	EngineName interface{}
	// Specifies the major version of the engine that this option group should be associated with.
	MajorEngineVersion interface{}
	// The Name of the setting.
	Name interface{}
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
	NamePrefix interface{}
	// A list of Options to apply.
	Options interface{}
	// The description of the option group. Defaults to "Managed by Pulumi".
	OptionGroupDescription interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
