// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Snapshot extends lumi.NamedResource implements SnapshotArgs {
    public /*out*/ readonly allocatedStorage: number;
    public /*out*/ readonly availabilityZone: string;
    public readonly dbInstanceIdentifier: string;
    public /*out*/ readonly dbSnapshotArn: string;
    public readonly dbSnapshotIdentifier: string;
    public /*out*/ readonly encrypted: boolean;
    public /*out*/ readonly engine: string;
    public /*out*/ readonly engineVersion: string;
    public /*out*/ readonly iops: number;
    public /*out*/ readonly kmsKeyId: string;
    public /*out*/ readonly licenseModel: string;
    public /*out*/ readonly optionGroupName: string;
    public /*out*/ readonly port: number;
    public /*out*/ readonly snapshotType: string;
    public /*out*/ readonly sourceDbSnapshotIdentifier: string;
    public /*out*/ readonly sourceRegion: string;
    public /*out*/ readonly status: string;
    public /*out*/ readonly storageType: string;
    public /*out*/ readonly vpcId: string;

    public static get(id: lumi.ID): Snapshot {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): Snapshot[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: SnapshotArgs) {
        super(name);
        if (lumirt.defaultIfComputed(args.dbInstanceIdentifier, "") === undefined) {
            throw new Error("Property argument 'dbInstanceIdentifier' is required, but was missing");
        }
        this.dbInstanceIdentifier = args.dbInstanceIdentifier;
        if (lumirt.defaultIfComputed(args.dbSnapshotIdentifier, "") === undefined) {
            throw new Error("Property argument 'dbSnapshotIdentifier' is required, but was missing");
        }
        this.dbSnapshotIdentifier = args.dbSnapshotIdentifier;
    }
}

export interface SnapshotArgs {
    readonly dbInstanceIdentifier: string;
    readonly dbSnapshotIdentifier: string;
}

