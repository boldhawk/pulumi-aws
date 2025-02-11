// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides information about a "classic" Elastic Load Balancer (ELB).
 * See [LB Data Source](https://www.terraform.io/docs/providers/aws/d/lb.html) if you are looking for "v2"
 * Application Load Balancer (ALB) or Network Load Balancer (NLB).
 * 
 * This data source can prove useful when a module accepts an LB as an input
 * variable and needs to, for example, determine the security groups associated
 * with it, etc.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const config = new pulumi.Config();
 * const lbName = config.get("lbName") || "";
 * 
 * const test = aws.elb.getLoadBalancer({
 *     name: lbName,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elb_legacy.html.markdown.
 */
export function getLoadBalancer(args: GetLoadBalancerArgs, opts?: pulumi.InvokeOptions): Promise<GetLoadBalancerResult> & GetLoadBalancerResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetLoadBalancerResult> = pulumi.runtime.invoke("aws:elasticloadbalancing/getLoadBalancer:getLoadBalancer", {
        "name": args.name,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getLoadBalancer.
 */
export interface GetLoadBalancerArgs {
    /**
     * The unique name of the load balancer.
     */
    readonly name: string;
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getLoadBalancer.
 */
export interface GetLoadBalancerResult {
    readonly accessLogs: outputs.elasticloadbalancing.GetLoadBalancerAccessLogs;
    readonly availabilityZones: string[];
    readonly connectionDraining: boolean;
    readonly connectionDrainingTimeout: number;
    readonly crossZoneLoadBalancing: boolean;
    readonly dnsName: string;
    readonly healthCheck: outputs.elasticloadbalancing.GetLoadBalancerHealthCheck;
    readonly idleTimeout: number;
    readonly instances: string[];
    readonly internal: boolean;
    readonly listeners: outputs.elasticloadbalancing.GetLoadBalancerListener[];
    readonly name: string;
    readonly securityGroups: string[];
    readonly sourceSecurityGroup: string;
    readonly sourceSecurityGroupId: string;
    readonly subnets: string[];
    readonly tags: {[key: string]: any};
    readonly zoneId: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
