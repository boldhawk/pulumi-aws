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
	s *pulumi.ResourceState
}

// NewPreset registers a new resource with the given unique name, arguments, and options.
func NewPreset(ctx *pulumi.Context,
	name string, args *PresetArgs, opts ...pulumi.ResourceOpt) (*Preset, error) {
	if args == nil || args.Container == nil {
		return nil, errors.New("missing required argument 'Container'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["audio"] = nil
		inputs["audioCodecOptions"] = nil
		inputs["container"] = nil
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["thumbnails"] = nil
		inputs["type"] = nil
		inputs["video"] = nil
		inputs["videoCodecOptions"] = nil
		inputs["videoWatermarks"] = nil
	} else {
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
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:elastictranscoder/preset:Preset", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Preset{s: s}, nil
}

// GetPreset gets an existing Preset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPreset(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PresetState, opts ...pulumi.ResourceOpt) (*Preset, error) {
	inputs := make(map[string]interface{})
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
	s, err := ctx.ReadResource("aws:elastictranscoder/preset:Preset", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Preset{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Preset) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Preset) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *Preset) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// Audio parameters object (documented below).
func (r *Preset) Audio() pulumi.Output {
	return r.s.State["audio"]
}

// Codec options for the audio parameters (documented below)
func (r *Preset) AudioCodecOptions() pulumi.Output {
	return r.s.State["audioCodecOptions"]
}

// The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
func (r *Preset) Container() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["container"])
}

// A description of the preset (maximum 255 characters)
func (r *Preset) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// The name of the preset. (maximum 40 characters)
func (r *Preset) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Thumbnail parameters object (documented below)
func (r *Preset) Thumbnails() pulumi.Output {
	return r.s.State["thumbnails"]
}

func (r *Preset) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

// Video parameters object (documented below)
func (r *Preset) Video() pulumi.Output {
	return r.s.State["video"]
}

func (r *Preset) VideoCodecOptions() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["videoCodecOptions"])
}

// Watermark parameters for the video parameters (documented below)
// * `videoCodecOptions` (Optional, Forces new resource) Codec options for the video parameters
func (r *Preset) VideoWatermarks() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["videoWatermarks"])
}

// Input properties used for looking up and filtering Preset resources.
type PresetState struct {
	Arn interface{}
	// Audio parameters object (documented below).
	Audio interface{}
	// Codec options for the audio parameters (documented below)
	AudioCodecOptions interface{}
	// The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
	Container interface{}
	// A description of the preset (maximum 255 characters)
	Description interface{}
	// The name of the preset. (maximum 40 characters)
	Name interface{}
	// Thumbnail parameters object (documented below)
	Thumbnails interface{}
	Type interface{}
	// Video parameters object (documented below)
	Video interface{}
	VideoCodecOptions interface{}
	// Watermark parameters for the video parameters (documented below)
	// * `videoCodecOptions` (Optional, Forces new resource) Codec options for the video parameters
	VideoWatermarks interface{}
}

// The set of arguments for constructing a Preset resource.
type PresetArgs struct {
	// Audio parameters object (documented below).
	Audio interface{}
	// Codec options for the audio parameters (documented below)
	AudioCodecOptions interface{}
	// The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
	Container interface{}
	// A description of the preset (maximum 255 characters)
	Description interface{}
	// The name of the preset. (maximum 40 characters)
	Name interface{}
	// Thumbnail parameters object (documented below)
	Thumbnails interface{}
	Type interface{}
	// Video parameters object (documented below)
	Video interface{}
	VideoCodecOptions interface{}
	// Watermark parameters for the video parameters (documented below)
	// * `videoCodecOptions` (Optional, Forces new resource) Codec options for the video parameters
	VideoWatermarks interface{}
}
