// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS App Mesh service mesh resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appmesh_mesh.html.markdown.
type Mesh struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN of the service mesh.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The creation date of the service mesh.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`

	// The last update date of the service mesh.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`

	// The name to use for the service mesh.
	Name pulumi.StringOutput `pulumi:"name"`

	// The service mesh specification to apply.
	Spec pulumi.AnyOutput `pulumi:"spec"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewMesh registers a new resource with the given unique name, arguments, and options.
func NewMesh(ctx *pulumi.Context,
	name string, args *MeshArgs, opts ...pulumi.ResourceOpt) (*Mesh, error) {
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["name"] = args.Name
		inputs["spec"] = args.Spec
		inputs["tags"] = args.Tags
	}
	var resource Mesh
	err := ctx.RegisterResource("aws:appmesh/mesh:Mesh", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMesh gets an existing Mesh resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMesh(ctx *pulumi.Context,
	name string, id pulumi.ID, state *MeshState, opts ...pulumi.ResourceOpt) (*Mesh, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["createdDate"] = state.CreatedDate
		inputs["lastUpdatedDate"] = state.LastUpdatedDate
		inputs["name"] = state.Name
		inputs["spec"] = state.Spec
		inputs["tags"] = state.Tags
	}
	var resource Mesh
	err := ctx.ReadResource("aws:appmesh/mesh:Mesh", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *Mesh) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *Mesh) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering Mesh resources.
type MeshState struct {
	// The ARN of the service mesh.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The creation date of the service mesh.
	CreatedDate pulumi.StringInput `pulumi:"createdDate"`
	// The last update date of the service mesh.
	LastUpdatedDate pulumi.StringInput `pulumi:"lastUpdatedDate"`
	// The name to use for the service mesh.
	Name pulumi.StringInput `pulumi:"name"`
	// The service mesh specification to apply.
	Spec pulumi.AnyInput `pulumi:"spec"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a Mesh resource.
type MeshArgs struct {
	// The name to use for the service mesh.
	Name pulumi.StringInput `pulumi:"name"`
	// The service mesh specification to apply.
	Spec pulumi.AnyInput `pulumi:"spec"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
