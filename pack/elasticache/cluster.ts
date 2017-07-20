// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Cluster extends lumi.NamedResource implements ClusterArgs {
    public readonly applyImmediately: boolean;
    public readonly availabilityZone: string;
    public readonly availabilityZones?: string[];
    public readonly azMode: string;
    public /*out*/ readonly cacheNodes: { address: string, availabilityZone: string, id: string, port: number }[];
    public /*out*/ readonly clusterAddress: string;
    public readonly clusterId: string;
    public /*out*/ readonly configurationEndpoint: string;
    public readonly engine: string;
    public readonly engineVersion: string;
    public readonly maintenanceWindow: string;
    public readonly nodeType: string;
    public readonly notificationTopicArn?: string;
    public readonly numCacheNodes: number;
    public readonly parameterGroupName: string;
    public readonly port: number;
    public /*out*/ readonly replicationGroupId: string;
    public readonly securityGroupIds: string[];
    public readonly securityGroupNames: string[];
    public readonly snapshotArns?: string[];
    public readonly snapshotName?: string;
    public readonly snapshotRetentionLimit?: number;
    public readonly snapshotWindow: string;
    public readonly subnetGroupName: string;
    public readonly tags?: {[key: string]: any};

    public static get(id: lumi.ID): Cluster {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): Cluster[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: ClusterArgs) {
        super(name);
        this.applyImmediately = args.applyImmediately;
        this.availabilityZone = args.availabilityZone;
        this.availabilityZones = args.availabilityZones;
        this.azMode = args.azMode;
        if (lumirt.defaultIfComputed(args.clusterId, "") === undefined) {
            throw new Error("Property argument 'clusterId' is required, but was missing");
        }
        this.clusterId = args.clusterId;
        if (lumirt.defaultIfComputed(args.engine, "") === undefined) {
            throw new Error("Property argument 'engine' is required, but was missing");
        }
        this.engine = args.engine;
        this.engineVersion = args.engineVersion;
        this.maintenanceWindow = args.maintenanceWindow;
        if (lumirt.defaultIfComputed(args.nodeType, "") === undefined) {
            throw new Error("Property argument 'nodeType' is required, but was missing");
        }
        this.nodeType = args.nodeType;
        this.notificationTopicArn = args.notificationTopicArn;
        if (lumirt.defaultIfComputed(args.numCacheNodes, "") === undefined) {
            throw new Error("Property argument 'numCacheNodes' is required, but was missing");
        }
        this.numCacheNodes = args.numCacheNodes;
        this.parameterGroupName = args.parameterGroupName;
        if (lumirt.defaultIfComputed(args.port, "") === undefined) {
            throw new Error("Property argument 'port' is required, but was missing");
        }
        this.port = args.port;
        this.securityGroupIds = args.securityGroupIds;
        this.securityGroupNames = args.securityGroupNames;
        this.snapshotArns = args.snapshotArns;
        this.snapshotName = args.snapshotName;
        this.snapshotRetentionLimit = args.snapshotRetentionLimit;
        this.snapshotWindow = args.snapshotWindow;
        this.subnetGroupName = args.subnetGroupName;
        this.tags = args.tags;
    }
}

export interface ClusterArgs {
    readonly applyImmediately?: boolean;
    readonly availabilityZone?: string;
    readonly availabilityZones?: string[];
    readonly azMode?: string;
    readonly clusterId: string;
    readonly engine: string;
    readonly engineVersion?: string;
    readonly maintenanceWindow?: string;
    readonly nodeType: string;
    readonly notificationTopicArn?: string;
    readonly numCacheNodes: number;
    readonly parameterGroupName?: string;
    readonly port: number;
    readonly securityGroupIds?: string[];
    readonly securityGroupNames?: string[];
    readonly snapshotArns?: string[];
    readonly snapshotName?: string;
    readonly snapshotRetentionLimit?: number;
    readonly snapshotWindow?: string;
    readonly subnetGroupName?: string;
    readonly tags?: {[key: string]: any};
}

