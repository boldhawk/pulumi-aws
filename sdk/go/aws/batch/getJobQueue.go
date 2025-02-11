// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The Batch Job Queue data source allows access to details of a specific
// job queue within AWS Batch.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/batch_job_queue.html.markdown.
func LookupJobQueue(ctx *pulumi.Context, args *GetJobQueueArgs) (*GetJobQueueResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("aws:batch/getJobQueue:getJobQueue", inputs)
	if err != nil {
		return nil, err
	}
	return &GetJobQueueResult{
		Arn: outputs["arn"],
		ComputeEnvironmentOrders: outputs["computeEnvironmentOrders"],
		Name: outputs["name"],
		Priority: outputs["priority"],
		State: outputs["state"],
		Status: outputs["status"],
		StatusReason: outputs["statusReason"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getJobQueue.
type GetJobQueueArgs struct {
	// The name of the job queue.
	Name interface{}
}

// A collection of values returned by getJobQueue.
type GetJobQueueResult struct {
	// The ARN of the job queue.
	Arn interface{}
	// The compute environments that are attached to the job queue and the order in
	// which job placement is preferred. Compute environments are selected for job placement in ascending order.
	// * `compute_environment_order.#.order` - The order of the compute environment.
	// * `compute_environment_order.#.compute_environment` - The ARN of the compute environment.
	ComputeEnvironmentOrders interface{}
	Name interface{}
	// The priority of the job queue. Job queues with a higher priority are evaluated first when
	// associated with the same compute environment.
	Priority interface{}
	// Describes the ability of the queue to accept new jobs (for example, `ENABLED` or `DISABLED`).
	State interface{}
	// The current status of the job queue (for example, `CREATING` or `VALID`).
	Status interface{}
	// A short, human-readable string to provide additional details about the current status
	// of the job queue.
	StatusReason interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
