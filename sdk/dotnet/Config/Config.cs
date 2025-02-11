// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.Aws
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("aws");

        /// <summary>
        /// The access key for API operations. You can retrieve this from the 'Security &amp; Credentials' section of
        /// the AWS console.
        /// </summary>
        public static string? AccessKey { get; set; } = __config.Get("accessKey");

        public static ImmutableArray<string> AllowedAccountIds { get; set; } = __config.GetObject<ImmutableArray<string>>("allowedAccountIds");

        public static ConfigTypes.AssumeRole? AssumeRole { get; set; } = __config.GetObject<ConfigTypes.AssumeRole>("assumeRole");

        public static ImmutableArray<ConfigTypes.Endpoints> Endpoints { get; set; } = __config.GetObject<ImmutableArray<ConfigTypes.Endpoints>>("endpoints");

        public static ImmutableArray<string> ForbiddenAccountIds { get; set; } = __config.GetObject<ImmutableArray<string>>("forbiddenAccountIds");

        /// <summary>
        /// Resource tag key prefixes to ignore across all resources.
        /// </summary>
        public static ImmutableArray<string> IgnoreTagPrefixes { get; set; } = __config.GetObject<ImmutableArray<string>>("ignoreTagPrefixes");

        /// <summary>
        /// Resource tag keys to ignore across all resources.
        /// </summary>
        public static ImmutableArray<string> IgnoreTags { get; set; } = __config.GetObject<ImmutableArray<string>>("ignoreTags");

        /// <summary>
        /// Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`
        /// </summary>
        public static bool? Insecure { get; set; } = __config.GetBoolean("insecure");

        /// <summary>
        /// The maximum number of times an AWS API request is being executed. If the API request still fails, an error
        /// is thrown.
        /// </summary>
        public static int? MaxRetries { get; set; } = __config.GetInt32("maxRetries");

        /// <summary>
        /// The profile for API operations. If not set, the default profile created with `aws configure` will be used.
        /// </summary>
        public static string? Profile { get; set; } = __config.Get("profile") ?? Utilities.GetEnv("AWS_PROFILE");

        /// <summary>
        /// The region where AWS operations will take place. Examples are us-east-1, us-west-2, etc.
        /// </summary>
        public static string? Region { get; set; } = __config.Get("region") ?? Utilities.GetEnv("AWS_REGION", "AWS_DEFAULT_REGION");

        /// <summary>
        /// Set this to true to force the request to use path-style addressing, i.e.,
        /// http://s3.amazonaws.com/BUCKET/KEY. By default, the S3 client will use virtual hosted bucket addressing when
        /// possible (http://BUCKET.s3.amazonaws.com/KEY). Specific to the Amazon S3 service.
        /// </summary>
        public static bool? S3ForcePathStyle { get; set; } = __config.GetBoolean("s3ForcePathStyle");

        /// <summary>
        /// The secret key for API operations. You can retrieve this from the 'Security &amp; Credentials' section of
        /// the AWS console.
        /// </summary>
        public static string? SecretKey { get; set; } = __config.Get("secretKey");

        /// <summary>
        /// The path to the shared credentials file. If not set this defaults to ~/.aws/credentials.
        /// </summary>
        public static string? SharedCredentialsFile { get; set; } = __config.Get("sharedCredentialsFile");

        /// <summary>
        /// Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS
        /// available/implemented.
        /// </summary>
        public static bool? SkipCredentialsValidation { get; set; } = __config.GetBoolean("skipCredentialsValidation");

        /// <summary>
        /// Skip getting the supported EC2 platforms. Used by users that don't have ec2:DescribeAccountAttributes
        /// permissions.
        /// </summary>
        public static bool? SkipGetEc2Platforms { get; set; } = __config.GetBoolean("skipGetEc2Platforms");

        public static bool? SkipMetadataApiCheck { get; set; } = __config.GetBoolean("skipMetadataApiCheck");

        /// <summary>
        /// Skip static validation of region name. Used by users of alternative AWS-like APIs or users w/ access to
        /// regions that are not public (yet).
        /// </summary>
        public static bool? SkipRegionValidation { get; set; } = __config.GetBoolean("skipRegionValidation");

        /// <summary>
        /// Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or
        /// metadata API.
        /// </summary>
        public static bool? SkipRequestingAccountId { get; set; } = __config.GetBoolean("skipRequestingAccountId");

        /// <summary>
        /// session token. A session token is only required if you are using temporary security credentials.
        /// </summary>
        public static string? Token { get; set; } = __config.Get("token");

    }
    namespace ConfigTypes
    {

    public class AssumeRole
    {
        public string? ExternalId { get; set; }
        public string? Policy { get; set; }
        public string? RoleArn { get; set; }
        public string? SessionName { get; set; }
    }

    public class Endpoints
    {
        public string? Acm { get; set; }
        public string? Acmpca { get; set; }
        public string? Amplify { get; set; }
        public string? Apigateway { get; set; }
        public string? Applicationautoscaling { get; set; }
        public string? Applicationinsights { get; set; }
        public string? Appmesh { get; set; }
        public string? Appstream { get; set; }
        public string? Appsync { get; set; }
        public string? Athena { get; set; }
        public string? Autoscaling { get; set; }
        public string? Autoscalingplans { get; set; }
        public string? Backup { get; set; }
        public string? Batch { get; set; }
        public string? Budgets { get; set; }
        public string? Cloud9 { get; set; }
        public string? Cloudformation { get; set; }
        public string? Cloudfront { get; set; }
        public string? Cloudhsm { get; set; }
        public string? Cloudsearch { get; set; }
        public string? Cloudtrail { get; set; }
        public string? Cloudwatch { get; set; }
        public string? Cloudwatchevents { get; set; }
        public string? Cloudwatchlogs { get; set; }
        public string? Codebuild { get; set; }
        public string? Codecommit { get; set; }
        public string? Codedeploy { get; set; }
        public string? Codepipeline { get; set; }
        public string? Cognitoidentity { get; set; }
        public string? Cognitoidp { get; set; }
        public string? Configservice { get; set; }
        public string? Cur { get; set; }
        public string? Datapipeline { get; set; }
        public string? Datasync { get; set; }
        public string? Dax { get; set; }
        public string? Devicefarm { get; set; }
        public string? Directconnect { get; set; }
        public string? Dlm { get; set; }
        public string? Dms { get; set; }
        public string? Docdb { get; set; }
        public string? Ds { get; set; }
        public string? Dynamodb { get; set; }
        public string? Ec2 { get; set; }
        public string? Ecr { get; set; }
        public string? Ecs { get; set; }
        public string? Efs { get; set; }
        public string? Eks { get; set; }
        public string? Elasticache { get; set; }
        public string? Elasticbeanstalk { get; set; }
        public string? Elastictranscoder { get; set; }
        public string? Elb { get; set; }
        public string? Emr { get; set; }
        public string? Es { get; set; }
        public string? Firehose { get; set; }
        public string? Fms { get; set; }
        public string? Forecast { get; set; }
        public string? Fsx { get; set; }
        public string? Gamelift { get; set; }
        public string? Glacier { get; set; }
        public string? Globalaccelerator { get; set; }
        public string? Glue { get; set; }
        public string? Greengrass { get; set; }
        public string? Guardduty { get; set; }
        public string? Iam { get; set; }
        public string? Inspector { get; set; }
        public string? Iot { get; set; }
        public string? Iotanalytics { get; set; }
        public string? Iotevents { get; set; }
        public string? Kafka { get; set; }
        public string? Kinesis { get; set; }
        public string? KinesisAnalytics { get; set; }
        public string? Kinesisanalytics { get; set; }
        public string? Kinesisvideo { get; set; }
        public string? Kms { get; set; }
        public string? Lakeformation { get; set; }
        public string? Lambda { get; set; }
        public string? Lexmodels { get; set; }
        public string? Licensemanager { get; set; }
        public string? Lightsail { get; set; }
        public string? Macie { get; set; }
        public string? Managedblockchain { get; set; }
        public string? Mediaconnect { get; set; }
        public string? Mediaconvert { get; set; }
        public string? Medialive { get; set; }
        public string? Mediapackage { get; set; }
        public string? Mediastore { get; set; }
        public string? Mediastoredata { get; set; }
        public string? Mq { get; set; }
        public string? Neptune { get; set; }
        public string? Opsworks { get; set; }
        public string? Organizations { get; set; }
        public string? Personalize { get; set; }
        public string? Pinpoint { get; set; }
        public string? Pricing { get; set; }
        public string? Qldb { get; set; }
        public string? Quicksight { get; set; }
        public string? R53 { get; set; }
        public string? Ram { get; set; }
        public string? Rds { get; set; }
        public string? Redshift { get; set; }
        public string? Resourcegroups { get; set; }
        public string? Route53 { get; set; }
        public string? Route53resolver { get; set; }
        public string? S3 { get; set; }
        public string? S3control { get; set; }
        public string? Sagemaker { get; set; }
        public string? Sdb { get; set; }
        public string? Secretsmanager { get; set; }
        public string? Securityhub { get; set; }
        public string? Serverlessrepo { get; set; }
        public string? Servicecatalog { get; set; }
        public string? Servicediscovery { get; set; }
        public string? Servicequotas { get; set; }
        public string? Ses { get; set; }
        public string? Shield { get; set; }
        public string? Sns { get; set; }
        public string? Sqs { get; set; }
        public string? Ssm { get; set; }
        public string? Stepfunctions { get; set; }
        public string? Storagegateway { get; set; }
        public string? Sts { get; set; }
        public string? Swf { get; set; }
        public string? Transfer { get; set; }
        public string? Waf { get; set; }
        public string? Wafregional { get; set; }
        public string? Worklink { get; set; }
        public string? Workspaces { get; set; }
        public string? Xray { get; set; }
    }
    }
}
