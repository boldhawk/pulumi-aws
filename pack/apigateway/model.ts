// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

import {RestApi} from "./restApi";

export class Model extends lumi.NamedResource implements ModelArgs {
    public readonly contentType: string;
    public readonly description?: string;
    public readonly modelName?: string;
    public readonly restApi: RestApi;
    public readonly schema?: string;

    public static get(id: lumi.ID): Model {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): Model[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: ModelArgs) {
        super(name);
        if (lumirt.defaultIfComputed(args.contentType, "") === undefined) {
            throw new Error("Property argument 'contentType' is required, but was missing");
        }
        this.contentType = args.contentType;
        this.description = args.description;
        this.modelName = args.modelName;
        if (lumirt.defaultIfComputed(args.restApi, "") === undefined) {
            throw new Error("Property argument 'restApi' is required, but was missing");
        }
        this.restApi = args.restApi;
        this.schema = args.schema;
    }
}

export interface ModelArgs {
    readonly contentType: string;
    readonly description?: string;
    readonly modelName?: string;
    readonly restApi: RestApi;
    readonly schema?: string;
}

