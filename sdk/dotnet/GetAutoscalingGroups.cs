// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws
{
    public static partial class Invokes
    {
        /// <summary>
        /// The Autoscaling Groups data source allows access to the list of AWS
        /// ASGs within a specific region. This will allow you to pass a list of AutoScaling Groups to other resources.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/autoscaling_groups.html.markdown.
        /// </summary>
        public static Task<GetAutoscalingGroupsResult> GetAutoscalingGroups(GetAutoscalingGroupsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAutoscalingGroupsResult>("aws:index/getAutoscalingGroups:getAutoscalingGroups", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetAutoscalingGroupsArgs : Pulumi.ResourceArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetAutoscalingGroupsFiltersArgs>? _filters;

        /// <summary>
        /// A filter used to scope the list e.g. by tags. See [related docs](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_Filter.html).
        /// </summary>
        public InputList<Inputs.GetAutoscalingGroupsFiltersArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetAutoscalingGroupsFiltersArgs>());
            set => _filters = value;
        }

        public GetAutoscalingGroupsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetAutoscalingGroupsResult
    {
        /// <summary>
        /// A list of the Autoscaling Groups Arns in the current region.
        /// </summary>
        public readonly ImmutableArray<string> Arns;
        public readonly ImmutableArray<Outputs.GetAutoscalingGroupsFiltersResult> Filters;
        /// <summary>
        /// A list of the Autoscaling Groups in the current region.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAutoscalingGroupsResult(
            ImmutableArray<string> arns,
            ImmutableArray<Outputs.GetAutoscalingGroupsFiltersResult> filters,
            ImmutableArray<string> names,
            string id)
        {
            Arns = arns;
            Filters = filters;
            Names = names;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetAutoscalingGroupsFiltersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the filter. The valid values are: `auto-scaling-group`, `key`, `value`, and `propagate-at-launch`.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// The value of the filter.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public GetAutoscalingGroupsFiltersArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetAutoscalingGroupsFiltersResult
    {
        /// <summary>
        /// The name of the filter. The valid values are: `auto-scaling-group`, `key`, `value`, and `propagate-at-launch`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value of the filter.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetAutoscalingGroupsFiltersResult(
            string name,
            ImmutableArray<string> values)
        {
            Name = name;
            Values = values;
        }
    }
    }
}
