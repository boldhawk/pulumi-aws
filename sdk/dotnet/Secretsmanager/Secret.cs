// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecretsManager
{
    /// <summary>
    /// Provides a resource to manage AWS Secrets Manager secret metadata. To manage a secret value, see the [`aws.secretsmanager.SecretVersion` resource](https://www.terraform.io/docs/providers/aws/r/secretsmanager_secret_version.html).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/secretsmanager_secret.html.markdown.
    /// </summary>
    public partial class Secret : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the secret.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A description of the secret.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies the ARN or alias of the AWS KMS customer master key (CMK) to be used to encrypt the secret values in the versions stored in this secret. If you don't specify this value, then Secrets Manager defaults to using the AWS account's default CMK (the one named `aws/secretsmanager`). If the default KMS CMK with that name doesn't yet exist, then AWS Secrets Manager creates it for you automatically the first time.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string?> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// Specifies the friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `name_prefix`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string> NamePrefix { get; private set; } = null!;

        [Output("policy")]
        public Output<string?> Policy { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
        /// </summary>
        [Output("recoveryWindowInDays")]
        public Output<int?> RecoveryWindowInDays { get; private set; } = null!;

        /// <summary>
        /// Specifies whether automatic rotation is enabled for this secret.
        /// </summary>
        [Output("rotationEnabled")]
        public Output<bool> RotationEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the ARN of the Lambda function that can rotate the secret.
        /// </summary>
        [Output("rotationLambdaArn")]
        public Output<string?> RotationLambdaArn { get; private set; } = null!;

        /// <summary>
        /// A structure that defines the rotation configuration for this secret. Defined below.
        /// </summary>
        [Output("rotationRules")]
        public Output<Outputs.SecretRotationRules?> RotationRules { get; private set; } = null!;

        /// <summary>
        /// Specifies a key-value map of user-defined tags that are attached to the secret.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Secret resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Secret(string name, SecretArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:secretsmanager/secret:Secret", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Secret(string name, Input<string> id, SecretState? state = null, CustomResourceOptions? options = null)
            : base("aws:secretsmanager/secret:Secret", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Secret resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Secret Get(string name, Input<string> id, SecretState? state = null, CustomResourceOptions? options = null)
        {
            return new Secret(name, id, state, options);
        }
    }

    public sealed class SecretArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the secret.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the ARN or alias of the AWS KMS customer master key (CMK) to be used to encrypt the secret values in the versions stored in this secret. If you don't specify this value, then Secrets Manager defaults to using the AWS account's default CMK (the one named `aws/secretsmanager`). If the default KMS CMK with that name doesn't yet exist, then AWS Secrets Manager creates it for you automatically the first time.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Specifies the friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `name_prefix`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// Specifies the number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
        /// </summary>
        [Input("recoveryWindowInDays")]
        public Input<int>? RecoveryWindowInDays { get; set; }

        /// <summary>
        /// Specifies the ARN of the Lambda function that can rotate the secret.
        /// </summary>
        [Input("rotationLambdaArn")]
        public Input<string>? RotationLambdaArn { get; set; }

        /// <summary>
        /// A structure that defines the rotation configuration for this secret. Defined below.
        /// </summary>
        [Input("rotationRules")]
        public Input<Inputs.SecretRotationRulesArgs>? RotationRules { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Specifies a key-value map of user-defined tags that are attached to the secret.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public SecretArgs()
        {
        }
    }

    public sealed class SecretState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the secret.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// A description of the secret.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the ARN or alias of the AWS KMS customer master key (CMK) to be used to encrypt the secret values in the versions stored in this secret. If you don't specify this value, then Secrets Manager defaults to using the AWS account's default CMK (the one named `aws/secretsmanager`). If the default KMS CMK with that name doesn't yet exist, then AWS Secrets Manager creates it for you automatically the first time.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Specifies the friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `name_prefix`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// Specifies the number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
        /// </summary>
        [Input("recoveryWindowInDays")]
        public Input<int>? RecoveryWindowInDays { get; set; }

        /// <summary>
        /// Specifies whether automatic rotation is enabled for this secret.
        /// </summary>
        [Input("rotationEnabled")]
        public Input<bool>? RotationEnabled { get; set; }

        /// <summary>
        /// Specifies the ARN of the Lambda function that can rotate the secret.
        /// </summary>
        [Input("rotationLambdaArn")]
        public Input<string>? RotationLambdaArn { get; set; }

        /// <summary>
        /// A structure that defines the rotation configuration for this secret. Defined below.
        /// </summary>
        [Input("rotationRules")]
        public Input<Inputs.SecretRotationRulesGetArgs>? RotationRules { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Specifies a key-value map of user-defined tags that are attached to the secret.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public SecretState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class SecretRotationRulesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the number of days between automatic scheduled rotations of the secret.
        /// </summary>
        [Input("automaticallyAfterDays", required: true)]
        public Input<int> AutomaticallyAfterDays { get; set; } = null!;

        public SecretRotationRulesArgs()
        {
        }
    }

    public sealed class SecretRotationRulesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the number of days between automatic scheduled rotations of the secret.
        /// </summary>
        [Input("automaticallyAfterDays", required: true)]
        public Input<int> AutomaticallyAfterDays { get; set; } = null!;

        public SecretRotationRulesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class SecretRotationRules
    {
        /// <summary>
        /// Specifies the number of days between automatic scheduled rotations of the secret.
        /// </summary>
        public readonly int AutomaticallyAfterDays;

        [OutputConstructor]
        private SecretRotationRules(int automaticallyAfterDays)
        {
            AutomaticallyAfterDays = automaticallyAfterDays;
        }
    }
    }
}
