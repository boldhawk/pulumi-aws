// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Load Balancer resource.
// 
// > **Note:** `alb.LoadBalancer` is known as `lb.LoadBalancer`. The functionality is identical.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_legacy.html.markdown.
type LoadBalancer struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// An Access Logs block. Access Logs documented below.
	AccessLogs pulumi.AnyOutput `pulumi:"accessLogs"`

	// The ARN of the load balancer (matches `id`).
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The ARN suffix for use with CloudWatch Metrics.
	ArnSuffix pulumi.StringOutput `pulumi:"arnSuffix"`

	// The DNS name of the load balancer.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`

	// If true, cross-zone load balancing of the load balancer will be enabled.
	// This is a `network` load balancer feature. Defaults to `false`.
	EnableCrossZoneLoadBalancing pulumi.BoolOutput `pulumi:"enableCrossZoneLoadBalancing"`

	// If true, deletion of the load balancer will be disabled via
	// the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
	EnableDeletionProtection pulumi.BoolOutput `pulumi:"enableDeletionProtection"`

	// Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
	EnableHttp2 pulumi.BoolOutput `pulumi:"enableHttp2"`

	// The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
	IdleTimeout pulumi.IntOutput `pulumi:"idleTimeout"`

	// If true, the LB will be internal.
	Internal pulumi.BoolOutput `pulumi:"internal"`

	// The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
	IpAddressType pulumi.StringOutput `pulumi:"ipAddressType"`

	// The type of load balancer to create. Possible values are `application` or `network`. The default value is `application`.
	LoadBalancerType pulumi.StringOutput `pulumi:"loadBalancerType"`

	// The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
	// must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
	// this provider will autogenerate a name beginning with `tf-lb`.
	Name pulumi.StringOutput `pulumi:"name"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`

	// A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
	SecurityGroups pulumi.ArrayOutput `pulumi:"securityGroups"`

	// A subnet mapping block as documented below.
	SubnetMappings pulumi.ArrayOutput `pulumi:"subnetMappings"`

	// A list of subnet IDs to attach to the LB. Subnets
	// cannot be updated for Load Balancers of type `network`. Changing this value
	// for load balancers of type `network` will force a recreation of the resource.
	Subnets pulumi.ArrayOutput `pulumi:"subnets"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	VpcId pulumi.StringOutput `pulumi:"vpcId"`

	// The canonical hosted zone ID of the load balancer (to be used in a Route 53 Alias record).
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOpt) (*LoadBalancer, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
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
	var resource LoadBalancer
	err := ctx.RegisterResource("aws:elasticloadbalancingv2/loadBalancer:LoadBalancer", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancer gets an existing LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LoadBalancerState, opts ...pulumi.ResourceOpt) (*LoadBalancer, error) {
	inputs := map[string]pulumi.Input{}
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
	var resource LoadBalancer
	err := ctx.ReadResource("aws:elasticloadbalancingv2/loadBalancer:LoadBalancer", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *LoadBalancer) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *LoadBalancer) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering LoadBalancer resources.
type LoadBalancerState struct {
	// An Access Logs block. Access Logs documented below.
	AccessLogs pulumi.AnyInput `pulumi:"accessLogs"`
	// The ARN of the load balancer (matches `id`).
	Arn pulumi.StringInput `pulumi:"arn"`
	// The ARN suffix for use with CloudWatch Metrics.
	ArnSuffix pulumi.StringInput `pulumi:"arnSuffix"`
	// The DNS name of the load balancer.
	DnsName pulumi.StringInput `pulumi:"dnsName"`
	// If true, cross-zone load balancing of the load balancer will be enabled.
	// This is a `network` load balancer feature. Defaults to `false`.
	EnableCrossZoneLoadBalancing pulumi.BoolInput `pulumi:"enableCrossZoneLoadBalancing"`
	// If true, deletion of the load balancer will be disabled via
	// the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
	EnableDeletionProtection pulumi.BoolInput `pulumi:"enableDeletionProtection"`
	// Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
	EnableHttp2 pulumi.BoolInput `pulumi:"enableHttp2"`
	// The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
	IdleTimeout pulumi.IntInput `pulumi:"idleTimeout"`
	// If true, the LB will be internal.
	Internal pulumi.BoolInput `pulumi:"internal"`
	// The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
	IpAddressType pulumi.StringInput `pulumi:"ipAddressType"`
	// The type of load balancer to create. Possible values are `application` or `network`. The default value is `application`.
	LoadBalancerType pulumi.StringInput `pulumi:"loadBalancerType"`
	// The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
	// must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
	// this provider will autogenerate a name beginning with `tf-lb`.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
	SecurityGroups pulumi.ArrayInput `pulumi:"securityGroups"`
	// A subnet mapping block as documented below.
	SubnetMappings pulumi.ArrayInput `pulumi:"subnetMappings"`
	// A list of subnet IDs to attach to the LB. Subnets
	// cannot be updated for Load Balancers of type `network`. Changing this value
	// for load balancers of type `network` will force a recreation of the resource.
	Subnets pulumi.ArrayInput `pulumi:"subnets"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	VpcId pulumi.StringInput `pulumi:"vpcId"`
	// The canonical hosted zone ID of the load balancer (to be used in a Route 53 Alias record).
	ZoneId pulumi.StringInput `pulumi:"zoneId"`
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// An Access Logs block. Access Logs documented below.
	AccessLogs pulumi.AnyInput `pulumi:"accessLogs"`
	// If true, cross-zone load balancing of the load balancer will be enabled.
	// This is a `network` load balancer feature. Defaults to `false`.
	EnableCrossZoneLoadBalancing pulumi.BoolInput `pulumi:"enableCrossZoneLoadBalancing"`
	// If true, deletion of the load balancer will be disabled via
	// the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
	EnableDeletionProtection pulumi.BoolInput `pulumi:"enableDeletionProtection"`
	// Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
	EnableHttp2 pulumi.BoolInput `pulumi:"enableHttp2"`
	// The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
	IdleTimeout pulumi.IntInput `pulumi:"idleTimeout"`
	// If true, the LB will be internal.
	Internal pulumi.BoolInput `pulumi:"internal"`
	// The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
	IpAddressType pulumi.StringInput `pulumi:"ipAddressType"`
	// The type of load balancer to create. Possible values are `application` or `network`. The default value is `application`.
	LoadBalancerType pulumi.StringInput `pulumi:"loadBalancerType"`
	// The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
	// must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
	// this provider will autogenerate a name beginning with `tf-lb`.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
	SecurityGroups pulumi.ArrayInput `pulumi:"securityGroups"`
	// A subnet mapping block as documented below.
	SubnetMappings pulumi.ArrayInput `pulumi:"subnetMappings"`
	// A list of subnet IDs to attach to the LB. Subnets
	// cannot be updated for Load Balancers of type `network`. Changing this value
	// for load balancers of type `network` will force a recreation of the resource.
	Subnets pulumi.ArrayInput `pulumi:"subnets"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
