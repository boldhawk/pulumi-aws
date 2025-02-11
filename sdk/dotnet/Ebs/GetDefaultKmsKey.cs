// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ebs
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get the default EBS encryption KMS key in the current region.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ebs_default_kms_key.html.markdown.
        /// </summary>
        public static Task<GetDefaultKmsKeyResult> GetDefaultKmsKey(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDefaultKmsKeyResult>("aws:ebs/getDefaultKmsKey:getDefaultKmsKey", ResourceArgs.Empty, options.WithVersion());
    }

    [OutputType]
    public sealed class GetDefaultKmsKeyResult
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the default KMS key uses to encrypt an EBS volume in this region when no key is specified in an API call that creates the volume and encryption by default is enabled.
        /// </summary>
        public readonly string KeyArn;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDefaultKmsKeyResult(
            string keyArn,
            string id)
        {
            KeyArn = keyArn;
            Id = id;
        }
    }
}
