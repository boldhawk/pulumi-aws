// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an [IAM service-linked role](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_service_linked_role.html.markdown.
type ServiceLinkedRole struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The Amazon Resource Name (ARN) specifying the role.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The AWS service to which this role is attached. You use a string similar to a URL but without the `http://` in front. For example: `elasticbeanstalk.amazonaws.com`. To find the full list of services that support service-linked roles, check [the docs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html).
	AwsServiceName pulumi.StringOutput `pulumi:"awsServiceName"`

	// The creation date of the IAM role.
	CreateDate pulumi.StringOutput `pulumi:"createDate"`

	// Additional string appended to the role name. Not all AWS services support custom suffixes.
	CustomSuffix pulumi.StringOutput `pulumi:"customSuffix"`

	// The description of the role.
	Description pulumi.StringOutput `pulumi:"description"`

	// The name of the role.
	Name pulumi.StringOutput `pulumi:"name"`

	// The path of the role.
	Path pulumi.StringOutput `pulumi:"path"`

	// The stable and unique string identifying the role.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
}

// NewServiceLinkedRole registers a new resource with the given unique name, arguments, and options.
func NewServiceLinkedRole(ctx *pulumi.Context,
	name string, args *ServiceLinkedRoleArgs, opts ...pulumi.ResourceOpt) (*ServiceLinkedRole, error) {
	if args == nil || args.AwsServiceName == nil {
		return nil, errors.New("missing required argument 'AwsServiceName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		inputs["awsServiceName"] = args.AwsServiceName
		inputs["customSuffix"] = args.CustomSuffix
		inputs["description"] = args.Description
	}
	var resource ServiceLinkedRole
	err := ctx.RegisterResource("aws:iam/serviceLinkedRole:ServiceLinkedRole", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceLinkedRole gets an existing ServiceLinkedRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceLinkedRole(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ServiceLinkedRoleState, opts ...pulumi.ResourceOpt) (*ServiceLinkedRole, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["awsServiceName"] = state.AwsServiceName
		inputs["createDate"] = state.CreateDate
		inputs["customSuffix"] = state.CustomSuffix
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["path"] = state.Path
		inputs["uniqueId"] = state.UniqueId
	}
	var resource ServiceLinkedRole
	err := ctx.ReadResource("aws:iam/serviceLinkedRole:ServiceLinkedRole", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *ServiceLinkedRole) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *ServiceLinkedRole) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering ServiceLinkedRole resources.
type ServiceLinkedRoleState struct {
	// The Amazon Resource Name (ARN) specifying the role.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The AWS service to which this role is attached. You use a string similar to a URL but without the `http://` in front. For example: `elasticbeanstalk.amazonaws.com`. To find the full list of services that support service-linked roles, check [the docs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html).
	AwsServiceName pulumi.StringInput `pulumi:"awsServiceName"`
	// The creation date of the IAM role.
	CreateDate pulumi.StringInput `pulumi:"createDate"`
	// Additional string appended to the role name. Not all AWS services support custom suffixes.
	CustomSuffix pulumi.StringInput `pulumi:"customSuffix"`
	// The description of the role.
	Description pulumi.StringInput `pulumi:"description"`
	// The name of the role.
	Name pulumi.StringInput `pulumi:"name"`
	// The path of the role.
	Path pulumi.StringInput `pulumi:"path"`
	// The stable and unique string identifying the role.
	UniqueId pulumi.StringInput `pulumi:"uniqueId"`
}

// The set of arguments for constructing a ServiceLinkedRole resource.
type ServiceLinkedRoleArgs struct {
	// The AWS service to which this role is attached. You use a string similar to a URL but without the `http://` in front. For example: `elasticbeanstalk.amazonaws.com`. To find the full list of services that support service-linked roles, check [the docs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html).
	AwsServiceName pulumi.StringInput `pulumi:"awsServiceName"`
	// Additional string appended to the role name. Not all AWS services support custom suffixes.
	CustomSuffix pulumi.StringInput `pulumi:"customSuffix"`
	// The description of the role.
	Description pulumi.StringInput `pulumi:"description"`
}
