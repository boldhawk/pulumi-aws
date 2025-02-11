// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cfg
{
    /// <summary>
    /// Provides an AWS Config Delivery Channel.
    /// 
    /// &gt; **Note:** Delivery Channel requires a [Configuration Recorder](https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html) to be present. Use of `depends_on` (as shown below) is recommended to avoid race conditions.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/config_delivery_channel.html.markdown.
    /// </summary>
    public partial class DeliveryChannel : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the S3 bucket used to store the configuration history.
        /// </summary>
        [Output("s3BucketName")]
        public Output<string> S3BucketName { get; private set; } = null!;

        /// <summary>
        /// The prefix for the specified S3 bucket.
        /// </summary>
        [Output("s3KeyPrefix")]
        public Output<string?> S3KeyPrefix { get; private set; } = null!;

        /// <summary>
        /// Options for how AWS Config delivers configuration snapshots. See below
        /// </summary>
        [Output("snapshotDeliveryProperties")]
        public Output<Outputs.DeliveryChannelSnapshotDeliveryProperties?> SnapshotDeliveryProperties { get; private set; } = null!;

        /// <summary>
        /// The ARN of the SNS topic that AWS Config delivers notifications to.
        /// </summary>
        [Output("snsTopicArn")]
        public Output<string?> SnsTopicArn { get; private set; } = null!;


        /// <summary>
        /// Create a DeliveryChannel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeliveryChannel(string name, DeliveryChannelArgs args, CustomResourceOptions? options = null)
            : base("aws:cfg/deliveryChannel:DeliveryChannel", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private DeliveryChannel(string name, Input<string> id, DeliveryChannelState? state = null, CustomResourceOptions? options = null)
            : base("aws:cfg/deliveryChannel:DeliveryChannel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DeliveryChannel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeliveryChannel Get(string name, Input<string> id, DeliveryChannelState? state = null, CustomResourceOptions? options = null)
        {
            return new DeliveryChannel(name, id, state, options);
        }
    }

    public sealed class DeliveryChannelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the S3 bucket used to store the configuration history.
        /// </summary>
        [Input("s3BucketName", required: true)]
        public Input<string> S3BucketName { get; set; } = null!;

        /// <summary>
        /// The prefix for the specified S3 bucket.
        /// </summary>
        [Input("s3KeyPrefix")]
        public Input<string>? S3KeyPrefix { get; set; }

        /// <summary>
        /// Options for how AWS Config delivers configuration snapshots. See below
        /// </summary>
        [Input("snapshotDeliveryProperties")]
        public Input<Inputs.DeliveryChannelSnapshotDeliveryPropertiesArgs>? SnapshotDeliveryProperties { get; set; }

        /// <summary>
        /// The ARN of the SNS topic that AWS Config delivers notifications to.
        /// </summary>
        [Input("snsTopicArn")]
        public Input<string>? SnsTopicArn { get; set; }

        public DeliveryChannelArgs()
        {
        }
    }

    public sealed class DeliveryChannelState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the S3 bucket used to store the configuration history.
        /// </summary>
        [Input("s3BucketName")]
        public Input<string>? S3BucketName { get; set; }

        /// <summary>
        /// The prefix for the specified S3 bucket.
        /// </summary>
        [Input("s3KeyPrefix")]
        public Input<string>? S3KeyPrefix { get; set; }

        /// <summary>
        /// Options for how AWS Config delivers configuration snapshots. See below
        /// </summary>
        [Input("snapshotDeliveryProperties")]
        public Input<Inputs.DeliveryChannelSnapshotDeliveryPropertiesGetArgs>? SnapshotDeliveryProperties { get; set; }

        /// <summary>
        /// The ARN of the SNS topic that AWS Config delivers notifications to.
        /// </summary>
        [Input("snsTopicArn")]
        public Input<string>? SnsTopicArn { get; set; }

        public DeliveryChannelState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class DeliveryChannelSnapshotDeliveryPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// - The frequency with which AWS Config recurringly delivers configuration snapshots.
        /// e.g. `One_Hour` or `Three_Hours`.
        /// Valid values are listed [here](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigSnapshotDeliveryProperties.html#API_ConfigSnapshotDeliveryProperties_Contents).
        /// </summary>
        [Input("deliveryFrequency")]
        public Input<string>? DeliveryFrequency { get; set; }

        public DeliveryChannelSnapshotDeliveryPropertiesArgs()
        {
        }
    }

    public sealed class DeliveryChannelSnapshotDeliveryPropertiesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// - The frequency with which AWS Config recurringly delivers configuration snapshots.
        /// e.g. `One_Hour` or `Three_Hours`.
        /// Valid values are listed [here](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigSnapshotDeliveryProperties.html#API_ConfigSnapshotDeliveryProperties_Contents).
        /// </summary>
        [Input("deliveryFrequency")]
        public Input<string>? DeliveryFrequency { get; set; }

        public DeliveryChannelSnapshotDeliveryPropertiesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class DeliveryChannelSnapshotDeliveryProperties
    {
        /// <summary>
        /// - The frequency with which AWS Config recurringly delivers configuration snapshots.
        /// e.g. `One_Hour` or `Three_Hours`.
        /// Valid values are listed [here](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigSnapshotDeliveryProperties.html#API_ConfigSnapshotDeliveryProperties_Contents).
        /// </summary>
        public readonly string? DeliveryFrequency;

        [OutputConstructor]
        private DeliveryChannelSnapshotDeliveryProperties(string? deliveryFrequency)
        {
            DeliveryFrequency = deliveryFrequency;
        }
    }
    }
}
