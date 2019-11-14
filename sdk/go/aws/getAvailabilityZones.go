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
	var rv GetAvailabilityZonesResult
	err := ctx.Invoke("aws:index/getAvailabilityZones:getAvailabilityZones", args, &rv)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAvailabilityZones.
type GetAvailabilityZonesArgs struct {
	// List of blacklisted Availability Zone names.
	BlacklistedNames []string `pulumi:"blacklistedNames"`
	// List of blacklisted Availability Zone IDs.
	BlacklistedZoneIds []string `pulumi:"blacklistedZoneIds"`
	// Allows to filter list of Availability Zones based on their
	// current state. Can be either `"available"`, `"information"`, `"impaired"` or
	// `"unavailable"`. By default the list includes a complete set of Availability Zones
	// to which the underlying AWS account has access, regardless of their state.
	State string `pulumi:"state"`
}

// A collection of values returned by getAvailabilityZones.
type GetAvailabilityZonesResult struct {
	BlacklistedNames []string `pulumi:"blacklistedNames"`
	BlacklistedZoneIds []string `pulumi:"blacklistedZoneIds"`
	// A list of the Availability Zone names available to the account.
	Names []string `pulumi:"names"`
	State string `pulumi:"state"`
	// A list of the Availability Zone IDs available to the account.
	ZoneIds []string `pulumi:"zoneIds"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
