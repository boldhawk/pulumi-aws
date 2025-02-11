// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Pinpoint Baidu Channel resource.
// 
// > **Note:** All arguments including the Api Key and Secret Key will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/pinpoint_baidu_channel.html.markdown.
type BaiduChannel struct {
	s *pulumi.ResourceState
}

// NewBaiduChannel registers a new resource with the given unique name, arguments, and options.
func NewBaiduChannel(ctx *pulumi.Context,
	name string, args *BaiduChannelArgs, opts ...pulumi.ResourceOpt) (*BaiduChannel, error) {
	if args == nil || args.ApiKey == nil {
		return nil, errors.New("missing required argument 'ApiKey'")
	}
	if args == nil || args.ApplicationId == nil {
		return nil, errors.New("missing required argument 'ApplicationId'")
	}
	if args == nil || args.SecretKey == nil {
		return nil, errors.New("missing required argument 'SecretKey'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["apiKey"] = nil
		inputs["applicationId"] = nil
		inputs["enabled"] = nil
		inputs["secretKey"] = nil
	} else {
		inputs["apiKey"] = args.ApiKey
		inputs["applicationId"] = args.ApplicationId
		inputs["enabled"] = args.Enabled
		inputs["secretKey"] = args.SecretKey
	}
	s, err := ctx.RegisterResource("aws:pinpoint/baiduChannel:BaiduChannel", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BaiduChannel{s: s}, nil
}

// GetBaiduChannel gets an existing BaiduChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBaiduChannel(ctx *pulumi.Context,
	name string, id pulumi.ID, state *BaiduChannelState, opts ...pulumi.ResourceOpt) (*BaiduChannel, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["apiKey"] = state.ApiKey
		inputs["applicationId"] = state.ApplicationId
		inputs["enabled"] = state.Enabled
		inputs["secretKey"] = state.SecretKey
	}
	s, err := ctx.ReadResource("aws:pinpoint/baiduChannel:BaiduChannel", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BaiduChannel{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *BaiduChannel) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *BaiduChannel) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Platform credential API key from Baidu.
func (r *BaiduChannel) ApiKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["apiKey"])
}

// The application ID.
func (r *BaiduChannel) ApplicationId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["applicationId"])
}

// Specifies whether to enable the channel. Defaults to `true`.
func (r *BaiduChannel) Enabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enabled"])
}

// Platform credential Secret key from Baidu.
func (r *BaiduChannel) SecretKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["secretKey"])
}

// Input properties used for looking up and filtering BaiduChannel resources.
type BaiduChannelState struct {
	// Platform credential API key from Baidu.
	ApiKey interface{}
	// The application ID.
	ApplicationId interface{}
	// Specifies whether to enable the channel. Defaults to `true`.
	Enabled interface{}
	// Platform credential Secret key from Baidu.
	SecretKey interface{}
}

// The set of arguments for constructing a BaiduChannel resource.
type BaiduChannelArgs struct {
	// Platform credential API key from Baidu.
	ApiKey interface{}
	// The application ID.
	ApplicationId interface{}
	// Specifies whether to enable the channel. Defaults to `true`.
	Enabled interface{}
	// Platform credential Secret key from Baidu.
	SecretKey interface{}
}
