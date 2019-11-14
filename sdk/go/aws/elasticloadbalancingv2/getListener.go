// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > **Note:** `alb.Listener` is known as `lb.Listener`. The functionality is identical.
// 
// Provides information about a Load Balancer Listener.
// 
// This data source can prove useful when a module accepts an LB Listener as an
// input variable and needs to know the LB it is attached to, or other
// information specific to the listener in question.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/lb_listener_legacy.html.markdown.
func LookupListener(ctx *pulumi.Context, args *GetListenerArgs) (*GetListenerResult, error) {
	var rv GetListenerResult
	err := ctx.Invoke("aws:elasticloadbalancingv2/getListener:getListener", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getListener.
type GetListenerArgs struct {
	// The arn of the listener. Required if `loadBalancerArn` and `port` is not set.
	Arn string `pulumi:"arn"`
	// The arn of the load balancer. Required if `arn` is not set.
	LoadBalancerArn string `pulumi:"loadBalancerArn"`
	// The port of the listener. Required if `arn` is not set.
	Port int `pulumi:"port"`
}

// A collection of values returned by getListener.
type GetListenerResult struct {
	Arn string `pulumi:"arn"`
	CertificateArn string `pulumi:"certificateArn"`
	DefaultActions []interface{} `pulumi:"defaultActions"`
	LoadBalancerArn string `pulumi:"loadBalancerArn"`
	Port int `pulumi:"port"`
	Protocol string `pulumi:"protocol"`
	SslPolicy string `pulumi:"sslPolicy"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
