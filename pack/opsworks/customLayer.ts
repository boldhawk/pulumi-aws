// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class CustomLayer extends lumi.NamedResource implements CustomLayerArgs {
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
    public readonly customLayerName?: string;
    public readonly shortName: string;
    public readonly stackId: string;
    public readonly systemPackages?: string[];
    public readonly useEbsOptimizedInstances?: boolean;

    public static get(id: lumi.ID): CustomLayer {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): CustomLayer[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: CustomLayerArgs) {
        super(name);
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
        this.customLayerName = args.customLayerName;
        if (lumirt.defaultIfComputed(args.shortName, "") === undefined) {
            throw new Error("Property argument 'shortName' is required, but was missing");
        }
        this.shortName = args.shortName;
        if (lumirt.defaultIfComputed(args.stackId, "") === undefined) {
            throw new Error("Property argument 'stackId' is required, but was missing");
        }
        this.stackId = args.stackId;
        this.systemPackages = args.systemPackages;
        this.useEbsOptimizedInstances = args.useEbsOptimizedInstances;
    }
}

export interface CustomLayerArgs {
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
    readonly customLayerName?: string;
    readonly shortName: string;
    readonly stackId: string;
    readonly systemPackages?: string[];
    readonly useEbsOptimizedInstances?: boolean;
}

