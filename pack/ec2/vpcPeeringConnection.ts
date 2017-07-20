// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class VpcPeeringConnection extends lumi.NamedResource implements VpcPeeringConnectionArgs {
    public /*out*/ readonly acceptStatus: string;
    public readonly accepter: { allowClassicLinkToRemoteVpc?: boolean, allowRemoteVpcDnsResolution?: boolean, allowVpcToRemoteClassicLink?: boolean }[];
    public readonly autoAccept?: boolean;
    public readonly peerOwnerId: string;
    public readonly peerVpcId: string;
    public readonly requester: { allowClassicLinkToRemoteVpc?: boolean, allowRemoteVpcDnsResolution?: boolean, allowVpcToRemoteClassicLink?: boolean }[];
    public readonly tags?: {[key: string]: any};
    public readonly vpcId: string;

    public static get(id: lumi.ID): VpcPeeringConnection {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): VpcPeeringConnection[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: VpcPeeringConnectionArgs) {
        super(name);
        this.accepter = args.accepter;
        this.autoAccept = args.autoAccept;
        this.peerOwnerId = args.peerOwnerId;
        if (lumirt.defaultIfComputed(args.peerVpcId, "") === undefined) {
            throw new Error("Property argument 'peerVpcId' is required, but was missing");
        }
        this.peerVpcId = args.peerVpcId;
        this.requester = args.requester;
        this.tags = args.tags;
        if (lumirt.defaultIfComputed(args.vpcId, "") === undefined) {
            throw new Error("Property argument 'vpcId' is required, but was missing");
        }
        this.vpcId = args.vpcId;
    }
}

export interface VpcPeeringConnectionArgs {
    readonly accepter?: { allowClassicLinkToRemoteVpc?: boolean, allowRemoteVpcDnsResolution?: boolean, allowVpcToRemoteClassicLink?: boolean }[];
    readonly autoAccept?: boolean;
    readonly peerOwnerId?: string;
    readonly peerVpcId: string;
    readonly requester?: { allowClassicLinkToRemoteVpc?: boolean, allowRemoteVpcDnsResolution?: boolean, allowVpcToRemoteClassicLink?: boolean }[];
    readonly tags?: {[key: string]: any};
    readonly vpcId: string;
}

