// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GameLift
{
    /// <summary>
    /// Provides an Gamelift Build resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/gamelift_build.html.markdown.
    /// </summary>
    public partial class Build : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the build
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
        /// </summary>
        [Output("operatingSystem")]
        public Output<string> OperatingSystem { get; private set; } = null!;

        /// <summary>
        /// Information indicating where your game build files are stored. See below.
        /// </summary>
        [Output("storageLocation")]
        public Output<Outputs.BuildStorageLocation> StorageLocation { get; private set; } = null!;

        /// <summary>
        /// Version that is associated with this build.
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Build resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Build(string name, BuildArgs args, CustomResourceOptions? options = null)
            : base("aws:gamelift/build:Build", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Build(string name, Input<string> id, BuildState? state = null, CustomResourceOptions? options = null)
            : base("aws:gamelift/build:Build", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Build resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Build Get(string name, Input<string> id, BuildState? state = null, CustomResourceOptions? options = null)
        {
            return new Build(name, id, state, options);
        }
    }

    public sealed class BuildArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the build
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
        /// </summary>
        [Input("operatingSystem", required: true)]
        public Input<string> OperatingSystem { get; set; } = null!;

        /// <summary>
        /// Information indicating where your game build files are stored. See below.
        /// </summary>
        [Input("storageLocation", required: true)]
        public Input<Inputs.BuildStorageLocationArgs> StorageLocation { get; set; } = null!;

        /// <summary>
        /// Version that is associated with this build.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public BuildArgs()
        {
        }
    }

    public sealed class BuildState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the build
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
        /// </summary>
        [Input("operatingSystem")]
        public Input<string>? OperatingSystem { get; set; }

        /// <summary>
        /// Information indicating where your game build files are stored. See below.
        /// </summary>
        [Input("storageLocation")]
        public Input<Inputs.BuildStorageLocationGetArgs>? StorageLocation { get; set; }

        /// <summary>
        /// Version that is associated with this build.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public BuildState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class BuildStorageLocationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of your S3 bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Name of the zip file containing your build files.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// ARN of the access role that allows Amazon GameLift to access your S3 bucket.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        public BuildStorageLocationArgs()
        {
        }
    }

    public sealed class BuildStorageLocationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of your S3 bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Name of the zip file containing your build files.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// ARN of the access role that allows Amazon GameLift to access your S3 bucket.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        public BuildStorageLocationGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class BuildStorageLocation
    {
        /// <summary>
        /// Name of your S3 bucket.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// Name of the zip file containing your build files.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// ARN of the access role that allows Amazon GameLift to access your S3 bucket.
        /// </summary>
        public readonly string RoleArn;

        [OutputConstructor]
        private BuildStorageLocation(
            string bucket,
            string key,
            string roleArn)
        {
            Bucket = bucket;
            Key = key;
            RoleArn = roleArn;
        }
    }
    }
}
