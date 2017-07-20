// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class UserProfile extends lumi.NamedResource implements UserProfileArgs {
    public readonly allowSelfManagement?: boolean;
    public /*out*/ readonly profileId: string;
    public readonly sshPublicKey?: string;
    public readonly sshUsername: string;
    public readonly userArn: string;

    public static get(id: lumi.ID): UserProfile {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): UserProfile[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: UserProfileArgs) {
        super(name);
        this.allowSelfManagement = args.allowSelfManagement;
        this.sshPublicKey = args.sshPublicKey;
        if (lumirt.defaultIfComputed(args.sshUsername, "") === undefined) {
            throw new Error("Property argument 'sshUsername' is required, but was missing");
        }
        this.sshUsername = args.sshUsername;
        if (lumirt.defaultIfComputed(args.userArn, "") === undefined) {
            throw new Error("Property argument 'userArn' is required, but was missing");
        }
        this.userArn = args.userArn;
    }
}

export interface UserProfileArgs {
    readonly allowSelfManagement?: boolean;
    readonly sshPublicKey?: string;
    readonly sshUsername: string;
    readonly userArn: string;
}

