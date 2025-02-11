// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Elastic Beanstalk Application Resource. Elastic Beanstalk allows
// you to deploy and manage applications in the AWS cloud without worrying about
// the infrastructure that runs those applications.
// 
// This resource creates an application that has one configuration template named
// `default`, and no application versions
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/elastic_beanstalk_application.html.markdown.
type Application struct {
	s *pulumi.ResourceState
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOpt) (*Application, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["appversionLifecycle"] = nil
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["tags"] = nil
	} else {
		inputs["appversionLifecycle"] = args.AppversionLifecycle
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["tags"] = args.Tags
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:elasticbeanstalk/application:Application", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Application{s: s}, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ApplicationState, opts ...pulumi.ResourceOpt) (*Application, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["appversionLifecycle"] = state.AppversionLifecycle
		inputs["arn"] = state.Arn
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("aws:elasticbeanstalk/application:Application", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Application{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Application) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Application) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *Application) AppversionLifecycle() pulumi.Output {
	return r.s.State["appversionLifecycle"]
}

// The ARN assigned by AWS for this Elastic Beanstalk Application.
func (r *Application) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// Short description of the application
func (r *Application) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// The name of the application, must be unique within your account
func (r *Application) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Key-value mapping of tags for the Elastic Beanstalk Application.
func (r *Application) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering Application resources.
type ApplicationState struct {
	AppversionLifecycle interface{}
	// The ARN assigned by AWS for this Elastic Beanstalk Application.
	Arn interface{}
	// Short description of the application
	Description interface{}
	// The name of the application, must be unique within your account
	Name interface{}
	// Key-value mapping of tags for the Elastic Beanstalk Application.
	Tags interface{}
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	AppversionLifecycle interface{}
	// Short description of the application
	Description interface{}
	// The name of the application, must be unique within your account
	Name interface{}
	// Key-value mapping of tags for the Elastic Beanstalk Application.
	Tags interface{}
}
