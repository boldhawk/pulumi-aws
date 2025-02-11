// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DynamoDB
{
    public static partial class Invokes
    {
        /// <summary>
        /// Provides information about a DynamoDB table.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/dynamodb_table.html.markdown.
        /// </summary>
        public static Task<GetTableResult> GetTable(GetTableArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTableResult>("aws:dynamodb/getTable:getTable", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetTableArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the DynamoDB table.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("serverSideEncryption")]
        public Input<Inputs.GetTableServerSideEncryptionArgs>? ServerSideEncryption { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetTableArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetTableResult
    {
        public readonly string Arn;
        public readonly ImmutableArray<Outputs.GetTableAttributesResult> Attributes;
        public readonly string BillingMode;
        public readonly ImmutableArray<Outputs.GetTableGlobalSecondaryIndexesResult> GlobalSecondaryIndexes;
        public readonly string HashKey;
        public readonly ImmutableArray<Outputs.GetTableLocalSecondaryIndexesResult> LocalSecondaryIndexes;
        public readonly string Name;
        public readonly Outputs.GetTablePointInTimeRecoveryResult PointInTimeRecovery;
        public readonly string RangeKey;
        public readonly int ReadCapacity;
        public readonly Outputs.GetTableServerSideEncryptionResult ServerSideEncryption;
        public readonly string StreamArn;
        public readonly bool StreamEnabled;
        public readonly string StreamLabel;
        public readonly string StreamViewType;
        public readonly ImmutableDictionary<string, object> Tags;
        public readonly Outputs.GetTableTtlResult Ttl;
        public readonly int WriteCapacity;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetTableResult(
            string arn,
            ImmutableArray<Outputs.GetTableAttributesResult> attributes,
            string billingMode,
            ImmutableArray<Outputs.GetTableGlobalSecondaryIndexesResult> globalSecondaryIndexes,
            string hashKey,
            ImmutableArray<Outputs.GetTableLocalSecondaryIndexesResult> localSecondaryIndexes,
            string name,
            Outputs.GetTablePointInTimeRecoveryResult pointInTimeRecovery,
            string rangeKey,
            int readCapacity,
            Outputs.GetTableServerSideEncryptionResult serverSideEncryption,
            string streamArn,
            bool streamEnabled,
            string streamLabel,
            string streamViewType,
            ImmutableDictionary<string, object> tags,
            Outputs.GetTableTtlResult ttl,
            int writeCapacity,
            string id)
        {
            Arn = arn;
            Attributes = attributes;
            BillingMode = billingMode;
            GlobalSecondaryIndexes = globalSecondaryIndexes;
            HashKey = hashKey;
            LocalSecondaryIndexes = localSecondaryIndexes;
            Name = name;
            PointInTimeRecovery = pointInTimeRecovery;
            RangeKey = rangeKey;
            ReadCapacity = readCapacity;
            ServerSideEncryption = serverSideEncryption;
            StreamArn = streamArn;
            StreamEnabled = streamEnabled;
            StreamLabel = streamLabel;
            StreamViewType = streamViewType;
            Tags = tags;
            Ttl = ttl;
            WriteCapacity = writeCapacity;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetTableServerSideEncryptionArgs : Pulumi.ResourceArgs
    {
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public GetTableServerSideEncryptionArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetTableAttributesResult
    {
        /// <summary>
        /// The name of the DynamoDB table.
        /// </summary>
        public readonly string Name;
        public readonly string Type;

        [OutputConstructor]
        private GetTableAttributesResult(
            string name,
            string type)
        {
            Name = name;
            Type = type;
        }
    }

    [OutputType]
    public sealed class GetTableGlobalSecondaryIndexesResult
    {
        public readonly string HashKey;
        /// <summary>
        /// The name of the DynamoDB table.
        /// </summary>
        public readonly string Name;
        public readonly ImmutableArray<string> NonKeyAttributes;
        public readonly string ProjectionType;
        public readonly string RangeKey;
        public readonly int ReadCapacity;
        public readonly int WriteCapacity;

        [OutputConstructor]
        private GetTableGlobalSecondaryIndexesResult(
            string hashKey,
            string name,
            ImmutableArray<string> nonKeyAttributes,
            string projectionType,
            string rangeKey,
            int readCapacity,
            int writeCapacity)
        {
            HashKey = hashKey;
            Name = name;
            NonKeyAttributes = nonKeyAttributes;
            ProjectionType = projectionType;
            RangeKey = rangeKey;
            ReadCapacity = readCapacity;
            WriteCapacity = writeCapacity;
        }
    }

    [OutputType]
    public sealed class GetTableLocalSecondaryIndexesResult
    {
        /// <summary>
        /// The name of the DynamoDB table.
        /// </summary>
        public readonly string Name;
        public readonly ImmutableArray<string> NonKeyAttributes;
        public readonly string ProjectionType;
        public readonly string RangeKey;

        [OutputConstructor]
        private GetTableLocalSecondaryIndexesResult(
            string name,
            ImmutableArray<string> nonKeyAttributes,
            string projectionType,
            string rangeKey)
        {
            Name = name;
            NonKeyAttributes = nonKeyAttributes;
            ProjectionType = projectionType;
            RangeKey = rangeKey;
        }
    }

    [OutputType]
    public sealed class GetTablePointInTimeRecoveryResult
    {
        public readonly bool Enabled;

        [OutputConstructor]
        private GetTablePointInTimeRecoveryResult(bool enabled)
        {
            Enabled = enabled;
        }
    }

    [OutputType]
    public sealed class GetTableServerSideEncryptionResult
    {
        public readonly bool Enabled;

        [OutputConstructor]
        private GetTableServerSideEncryptionResult(bool enabled)
        {
            Enabled = enabled;
        }
    }

    [OutputType]
    public sealed class GetTableTtlResult
    {
        public readonly string AttributeName;
        public readonly bool Enabled;

        [OutputConstructor]
        private GetTableTtlResult(
            string attributeName,
            bool enabled)
        {
            AttributeName = attributeName;
            Enabled = enabled;
        }
    }
    }
}
