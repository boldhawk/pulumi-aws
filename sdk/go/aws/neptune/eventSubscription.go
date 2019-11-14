// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## Attributes
// 
// The following additional atttributes are provided:
// 
// * `id` - The name of the Neptune event notification subscription.
// * `arn` - The Amazon Resource Name of the Neptune event notification subscription.
// * `customerAwsId` - The AWS customer account associated with the Neptune event notification subscription.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/neptune_event_subscription.html.markdown.
type EventSubscription struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	Arn pulumi.StringOutput `pulumi:"arn"`

	CustomerAwsId pulumi.StringOutput `pulumi:"customerAwsId"`

	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`

	// A list of event categories for a `sourceType` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
	EventCategories pulumi.ArrayOutput `pulumi:"eventCategories"`

	// The name of the Neptune event subscription. By default generated by this provider.
	Name pulumi.StringOutput `pulumi:"name"`

	// The name of the Neptune event subscription. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`

	// The ARN of the SNS topic to send events to.
	SnsTopicArn pulumi.StringOutput `pulumi:"snsTopicArn"`

	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `sourceType` must also be specified.
	SourceIds pulumi.ArrayOutput `pulumi:"sourceIds"`

	// The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
	SourceType pulumi.StringOutput `pulumi:"sourceType"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewEventSubscription registers a new resource with the given unique name, arguments, and options.
func NewEventSubscription(ctx *pulumi.Context,
	name string, args *EventSubscriptionArgs, opts ...pulumi.ResourceOpt) (*EventSubscription, error) {
	if args == nil || args.SnsTopicArn == nil {
		return nil, errors.New("missing required argument 'SnsTopicArn'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["enabled"] = args.Enabled
		inputs["eventCategories"] = args.EventCategories
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["snsTopicArn"] = args.SnsTopicArn
		inputs["sourceIds"] = args.SourceIds
		inputs["sourceType"] = args.SourceType
		inputs["tags"] = args.Tags
	}
	var resource EventSubscription
	err := ctx.RegisterResource("aws:neptune/eventSubscription:EventSubscription", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventSubscription gets an existing EventSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.ID, state *EventSubscriptionState, opts ...pulumi.ResourceOpt) (*EventSubscription, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["customerAwsId"] = state.CustomerAwsId
		inputs["enabled"] = state.Enabled
		inputs["eventCategories"] = state.EventCategories
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["snsTopicArn"] = state.SnsTopicArn
		inputs["sourceIds"] = state.SourceIds
		inputs["sourceType"] = state.SourceType
		inputs["tags"] = state.Tags
	}
	var resource EventSubscription
	err := ctx.ReadResource("aws:neptune/eventSubscription:EventSubscription", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *EventSubscription) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *EventSubscription) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering EventSubscription resources.
type EventSubscriptionState struct {
	Arn pulumi.StringInput `pulumi:"arn"`
	CustomerAwsId pulumi.StringInput `pulumi:"customerAwsId"`
	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled pulumi.BoolInput `pulumi:"enabled"`
	// A list of event categories for a `sourceType` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
	EventCategories pulumi.ArrayInput `pulumi:"eventCategories"`
	// The name of the Neptune event subscription. By default generated by this provider.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Neptune event subscription. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// The ARN of the SNS topic to send events to.
	SnsTopicArn pulumi.StringInput `pulumi:"snsTopicArn"`
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `sourceType` must also be specified.
	SourceIds pulumi.ArrayInput `pulumi:"sourceIds"`
	// The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
	SourceType pulumi.StringInput `pulumi:"sourceType"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a EventSubscription resource.
type EventSubscriptionArgs struct {
	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled pulumi.BoolInput `pulumi:"enabled"`
	// A list of event categories for a `sourceType` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
	EventCategories pulumi.ArrayInput `pulumi:"eventCategories"`
	// The name of the Neptune event subscription. By default generated by this provider.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the Neptune event subscription. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// The ARN of the SNS topic to send events to.
	SnsTopicArn pulumi.StringInput `pulumi:"snsTopicArn"`
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `sourceType` must also be specified.
	SourceIds pulumi.ArrayInput `pulumi:"sourceIds"`
	// The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
	SourceType pulumi.StringInput `pulumi:"sourceType"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
