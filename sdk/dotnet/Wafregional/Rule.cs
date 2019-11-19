// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafRegional
{
    /// <summary>
    /// Provides an WAF Regional Rule Resource for use with Application Load Balancer.
    /// 
    /// ## Nested Fields
    /// 
    /// ### `predicate`
    /// 
    /// See the [WAF Documentation](https://docs.aws.amazon.com/waf/latest/APIReference/API_Predicate.html) for more information.
    /// 
    /// #### Arguments
    /// 
    /// * `type` - (Required) The type of predicate in a rule. Valid values: `ByteMatch`, `GeoMatch`, `IPMatch`, `RegexMatch`, `SizeConstraint`, `SqlInjectionMatch`, or `XssMatch`
    /// * `data_id` - (Required) The unique identifier of a predicate, such as the ID of a `ByteMatchSet` or `IPSet`.
    /// * `negated` - (Required) Whether to use the settings or the negated settings that you specified in the objects.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/wafregional_rule.html.markdown.
    /// </summary>
    public partial class Rule : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the WAF Regional Rule.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name or description for the Amazon CloudWatch metric of this rule.
        /// </summary>
        [Output("metricName")]
        public Output<string> MetricName { get; private set; } = null!;

        /// <summary>
        /// The name or description of the rule.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The objects to include in a rule (documented below).
        /// </summary>
        [Output("predicates")]
        public Output<ImmutableArray<Outputs.RulePredicates>> Predicates { get; private set; } = null!;

        /// <summary>
        /// Key-value mapping of resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Rule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Rule(string name, RuleArgs args, CustomResourceOptions? options = null)
            : base("aws:wafregional/rule:Rule", name, args, MakeResourceOptions(options, ""))
        {
        }

        private Rule(string name, Input<string> id, RuleState? state = null, CustomResourceOptions? options = null)
            : base("aws:wafregional/rule:Rule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Rule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Rule Get(string name, Input<string> id, RuleState? state = null, CustomResourceOptions? options = null)
        {
            return new Rule(name, id, state, options);
        }
    }

    public sealed class RuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name or description for the Amazon CloudWatch metric of this rule.
        /// </summary>
        [Input("metricName", required: true)]
        public Input<string> MetricName { get; set; } = null!;

        /// <summary>
        /// The name or description of the rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("predicates")]
        private InputList<Inputs.RulePredicatesArgs>? _predicates;

        /// <summary>
        /// The objects to include in a rule (documented below).
        /// </summary>
        public InputList<Inputs.RulePredicatesArgs> Predicates
        {
            get => _predicates ?? (_predicates = new InputList<Inputs.RulePredicatesArgs>());
            set => _predicates = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public RuleArgs()
        {
        }
    }

    public sealed class RuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the WAF Regional Rule.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name or description for the Amazon CloudWatch metric of this rule.
        /// </summary>
        [Input("metricName")]
        public Input<string>? MetricName { get; set; }

        /// <summary>
        /// The name or description of the rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("predicates")]
        private InputList<Inputs.RulePredicatesGetArgs>? _predicates;

        /// <summary>
        /// The objects to include in a rule (documented below).
        /// </summary>
        public InputList<Inputs.RulePredicatesGetArgs> Predicates
        {
            get => _predicates ?? (_predicates = new InputList<Inputs.RulePredicatesGetArgs>());
            set => _predicates = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public RuleState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class RulePredicatesArgs : Pulumi.ResourceArgs
    {
        [Input("dataId", required: true)]
        public Input<string> DataId { get; set; } = null!;

        [Input("negated", required: true)]
        public Input<bool> Negated { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public RulePredicatesArgs()
        {
        }
    }

    public sealed class RulePredicatesGetArgs : Pulumi.ResourceArgs
    {
        [Input("dataId", required: true)]
        public Input<string> DataId { get; set; } = null!;

        [Input("negated", required: true)]
        public Input<bool> Negated { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public RulePredicatesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class RulePredicates
    {
        public readonly string DataId;
        public readonly bool Negated;
        public readonly string Type;

        [OutputConstructor]
        private RulePredicates(
            string dataId,
            bool negated,
            string type)
        {
            DataId = dataId;
            Negated = negated;
            Type = type;
        }
    }
    }
}
