// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway
{
    /// <summary>
    /// Provides an API Gateway Stage.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_stage.html.markdown.
    /// </summary>
    public partial class Stage : Pulumi.CustomResource
    {
        /// <summary>
        /// Enables access logs for the API stage. Detailed below.
        /// </summary>
        [Output("accessLogSettings")]
        public Output<Outputs.StageAccessLogSettings?> AccessLogSettings { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Specifies whether a cache cluster is enabled for the stage
        /// </summary>
        [Output("cacheClusterEnabled")]
        public Output<bool?> CacheClusterEnabled { get; private set; } = null!;

        /// <summary>
        /// The size of the cache cluster for the stage, if enabled.
        /// Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
        /// </summary>
        [Output("cacheClusterSize")]
        public Output<string?> CacheClusterSize { get; private set; } = null!;

        /// <summary>
        /// The identifier of a client certificate for the stage.
        /// </summary>
        [Output("clientCertificateId")]
        public Output<string?> ClientCertificateId { get; private set; } = null!;

        /// <summary>
        /// The ID of the deployment that the stage points to
        /// </summary>
        [Output("deployment")]
        public Output<string> Deployment { get; private set; } = null!;

        /// <summary>
        /// The description of the stage
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The version of the associated API documentation
        /// </summary>
        [Output("documentationVersion")]
        public Output<string?> DocumentationVersion { get; private set; } = null!;

        /// <summary>
        /// The execution ARN to be used in [`lambda_permission`](https://www.terraform.io/docs/providers/aws/r/lambda_permission.html)'s `source_arn`
        /// when allowing API Gateway to invoke a Lambda function,
        /// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
        /// </summary>
        [Output("executionArn")]
        public Output<string> ExecutionArn { get; private set; } = null!;

        /// <summary>
        /// The URL to invoke the API pointing to the stage,
        /// e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
        /// </summary>
        [Output("invokeUrl")]
        public Output<string> InvokeUrl { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated REST API
        /// </summary>
        [Output("restApi")]
        public Output<string> RestApi { get; private set; } = null!;

        /// <summary>
        /// The name of the stage
        /// </summary>
        [Output("stageName")]
        public Output<string> StageName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map that defines the stage variables
        /// </summary>
        [Output("variables")]
        public Output<ImmutableDictionary<string, object>?> Variables { get; private set; } = null!;

        /// <summary>
        /// Whether active tracing with X-ray is enabled. Defaults to `false`.
        /// </summary>
        [Output("xrayTracingEnabled")]
        public Output<bool?> XrayTracingEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a Stage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Stage(string name, StageArgs args, CustomResourceOptions? options = null)
            : base("aws:apigateway/stage:Stage", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Stage(string name, Input<string> id, StageState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigateway/stage:Stage", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Stage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Stage Get(string name, Input<string> id, StageState? state = null, CustomResourceOptions? options = null)
        {
            return new Stage(name, id, state, options);
        }
    }

    public sealed class StageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables access logs for the API stage. Detailed below.
        /// </summary>
        [Input("accessLogSettings")]
        public Input<Inputs.StageAccessLogSettingsArgs>? AccessLogSettings { get; set; }

        /// <summary>
        /// Specifies whether a cache cluster is enabled for the stage
        /// </summary>
        [Input("cacheClusterEnabled")]
        public Input<bool>? CacheClusterEnabled { get; set; }

        /// <summary>
        /// The size of the cache cluster for the stage, if enabled.
        /// Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
        /// </summary>
        [Input("cacheClusterSize")]
        public Input<string>? CacheClusterSize { get; set; }

        /// <summary>
        /// The identifier of a client certificate for the stage.
        /// </summary>
        [Input("clientCertificateId")]
        public Input<string>? ClientCertificateId { get; set; }

        /// <summary>
        /// The ID of the deployment that the stage points to
        /// </summary>
        [Input("deployment", required: true)]
        public Input<string> Deployment { get; set; } = null!;

        /// <summary>
        /// The description of the stage
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The version of the associated API documentation
        /// </summary>
        [Input("documentationVersion")]
        public Input<string>? DocumentationVersion { get; set; }

        /// <summary>
        /// The ID of the associated REST API
        /// </summary>
        [Input("restApi", required: true)]
        public Input<string> RestApi { get; set; } = null!;

        /// <summary>
        /// The name of the stage
        /// </summary>
        [Input("stageName", required: true)]
        public Input<string> StageName { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("variables")]
        private InputMap<object>? _variables;

        /// <summary>
        /// A map that defines the stage variables
        /// </summary>
        public InputMap<object> Variables
        {
            get => _variables ?? (_variables = new InputMap<object>());
            set => _variables = value;
        }

        /// <summary>
        /// Whether active tracing with X-ray is enabled. Defaults to `false`.
        /// </summary>
        [Input("xrayTracingEnabled")]
        public Input<bool>? XrayTracingEnabled { get; set; }

        public StageArgs()
        {
        }
    }

    public sealed class StageState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables access logs for the API stage. Detailed below.
        /// </summary>
        [Input("accessLogSettings")]
        public Input<Inputs.StageAccessLogSettingsGetArgs>? AccessLogSettings { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Specifies whether a cache cluster is enabled for the stage
        /// </summary>
        [Input("cacheClusterEnabled")]
        public Input<bool>? CacheClusterEnabled { get; set; }

        /// <summary>
        /// The size of the cache cluster for the stage, if enabled.
        /// Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
        /// </summary>
        [Input("cacheClusterSize")]
        public Input<string>? CacheClusterSize { get; set; }

        /// <summary>
        /// The identifier of a client certificate for the stage.
        /// </summary>
        [Input("clientCertificateId")]
        public Input<string>? ClientCertificateId { get; set; }

        /// <summary>
        /// The ID of the deployment that the stage points to
        /// </summary>
        [Input("deployment")]
        public Input<string>? Deployment { get; set; }

        /// <summary>
        /// The description of the stage
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The version of the associated API documentation
        /// </summary>
        [Input("documentationVersion")]
        public Input<string>? DocumentationVersion { get; set; }

        /// <summary>
        /// The execution ARN to be used in [`lambda_permission`](https://www.terraform.io/docs/providers/aws/r/lambda_permission.html)'s `source_arn`
        /// when allowing API Gateway to invoke a Lambda function,
        /// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
        /// </summary>
        [Input("executionArn")]
        public Input<string>? ExecutionArn { get; set; }

        /// <summary>
        /// The URL to invoke the API pointing to the stage,
        /// e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
        /// </summary>
        [Input("invokeUrl")]
        public Input<string>? InvokeUrl { get; set; }

        /// <summary>
        /// The ID of the associated REST API
        /// </summary>
        [Input("restApi")]
        public Input<string>? RestApi { get; set; }

        /// <summary>
        /// The name of the stage
        /// </summary>
        [Input("stageName")]
        public Input<string>? StageName { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("variables")]
        private InputMap<object>? _variables;

        /// <summary>
        /// A map that defines the stage variables
        /// </summary>
        public InputMap<object> Variables
        {
            get => _variables ?? (_variables = new InputMap<object>());
            set => _variables = value;
        }

        /// <summary>
        /// Whether active tracing with X-ray is enabled. Defaults to `false`.
        /// </summary>
        [Input("xrayTracingEnabled")]
        public Input<bool>? XrayTracingEnabled { get; set; }

        public StageState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class StageAccessLogSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the log group to send the logs to. Automatically removes trailing `:*` if present.
        /// </summary>
        [Input("destinationArn", required: true)]
        public Input<string> DestinationArn { get; set; } = null!;

        /// <summary>
        /// The formatting and values recorded in the logs. 
        /// For more information on configuring the log format rules visit the AWS [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html)
        /// </summary>
        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        public StageAccessLogSettingsArgs()
        {
        }
    }

    public sealed class StageAccessLogSettingsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the log group to send the logs to. Automatically removes trailing `:*` if present.
        /// </summary>
        [Input("destinationArn", required: true)]
        public Input<string> DestinationArn { get; set; } = null!;

        /// <summary>
        /// The formatting and values recorded in the logs. 
        /// For more information on configuring the log format rules visit the AWS [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html)
        /// </summary>
        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        public StageAccessLogSettingsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class StageAccessLogSettings
    {
        /// <summary>
        /// ARN of the log group to send the logs to. Automatically removes trailing `:*` if present.
        /// </summary>
        public readonly string DestinationArn;
        /// <summary>
        /// The formatting and values recorded in the logs. 
        /// For more information on configuring the log format rules visit the AWS [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html)
        /// </summary>
        public readonly string Format;

        [OutputConstructor]
        private StageAccessLogSettings(
            string destinationArn,
            string format)
        {
            DestinationArn = destinationArn;
            Format = format;
        }
    }
    }
}
