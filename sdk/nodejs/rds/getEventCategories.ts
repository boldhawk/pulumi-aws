// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * 
 * List the event categories of all the RDS resources. 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleEventCategories = pulumi.output(aws.rds.getEventCategories({}));
 * 
 * export const example = exampleEventCategories.eventCategories;
 * ```
 * 
 * List the event categories specific to the RDS resource `db-snapshot`.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleEventCategories = pulumi.output(aws.rds.getEventCategories({
 *     sourceType: "db-snapshot",
 * }));
 * 
 * export const example = exampleEventCategories.eventCategories;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/db_event_categories.html.markdown.
 */
export function getEventCategories(args?: GetEventCategoriesArgs, opts?: pulumi.InvokeOptions): Promise<GetEventCategoriesResult> & GetEventCategoriesResult {
    args = args || {};
    const promise: Promise<GetEventCategoriesResult> = pulumi.runtime.invoke("aws:rds/getEventCategories:getEventCategories", {
        "sourceType": args.sourceType,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getEventCategories.
 */
export interface GetEventCategoriesArgs {
    /**
     * The type of source that will be generating the events. Valid options are db-instance, db-security-group, db-parameter-group, db-snapshot, db-cluster or db-cluster-snapshot.
     */
    readonly sourceType?: string;
}

/**
 * A collection of values returned by getEventCategories.
 */
export interface GetEventCategoriesResult {
    /**
     * A list of the event categories.
     */
    readonly eventCategories: string[];
    readonly sourceType?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
