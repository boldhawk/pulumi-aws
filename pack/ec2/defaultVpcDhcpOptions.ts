// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class DefaultVpcDhcpOptions extends lumi.NamedResource implements DefaultVpcDhcpOptionsArgs {
    public /*out*/ readonly domainName: string;
    public /*out*/ readonly domainNameServers: string;
    public readonly netbiosNameServers?: string[];
    public readonly netbiosNodeType?: string;
    public /*out*/ readonly ntpServers: string;
    public readonly tags?: {[key: string]: any};

    public static get(id: lumi.ID): DefaultVpcDhcpOptions {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): DefaultVpcDhcpOptions[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: DefaultVpcDhcpOptionsArgs) {
        super(name);
        this.netbiosNameServers = args.netbiosNameServers;
        this.netbiosNodeType = args.netbiosNodeType;
        this.tags = args.tags;
    }
}

export interface DefaultVpcDhcpOptionsArgs {
    readonly netbiosNameServers?: string[];
    readonly netbiosNodeType?: string;
    readonly tags?: {[key: string]: any};
}

