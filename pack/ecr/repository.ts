// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Repository extends lumi.NamedResource implements RepositoryArgs {
    public /*out*/ readonly arn: string;
    public readonly repositoryName?: string;
    public /*out*/ readonly registryId: string;
    public /*out*/ readonly repositoryUrl: string;

    public static get(id: lumi.ID): Repository {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): Repository[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: RepositoryArgs) {
        super(name);
        this.repositoryName = args.repositoryName;
    }
}

export interface RepositoryArgs {
    readonly repositoryName?: string;
}

