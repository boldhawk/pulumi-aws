// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GuardDuty
{
    /// <summary>
    /// Provides a resource to accept a pending GuardDuty invite on creation, ensure the detector has the correct master account on read, and disassociate with the master account upon removal.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/guardduty_invite_accepter.html.markdown.
    /// </summary>
    public partial class InviteAccepter : Pulumi.CustomResource
    {
        /// <summary>
        /// The detector ID of the member GuardDuty account.
        /// </summary>
        [Output("detectorId")]
        public Output<string> DetectorId { get; private set; } = null!;

        /// <summary>
        /// AWS account ID for master account.
        /// </summary>
        [Output("masterAccountId")]
        public Output<string> MasterAccountId { get; private set; } = null!;


        /// <summary>
        /// Create a InviteAccepter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InviteAccepter(string name, InviteAccepterArgs args, CustomResourceOptions? options = null)
            : base("aws:guardduty/inviteAccepter:InviteAccepter", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private InviteAccepter(string name, Input<string> id, InviteAccepterState? state = null, CustomResourceOptions? options = null)
            : base("aws:guardduty/inviteAccepter:InviteAccepter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InviteAccepter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InviteAccepter Get(string name, Input<string> id, InviteAccepterState? state = null, CustomResourceOptions? options = null)
        {
            return new InviteAccepter(name, id, state, options);
        }
    }

    public sealed class InviteAccepterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The detector ID of the member GuardDuty account.
        /// </summary>
        [Input("detectorId", required: true)]
        public Input<string> DetectorId { get; set; } = null!;

        /// <summary>
        /// AWS account ID for master account.
        /// </summary>
        [Input("masterAccountId", required: true)]
        public Input<string> MasterAccountId { get; set; } = null!;

        public InviteAccepterArgs()
        {
        }
    }

    public sealed class InviteAccepterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The detector ID of the member GuardDuty account.
        /// </summary>
        [Input("detectorId")]
        public Input<string>? DetectorId { get; set; }

        /// <summary>
        /// AWS account ID for master account.
        /// </summary>
        [Input("masterAccountId")]
        public Input<string>? MasterAccountId { get; set; }

        public InviteAccepterState()
        {
        }
    }
}
