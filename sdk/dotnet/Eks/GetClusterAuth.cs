// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Eks
{
    public static partial class Invokes
    {
        /// <summary>
        /// Get an authentication token to communicate with an EKS cluster.
        /// 
        /// Uses IAM credentials from the AWS provider to generate a temporary token that is compatible with
        /// [AWS IAM Authenticator](https://github.com/kubernetes-sigs/aws-iam-authenticator) authentication.
        /// This can be used to authenticate to an EKS cluster or to a cluster that has the AWS IAM Authenticator
        /// server configured.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/eks_cluster_auth.html.markdown.
        /// </summary>
        public static Task<GetClusterAuthResult> GetClusterAuth(GetClusterAuthArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterAuthResult>("aws:eks/getClusterAuth:getClusterAuth", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetClusterAuthArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the cluster
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetClusterAuthArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetClusterAuthResult
    {
        public readonly string Name;
        /// <summary>
        /// The token to use to authenticate with the cluster.
        /// </summary>
        public readonly string Token;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetClusterAuthResult(
            string name,
            string token,
            string id)
        {
            Name = name;
            Token = token;
            Id = id;
        }
    }
}
