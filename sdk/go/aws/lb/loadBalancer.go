// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lb

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Load Balancer resource.
// 
// > **Note:** `alb.LoadBalancer` is known as `lb.LoadBalancer`. The functionality is identical.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb.html.markdown.
type LoadBalancer struct {
	s *pulumi.ResourceState
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOpt) (*LoadBalancer, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accessLogs"] = nil
		inputs["enableCrossZoneLoadBalancing"] = nil
		inputs["enableDeletionProtection"] = nil
		inputs["enableHttp2"] = nil
		inputs["idleTimeout"] = nil
		inputs["internal"] = nil
		inputs["ipAddressType"] = nil
		inputs["loadBalancerType"] = nil
		inputs["name"] = nil
		inputs["namePrefix"] = nil
		inputs["securityGroups"] = nil
		inputs["subnetMappings"] = nil
		inputs["subnets"] = nil
		inputs["tags"] = nil
	} else {
		inputs["accessLogs"] = args.AccessLogs
		inputs["enableCrossZoneLoadBalancing"] = args.EnableCrossZoneLoadBalancing
		inputs["enableDeletionProtection"] = args.EnableDeletionProtection
		inputs["enableHttp2"] = args.EnableHttp2
		inputs["idleTimeout"] = args.IdleTimeout
		inputs["internal"] = args.Internal
		inputs["ipAddressType"] = args.IpAddressType
		inputs["loadBalancerType"] = args.LoadBalancerType
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["securityGroups"] = args.SecurityGroups
		inputs["subnetMappings"] = args.SubnetMappings
		inputs["subnets"] = args.Subnets
		inputs["tags"] = args.Tags
	}
	inputs["arn"] = nil
	inputs["arnSuffix"] = nil
	inputs["dnsName"] = nil
	inputs["vpcId"] = nil
	inputs["zoneId"] = nil
	s, err := ctx.RegisterResource("aws:lb/loadBalancer:LoadBalancer", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LoadBalancer{s: s}, nil
}

// GetLoadBalancer gets an existing LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LoadBalancerState, opts ...pulumi.ResourceOpt) (*LoadBalancer, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accessLogs"] = state.AccessLogs
		inputs["arn"] = state.Arn
		inputs["arnSuffix"] = state.ArnSuffix
		inputs["dnsName"] = state.DnsName
		inputs["enableCrossZoneLoadBalancing"] = state.EnableCrossZoneLoadBalancing
		inputs["enableDeletionProtection"] = state.EnableDeletionProtection
		inputs["enableHttp2"] = state.EnableHttp2
		inputs["idleTimeout"] = state.IdleTimeout
		inputs["internal"] = state.Internal
		inputs["ipAddressType"] = state.IpAddressType
		inputs["loadBalancerType"] = state.LoadBalancerType
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["securityGroups"] = state.SecurityGroups
		inputs["subnetMappings"] = state.SubnetMappings
		inputs["subnets"] = state.Subnets
		inputs["tags"] = state.Tags
		inputs["vpcId"] = state.VpcId
		inputs["zoneId"] = state.ZoneId
	}
	s, err := ctx.ReadResource("aws:lb/loadBalancer:LoadBalancer", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LoadBalancer{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *LoadBalancer) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *LoadBalancer) ID() pulumi.IDOutput {
	return r.s.ID()
}

// An Access Logs block. Access Logs documented below.
func (r *LoadBalancer) AccessLogs() pulumi.Output {
	return r.s.State["accessLogs"]
}

// The ARN of the load balancer (matches `id`).
func (r *LoadBalancer) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// The ARN suffix for use with CloudWatch Metrics.
func (r *LoadBalancer) ArnSuffix() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arnSuffix"])
}

// The DNS name of the load balancer.
func (r *LoadBalancer) DnsName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dnsName"])
}

// If true, cross-zone load balancing of the load balancer will be enabled.
// This is a `network` load balancer feature. Defaults to `false`.
func (r *LoadBalancer) EnableCrossZoneLoadBalancing() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enableCrossZoneLoadBalancing"])
}

// If true, deletion of the load balancer will be disabled via
// the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
func (r *LoadBalancer) EnableDeletionProtection() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enableDeletionProtection"])
}

// Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
func (r *LoadBalancer) EnableHttp2() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enableHttp2"])
}

// The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
func (r *LoadBalancer) IdleTimeout() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["idleTimeout"])
}

// If true, the LB will be internal.
func (r *LoadBalancer) Internal() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["internal"])
}

// The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
func (r *LoadBalancer) IpAddressType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ipAddressType"])
}

// The type of load balancer to create. Possible values are `application` or `network`. The default value is `application`.
func (r *LoadBalancer) LoadBalancerType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["loadBalancerType"])
}

// The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
// must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
// this provider will autogenerate a name beginning with `tf-lb`.
func (r *LoadBalancer) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
func (r *LoadBalancer) NamePrefix() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["namePrefix"])
}

// A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
func (r *LoadBalancer) SecurityGroups() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["securityGroups"])
}

// A subnet mapping block as documented below.
func (r *LoadBalancer) SubnetMappings() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["subnetMappings"])
}

// A list of subnet IDs to attach to the LB. Subnets
// cannot be updated for Load Balancers of type `network`. Changing this value
// for load balancers of type `network` will force a recreation of the resource.
func (r *LoadBalancer) Subnets() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["subnets"])
}

// A mapping of tags to assign to the resource.
func (r *LoadBalancer) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

func (r *LoadBalancer) VpcId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["vpcId"])
}

// The canonical hosted zone ID of the load balancer (to be used in a Route 53 Alias record).
func (r *LoadBalancer) ZoneId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["zoneId"])
}

// Input properties used for looking up and filtering LoadBalancer resources.
type LoadBalancerState struct {
	// An Access Logs block. Access Logs documented below.
	AccessLogs interface{}
	// The ARN of the load balancer (matches `id`).
	Arn interface{}
	// The ARN suffix for use with CloudWatch Metrics.
	ArnSuffix interface{}
	// The DNS name of the load balancer.
	DnsName interface{}
	// If true, cross-zone load balancing of the load balancer will be enabled.
	// This is a `network` load balancer feature. Defaults to `false`.
	EnableCrossZoneLoadBalancing interface{}
	// If true, deletion of the load balancer will be disabled via
	// the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
	EnableDeletionProtection interface{}
	// Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
	EnableHttp2 interface{}
	// The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
	IdleTimeout interface{}
	// If true, the LB will be internal.
	Internal interface{}
	// The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
	IpAddressType interface{}
	// The type of load balancer to create. Possible values are `application` or `network`. The default value is `application`.
	LoadBalancerType interface{}
	// The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
	// must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
	// this provider will autogenerate a name beginning with `tf-lb`.
	Name interface{}
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix interface{}
	// A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
	SecurityGroups interface{}
	// A subnet mapping block as documented below.
	SubnetMappings interface{}
	// A list of subnet IDs to attach to the LB. Subnets
	// cannot be updated for Load Balancers of type `network`. Changing this value
	// for load balancers of type `network` will force a recreation of the resource.
	Subnets interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	VpcId interface{}
	// The canonical hosted zone ID of the load balancer (to be used in a Route 53 Alias record).
	ZoneId interface{}
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// An Access Logs block. Access Logs documented below.
	AccessLogs interface{}
	// If true, cross-zone load balancing of the load balancer will be enabled.
	// This is a `network` load balancer feature. Defaults to `false`.
	EnableCrossZoneLoadBalancing interface{}
	// If true, deletion of the load balancer will be disabled via
	// the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
	EnableDeletionProtection interface{}
	// Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
	EnableHttp2 interface{}
	// The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
	IdleTimeout interface{}
	// If true, the LB will be internal.
	Internal interface{}
	// The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
	IpAddressType interface{}
	// The type of load balancer to create. Possible values are `application` or `network`. The default value is `application`.
	LoadBalancerType interface{}
	// The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
	// must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
	// this provider will autogenerate a name beginning with `tf-lb`.
	Name interface{}
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix interface{}
	// A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
	SecurityGroups interface{}
	// A subnet mapping block as documented below.
	SubnetMappings interface{}
	// A list of subnet IDs to attach to the LB. Subnets
	// cannot be updated for Load Balancers of type `network`. Changing this value
	// for load balancers of type `network` will force a recreation of the resource.
	Subnets interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
