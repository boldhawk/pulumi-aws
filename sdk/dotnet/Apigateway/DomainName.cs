// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway
{
    /// <summary>
    /// Registers a custom domain name for use with AWS API Gateway. Additional information about this functionality
    /// can be found in the [API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html).
    /// 
    /// This resource just establishes ownership of and the TLS settings for
    /// a particular domain name. An API can be attached to a particular path
    /// under the registered domain name using
    /// the `aws.apigateway.BasePathMapping` resource.
    /// 
    /// API Gateway domains can be defined as either 'edge-optimized' or 'regional'.  In an edge-optimized configuration,
    /// API Gateway internally creates and manages a CloudFront distribution to route requests on the given hostname. In
    /// addition to this resource it's necessary to create a DNS record corresponding to the given domain name which is an alias
    /// (either Route53 alias or traditional CNAME) to the Cloudfront domain name exported in the `cloudfront_domain_name`
    /// attribute.
    /// 
    /// In a regional configuration, API Gateway does not create a CloudFront distribution to route requests to the API, though
    /// a distribution can be created if needed. In either case, it is necessary to create a DNS record corresponding to the
    /// given domain name which is an alias (either Route53 alias or traditional CNAME) to the regional domain name exported in
    /// the `regional_domain_name` attribute.
    /// 
    /// &gt; **Note:** API Gateway requires the use of AWS Certificate Manager (ACM) certificates instead of Identity and Access Management (IAM) certificates in regions that support ACM. Regions that support ACM can be found in the [Regions and Endpoints Documentation](https://docs.aws.amazon.com/general/latest/gr/rande.html#acm_region). To import an existing private key and certificate into ACM or request an ACM certificate, see the [`aws.acm.Certificate` resource](https://www.terraform.io/docs/providers/aws/r/acm_certificate.html).
    /// 
    /// &gt; **Note:** The `aws.apigateway.DomainName` resource expects dependency on the `aws.acm.CertificateValidation` as 
    /// only verified certificates can be used. This can be made either explicitly by adding the 
    /// `depends_on = [aws_acm_certificate_validation.cert]` attribute. Or implicitly by referring certificate ARN 
    /// from the validation resource where it will be available after the resource creation: 
    /// `regional_certificate_arn = aws_acm_certificate_validation.cert.certificate_arn`.
    /// 
    /// &gt; **Note:** All arguments including the private key will be stored in the raw state as plain-text.
    /// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_domain_name.html.markdown.
    /// </summary>
    public partial class DomainName : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificate_name`, `certificate_body`, `certificate_chain`, `certificate_private_key`, `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Output("certificateArn")]
        public Output<string?> CertificateArn { get; private set; } = null!;

        /// <summary>
        /// The certificate issued for the domain name
        /// being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
        /// `regional_certificate_name`.
        /// </summary>
        [Output("certificateBody")]
        public Output<string?> CertificateBody { get; private set; } = null!;

        /// <summary>
        /// The certificate for the CA that issued the
        /// certificate, along with any intermediate CA certificates required to
        /// create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`,
        /// `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Output("certificateChain")]
        public Output<string?> CertificateChain { get; private set; } = null!;

        /// <summary>
        /// The unique name to use when registering this
        /// certificate as an IAM server certificate. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
        /// `regional_certificate_name`. Required if `certificate_arn` is not set.
        /// </summary>
        [Output("certificateName")]
        public Output<string?> CertificateName { get; private set; } = null!;

        /// <summary>
        /// The private key associated with the
        /// domain certificate given in `certificate_body`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Output("certificatePrivateKey")]
        public Output<string?> CertificatePrivateKey { get; private set; } = null!;

        /// <summary>
        /// The upload date associated with the domain certificate.
        /// </summary>
        [Output("certificateUploadDate")]
        public Output<string> CertificateUploadDate { get; private set; } = null!;

        /// <summary>
        /// The hostname created by Cloudfront to represent
        /// the distribution that implements this domain name mapping.
        /// </summary>
        [Output("cloudfrontDomainName")]
        public Output<string> CloudfrontDomainName { get; private set; } = null!;

        /// <summary>
        /// For convenience, the hosted zone ID (`Z2FDTNDATAQYW2`)
        /// that can be used to create a Route53 alias record for the distribution.
        /// </summary>
        [Output("cloudfrontZoneId")]
        public Output<string> CloudfrontZoneId { get; private set; } = null!;

        /// <summary>
        /// The fully-qualified domain name to register
        /// </summary>
        [Output("domainName")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Configuration block defining API endpoint information including type. Defined below.
        /// </summary>
        [Output("endpointConfiguration")]
        public Output<Outputs.DomainNameEndpointConfiguration> EndpointConfiguration { get; private set; } = null!;

        /// <summary>
        /// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and `certificate_private_key`.
        /// </summary>
        [Output("regionalCertificateArn")]
        public Output<string?> RegionalCertificateArn { get; private set; } = null!;

        /// <summary>
        /// The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and
        /// `certificate_private_key`.
        /// </summary>
        [Output("regionalCertificateName")]
        public Output<string?> RegionalCertificateName { get; private set; } = null!;

        /// <summary>
        /// The hostname for the custom domain's regional endpoint.
        /// </summary>
        [Output("regionalDomainName")]
        public Output<string> RegionalDomainName { get; private set; } = null!;

        /// <summary>
        /// The hosted zone ID that can be used to create a Route53 alias record for the regional endpoint.
        /// </summary>
        [Output("regionalZoneId")]
        public Output<string> RegionalZoneId { get; private set; } = null!;

        /// <summary>
        /// The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
        /// </summary>
        [Output("securityPolicy")]
        public Output<string> SecurityPolicy { get; private set; } = null!;

        /// <summary>
        /// Key-value mapping of resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DomainName resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainName(string name, DomainNameArgs args, CustomResourceOptions? options = null)
            : base("aws:apigateway/domainName:DomainName", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private DomainName(string name, Input<string> id, DomainNameState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigateway/domainName:DomainName", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DomainName resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainName Get(string name, Input<string> id, DomainNameState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainName(name, id, state, options);
        }
    }

    public sealed class DomainNameArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificate_name`, `certificate_body`, `certificate_chain`, `certificate_private_key`, `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Input("certificateArn")]
        public Input<string>? CertificateArn { get; set; }

        /// <summary>
        /// The certificate issued for the domain name
        /// being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
        /// `regional_certificate_name`.
        /// </summary>
        [Input("certificateBody")]
        public Input<string>? CertificateBody { get; set; }

        /// <summary>
        /// The certificate for the CA that issued the
        /// certificate, along with any intermediate CA certificates required to
        /// create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`,
        /// `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Input("certificateChain")]
        public Input<string>? CertificateChain { get; set; }

        /// <summary>
        /// The unique name to use when registering this
        /// certificate as an IAM server certificate. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
        /// `regional_certificate_name`. Required if `certificate_arn` is not set.
        /// </summary>
        [Input("certificateName")]
        public Input<string>? CertificateName { get; set; }

        /// <summary>
        /// The private key associated with the
        /// domain certificate given in `certificate_body`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Input("certificatePrivateKey")]
        public Input<string>? CertificatePrivateKey { get; set; }

        /// <summary>
        /// The fully-qualified domain name to register
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Configuration block defining API endpoint information including type. Defined below.
        /// </summary>
        [Input("endpointConfiguration")]
        public Input<Inputs.DomainNameEndpointConfigurationArgs>? EndpointConfiguration { get; set; }

        /// <summary>
        /// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and `certificate_private_key`.
        /// </summary>
        [Input("regionalCertificateArn")]
        public Input<string>? RegionalCertificateArn { get; set; }

        /// <summary>
        /// The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and
        /// `certificate_private_key`.
        /// </summary>
        [Input("regionalCertificateName")]
        public Input<string>? RegionalCertificateName { get; set; }

        /// <summary>
        /// The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
        /// </summary>
        [Input("securityPolicy")]
        public Input<string>? SecurityPolicy { get; set; }

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

        public DomainNameArgs()
        {
        }
    }

    public sealed class DomainNameState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificate_name`, `certificate_body`, `certificate_chain`, `certificate_private_key`, `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Input("certificateArn")]
        public Input<string>? CertificateArn { get; set; }

        /// <summary>
        /// The certificate issued for the domain name
        /// being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
        /// `regional_certificate_name`.
        /// </summary>
        [Input("certificateBody")]
        public Input<string>? CertificateBody { get; set; }

        /// <summary>
        /// The certificate for the CA that issued the
        /// certificate, along with any intermediate CA certificates required to
        /// create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`,
        /// `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Input("certificateChain")]
        public Input<string>? CertificateChain { get; set; }

        /// <summary>
        /// The unique name to use when registering this
        /// certificate as an IAM server certificate. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
        /// `regional_certificate_name`. Required if `certificate_arn` is not set.
        /// </summary>
        [Input("certificateName")]
        public Input<string>? CertificateName { get; set; }

        /// <summary>
        /// The private key associated with the
        /// domain certificate given in `certificate_body`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and `regional_certificate_name`.
        /// </summary>
        [Input("certificatePrivateKey")]
        public Input<string>? CertificatePrivateKey { get; set; }

        /// <summary>
        /// The upload date associated with the domain certificate.
        /// </summary>
        [Input("certificateUploadDate")]
        public Input<string>? CertificateUploadDate { get; set; }

        /// <summary>
        /// The hostname created by Cloudfront to represent
        /// the distribution that implements this domain name mapping.
        /// </summary>
        [Input("cloudfrontDomainName")]
        public Input<string>? CloudfrontDomainName { get; set; }

        /// <summary>
        /// For convenience, the hosted zone ID (`Z2FDTNDATAQYW2`)
        /// that can be used to create a Route53 alias record for the distribution.
        /// </summary>
        [Input("cloudfrontZoneId")]
        public Input<string>? CloudfrontZoneId { get; set; }

        /// <summary>
        /// The fully-qualified domain name to register
        /// </summary>
        [Input("domainName")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Configuration block defining API endpoint information including type. Defined below.
        /// </summary>
        [Input("endpointConfiguration")]
        public Input<Inputs.DomainNameEndpointConfigurationGetArgs>? EndpointConfiguration { get; set; }

        /// <summary>
        /// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and `certificate_private_key`.
        /// </summary>
        [Input("regionalCertificateArn")]
        public Input<string>? RegionalCertificateArn { get; set; }

        /// <summary>
        /// The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and
        /// `certificate_private_key`.
        /// </summary>
        [Input("regionalCertificateName")]
        public Input<string>? RegionalCertificateName { get; set; }

        /// <summary>
        /// The hostname for the custom domain's regional endpoint.
        /// </summary>
        [Input("regionalDomainName")]
        public Input<string>? RegionalDomainName { get; set; }

        /// <summary>
        /// The hosted zone ID that can be used to create a Route53 alias record for the regional endpoint.
        /// </summary>
        [Input("regionalZoneId")]
        public Input<string>? RegionalZoneId { get; set; }

        /// <summary>
        /// The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
        /// </summary>
        [Input("securityPolicy")]
        public Input<string>? SecurityPolicy { get; set; }

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

        public DomainNameState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class DomainNameEndpointConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A list of endpoint types. This resource currently only supports managing a single value. Valid values: `EDGE` or `REGIONAL`. If unspecified, defaults to `EDGE`. Must be declared as `REGIONAL` in non-Commercial partitions. Refer to the [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/create-regional-api.html) for more information on the difference between edge-optimized and regional APIs.
        /// </summary>
        [Input("types", required: true)]
        public Input<string> Types { get; set; } = null!;

        public DomainNameEndpointConfigurationArgs()
        {
        }
    }

    public sealed class DomainNameEndpointConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A list of endpoint types. This resource currently only supports managing a single value. Valid values: `EDGE` or `REGIONAL`. If unspecified, defaults to `EDGE`. Must be declared as `REGIONAL` in non-Commercial partitions. Refer to the [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/create-regional-api.html) for more information on the difference between edge-optimized and regional APIs.
        /// </summary>
        [Input("types", required: true)]
        public Input<string> Types { get; set; } = null!;

        public DomainNameEndpointConfigurationGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class DomainNameEndpointConfiguration
    {
        /// <summary>
        /// A list of endpoint types. This resource currently only supports managing a single value. Valid values: `EDGE` or `REGIONAL`. If unspecified, defaults to `EDGE`. Must be declared as `REGIONAL` in non-Commercial partitions. Refer to the [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/create-regional-api.html) for more information on the difference between edge-optimized and regional APIs.
        /// </summary>
        public readonly string Types;

        [OutputConstructor]
        private DomainNameEndpointConfiguration(string types)
        {
            Types = types;
        }
    }
    }
}
