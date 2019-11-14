// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > **Note:** `alb.TargetGroup` is known as `lb.TargetGroup`. The functionality is identical.
// 
// Provides information about a Load Balancer Target Group.
// 
// This data source can prove useful when a module accepts an LB Target Group as an
// input variable and needs to know its attributes. It can also be used to get the ARN of
// an LB Target Group for use in other resources, given LB Target Group name.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/alb_target_group.html.markdown.
func LookupTargetGroup(ctx *pulumi.Context, args *GetTargetGroupArgs) (*GetTargetGroupResult, error) {
var rv GetTargetGroupResult
	err := ctx.Invoke("aws:alb/getTargetGroup:getTargetGroup", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTargetGroup.
type GetTargetGroupArgs struct {
	// The full ARN of the target group.
	Arn string `pulumi:"arn"`
	// The unique name of the target group.
	Name string `pulumi:"name"`
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getTargetGroup.
type GetTargetGroupResult struct {
	Arn string `pulumi:"arn"`
	ArnSuffix string `pulumi:"arnSuffix"`
	DeregistrationDelay int `pulumi:"deregistrationDelay"`
	HealthCheck interface{} `pulumi:"healthCheck"`
	LambdaMultiValueHeadersEnabled bool `pulumi:"lambdaMultiValueHeadersEnabled"`
	Name string `pulumi:"name"`
	Port int `pulumi:"port"`
	Protocol string `pulumi:"protocol"`
	ProxyProtocolV2 bool `pulumi:"proxyProtocolV2"`
	SlowStart int `pulumi:"slowStart"`
	Stickiness interface{} `pulumi:"stickiness"`
	Tags map[string]interface{} `pulumi:"tags"`
	TargetType string `pulumi:"targetType"`
	VpcId string `pulumi:"vpcId"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
