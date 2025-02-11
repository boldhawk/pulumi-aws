// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Neptune
{
    /// <summary>
    /// Manages a Neptune Cluster Parameter Group
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/neptune_cluster_parameter_group.html.markdown.
    /// </summary>
    public partial class ClusterParameterGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the neptune cluster parameter group.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The description of the neptune cluster parameter group. Defaults to "Managed by Pulumi".
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The family of the neptune cluster parameter group.
        /// </summary>
        [Output("family")]
        public Output<string> Family { get; private set; } = null!;

        /// <summary>
        /// The name of the neptune parameter.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// A list of neptune parameters to apply.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableArray<Outputs.ClusterParameterGroupParameters>> Parameters { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterParameterGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterParameterGroup(string name, ClusterParameterGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:neptune/clusterParameterGroup:ClusterParameterGroup", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ClusterParameterGroup(string name, Input<string> id, ClusterParameterGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:neptune/clusterParameterGroup:ClusterParameterGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClusterParameterGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterParameterGroup Get(string name, Input<string> id, ClusterParameterGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ClusterParameterGroup(name, id, state, options);
        }
    }

    public sealed class ClusterParameterGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the neptune cluster parameter group. Defaults to "Managed by Pulumi".
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The family of the neptune cluster parameter group.
        /// </summary>
        [Input("family", required: true)]
        public Input<string> Family { get; set; } = null!;

        /// <summary>
        /// The name of the neptune parameter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("parameters")]
        private InputList<Inputs.ClusterParameterGroupParametersArgs>? _parameters;

        /// <summary>
        /// A list of neptune parameters to apply.
        /// </summary>
        public InputList<Inputs.ClusterParameterGroupParametersArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ClusterParameterGroupParametersArgs>());
            set => _parameters = value;
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

        public ClusterParameterGroupArgs()
        {
        }
    }

    public sealed class ClusterParameterGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the neptune cluster parameter group.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The description of the neptune cluster parameter group. Defaults to "Managed by Pulumi".
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The family of the neptune cluster parameter group.
        /// </summary>
        [Input("family")]
        public Input<string>? Family { get; set; }

        /// <summary>
        /// The name of the neptune parameter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("parameters")]
        private InputList<Inputs.ClusterParameterGroupParametersGetArgs>? _parameters;

        /// <summary>
        /// A list of neptune parameters to apply.
        /// </summary>
        public InputList<Inputs.ClusterParameterGroupParametersGetArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ClusterParameterGroupParametersGetArgs>());
            set => _parameters = value;
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

        public ClusterParameterGroupState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ClusterParameterGroupParametersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
        /// </summary>
        [Input("applyMethod")]
        public Input<string>? ApplyMethod { get; set; }

        /// <summary>
        /// The name of the neptune parameter.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The value of the neptune parameter.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ClusterParameterGroupParametersArgs()
        {
        }
    }

    public sealed class ClusterParameterGroupParametersGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
        /// </summary>
        [Input("applyMethod")]
        public Input<string>? ApplyMethod { get; set; }

        /// <summary>
        /// The name of the neptune parameter.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The value of the neptune parameter.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ClusterParameterGroupParametersGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ClusterParameterGroupParameters
    {
        /// <summary>
        /// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
        /// </summary>
        public readonly string? ApplyMethod;
        /// <summary>
        /// The name of the neptune parameter.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value of the neptune parameter.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ClusterParameterGroupParameters(
            string? applyMethod,
            string name,
            string value)
        {
            ApplyMethod = applyMethod;
            Name = name;
            Value = value;
        }
    }
    }
}
