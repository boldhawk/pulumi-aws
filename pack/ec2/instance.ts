// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

import {InstanceType} from "./instanceType";

export class Instance extends lumi.NamedResource implements InstanceArgs {
    public readonly ami: string;
    public readonly associatePublicIpAddress: boolean;
    public readonly availabilityZone: string;
    public readonly disableApiTermination?: boolean;
    public readonly ebsBlockDevice: { deleteOnTermination?: boolean, deviceName: string, encrypted: boolean, iops: number, snapshotId: string, volumeSize: number, volumeType: string }[];
    public readonly ebsOptimized?: boolean;
    public readonly ephemeralBlockDevice: { deviceName: string, noDevice?: boolean, virtualName?: string }[];
    public readonly iamInstanceProfile?: string;
    public readonly instanceInitiatedShutdownBehavior?: string;
    public /*out*/ readonly instanceState: string;
    public readonly instanceType: InstanceType;
    public readonly ipv6AddressCount: number;
    public readonly ipv6Addresses: string[];
    public readonly keyName: string;
    public readonly monitoring?: boolean;
    public readonly networkInterface: { deleteOnTermination?: boolean, deviceIndex: number, networkInterfaceId: string }[];
    public /*out*/ readonly networkInterfaceId: string;
    public readonly placementGroup: string;
    public /*out*/ readonly primaryNetworkInterfaceId: string;
    public /*out*/ readonly privateDns: string;
    public readonly privateIp: string;
    public /*out*/ readonly publicDns: string;
    public /*out*/ readonly publicIp: string;
    public readonly rootBlockDevice: { deleteOnTermination?: boolean, iops: number, volumeSize: number, volumeType: string }[];
    public readonly securityGroups: string[];
    public readonly sourceDestCheck?: boolean;
    public readonly subnetId: string;
    public readonly tags?: {[key: string]: any};
    public readonly tenancy: string;
    public readonly userData?: string;
    public readonly volumeTags: {[key: string]: any};
    public readonly vpcSecurityGroupIds: string[];

    public static get(id: lumi.ID): Instance {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): Instance[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: InstanceArgs) {
        super(name);
        if (lumirt.defaultIfComputed(args.ami, "") === undefined) {
            throw new Error("Property argument 'ami' is required, but was missing");
        }
        this.ami = args.ami;
        this.associatePublicIpAddress = args.associatePublicIpAddress;
        this.availabilityZone = args.availabilityZone;
        this.disableApiTermination = args.disableApiTermination;
        this.ebsBlockDevice = args.ebsBlockDevice;
        this.ebsOptimized = args.ebsOptimized;
        this.ephemeralBlockDevice = args.ephemeralBlockDevice;
        this.iamInstanceProfile = args.iamInstanceProfile;
        this.instanceInitiatedShutdownBehavior = args.instanceInitiatedShutdownBehavior;
        if (lumirt.defaultIfComputed(args.instanceType, "") === undefined) {
            throw new Error("Property argument 'instanceType' is required, but was missing");
        }
        this.instanceType = args.instanceType;
        this.ipv6AddressCount = args.ipv6AddressCount;
        this.ipv6Addresses = args.ipv6Addresses;
        this.keyName = args.keyName;
        this.monitoring = args.monitoring;
        this.networkInterface = args.networkInterface;
        this.placementGroup = args.placementGroup;
        this.privateIp = args.privateIp;
        this.rootBlockDevice = args.rootBlockDevice;
        this.securityGroups = args.securityGroups;
        this.sourceDestCheck = args.sourceDestCheck;
        this.subnetId = args.subnetId;
        this.tags = args.tags;
        this.tenancy = args.tenancy;
        this.userData = args.userData;
        this.volumeTags = args.volumeTags;
        this.vpcSecurityGroupIds = args.vpcSecurityGroupIds;
    }
}

export interface InstanceArgs {
    readonly ami: string;
    readonly associatePublicIpAddress?: boolean;
    readonly availabilityZone?: string;
    readonly disableApiTermination?: boolean;
    readonly ebsBlockDevice?: { deleteOnTermination?: boolean, deviceName: string, encrypted: boolean, iops: number, snapshotId: string, volumeSize: number, volumeType: string }[];
    readonly ebsOptimized?: boolean;
    readonly ephemeralBlockDevice?: { deviceName: string, noDevice?: boolean, virtualName?: string }[];
    readonly iamInstanceProfile?: string;
    readonly instanceInitiatedShutdownBehavior?: string;
    readonly instanceType: InstanceType;
    readonly ipv6AddressCount?: number;
    readonly ipv6Addresses?: string[];
    readonly keyName?: string;
    readonly monitoring?: boolean;
    readonly networkInterface?: { deleteOnTermination?: boolean, deviceIndex: number, networkInterfaceId: string }[];
    readonly placementGroup?: string;
    readonly privateIp?: string;
    readonly rootBlockDevice?: { deleteOnTermination?: boolean, iops: number, volumeSize: number, volumeType: string }[];
    readonly securityGroups?: string[];
    readonly sourceDestCheck?: boolean;
    readonly subnetId?: string;
    readonly tags?: {[key: string]: any};
    readonly tenancy?: string;
    readonly userData?: string;
    readonly volumeTags?: {[key: string]: any};
    readonly vpcSecurityGroupIds?: string[];
}

