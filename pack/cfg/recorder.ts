// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Recorder extends lumi.NamedResource implements RecorderArgs {
    public readonly recorderName?: string;
    public readonly recordingGroup: { allSupported?: boolean, includeGlobalResourceTypes?: boolean, resourceTypes?: string[] }[];
    public readonly roleArn: string;

    public static get(id: lumi.ID): Recorder {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): Recorder[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: RecorderArgs) {
        super(name);
        this.recorderName = args.recorderName;
        this.recordingGroup = args.recordingGroup;
        if (lumirt.defaultIfComputed(args.roleArn, "") === undefined) {
            throw new Error("Property argument 'roleArn' is required, but was missing");
        }
        this.roleArn = args.roleArn;
    }
}

export interface RecorderArgs {
    readonly recorderName?: string;
    readonly recordingGroup?: { allSupported?: boolean, includeGlobalResourceTypes?: boolean, resourceTypes?: string[] }[];
    readonly roleArn: string;
}

