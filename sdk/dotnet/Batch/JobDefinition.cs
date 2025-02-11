// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Batch
{
    /// <summary>
    /// Provides a Batch Job Definition resource.
    /// 
    /// ## retry_strategy
    /// 
    /// `retry_strategy` supports the following:
    /// 
    /// * `attempts` - (Optional) The number of times to move a job to the `RUNNABLE` status. You may specify between `1` and `10` attempts.
    /// 
    /// ## timeout
    /// 
    /// `timeout` supports the following:
    /// 
    /// * `attempt_duration_seconds` - (Optional) The time duration in seconds after which AWS Batch terminates your jobs if they have not finished. The minimum value for the timeout is `60` seconds.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/batch_job_definition.html.markdown.
    /// </summary>
    public partial class JobDefinition : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name of the job definition.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
        /// provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
        /// </summary>
        [Output("containerProperties")]
        public Output<string?> ContainerProperties { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the job definition.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the parameter substitution placeholders to set in the job definition.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, string>?> Parameters { get; private set; } = null!;

        /// <summary>
        /// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
        /// Maximum number of `retry_strategy` is `1`.  Defined below.
        /// </summary>
        [Output("retryStrategy")]
        public Output<Outputs.JobDefinitionRetryStrategy?> RetryStrategy { get; private set; } = null!;

        /// <summary>
        /// The revision of the job definition.
        /// </summary>
        [Output("revision")]
        public Output<int> Revision { get; private set; } = null!;

        /// <summary>
        /// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
        /// </summary>
        [Output("timeout")]
        public Output<Outputs.JobDefinitionTimeout?> Timeout { get; private set; } = null!;

        /// <summary>
        /// The type of job definition.  Must be `container`
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a JobDefinition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public JobDefinition(string name, JobDefinitionArgs args, CustomResourceOptions? options = null)
            : base("aws:batch/jobDefinition:JobDefinition", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private JobDefinition(string name, Input<string> id, JobDefinitionState? state = null, CustomResourceOptions? options = null)
            : base("aws:batch/jobDefinition:JobDefinition", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing JobDefinition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static JobDefinition Get(string name, Input<string> id, JobDefinitionState? state = null, CustomResourceOptions? options = null)
        {
            return new JobDefinition(name, id, state, options);
        }
    }

    public sealed class JobDefinitionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
        /// provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
        /// </summary>
        [Input("containerProperties")]
        public Input<string>? ContainerProperties { get; set; }

        /// <summary>
        /// Specifies the name of the job definition.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// Specifies the parameter substitution placeholders to set in the job definition.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
        /// Maximum number of `retry_strategy` is `1`.  Defined below.
        /// </summary>
        [Input("retryStrategy")]
        public Input<Inputs.JobDefinitionRetryStrategyArgs>? RetryStrategy { get; set; }

        /// <summary>
        /// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
        /// </summary>
        [Input("timeout")]
        public Input<Inputs.JobDefinitionTimeoutArgs>? Timeout { get; set; }

        /// <summary>
        /// The type of job definition.  Must be `container`
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public JobDefinitionArgs()
        {
        }
    }

    public sealed class JobDefinitionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name of the job definition.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
        /// provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
        /// </summary>
        [Input("containerProperties")]
        public Input<string>? ContainerProperties { get; set; }

        /// <summary>
        /// Specifies the name of the job definition.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// Specifies the parameter substitution placeholders to set in the job definition.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
        /// Maximum number of `retry_strategy` is `1`.  Defined below.
        /// </summary>
        [Input("retryStrategy")]
        public Input<Inputs.JobDefinitionRetryStrategyGetArgs>? RetryStrategy { get; set; }

        /// <summary>
        /// The revision of the job definition.
        /// </summary>
        [Input("revision")]
        public Input<int>? Revision { get; set; }

        /// <summary>
        /// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
        /// </summary>
        [Input("timeout")]
        public Input<Inputs.JobDefinitionTimeoutGetArgs>? Timeout { get; set; }

        /// <summary>
        /// The type of job definition.  Must be `container`
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public JobDefinitionState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class JobDefinitionRetryStrategyArgs : Pulumi.ResourceArgs
    {
        [Input("attempts")]
        public Input<int>? Attempts { get; set; }

        public JobDefinitionRetryStrategyArgs()
        {
        }
    }

    public sealed class JobDefinitionRetryStrategyGetArgs : Pulumi.ResourceArgs
    {
        [Input("attempts")]
        public Input<int>? Attempts { get; set; }

        public JobDefinitionRetryStrategyGetArgs()
        {
        }
    }

    public sealed class JobDefinitionTimeoutArgs : Pulumi.ResourceArgs
    {
        [Input("attemptDurationSeconds")]
        public Input<int>? AttemptDurationSeconds { get; set; }

        public JobDefinitionTimeoutArgs()
        {
        }
    }

    public sealed class JobDefinitionTimeoutGetArgs : Pulumi.ResourceArgs
    {
        [Input("attemptDurationSeconds")]
        public Input<int>? AttemptDurationSeconds { get; set; }

        public JobDefinitionTimeoutGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class JobDefinitionRetryStrategy
    {
        public readonly int? Attempts;

        [OutputConstructor]
        private JobDefinitionRetryStrategy(int? attempts)
        {
            Attempts = attempts;
        }
    }

    [OutputType]
    public sealed class JobDefinitionTimeout
    {
        public readonly int? AttemptDurationSeconds;

        [OutputConstructor]
        private JobDefinitionTimeout(int? attemptDurationSeconds)
        {
            AttemptDurationSeconds = attemptDurationSeconds;
        }
    }
    }
}
