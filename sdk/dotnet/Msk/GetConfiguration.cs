// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Msk
{
    public static partial class Invokes
    {
        /// <summary>
        /// Get information on an Amazon MSK Configuration.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/msk_configuration.html.markdown.
        /// </summary>
        public static Task<GetConfigurationResult> GetConfiguration(GetConfigurationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConfigurationResult>("aws:msk/getConfiguration:getConfiguration", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the configuration.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetConfigurationArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetConfigurationResult
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the configuration.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Description of the configuration.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// List of Apache Kafka versions which can use this configuration.
        /// </summary>
        public readonly ImmutableArray<string> KafkaVersions;
        /// <summary>
        /// Latest revision of the configuration.
        /// </summary>
        public readonly int LatestRevision;
        public readonly string Name;
        /// <summary>
        /// Contents of the server.properties file.
        /// </summary>
        public readonly string ServerProperties;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetConfigurationResult(
            string arn,
            string description,
            ImmutableArray<string> kafkaVersions,
            int latestRevision,
            string name,
            string serverProperties,
            string id)
        {
            Arn = arn;
            Description = description;
            KafkaVersions = kafkaVersions;
            LatestRevision = latestRevision;
            Name = name;
            ServerProperties = serverProperties;
            Id = id;
        }
    }
}
