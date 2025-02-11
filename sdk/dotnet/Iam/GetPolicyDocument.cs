// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iam
{
    public static partial class Invokes
    {
        /// <summary>
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/iam_policy_document.html.markdown.
        /// </summary>
        public static Task<GetPolicyDocumentResult> GetPolicyDocument(GetPolicyDocumentArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyDocumentResult>("aws:iam/getPolicyDocument:getPolicyDocument", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetPolicyDocumentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An IAM policy document to import and override the
        /// current policy document.  Statements with non-blank `sid`s in the override
        /// document will overwrite statements with the same `sid` in the current document.
        /// Statements without an `sid` cannot be overwritten.
        /// </summary>
        [Input("overrideJson")]
        public Input<string>? OverrideJson { get; set; }

        /// <summary>
        /// An ID for the policy document.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// An IAM policy document to import as a base for the
        /// current policy document.  Statements with non-blank `sid`s in the current
        /// policy document will overwrite statements with the same `sid` in the source
        /// json.  Statements without an `sid` cannot be overwritten.
        /// </summary>
        [Input("sourceJson")]
        public Input<string>? SourceJson { get; set; }

        [Input("statements")]
        private InputList<Inputs.GetPolicyDocumentStatementsArgs>? _statements;

        /// <summary>
        /// A nested configuration block (described below)
        /// configuring one *statement* to be included in the policy document.
        /// </summary>
        public InputList<Inputs.GetPolicyDocumentStatementsArgs> Statements
        {
            get => _statements ?? (_statements = new InputList<Inputs.GetPolicyDocumentStatementsArgs>());
            set => _statements = value;
        }

        /// <summary>
        /// IAM policy document version. Valid values: `2008-10-17`, `2012-10-17`. Defaults to `2012-10-17`. For more information, see the [AWS IAM User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_version.html).
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public GetPolicyDocumentArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetPolicyDocumentResult
    {
        /// <summary>
        /// The above arguments serialized as a standard JSON policy document.
        /// </summary>
        public readonly string Json;
        public readonly string? OverrideJson;
        public readonly string? PolicyId;
        public readonly string? SourceJson;
        public readonly ImmutableArray<Outputs.GetPolicyDocumentStatementsResult> Statements;
        public readonly string? Version;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetPolicyDocumentResult(
            string json,
            string? overrideJson,
            string? policyId,
            string? sourceJson,
            ImmutableArray<Outputs.GetPolicyDocumentStatementsResult> statements,
            string? version,
            string id)
        {
            Json = json;
            OverrideJson = overrideJson;
            PolicyId = policyId;
            SourceJson = sourceJson;
            Statements = statements;
            Version = version;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetPolicyDocumentStatementsArgs : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<string>? _actions;

        /// <summary>
        /// A list of actions that this statement either allows
        /// or denies. For example, ``["ec2:RunInstances", "s3:*"]``.
        /// </summary>
        public InputList<string> Actions
        {
            get => _actions ?? (_actions = new InputList<string>());
            set => _actions = value;
        }

        [Input("conditions")]
        private InputList<GetPolicyDocumentStatementsConditionsArgs>? _conditions;

        /// <summary>
        /// A nested configuration block (described below)
        /// that defines a further, possibly-service-specific condition that constrains
        /// whether this statement applies.
        /// </summary>
        public InputList<GetPolicyDocumentStatementsConditionsArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<GetPolicyDocumentStatementsConditionsArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// Either "Allow" or "Deny", to specify whether this
        /// statement allows or denies the given actions. The default is "Allow".
        /// </summary>
        [Input("effect")]
        public Input<string>? Effect { get; set; }

        [Input("notActions")]
        private InputList<string>? _notActions;

        /// <summary>
        /// A list of actions that this statement does *not*
        /// apply to. Used to apply a policy statement to all actions *except* those
        /// listed.
        /// </summary>
        public InputList<string> NotActions
        {
            get => _notActions ?? (_notActions = new InputList<string>());
            set => _notActions = value;
        }

        [Input("notPrincipals")]
        private InputList<GetPolicyDocumentStatementsNotPrincipalsArgs>? _notPrincipals;

        /// <summary>
        /// Like `principals` except gives resources that
        /// the statement does *not* apply to.
        /// </summary>
        public InputList<GetPolicyDocumentStatementsNotPrincipalsArgs> NotPrincipals
        {
            get => _notPrincipals ?? (_notPrincipals = new InputList<GetPolicyDocumentStatementsNotPrincipalsArgs>());
            set => _notPrincipals = value;
        }

        [Input("notResources")]
        private InputList<string>? _notResources;

        /// <summary>
        /// A list of resource ARNs that this statement
        /// does *not* apply to. Used to apply a policy statement to all resources
        /// *except* those listed.
        /// </summary>
        public InputList<string> NotResources
        {
            get => _notResources ?? (_notResources = new InputList<string>());
            set => _notResources = value;
        }

        [Input("principals")]
        private InputList<GetPolicyDocumentStatementsPrincipalsArgs>? _principals;

        /// <summary>
        /// A nested configuration block (described below)
        /// specifying a resource (or resource pattern) to which this statement applies.
        /// </summary>
        public InputList<GetPolicyDocumentStatementsPrincipalsArgs> Principals
        {
            get => _principals ?? (_principals = new InputList<GetPolicyDocumentStatementsPrincipalsArgs>());
            set => _principals = value;
        }

        [Input("resources")]
        private InputList<string>? _resources;

        /// <summary>
        /// A list of resource ARNs that this statement applies
        /// to. This is required by AWS if used for an IAM policy.
        /// </summary>
        public InputList<string> Resources
        {
            get => _resources ?? (_resources = new InputList<string>());
            set => _resources = value;
        }

        /// <summary>
        /// An ID for the policy statement.
        /// </summary>
        [Input("sid")]
        public Input<string>? Sid { get; set; }

        public GetPolicyDocumentStatementsArgs()
        {
        }
    }

    public sealed class GetPolicyDocumentStatementsConditionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the
        /// [IAM condition operator](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition_operators.html)
        /// to evaluate.
        /// </summary>
        [Input("test", required: true)]
        public Input<string> Test { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// The values to evaluate the condition against. If multiple
        /// values are provided, the condition matches if at least one of them applies.
        /// (That is, the tests are combined with the "OR" boolean operation.)
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        /// <summary>
        /// The name of a
        /// [Context Variable](http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html#AvailableKeys)
        /// to apply the condition to. Context variables may either be standard AWS
        /// variables starting with `aws:`, or service-specific variables prefixed with
        /// the service name.
        /// </summary>
        [Input("variable", required: true)]
        public Input<string> Variable { get; set; } = null!;

        public GetPolicyDocumentStatementsConditionsArgs()
        {
        }
    }

    public sealed class GetPolicyDocumentStatementsNotPrincipalsArgs : Pulumi.ResourceArgs
    {
        [Input("identifiers", required: true)]
        private InputList<string>? _identifiers;

        /// <summary>
        /// List of identifiers for principals. When `type`
        /// is "AWS", these are IAM user or role ARNs.  When `type` is "Service", these are AWS Service roles e.g. `lambda.amazonaws.com`.
        /// </summary>
        public InputList<string> Identifiers
        {
            get => _identifiers ?? (_identifiers = new InputList<string>());
            set => _identifiers = value;
        }

        /// <summary>
        /// The type of principal. For AWS ARNs this is "AWS".  For AWS services (e.g. Lambda), this is "Service".
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetPolicyDocumentStatementsNotPrincipalsArgs()
        {
        }
    }

    public sealed class GetPolicyDocumentStatementsPrincipalsArgs : Pulumi.ResourceArgs
    {
        [Input("identifiers", required: true)]
        private InputList<string>? _identifiers;

        /// <summary>
        /// List of identifiers for principals. When `type`
        /// is "AWS", these are IAM user or role ARNs.  When `type` is "Service", these are AWS Service roles e.g. `lambda.amazonaws.com`.
        /// </summary>
        public InputList<string> Identifiers
        {
            get => _identifiers ?? (_identifiers = new InputList<string>());
            set => _identifiers = value;
        }

        /// <summary>
        /// The type of principal. For AWS ARNs this is "AWS".  For AWS services (e.g. Lambda), this is "Service".
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetPolicyDocumentStatementsPrincipalsArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetPolicyDocumentStatementsConditionsResult
    {
        /// <summary>
        /// The name of the
        /// [IAM condition operator](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition_operators.html)
        /// to evaluate.
        /// </summary>
        public readonly string Test;
        /// <summary>
        /// The values to evaluate the condition against. If multiple
        /// values are provided, the condition matches if at least one of them applies.
        /// (That is, the tests are combined with the "OR" boolean operation.)
        /// </summary>
        public readonly ImmutableArray<string> Values;
        /// <summary>
        /// The name of a
        /// [Context Variable](http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html#AvailableKeys)
        /// to apply the condition to. Context variables may either be standard AWS
        /// variables starting with `aws:`, or service-specific variables prefixed with
        /// the service name.
        /// </summary>
        public readonly string Variable;

        [OutputConstructor]
        private GetPolicyDocumentStatementsConditionsResult(
            string test,
            ImmutableArray<string> values,
            string variable)
        {
            Test = test;
            Values = values;
            Variable = variable;
        }
    }

    [OutputType]
    public sealed class GetPolicyDocumentStatementsNotPrincipalsResult
    {
        /// <summary>
        /// List of identifiers for principals. When `type`
        /// is "AWS", these are IAM user or role ARNs.  When `type` is "Service", these are AWS Service roles e.g. `lambda.amazonaws.com`.
        /// </summary>
        public readonly ImmutableArray<string> Identifiers;
        /// <summary>
        /// The type of principal. For AWS ARNs this is "AWS".  For AWS services (e.g. Lambda), this is "Service".
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPolicyDocumentStatementsNotPrincipalsResult(
            ImmutableArray<string> identifiers,
            string type)
        {
            Identifiers = identifiers;
            Type = type;
        }
    }

    [OutputType]
    public sealed class GetPolicyDocumentStatementsPrincipalsResult
    {
        /// <summary>
        /// List of identifiers for principals. When `type`
        /// is "AWS", these are IAM user or role ARNs.  When `type` is "Service", these are AWS Service roles e.g. `lambda.amazonaws.com`.
        /// </summary>
        public readonly ImmutableArray<string> Identifiers;
        /// <summary>
        /// The type of principal. For AWS ARNs this is "AWS".  For AWS services (e.g. Lambda), this is "Service".
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPolicyDocumentStatementsPrincipalsResult(
            ImmutableArray<string> identifiers,
            string type)
        {
            Identifiers = identifiers;
            Type = type;
        }
    }

    [OutputType]
    public sealed class GetPolicyDocumentStatementsResult
    {
        /// <summary>
        /// A list of actions that this statement either allows
        /// or denies. For example, ``["ec2:RunInstances", "s3:*"]``.
        /// </summary>
        public readonly ImmutableArray<string> Actions;
        /// <summary>
        /// A nested configuration block (described below)
        /// that defines a further, possibly-service-specific condition that constrains
        /// whether this statement applies.
        /// </summary>
        public readonly ImmutableArray<GetPolicyDocumentStatementsConditionsResult> Conditions;
        /// <summary>
        /// Either "Allow" or "Deny", to specify whether this
        /// statement allows or denies the given actions. The default is "Allow".
        /// </summary>
        public readonly string? Effect;
        /// <summary>
        /// A list of actions that this statement does *not*
        /// apply to. Used to apply a policy statement to all actions *except* those
        /// listed.
        /// </summary>
        public readonly ImmutableArray<string> NotActions;
        /// <summary>
        /// Like `principals` except gives resources that
        /// the statement does *not* apply to.
        /// </summary>
        public readonly ImmutableArray<GetPolicyDocumentStatementsNotPrincipalsResult> NotPrincipals;
        /// <summary>
        /// A list of resource ARNs that this statement
        /// does *not* apply to. Used to apply a policy statement to all resources
        /// *except* those listed.
        /// </summary>
        public readonly ImmutableArray<string> NotResources;
        /// <summary>
        /// A nested configuration block (described below)
        /// specifying a resource (or resource pattern) to which this statement applies.
        /// </summary>
        public readonly ImmutableArray<GetPolicyDocumentStatementsPrincipalsResult> Principals;
        /// <summary>
        /// A list of resource ARNs that this statement applies
        /// to. This is required by AWS if used for an IAM policy.
        /// </summary>
        public readonly ImmutableArray<string> Resources;
        /// <summary>
        /// An ID for the policy statement.
        /// </summary>
        public readonly string? Sid;

        [OutputConstructor]
        private GetPolicyDocumentStatementsResult(
            ImmutableArray<string> actions,
            ImmutableArray<GetPolicyDocumentStatementsConditionsResult> conditions,
            string? effect,
            ImmutableArray<string> notActions,
            ImmutableArray<GetPolicyDocumentStatementsNotPrincipalsResult> notPrincipals,
            ImmutableArray<string> notResources,
            ImmutableArray<GetPolicyDocumentStatementsPrincipalsResult> principals,
            ImmutableArray<string> resources,
            string? sid)
        {
            Actions = actions;
            Conditions = conditions;
            Effect = effect;
            NotActions = notActions;
            NotPrincipals = notPrincipals;
            NotResources = notResources;
            Principals = principals;
            Resources = resources;
            Sid = sid;
        }
    }
    }
}
