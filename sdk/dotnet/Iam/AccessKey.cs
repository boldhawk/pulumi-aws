// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iam
{
    /// <summary>
    /// Provides an IAM access key. This is a set of credentials that allow API requests to be made as an IAM user.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_access_key.html.markdown.
    /// </summary>
    public partial class AccessKey : Pulumi.CustomResource
    {
        /// <summary>
        /// The encrypted secret, base64 encoded.
        /// &gt; **NOTE:** The encrypted secret may be decrypted using the command line,
        /// for example: `... | base64 --decode | keybase pgp decrypt`.
        /// </summary>
        [Output("encryptedSecret")]
        public Output<string> EncryptedSecret { get; private set; } = null!;

        /// <summary>
        /// The fingerprint of the PGP key used to encrypt
        /// the secret
        /// </summary>
        [Output("keyFingerprint")]
        public Output<string> KeyFingerprint { get; private set; } = null!;

        /// <summary>
        /// Either a base-64 encoded PGP public key, or a
        /// keybase username in the form `keybase:some_person_that_exists`.
        /// </summary>
        [Output("pgpKey")]
        public Output<string?> PgpKey { get; private set; } = null!;

        /// <summary>
        /// The secret access key. Note that this will be written
        /// to the state file. Please supply a `pgp_key` instead, which will prevent the
        /// secret from being stored in plain text
        /// </summary>
        [Output("secret")]
        public Output<string> Secret { get; private set; } = null!;

        /// <summary>
        /// The secret access key converted into an SES SMTP
        /// password by applying [AWS's documented conversion
        /// algorithm](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/smtp-credentials.html#smtp-credentials-convert).
        /// </summary>
        [Output("sesSmtpPassword")]
        public Output<string> SesSmtpPassword { get; private set; } = null!;

        /// <summary>
        /// The access key status to apply. Defaults to `Active`.
        /// Valid values are `Active` and `Inactive`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The IAM user to associate with this access key.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;


        /// <summary>
        /// Create a AccessKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessKey(string name, AccessKeyArgs args, CustomResourceOptions? options = null)
            : base("aws:iam/accessKey:AccessKey", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private AccessKey(string name, Input<string> id, AccessKeyState? state = null, CustomResourceOptions? options = null)
            : base("aws:iam/accessKey:AccessKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccessKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessKey Get(string name, Input<string> id, AccessKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new AccessKey(name, id, state, options);
        }
    }

    public sealed class AccessKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Either a base-64 encoded PGP public key, or a
        /// keybase username in the form `keybase:some_person_that_exists`.
        /// </summary>
        [Input("pgpKey")]
        public Input<string>? PgpKey { get; set; }

        /// <summary>
        /// The access key status to apply. Defaults to `Active`.
        /// Valid values are `Active` and `Inactive`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The IAM user to associate with this access key.
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        public AccessKeyArgs()
        {
        }
    }

    public sealed class AccessKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The encrypted secret, base64 encoded.
        /// &gt; **NOTE:** The encrypted secret may be decrypted using the command line,
        /// for example: `... | base64 --decode | keybase pgp decrypt`.
        /// </summary>
        [Input("encryptedSecret")]
        public Input<string>? EncryptedSecret { get; set; }

        /// <summary>
        /// The fingerprint of the PGP key used to encrypt
        /// the secret
        /// </summary>
        [Input("keyFingerprint")]
        public Input<string>? KeyFingerprint { get; set; }

        /// <summary>
        /// Either a base-64 encoded PGP public key, or a
        /// keybase username in the form `keybase:some_person_that_exists`.
        /// </summary>
        [Input("pgpKey")]
        public Input<string>? PgpKey { get; set; }

        /// <summary>
        /// The secret access key. Note that this will be written
        /// to the state file. Please supply a `pgp_key` instead, which will prevent the
        /// secret from being stored in plain text
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        /// <summary>
        /// The secret access key converted into an SES SMTP
        /// password by applying [AWS's documented conversion
        /// algorithm](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/smtp-credentials.html#smtp-credentials-convert).
        /// </summary>
        [Input("sesSmtpPassword")]
        public Input<string>? SesSmtpPassword { get; set; }

        /// <summary>
        /// The access key status to apply. Defaults to `Active`.
        /// Valid values are `Active` and `Inactive`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The IAM user to associate with this access key.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public AccessKeyState()
        {
        }
    }
}
