// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Provides a VPC Endpoint resource.
    /// 
    /// &gt; **NOTE on VPC Endpoints and VPC Endpoint Associations:** This provider provides both standalone VPC Endpoint Associations for
    /// Route Tables - (an association between a VPC endpoint and a single `route_table_id`) and
    /// Subnets - (an association between a VPC endpoint and a single `subnet_id`) and
    /// a VPC Endpoint resource with `route_table_ids` and `subnet_ids` attributes.
    /// Do not use the same resource ID in both a VPC Endpoint resource and a VPC Endpoint Association resource.
    /// Doing so will cause a conflict of associations and will overwrite the association.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpc_endpoint.html.markdown.
    /// </summary>
    public partial class VpcEndpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
        /// </summary>
        [Output("autoAccept")]
        public Output<bool?> AutoAccept { get; private set; } = null!;

        /// <summary>
        /// The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
        /// </summary>
        [Output("cidrBlocks")]
        public Output<ImmutableArray<string>> CidrBlocks { get; private set; } = null!;

        /// <summary>
        /// The DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS blocks are documented below.
        /// </summary>
        [Output("dnsEntries")]
        public Output<ImmutableArray<Outputs.VpcEndpointDnsEntries>> DnsEntries { get; private set; } = null!;

        /// <summary>
        /// One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
        /// </summary>
        [Output("networkInterfaceIds")]
        public Output<ImmutableArray<string>> NetworkInterfaceIds { get; private set; } = null!;

        /// <summary>
        /// The ID of the AWS account that owns the VPC endpoint.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;

        /// <summary>
        /// The prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
        /// </summary>
        [Output("prefixListId")]
        public Output<string> PrefixListId { get; private set; } = null!;

        /// <summary>
        /// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
        /// Defaults to `false`.
        /// </summary>
        [Output("privateDnsEnabled")]
        public Output<bool?> PrivateDnsEnabled { get; private set; } = null!;

        /// <summary>
        /// Whether or not the VPC Endpoint is being managed by its service - `true` or `false`.
        /// </summary>
        [Output("requesterManaged")]
        public Output<bool> RequesterManaged { get; private set; } = null!;

        /// <summary>
        /// One or more route table IDs. Applicable for endpoints of type `Gateway`.
        /// </summary>
        [Output("routeTableIds")]
        public Output<ImmutableArray<string>> RouteTableIds { get; private set; } = null!;

        /// <summary>
        /// The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// The service name, in the form `com.amazonaws.region.service` for AWS services.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// The state of the VPC endpoint.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `Interface`.
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The VPC endpoint type, `Gateway` or `Interface`. Defaults to `Gateway`.
        /// </summary>
        [Output("vpcEndpointType")]
        public Output<string?> VpcEndpointType { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC in which the endpoint will be used.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcEndpoint(string name, VpcEndpointArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/vpcEndpoint:VpcEndpoint", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private VpcEndpoint(string name, Input<string> id, VpcEndpointState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/vpcEndpoint:VpcEndpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcEndpoint Get(string name, Input<string> id, VpcEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcEndpoint(name, id, state, options);
        }
    }

    public sealed class VpcEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
        /// </summary>
        [Input("autoAccept")]
        public Input<bool>? AutoAccept { get; set; }

        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
        /// Defaults to `false`.
        /// </summary>
        [Input("privateDnsEnabled")]
        public Input<bool>? PrivateDnsEnabled { get; set; }

        [Input("routeTableIds")]
        private InputList<string>? _routeTableIds;

        /// <summary>
        /// One or more route table IDs. Applicable for endpoints of type `Gateway`.
        /// </summary>
        public InputList<string> RouteTableIds
        {
            get => _routeTableIds ?? (_routeTableIds = new InputList<string>());
            set => _routeTableIds = value;
        }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The service name, in the form `com.amazonaws.region.service` for AWS services.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `Interface`.
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
        /// The VPC endpoint type, `Gateway` or `Interface`. Defaults to `Gateway`.
        /// </summary>
        [Input("vpcEndpointType")]
        public Input<string>? VpcEndpointType { get; set; }

        /// <summary>
        /// The ID of the VPC in which the endpoint will be used.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public VpcEndpointArgs()
        {
        }
    }

    public sealed class VpcEndpointState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
        /// </summary>
        [Input("autoAccept")]
        public Input<bool>? AutoAccept { get; set; }

        [Input("cidrBlocks")]
        private InputList<string>? _cidrBlocks;

        /// <summary>
        /// The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
        /// </summary>
        public InputList<string> CidrBlocks
        {
            get => _cidrBlocks ?? (_cidrBlocks = new InputList<string>());
            set => _cidrBlocks = value;
        }

        [Input("dnsEntries")]
        private InputList<Inputs.VpcEndpointDnsEntriesGetArgs>? _dnsEntries;

        /// <summary>
        /// The DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS blocks are documented below.
        /// </summary>
        public InputList<Inputs.VpcEndpointDnsEntriesGetArgs> DnsEntries
        {
            get => _dnsEntries ?? (_dnsEntries = new InputList<Inputs.VpcEndpointDnsEntriesGetArgs>());
            set => _dnsEntries = value;
        }

        [Input("networkInterfaceIds")]
        private InputList<string>? _networkInterfaceIds;

        /// <summary>
        /// One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
        /// </summary>
        public InputList<string> NetworkInterfaceIds
        {
            get => _networkInterfaceIds ?? (_networkInterfaceIds = new InputList<string>());
            set => _networkInterfaceIds = value;
        }

        /// <summary>
        /// The ID of the AWS account that owns the VPC endpoint.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
        /// </summary>
        [Input("prefixListId")]
        public Input<string>? PrefixListId { get; set; }

        /// <summary>
        /// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
        /// Defaults to `false`.
        /// </summary>
        [Input("privateDnsEnabled")]
        public Input<bool>? PrivateDnsEnabled { get; set; }

        /// <summary>
        /// Whether or not the VPC Endpoint is being managed by its service - `true` or `false`.
        /// </summary>
        [Input("requesterManaged")]
        public Input<bool>? RequesterManaged { get; set; }

        [Input("routeTableIds")]
        private InputList<string>? _routeTableIds;

        /// <summary>
        /// One or more route table IDs. Applicable for endpoints of type `Gateway`.
        /// </summary>
        public InputList<string> RouteTableIds
        {
            get => _routeTableIds ?? (_routeTableIds = new InputList<string>());
            set => _routeTableIds = value;
        }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The service name, in the form `com.amazonaws.region.service` for AWS services.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// The state of the VPC endpoint.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `Interface`.
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
        /// The VPC endpoint type, `Gateway` or `Interface`. Defaults to `Gateway`.
        /// </summary>
        [Input("vpcEndpointType")]
        public Input<string>? VpcEndpointType { get; set; }

        /// <summary>
        /// The ID of the VPC in which the endpoint will be used.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public VpcEndpointState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class VpcEndpointDnsEntriesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DNS name.
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        /// <summary>
        /// The ID of the private hosted zone.
        /// </summary>
        [Input("hostedZoneId")]
        public Input<string>? HostedZoneId { get; set; }

        public VpcEndpointDnsEntriesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class VpcEndpointDnsEntries
    {
        /// <summary>
        /// The DNS name.
        /// </summary>
        public readonly string DnsName;
        /// <summary>
        /// The ID of the private hosted zone.
        /// </summary>
        public readonly string HostedZoneId;

        [OutputConstructor]
        private VpcEndpointDnsEntries(
            string dnsName,
            string hostedZoneId)
        {
            DnsName = dnsName;
            HostedZoneId = hostedZoneId;
        }
    }
    }
}
