// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Provides a resource to manage a Default VPC Routing Table.
    /// 
    /// Each VPC created in AWS comes with a Default Route Table that can be managed, but not
    /// destroyed. **This is an advanced resource**, and has special caveats to be aware
    /// of when using it. Please read this document in its entirety before using this
    /// resource. It is recommended you **do not** use both `aws.ec2.DefaultRouteTable` to
    /// manage the default route table **and** use the `aws.ec2.MainRouteTableAssociation`,
    /// due to possible conflict in routes.
    /// 
    /// The `aws.ec2.DefaultRouteTable` behaves differently from normal resources, in that
    /// this provider does not _create_ this resource, but instead attempts to "adopt" it
    /// into management. We can do this because each VPC created has a Default Route
    /// Table that cannot be destroyed, and is created with a single route.
    /// 
    /// When this provider first adopts the Default Route Table, it **immediately removes all
    /// defined routes**. It then proceeds to create any routes specified in the
    /// configuration. This step is required so that only the routes specified in the
    /// configuration present in the Default Route Table.
    /// 
    /// For more information about Route Tables, see the AWS Documentation on
    /// [Route Tables][aws-route-tables].
    /// 
    /// For more information about managing normal Route Tables in this provider, see our
    /// documentation on [aws.ec2.RouteTable][tf-route-tables].
    /// 
    /// &gt; **NOTE on Route Tables and Routes:** This provider currently
    /// provides both a standalone Route resource and a Route Table resource with routes
    /// defined in-line. At this time you cannot use a Route Table with in-line routes
    /// in conjunction with any Route resources. Doing so will cause
    /// a conflict of rule settings and will overwrite routes.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/default_route_table.html.markdown.
    /// </summary>
    public partial class DefaultRouteTable : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Default Routing Table.
        /// </summary>
        [Output("defaultRouteTableId")]
        public Output<string> DefaultRouteTableId { get; private set; } = null!;

        /// <summary>
        /// The ID of the AWS account that owns the route table
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// A list of virtual gateways for propagation.
        /// </summary>
        [Output("propagatingVgws")]
        public Output<ImmutableArray<string>> PropagatingVgws { get; private set; } = null!;

        /// <summary>
        /// A list of route objects. Their keys are documented below.
        /// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        [Output("routes")]
        public Output<ImmutableArray<Outputs.DefaultRouteTableRoutes>> Routes { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a DefaultRouteTable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DefaultRouteTable(string name, DefaultRouteTableArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/defaultRouteTable:DefaultRouteTable", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private DefaultRouteTable(string name, Input<string> id, DefaultRouteTableState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/defaultRouteTable:DefaultRouteTable", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DefaultRouteTable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DefaultRouteTable Get(string name, Input<string> id, DefaultRouteTableState? state = null, CustomResourceOptions? options = null)
        {
            return new DefaultRouteTable(name, id, state, options);
        }
    }

    public sealed class DefaultRouteTableArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Default Routing Table.
        /// </summary>
        [Input("defaultRouteTableId", required: true)]
        public Input<string> DefaultRouteTableId { get; set; } = null!;

        [Input("propagatingVgws")]
        private InputList<string>? _propagatingVgws;

        /// <summary>
        /// A list of virtual gateways for propagation.
        /// </summary>
        public InputList<string> PropagatingVgws
        {
            get => _propagatingVgws ?? (_propagatingVgws = new InputList<string>());
            set => _propagatingVgws = value;
        }

        [Input("routes")]
        private InputList<Inputs.DefaultRouteTableRoutesArgs>? _routes;

        /// <summary>
        /// A list of route objects. Their keys are documented below.
        /// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        public InputList<Inputs.DefaultRouteTableRoutesArgs> Routes
        {
            get => _routes ?? (_routes = new InputList<Inputs.DefaultRouteTableRoutesArgs>());
            set => _routes = value;
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

        public DefaultRouteTableArgs()
        {
        }
    }

    public sealed class DefaultRouteTableState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Default Routing Table.
        /// </summary>
        [Input("defaultRouteTableId")]
        public Input<string>? DefaultRouteTableId { get; set; }

        /// <summary>
        /// The ID of the AWS account that owns the route table
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        [Input("propagatingVgws")]
        private InputList<string>? _propagatingVgws;

        /// <summary>
        /// A list of virtual gateways for propagation.
        /// </summary>
        public InputList<string> PropagatingVgws
        {
            get => _propagatingVgws ?? (_propagatingVgws = new InputList<string>());
            set => _propagatingVgws = value;
        }

        [Input("routes")]
        private InputList<Inputs.DefaultRouteTableRoutesGetArgs>? _routes;

        /// <summary>
        /// A list of route objects. Their keys are documented below.
        /// This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
        /// </summary>
        public InputList<Inputs.DefaultRouteTableRoutesGetArgs> Routes
        {
            get => _routes ?? (_routes = new InputList<Inputs.DefaultRouteTableRoutesGetArgs>());
            set => _routes = value;
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

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public DefaultRouteTableState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class DefaultRouteTableRoutesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CIDR block of the route.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Identifier of a VPC Egress Only Internet Gateway.
        /// </summary>
        [Input("egressOnlyGatewayId")]
        public Input<string>? EgressOnlyGatewayId { get; set; }

        /// <summary>
        /// Identifier of a VPC internet gateway or a virtual private gateway.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        /// <summary>
        /// Identifier of an EC2 instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The Ipv6 CIDR block of the route
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// Identifier of a VPC NAT gateway.
        /// </summary>
        [Input("natGatewayId")]
        public Input<string>? NatGatewayId { get; set; }

        /// <summary>
        /// Identifier of an EC2 network interface.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// Identifier of an EC2 Transit Gateway.
        /// </summary>
        [Input("transitGatewayId")]
        public Input<string>? TransitGatewayId { get; set; }

        /// <summary>
        /// Identifier of a VPC peering connection.
        /// </summary>
        [Input("vpcPeeringConnectionId")]
        public Input<string>? VpcPeeringConnectionId { get; set; }

        public DefaultRouteTableRoutesArgs()
        {
        }
    }

    public sealed class DefaultRouteTableRoutesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CIDR block of the route.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Identifier of a VPC Egress Only Internet Gateway.
        /// </summary>
        [Input("egressOnlyGatewayId")]
        public Input<string>? EgressOnlyGatewayId { get; set; }

        /// <summary>
        /// Identifier of a VPC internet gateway or a virtual private gateway.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        /// <summary>
        /// Identifier of an EC2 instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The Ipv6 CIDR block of the route
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// Identifier of a VPC NAT gateway.
        /// </summary>
        [Input("natGatewayId")]
        public Input<string>? NatGatewayId { get; set; }

        /// <summary>
        /// Identifier of an EC2 network interface.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// Identifier of an EC2 Transit Gateway.
        /// </summary>
        [Input("transitGatewayId")]
        public Input<string>? TransitGatewayId { get; set; }

        /// <summary>
        /// Identifier of a VPC peering connection.
        /// </summary>
        [Input("vpcPeeringConnectionId")]
        public Input<string>? VpcPeeringConnectionId { get; set; }

        public DefaultRouteTableRoutesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class DefaultRouteTableRoutes
    {
        /// <summary>
        /// The CIDR block of the route.
        /// </summary>
        public readonly string? CidrBlock;
        /// <summary>
        /// Identifier of a VPC Egress Only Internet Gateway.
        /// </summary>
        public readonly string? EgressOnlyGatewayId;
        /// <summary>
        /// Identifier of a VPC internet gateway or a virtual private gateway.
        /// </summary>
        public readonly string? GatewayId;
        /// <summary>
        /// Identifier of an EC2 instance.
        /// </summary>
        public readonly string? InstanceId;
        /// <summary>
        /// The Ipv6 CIDR block of the route
        /// </summary>
        public readonly string? Ipv6CidrBlock;
        /// <summary>
        /// Identifier of a VPC NAT gateway.
        /// </summary>
        public readonly string? NatGatewayId;
        /// <summary>
        /// Identifier of an EC2 network interface.
        /// </summary>
        public readonly string? NetworkInterfaceId;
        /// <summary>
        /// Identifier of an EC2 Transit Gateway.
        /// </summary>
        public readonly string? TransitGatewayId;
        /// <summary>
        /// Identifier of a VPC peering connection.
        /// </summary>
        public readonly string? VpcPeeringConnectionId;

        [OutputConstructor]
        private DefaultRouteTableRoutes(
            string? cidrBlock,
            string? egressOnlyGatewayId,
            string? gatewayId,
            string? instanceId,
            string? ipv6CidrBlock,
            string? natGatewayId,
            string? networkInterfaceId,
            string? transitGatewayId,
            string? vpcPeeringConnectionId)
        {
            CidrBlock = cidrBlock;
            EgressOnlyGatewayId = egressOnlyGatewayId;
            GatewayId = gatewayId;
            InstanceId = instanceId;
            Ipv6CidrBlock = ipv6CidrBlock;
            NatGatewayId = natGatewayId;
            NetworkInterfaceId = networkInterfaceId;
            TransitGatewayId = transitGatewayId;
            VpcPeeringConnectionId = vpcPeeringConnectionId;
        }
    }
    }
}
