// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dynamodb

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a DynamoDB table item resource
// 
// > **Note:** This resource is not meant to be used for managing large amounts of data in your table, it is not designed to scale.
//   You should perform **regular backups** of all data in the table, see [AWS docs for more](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/BackupRestore.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dynamodb_table_item.html.markdown.
type TableItem struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// Hash key to use for lookups and identification of the item
	HashKey pulumi.StringOutput `pulumi:"hashKey"`

	// JSON representation of a map of attribute name/value pairs, one for each attribute.
	// Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
	Item pulumi.StringOutput `pulumi:"item"`

	// Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
	RangeKey pulumi.StringOutput `pulumi:"rangeKey"`

	// The name of the table to contain the item.
	TableName pulumi.StringOutput `pulumi:"tableName"`
}

// NewTableItem registers a new resource with the given unique name, arguments, and options.
func NewTableItem(ctx *pulumi.Context,
	name string, args *TableItemArgs, opts ...pulumi.ResourceOpt) (*TableItem, error) {
	if args == nil || args.HashKey == nil {
		return nil, errors.New("missing required argument 'HashKey'")
	}
	if args == nil || args.Item == nil {
		return nil, errors.New("missing required argument 'Item'")
	}
	if args == nil || args.TableName == nil {
		return nil, errors.New("missing required argument 'TableName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["hashKey"] = args.HashKey
		inputs["item"] = args.Item
		inputs["rangeKey"] = args.RangeKey
		inputs["tableName"] = args.TableName
	}
	var resource TableItem
	err := ctx.RegisterResource("aws:dynamodb/tableItem:TableItem", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTableItem gets an existing TableItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTableItem(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TableItemState, opts ...pulumi.ResourceOpt) (*TableItem, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["hashKey"] = state.HashKey
		inputs["item"] = state.Item
		inputs["rangeKey"] = state.RangeKey
		inputs["tableName"] = state.TableName
	}
	var resource TableItem
	err := ctx.ReadResource("aws:dynamodb/tableItem:TableItem", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *TableItem) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *TableItem) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering TableItem resources.
type TableItemState struct {
	// Hash key to use for lookups and identification of the item
	HashKey pulumi.StringInput `pulumi:"hashKey"`
	// JSON representation of a map of attribute name/value pairs, one for each attribute.
	// Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
	Item pulumi.StringInput `pulumi:"item"`
	// Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
	RangeKey pulumi.StringInput `pulumi:"rangeKey"`
	// The name of the table to contain the item.
	TableName pulumi.StringInput `pulumi:"tableName"`
}

// The set of arguments for constructing a TableItem resource.
type TableItemArgs struct {
	// Hash key to use for lookups and identification of the item
	HashKey pulumi.StringInput `pulumi:"hashKey"`
	// JSON representation of a map of attribute name/value pairs, one for each attribute.
	// Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
	Item pulumi.StringInput `pulumi:"item"`
	// Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
	RangeKey pulumi.StringInput `pulumi:"rangeKey"`
	// The name of the table to contain the item.
	TableName pulumi.StringInput `pulumi:"tableName"`
}
