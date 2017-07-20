// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Instance extends lumi.NamedResource implements InstanceArgs {
    public /*out*/ readonly arn: string;
    public readonly availabilityZone: string;
    public readonly blueprintId: string;
    public readonly bundleId: string;
    public /*out*/ readonly cpuCount: number;
    public /*out*/ readonly createdAt: string;
    public /*out*/ readonly ipv6Address: string;
    public /*out*/ readonly isStaticIp: boolean;
    public readonly keyPairName?: string;
    public readonly instanceName?: string;
    public /*out*/ readonly privateIpAddress: string;
    public /*out*/ readonly publicIpAddress: string;
    public /*out*/ readonly ramSize: number;
    public readonly userData?: string;
    public /*out*/ readonly username: string;

    public static get(id: lumi.ID): Instance {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): Instance[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: InstanceArgs) {
        super(name);
        if (lumirt.defaultIfComputed(args.availabilityZone, "") === undefined) {
            throw new Error("Property argument 'availabilityZone' is required, but was missing");
        }
        this.availabilityZone = args.availabilityZone;
        if (lumirt.defaultIfComputed(args.blueprintId, "") === undefined) {
            throw new Error("Property argument 'blueprintId' is required, but was missing");
        }
        this.blueprintId = args.blueprintId;
        if (lumirt.defaultIfComputed(args.bundleId, "") === undefined) {
            throw new Error("Property argument 'bundleId' is required, but was missing");
        }
        this.bundleId = args.bundleId;
        this.keyPairName = args.keyPairName;
        this.instanceName = args.instanceName;
        this.userData = args.userData;
    }
}

export interface InstanceArgs {
    readonly availabilityZone: string;
    readonly blueprintId: string;
    readonly bundleId: string;
    readonly keyPairName?: string;
    readonly instanceName?: string;
    readonly userData?: string;
}

