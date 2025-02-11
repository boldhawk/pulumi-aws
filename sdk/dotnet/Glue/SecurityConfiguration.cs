// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue
{
    /// <summary>
    /// Manages a Glue Security Configuration.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/glue_security_configuration.html.markdown.
    /// </summary>
    public partial class SecurityConfiguration : Pulumi.CustomResource
    {
        /// <summary>
        /// Configuration block containing encryption configuration. Detailed below.
        /// </summary>
        [Output("encryptionConfiguration")]
        public Output<Outputs.SecurityConfigurationEncryptionConfiguration> EncryptionConfiguration { get; private set; } = null!;

        /// <summary>
        /// Name of the security configuration.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityConfiguration(string name, SecurityConfigurationArgs args, CustomResourceOptions? options = null)
            : base("aws:glue/securityConfiguration:SecurityConfiguration", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SecurityConfiguration(string name, Input<string> id, SecurityConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("aws:glue/securityConfiguration:SecurityConfiguration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityConfiguration Get(string name, Input<string> id, SecurityConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityConfiguration(name, id, state, options);
        }
    }

    public sealed class SecurityConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration block containing encryption configuration. Detailed below.
        /// </summary>
        [Input("encryptionConfiguration", required: true)]
        public Input<Inputs.SecurityConfigurationEncryptionConfigurationArgs> EncryptionConfiguration { get; set; } = null!;

        /// <summary>
        /// Name of the security configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SecurityConfigurationArgs()
        {
        }
    }

    public sealed class SecurityConfigurationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration block containing encryption configuration. Detailed below.
        /// </summary>
        [Input("encryptionConfiguration")]
        public Input<Inputs.SecurityConfigurationEncryptionConfigurationGetArgs>? EncryptionConfiguration { get; set; }

        /// <summary>
        /// Name of the security configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SecurityConfigurationState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class SecurityConfigurationEncryptionConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("cloudwatchEncryption", required: true)]
        public Input<SecurityConfigurationEncryptionConfigurationCloudwatchEncryptionArgs> CloudwatchEncryption { get; set; } = null!;

        [Input("jobBookmarksEncryption", required: true)]
        public Input<SecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionArgs> JobBookmarksEncryption { get; set; } = null!;

        /// <summary>
        /// A `s3_encryption ` block as described below, which contains encryption configuration for S3 data.
        /// </summary>
        [Input("s3Encryption", required: true)]
        public Input<SecurityConfigurationEncryptionConfigurationS3EncryptionArgs> S3Encryption { get; set; } = null!;

        public SecurityConfigurationEncryptionConfigurationArgs()
        {
        }
    }

    public sealed class SecurityConfigurationEncryptionConfigurationCloudwatchEncryptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Encryption mode to use for CloudWatch data. Valid values: `DISABLED`, `SSE-KMS`. Default value: `DISABLED`.
        /// </summary>
        [Input("cloudwatchEncryptionMode")]
        public Input<string>? CloudwatchEncryptionMode { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        public SecurityConfigurationEncryptionConfigurationCloudwatchEncryptionArgs()
        {
        }
    }

    public sealed class SecurityConfigurationEncryptionConfigurationCloudwatchEncryptionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Encryption mode to use for CloudWatch data. Valid values: `DISABLED`, `SSE-KMS`. Default value: `DISABLED`.
        /// </summary>
        [Input("cloudwatchEncryptionMode")]
        public Input<string>? CloudwatchEncryptionMode { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        public SecurityConfigurationEncryptionConfigurationCloudwatchEncryptionGetArgs()
        {
        }
    }

    public sealed class SecurityConfigurationEncryptionConfigurationGetArgs : Pulumi.ResourceArgs
    {
        [Input("cloudwatchEncryption", required: true)]
        public Input<SecurityConfigurationEncryptionConfigurationCloudwatchEncryptionGetArgs> CloudwatchEncryption { get; set; } = null!;

        [Input("jobBookmarksEncryption", required: true)]
        public Input<SecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionGetArgs> JobBookmarksEncryption { get; set; } = null!;

        /// <summary>
        /// A `s3_encryption ` block as described below, which contains encryption configuration for S3 data.
        /// </summary>
        [Input("s3Encryption", required: true)]
        public Input<SecurityConfigurationEncryptionConfigurationS3EncryptionGetArgs> S3Encryption { get; set; } = null!;

        public SecurityConfigurationEncryptionConfigurationGetArgs()
        {
        }
    }

    public sealed class SecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Encryption mode to use for job bookmarks data. Valid values: `CSE-KMS`, `DISABLED`. Default value: `DISABLED`.
        /// </summary>
        [Input("jobBookmarksEncryptionMode")]
        public Input<string>? JobBookmarksEncryptionMode { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        public SecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionArgs()
        {
        }
    }

    public sealed class SecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Encryption mode to use for job bookmarks data. Valid values: `CSE-KMS`, `DISABLED`. Default value: `DISABLED`.
        /// </summary>
        [Input("jobBookmarksEncryptionMode")]
        public Input<string>? JobBookmarksEncryptionMode { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        public SecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionGetArgs()
        {
        }
    }

    public sealed class SecurityConfigurationEncryptionConfigurationS3EncryptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// Encryption mode to use for S3 data. Valid values: `DISABLED`, `SSE-KMS`, `SSE-S3`. Default value: `DISABLED`.
        /// </summary>
        [Input("s3EncryptionMode")]
        public Input<string>? S3EncryptionMode { get; set; }

        public SecurityConfigurationEncryptionConfigurationS3EncryptionArgs()
        {
        }
    }

    public sealed class SecurityConfigurationEncryptionConfigurationS3EncryptionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// Encryption mode to use for S3 data. Valid values: `DISABLED`, `SSE-KMS`, `SSE-S3`. Default value: `DISABLED`.
        /// </summary>
        [Input("s3EncryptionMode")]
        public Input<string>? S3EncryptionMode { get; set; }

        public SecurityConfigurationEncryptionConfigurationS3EncryptionGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class SecurityConfigurationEncryptionConfiguration
    {
        public readonly SecurityConfigurationEncryptionConfigurationCloudwatchEncryption CloudwatchEncryption;
        public readonly SecurityConfigurationEncryptionConfigurationJobBookmarksEncryption JobBookmarksEncryption;
        /// <summary>
        /// A `s3_encryption ` block as described below, which contains encryption configuration for S3 data.
        /// </summary>
        public readonly SecurityConfigurationEncryptionConfigurationS3Encryption S3Encryption;

        [OutputConstructor]
        private SecurityConfigurationEncryptionConfiguration(
            SecurityConfigurationEncryptionConfigurationCloudwatchEncryption cloudwatchEncryption,
            SecurityConfigurationEncryptionConfigurationJobBookmarksEncryption jobBookmarksEncryption,
            SecurityConfigurationEncryptionConfigurationS3Encryption s3Encryption)
        {
            CloudwatchEncryption = cloudwatchEncryption;
            JobBookmarksEncryption = jobBookmarksEncryption;
            S3Encryption = s3Encryption;
        }
    }

    [OutputType]
    public sealed class SecurityConfigurationEncryptionConfigurationCloudwatchEncryption
    {
        /// <summary>
        /// Encryption mode to use for CloudWatch data. Valid values: `DISABLED`, `SSE-KMS`. Default value: `DISABLED`.
        /// </summary>
        public readonly string? CloudwatchEncryptionMode;
        /// <summary>
        /// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
        /// </summary>
        public readonly string? KmsKeyArn;

        [OutputConstructor]
        private SecurityConfigurationEncryptionConfigurationCloudwatchEncryption(
            string? cloudwatchEncryptionMode,
            string? kmsKeyArn)
        {
            CloudwatchEncryptionMode = cloudwatchEncryptionMode;
            KmsKeyArn = kmsKeyArn;
        }
    }

    [OutputType]
    public sealed class SecurityConfigurationEncryptionConfigurationJobBookmarksEncryption
    {
        /// <summary>
        /// Encryption mode to use for job bookmarks data. Valid values: `CSE-KMS`, `DISABLED`. Default value: `DISABLED`.
        /// </summary>
        public readonly string? JobBookmarksEncryptionMode;
        /// <summary>
        /// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
        /// </summary>
        public readonly string? KmsKeyArn;

        [OutputConstructor]
        private SecurityConfigurationEncryptionConfigurationJobBookmarksEncryption(
            string? jobBookmarksEncryptionMode,
            string? kmsKeyArn)
        {
            JobBookmarksEncryptionMode = jobBookmarksEncryptionMode;
            KmsKeyArn = kmsKeyArn;
        }
    }

    [OutputType]
    public sealed class SecurityConfigurationEncryptionConfigurationS3Encryption
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
        /// </summary>
        public readonly string? KmsKeyArn;
        /// <summary>
        /// Encryption mode to use for S3 data. Valid values: `DISABLED`, `SSE-KMS`, `SSE-S3`. Default value: `DISABLED`.
        /// </summary>
        public readonly string? S3EncryptionMode;

        [OutputConstructor]
        private SecurityConfigurationEncryptionConfigurationS3Encryption(
            string? kmsKeyArn,
            string? s3EncryptionMode)
        {
            KmsKeyArn = kmsKeyArn;
            S3EncryptionMode = s3EncryptionMode;
        }
    }
    }
}
