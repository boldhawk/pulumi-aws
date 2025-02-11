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
        /// `aws.wafregional.Rule` Retrieves a WAF Regional Rule Resource Id.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/wafregional_rule.html.markdown.
        /// </summary>
        public static Task<GetRuleResult> GetRule(GetRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRuleResult>("aws:wafregional/getRule:getRule", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the WAF Regional rule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetRuleArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetRuleResult
    {
        public readonly string Name;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetRuleResult(
            string name,
            string id)
        {
            Name = name;
            Id = id;
        }
    }
}
