// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront
{
    /// <summary>
    /// Creates an Amazon CloudFront origin access identity.
    /// 
    /// For information about CloudFront distributions, see the
    /// [Amazon CloudFront Developer Guide][1]. For more information on generating
    /// origin access identities, see
    /// [Using an Origin Access Identity to Restrict Access to Your Amazon S3 Content][2].
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudfront_origin_access_identity.html.markdown.
    /// </summary>
    public partial class OriginAccessIdentity : Pulumi.CustomResource
    {
        /// <summary>
        /// Internal value used by CloudFront to allow future
        /// updates to the origin access identity.
        /// </summary>
        [Output("callerReference")]
        public Output<string> CallerReference { get; private set; } = null!;

        /// <summary>
        /// A shortcut to the full path for the
        /// origin access identity to use in CloudFront, see below.
        /// </summary>
        [Output("cloudfrontAccessIdentityPath")]
        public Output<string> CloudfrontAccessIdentityPath { get; private set; } = null!;

        /// <summary>
        /// An optional comment for the origin access identity.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// The current version of the origin access identity's information.
        /// For example: `E2QWRUHAPOMQZL`.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// A pre-generated ARN for use in S3 bucket policies (see below).
        /// Example: `arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity
        /// E2QWRUHAPOMQZL`.
        /// </summary>
        [Output("iamArn")]
        public Output<string> IamArn { get; private set; } = null!;

        /// <summary>
        /// The Amazon S3 canonical user ID for the origin
        /// access identity, which you use when giving the origin access identity read
        /// permission to an object in Amazon S3.
        /// </summary>
        [Output("s3CanonicalUserId")]
        public Output<string> S3CanonicalUserId { get; private set; } = null!;


        /// <summary>
        /// Create a OriginAccessIdentity resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OriginAccessIdentity(string name, OriginAccessIdentityArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:cloudfront/originAccessIdentity:OriginAccessIdentity", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private OriginAccessIdentity(string name, Input<string> id, OriginAccessIdentityState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudfront/originAccessIdentity:OriginAccessIdentity", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OriginAccessIdentity resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OriginAccessIdentity Get(string name, Input<string> id, OriginAccessIdentityState? state = null, CustomResourceOptions? options = null)
        {
            return new OriginAccessIdentity(name, id, state, options);
        }
    }

    public sealed class OriginAccessIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional comment for the origin access identity.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        public OriginAccessIdentityArgs()
        {
        }
    }

    public sealed class OriginAccessIdentityState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Internal value used by CloudFront to allow future
        /// updates to the origin access identity.
        /// </summary>
        [Input("callerReference")]
        public Input<string>? CallerReference { get; set; }

        /// <summary>
        /// A shortcut to the full path for the
        /// origin access identity to use in CloudFront, see below.
        /// </summary>
        [Input("cloudfrontAccessIdentityPath")]
        public Input<string>? CloudfrontAccessIdentityPath { get; set; }

        /// <summary>
        /// An optional comment for the origin access identity.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// The current version of the origin access identity's information.
        /// For example: `E2QWRUHAPOMQZL`.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// A pre-generated ARN for use in S3 bucket policies (see below).
        /// Example: `arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity
        /// E2QWRUHAPOMQZL`.
        /// </summary>
        [Input("iamArn")]
        public Input<string>? IamArn { get; set; }

        /// <summary>
        /// The Amazon S3 canonical user ID for the origin
        /// access identity, which you use when giving the origin access identity read
        /// permission to an object in Amazon S3.
        /// </summary>
        [Input("s3CanonicalUserId")]
        public Input<string>? S3CanonicalUserId { get; set; }

        public OriginAccessIdentityState()
        {
        }
    }
}
