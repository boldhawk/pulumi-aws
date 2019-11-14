// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS App Mesh virtual node resource.
// 
// ## Breaking Changes
// 
// Because of backward incompatible API changes (read [here](https://github.com/awslabs/aws-app-mesh-examples/issues/92)), `appmesh.VirtualNode` resource definitions created with provider versions earlier than v2.3.0 will need to be modified:
// 
// * Rename the `serviceName` attribute of the `dns` object to `hostname`.
// 
// * Replace the `backends` attribute of the `spec` object with one or more `backend` configuration blocks,
// setting `virtualServiceName` to the name of the service.
// 
// The state associated with existing resources will automatically be migrated.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appmesh_virtual_node.html.markdown.
type VirtualNode struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN of the virtual node.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The creation date of the virtual node.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`

	// The last update date of the virtual node.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`

	// The name of the service mesh in which to create the virtual node.
	MeshName pulumi.StringOutput `pulumi:"meshName"`

	// The name to use for the virtual node.
	Name pulumi.StringOutput `pulumi:"name"`

	// The virtual node specification to apply.
	Spec pulumi.AnyOutput `pulumi:"spec"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewVirtualNode registers a new resource with the given unique name, arguments, and options.
func NewVirtualNode(ctx *pulumi.Context,
	name string, args *VirtualNodeArgs, opts ...pulumi.ResourceOpt) (*VirtualNode, error) {
	if args == nil || args.MeshName == nil {
		return nil, errors.New("missing required argument 'MeshName'")
	}
	if args == nil || args.Spec == nil {
		return nil, errors.New("missing required argument 'Spec'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["meshName"] = args.MeshName
		inputs["name"] = args.Name
		inputs["spec"] = args.Spec
		inputs["tags"] = args.Tags
	}
	var resource VirtualNode
	err := ctx.RegisterResource("aws:appmesh/virtualNode:VirtualNode", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNode gets an existing VirtualNode resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNode(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VirtualNodeState, opts ...pulumi.ResourceOpt) (*VirtualNode, error) {
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
	var resource VirtualNode
	err := ctx.ReadResource("aws:appmesh/virtualNode:VirtualNode", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *VirtualNode) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *VirtualNode) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering VirtualNode resources.
type VirtualNodeState struct {
	// The ARN of the virtual node.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The creation date of the virtual node.
	CreatedDate pulumi.StringInput `pulumi:"createdDate"`
	// The last update date of the virtual node.
	LastUpdatedDate pulumi.StringInput `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the virtual node.
	MeshName pulumi.StringInput `pulumi:"meshName"`
	// The name to use for the virtual node.
	Name pulumi.StringInput `pulumi:"name"`
	// The virtual node specification to apply.
	Spec pulumi.AnyInput `pulumi:"spec"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a VirtualNode resource.
type VirtualNodeArgs struct {
	// The name of the service mesh in which to create the virtual node.
	MeshName pulumi.StringInput `pulumi:"meshName"`
	// The name to use for the virtual node.
	Name pulumi.StringInput `pulumi:"name"`
	// The virtual node specification to apply.
	Spec pulumi.AnyInput `pulumi:"spec"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
