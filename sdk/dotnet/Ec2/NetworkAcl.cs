// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Provides an network ACL resource. You might set up network ACLs with rules similar
    /// to your security groups in order to add an additional layer of security to your VPC.
    /// 
    /// &gt; **NOTE on Network ACLs and Network ACL Rules:** This provider currently
    /// provides both a standalone Network ACL Rule resource and a Network ACL resource with rules
    /// defined in-line. At this time you cannot use a Network ACL with in-line rules
    /// in conjunction with any Network ACL Rule resources. Doing so will cause
    /// a conflict of rule settings and will overwrite rules.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/network_acl.html.markdown.
    /// </summary>
    public partial class NetworkAcl : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies an egress rule. Parameters defined below.
        /// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        [Output("egress")]
        public Output<ImmutableArray<Outputs.NetworkAclEgress>> Egress { get; private set; } = null!;

        /// <summary>
        /// Specifies an ingress rule. Parameters defined below.
        /// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        [Output("ingress")]
        public Output<ImmutableArray<Outputs.NetworkAclIngress>> Ingress { get; private set; } = null!;

        /// <summary>
        /// The ID of the AWS account that owns the network ACL.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// A list of Subnet IDs to apply the ACL to
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkAcl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkAcl(string name, NetworkAclArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/networkAcl:NetworkAcl", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private NetworkAcl(string name, Input<string> id, NetworkAclState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/networkAcl:NetworkAcl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkAcl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkAcl Get(string name, Input<string> id, NetworkAclState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkAcl(name, id, state, options);
        }
    }

    public sealed class NetworkAclArgs : Pulumi.ResourceArgs
    {
        [Input("egress")]
        private InputList<Inputs.NetworkAclEgressArgs>? _egress;

        /// <summary>
        /// Specifies an egress rule. Parameters defined below.
        /// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        public InputList<Inputs.NetworkAclEgressArgs> Egress
        {
            get => _egress ?? (_egress = new InputList<Inputs.NetworkAclEgressArgs>());
            set => _egress = value;
        }

        [Input("ingress")]
        private InputList<Inputs.NetworkAclIngressArgs>? _ingress;

        /// <summary>
        /// Specifies an ingress rule. Parameters defined below.
        /// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        public InputList<Inputs.NetworkAclIngressArgs> Ingress
        {
            get => _ingress ?? (_ingress = new InputList<Inputs.NetworkAclIngressArgs>());
            set => _ingress = value;
        }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// A list of Subnet IDs to apply the ACL to
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the associated VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public NetworkAclArgs()
        {
        }
    }

    public sealed class NetworkAclState : Pulumi.ResourceArgs
    {
        [Input("egress")]
        private InputList<Inputs.NetworkAclEgressGetArgs>? _egress;

        /// <summary>
        /// Specifies an egress rule. Parameters defined below.
        /// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        public InputList<Inputs.NetworkAclEgressGetArgs> Egress
        {
            get => _egress ?? (_egress = new InputList<Inputs.NetworkAclEgressGetArgs>());
            set => _egress = value;
        }

        [Input("ingress")]
        private InputList<Inputs.NetworkAclIngressGetArgs>? _ingress;

        /// <summary>
        /// Specifies an ingress rule. Parameters defined below.
        /// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        public InputList<Inputs.NetworkAclIngressGetArgs> Ingress
        {
            get => _ingress ?? (_ingress = new InputList<Inputs.NetworkAclIngressGetArgs>());
            set => _ingress = value;
        }

        /// <summary>
        /// The ID of the AWS account that owns the network ACL.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// A list of Subnet IDs to apply the ACL to
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the associated VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public NetworkAclState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class NetworkAclEgressArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to take.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// The CIDR block to match. This must be a
        /// valid network mask.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// The from port to match.
        /// </summary>
        [Input("fromPort", required: true)]
        public Input<int> FromPort { get; set; } = null!;

        /// <summary>
        /// The ICMP type code to be used. Default 0.
        /// </summary>
        [Input("icmpCode")]
        public Input<int>? IcmpCode { get; set; }

        /// <summary>
        /// The ICMP type to be used. Default 0.
        /// </summary>
        [Input("icmpType")]
        public Input<int>? IcmpType { get; set; }

        /// <summary>
        /// The IPv6 CIDR block.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// The protocol to match. If using the -1 'all'
        /// protocol, you must specify a from and to port of 0.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The rule number. Used for ordering.
        /// </summary>
        [Input("ruleNo", required: true)]
        public Input<int> RuleNo { get; set; } = null!;

        /// <summary>
        /// The to port to match.
        /// </summary>
        [Input("toPort", required: true)]
        public Input<int> ToPort { get; set; } = null!;

        public NetworkAclEgressArgs()
        {
        }
    }

    public sealed class NetworkAclEgressGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to take.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// The CIDR block to match. This must be a
        /// valid network mask.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// The from port to match.
        /// </summary>
        [Input("fromPort", required: true)]
        public Input<int> FromPort { get; set; } = null!;

        /// <summary>
        /// The ICMP type code to be used. Default 0.
        /// </summary>
        [Input("icmpCode")]
        public Input<int>? IcmpCode { get; set; }

        /// <summary>
        /// The ICMP type to be used. Default 0.
        /// </summary>
        [Input("icmpType")]
        public Input<int>? IcmpType { get; set; }

        /// <summary>
        /// The IPv6 CIDR block.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// The protocol to match. If using the -1 'all'
        /// protocol, you must specify a from and to port of 0.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The rule number. Used for ordering.
        /// </summary>
        [Input("ruleNo", required: true)]
        public Input<int> RuleNo { get; set; } = null!;

        /// <summary>
        /// The to port to match.
        /// </summary>
        [Input("toPort", required: true)]
        public Input<int> ToPort { get; set; } = null!;

        public NetworkAclEgressGetArgs()
        {
        }
    }

    public sealed class NetworkAclIngressArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to take.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// The CIDR block to match. This must be a
        /// valid network mask.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// The from port to match.
        /// </summary>
        [Input("fromPort", required: true)]
        public Input<int> FromPort { get; set; } = null!;

        /// <summary>
        /// The ICMP type code to be used. Default 0.
        /// </summary>
        [Input("icmpCode")]
        public Input<int>? IcmpCode { get; set; }

        /// <summary>
        /// The ICMP type to be used. Default 0.
        /// </summary>
        [Input("icmpType")]
        public Input<int>? IcmpType { get; set; }

        /// <summary>
        /// The IPv6 CIDR block.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// The protocol to match. If using the -1 'all'
        /// protocol, you must specify a from and to port of 0.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The rule number. Used for ordering.
        /// </summary>
        [Input("ruleNo", required: true)]
        public Input<int> RuleNo { get; set; } = null!;

        /// <summary>
        /// The to port to match.
        /// </summary>
        [Input("toPort", required: true)]
        public Input<int> ToPort { get; set; } = null!;

        public NetworkAclIngressArgs()
        {
        }
    }

    public sealed class NetworkAclIngressGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to take.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// The CIDR block to match. This must be a
        /// valid network mask.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// The from port to match.
        /// </summary>
        [Input("fromPort", required: true)]
        public Input<int> FromPort { get; set; } = null!;

        /// <summary>
        /// The ICMP type code to be used. Default 0.
        /// </summary>
        [Input("icmpCode")]
        public Input<int>? IcmpCode { get; set; }

        /// <summary>
        /// The ICMP type to be used. Default 0.
        /// </summary>
        [Input("icmpType")]
        public Input<int>? IcmpType { get; set; }

        /// <summary>
        /// The IPv6 CIDR block.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// The protocol to match. If using the -1 'all'
        /// protocol, you must specify a from and to port of 0.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The rule number. Used for ordering.
        /// </summary>
        [Input("ruleNo", required: true)]
        public Input<int> RuleNo { get; set; } = null!;

        /// <summary>
        /// The to port to match.
        /// </summary>
        [Input("toPort", required: true)]
        public Input<int> ToPort { get; set; } = null!;

        public NetworkAclIngressGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class NetworkAclEgress
    {
        /// <summary>
        /// The action to take.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// The CIDR block to match. This must be a
        /// valid network mask.
        /// </summary>
        public readonly string? CidrBlock;
        /// <summary>
        /// The from port to match.
        /// </summary>
        public readonly int FromPort;
        /// <summary>
        /// The ICMP type code to be used. Default 0.
        /// </summary>
        public readonly int? IcmpCode;
        /// <summary>
        /// The ICMP type to be used. Default 0.
        /// </summary>
        public readonly int? IcmpType;
        /// <summary>
        /// The IPv6 CIDR block.
        /// </summary>
        public readonly string? Ipv6CidrBlock;
        /// <summary>
        /// The protocol to match. If using the -1 'all'
        /// protocol, you must specify a from and to port of 0.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The rule number. Used for ordering.
        /// </summary>
        public readonly int RuleNo;
        /// <summary>
        /// The to port to match.
        /// </summary>
        public readonly int ToPort;

        [OutputConstructor]
        private NetworkAclEgress(
            string action,
            string? cidrBlock,
            int fromPort,
            int? icmpCode,
            int? icmpType,
            string? ipv6CidrBlock,
            string protocol,
            int ruleNo,
            int toPort)
        {
            Action = action;
            CidrBlock = cidrBlock;
            FromPort = fromPort;
            IcmpCode = icmpCode;
            IcmpType = icmpType;
            Ipv6CidrBlock = ipv6CidrBlock;
            Protocol = protocol;
            RuleNo = ruleNo;
            ToPort = toPort;
        }
    }

    [OutputType]
    public sealed class NetworkAclIngress
    {
        /// <summary>
        /// The action to take.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// The CIDR block to match. This must be a
        /// valid network mask.
        /// </summary>
        public readonly string? CidrBlock;
        /// <summary>
        /// The from port to match.
        /// </summary>
        public readonly int FromPort;
        /// <summary>
        /// The ICMP type code to be used. Default 0.
        /// </summary>
        public readonly int? IcmpCode;
        /// <summary>
        /// The ICMP type to be used. Default 0.
        /// </summary>
        public readonly int? IcmpType;
        /// <summary>
        /// The IPv6 CIDR block.
        /// </summary>
        public readonly string? Ipv6CidrBlock;
        /// <summary>
        /// The protocol to match. If using the -1 'all'
        /// protocol, you must specify a from and to port of 0.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The rule number. Used for ordering.
        /// </summary>
        public readonly int RuleNo;
        /// <summary>
        /// The to port to match.
        /// </summary>
        public readonly int ToPort;

        [OutputConstructor]
        private NetworkAclIngress(
            string action,
            string? cidrBlock,
            int fromPort,
            int? icmpCode,
            int? icmpType,
            string? ipv6CidrBlock,
            string protocol,
            int ruleNo,
            int toPort)
        {
            Action = action;
            CidrBlock = cidrBlock;
            FromPort = fromPort;
            IcmpCode = icmpCode;
            IcmpType = icmpType;
            Ipv6CidrBlock = ipv6CidrBlock;
            Protocol = protocol;
            RuleNo = ruleNo;
            ToPort = toPort;
        }
    }
    }
}
