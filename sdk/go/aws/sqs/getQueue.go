// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sqs

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the ARN and URL of queue in AWS Simple Queue Service (SQS).
// By using this data source, you can reference SQS queues without having to hardcode
// the ARNs as input.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/sqs_queue.html.markdown.
func LookupQueue(ctx *pulumi.Context, args *GetQueueArgs) (*GetQueueResult, error) {
	var rv GetQueueResult
	err := ctx.Invoke("aws:sqs/getQueue:getQueue", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQueue.
type GetQueueArgs struct {
	// The name of the queue to match.
	Name string `pulumi:"name"`
}

// A collection of values returned by getQueue.
type GetQueueResult struct {
	// The Amazon Resource Name (ARN) of the queue.
	Arn string `pulumi:"arn"`
	Name string `pulumi:"name"`
	// The URL of the queue.
	Url string `pulumi:"url"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
