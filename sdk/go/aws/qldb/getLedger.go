// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package qldb

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to fetch information about a Quantum Ledger Database.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/qldb_ledger.html.markdown.
func LookupLedger(ctx *pulumi.Context, args *GetLedgerArgs) (*GetLedgerResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("aws:qldb/getLedger:getLedger", inputs)
	if err != nil {
		return nil, err
	}
	return &GetLedgerResult{
		Arn: outputs["arn"],
		DeletionProtection: outputs["deletionProtection"],
		Name: outputs["name"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getLedger.
type GetLedgerArgs struct {
	// The friendly name of the ledger to match.
	Name interface{}
}

// A collection of values returned by getLedger.
type GetLedgerResult struct {
	// Amazon Resource Name (ARN) of the ledger.
	Arn interface{}
	// Deletion protection on the QLDB Ledger instance. Set to `true` by default. 
	DeletionProtection interface{}
	Name interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
