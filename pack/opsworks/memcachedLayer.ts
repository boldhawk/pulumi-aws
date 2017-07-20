// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class MemcachedLayer extends lumi.NamedResource implements MemcachedLayerArgs {
    public readonly allocatedMemory?: number;
    public readonly autoAssignElasticIps?: boolean;
    public readonly autoAssignPublicIps?: boolean;
    public readonly autoHealing?: boolean;
    public readonly customConfigureRecipes?: string[];
    public readonly customDeployRecipes?: string[];
    public readonly customInstanceProfileArn?: string;
    public readonly customJson?: string;
    public readonly customSecurityGroupIds?: string[];
    public readonly customSetupRecipes?: string[];
    public readonly customShutdownRecipes?: string[];
    public readonly customUndeployRecipes?: string[];
    public readonly drainElbOnShutdown?: boolean;
    public readonly ebsVolume?: { iops?: number, mountPoint: string, numberOfDisks: number, raidLevel?: string, size: number, type?: string }[];
    public readonly elasticLoadBalancer?: string;
    public /*out*/ readonly layerId: string;
    public readonly installUpdatesOnBoot?: boolean;
    public readonly instanceShutdownTimeout?: number;
    public readonly memcachedLayerName?: string;
    public readonly stackId: string;
    public readonly systemPackages?: string[];
    public readonly useEbsOptimizedInstances?: boolean;

    public static get(id: lumi.ID): MemcachedLayer {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): MemcachedLayer[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: MemcachedLayerArgs) {
        super(name);
        this.allocatedMemory = args.allocatedMemory;
        this.autoAssignElasticIps = args.autoAssignElasticIps;
        this.autoAssignPublicIps = args.autoAssignPublicIps;
        this.autoHealing = args.autoHealing;
        this.customConfigureRecipes = args.customConfigureRecipes;
        this.customDeployRecipes = args.customDeployRecipes;
        this.customInstanceProfileArn = args.customInstanceProfileArn;
        this.customJson = args.customJson;
        this.customSecurityGroupIds = args.customSecurityGroupIds;
        this.customSetupRecipes = args.customSetupRecipes;
        this.customShutdownRecipes = args.customShutdownRecipes;
        this.customUndeployRecipes = args.customUndeployRecipes;
        this.drainElbOnShutdown = args.drainElbOnShutdown;
        this.ebsVolume = args.ebsVolume;
        this.elasticLoadBalancer = args.elasticLoadBalancer;
        this.installUpdatesOnBoot = args.installUpdatesOnBoot;
        this.instanceShutdownTimeout = args.instanceShutdownTimeout;
        this.memcachedLayerName = args.memcachedLayerName;
        if (lumirt.defaultIfComputed(args.stackId, "") === undefined) {
            throw new Error("Property argument 'stackId' is required, but was missing");
        }
        this.stackId = args.stackId;
        this.systemPackages = args.systemPackages;
        this.useEbsOptimizedInstances = args.useEbsOptimizedInstances;
    }
}

export interface MemcachedLayerArgs {
    readonly allocatedMemory?: number;
    readonly autoAssignElasticIps?: boolean;
    readonly autoAssignPublicIps?: boolean;
    readonly autoHealing?: boolean;
    readonly customConfigureRecipes?: string[];
    readonly customDeployRecipes?: string[];
    readonly customInstanceProfileArn?: string;
    readonly customJson?: string;
    readonly customSecurityGroupIds?: string[];
    readonly customSetupRecipes?: string[];
    readonly customShutdownRecipes?: string[];
    readonly customUndeployRecipes?: string[];
    readonly drainElbOnShutdown?: boolean;
    readonly ebsVolume?: { iops?: number, mountPoint: string, numberOfDisks: number, raidLevel?: string, size: number, type?: string }[];
    readonly elasticLoadBalancer?: string;
    readonly installUpdatesOnBoot?: boolean;
    readonly instanceShutdownTimeout?: number;
    readonly memcachedLayerName?: string;
    readonly stackId: string;
    readonly systemPackages?: string[];
    readonly useEbsOptimizedInstances?: boolean;
}

