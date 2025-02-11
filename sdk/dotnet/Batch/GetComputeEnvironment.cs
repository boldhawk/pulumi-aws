// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Batch
{
    public static partial class Invokes
    {
        /// <summary>
        /// The Batch Compute Environment data source allows access to details of a specific
        /// compute environment within AWS Batch.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/batch_compute_environment.html.markdown.
        /// </summary>
        public static Task<GetComputeEnvironmentResult> GetComputeEnvironment(GetComputeEnvironmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetComputeEnvironmentResult>("aws:batch/getComputeEnvironment:getComputeEnvironment", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetComputeEnvironmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Batch Compute Environment
        /// </summary>
        [Input("computeEnvironmentName", required: true)]
        public Input<string> ComputeEnvironmentName { get; set; } = null!;

        public GetComputeEnvironmentArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetComputeEnvironmentResult
    {
        /// <summary>
        /// The ARN of the compute environment.
        /// </summary>
        public readonly string Arn;
        public readonly string ComputeEnvironmentName;
        /// <summary>
        /// The ARN of the underlying Amazon ECS cluster used by the compute environment.
        /// </summary>
        public readonly string EcsClusterArn;
        /// <summary>
        /// The ARN of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
        /// </summary>
        public readonly string ServiceRole;
        /// <summary>
        /// The state of the compute environment (for example, `ENABLED` or `DISABLED`). If the state is `ENABLED`, then the compute environment accepts jobs from a queue and can scale out automatically based on queues.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The current status of the compute environment (for example, `CREATING` or `VALID`).
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A short, human-readable string to provide additional details about the current status of the compute environment.
        /// </summary>
        public readonly string StatusReason;
        /// <summary>
        /// The type of the compute environment (for example, `MANAGED` or `UNMANAGED`).
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetComputeEnvironmentResult(
            string arn,
            string computeEnvironmentName,
            string ecsClusterArn,
            string serviceRole,
            string state,
            string status,
            string statusReason,
            string type,
            string id)
        {
            Arn = arn;
            ComputeEnvironmentName = computeEnvironmentName;
            EcsClusterArn = ecsClusterArn;
            ServiceRole = serviceRole;
            State = state;
            Status = status;
            StatusReason = statusReason;
            Type = type;
            Id = id;
        }
    }
}
