// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecs
{
    public static partial class Invokes
    {
        /// <summary>
        /// The ECS task definition data source allows access to details of
        /// a specific AWS ECS task definition.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ecs_task_definition.html.markdown.
        /// </summary>
        public static Task<GetTaskDefinitionResult> GetTaskDefinition(GetTaskDefinitionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTaskDefinitionResult>("aws:ecs/getTaskDefinition:getTaskDefinition", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetTaskDefinitionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The family for the latest ACTIVE revision, family and revision (family:revision) for a specific revision in the family, the ARN of the task definition to access to.
        /// </summary>
        [Input("taskDefinition", required: true)]
        public Input<string> TaskDefinition { get; set; } = null!;

        public GetTaskDefinitionArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetTaskDefinitionResult
    {
        /// <summary>
        /// The family of this task definition
        /// </summary>
        public readonly string Family;
        /// <summary>
        /// The Docker networking mode to use for the containers in this task.
        /// </summary>
        public readonly string NetworkMode;
        /// <summary>
        /// The revision of this task definition
        /// </summary>
        public readonly int Revision;
        /// <summary>
        /// The status of this task definition
        /// </summary>
        public readonly string Status;
        public readonly string TaskDefinition;
        /// <summary>
        /// The ARN of the IAM role that containers in this task can assume
        /// </summary>
        public readonly string TaskRoleArn;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetTaskDefinitionResult(
            string family,
            string networkMode,
            int revision,
            string status,
            string taskDefinition,
            string taskRoleArn,
            string id)
        {
            Family = family;
            NetworkMode = networkMode;
            Revision = revision;
            Status = status;
            TaskDefinition = taskDefinition;
            TaskRoleArn = taskRoleArn;
            Id = id;
        }
    }
}
