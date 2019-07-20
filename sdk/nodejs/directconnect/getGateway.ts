// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieve information about a Direct Connect Gateway.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = pulumi.output(aws.directconnect.getGateway({
 *     name: "example",
 * }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/dx_gateway.html.markdown.
 */
export function getGateway(args: GetGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayResult> & GetGatewayResult {
    const promise: Promise<GetGatewayResult> = pulumi.runtime.invoke("aws:directconnect/getGateway:getGateway", {
        "name": args.name,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getGateway.
 */
export interface GetGatewayArgs {
    /**
     * The name of the gateway to retrieve.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getGateway.
 */
export interface GetGatewayResult {
    /**
     * The ASN on the Amazon side of the connection.
     */
    readonly amazonSideAsn: string;
    readonly name: string;
    /**
     * AWS Account ID of the gateway.
     */
    readonly ownerAccountId: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
