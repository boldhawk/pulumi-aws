// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to lookup information about IAM Server Certificates.
// 
// ## Import 
// 
// The import function will read in certificate body, certificate chain (if it exists), id, name, path, and arn. 
// It will not retrieve the private key which is not available through the AWS API.   
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/iam_server_certificate.html.markdown.
func LookupServerCertificate(ctx *pulumi.Context, args *GetServerCertificateArgs) (*GetServerCertificateResult, error) {
var rv GetServerCertificateResult
	err := ctx.Invoke("aws:iam/getServerCertificate:getServerCertificate", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServerCertificate.
type GetServerCertificateArgs struct {
	// sort results by expiration date. returns the certificate with expiration date in furthest in the future.
	Latest bool `pulumi:"latest"`
	// exact name of the cert to lookup
	Name string `pulumi:"name"`
	// prefix of cert to filter by
	NamePrefix string `pulumi:"namePrefix"`
	// prefix of path to filter by
	PathPrefix string `pulumi:"pathPrefix"`
}

// A collection of values returned by getServerCertificate.
type GetServerCertificateResult struct {
	Arn string `pulumi:"arn"`
	CertificateBody string `pulumi:"certificateBody"`
	CertificateChain string `pulumi:"certificateChain"`
	ExpirationDate string `pulumi:"expirationDate"`
	Latest bool `pulumi:"latest"`
	Name string `pulumi:"name"`
	NamePrefix string `pulumi:"namePrefix"`
	Path string `pulumi:"path"`
	PathPrefix string `pulumi:"pathPrefix"`
	UploadDate string `pulumi:"uploadDate"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
