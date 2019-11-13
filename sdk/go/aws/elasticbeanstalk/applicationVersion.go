// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Elastic Beanstalk Application Version Resource. Elastic Beanstalk allows
// you to deploy and manage applications in the AWS cloud without worrying about
// the infrastructure that runs those applications.
// 
// This resource creates a Beanstalk Application Version that can be deployed to a Beanstalk
// Environment.
// 
// > **NOTE on Application Version Resource:**  When using the Application Version resource with multiple 
// Elastic Beanstalk Environments it is possible that an error may be returned
// when attempting to delete an Application Version while it is still in use by a different environment.
// To work around this you can:
// <ol>
// <li>Create each environment in a separate AWS account</li>
// <li>Create your `elasticbeanstalk.ApplicationVersion` resources with a unique names in your 
// Elastic Beanstalk Application. For example &lt;revision&gt;-&lt;environment&gt;.</li>
// </ol>
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/elastic_beanstalk_application_version.html.markdown.
type ApplicationVersion struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// Name of the Beanstalk Application the version is associated with.
	Application pulumi.StringOutput `pulumi:"application"`

	// The ARN assigned by AWS for this Elastic Beanstalk Application.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// S3 bucket that contains the Application Version source bundle.
	Bucket pulumi.StringOutput `pulumi:"bucket"`

	// Short description of the Application Version.
	Description pulumi.StringOutput `pulumi:"description"`

	// On delete, force an Application Version to be deleted when it may be in use
	// by multiple Elastic Beanstalk Environments.
	ForceDelete pulumi.BoolOutput `pulumi:"forceDelete"`

	// S3 object that is the Application Version source bundle.
	Key pulumi.StringOutput `pulumi:"key"`

	// A unique name for the this Application Version.
	Name pulumi.StringOutput `pulumi:"name"`

	// Key-value mapping of tags for the Elastic Beanstalk Application Version.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewApplicationVersion registers a new resource with the given unique name, arguments, and options.
func NewApplicationVersion(ctx *pulumi.Context,
	name string, args *ApplicationVersionArgs, opts ...pulumi.ResourceOpt) (*ApplicationVersion, error) {
	if args == nil || args.Application == nil {
		return nil, errors.New("missing required argument 'Application'")
	}
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	if args == nil || args.Key == nil {
		return nil, errors.New("missing required argument 'Key'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["application"] = args.Application
		inputs["bucket"] = args.Bucket
		inputs["description"] = args.Description
		inputs["forceDelete"] = args.ForceDelete
		inputs["key"] = args.Key
		inputs["name"] = args.Name
		inputs["tags"] = args.Tags
	}
	var resource ApplicationVersion
	err := ctx.RegisterResource("aws:elasticbeanstalk/applicationVersion:ApplicationVersion", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationVersion gets an existing ApplicationVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationVersion(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ApplicationVersionState, opts ...pulumi.ResourceOpt) (*ApplicationVersion, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["application"] = state.Application
		inputs["arn"] = state.Arn
		inputs["bucket"] = state.Bucket
		inputs["description"] = state.Description
		inputs["forceDelete"] = state.ForceDelete
		inputs["key"] = state.Key
		inputs["name"] = state.Name
		inputs["tags"] = state.Tags
	}
	var resource ApplicationVersion
	err := ctx.ReadResource("aws:elasticbeanstalk/applicationVersion:ApplicationVersion", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *ApplicationVersion) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *ApplicationVersion) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering ApplicationVersion resources.
type ApplicationVersionState struct {
	// Name of the Beanstalk Application the version is associated with.
	Application pulumi.StringInput `pulumi:"application"`
	// The ARN assigned by AWS for this Elastic Beanstalk Application.
	Arn pulumi.StringInput `pulumi:"arn"`
	// S3 bucket that contains the Application Version source bundle.
	Bucket pulumi.StringInput `pulumi:"bucket"`
	// Short description of the Application Version.
	Description pulumi.StringInput `pulumi:"description"`
	// On delete, force an Application Version to be deleted when it may be in use
	// by multiple Elastic Beanstalk Environments.
	ForceDelete pulumi.BoolInput `pulumi:"forceDelete"`
	// S3 object that is the Application Version source bundle.
	Key pulumi.StringInput `pulumi:"key"`
	// A unique name for the this Application Version.
	Name pulumi.StringInput `pulumi:"name"`
	// Key-value mapping of tags for the Elastic Beanstalk Application Version.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a ApplicationVersion resource.
type ApplicationVersionArgs struct {
	// Name of the Beanstalk Application the version is associated with.
	Application pulumi.StringInput `pulumi:"application"`
	// S3 bucket that contains the Application Version source bundle.
	Bucket pulumi.StringInput `pulumi:"bucket"`
	// Short description of the Application Version.
	Description pulumi.StringInput `pulumi:"description"`
	// On delete, force an Application Version to be deleted when it may be in use
	// by multiple Elastic Beanstalk Environments.
	ForceDelete pulumi.BoolInput `pulumi:"forceDelete"`
	// S3 object that is the Application Version source bundle.
	Key pulumi.StringInput `pulumi:"key"`
	// A unique name for the this Application Version.
	Name pulumi.StringInput `pulumi:"name"`
	// Key-value mapping of tags for the Elastic Beanstalk Application Version.
	Tags pulumi.MapInput `pulumi:"tags"`
}
