// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class ReplicationGroup extends lumi.NamedResource implements ReplicationGroupArgs {
    public readonly applyImmediately: boolean;
    public readonly autoMinorVersionUpgrade?: boolean;
    public readonly automaticFailoverEnabled?: boolean;
    public readonly availabilityZones?: string[];
    public readonly clusterMode?: { numNodeGroups: number, replicasPerNodeGroup: number }[];
    public /*out*/ readonly configurationEndpointAddress: string;
    public readonly engine?: string;
    public readonly engineVersion: string;
    public readonly maintenanceWindow: string;
    public readonly nodeType: string;
    public readonly notificationTopicArn?: string;
    public readonly numberCacheClusters: number;
    public readonly parameterGroupName: string;
    public readonly port: number;
    public /*out*/ readonly primaryEndpointAddress: string;
    public readonly replicationGroupDescription: string;
    public readonly replicationGroupId: string;
    public readonly securityGroupIds: string[];
    public readonly securityGroupNames: string[];
    public readonly snapshotArns?: string[];
    public readonly snapshotName?: string;
    public readonly snapshotRetentionLimit?: number;
    public readonly snapshotWindow: string;
    public readonly subnetGroupName: string;
    public readonly tags?: {[key: string]: any};

    public static get(id: lumi.ID): ReplicationGroup {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): ReplicationGroup[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: ReplicationGroupArgs) {
        super(name);
        this.applyImmediately = args.applyImmediately;
        this.autoMinorVersionUpgrade = args.autoMinorVersionUpgrade;
        this.automaticFailoverEnabled = args.automaticFailoverEnabled;
        this.availabilityZones = args.availabilityZones;
        this.clusterMode = args.clusterMode;
        this.engine = args.engine;
        this.engineVersion = args.engineVersion;
        this.maintenanceWindow = args.maintenanceWindow;
        if (lumirt.defaultIfComputed(args.nodeType, "") === undefined) {
            throw new Error("Property argument 'nodeType' is required, but was missing");
        }
        this.nodeType = args.nodeType;
        this.notificationTopicArn = args.notificationTopicArn;
        this.numberCacheClusters = args.numberCacheClusters;
        this.parameterGroupName = args.parameterGroupName;
        if (lumirt.defaultIfComputed(args.port, "") === undefined) {
            throw new Error("Property argument 'port' is required, but was missing");
        }
        this.port = args.port;
        if (lumirt.defaultIfComputed(args.replicationGroupDescription, "") === undefined) {
            throw new Error("Property argument 'replicationGroupDescription' is required, but was missing");
        }
        this.replicationGroupDescription = args.replicationGroupDescription;
        if (lumirt.defaultIfComputed(args.replicationGroupId, "") === undefined) {
            throw new Error("Property argument 'replicationGroupId' is required, but was missing");
        }
        this.replicationGroupId = args.replicationGroupId;
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

export interface ReplicationGroupArgs {
    readonly applyImmediately?: boolean;
    readonly autoMinorVersionUpgrade?: boolean;
    readonly automaticFailoverEnabled?: boolean;
    readonly availabilityZones?: string[];
    readonly clusterMode?: { numNodeGroups: number, replicasPerNodeGroup: number }[];
    readonly engine?: string;
    readonly engineVersion?: string;
    readonly maintenanceWindow?: string;
    readonly nodeType: string;
    readonly notificationTopicArn?: string;
    readonly numberCacheClusters?: number;
    readonly parameterGroupName?: string;
    readonly port: number;
    readonly replicationGroupDescription: string;
    readonly replicationGroupId: string;
    readonly securityGroupIds?: string[];
    readonly securityGroupNames?: string[];
    readonly snapshotArns?: string[];
    readonly snapshotName?: string;
    readonly snapshotRetentionLimit?: number;
    readonly snapshotWindow?: string;
    readonly subnetGroupName?: string;
    readonly tags?: {[key: string]: any};
}

