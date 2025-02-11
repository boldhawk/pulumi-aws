// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2TransitGateway
{
    public static partial class Invokes
    {
        /// <summary>
        /// Get information on an EC2 Transit Gateway VPN Attachment.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ec2_transit_gateway_vpn_attachment.html.markdown.
        /// </summary>
        public static Task<GetVpnAttachmentResult> GetVpnAttachment(GetVpnAttachmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpnAttachmentResult>("aws:ec2transitgateway/getVpnAttachment:getVpnAttachment", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetVpnAttachmentArgs : Pulumi.ResourceArgs
    {
        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Identifier of the EC2 Transit Gateway.
        /// </summary>
        [Input("transitGatewayId", required: true)]
        public Input<string> TransitGatewayId { get; set; } = null!;

        /// <summary>
        /// Identifier of the EC2 VPN Connection.
        /// </summary>
        [Input("vpnConnectionId", required: true)]
        public Input<string> VpnConnectionId { get; set; } = null!;

        public GetVpnAttachmentArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetVpnAttachmentResult
    {
        /// <summary>
        /// Key-value tags for the EC2 Transit Gateway VPN Attachment
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        public readonly string TransitGatewayId;
        public readonly string VpnConnectionId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetVpnAttachmentResult(
            ImmutableDictionary<string, object> tags,
            string transitGatewayId,
            string vpnConnectionId,
            string id)
        {
            Tags = tags;
            TransitGatewayId = transitGatewayId;
            VpnConnectionId = vpnConnectionId;
            Id = id;
        }
    }
}
