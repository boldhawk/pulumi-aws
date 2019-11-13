// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sns

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

//   Provides a resource for subscribing to SNS topics. Requires that an SNS topic exist for the subscription to attach to.
// This resource allows you to automatically place messages sent to SNS topics in SQS queues, send them as HTTP(S) POST requests
// to a given endpoint, send SMS messages, or notify devices / applications. The most likely use case will
// probably be SQS queues.
// 
// > **NOTE:** If the SNS topic and SQS queue are in different AWS regions, it is important for the "sns.TopicSubscription" to use an AWS provider that is in the same region of the SNS topic. If the "sns.TopicSubscription" is using a provider with a different region than the SNS topic, the subscription will fail to create.
// 
// > **NOTE:** Setup of cross-account subscriptions from SNS topics to SQS queues requires the provider to have access to BOTH accounts.
// 
// > **NOTE:** If SNS topic and SQS queue are in different AWS accounts but the same region it is important for the "sns.TopicSubscription" to use the AWS provider of the account with the SQS queue. If "sns.TopicSubscription" is using a Provider with a different account than the SQS queue, the provider creates the subscriptions but does not keep state and tries to re-create the subscription at every apply.
// 
// > **NOTE:** If SNS topic and SQS queue are in different AWS accounts and different AWS regions it is important to recognize that the subscription needs to be initiated from the account with the SQS queue but in the region of the SNS topic.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/sns_topic_subscription.html.markdown.
type TopicSubscription struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN of the subscription stored as a more user-friendly property
	Arn pulumi.StringOutput `pulumi:"arn"`

	// Integer indicating number of minutes to wait in retying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols (default is 1 minute).
	ConfirmationTimeoutInMinutes pulumi.IntOutput `pulumi:"confirmationTimeoutInMinutes"`

	// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
	DeliveryPolicy pulumi.StringOutput `pulumi:"deliveryPolicy"`

	// The endpoint to send data to, the contents will vary with the protocol. (see below for more information)
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`

	// Boolean indicating whether the end point is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) e.g., PagerDuty (default is false)
	EndpointAutoConfirms pulumi.BoolOutput `pulumi:"endpointAutoConfirms"`

	// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
	FilterPolicy pulumi.StringOutput `pulumi:"filterPolicy"`

	// The protocol to use. The possible values for this are: `sqs`, `sms`, `lambda`, `application`. (`http` or `https` are partially supported, see below) (`email` is option but unsupported, see below).
	Protocol pulumi.StringOutput `pulumi:"protocol"`

	// Boolean indicating whether or not to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property) (default is false).
	RawMessageDelivery pulumi.BoolOutput `pulumi:"rawMessageDelivery"`

	// The ARN of the SNS topic to subscribe to
	Topic pulumi.StringOutput `pulumi:"topic"`
}

// NewTopicSubscription registers a new resource with the given unique name, arguments, and options.
func NewTopicSubscription(ctx *pulumi.Context,
	name string, args *TopicSubscriptionArgs, opts ...pulumi.ResourceOpt) (*TopicSubscription, error) {
	if args == nil || args.Endpoint == nil {
		return nil, errors.New("missing required argument 'Endpoint'")
	}
	if args == nil || args.Protocol == nil {
		return nil, errors.New("missing required argument 'Protocol'")
	}
	if args == nil || args.Topic == nil {
		return nil, errors.New("missing required argument 'Topic'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["confirmationTimeoutInMinutes"] = args.ConfirmationTimeoutInMinutes
		inputs["deliveryPolicy"] = args.DeliveryPolicy
		inputs["endpoint"] = args.Endpoint
		inputs["endpointAutoConfirms"] = args.EndpointAutoConfirms
		inputs["filterPolicy"] = args.FilterPolicy
		inputs["protocol"] = args.Protocol
		inputs["rawMessageDelivery"] = args.RawMessageDelivery
		inputs["topic"] = args.Topic
	}
	var resource TopicSubscription
	err := ctx.RegisterResource("aws:sns/topicSubscription:TopicSubscription", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopicSubscription gets an existing TopicSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicSubscription(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TopicSubscriptionState, opts ...pulumi.ResourceOpt) (*TopicSubscription, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["confirmationTimeoutInMinutes"] = state.ConfirmationTimeoutInMinutes
		inputs["deliveryPolicy"] = state.DeliveryPolicy
		inputs["endpoint"] = state.Endpoint
		inputs["endpointAutoConfirms"] = state.EndpointAutoConfirms
		inputs["filterPolicy"] = state.FilterPolicy
		inputs["protocol"] = state.Protocol
		inputs["rawMessageDelivery"] = state.RawMessageDelivery
		inputs["topic"] = state.Topic
	}
	var resource TopicSubscription
	err := ctx.ReadResource("aws:sns/topicSubscription:TopicSubscription", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *TopicSubscription) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *TopicSubscription) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering TopicSubscription resources.
type TopicSubscriptionState struct {
	// The ARN of the subscription stored as a more user-friendly property
	Arn pulumi.StringInput `pulumi:"arn"`
	// Integer indicating number of minutes to wait in retying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols (default is 1 minute).
	ConfirmationTimeoutInMinutes pulumi.IntInput `pulumi:"confirmationTimeoutInMinutes"`
	// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
	DeliveryPolicy pulumi.StringInput `pulumi:"deliveryPolicy"`
	// The endpoint to send data to, the contents will vary with the protocol. (see below for more information)
	Endpoint pulumi.StringInput `pulumi:"endpoint"`
	// Boolean indicating whether the end point is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) e.g., PagerDuty (default is false)
	EndpointAutoConfirms pulumi.BoolInput `pulumi:"endpointAutoConfirms"`
	// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
	FilterPolicy pulumi.StringInput `pulumi:"filterPolicy"`
	// The protocol to use. The possible values for this are: `sqs`, `sms`, `lambda`, `application`. (`http` or `https` are partially supported, see below) (`email` is option but unsupported, see below).
	Protocol pulumi.StringInput `pulumi:"protocol"`
	// Boolean indicating whether or not to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property) (default is false).
	RawMessageDelivery pulumi.BoolInput `pulumi:"rawMessageDelivery"`
	// The ARN of the SNS topic to subscribe to
	Topic pulumi.StringInput `pulumi:"topic"`
}

// The set of arguments for constructing a TopicSubscription resource.
type TopicSubscriptionArgs struct {
	// Integer indicating number of minutes to wait in retying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols (default is 1 minute).
	ConfirmationTimeoutInMinutes pulumi.IntInput `pulumi:"confirmationTimeoutInMinutes"`
	// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
	DeliveryPolicy pulumi.StringInput `pulumi:"deliveryPolicy"`
	// The endpoint to send data to, the contents will vary with the protocol. (see below for more information)
	Endpoint pulumi.StringInput `pulumi:"endpoint"`
	// Boolean indicating whether the end point is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) e.g., PagerDuty (default is false)
	EndpointAutoConfirms pulumi.BoolInput `pulumi:"endpointAutoConfirms"`
	// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
	FilterPolicy pulumi.StringInput `pulumi:"filterPolicy"`
	// The protocol to use. The possible values for this are: `sqs`, `sms`, `lambda`, `application`. (`http` or `https` are partially supported, see below) (`email` is option but unsupported, see below).
	Protocol pulumi.StringInput `pulumi:"protocol"`
	// Boolean indicating whether or not to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property) (default is false).
	RawMessageDelivery pulumi.BoolInput `pulumi:"rawMessageDelivery"`
	// The ARN of the SNS topic to subscribe to
	Topic pulumi.StringInput `pulumi:"topic"`
}
