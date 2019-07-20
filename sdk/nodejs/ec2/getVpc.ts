// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `aws_vpc` provides details about a specific VPC.
 * 
 * This resource can prove useful when a module accepts a vpc id as
 * an input variable and needs to, for example, determine the CIDR block of that
 * VPC.
 * 
 * ## Example Usage
 * 
 * The following example shows how one might accept a VPC id as a variable
 * and use this data source to obtain the data necessary to create a subnet
 * within it.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const config = new pulumi.Config();
 * const vpcId = config.require("vpcId");
 * 
 * const selected = pulumi.output(aws.ec2.getVpc({
 *     id: vpcId,
 * }));
 * const example = new aws.ec2.Subnet("example", {
 *     availabilityZone: "us-west-2a",
 *     cidrBlock: selected.apply(selected => (() => {
 *         throw "tf2pulumi error: NYI: call to cidrsubnet";
 *         return (() => { throw "NYI: call to cidrsubnet"; })();
 *     })()),
 *     vpcId: selected.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/vpc.html.markdown.
 */
export function getVpc(args?: GetVpcArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcResult> & GetVpcResult {
    args = args || {};
    const promise: Promise<GetVpcResult> = pulumi.runtime.invoke("aws:ec2/getVpc:getVpc", {
        "cidrBlock": args.cidrBlock,
        "default": args.default,
        "dhcpOptionsId": args.dhcpOptionsId,
        "filters": args.filters,
        "id": args.id,
        "state": args.state,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getVpc.
 */
export interface GetVpcArgs {
    /**
     * The cidr block of the desired VPC.
     */
    readonly cidrBlock?: string;
    /**
     * Boolean constraint on whether the desired VPC is
     * the default VPC for the region.
     */
    readonly default?: boolean;
    /**
     * The DHCP options id of the desired VPC.
     */
    readonly dhcpOptionsId?: string;
    /**
     * Custom filter block as described below.
     */
    readonly filters?: { name: string, values: string[] }[];
    /**
     * The id of the specific VPC to retrieve.
     */
    readonly id?: string;
    /**
     * The current state of the desired VPC.
     * Can be either `"pending"` or `"available"`.
     */
    readonly state?: string;
    /**
     * A mapping of tags, each pair of which must exactly match
     * a pair on the desired VPC.
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getVpc.
 */
export interface GetVpcResult {
    /**
     * Amazon Resource Name (ARN) of VPC
     */
    readonly arn: string;
    /**
     * The CIDR block for the association.
     */
    readonly cidrBlock: string;
    readonly cidrBlockAssociations: { associationId: string, cidrBlock: string, state: string }[];
    readonly default: boolean;
    readonly dhcpOptionsId: string;
    /**
     * Whether or not the VPC has DNS hostname support
     */
    readonly enableDnsHostnames: boolean;
    /**
     * Whether or not the VPC has DNS support
     */
    readonly enableDnsSupport: boolean;
    readonly filters?: { name: string, values: string[] }[];
    readonly id: string;
    /**
     * The allowed tenancy of instances launched into the
     * selected VPC. May be any of `"default"`, `"dedicated"`, or `"host"`.
     */
    readonly instanceTenancy: string;
    /**
     * The association ID for the IPv6 CIDR block.
     */
    readonly ipv6AssociationId: string;
    /**
     * The IPv6 CIDR block.
     */
    readonly ipv6CidrBlock: string;
    /**
     * The ID of the main route table associated with this VPC.
     */
    readonly mainRouteTableId: string;
    /**
     * The ID of the AWS account that owns the VPC.
     */
    readonly ownerId: string;
    /**
     * The State of the association.
     */
    readonly state: string;
    readonly tags: {[key: string]: any};
}
