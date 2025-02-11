# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class UsagePlan(pulumi.CustomResource):
    api_stages: pulumi.Output[list]
    """
    The associated API stages of the usage plan.
    
      * `apiId` (`str`) - API Id of the associated API stage in a usage plan.
      * `stage` (`str`) - API stage name of the associated API stage in a usage plan.
    """
    description: pulumi.Output[str]
    """
    The description of a usage plan.
    """
    name: pulumi.Output[str]
    """
    The name of the usage plan.
    """
    product_code: pulumi.Output[str]
    """
    The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
    """
    quota_settings: pulumi.Output[dict]
    """
    The quota settings of the usage plan.
    
      * `limit` (`float`) - The maximum number of requests that can be made in a given time period.
      * `offset` (`float`) - The number of requests subtracted from the given limit in the initial time period.
      * `period` (`str`) - The time period in which the limit applies. Valid values are "DAY", "WEEK" or "MONTH".
    """
    throttle_settings: pulumi.Output[dict]
    """
    The throttling limits of the usage plan.
    
      * `burstLimit` (`float`) - The API request burst limit, the maximum rate limit over a time ranging from one to a few seconds, depending upon whether the underlying token bucket is at its full capacity.
      * `rateLimit` (`float`) - The API request steady-state rate limit.
    """
    def __init__(__self__, resource_name, opts=None, api_stages=None, description=None, name=None, product_code=None, quota_settings=None, throttle_settings=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an API Gateway Usage Plan.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] api_stages: The associated API stages of the usage plan.
        :param pulumi.Input[str] description: The description of a usage plan.
        :param pulumi.Input[str] name: The name of the usage plan.
        :param pulumi.Input[str] product_code: The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
        :param pulumi.Input[dict] quota_settings: The quota settings of the usage plan.
        :param pulumi.Input[dict] throttle_settings: The throttling limits of the usage plan.
        
        The **api_stages** object supports the following:
        
          * `apiId` (`pulumi.Input[str]`) - API Id of the associated API stage in a usage plan.
          * `stage` (`pulumi.Input[str]`) - API stage name of the associated API stage in a usage plan.
        
        The **quota_settings** object supports the following:
        
          * `limit` (`pulumi.Input[float]`) - The maximum number of requests that can be made in a given time period.
          * `offset` (`pulumi.Input[float]`) - The number of requests subtracted from the given limit in the initial time period.
          * `period` (`pulumi.Input[str]`) - The time period in which the limit applies. Valid values are "DAY", "WEEK" or "MONTH".
        
        The **throttle_settings** object supports the following:
        
          * `burstLimit` (`pulumi.Input[float]`) - The API request burst limit, the maximum rate limit over a time ranging from one to a few seconds, depending upon whether the underlying token bucket is at its full capacity.
          * `rateLimit` (`pulumi.Input[float]`) - The API request steady-state rate limit.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_usage_plan.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['api_stages'] = api_stages
            __props__['description'] = description
            __props__['name'] = name
            __props__['product_code'] = product_code
            __props__['quota_settings'] = quota_settings
            __props__['throttle_settings'] = throttle_settings
        super(UsagePlan, __self__).__init__(
            'aws:apigateway/usagePlan:UsagePlan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, api_stages=None, description=None, name=None, product_code=None, quota_settings=None, throttle_settings=None):
        """
        Get an existing UsagePlan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] api_stages: The associated API stages of the usage plan.
        :param pulumi.Input[str] description: The description of a usage plan.
        :param pulumi.Input[str] name: The name of the usage plan.
        :param pulumi.Input[str] product_code: The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
        :param pulumi.Input[dict] quota_settings: The quota settings of the usage plan.
        :param pulumi.Input[dict] throttle_settings: The throttling limits of the usage plan.
        
        The **api_stages** object supports the following:
        
          * `apiId` (`pulumi.Input[str]`) - API Id of the associated API stage in a usage plan.
          * `stage` (`pulumi.Input[str]`) - API stage name of the associated API stage in a usage plan.
        
        The **quota_settings** object supports the following:
        
          * `limit` (`pulumi.Input[float]`) - The maximum number of requests that can be made in a given time period.
          * `offset` (`pulumi.Input[float]`) - The number of requests subtracted from the given limit in the initial time period.
          * `period` (`pulumi.Input[str]`) - The time period in which the limit applies. Valid values are "DAY", "WEEK" or "MONTH".
        
        The **throttle_settings** object supports the following:
        
          * `burstLimit` (`pulumi.Input[float]`) - The API request burst limit, the maximum rate limit over a time ranging from one to a few seconds, depending upon whether the underlying token bucket is at its full capacity.
          * `rateLimit` (`pulumi.Input[float]`) - The API request steady-state rate limit.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_usage_plan.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["api_stages"] = api_stages
        __props__["description"] = description
        __props__["name"] = name
        __props__["product_code"] = product_code
        __props__["quota_settings"] = quota_settings
        __props__["throttle_settings"] = throttle_settings
        return UsagePlan(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

