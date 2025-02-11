// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a S3 bucket resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/s3_bucket.html.markdown.
type Bucket struct {
	s *pulumi.ResourceState
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOpt) (*Bucket, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accelerationStatus"] = nil
		inputs["acl"] = nil
		inputs["arn"] = nil
		inputs["bucket"] = nil
		inputs["bucketPrefix"] = nil
		inputs["corsRules"] = nil
		inputs["forceDestroy"] = nil
		inputs["hostedZoneId"] = nil
		inputs["lifecycleRules"] = nil
		inputs["loggings"] = nil
		inputs["objectLockConfiguration"] = nil
		inputs["policy"] = nil
		inputs["region"] = nil
		inputs["replicationConfiguration"] = nil
		inputs["requestPayer"] = nil
		inputs["serverSideEncryptionConfiguration"] = nil
		inputs["tags"] = nil
		inputs["versioning"] = nil
		inputs["website"] = nil
		inputs["websiteDomain"] = nil
		inputs["websiteEndpoint"] = nil
	} else {
		inputs["accelerationStatus"] = args.AccelerationStatus
		inputs["acl"] = args.Acl
		inputs["arn"] = args.Arn
		inputs["bucket"] = args.Bucket
		inputs["bucketPrefix"] = args.BucketPrefix
		inputs["corsRules"] = args.CorsRules
		inputs["forceDestroy"] = args.ForceDestroy
		inputs["hostedZoneId"] = args.HostedZoneId
		inputs["lifecycleRules"] = args.LifecycleRules
		inputs["loggings"] = args.Loggings
		inputs["objectLockConfiguration"] = args.ObjectLockConfiguration
		inputs["policy"] = args.Policy
		inputs["region"] = args.Region
		inputs["replicationConfiguration"] = args.ReplicationConfiguration
		inputs["requestPayer"] = args.RequestPayer
		inputs["serverSideEncryptionConfiguration"] = args.ServerSideEncryptionConfiguration
		inputs["tags"] = args.Tags
		inputs["versioning"] = args.Versioning
		inputs["website"] = args.Website
		inputs["websiteDomain"] = args.WebsiteDomain
		inputs["websiteEndpoint"] = args.WebsiteEndpoint
	}
	inputs["bucketDomainName"] = nil
	inputs["bucketRegionalDomainName"] = nil
	s, err := ctx.RegisterResource("aws:s3/bucket:Bucket", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Bucket{s: s}, nil
}

// GetBucket gets an existing Bucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucket(ctx *pulumi.Context,
	name string, id pulumi.ID, state *BucketState, opts ...pulumi.ResourceOpt) (*Bucket, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accelerationStatus"] = state.AccelerationStatus
		inputs["acl"] = state.Acl
		inputs["arn"] = state.Arn
		inputs["bucket"] = state.Bucket
		inputs["bucketDomainName"] = state.BucketDomainName
		inputs["bucketPrefix"] = state.BucketPrefix
		inputs["bucketRegionalDomainName"] = state.BucketRegionalDomainName
		inputs["corsRules"] = state.CorsRules
		inputs["forceDestroy"] = state.ForceDestroy
		inputs["hostedZoneId"] = state.HostedZoneId
		inputs["lifecycleRules"] = state.LifecycleRules
		inputs["loggings"] = state.Loggings
		inputs["objectLockConfiguration"] = state.ObjectLockConfiguration
		inputs["policy"] = state.Policy
		inputs["region"] = state.Region
		inputs["replicationConfiguration"] = state.ReplicationConfiguration
		inputs["requestPayer"] = state.RequestPayer
		inputs["serverSideEncryptionConfiguration"] = state.ServerSideEncryptionConfiguration
		inputs["tags"] = state.Tags
		inputs["versioning"] = state.Versioning
		inputs["website"] = state.Website
		inputs["websiteDomain"] = state.WebsiteDomain
		inputs["websiteEndpoint"] = state.WebsiteEndpoint
	}
	s, err := ctx.ReadResource("aws:s3/bucket:Bucket", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Bucket{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Bucket) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Bucket) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
func (r *Bucket) AccelerationStatus() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["accelerationStatus"])
}

// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".
func (r *Bucket) Acl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["acl"])
}

// The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
func (r *Bucket) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// The ARN of the S3 bucket where you want Amazon S3 to store replicas of the object identified by the rule.
func (r *Bucket) Bucket() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["bucket"])
}

// The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
func (r *Bucket) BucketDomainName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["bucketDomainName"])
}

// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
func (r *Bucket) BucketPrefix() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["bucketPrefix"])
}

// The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
func (r *Bucket) BucketRegionalDomainName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["bucketRegionalDomainName"])
}

// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
func (r *Bucket) CorsRules() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["corsRules"])
}

// A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
func (r *Bucket) ForceDestroy() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["forceDestroy"])
}

// The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
func (r *Bucket) HostedZoneId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["hostedZoneId"])
}

// A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
func (r *Bucket) LifecycleRules() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["lifecycleRules"])
}

// A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
func (r *Bucket) Loggings() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["loggings"])
}

// A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
func (r *Bucket) ObjectLockConfiguration() pulumi.Output {
	return r.s.State["objectLockConfiguration"]
}

// A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), this provider may view the policy as constantly changing in a deployment. In this case, please make sure you use the verbose/specific version of the policy.
func (r *Bucket) Policy() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["policy"])
}

// If specified, the AWS region this bucket should reside in. Otherwise, the region used by the callee.
func (r *Bucket) Region() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["region"])
}

// A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
func (r *Bucket) ReplicationConfiguration() pulumi.Output {
	return r.s.State["replicationConfiguration"]
}

// Specifies who should bear the cost of Amazon S3 data transfer.
// Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
// the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
// developer guide for more information.
func (r *Bucket) RequestPayer() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["requestPayer"])
}

// A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
func (r *Bucket) ServerSideEncryptionConfiguration() pulumi.Output {
	return r.s.State["serverSideEncryptionConfiguration"]
}

// A mapping of tags that identifies subset of objects to which the rule applies.
// The rule applies only to objects having all the tags in its tagset.
func (r *Bucket) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
func (r *Bucket) Versioning() pulumi.Output {
	return r.s.State["versioning"]
}

// A website object (documented below).
func (r *Bucket) Website() pulumi.Output {
	return r.s.State["website"]
}

// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
func (r *Bucket) WebsiteDomain() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["websiteDomain"])
}

// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
func (r *Bucket) WebsiteEndpoint() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["websiteEndpoint"])
}

// Input properties used for looking up and filtering Bucket resources.
type BucketState struct {
	// Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
	AccelerationStatus interface{}
	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".
	Acl interface{}
	// The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
	Arn interface{}
	// The ARN of the S3 bucket where you want Amazon S3 to store replicas of the object identified by the rule.
	Bucket interface{}
	// The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
	BucketDomainName interface{}
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix interface{}
	// The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
	BucketRegionalDomainName interface{}
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules interface{}
	// A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy interface{}
	// The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
	HostedZoneId interface{}
	// A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
	LifecycleRules interface{}
	// A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
	Loggings interface{}
	// A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
	ObjectLockConfiguration interface{}
	// A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), this provider may view the policy as constantly changing in a deployment. In this case, please make sure you use the verbose/specific version of the policy.
	Policy interface{}
	// If specified, the AWS region this bucket should reside in. Otherwise, the region used by the callee.
	Region interface{}
	// A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
	ReplicationConfiguration interface{}
	// Specifies who should bear the cost of Amazon S3 data transfer.
	// Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
	// the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
	// developer guide for more information.
	RequestPayer interface{}
	// A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
	ServerSideEncryptionConfiguration interface{}
	// A mapping of tags that identifies subset of objects to which the rule applies.
	// The rule applies only to objects having all the tags in its tagset.
	Tags interface{}
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning interface{}
	// A website object (documented below).
	Website interface{}
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain interface{}
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint interface{}
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
	AccelerationStatus interface{}
	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".
	Acl interface{}
	// The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
	Arn interface{}
	// The ARN of the S3 bucket where you want Amazon S3 to store replicas of the object identified by the rule.
	Bucket interface{}
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix interface{}
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules interface{}
	// A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy interface{}
	// The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
	HostedZoneId interface{}
	// A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
	LifecycleRules interface{}
	// A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
	Loggings interface{}
	// A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
	ObjectLockConfiguration interface{}
	// A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), this provider may view the policy as constantly changing in a deployment. In this case, please make sure you use the verbose/specific version of the policy.
	Policy interface{}
	// If specified, the AWS region this bucket should reside in. Otherwise, the region used by the callee.
	Region interface{}
	// A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
	ReplicationConfiguration interface{}
	// Specifies who should bear the cost of Amazon S3 data transfer.
	// Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
	// the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
	// developer guide for more information.
	RequestPayer interface{}
	// A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
	ServerSideEncryptionConfiguration interface{}
	// A mapping of tags that identifies subset of objects to which the rule applies.
	// The rule applies only to objects having all the tags in its tagset.
	Tags interface{}
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning interface{}
	// A website object (documented below).
	Website interface{}
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain interface{}
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint interface{}
}
