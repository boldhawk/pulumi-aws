// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cur

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get information on an AWS Cost and Usage Report Definition.
// 
// > *NOTE:* The AWS Cost and Usage Report service is only available in `us-east-1` currently.
// 
// > *NOTE:* If AWS Organizations is enabled, only the master account can use this resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/cur_report_definition.html.markdown.
func LookupReportDefinition(ctx *pulumi.Context, args *GetReportDefinitionArgs) (*GetReportDefinitionResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["reportName"] = args.ReportName
	}
	outputs, err := ctx.Invoke("aws:cur/getReportDefinition:getReportDefinition", inputs)
	if err != nil {
		return nil, err
	}
	return &GetReportDefinitionResult{
		AdditionalArtifacts: outputs["additionalArtifacts"],
		AdditionalSchemaElements: outputs["additionalSchemaElements"],
		Compression: outputs["compression"],
		Format: outputs["format"],
		ReportName: outputs["reportName"],
		S3Bucket: outputs["s3Bucket"],
		S3Prefix: outputs["s3Prefix"],
		S3Region: outputs["s3Region"],
		TimeUnit: outputs["timeUnit"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getReportDefinition.
type GetReportDefinitionArgs struct {
	// The name of the report definition to match.
	ReportName interface{}
}

// A collection of values returned by getReportDefinition.
type GetReportDefinitionResult struct {
	// A list of additional artifacts.
	AdditionalArtifacts interface{}
	// A list of schema elements.
	AdditionalSchemaElements interface{}
	// Preferred format for report.
	Compression interface{}
	// Preferred compression format for report.
	Format interface{}
	ReportName interface{}
	// Name of customer S3 bucket.
	S3Bucket interface{}
	// Preferred report path prefix.
	S3Prefix interface{}
	// Region of customer S3 bucket.
	S3Region interface{}
	// The frequency on which report data are measured and displayed.
	TimeUnit interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
