// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3
{
    /// <summary>
    /// Attaches a policy to an S3 bucket resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/s3_bucket_policy.html.markdown.
    /// </summary>
    public partial class BucketPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the bucket to which to apply the policy.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;


        /// <summary>
        /// Create a BucketPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketPolicy(string name, BucketPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:s3/bucketPolicy:BucketPolicy", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private BucketPolicy(string name, Input<string> id, BucketPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:s3/bucketPolicy:BucketPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BucketPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketPolicy Get(string name, Input<string> id, BucketPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new BucketPolicy(name, id, state, options);
        }
    }

    public sealed class BucketPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket to which to apply the policy.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        public BucketPolicyArgs()
        {
        }
    }

    public sealed class BucketPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket to which to apply the policy.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("policy")]
        public Input<string>? Policy { get; set; }

        public BucketPolicyState()
        {
        }
    }
}
