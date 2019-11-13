// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS App Mesh virtual router resource.
// 
// ## Breaking Changes
// 
// Because of backward incompatible API changes (read [here](https://github.com/awslabs/aws-app-mesh-examples/issues/92) and [here](https://github.com/awslabs/aws-app-mesh-examples/issues/94)), `appmesh.VirtualRouter` resource definitions created with provider versions earlier than v2.3.0 will need to be modified:
// 
// * Remove service `serviceNames` from the `spec` argument.
// AWS has created a `appmesh.VirtualService` resource for each of service names.
// These resource can be imported using `import`.
// 
// * Add a `listener` configuration block to the `spec` argument.
// 
// The state associated with existing resources will automatically be migrated.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appmesh_virtual_router.html.markdown.
type VirtualRouter struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN of the virtual router.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The creation date of the virtual router.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`

	// The last update date of the virtual router.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`

	// The name of the service mesh in which to create the virtual router.
	MeshName pulumi.StringOutput `pulumi:"meshName"`

	// The name to use for the virtual router.
	Name pulumi.StringOutput `pulumi:"name"`

	// The virtual router specification to apply.
	Spec pulumi.AnyOutput `pulumi:"spec"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewVirtualRouter registers a new resource with the given unique name, arguments, and options.
func NewVirtualRouter(ctx *pulumi.Context,
	name string, args *VirtualRouterArgs, opts ...pulumi.ResourceOpt) (*VirtualRouter, error) {
	if args == nil || args.MeshName == nil {
		return nil, errors.New("missing required argument 'MeshName'")
	}
	if args == nil || args.Spec == nil {
		return nil, errors.New("missing required argument 'Spec'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["meshName"] = args.MeshName
		inputs["name"] = args.Name
		inputs["spec"] = args.Spec
		inputs["tags"] = args.Tags
	}
	var resource VirtualRouter
	err := ctx.RegisterResource("aws:appmesh/virtualRouter:VirtualRouter", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualRouter gets an existing VirtualRouter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualRouter(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VirtualRouterState, opts ...pulumi.ResourceOpt) (*VirtualRouter, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["createdDate"] = state.CreatedDate
		inputs["lastUpdatedDate"] = state.LastUpdatedDate
		inputs["meshName"] = state.MeshName
		inputs["name"] = state.Name
		inputs["spec"] = state.Spec
		inputs["tags"] = state.Tags
	}
	var resource VirtualRouter
	err := ctx.ReadResource("aws:appmesh/virtualRouter:VirtualRouter", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *VirtualRouter) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *VirtualRouter) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering VirtualRouter resources.
type VirtualRouterState struct {
	// The ARN of the virtual router.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The creation date of the virtual router.
	CreatedDate pulumi.StringInput `pulumi:"createdDate"`
	// The last update date of the virtual router.
	LastUpdatedDate pulumi.StringInput `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the virtual router.
	MeshName pulumi.StringInput `pulumi:"meshName"`
	// The name to use for the virtual router.
	Name pulumi.StringInput `pulumi:"name"`
	// The virtual router specification to apply.
	Spec pulumi.AnyInput `pulumi:"spec"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a VirtualRouter resource.
type VirtualRouterArgs struct {
	// The name of the service mesh in which to create the virtual router.
	MeshName pulumi.StringInput `pulumi:"meshName"`
	// The name to use for the virtual router.
	Name pulumi.StringInput `pulumi:"name"`
	// The virtual router specification to apply.
	Spec pulumi.AnyInput `pulumi:"spec"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
