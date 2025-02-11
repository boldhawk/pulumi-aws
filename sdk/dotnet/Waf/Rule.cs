// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Waf
{
    /// <summary>
    /// Provides a WAF Rule Resource
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/waf_rule.html.markdown.
    /// </summary>
    public partial class Rule : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the WAF rule.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
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
            : base("aws:waf/rule:Rule", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Rule(string name, Input<string> id, RuleState? state = null, CustomResourceOptions? options = null)
            : base("aws:waf/rule:Rule", name, state, MakeResourceOptions(options, id))
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
        /// The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
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
        /// The ARN of the WAF rule.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
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
        /// <summary>
        /// A unique identifier for a predicate in the rule, such as Byte Match Set ID or IPSet ID.
        /// </summary>
        [Input("dataId", required: true)]
        public Input<string> DataId { get; set; } = null!;

        /// <summary>
        /// Set this to `false` if you want to allow, block, or count requests
        /// based on the settings in the specified [waf_byte_match_set](https://www.terraform.io/docs/providers/aws/r/waf_byte_match_set.html), [waf_ipset](https://www.terraform.io/docs/providers/aws/r/waf_ipset.html), [aws.waf.SizeConstraintSet](https://www.terraform.io/docs/providers/aws/r/waf_size_constraint_set.html), [aws.waf.SqlInjectionMatchSet](https://www.terraform.io/docs/providers/aws/r/waf_sql_injection_match_set.html) or [aws.waf.XssMatchSet](https://www.terraform.io/docs/providers/aws/r/waf_xss_match_set.html).
        /// For example, if an IPSet includes the IP address `192.0.2.44`, AWS WAF will allow or block requests based on that IP address.
        /// If set to `true`, AWS WAF will allow, block, or count requests based on all IP addresses _except_ `192.0.2.44`.
        /// </summary>
        [Input("negated", required: true)]
        public Input<bool> Negated { get; set; } = null!;

        /// <summary>
        /// The type of predicate in a rule. Valid values: `ByteMatch`, `GeoMatch`, `IPMatch`, `RegexMatch`, `SizeConstraint`, `SqlInjectionMatch`, or `XssMatch`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public RulePredicatesArgs()
        {
        }
    }

    public sealed class RulePredicatesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique identifier for a predicate in the rule, such as Byte Match Set ID or IPSet ID.
        /// </summary>
        [Input("dataId", required: true)]
        public Input<string> DataId { get; set; } = null!;

        /// <summary>
        /// Set this to `false` if you want to allow, block, or count requests
        /// based on the settings in the specified [waf_byte_match_set](https://www.terraform.io/docs/providers/aws/r/waf_byte_match_set.html), [waf_ipset](https://www.terraform.io/docs/providers/aws/r/waf_ipset.html), [aws.waf.SizeConstraintSet](https://www.terraform.io/docs/providers/aws/r/waf_size_constraint_set.html), [aws.waf.SqlInjectionMatchSet](https://www.terraform.io/docs/providers/aws/r/waf_sql_injection_match_set.html) or [aws.waf.XssMatchSet](https://www.terraform.io/docs/providers/aws/r/waf_xss_match_set.html).
        /// For example, if an IPSet includes the IP address `192.0.2.44`, AWS WAF will allow or block requests based on that IP address.
        /// If set to `true`, AWS WAF will allow, block, or count requests based on all IP addresses _except_ `192.0.2.44`.
        /// </summary>
        [Input("negated", required: true)]
        public Input<bool> Negated { get; set; } = null!;

        /// <summary>
        /// The type of predicate in a rule. Valid values: `ByteMatch`, `GeoMatch`, `IPMatch`, `RegexMatch`, `SizeConstraint`, `SqlInjectionMatch`, or `XssMatch`.
        /// </summary>
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
        /// <summary>
        /// A unique identifier for a predicate in the rule, such as Byte Match Set ID or IPSet ID.
        /// </summary>
        public readonly string DataId;
        /// <summary>
        /// Set this to `false` if you want to allow, block, or count requests
        /// based on the settings in the specified [waf_byte_match_set](https://www.terraform.io/docs/providers/aws/r/waf_byte_match_set.html), [waf_ipset](https://www.terraform.io/docs/providers/aws/r/waf_ipset.html), [aws.waf.SizeConstraintSet](https://www.terraform.io/docs/providers/aws/r/waf_size_constraint_set.html), [aws.waf.SqlInjectionMatchSet](https://www.terraform.io/docs/providers/aws/r/waf_sql_injection_match_set.html) or [aws.waf.XssMatchSet](https://www.terraform.io/docs/providers/aws/r/waf_xss_match_set.html).
        /// For example, if an IPSet includes the IP address `192.0.2.44`, AWS WAF will allow or block requests based on that IP address.
        /// If set to `true`, AWS WAF will allow, block, or count requests based on all IP addresses _except_ `192.0.2.44`.
        /// </summary>
        public readonly bool Negated;
        /// <summary>
        /// The type of predicate in a rule. Valid values: `ByteMatch`, `GeoMatch`, `IPMatch`, `RegexMatch`, `SizeConstraint`, `SqlInjectionMatch`, or `XssMatch`.
        /// </summary>
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
