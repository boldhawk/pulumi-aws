// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch
{
    /// <summary>
    /// Provides a CloudWatch Logs subscription filter resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_subscription_filter.html.markdown.
    /// </summary>
    public partial class LogSubscriptionFilter : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
        /// </summary>
        [Output("destinationArn")]
        public Output<string> DestinationArn { get; private set; } = null!;

        /// <summary>
        /// The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
        /// </summary>
        [Output("distribution")]
        public Output<string?> Distribution { get; private set; } = null!;

        /// <summary>
        /// A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
        /// </summary>
        [Output("filterPattern")]
        public Output<string> FilterPattern { get; private set; } = null!;

        /// <summary>
        /// The name of the log group to associate the subscription filter with
        /// </summary>
        [Output("logGroup")]
        public Output<string> LogGroup { get; private set; } = null!;

        /// <summary>
        /// A name for the subscription filter
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `aws.lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function. 
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;


        /// <summary>
        /// Create a LogSubscriptionFilter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LogSubscriptionFilter(string name, LogSubscriptionFilterArgs args, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/logSubscriptionFilter:LogSubscriptionFilter", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private LogSubscriptionFilter(string name, Input<string> id, LogSubscriptionFilterState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/logSubscriptionFilter:LogSubscriptionFilter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LogSubscriptionFilter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LogSubscriptionFilter Get(string name, Input<string> id, LogSubscriptionFilterState? state = null, CustomResourceOptions? options = null)
        {
            return new LogSubscriptionFilter(name, id, state, options);
        }
    }

    public sealed class LogSubscriptionFilterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
        /// </summary>
        [Input("destinationArn", required: true)]
        public Input<string> DestinationArn { get; set; } = null!;

        /// <summary>
        /// The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
        /// </summary>
        [Input("distribution")]
        public Input<string>? Distribution { get; set; }

        /// <summary>
        /// A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
        /// </summary>
        [Input("filterPattern", required: true)]
        public Input<string> FilterPattern { get; set; } = null!;

        /// <summary>
        /// The name of the log group to associate the subscription filter with
        /// </summary>
        [Input("logGroup", required: true)]
        public Input<string> LogGroup { get; set; } = null!;

        /// <summary>
        /// A name for the subscription filter
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `aws.lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function. 
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        public LogSubscriptionFilterArgs()
        {
        }
    }

    public sealed class LogSubscriptionFilterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
        /// </summary>
        [Input("destinationArn")]
        public Input<string>? DestinationArn { get; set; }

        /// <summary>
        /// The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
        /// </summary>
        [Input("distribution")]
        public Input<string>? Distribution { get; set; }

        /// <summary>
        /// A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
        /// </summary>
        [Input("filterPattern")]
        public Input<string>? FilterPattern { get; set; }

        /// <summary>
        /// The name of the log group to associate the subscription filter with
        /// </summary>
        [Input("logGroup")]
        public Input<string>? LogGroup { get; set; }

        /// <summary>
        /// A name for the subscription filter
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `aws.lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function. 
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        public LogSubscriptionFilterState()
        {
        }
    }
}
