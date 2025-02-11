// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Fsx
{
    /// <summary>
    /// Manages a FSx Lustre File System. See the [FSx Lustre Guide](https://docs.aws.amazon.com/fsx/latest/LustreGuide/what-is.html) for more information.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/fsx_lustre_file_system.html.markdown.
    /// </summary>
    public partial class LustreFileSystem : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name of the file system.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// DNS name for the file system, e.g. `fs-12345678.fsx.us-west-2.amazonaws.com`
        /// </summary>
        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        /// <summary>
        /// S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `import_path` argument and the path must use the same Amazon S3 bucket as specified in `import_path`. Set equal to `import_path` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
        /// </summary>
        [Output("exportPath")]
        public Output<string> ExportPath { get; private set; } = null!;

        /// <summary>
        /// S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
        /// </summary>
        [Output("importPath")]
        public Output<string?> ImportPath { get; private set; } = null!;

        /// <summary>
        /// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `import_path` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
        /// </summary>
        [Output("importedFileChunkSize")]
        public Output<int> ImportedFileChunkSize { get; private set; } = null!;

        /// <summary>
        /// Set of Elastic Network Interface identifiers from which the file system is accessible.
        /// </summary>
        [Output("networkInterfaceIds")]
        public Output<ImmutableArray<string>> NetworkInterfaceIds { get; private set; } = null!;

        /// <summary>
        /// AWS account identifier that created the file system.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// The storage capacity (GiB) of the file system. Minimum of `3600`. Storage capacity is provisioned in increments of 3,600 GiB.
        /// </summary>
        [Output("storageCapacity")]
        public Output<int> StorageCapacity { get; private set; } = null!;

        /// <summary>
        /// A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
        /// </summary>
        [Output("subnetIds")]
        public Output<string> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the file system.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Identifier of the Virtual Private Cloud for the file system.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
        /// </summary>
        [Output("weeklyMaintenanceStartTime")]
        public Output<string> WeeklyMaintenanceStartTime { get; private set; } = null!;


        /// <summary>
        /// Create a LustreFileSystem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LustreFileSystem(string name, LustreFileSystemArgs args, CustomResourceOptions? options = null)
            : base("aws:fsx/lustreFileSystem:LustreFileSystem", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private LustreFileSystem(string name, Input<string> id, LustreFileSystemState? state = null, CustomResourceOptions? options = null)
            : base("aws:fsx/lustreFileSystem:LustreFileSystem", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LustreFileSystem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LustreFileSystem Get(string name, Input<string> id, LustreFileSystemState? state = null, CustomResourceOptions? options = null)
        {
            return new LustreFileSystem(name, id, state, options);
        }
    }

    public sealed class LustreFileSystemArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `import_path` argument and the path must use the same Amazon S3 bucket as specified in `import_path`. Set equal to `import_path` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
        /// </summary>
        [Input("exportPath")]
        public Input<string>? ExportPath { get; set; }

        /// <summary>
        /// S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
        /// </summary>
        [Input("importPath")]
        public Input<string>? ImportPath { get; set; }

        /// <summary>
        /// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `import_path` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
        /// </summary>
        [Input("importedFileChunkSize")]
        public Input<int>? ImportedFileChunkSize { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The storage capacity (GiB) of the file system. Minimum of `3600`. Storage capacity is provisioned in increments of 3,600 GiB.
        /// </summary>
        [Input("storageCapacity", required: true)]
        public Input<int> StorageCapacity { get; set; } = null!;

        /// <summary>
        /// A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
        /// </summary>
        [Input("subnetIds", required: true)]
        public Input<string> SubnetIds { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the file system.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
        /// </summary>
        [Input("weeklyMaintenanceStartTime")]
        public Input<string>? WeeklyMaintenanceStartTime { get; set; }

        public LustreFileSystemArgs()
        {
        }
    }

    public sealed class LustreFileSystemState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name of the file system.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// DNS name for the file system, e.g. `fs-12345678.fsx.us-west-2.amazonaws.com`
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        /// <summary>
        /// S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `import_path` argument and the path must use the same Amazon S3 bucket as specified in `import_path`. Set equal to `import_path` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
        /// </summary>
        [Input("exportPath")]
        public Input<string>? ExportPath { get; set; }

        /// <summary>
        /// S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
        /// </summary>
        [Input("importPath")]
        public Input<string>? ImportPath { get; set; }

        /// <summary>
        /// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `import_path` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
        /// </summary>
        [Input("importedFileChunkSize")]
        public Input<int>? ImportedFileChunkSize { get; set; }

        [Input("networkInterfaceIds")]
        private InputList<string>? _networkInterfaceIds;

        /// <summary>
        /// Set of Elastic Network Interface identifiers from which the file system is accessible.
        /// </summary>
        public InputList<string> NetworkInterfaceIds
        {
            get => _networkInterfaceIds ?? (_networkInterfaceIds = new InputList<string>());
            set => _networkInterfaceIds = value;
        }

        /// <summary>
        /// AWS account identifier that created the file system.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The storage capacity (GiB) of the file system. Minimum of `3600`. Storage capacity is provisioned in increments of 3,600 GiB.
        /// </summary>
        [Input("storageCapacity")]
        public Input<int>? StorageCapacity { get; set; }

        /// <summary>
        /// A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
        /// </summary>
        [Input("subnetIds")]
        public Input<string>? SubnetIds { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the file system.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Identifier of the Virtual Private Cloud for the file system.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
        /// </summary>
        [Input("weeklyMaintenanceStartTime")]
        public Input<string>? WeeklyMaintenanceStartTime { get; set; }

        public LustreFileSystemState()
        {
        }
    }
}
