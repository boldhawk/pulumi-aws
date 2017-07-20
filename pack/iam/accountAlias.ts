// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class AccountAlias extends lumi.NamedResource implements AccountAliasArgs {
    public readonly accountAlias: string;

    public static get(id: lumi.ID): AccountAlias {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): AccountAlias[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: AccountAliasArgs) {
        super(name);
        if (lumirt.defaultIfComputed(args.accountAlias, "") === undefined) {
            throw new Error("Property argument 'accountAlias' is required, but was missing");
        }
        this.accountAlias = args.accountAlias;
    }
}

export interface AccountAliasArgs {
    readonly accountAlias: string;
}

