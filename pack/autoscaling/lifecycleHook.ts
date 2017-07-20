// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class LifecycleHook extends lumi.NamedResource implements LifecycleHookArgs {
    public readonly autoscalingGroupName: string;
    public readonly defaultResult: string;
    public readonly heartbeatTimeout?: number;
    public readonly lifecycleTransition: string;
    public readonly lifecycleHookName?: string;
    public readonly notificationMetadata?: string;
    public readonly notificationTargetArn?: string;
    public readonly roleArn?: string;

    public static get(id: lumi.ID): LifecycleHook {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): LifecycleHook[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: LifecycleHookArgs) {
        super(name);
        if (lumirt.defaultIfComputed(args.autoscalingGroupName, "") === undefined) {
            throw new Error("Property argument 'autoscalingGroupName' is required, but was missing");
        }
        this.autoscalingGroupName = args.autoscalingGroupName;
        this.defaultResult = args.defaultResult;
        this.heartbeatTimeout = args.heartbeatTimeout;
        if (lumirt.defaultIfComputed(args.lifecycleTransition, "") === undefined) {
            throw new Error("Property argument 'lifecycleTransition' is required, but was missing");
        }
        this.lifecycleTransition = args.lifecycleTransition;
        this.lifecycleHookName = args.lifecycleHookName;
        this.notificationMetadata = args.notificationMetadata;
        this.notificationTargetArn = args.notificationTargetArn;
        this.roleArn = args.roleArn;
    }
}

export interface LifecycleHookArgs {
    readonly autoscalingGroupName: string;
    readonly defaultResult?: string;
    readonly heartbeatTimeout?: number;
    readonly lifecycleTransition: string;
    readonly lifecycleHookName?: string;
    readonly notificationMetadata?: string;
    readonly notificationTargetArn?: string;
    readonly roleArn?: string;
}

