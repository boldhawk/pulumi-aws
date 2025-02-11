// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafRegional
{
    public static partial class Invokes
    {
        /// <summary>
        /// `aws.wafregional.IpSet` Retrieves a WAF Regional IP Set Resource Id.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/wafregional_ipset.html.markdown.
        /// </summary>
        public static Task<GetIpsetResult> GetIpset(GetIpsetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIpsetResult>("aws:wafregional/getIpset:getIpset", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetIpsetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the WAF Regional IP set.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetIpsetArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetIpsetResult
    {
        public readonly string Name;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetIpsetResult(
            string name,
            string id)
        {
            Name = name;
            Id = id;
        }
    }
}
