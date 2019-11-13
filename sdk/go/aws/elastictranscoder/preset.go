// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elastictranscoder

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Elastic Transcoder preset resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/elastictranscoder_preset.html.markdown.
type Preset struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	Arn pulumi.StringOutput `pulumi:"arn"`

	// Audio parameters object (documented below).
	Audio pulumi.AnyOutput `pulumi:"audio"`

	// Codec options for the audio parameters (documented below)
	AudioCodecOptions pulumi.AnyOutput `pulumi:"audioCodecOptions"`

	// The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
	Container pulumi.StringOutput `pulumi:"container"`

	// A description of the preset (maximum 255 characters)
	Description pulumi.StringOutput `pulumi:"description"`

	// The name of the preset. (maximum 40 characters)
	Name pulumi.StringOutput `pulumi:"name"`

	// Thumbnail parameters object (documented below)
	Thumbnails pulumi.AnyOutput `pulumi:"thumbnails"`

	Type pulumi.StringOutput `pulumi:"type"`

	// Video parameters object (documented below)
	Video pulumi.AnyOutput `pulumi:"video"`

	VideoCodecOptions pulumi.MapOutput `pulumi:"videoCodecOptions"`

	// Watermark parameters for the video parameters (documented below)
	// * `videoCodecOptions` (Optional, Forces new resource) Codec options for the video parameters
	VideoWatermarks pulumi.ArrayOutput `pulumi:"videoWatermarks"`
}

// NewPreset registers a new resource with the given unique name, arguments, and options.
func NewPreset(ctx *pulumi.Context,
	name string, args *PresetArgs, opts ...pulumi.ResourceOpt) (*Preset, error) {
	if args == nil || args.Container == nil {
		return nil, errors.New("missing required argument 'Container'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["audio"] = args.Audio
		inputs["audioCodecOptions"] = args.AudioCodecOptions
		inputs["container"] = args.Container
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["thumbnails"] = args.Thumbnails
		inputs["type"] = args.Type
		inputs["video"] = args.Video
		inputs["videoCodecOptions"] = args.VideoCodecOptions
		inputs["videoWatermarks"] = args.VideoWatermarks
	}
	var resource Preset
	err := ctx.RegisterResource("aws:elastictranscoder/preset:Preset", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPreset gets an existing Preset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPreset(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PresetState, opts ...pulumi.ResourceOpt) (*Preset, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["audio"] = state.Audio
		inputs["audioCodecOptions"] = state.AudioCodecOptions
		inputs["container"] = state.Container
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["thumbnails"] = state.Thumbnails
		inputs["type"] = state.Type
		inputs["video"] = state.Video
		inputs["videoCodecOptions"] = state.VideoCodecOptions
		inputs["videoWatermarks"] = state.VideoWatermarks
	}
	var resource Preset
	err := ctx.ReadResource("aws:elastictranscoder/preset:Preset", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *Preset) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *Preset) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering Preset resources.
type PresetState struct {
	Arn pulumi.StringInput `pulumi:"arn"`
	// Audio parameters object (documented below).
	Audio pulumi.AnyInput `pulumi:"audio"`
	// Codec options for the audio parameters (documented below)
	AudioCodecOptions pulumi.AnyInput `pulumi:"audioCodecOptions"`
	// The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
	Container pulumi.StringInput `pulumi:"container"`
	// A description of the preset (maximum 255 characters)
	Description pulumi.StringInput `pulumi:"description"`
	// The name of the preset. (maximum 40 characters)
	Name pulumi.StringInput `pulumi:"name"`
	// Thumbnail parameters object (documented below)
	Thumbnails pulumi.AnyInput `pulumi:"thumbnails"`
	Type pulumi.StringInput `pulumi:"type"`
	// Video parameters object (documented below)
	Video pulumi.AnyInput `pulumi:"video"`
	VideoCodecOptions pulumi.MapInput `pulumi:"videoCodecOptions"`
	// Watermark parameters for the video parameters (documented below)
	// * `videoCodecOptions` (Optional, Forces new resource) Codec options for the video parameters
	VideoWatermarks pulumi.ArrayInput `pulumi:"videoWatermarks"`
}

// The set of arguments for constructing a Preset resource.
type PresetArgs struct {
	// Audio parameters object (documented below).
	Audio pulumi.AnyInput `pulumi:"audio"`
	// Codec options for the audio parameters (documented below)
	AudioCodecOptions pulumi.AnyInput `pulumi:"audioCodecOptions"`
	// The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
	Container pulumi.StringInput `pulumi:"container"`
	// A description of the preset (maximum 255 characters)
	Description pulumi.StringInput `pulumi:"description"`
	// The name of the preset. (maximum 40 characters)
	Name pulumi.StringInput `pulumi:"name"`
	// Thumbnail parameters object (documented below)
	Thumbnails pulumi.AnyInput `pulumi:"thumbnails"`
	Type pulumi.StringInput `pulumi:"type"`
	// Video parameters object (documented below)
	Video pulumi.AnyInput `pulumi:"video"`
	VideoCodecOptions pulumi.MapInput `pulumi:"videoCodecOptions"`
	// Watermark parameters for the video parameters (documented below)
	// * `videoCodecOptions` (Optional, Forces new resource) Codec options for the video parameters
	VideoWatermarks pulumi.ArrayInput `pulumi:"videoWatermarks"`
}
