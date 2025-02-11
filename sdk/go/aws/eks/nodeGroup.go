// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an EKS Node Group, which can provision and optionally update an Auto Scaling Group of Kubernetes worker nodes compatible with EKS. Additional documentation about this functionality can be found in the [EKS User Guide](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/eks_node_group.html.markdown.
type NodeGroup struct {
	s *pulumi.ResourceState
}

// NewNodeGroup registers a new resource with the given unique name, arguments, and options.
func NewNodeGroup(ctx *pulumi.Context,
	name string, args *NodeGroupArgs, opts ...pulumi.ResourceOpt) (*NodeGroup, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.NodeGroupName == nil {
		return nil, errors.New("missing required argument 'NodeGroupName'")
	}
	if args == nil || args.NodeRoleArn == nil {
		return nil, errors.New("missing required argument 'NodeRoleArn'")
	}
	if args == nil || args.ScalingConfig == nil {
		return nil, errors.New("missing required argument 'ScalingConfig'")
	}
	if args == nil || args.SubnetIds == nil {
		return nil, errors.New("missing required argument 'SubnetIds'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["amiType"] = nil
		inputs["clusterName"] = nil
		inputs["diskSize"] = nil
		inputs["instanceTypes"] = nil
		inputs["labels"] = nil
		inputs["nodeGroupName"] = nil
		inputs["nodeRoleArn"] = nil
		inputs["releaseVersion"] = nil
		inputs["remoteAccess"] = nil
		inputs["scalingConfig"] = nil
		inputs["subnetIds"] = nil
		inputs["tags"] = nil
		inputs["version"] = nil
	} else {
		inputs["amiType"] = args.AmiType
		inputs["clusterName"] = args.ClusterName
		inputs["diskSize"] = args.DiskSize
		inputs["instanceTypes"] = args.InstanceTypes
		inputs["labels"] = args.Labels
		inputs["nodeGroupName"] = args.NodeGroupName
		inputs["nodeRoleArn"] = args.NodeRoleArn
		inputs["releaseVersion"] = args.ReleaseVersion
		inputs["remoteAccess"] = args.RemoteAccess
		inputs["scalingConfig"] = args.ScalingConfig
		inputs["subnetIds"] = args.SubnetIds
		inputs["tags"] = args.Tags
		inputs["version"] = args.Version
	}
	inputs["arn"] = nil
	inputs["resources"] = nil
	inputs["status"] = nil
	s, err := ctx.RegisterResource("aws:eks/nodeGroup:NodeGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NodeGroup{s: s}, nil
}

// GetNodeGroup gets an existing NodeGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodeGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NodeGroupState, opts ...pulumi.ResourceOpt) (*NodeGroup, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["amiType"] = state.AmiType
		inputs["arn"] = state.Arn
		inputs["clusterName"] = state.ClusterName
		inputs["diskSize"] = state.DiskSize
		inputs["instanceTypes"] = state.InstanceTypes
		inputs["labels"] = state.Labels
		inputs["nodeGroupName"] = state.NodeGroupName
		inputs["nodeRoleArn"] = state.NodeRoleArn
		inputs["releaseVersion"] = state.ReleaseVersion
		inputs["remoteAccess"] = state.RemoteAccess
		inputs["resources"] = state.Resources
		inputs["scalingConfig"] = state.ScalingConfig
		inputs["status"] = state.Status
		inputs["subnetIds"] = state.SubnetIds
		inputs["tags"] = state.Tags
		inputs["version"] = state.Version
	}
	s, err := ctx.ReadResource("aws:eks/nodeGroup:NodeGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NodeGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *NodeGroup) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *NodeGroup) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *NodeGroup) AmiType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["amiType"])
}

// Amazon Resource Name (ARN) of the EKS Node Group.
func (r *NodeGroup) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// Name of the EKS Cluster.
func (r *NodeGroup) ClusterName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clusterName"])
}

func (r *NodeGroup) DiskSize() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["diskSize"])
}

func (r *NodeGroup) InstanceTypes() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["instanceTypes"])
}

// Key-value mapping of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
func (r *NodeGroup) Labels() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["labels"])
}

// Name of the EKS Node Group.
func (r *NodeGroup) NodeGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["nodeGroupName"])
}

// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
func (r *NodeGroup) NodeRoleArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["nodeRoleArn"])
}

// AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
func (r *NodeGroup) ReleaseVersion() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["releaseVersion"])
}

// Configuration block with remote access settings. Detailed below.
func (r *NodeGroup) RemoteAccess() pulumi.Output {
	return r.s.State["remoteAccess"]
}

// List of objects containing information about underlying resources.
func (r *NodeGroup) Resources() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["resources"])
}

// Configuration block with scaling settings. Detailed below.
func (r *NodeGroup) ScalingConfig() pulumi.Output {
	return r.s.State["scalingConfig"]
}

// Status of the EKS Node Group.
func (r *NodeGroup) Status() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["status"])
}

// Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
func (r *NodeGroup) SubnetIds() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["subnetIds"])
}

// Key-value mapping of resource tags.
func (r *NodeGroup) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

func (r *NodeGroup) Version() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["version"])
}

// Input properties used for looking up and filtering NodeGroup resources.
type NodeGroupState struct {
	AmiType interface{}
	// Amazon Resource Name (ARN) of the EKS Node Group.
	Arn interface{}
	// Name of the EKS Cluster.
	ClusterName interface{}
	DiskSize interface{}
	InstanceTypes interface{}
	// Key-value mapping of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
	Labels interface{}
	// Name of the EKS Node Group.
	NodeGroupName interface{}
	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
	NodeRoleArn interface{}
	// AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
	ReleaseVersion interface{}
	// Configuration block with remote access settings. Detailed below.
	RemoteAccess interface{}
	// List of objects containing information about underlying resources.
	Resources interface{}
	// Configuration block with scaling settings. Detailed below.
	ScalingConfig interface{}
	// Status of the EKS Node Group.
	Status interface{}
	// Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
	SubnetIds interface{}
	// Key-value mapping of resource tags.
	Tags interface{}
	Version interface{}
}

// The set of arguments for constructing a NodeGroup resource.
type NodeGroupArgs struct {
	AmiType interface{}
	// Name of the EKS Cluster.
	ClusterName interface{}
	DiskSize interface{}
	InstanceTypes interface{}
	// Key-value mapping of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
	Labels interface{}
	// Name of the EKS Node Group.
	NodeGroupName interface{}
	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
	NodeRoleArn interface{}
	// AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
	ReleaseVersion interface{}
	// Configuration block with remote access settings. Detailed below.
	RemoteAccess interface{}
	// Configuration block with scaling settings. Detailed below.
	ScalingConfig interface{}
	// Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
	SubnetIds interface{}
	// Key-value mapping of resource tags.
	Tags interface{}
	Version interface{}
}
