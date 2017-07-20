// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class FirehoseDeliveryStream extends lumi.NamedResource implements FirehoseDeliveryStreamArgs {
    public readonly arn: string;
    public readonly destination: string;
    public readonly destinationId: string;
    public readonly elasticsearchConfiguration?: { bufferingInterval?: number, bufferingSize?: number, cloudwatchLoggingOptions: { enabled?: boolean, logGroupName?: string, logStreamName?: string }[], domainArn: string, indexName: string, indexRotationPeriod?: string, retryDuration?: number, roleArn: string, s3BackupMode?: string, typeName?: string }[];
    public readonly firehoseDeliveryStreamName?: string;
    public readonly redshiftConfiguration?: { cloudwatchLoggingOptions: { enabled?: boolean, logGroupName?: string, logStreamName?: string }[], clusterJdbcurl: string, copyOptions?: string, dataTableColumns?: string, dataTableName: string, password: string, retryDuration?: number, roleArn: string, username: string }[];
    public readonly s3Configuration: { bucketArn: string, bufferInterval?: number, bufferSize?: number, cloudwatchLoggingOptions: { enabled?: boolean, logGroupName?: string, logStreamName?: string }[], compressionFormat?: string, kmsKeyArn?: string, prefix?: string, roleArn: string }[];
    public readonly versionId: string;

    public static get(id: lumi.ID): FirehoseDeliveryStream {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): FirehoseDeliveryStream[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: FirehoseDeliveryStreamArgs) {
        super(name);
        this.arn = args.arn;
        if (lumirt.defaultIfComputed(args.destination, "") === undefined) {
            throw new Error("Property argument 'destination' is required, but was missing");
        }
        this.destination = args.destination;
        this.destinationId = args.destinationId;
        this.elasticsearchConfiguration = args.elasticsearchConfiguration;
        this.firehoseDeliveryStreamName = args.firehoseDeliveryStreamName;
        this.redshiftConfiguration = args.redshiftConfiguration;
        if (lumirt.defaultIfComputed(args.s3Configuration, "") === undefined) {
            throw new Error("Property argument 's3Configuration' is required, but was missing");
        }
        this.s3Configuration = args.s3Configuration;
        this.versionId = args.versionId;
    }
}

export interface FirehoseDeliveryStreamArgs {
    readonly arn?: string;
    readonly destination: string;
    readonly destinationId?: string;
    readonly elasticsearchConfiguration?: { bufferingInterval?: number, bufferingSize?: number, cloudwatchLoggingOptions: { enabled?: boolean, logGroupName?: string, logStreamName?: string }[], domainArn: string, indexName: string, indexRotationPeriod?: string, retryDuration?: number, roleArn: string, s3BackupMode?: string, typeName?: string }[];
    readonly firehoseDeliveryStreamName?: string;
    readonly redshiftConfiguration?: { cloudwatchLoggingOptions: { enabled?: boolean, logGroupName?: string, logStreamName?: string }[], clusterJdbcurl: string, copyOptions?: string, dataTableColumns?: string, dataTableName: string, password: string, retryDuration?: number, roleArn: string, username: string }[];
    readonly s3Configuration: { bucketArn: string, bufferInterval?: number, bufferSize?: number, cloudwatchLoggingOptions: { enabled?: boolean, logGroupName?: string, logStreamName?: string }[], compressionFormat?: string, kmsKeyArn?: string, prefix?: string, roleArn: string }[];
    readonly versionId?: string;
}

