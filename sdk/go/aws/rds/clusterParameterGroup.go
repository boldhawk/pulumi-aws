// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an RDS DB cluster parameter group resource. Documentation of the available parameters for various Aurora engines can be found at:
// 
// * [Aurora MySQL Parameters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Reference.html)
// * [Aurora PostgreSQL Parameters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraPostgreSQL.Reference.html)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/rds_cluster_parameter_group.html.markdown.
type ClusterParameterGroup struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN of the db cluster parameter group.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The description of the DB cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`

	// The family of the DB cluster parameter group.
	Family pulumi.StringOutput `pulumi:"family"`

	// The name of the DB parameter.
	Name pulumi.StringOutput `pulumi:"name"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`

	// A list of DB parameters to apply. Note that parameters may differ from a family to an other. Full list of all parameters can be discovered via [`aws rds describe-db-cluster-parameters`](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-cluster-parameters.html) after initial creation of the group.
	Parameters pulumi.ArrayOutput `pulumi:"parameters"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewClusterParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewClusterParameterGroup(ctx *pulumi.Context,
	name string, args *ClusterParameterGroupArgs, opts ...pulumi.ResourceOpt) (*ClusterParameterGroup, error) {
	if args == nil || args.Family == nil {
		return nil, errors.New("missing required argument 'Family'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["description"] = pulumi.Any("Managed by Pulumi")
	if args != nil {
		inputs["description"] = args.Description
		inputs["family"] = args.Family
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["parameters"] = args.Parameters
		inputs["tags"] = args.Tags
	}
	var resource ClusterParameterGroup
	err := ctx.RegisterResource("aws:rds/clusterParameterGroup:ClusterParameterGroup", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterParameterGroup gets an existing ClusterParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ClusterParameterGroupState, opts ...pulumi.ResourceOpt) (*ClusterParameterGroup, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["description"] = state.Description
		inputs["family"] = state.Family
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["parameters"] = state.Parameters
		inputs["tags"] = state.Tags
	}
	var resource ClusterParameterGroup
	err := ctx.ReadResource("aws:rds/clusterParameterGroup:ClusterParameterGroup", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *ClusterParameterGroup) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *ClusterParameterGroup) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering ClusterParameterGroup resources.
type ClusterParameterGroupState struct {
	// The ARN of the db cluster parameter group.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The description of the DB cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringInput `pulumi:"description"`
	// The family of the DB cluster parameter group.
	Family pulumi.StringInput `pulumi:"family"`
	// The name of the DB parameter.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// A list of DB parameters to apply. Note that parameters may differ from a family to an other. Full list of all parameters can be discovered via [`aws rds describe-db-cluster-parameters`](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-cluster-parameters.html) after initial creation of the group.
	Parameters pulumi.ArrayInput `pulumi:"parameters"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterParameterGroup resource.
type ClusterParameterGroupArgs struct {
	// The description of the DB cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringInput `pulumi:"description"`
	// The family of the DB cluster parameter group.
	Family pulumi.StringInput `pulumi:"family"`
	// The name of the DB parameter.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// A list of DB parameters to apply. Note that parameters may differ from a family to an other. Full list of all parameters can be discovered via [`aws rds describe-db-cluster-parameters`](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-cluster-parameters.html) after initial creation of the group.
	Parameters pulumi.ArrayInput `pulumi:"parameters"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
