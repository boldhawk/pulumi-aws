// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Lambda
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lambda_function.html.markdown.
    /// </summary>
    public partial class Function : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) identifying your Lambda Function.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Nested block to configure the function's *dead letter queue*. See details below.
        /// </summary>
        [Output("deadLetterConfig")]
        public Output<Outputs.FunctionDeadLetterConfig?> DeadLetterConfig { get; private set; } = null!;

        /// <summary>
        /// Description of what your Lambda Function does.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Lambda environment's configuration settings. Fields documented below.
        /// </summary>
        [Output("environment")]
        public Output<Outputs.FunctionEnvironment?> Environment { get; private set; } = null!;

        /// <summary>
        /// The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
        /// </summary>
        [Output("code")]
        public Output<Archive?> Code { get; private set; } = null!;

        /// <summary>
        /// A unique name for your Lambda Function.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The function [entrypoint][3] in your code.
        /// </summary>
        [Output("handler")]
        public Output<string> Handler { get; private set; } = null!;

        /// <summary>
        /// The ARN to be used for invoking Lambda Function from API Gateway - to be used in [`aws.apigateway.Integration`](https://www.terraform.io/docs/providers/aws/r/api_gateway_integration.html)'s `uri`
        /// </summary>
        [Output("invokeArn")]
        public Output<string> InvokeArn { get; private set; } = null!;

        /// <summary>
        /// The ARN for the KMS encryption key.
        /// </summary>
        [Output("kmsKeyArn")]
        public Output<string?> KmsKeyArn { get; private set; } = null!;

        /// <summary>
        /// The date this resource was last modified.
        /// </summary>
        [Output("lastModified")]
        public Output<string> LastModified { get; private set; } = null!;

        /// <summary>
        /// List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function. See [Lambda Layers][10]
        /// </summary>
        [Output("layers")]
        public Output<ImmutableArray<string>> Layers { get; private set; } = null!;

        /// <summary>
        /// Amount of memory in MB your Lambda Function can use at runtime. Defaults to `128`. See [Limits][5]
        /// </summary>
        [Output("memorySize")]
        public Output<int?> MemorySize { get; private set; } = null!;

        /// <summary>
        /// Whether to publish creation/change as new Lambda Function Version. Defaults to `false`.
        /// </summary>
        [Output("publish")]
        public Output<bool?> Publish { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) identifying your Lambda Function Version
        /// (if versioning is enabled via `publish = true`).
        /// </summary>
        [Output("qualifiedArn")]
        public Output<string> QualifiedArn { get; private set; } = null!;

        /// <summary>
        /// The amount of reserved concurrent executions for this lambda function. A value of `0` disables lambda from being triggered and `-1` removes any concurrency limitations. Defaults to Unreserved Concurrency Limits `-1`. See [Managing Concurrency][9]
        /// </summary>
        [Output("reservedConcurrentExecutions")]
        public Output<int?> ReservedConcurrentExecutions { get; private set; } = null!;

        /// <summary>
        /// IAM role attached to the Lambda Function. This governs both who / what can invoke your Lambda Function, as well as what resources our Lambda Function has access to. See [Lambda Permission Model][4] for more details.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// See [Runtimes][6] for valid values.
        /// </summary>
        [Output("runtime")]
        public Output<string> Runtime { get; private set; } = null!;

        /// <summary>
        /// The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
        /// </summary>
        [Output("s3Bucket")]
        public Output<string?> S3Bucket { get; private set; } = null!;

        /// <summary>
        /// The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
        /// </summary>
        [Output("s3Key")]
        public Output<string?> S3Key { get; private set; } = null!;

        /// <summary>
        /// The object version containing the function's deployment package. Conflicts with `filename`.
        /// </summary>
        [Output("s3ObjectVersion")]
        public Output<string?> S3ObjectVersion { get; private set; } = null!;

        /// <summary>
        /// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3_key`. The usual way to set this is `filebase64sha256("file.zip")` (this provider 0.11.12 and later) or `base64sha256(file("file.zip"))` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda function source archive.
        /// </summary>
        [Output("sourceCodeHash")]
        public Output<string> SourceCodeHash { get; private set; } = null!;

        /// <summary>
        /// The size in bytes of the function .zip file.
        /// </summary>
        [Output("sourceCodeSize")]
        public Output<int> SourceCodeSize { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the object.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The amount of time your Lambda Function has to run in seconds. Defaults to `3`. See [Limits][5]
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;

        [Output("tracingConfig")]
        public Output<Outputs.FunctionTracingConfig> TracingConfig { get; private set; } = null!;

        /// <summary>
        /// Latest published version of your Lambda Function.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// Provide this to allow your function to access your VPC. Fields documented below. See [Lambda in VPC][7]
        /// </summary>
        [Output("vpcConfig")]
        public Output<Outputs.FunctionVpcConfig?> VpcConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Function resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Function(string name, FunctionArgs args, CustomResourceOptions? options = null)
            : base("aws:lambda/function:Function", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Function(string name, Input<string> id, FunctionState? state = null, CustomResourceOptions? options = null)
            : base("aws:lambda/function:Function", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Function resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Function Get(string name, Input<string> id, FunctionState? state = null, CustomResourceOptions? options = null)
        {
            return new Function(name, id, state, options);
        }
    }

    public sealed class FunctionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Nested block to configure the function's *dead letter queue*. See details below.
        /// </summary>
        [Input("deadLetterConfig")]
        public Input<Inputs.FunctionDeadLetterConfigArgs>? DeadLetterConfig { get; set; }

        /// <summary>
        /// Description of what your Lambda Function does.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Lambda environment's configuration settings. Fields documented below.
        /// </summary>
        [Input("environment")]
        public Input<Inputs.FunctionEnvironmentArgs>? Environment { get; set; }

        /// <summary>
        /// The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
        /// </summary>
        [Input("code")]
        public Input<Archive>? Code { get; set; }

        /// <summary>
        /// A unique name for your Lambda Function.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The function [entrypoint][3] in your code.
        /// </summary>
        [Input("handler", required: true)]
        public Input<string> Handler { get; set; } = null!;

        /// <summary>
        /// The ARN for the KMS encryption key.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        [Input("layers")]
        private InputList<string>? _layers;

        /// <summary>
        /// List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function. See [Lambda Layers][10]
        /// </summary>
        public InputList<string> Layers
        {
            get => _layers ?? (_layers = new InputList<string>());
            set => _layers = value;
        }

        /// <summary>
        /// Amount of memory in MB your Lambda Function can use at runtime. Defaults to `128`. See [Limits][5]
        /// </summary>
        [Input("memorySize")]
        public Input<int>? MemorySize { get; set; }

        /// <summary>
        /// Whether to publish creation/change as new Lambda Function Version. Defaults to `false`.
        /// </summary>
        [Input("publish")]
        public Input<bool>? Publish { get; set; }

        /// <summary>
        /// The amount of reserved concurrent executions for this lambda function. A value of `0` disables lambda from being triggered and `-1` removes any concurrency limitations. Defaults to Unreserved Concurrency Limits `-1`. See [Managing Concurrency][9]
        /// </summary>
        [Input("reservedConcurrentExecutions")]
        public Input<int>? ReservedConcurrentExecutions { get; set; }

        /// <summary>
        /// IAM role attached to the Lambda Function. This governs both who / what can invoke your Lambda Function, as well as what resources our Lambda Function has access to. See [Lambda Permission Model][4] for more details.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// See [Runtimes][6] for valid values.
        /// </summary>
        [Input("runtime", required: true)]
        public Input<string> Runtime { get; set; } = null!;

        /// <summary>
        /// The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
        /// </summary>
        [Input("s3Bucket")]
        public Input<string>? S3Bucket { get; set; }

        /// <summary>
        /// The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
        /// </summary>
        [Input("s3Key")]
        public Input<string>? S3Key { get; set; }

        /// <summary>
        /// The object version containing the function's deployment package. Conflicts with `filename`.
        /// </summary>
        [Input("s3ObjectVersion")]
        public Input<string>? S3ObjectVersion { get; set; }

        /// <summary>
        /// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3_key`. The usual way to set this is `filebase64sha256("file.zip")` (this provider 0.11.12 and later) or `base64sha256(file("file.zip"))` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda function source archive.
        /// </summary>
        [Input("sourceCodeHash")]
        public Input<string>? SourceCodeHash { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the object.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The amount of time your Lambda Function has to run in seconds. Defaults to `3`. See [Limits][5]
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        [Input("tracingConfig")]
        public Input<Inputs.FunctionTracingConfigArgs>? TracingConfig { get; set; }

        /// <summary>
        /// Provide this to allow your function to access your VPC. Fields documented below. See [Lambda in VPC][7]
        /// </summary>
        [Input("vpcConfig")]
        public Input<Inputs.FunctionVpcConfigArgs>? VpcConfig { get; set; }

        public FunctionArgs()
        {
        }
    }

    public sealed class FunctionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) identifying your Lambda Function.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Nested block to configure the function's *dead letter queue*. See details below.
        /// </summary>
        [Input("deadLetterConfig")]
        public Input<Inputs.FunctionDeadLetterConfigGetArgs>? DeadLetterConfig { get; set; }

        /// <summary>
        /// Description of what your Lambda Function does.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Lambda environment's configuration settings. Fields documented below.
        /// </summary>
        [Input("environment")]
        public Input<Inputs.FunctionEnvironmentGetArgs>? Environment { get; set; }

        /// <summary>
        /// The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
        /// </summary>
        [Input("code")]
        public Input<Archive>? Code { get; set; }

        /// <summary>
        /// A unique name for your Lambda Function.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The function [entrypoint][3] in your code.
        /// </summary>
        [Input("handler")]
        public Input<string>? Handler { get; set; }

        /// <summary>
        /// The ARN to be used for invoking Lambda Function from API Gateway - to be used in [`aws.apigateway.Integration`](https://www.terraform.io/docs/providers/aws/r/api_gateway_integration.html)'s `uri`
        /// </summary>
        [Input("invokeArn")]
        public Input<string>? InvokeArn { get; set; }

        /// <summary>
        /// The ARN for the KMS encryption key.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// The date this resource was last modified.
        /// </summary>
        [Input("lastModified")]
        public Input<string>? LastModified { get; set; }

        [Input("layers")]
        private InputList<string>? _layers;

        /// <summary>
        /// List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function. See [Lambda Layers][10]
        /// </summary>
        public InputList<string> Layers
        {
            get => _layers ?? (_layers = new InputList<string>());
            set => _layers = value;
        }

        /// <summary>
        /// Amount of memory in MB your Lambda Function can use at runtime. Defaults to `128`. See [Limits][5]
        /// </summary>
        [Input("memorySize")]
        public Input<int>? MemorySize { get; set; }

        /// <summary>
        /// Whether to publish creation/change as new Lambda Function Version. Defaults to `false`.
        /// </summary>
        [Input("publish")]
        public Input<bool>? Publish { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) identifying your Lambda Function Version
        /// (if versioning is enabled via `publish = true`).
        /// </summary>
        [Input("qualifiedArn")]
        public Input<string>? QualifiedArn { get; set; }

        /// <summary>
        /// The amount of reserved concurrent executions for this lambda function. A value of `0` disables lambda from being triggered and `-1` removes any concurrency limitations. Defaults to Unreserved Concurrency Limits `-1`. See [Managing Concurrency][9]
        /// </summary>
        [Input("reservedConcurrentExecutions")]
        public Input<int>? ReservedConcurrentExecutions { get; set; }

        /// <summary>
        /// IAM role attached to the Lambda Function. This governs both who / what can invoke your Lambda Function, as well as what resources our Lambda Function has access to. See [Lambda Permission Model][4] for more details.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// See [Runtimes][6] for valid values.
        /// </summary>
        [Input("runtime")]
        public Input<string>? Runtime { get; set; }

        /// <summary>
        /// The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
        /// </summary>
        [Input("s3Bucket")]
        public Input<string>? S3Bucket { get; set; }

        /// <summary>
        /// The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
        /// </summary>
        [Input("s3Key")]
        public Input<string>? S3Key { get; set; }

        /// <summary>
        /// The object version containing the function's deployment package. Conflicts with `filename`.
        /// </summary>
        [Input("s3ObjectVersion")]
        public Input<string>? S3ObjectVersion { get; set; }

        /// <summary>
        /// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3_key`. The usual way to set this is `filebase64sha256("file.zip")` (this provider 0.11.12 and later) or `base64sha256(file("file.zip"))` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda function source archive.
        /// </summary>
        [Input("sourceCodeHash")]
        public Input<string>? SourceCodeHash { get; set; }

        /// <summary>
        /// The size in bytes of the function .zip file.
        /// </summary>
        [Input("sourceCodeSize")]
        public Input<int>? SourceCodeSize { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the object.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The amount of time your Lambda Function has to run in seconds. Defaults to `3`. See [Limits][5]
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        [Input("tracingConfig")]
        public Input<Inputs.FunctionTracingConfigGetArgs>? TracingConfig { get; set; }

        /// <summary>
        /// Latest published version of your Lambda Function.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// Provide this to allow your function to access your VPC. Fields documented below. See [Lambda in VPC][7]
        /// </summary>
        [Input("vpcConfig")]
        public Input<Inputs.FunctionVpcConfigGetArgs>? VpcConfig { get; set; }

        public FunctionState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class FunctionDeadLetterConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of an SNS topic or SQS queue to notify when an invocation fails. If this
        /// option is used, the function's IAM role must be granted suitable access to write to the target object,
        /// which means allowing either the `sns:Publish` or `sqs:SendMessage` action on this ARN, depending on
        /// which service is targeted.
        /// </summary>
        [Input("targetArn", required: true)]
        public Input<string> TargetArn { get; set; } = null!;

        public FunctionDeadLetterConfigArgs()
        {
        }
    }

    public sealed class FunctionDeadLetterConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of an SNS topic or SQS queue to notify when an invocation fails. If this
        /// option is used, the function's IAM role must be granted suitable access to write to the target object,
        /// which means allowing either the `sns:Publish` or `sqs:SendMessage` action on this ARN, depending on
        /// which service is targeted.
        /// </summary>
        [Input("targetArn", required: true)]
        public Input<string> TargetArn { get; set; } = null!;

        public FunctionDeadLetterConfigGetArgs()
        {
        }
    }

    public sealed class FunctionEnvironmentArgs : Pulumi.ResourceArgs
    {
        [Input("variables")]
        private InputMap<string>? _variables;

        /// <summary>
        /// A map that defines environment variables for the Lambda function.
        /// </summary>
        public InputMap<string> Variables
        {
            get => _variables ?? (_variables = new InputMap<string>());
            set => _variables = value;
        }

        public FunctionEnvironmentArgs()
        {
        }
    }

    public sealed class FunctionEnvironmentGetArgs : Pulumi.ResourceArgs
    {
        [Input("variables")]
        private InputMap<string>? _variables;

        /// <summary>
        /// A map that defines environment variables for the Lambda function.
        /// </summary>
        public InputMap<string> Variables
        {
            get => _variables ?? (_variables = new InputMap<string>());
            set => _variables = value;
        }

        public FunctionEnvironmentGetArgs()
        {
        }
    }

    public sealed class FunctionTracingConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Can be either `PassThrough` or `Active`. If PassThrough, Lambda will only trace
        /// the request from an upstream service if it contains a tracing header with
        /// "sampled=1". If Active, Lambda will respect any tracing header it receives
        /// from an upstream service. If no tracing header is received, Lambda will call
        /// X-Ray for a tracing decision.
        /// </summary>
        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        public FunctionTracingConfigArgs()
        {
        }
    }

    public sealed class FunctionTracingConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Can be either `PassThrough` or `Active`. If PassThrough, Lambda will only trace
        /// the request from an upstream service if it contains a tracing header with
        /// "sampled=1". If Active, Lambda will respect any tracing header it receives
        /// from an upstream service. If no tracing header is received, Lambda will call
        /// X-Ray for a tracing decision.
        /// </summary>
        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        public FunctionTracingConfigGetArgs()
        {
        }
    }

    public sealed class FunctionVpcConfigArgs : Pulumi.ResourceArgs
    {
        [Input("securityGroupIds", required: true)]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of security group IDs associated with the Lambda function.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// A list of subnet IDs associated with the Lambda function.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public FunctionVpcConfigArgs()
        {
        }
    }

    public sealed class FunctionVpcConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("securityGroupIds", required: true)]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of security group IDs associated with the Lambda function.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// A list of subnet IDs associated with the Lambda function.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public FunctionVpcConfigGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class FunctionDeadLetterConfig
    {
        /// <summary>
        /// The ARN of an SNS topic or SQS queue to notify when an invocation fails. If this
        /// option is used, the function's IAM role must be granted suitable access to write to the target object,
        /// which means allowing either the `sns:Publish` or `sqs:SendMessage` action on this ARN, depending on
        /// which service is targeted.
        /// </summary>
        public readonly string TargetArn;

        [OutputConstructor]
        private FunctionDeadLetterConfig(string targetArn)
        {
            TargetArn = targetArn;
        }
    }

    [OutputType]
    public sealed class FunctionEnvironment
    {
        /// <summary>
        /// A map that defines environment variables for the Lambda function.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Variables;

        [OutputConstructor]
        private FunctionEnvironment(ImmutableDictionary<string, string>? variables)
        {
            Variables = variables;
        }
    }

    [OutputType]
    public sealed class FunctionTracingConfig
    {
        /// <summary>
        /// Can be either `PassThrough` or `Active`. If PassThrough, Lambda will only trace
        /// the request from an upstream service if it contains a tracing header with
        /// "sampled=1". If Active, Lambda will respect any tracing header it receives
        /// from an upstream service. If no tracing header is received, Lambda will call
        /// X-Ray for a tracing decision.
        /// </summary>
        public readonly string Mode;

        [OutputConstructor]
        private FunctionTracingConfig(string mode)
        {
            Mode = mode;
        }
    }

    [OutputType]
    public sealed class FunctionVpcConfig
    {
        /// <summary>
        /// A list of security group IDs associated with the Lambda function.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// A list of subnet IDs associated with the Lambda function.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        public readonly string VpcId;

        [OutputConstructor]
        private FunctionVpcConfig(
            ImmutableArray<string> securityGroupIds,
            ImmutableArray<string> subnetIds,
            string vpcId)
        {
            SecurityGroupIds = securityGroupIds;
            SubnetIds = subnetIds;
            VpcId = vpcId;
        }
    }
    }
}
