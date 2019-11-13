// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alb

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
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/alb_listener.html.markdown.
func LookupListener(ctx *pulumi.Context, args *GetListenerArgs) (*GetListenerResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["arn"] = args.Arn
		inputs["loadBalancerArn"] = args.LoadBalancerArn
		inputs["port"] = args.Port
	}
	outputs, err := ctx.Invoke("aws:alb/getListener:getListener", inputs)
	if err != nil {
		return nil, err
	}
	return &GetListenerResult{
		Arn: outputs["arn"],
		CertificateArn: outputs["certificateArn"],
		DefaultActions: outputs["defaultActions"],
		LoadBalancerArn: outputs["loadBalancerArn"],
		Port: outputs["port"],
		Protocol: outputs["protocol"],
		SslPolicy: outputs["sslPolicy"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getListener.
type GetListenerArgs struct {
	// The arn of the listener. Required if `loadBalancerArn` and `port` is not set.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The arn of the load balancer. Required if `arn` is not set.
	LoadBalancerArn pulumi.StringInput `pulumi:"loadBalancerArn"`
	// The port of the listener. Required if `arn` is not set.
	Port pulumi.IntInput `pulumi:"port"`
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
