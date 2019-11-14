// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an IAM Server Certificate resource to upload Server Certificates.
// Certs uploaded to IAM can easily work with other AWS services such as:
// 
// - AWS Elastic Beanstalk
// - Elastic Load Balancing
// - CloudFront
// - AWS OpsWorks
// 
// For information about server certificates in IAM, see [Managing Server
// Certificates][2] in AWS Documentation.
// 
// > **Note:** All arguments including the private key will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_server_certificate.html.markdown.
type ServerCertificate struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The Amazon Resource Name (ARN) specifying the server certificate.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The contents of the public key certificate in
	// PEM-encoded format.
	CertificateBody pulumi.StringOutput `pulumi:"certificateBody"`

	// The contents of the certificate chain.
	// This is typically a concatenation of the PEM-encoded public key certificates
	// of the chain.
	CertificateChain pulumi.StringOutput `pulumi:"certificateChain"`

	// The name of the Server Certificate. Do not include the
	// path in this value. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`

	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`

	// The IAM path for the server certificate.  If it is not
	// included, it defaults to a slash (/). If this certificate is for use with
	// AWS CloudFront, the path must be in format `/cloudfront/your_path_here`.
	// See [IAM Identifiers][1] for more details on IAM Paths.
	Path pulumi.StringOutput `pulumi:"path"`

	// The contents of the private key in PEM-encoded format.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
}

// NewServerCertificate registers a new resource with the given unique name, arguments, and options.
func NewServerCertificate(ctx *pulumi.Context,
	name string, args *ServerCertificateArgs, opts ...pulumi.ResourceOpt) (*ServerCertificate, error) {
	if args == nil || args.CertificateBody == nil {
		return nil, errors.New("missing required argument 'CertificateBody'")
	}
	if args == nil || args.PrivateKey == nil {
		return nil, errors.New("missing required argument 'PrivateKey'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["arn"] = args.Arn
		inputs["certificateBody"] = args.CertificateBody
		inputs["certificateChain"] = args.CertificateChain
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["path"] = args.Path
		inputs["privateKey"] = args.PrivateKey
	}
	var resource ServerCertificate
	err := ctx.RegisterResource("aws:iam/serverCertificate:ServerCertificate", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerCertificate gets an existing ServerCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerCertificate(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ServerCertificateState, opts ...pulumi.ResourceOpt) (*ServerCertificate, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["certificateBody"] = state.CertificateBody
		inputs["certificateChain"] = state.CertificateChain
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["path"] = state.Path
		inputs["privateKey"] = state.PrivateKey
	}
	var resource ServerCertificate
	err := ctx.ReadResource("aws:iam/serverCertificate:ServerCertificate", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *ServerCertificate) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *ServerCertificate) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering ServerCertificate resources.
type ServerCertificateState struct {
	// The Amazon Resource Name (ARN) specifying the server certificate.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The contents of the public key certificate in
	// PEM-encoded format.
	CertificateBody pulumi.StringInput `pulumi:"certificateBody"`
	// The contents of the certificate chain.
	// This is typically a concatenation of the PEM-encoded public key certificates
	// of the chain.
	CertificateChain pulumi.StringInput `pulumi:"certificateChain"`
	// The name of the Server Certificate. Do not include the
	// path in this value. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// The IAM path for the server certificate.  If it is not
	// included, it defaults to a slash (/). If this certificate is for use with
	// AWS CloudFront, the path must be in format `/cloudfront/your_path_here`.
	// See [IAM Identifiers][1] for more details on IAM Paths.
	Path pulumi.StringInput `pulumi:"path"`
	// The contents of the private key in PEM-encoded format.
	PrivateKey pulumi.StringInput `pulumi:"privateKey"`
}

// The set of arguments for constructing a ServerCertificate resource.
type ServerCertificateArgs struct {
	// The Amazon Resource Name (ARN) specifying the server certificate.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The contents of the public key certificate in
	// PEM-encoded format.
	CertificateBody pulumi.StringInput `pulumi:"certificateBody"`
	// The contents of the certificate chain.
	// This is typically a concatenation of the PEM-encoded public key certificates
	// of the chain.
	CertificateChain pulumi.StringInput `pulumi:"certificateChain"`
	// The name of the Server Certificate. Do not include the
	// path in this value. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// The IAM path for the server certificate.  If it is not
	// included, it defaults to a slash (/). If this certificate is for use with
	// AWS CloudFront, the path must be in format `/cloudfront/your_path_here`.
	// See [IAM Identifiers][1] for more details on IAM Paths.
	Path pulumi.StringInput `pulumi:"path"`
	// The contents of the private key in PEM-encoded format.
	PrivateKey pulumi.StringInput `pulumi:"privateKey"`
}
