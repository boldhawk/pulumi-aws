// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sns

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an SNS platform application resource
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/sns_platform_application.html.markdown.
type PlatformApplication struct {
	s *pulumi.ResourceState
}

// NewPlatformApplication registers a new resource with the given unique name, arguments, and options.
func NewPlatformApplication(ctx *pulumi.Context,
	name string, args *PlatformApplicationArgs, opts ...pulumi.ResourceOpt) (*PlatformApplication, error) {
	if args == nil || args.Platform == nil {
		return nil, errors.New("missing required argument 'Platform'")
	}
	if args == nil || args.PlatformCredential == nil {
		return nil, errors.New("missing required argument 'PlatformCredential'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["eventDeliveryFailureTopicArn"] = nil
		inputs["eventEndpointCreatedTopicArn"] = nil
		inputs["eventEndpointDeletedTopicArn"] = nil
		inputs["eventEndpointUpdatedTopicArn"] = nil
		inputs["failureFeedbackRoleArn"] = nil
		inputs["name"] = nil
		inputs["platform"] = nil
		inputs["platformCredential"] = nil
		inputs["platformPrincipal"] = nil
		inputs["successFeedbackRoleArn"] = nil
		inputs["successFeedbackSampleRate"] = nil
	} else {
		inputs["eventDeliveryFailureTopicArn"] = args.EventDeliveryFailureTopicArn
		inputs["eventEndpointCreatedTopicArn"] = args.EventEndpointCreatedTopicArn
		inputs["eventEndpointDeletedTopicArn"] = args.EventEndpointDeletedTopicArn
		inputs["eventEndpointUpdatedTopicArn"] = args.EventEndpointUpdatedTopicArn
		inputs["failureFeedbackRoleArn"] = args.FailureFeedbackRoleArn
		inputs["name"] = args.Name
		inputs["platform"] = args.Platform
		inputs["platformCredential"] = args.PlatformCredential
		inputs["platformPrincipal"] = args.PlatformPrincipal
		inputs["successFeedbackRoleArn"] = args.SuccessFeedbackRoleArn
		inputs["successFeedbackSampleRate"] = args.SuccessFeedbackSampleRate
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:sns/platformApplication:PlatformApplication", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PlatformApplication{s: s}, nil
}

// GetPlatformApplication gets an existing PlatformApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlatformApplication(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PlatformApplicationState, opts ...pulumi.ResourceOpt) (*PlatformApplication, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["eventDeliveryFailureTopicArn"] = state.EventDeliveryFailureTopicArn
		inputs["eventEndpointCreatedTopicArn"] = state.EventEndpointCreatedTopicArn
		inputs["eventEndpointDeletedTopicArn"] = state.EventEndpointDeletedTopicArn
		inputs["eventEndpointUpdatedTopicArn"] = state.EventEndpointUpdatedTopicArn
		inputs["failureFeedbackRoleArn"] = state.FailureFeedbackRoleArn
		inputs["name"] = state.Name
		inputs["platform"] = state.Platform
		inputs["platformCredential"] = state.PlatformCredential
		inputs["platformPrincipal"] = state.PlatformPrincipal
		inputs["successFeedbackRoleArn"] = state.SuccessFeedbackRoleArn
		inputs["successFeedbackSampleRate"] = state.SuccessFeedbackSampleRate
	}
	s, err := ctx.ReadResource("aws:sns/platformApplication:PlatformApplication", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PlatformApplication{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *PlatformApplication) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *PlatformApplication) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the SNS platform application
func (r *PlatformApplication) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// SNS Topic triggered when a delivery to any of the platform endpoints associated with your platform application encounters a permanent failure.
func (r *PlatformApplication) EventDeliveryFailureTopicArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["eventDeliveryFailureTopicArn"])
}

// SNS Topic triggered when a new platform endpoint is added to your platform application.
func (r *PlatformApplication) EventEndpointCreatedTopicArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["eventEndpointCreatedTopicArn"])
}

// SNS Topic triggered when an existing platform endpoint is deleted from your platform application.
func (r *PlatformApplication) EventEndpointDeletedTopicArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["eventEndpointDeletedTopicArn"])
}

// SNS Topic triggered when an existing platform endpoint is changed from your platform application.
func (r *PlatformApplication) EventEndpointUpdatedTopicArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["eventEndpointUpdatedTopicArn"])
}

// The IAM role permitted to receive failure feedback for this application.
func (r *PlatformApplication) FailureFeedbackRoleArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["failureFeedbackRoleArn"])
}

// The friendly name for the SNS platform application
func (r *PlatformApplication) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The platform that the app is registered with. See [Platform][1] for supported platforms.
func (r *PlatformApplication) Platform() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["platform"])
}

// Application Platform credential. See [Credential][1] for type of credential required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
func (r *PlatformApplication) PlatformCredential() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["platformCredential"])
}

// Application Platform principal. See [Principal][2] for type of principal required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
func (r *PlatformApplication) PlatformPrincipal() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["platformPrincipal"])
}

// The IAM role permitted to receive success feedback for this application.
func (r *PlatformApplication) SuccessFeedbackRoleArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["successFeedbackRoleArn"])
}

// The percentage of success to sample (0-100)
func (r *PlatformApplication) SuccessFeedbackSampleRate() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["successFeedbackSampleRate"])
}

// Input properties used for looking up and filtering PlatformApplication resources.
type PlatformApplicationState struct {
	// The ARN of the SNS platform application
	Arn interface{}
	// SNS Topic triggered when a delivery to any of the platform endpoints associated with your platform application encounters a permanent failure.
	EventDeliveryFailureTopicArn interface{}
	// SNS Topic triggered when a new platform endpoint is added to your platform application.
	EventEndpointCreatedTopicArn interface{}
	// SNS Topic triggered when an existing platform endpoint is deleted from your platform application.
	EventEndpointDeletedTopicArn interface{}
	// SNS Topic triggered when an existing platform endpoint is changed from your platform application.
	EventEndpointUpdatedTopicArn interface{}
	// The IAM role permitted to receive failure feedback for this application.
	FailureFeedbackRoleArn interface{}
	// The friendly name for the SNS platform application
	Name interface{}
	// The platform that the app is registered with. See [Platform][1] for supported platforms.
	Platform interface{}
	// Application Platform credential. See [Credential][1] for type of credential required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
	PlatformCredential interface{}
	// Application Platform principal. See [Principal][2] for type of principal required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
	PlatformPrincipal interface{}
	// The IAM role permitted to receive success feedback for this application.
	SuccessFeedbackRoleArn interface{}
	// The percentage of success to sample (0-100)
	SuccessFeedbackSampleRate interface{}
}

// The set of arguments for constructing a PlatformApplication resource.
type PlatformApplicationArgs struct {
	// SNS Topic triggered when a delivery to any of the platform endpoints associated with your platform application encounters a permanent failure.
	EventDeliveryFailureTopicArn interface{}
	// SNS Topic triggered when a new platform endpoint is added to your platform application.
	EventEndpointCreatedTopicArn interface{}
	// SNS Topic triggered when an existing platform endpoint is deleted from your platform application.
	EventEndpointDeletedTopicArn interface{}
	// SNS Topic triggered when an existing platform endpoint is changed from your platform application.
	EventEndpointUpdatedTopicArn interface{}
	// The IAM role permitted to receive failure feedback for this application.
	FailureFeedbackRoleArn interface{}
	// The friendly name for the SNS platform application
	Name interface{}
	// The platform that the app is registered with. See [Platform][1] for supported platforms.
	Platform interface{}
	// Application Platform credential. See [Credential][1] for type of credential required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
	PlatformCredential interface{}
	// Application Platform principal. See [Principal][2] for type of principal required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
	PlatformPrincipal interface{}
	// The IAM role permitted to receive success feedback for this application.
	SuccessFeedbackRoleArn interface{}
	// The percentage of success to sample (0-100)
	SuccessFeedbackSampleRate interface{}
}
