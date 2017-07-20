// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class ListenerPolicy extends lumi.NamedResource implements ListenerPolicyArgs {
    public readonly loadBalancerName: string;
    public readonly loadBalancerPort: number;
    public readonly policyNames?: string[];

    public static get(id: lumi.ID): ListenerPolicy {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): ListenerPolicy[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: ListenerPolicyArgs) {
        super(name);
        if (lumirt.defaultIfComputed(args.loadBalancerName, "") === undefined) {
            throw new Error("Property argument 'loadBalancerName' is required, but was missing");
        }
        this.loadBalancerName = args.loadBalancerName;
        if (lumirt.defaultIfComputed(args.loadBalancerPort, "") === undefined) {
            throw new Error("Property argument 'loadBalancerPort' is required, but was missing");
        }
        this.loadBalancerPort = args.loadBalancerPort;
        this.policyNames = args.policyNames;
    }
}

export interface ListenerPolicyArgs {
    readonly loadBalancerName: string;
    readonly loadBalancerPort: number;
    readonly policyNames?: string[];
}

