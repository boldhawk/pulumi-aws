// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ram
{
    /// <summary>
    /// Provides a Resource Access Manager (RAM) principal association. Depending if [RAM Sharing with AWS Organizations is enabled](https://docs.aws.amazon.com/ram/latest/userguide/getting-started-sharing.html#getting-started-sharing-orgs), the RAM behavior with different principal types changes.
    /// 
    /// When RAM Sharing with AWS Organizations is enabled:
    /// 
    /// - For AWS Account ID, Organization, and Organizational Unit principals within the same AWS Organization, no resource share invitation is sent and resources become available automatically after creating the association.
    /// - For AWS Account ID principals outside the AWS Organization, a resource share invitation is sent and must be accepted before resources become available. See the [`aws.ram.ResourceShareAccepter` resource](https://www.terraform.io/docs/providers/aws/r/ram_resource_share_accepter.html) to accept these invitations.
    /// 
    /// When RAM Sharing with AWS Organizations is not enabled:
    /// 
    /// - Organization and Organizational Unit principals cannot be used.
    /// - For AWS Account ID principals, a resource share invitation is sent and must be accepted before resources become available. See the [`aws.ram.ResourceShareAccepter` resource](https://www.terraform.io/docs/providers/aws/r/ram_resource_share_accepter.html) to accept these invitations.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ram_principal_association.html.markdown.
    /// </summary>
    public partial class PrincipalAssociation : Pulumi.CustomResource
    {
        /// <summary>
        /// The principal to associate with the resource share. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN.
        /// </summary>
        [Output("principal")]
        public Output<string> Principal { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the resource share.
        /// </summary>
        [Output("resourceShareArn")]
        public Output<string> ResourceShareArn { get; private set; } = null!;


        /// <summary>
        /// Create a PrincipalAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrincipalAssociation(string name, PrincipalAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws:ram/principalAssociation:PrincipalAssociation", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private PrincipalAssociation(string name, Input<string> id, PrincipalAssociationState? state = null, CustomResourceOptions? options = null)
            : base("aws:ram/principalAssociation:PrincipalAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PrincipalAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrincipalAssociation Get(string name, Input<string> id, PrincipalAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new PrincipalAssociation(name, id, state, options);
        }
    }

    public sealed class PrincipalAssociationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The principal to associate with the resource share. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN.
        /// </summary>
        [Input("principal", required: true)]
        public Input<string> Principal { get; set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the resource share.
        /// </summary>
        [Input("resourceShareArn", required: true)]
        public Input<string> ResourceShareArn { get; set; } = null!;

        public PrincipalAssociationArgs()
        {
        }
    }

    public sealed class PrincipalAssociationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The principal to associate with the resource share. Possible values are an AWS account ID, an AWS Organizations Organization ARN, or an AWS Organizations Organization Unit ARN.
        /// </summary>
        [Input("principal")]
        public Input<string>? Principal { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the resource share.
        /// </summary>
        [Input("resourceShareArn")]
        public Input<string>? ResourceShareArn { get; set; }

        public PrincipalAssociationState()
        {
        }
    }
}
