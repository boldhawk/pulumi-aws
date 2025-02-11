// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFormation
{
    public static partial class Invokes
    {
        /// <summary>
        /// The CloudFormation Stack data source allows access to stack
        /// outputs and other useful data including the template body.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/cloudformation_stack.html.markdown.
        /// </summary>
        public static Task<GetStackResult> GetStack(GetStackArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStackResult>("aws:cloudformation/getStack:getStack", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetStackArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the stack
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetStackArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetStackResult
    {
        /// <summary>
        /// A list of capabilities
        /// </summary>
        public readonly ImmutableArray<string> Capabilities;
        /// <summary>
        /// Description of the stack
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Whether the rollback of the stack is disabled when stack creation fails
        /// </summary>
        public readonly bool DisableRollback;
        /// <summary>
        /// The ARN of the IAM role used to create the stack.
        /// </summary>
        public readonly string IamRoleArn;
        public readonly string Name;
        /// <summary>
        /// A list of SNS topic ARNs to publish stack related events
        /// </summary>
        public readonly ImmutableArray<string> NotificationArns;
        /// <summary>
        /// A map of outputs from the stack.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Outputs;
        /// <summary>
        /// A map of parameters that specify input parameters for the stack.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Parameters;
        /// <summary>
        /// A map of tags associated with this stack.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// Structure containing the template body.
        /// </summary>
        public readonly string TemplateBody;
        /// <summary>
        /// The amount of time that can pass before the stack status becomes `CREATE_FAILED`
        /// </summary>
        public readonly int TimeoutInMinutes;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetStackResult(
            ImmutableArray<string> capabilities,
            string description,
            bool disableRollback,
            string iamRoleArn,
            string name,
            ImmutableArray<string> notificationArns,
            ImmutableDictionary<string, object> outputs,
            ImmutableDictionary<string, object> parameters,
            ImmutableDictionary<string, object> tags,
            string templateBody,
            int timeoutInMinutes,
            string id)
        {
            Capabilities = capabilities;
            Description = description;
            DisableRollback = disableRollback;
            IamRoleArn = iamRoleArn;
            Name = name;
            NotificationArns = notificationArns;
            Outputs = outputs;
            Parameters = parameters;
            Tags = tags;
            TemplateBody = templateBody;
            TimeoutInMinutes = timeoutInMinutes;
            Id = id;
        }
    }
}
