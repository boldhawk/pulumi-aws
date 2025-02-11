// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Dms
{
    /// <summary>
    /// Provides a DMS (Data Migration Service) replication subnet group resource. DMS replication subnet groups can be created, updated, deleted, and imported.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dms_replication_subnet_group.html.markdown.
    /// </summary>
    public partial class ReplicationSubnetGroup : Pulumi.CustomResource
    {
        [Output("replicationSubnetGroupArn")]
        public Output<string> ReplicationSubnetGroupArn { get; private set; } = null!;

        /// <summary>
        /// The description for the subnet group.
        /// </summary>
        [Output("replicationSubnetGroupDescription")]
        public Output<string> ReplicationSubnetGroupDescription { get; private set; } = null!;

        /// <summary>
        /// The name for the replication subnet group. This value is stored as a lowercase string.
        /// </summary>
        [Output("replicationSubnetGroupId")]
        public Output<string> ReplicationSubnetGroupId { get; private set; } = null!;

        /// <summary>
        /// A list of the EC2 subnet IDs for the subnet group.
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC the subnet group is in.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a ReplicationSubnetGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplicationSubnetGroup(string name, ReplicationSubnetGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:dms/replicationSubnetGroup:ReplicationSubnetGroup", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ReplicationSubnetGroup(string name, Input<string> id, ReplicationSubnetGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:dms/replicationSubnetGroup:ReplicationSubnetGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReplicationSubnetGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReplicationSubnetGroup Get(string name, Input<string> id, ReplicationSubnetGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ReplicationSubnetGroup(name, id, state, options);
        }
    }

    public sealed class ReplicationSubnetGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description for the subnet group.
        /// </summary>
        [Input("replicationSubnetGroupDescription", required: true)]
        public Input<string> ReplicationSubnetGroupDescription { get; set; } = null!;

        /// <summary>
        /// The name for the replication subnet group. This value is stored as a lowercase string.
        /// </summary>
        [Input("replicationSubnetGroupId", required: true)]
        public Input<string> ReplicationSubnetGroupId { get; set; } = null!;

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// A list of the EC2 subnet IDs for the subnet group.
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

        public ReplicationSubnetGroupArgs()
        {
        }
    }

    public sealed class ReplicationSubnetGroupState : Pulumi.ResourceArgs
    {
        [Input("replicationSubnetGroupArn")]
        public Input<string>? ReplicationSubnetGroupArn { get; set; }

        /// <summary>
        /// The description for the subnet group.
        /// </summary>
        [Input("replicationSubnetGroupDescription")]
        public Input<string>? ReplicationSubnetGroupDescription { get; set; }

        /// <summary>
        /// The name for the replication subnet group. This value is stored as a lowercase string.
        /// </summary>
        [Input("replicationSubnetGroupId")]
        public Input<string>? ReplicationSubnetGroupId { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// A list of the EC2 subnet IDs for the subnet group.
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
        /// The ID of the VPC the subnet group is in.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public ReplicationSubnetGroupState()
        {
        }
    }
}
