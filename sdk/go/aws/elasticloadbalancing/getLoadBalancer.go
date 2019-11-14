// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancing

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides information about a "classic" Elastic Load Balancer (ELB).
// See [LB Data Source](https://www.terraform.io/docs/providers/aws/d/lb.html) if you are looking for "v2"
// Application Load Balancer (ALB) or Network Load Balancer (NLB).
// 
// This data source can prove useful when a module accepts an LB as an input
// variable and needs to, for example, determine the security groups associated
// with it, etc.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elb_legacy.html.markdown.
func LookupLoadBalancer(ctx *pulumi.Context, args *GetLoadBalancerArgs) (*GetLoadBalancerResult, error) {
	var rv GetLoadBalancerResult
	err := ctx.Invoke("aws:elasticloadbalancing/getLoadBalancer:getLoadBalancer", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLoadBalancer.
type GetLoadBalancerArgs struct {
	// The unique name of the load balancer.
	Name string `pulumi:"name"`
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getLoadBalancer.
type GetLoadBalancerResult struct {
	AccessLogs interface{} `pulumi:"accessLogs"`
	AvailabilityZones []string `pulumi:"availabilityZones"`
	ConnectionDraining bool `pulumi:"connectionDraining"`
	ConnectionDrainingTimeout int `pulumi:"connectionDrainingTimeout"`
	CrossZoneLoadBalancing bool `pulumi:"crossZoneLoadBalancing"`
	DnsName string `pulumi:"dnsName"`
	HealthCheck interface{} `pulumi:"healthCheck"`
	IdleTimeout int `pulumi:"idleTimeout"`
	Instances []string `pulumi:"instances"`
	Internal bool `pulumi:"internal"`
	Listeners []interface{} `pulumi:"listeners"`
	Name string `pulumi:"name"`
	SecurityGroups []string `pulumi:"securityGroups"`
	SourceSecurityGroup string `pulumi:"sourceSecurityGroup"`
	SourceSecurityGroupId string `pulumi:"sourceSecurityGroupId"`
	Subnets []string `pulumi:"subnets"`
	Tags map[string]string `pulumi:"tags"`
	ZoneId string `pulumi:"zoneId"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
