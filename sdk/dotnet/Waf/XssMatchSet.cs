// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Waf
{
    /// <summary>
    /// Provides a WAF XSS Match Set Resource
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/waf_xss_match_set.html.markdown.
    /// </summary>
    public partial class XssMatchSet : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name or description of the SizeConstraintSet.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The parts of web requests that you want to inspect for cross-site scripting attacks.
        /// </summary>
        [Output("xssMatchTuples")]
        public Output<ImmutableArray<Outputs.XssMatchSetXssMatchTuples>> XssMatchTuples { get; private set; } = null!;


        /// <summary>
        /// Create a XssMatchSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public XssMatchSet(string name, XssMatchSetArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:waf/xssMatchSet:XssMatchSet", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private XssMatchSet(string name, Input<string> id, XssMatchSetState? state = null, CustomResourceOptions? options = null)
            : base("aws:waf/xssMatchSet:XssMatchSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing XssMatchSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static XssMatchSet Get(string name, Input<string> id, XssMatchSetState? state = null, CustomResourceOptions? options = null)
        {
            return new XssMatchSet(name, id, state, options);
        }
    }

    public sealed class XssMatchSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name or description of the SizeConstraintSet.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("xssMatchTuples")]
        private InputList<Inputs.XssMatchSetXssMatchTuplesArgs>? _xssMatchTuples;

        /// <summary>
        /// The parts of web requests that you want to inspect for cross-site scripting attacks.
        /// </summary>
        public InputList<Inputs.XssMatchSetXssMatchTuplesArgs> XssMatchTuples
        {
            get => _xssMatchTuples ?? (_xssMatchTuples = new InputList<Inputs.XssMatchSetXssMatchTuplesArgs>());
            set => _xssMatchTuples = value;
        }

        public XssMatchSetArgs()
        {
        }
    }

    public sealed class XssMatchSetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name or description of the SizeConstraintSet.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("xssMatchTuples")]
        private InputList<Inputs.XssMatchSetXssMatchTuplesGetArgs>? _xssMatchTuples;

        /// <summary>
        /// The parts of web requests that you want to inspect for cross-site scripting attacks.
        /// </summary>
        public InputList<Inputs.XssMatchSetXssMatchTuplesGetArgs> XssMatchTuples
        {
            get => _xssMatchTuples ?? (_xssMatchTuples = new InputList<Inputs.XssMatchSetXssMatchTuplesGetArgs>());
            set => _xssMatchTuples = value;
        }

        public XssMatchSetState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class XssMatchSetXssMatchTuplesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies where in a web request to look for cross-site scripting attacks.
        /// </summary>
        [Input("fieldToMatch", required: true)]
        public Input<XssMatchSetXssMatchTuplesFieldToMatchArgs> FieldToMatch { get; set; } = null!;

        /// <summary>
        /// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
        /// If you specify a transformation, AWS WAF performs the transformation on `target_string` before inspecting a request for a match.
        /// e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
        /// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_XssMatchTuple.html#WAF-Type-XssMatchTuple-TextTransformation)
        /// for all supported values.
        /// </summary>
        [Input("textTransformation", required: true)]
        public Input<string> TextTransformation { get; set; } = null!;

        public XssMatchSetXssMatchTuplesArgs()
        {
        }
    }

    public sealed class XssMatchSetXssMatchTuplesFieldToMatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When `type` is `HEADER`, enter the name of the header that you want to search, e.g. `User-Agent` or `Referer`.
        /// If `type` is any other value, omit this field.
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// The part of the web request that you want AWS WAF to search for a specified string.
        /// e.g. `HEADER`, `METHOD` or `BODY`.
        /// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html)
        /// for all supported values.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public XssMatchSetXssMatchTuplesFieldToMatchArgs()
        {
        }
    }

    public sealed class XssMatchSetXssMatchTuplesFieldToMatchGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When `type` is `HEADER`, enter the name of the header that you want to search, e.g. `User-Agent` or `Referer`.
        /// If `type` is any other value, omit this field.
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// The part of the web request that you want AWS WAF to search for a specified string.
        /// e.g. `HEADER`, `METHOD` or `BODY`.
        /// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html)
        /// for all supported values.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public XssMatchSetXssMatchTuplesFieldToMatchGetArgs()
        {
        }
    }

    public sealed class XssMatchSetXssMatchTuplesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies where in a web request to look for cross-site scripting attacks.
        /// </summary>
        [Input("fieldToMatch", required: true)]
        public Input<XssMatchSetXssMatchTuplesFieldToMatchGetArgs> FieldToMatch { get; set; } = null!;

        /// <summary>
        /// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
        /// If you specify a transformation, AWS WAF performs the transformation on `target_string` before inspecting a request for a match.
        /// e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
        /// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_XssMatchTuple.html#WAF-Type-XssMatchTuple-TextTransformation)
        /// for all supported values.
        /// </summary>
        [Input("textTransformation", required: true)]
        public Input<string> TextTransformation { get; set; } = null!;

        public XssMatchSetXssMatchTuplesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class XssMatchSetXssMatchTuples
    {
        /// <summary>
        /// Specifies where in a web request to look for cross-site scripting attacks.
        /// </summary>
        public readonly XssMatchSetXssMatchTuplesFieldToMatch FieldToMatch;
        /// <summary>
        /// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
        /// If you specify a transformation, AWS WAF performs the transformation on `target_string` before inspecting a request for a match.
        /// e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
        /// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_XssMatchTuple.html#WAF-Type-XssMatchTuple-TextTransformation)
        /// for all supported values.
        /// </summary>
        public readonly string TextTransformation;

        [OutputConstructor]
        private XssMatchSetXssMatchTuples(
            XssMatchSetXssMatchTuplesFieldToMatch fieldToMatch,
            string textTransformation)
        {
            FieldToMatch = fieldToMatch;
            TextTransformation = textTransformation;
        }
    }

    [OutputType]
    public sealed class XssMatchSetXssMatchTuplesFieldToMatch
    {
        /// <summary>
        /// When `type` is `HEADER`, enter the name of the header that you want to search, e.g. `User-Agent` or `Referer`.
        /// If `type` is any other value, omit this field.
        /// </summary>
        public readonly string? Data;
        /// <summary>
        /// The part of the web request that you want AWS WAF to search for a specified string.
        /// e.g. `HEADER`, `METHOD` or `BODY`.
        /// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html)
        /// for all supported values.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private XssMatchSetXssMatchTuplesFieldToMatch(
            string? data,
            string type)
        {
            Data = data;
            Type = type;
        }
    }
    }
}
