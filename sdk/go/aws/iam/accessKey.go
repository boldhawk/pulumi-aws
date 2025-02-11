// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an IAM access key. This is a set of credentials that allow API requests to be made as an IAM user.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_access_key.html.markdown.
type AccessKey struct {
	s *pulumi.ResourceState
}

// NewAccessKey registers a new resource with the given unique name, arguments, and options.
func NewAccessKey(ctx *pulumi.Context,
	name string, args *AccessKeyArgs, opts ...pulumi.ResourceOpt) (*AccessKey, error) {
	if args == nil || args.User == nil {
		return nil, errors.New("missing required argument 'User'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["pgpKey"] = nil
		inputs["status"] = nil
		inputs["user"] = nil
	} else {
		inputs["pgpKey"] = args.PgpKey
		inputs["status"] = args.Status
		inputs["user"] = args.User
	}
	inputs["encryptedSecret"] = nil
	inputs["keyFingerprint"] = nil
	inputs["secret"] = nil
	inputs["sesSmtpPassword"] = nil
	s, err := ctx.RegisterResource("aws:iam/accessKey:AccessKey", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AccessKey{s: s}, nil
}

// GetAccessKey gets an existing AccessKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessKey(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AccessKeyState, opts ...pulumi.ResourceOpt) (*AccessKey, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["encryptedSecret"] = state.EncryptedSecret
		inputs["keyFingerprint"] = state.KeyFingerprint
		inputs["pgpKey"] = state.PgpKey
		inputs["secret"] = state.Secret
		inputs["sesSmtpPassword"] = state.SesSmtpPassword
		inputs["status"] = state.Status
		inputs["user"] = state.User
	}
	s, err := ctx.ReadResource("aws:iam/accessKey:AccessKey", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AccessKey{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AccessKey) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AccessKey) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The encrypted secret, base64 encoded.
// > **NOTE:** The encrypted secret may be decrypted using the command line,
// for example: `... | base64 --decode | keybase pgp decrypt`.
func (r *AccessKey) EncryptedSecret() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["encryptedSecret"])
}

// The fingerprint of the PGP key used to encrypt
// the secret
func (r *AccessKey) KeyFingerprint() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["keyFingerprint"])
}

// Either a base-64 encoded PGP public key, or a
// keybase username in the form `keybase:some_person_that_exists`.
func (r *AccessKey) PgpKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["pgpKey"])
}

// The secret access key. Note that this will be written
// to the state file. Please supply a `pgpKey` instead, which will prevent the
// secret from being stored in plain text
func (r *AccessKey) Secret() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["secret"])
}

// The secret access key converted into an SES SMTP
// password by applying [AWS's documented conversion
// algorithm](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/smtp-credentials.html#smtp-credentials-convert).
func (r *AccessKey) SesSmtpPassword() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sesSmtpPassword"])
}

// The access key status to apply. Defaults to `Active`.
// Valid values are `Active` and `Inactive`.
func (r *AccessKey) Status() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["status"])
}

// The IAM user to associate with this access key.
func (r *AccessKey) User() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["user"])
}

// Input properties used for looking up and filtering AccessKey resources.
type AccessKeyState struct {
	// The encrypted secret, base64 encoded.
	// > **NOTE:** The encrypted secret may be decrypted using the command line,
	// for example: `... | base64 --decode | keybase pgp decrypt`.
	EncryptedSecret interface{}
	// The fingerprint of the PGP key used to encrypt
	// the secret
	KeyFingerprint interface{}
	// Either a base-64 encoded PGP public key, or a
	// keybase username in the form `keybase:some_person_that_exists`.
	PgpKey interface{}
	// The secret access key. Note that this will be written
	// to the state file. Please supply a `pgpKey` instead, which will prevent the
	// secret from being stored in plain text
	Secret interface{}
	// The secret access key converted into an SES SMTP
	// password by applying [AWS's documented conversion
	// algorithm](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/smtp-credentials.html#smtp-credentials-convert).
	SesSmtpPassword interface{}
	// The access key status to apply. Defaults to `Active`.
	// Valid values are `Active` and `Inactive`.
	Status interface{}
	// The IAM user to associate with this access key.
	User interface{}
}

// The set of arguments for constructing a AccessKey resource.
type AccessKeyArgs struct {
	// Either a base-64 encoded PGP public key, or a
	// keybase username in the form `keybase:some_person_that_exists`.
	PgpKey interface{}
	// The access key status to apply. Defaults to `Active`.
	// Valid values are `Active` and `Inactive`.
	Status interface{}
	// The IAM user to associate with this access key.
	User interface{}
}
