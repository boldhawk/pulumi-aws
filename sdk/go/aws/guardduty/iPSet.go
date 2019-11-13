// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage a GuardDuty IPSet.
// 
// > **Note:** Currently in GuardDuty, users from member accounts cannot upload and further manage IPSets. IPSets that are uploaded by the master account are imposed on GuardDuty functionality in its member accounts. See the [GuardDuty API Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/create-ip-set.html)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/guardduty_ipset.html.markdown.
type IPSet struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// Specifies whether GuardDuty is to start using the uploaded IPSet.
	Activate pulumi.BoolOutput `pulumi:"activate"`

	// The detector ID of the GuardDuty.
	DetectorId pulumi.StringOutput `pulumi:"detectorId"`

	// The format of the file that contains the IPSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format pulumi.StringOutput `pulumi:"format"`

	// The URI of the file that contains the IPSet.
	Location pulumi.StringOutput `pulumi:"location"`

	// The friendly name to identify the IPSet.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewIPSet registers a new resource with the given unique name, arguments, and options.
func NewIPSet(ctx *pulumi.Context,
	name string, args *IPSetArgs, opts ...pulumi.ResourceOpt) (*IPSet, error) {
	if args == nil || args.Activate == nil {
		return nil, errors.New("missing required argument 'Activate'")
	}
	if args == nil || args.DetectorId == nil {
		return nil, errors.New("missing required argument 'DetectorId'")
	}
	if args == nil || args.Format == nil {
		return nil, errors.New("missing required argument 'Format'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["activate"] = args.Activate
		inputs["detectorId"] = args.DetectorId
		inputs["format"] = args.Format
		inputs["location"] = args.Location
		inputs["name"] = args.Name
	}
	var resource IPSet
	err := ctx.RegisterResource("aws:guardduty/iPSet:IPSet", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIPSet gets an existing IPSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIPSet(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IPSetState, opts ...pulumi.ResourceOpt) (*IPSet, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["activate"] = state.Activate
		inputs["detectorId"] = state.DetectorId
		inputs["format"] = state.Format
		inputs["location"] = state.Location
		inputs["name"] = state.Name
	}
	var resource IPSet
	err := ctx.ReadResource("aws:guardduty/iPSet:IPSet", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *IPSet) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *IPSet) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering IPSet resources.
type IPSetState struct {
	// Specifies whether GuardDuty is to start using the uploaded IPSet.
	Activate pulumi.BoolInput `pulumi:"activate"`
	// The detector ID of the GuardDuty.
	DetectorId pulumi.StringInput `pulumi:"detectorId"`
	// The format of the file that contains the IPSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format pulumi.StringInput `pulumi:"format"`
	// The URI of the file that contains the IPSet.
	Location pulumi.StringInput `pulumi:"location"`
	// The friendly name to identify the IPSet.
	Name pulumi.StringInput `pulumi:"name"`
}

// The set of arguments for constructing a IPSet resource.
type IPSetArgs struct {
	// Specifies whether GuardDuty is to start using the uploaded IPSet.
	Activate pulumi.BoolInput `pulumi:"activate"`
	// The detector ID of the GuardDuty.
	DetectorId pulumi.StringInput `pulumi:"detectorId"`
	// The format of the file that contains the IPSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format pulumi.StringInput `pulumi:"format"`
	// The URI of the file that contains the IPSet.
	Location pulumi.StringInput `pulumi:"location"`
	// The friendly name to identify the IPSet.
	Name pulumi.StringInput `pulumi:"name"`
}
