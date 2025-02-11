// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Pinpoint
{
    /// <summary>
    /// Provides a Pinpoint Baidu Channel resource.
    /// 
    /// &gt; **Note:** All arguments including the Api Key and Secret Key will be stored in the raw state as plain-text.
    /// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/pinpoint_baidu_channel.html.markdown.
    /// </summary>
    public partial class BaiduChannel : Pulumi.CustomResource
    {
        /// <summary>
        /// Platform credential API key from Baidu.
        /// </summary>
        [Output("apiKey")]
        public Output<string> ApiKey { get; private set; } = null!;

        /// <summary>
        /// The application ID.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable the channel. Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Platform credential Secret key from Baidu.
        /// </summary>
        [Output("secretKey")]
        public Output<string> SecretKey { get; private set; } = null!;


        /// <summary>
        /// Create a BaiduChannel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BaiduChannel(string name, BaiduChannelArgs args, CustomResourceOptions? options = null)
            : base("aws:pinpoint/baiduChannel:BaiduChannel", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private BaiduChannel(string name, Input<string> id, BaiduChannelState? state = null, CustomResourceOptions? options = null)
            : base("aws:pinpoint/baiduChannel:BaiduChannel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BaiduChannel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BaiduChannel Get(string name, Input<string> id, BaiduChannelState? state = null, CustomResourceOptions? options = null)
        {
            return new BaiduChannel(name, id, state, options);
        }
    }

    public sealed class BaiduChannelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Platform credential API key from Baidu.
        /// </summary>
        [Input("apiKey", required: true)]
        public Input<string> ApiKey { get; set; } = null!;

        /// <summary>
        /// The application ID.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// Specifies whether to enable the channel. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Platform credential Secret key from Baidu.
        /// </summary>
        [Input("secretKey", required: true)]
        public Input<string> SecretKey { get; set; } = null!;

        public BaiduChannelArgs()
        {
        }
    }

    public sealed class BaiduChannelState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Platform credential API key from Baidu.
        /// </summary>
        [Input("apiKey")]
        public Input<string>? ApiKey { get; set; }

        /// <summary>
        /// The application ID.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// Specifies whether to enable the channel. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Platform credential Secret key from Baidu.
        /// </summary>
        [Input("secretKey")]
        public Input<string>? SecretKey { get; set; }

        public BaiduChannelState()
        {
        }
    }
}
