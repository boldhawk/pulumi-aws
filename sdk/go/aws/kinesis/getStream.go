// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kinesis

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get information about a Kinesis Stream for use in other
// resources.
// 
// For more details, see the [Amazon Kinesis Documentation][1].
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/kinesis_stream.html.markdown.
func LookupStream(ctx *pulumi.Context, args *GetStreamArgs) (*GetStreamResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("aws:kinesis/getStream:getStream", inputs)
	if err != nil {
		return nil, err
	}
	return &GetStreamResult{
		Arn: outputs["arn"],
		ClosedShards: outputs["closedShards"],
		CreationTimestamp: outputs["creationTimestamp"],
		Name: outputs["name"],
		OpenShards: outputs["openShards"],
		RetentionPeriod: outputs["retentionPeriod"],
		ShardLevelMetrics: outputs["shardLevelMetrics"],
		Status: outputs["status"],
		Tags: outputs["tags"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getStream.
type GetStreamArgs struct {
	// The name of the Kinesis Stream.
	Name pulumi.StringInput `pulumi:"name"`
}

// A collection of values returned by getStream.
type GetStreamResult struct {
	// The Amazon Resource Name (ARN) of the Kinesis Stream (same as id).
	Arn string `pulumi:"arn"`
	// The list of shard ids in the CLOSED state. See [Shard State][2] for more.
	ClosedShards []interface{} `pulumi:"closedShards"`
	// The approximate UNIX timestamp that the stream was created.
	CreationTimestamp int `pulumi:"creationTimestamp"`
	// The name of the Kinesis Stream.
	Name string `pulumi:"name"`
	// The list of shard ids in the OPEN state. See [Shard State][2] for more.
	OpenShards []interface{} `pulumi:"openShards"`
	// Length of time (in hours) data records are accessible after they are added to the stream.
	RetentionPeriod int `pulumi:"retentionPeriod"`
	// A list of shard-level CloudWatch metrics which are enabled for the stream. See [Monitoring with CloudWatch][3] for more.
	ShardLevelMetrics []interface{} `pulumi:"shardLevelMetrics"`
	// The current status of the stream. The stream status is one of CREATING, DELETING, ACTIVE, or UPDATING.
	Status string `pulumi:"status"`
	// A mapping of tags to assigned to the stream.
	Tags map[string]interface{} `pulumi:"tags"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
