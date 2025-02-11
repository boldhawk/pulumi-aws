// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package applicationloadbalancing

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
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/alb_target_group_legacy.html.markdown.
func LookupTargetGroup(ctx *pulumi.Context, args *GetTargetGroupArgs) (*GetTargetGroupResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["arn"] = args.Arn
		inputs["name"] = args.Name
		inputs["tags"] = args.Tags
	}
	outputs, err := ctx.Invoke("aws:applicationloadbalancing/getTargetGroup:getTargetGroup", inputs)
	if err != nil {
		return nil, err
	}
	return &GetTargetGroupResult{
		Arn: outputs["arn"],
		ArnSuffix: outputs["arnSuffix"],
		DeregistrationDelay: outputs["deregistrationDelay"],
		HealthCheck: outputs["healthCheck"],
		LambdaMultiValueHeadersEnabled: outputs["lambdaMultiValueHeadersEnabled"],
		Name: outputs["name"],
		Port: outputs["port"],
		Protocol: outputs["protocol"],
		ProxyProtocolV2: outputs["proxyProtocolV2"],
		SlowStart: outputs["slowStart"],
		Stickiness: outputs["stickiness"],
		Tags: outputs["tags"],
		TargetType: outputs["targetType"],
		VpcId: outputs["vpcId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getTargetGroup.
type GetTargetGroupArgs struct {
	// The full ARN of the target group.
	Arn interface{}
	// The unique name of the target group.
	Name interface{}
	Tags interface{}
}

// A collection of values returned by getTargetGroup.
type GetTargetGroupResult struct {
	Arn interface{}
	ArnSuffix interface{}
	DeregistrationDelay interface{}
	HealthCheck interface{}
	LambdaMultiValueHeadersEnabled interface{}
	Name interface{}
	Port interface{}
	Protocol interface{}
	ProxyProtocolV2 interface{}
	SlowStart interface{}
	Stickiness interface{}
	Tags interface{}
	TargetType interface{}
	VpcId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
