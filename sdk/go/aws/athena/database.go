// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package athena

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Athena database.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/athena_database.html.markdown.
type Database struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// Name of s3 bucket to save the results of the query execution.
	Bucket pulumi.StringOutput `pulumi:"bucket"`

	// The encryption key block AWS Athena uses to decrypt the data in S3, such as an AWS Key Management Service (AWS KMS) key. An `encryptionConfiguration` block is documented below.
	EncryptionConfiguration pulumi.AnyOutput `pulumi:"encryptionConfiguration"`

	// A boolean that indicates all tables should be deleted from the database so that the database can be destroyed without error. The tables are *not* recoverable.
	ForceDestroy pulumi.BoolOutput `pulumi:"forceDestroy"`

	// Name of the database to create.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOpt) (*Database, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["bucket"] = args.Bucket
		inputs["encryptionConfiguration"] = args.EncryptionConfiguration
		inputs["forceDestroy"] = args.ForceDestroy
		inputs["name"] = args.Name
	}
	var resource Database
	err := ctx.RegisterResource("aws:athena/database:Database", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DatabaseState, opts ...pulumi.ResourceOpt) (*Database, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["bucket"] = state.Bucket
		inputs["encryptionConfiguration"] = state.EncryptionConfiguration
		inputs["forceDestroy"] = state.ForceDestroy
		inputs["name"] = state.Name
	}
	var resource Database
	err := ctx.ReadResource("aws:athena/database:Database", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *Database) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *Database) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering Database resources.
type DatabaseState struct {
	// Name of s3 bucket to save the results of the query execution.
	Bucket pulumi.StringInput `pulumi:"bucket"`
	// The encryption key block AWS Athena uses to decrypt the data in S3, such as an AWS Key Management Service (AWS KMS) key. An `encryptionConfiguration` block is documented below.
	EncryptionConfiguration pulumi.AnyInput `pulumi:"encryptionConfiguration"`
	// A boolean that indicates all tables should be deleted from the database so that the database can be destroyed without error. The tables are *not* recoverable.
	ForceDestroy pulumi.BoolInput `pulumi:"forceDestroy"`
	// Name of the database to create.
	Name pulumi.StringInput `pulumi:"name"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// Name of s3 bucket to save the results of the query execution.
	Bucket pulumi.StringInput `pulumi:"bucket"`
	// The encryption key block AWS Athena uses to decrypt the data in S3, such as an AWS Key Management Service (AWS KMS) key. An `encryptionConfiguration` block is documented below.
	EncryptionConfiguration pulumi.AnyInput `pulumi:"encryptionConfiguration"`
	// A boolean that indicates all tables should be deleted from the database so that the database can be destroyed without error. The tables are *not* recoverable.
	ForceDestroy pulumi.BoolInput `pulumi:"forceDestroy"`
	// Name of the database to create.
	Name pulumi.StringInput `pulumi:"name"`
}
