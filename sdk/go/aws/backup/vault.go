// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS Backup vault resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/backup_vault.html.markdown.
type Vault struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN of the vault.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The server-side encryption key that is used to protect your backups.
	KmsKeyArn pulumi.StringOutput `pulumi:"kmsKeyArn"`

	// Name of the backup vault to create.
	Name pulumi.StringOutput `pulumi:"name"`

	// The number of recovery points that are stored in a backup vault.
	RecoveryPoints pulumi.IntOutput `pulumi:"recoveryPoints"`

	// Metadata that you can assign to help organize the resources that you create.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewVault registers a new resource with the given unique name, arguments, and options.
func NewVault(ctx *pulumi.Context,
	name string, args *VaultArgs, opts ...pulumi.ResourceOpt) (*Vault, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["kmsKeyArn"] = args.KmsKeyArn
		inputs["name"] = args.Name
		inputs["tags"] = args.Tags
	}
	var resource Vault
	err := ctx.RegisterResource("aws:backup/vault:Vault", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVault gets an existing Vault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVault(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VaultState, opts ...pulumi.ResourceOpt) (*Vault, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["kmsKeyArn"] = state.KmsKeyArn
		inputs["name"] = state.Name
		inputs["recoveryPoints"] = state.RecoveryPoints
		inputs["tags"] = state.Tags
	}
	var resource Vault
	err := ctx.ReadResource("aws:backup/vault:Vault", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *Vault) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *Vault) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering Vault resources.
type VaultState struct {
	// The ARN of the vault.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The server-side encryption key that is used to protect your backups.
	KmsKeyArn pulumi.StringInput `pulumi:"kmsKeyArn"`
	// Name of the backup vault to create.
	Name pulumi.StringInput `pulumi:"name"`
	// The number of recovery points that are stored in a backup vault.
	RecoveryPoints pulumi.IntInput `pulumi:"recoveryPoints"`
	// Metadata that you can assign to help organize the resources that you create.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a Vault resource.
type VaultArgs struct {
	// The server-side encryption key that is used to protect your backups.
	KmsKeyArn pulumi.StringInput `pulumi:"kmsKeyArn"`
	// Name of the backup vault to create.
	Name pulumi.StringInput `pulumi:"name"`
	// Metadata that you can assign to help organize the resources that you create.
	Tags pulumi.MapInput `pulumi:"tags"`
}
