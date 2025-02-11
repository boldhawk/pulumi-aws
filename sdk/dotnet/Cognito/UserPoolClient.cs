// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito
{
    /// <summary>
    /// Provides a Cognito User Pool Client resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cognito_user_pool_client.html.markdown.
    /// </summary>
    public partial class UserPoolClient : Pulumi.CustomResource
    {
        /// <summary>
        /// List of allowed OAuth flows (code, implicit, client_credentials).
        /// </summary>
        [Output("allowedOauthFlows")]
        public Output<ImmutableArray<string>> AllowedOauthFlows { get; private set; } = null!;

        /// <summary>
        /// Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
        /// </summary>
        [Output("allowedOauthFlowsUserPoolClient")]
        public Output<bool?> AllowedOauthFlowsUserPoolClient { get; private set; } = null!;

        /// <summary>
        /// List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
        /// </summary>
        [Output("allowedOauthScopes")]
        public Output<ImmutableArray<string>> AllowedOauthScopes { get; private set; } = null!;

        /// <summary>
        /// List of allowed callback URLs for the identity providers.
        /// </summary>
        [Output("callbackUrls")]
        public Output<ImmutableArray<string>> CallbackUrls { get; private set; } = null!;

        /// <summary>
        /// The client secret of the user pool client.
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// The default redirect URI. Must be in the list of callback URLs.
        /// </summary>
        [Output("defaultRedirectUri")]
        public Output<string?> DefaultRedirectUri { get; private set; } = null!;

        /// <summary>
        /// List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY, USER_PASSWORD_AUTH).
        /// </summary>
        [Output("explicitAuthFlows")]
        public Output<ImmutableArray<string>> ExplicitAuthFlows { get; private set; } = null!;

        /// <summary>
        /// Should an application secret be generated.
        /// </summary>
        [Output("generateSecret")]
        public Output<bool?> GenerateSecret { get; private set; } = null!;

        /// <summary>
        /// List of allowed logout URLs for the identity providers.
        /// </summary>
        [Output("logoutUrls")]
        public Output<ImmutableArray<string>> LogoutUrls { get; private set; } = null!;

        /// <summary>
        /// The name of the application client.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of user pool attributes the application client can read from.
        /// </summary>
        [Output("readAttributes")]
        public Output<ImmutableArray<string>> ReadAttributes { get; private set; } = null!;

        /// <summary>
        /// The time limit in days refresh tokens are valid for.
        /// </summary>
        [Output("refreshTokenValidity")]
        public Output<int?> RefreshTokenValidity { get; private set; } = null!;

        /// <summary>
        /// List of provider names for the identity providers that are supported on this client.
        /// </summary>
        [Output("supportedIdentityProviders")]
        public Output<ImmutableArray<string>> SupportedIdentityProviders { get; private set; } = null!;

        /// <summary>
        /// The user pool the client belongs to.
        /// </summary>
        [Output("userPoolId")]
        public Output<string> UserPoolId { get; private set; } = null!;

        /// <summary>
        /// List of user pool attributes the application client can write to.
        /// </summary>
        [Output("writeAttributes")]
        public Output<ImmutableArray<string>> WriteAttributes { get; private set; } = null!;


        /// <summary>
        /// Create a UserPoolClient resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserPoolClient(string name, UserPoolClientArgs args, CustomResourceOptions? options = null)
            : base("aws:cognito/userPoolClient:UserPoolClient", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private UserPoolClient(string name, Input<string> id, UserPoolClientState? state = null, CustomResourceOptions? options = null)
            : base("aws:cognito/userPoolClient:UserPoolClient", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserPoolClient resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserPoolClient Get(string name, Input<string> id, UserPoolClientState? state = null, CustomResourceOptions? options = null)
        {
            return new UserPoolClient(name, id, state, options);
        }
    }

    public sealed class UserPoolClientArgs : Pulumi.ResourceArgs
    {
        [Input("allowedOauthFlows")]
        private InputList<string>? _allowedOauthFlows;

        /// <summary>
        /// List of allowed OAuth flows (code, implicit, client_credentials).
        /// </summary>
        public InputList<string> AllowedOauthFlows
        {
            get => _allowedOauthFlows ?? (_allowedOauthFlows = new InputList<string>());
            set => _allowedOauthFlows = value;
        }

        /// <summary>
        /// Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
        /// </summary>
        [Input("allowedOauthFlowsUserPoolClient")]
        public Input<bool>? AllowedOauthFlowsUserPoolClient { get; set; }

        [Input("allowedOauthScopes")]
        private InputList<string>? _allowedOauthScopes;

        /// <summary>
        /// List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
        /// </summary>
        public InputList<string> AllowedOauthScopes
        {
            get => _allowedOauthScopes ?? (_allowedOauthScopes = new InputList<string>());
            set => _allowedOauthScopes = value;
        }

        [Input("callbackUrls")]
        private InputList<string>? _callbackUrls;

        /// <summary>
        /// List of allowed callback URLs for the identity providers.
        /// </summary>
        public InputList<string> CallbackUrls
        {
            get => _callbackUrls ?? (_callbackUrls = new InputList<string>());
            set => _callbackUrls = value;
        }

        /// <summary>
        /// The default redirect URI. Must be in the list of callback URLs.
        /// </summary>
        [Input("defaultRedirectUri")]
        public Input<string>? DefaultRedirectUri { get; set; }

        [Input("explicitAuthFlows")]
        private InputList<string>? _explicitAuthFlows;

        /// <summary>
        /// List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY, USER_PASSWORD_AUTH).
        /// </summary>
        public InputList<string> ExplicitAuthFlows
        {
            get => _explicitAuthFlows ?? (_explicitAuthFlows = new InputList<string>());
            set => _explicitAuthFlows = value;
        }

        /// <summary>
        /// Should an application secret be generated.
        /// </summary>
        [Input("generateSecret")]
        public Input<bool>? GenerateSecret { get; set; }

        [Input("logoutUrls")]
        private InputList<string>? _logoutUrls;

        /// <summary>
        /// List of allowed logout URLs for the identity providers.
        /// </summary>
        public InputList<string> LogoutUrls
        {
            get => _logoutUrls ?? (_logoutUrls = new InputList<string>());
            set => _logoutUrls = value;
        }

        /// <summary>
        /// The name of the application client.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("readAttributes")]
        private InputList<string>? _readAttributes;

        /// <summary>
        /// List of user pool attributes the application client can read from.
        /// </summary>
        public InputList<string> ReadAttributes
        {
            get => _readAttributes ?? (_readAttributes = new InputList<string>());
            set => _readAttributes = value;
        }

        /// <summary>
        /// The time limit in days refresh tokens are valid for.
        /// </summary>
        [Input("refreshTokenValidity")]
        public Input<int>? RefreshTokenValidity { get; set; }

        [Input("supportedIdentityProviders")]
        private InputList<string>? _supportedIdentityProviders;

        /// <summary>
        /// List of provider names for the identity providers that are supported on this client.
        /// </summary>
        public InputList<string> SupportedIdentityProviders
        {
            get => _supportedIdentityProviders ?? (_supportedIdentityProviders = new InputList<string>());
            set => _supportedIdentityProviders = value;
        }

        /// <summary>
        /// The user pool the client belongs to.
        /// </summary>
        [Input("userPoolId", required: true)]
        public Input<string> UserPoolId { get; set; } = null!;

        [Input("writeAttributes")]
        private InputList<string>? _writeAttributes;

        /// <summary>
        /// List of user pool attributes the application client can write to.
        /// </summary>
        public InputList<string> WriteAttributes
        {
            get => _writeAttributes ?? (_writeAttributes = new InputList<string>());
            set => _writeAttributes = value;
        }

        public UserPoolClientArgs()
        {
        }
    }

    public sealed class UserPoolClientState : Pulumi.ResourceArgs
    {
        [Input("allowedOauthFlows")]
        private InputList<string>? _allowedOauthFlows;

        /// <summary>
        /// List of allowed OAuth flows (code, implicit, client_credentials).
        /// </summary>
        public InputList<string> AllowedOauthFlows
        {
            get => _allowedOauthFlows ?? (_allowedOauthFlows = new InputList<string>());
            set => _allowedOauthFlows = value;
        }

        /// <summary>
        /// Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
        /// </summary>
        [Input("allowedOauthFlowsUserPoolClient")]
        public Input<bool>? AllowedOauthFlowsUserPoolClient { get; set; }

        [Input("allowedOauthScopes")]
        private InputList<string>? _allowedOauthScopes;

        /// <summary>
        /// List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
        /// </summary>
        public InputList<string> AllowedOauthScopes
        {
            get => _allowedOauthScopes ?? (_allowedOauthScopes = new InputList<string>());
            set => _allowedOauthScopes = value;
        }

        [Input("callbackUrls")]
        private InputList<string>? _callbackUrls;

        /// <summary>
        /// List of allowed callback URLs for the identity providers.
        /// </summary>
        public InputList<string> CallbackUrls
        {
            get => _callbackUrls ?? (_callbackUrls = new InputList<string>());
            set => _callbackUrls = value;
        }

        /// <summary>
        /// The client secret of the user pool client.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// The default redirect URI. Must be in the list of callback URLs.
        /// </summary>
        [Input("defaultRedirectUri")]
        public Input<string>? DefaultRedirectUri { get; set; }

        [Input("explicitAuthFlows")]
        private InputList<string>? _explicitAuthFlows;

        /// <summary>
        /// List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY, USER_PASSWORD_AUTH).
        /// </summary>
        public InputList<string> ExplicitAuthFlows
        {
            get => _explicitAuthFlows ?? (_explicitAuthFlows = new InputList<string>());
            set => _explicitAuthFlows = value;
        }

        /// <summary>
        /// Should an application secret be generated.
        /// </summary>
        [Input("generateSecret")]
        public Input<bool>? GenerateSecret { get; set; }

        [Input("logoutUrls")]
        private InputList<string>? _logoutUrls;

        /// <summary>
        /// List of allowed logout URLs for the identity providers.
        /// </summary>
        public InputList<string> LogoutUrls
        {
            get => _logoutUrls ?? (_logoutUrls = new InputList<string>());
            set => _logoutUrls = value;
        }

        /// <summary>
        /// The name of the application client.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("readAttributes")]
        private InputList<string>? _readAttributes;

        /// <summary>
        /// List of user pool attributes the application client can read from.
        /// </summary>
        public InputList<string> ReadAttributes
        {
            get => _readAttributes ?? (_readAttributes = new InputList<string>());
            set => _readAttributes = value;
        }

        /// <summary>
        /// The time limit in days refresh tokens are valid for.
        /// </summary>
        [Input("refreshTokenValidity")]
        public Input<int>? RefreshTokenValidity { get; set; }

        [Input("supportedIdentityProviders")]
        private InputList<string>? _supportedIdentityProviders;

        /// <summary>
        /// List of provider names for the identity providers that are supported on this client.
        /// </summary>
        public InputList<string> SupportedIdentityProviders
        {
            get => _supportedIdentityProviders ?? (_supportedIdentityProviders = new InputList<string>());
            set => _supportedIdentityProviders = value;
        }

        /// <summary>
        /// The user pool the client belongs to.
        /// </summary>
        [Input("userPoolId")]
        public Input<string>? UserPoolId { get; set; }

        [Input("writeAttributes")]
        private InputList<string>? _writeAttributes;

        /// <summary>
        /// List of user pool attributes the application client can write to.
        /// </summary>
        public InputList<string> WriteAttributes
        {
            get => _writeAttributes ?? (_writeAttributes = new InputList<string>());
            set => _writeAttributes = value;
        }

        public UserPoolClientState()
        {
        }
    }
}
