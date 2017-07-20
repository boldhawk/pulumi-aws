// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class SecurityGroup extends lumi.NamedResource implements SecurityGroupArgs {
    public readonly description?: string;
    public readonly securityGroupName?: string;
    public readonly securityGroupNames: string[];

    public static get(id: lumi.ID): SecurityGroup {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): SecurityGroup[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: SecurityGroupArgs) {
        super(name);
        this.description = args.description;
        this.securityGroupName = args.securityGroupName;
        if (lumirt.defaultIfComputed(args.securityGroupNames, "") === undefined) {
            throw new Error("Property argument 'securityGroupNames' is required, but was missing");
        }
        this.securityGroupNames = args.securityGroupNames;
    }
}

export interface SecurityGroupArgs {
    readonly description?: string;
    readonly securityGroupName?: string;
    readonly securityGroupNames: string[];
}

