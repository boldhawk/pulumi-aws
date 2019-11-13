// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The Availability Zones data source allows access to the list of AWS
// Availability Zones which can be accessed by an AWS account within the region
// configured in the provider.
// 
// This is different from the `.getAvailabilityZone` (singular) data source,
// which provides some details about a specific availability zone.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/availability_zones.html.markdown.
func LookupAvailabilityZones(ctx *pulumi.Context, args *GetAvailabilityZonesArgs) (*GetAvailabilityZonesResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["blacklistedNames"] = args.BlacklistedNames
		inputs["blacklistedZoneIds"] = args.BlacklistedZoneIds
		inputs["state"] = args.State
	}
	outputs, err := ctx.Invoke("aws:index/getAvailabilityZones:getAvailabilityZones", inputs)
	if err != nil {
		return nil, err
	}
	return &GetAvailabilityZonesResult{
		BlacklistedNames: outputs["blacklistedNames"],
		BlacklistedZoneIds: outputs["blacklistedZoneIds"],
		Names: outputs["names"],
		State: outputs["state"],
		ZoneIds: outputs["zoneIds"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getAvailabilityZones.
type GetAvailabilityZonesArgs struct {
	// List of blacklisted Availability Zone names.
	BlacklistedNames pulumi.ArrayInput `pulumi:"blacklistedNames"`
	// List of blacklisted Availability Zone IDs.
	BlacklistedZoneIds pulumi.ArrayInput `pulumi:"blacklistedZoneIds"`
	// Allows to filter list of Availability Zones based on their
	// current state. Can be either `"available"`, `"information"`, `"impaired"` or
	// `"unavailable"`. By default the list includes a complete set of Availability Zones
	// to which the underlying AWS account has access, regardless of their state.
	State pulumi.StringInput `pulumi:"state"`
}

// A collection of values returned by getAvailabilityZones.
type GetAvailabilityZonesResult struct {
	BlacklistedNames []interface{} `pulumi:"blacklistedNames"`
	BlacklistedZoneIds []interface{} `pulumi:"blacklistedZoneIds"`
	// A list of the Availability Zone names available to the account.
	Names []interface{} `pulumi:"names"`
	State string `pulumi:"state"`
	// A list of the Availability Zone IDs available to the account.
	ZoneIds []interface{} `pulumi:"zoneIds"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
