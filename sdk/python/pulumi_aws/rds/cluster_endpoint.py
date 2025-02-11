# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ClusterEndpoint(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of cluster
    """
    cluster_endpoint_identifier: pulumi.Output[str]
    """
    The identifier to use for the new endpoint. This parameter is stored as a lowercase string.
    """
    cluster_identifier: pulumi.Output[str]
    """
    The cluster identifier.
    """
    custom_endpoint_type: pulumi.Output[str]
    """
    The type of the endpoint. One of: READER , ANY .
    """
    endpoint: pulumi.Output[str]
    """
    A custom endpoint for the Aurora cluster
    """
    excluded_members: pulumi.Output[list]
    """
    List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with `static_members`.
    """
    static_members: pulumi.Output[list]
    """
    List of DB instance identifiers that are part of the custom endpoint group. Conflicts with `excluded_members`.
    """
    def __init__(__self__, resource_name, opts=None, cluster_endpoint_identifier=None, cluster_identifier=None, custom_endpoint_type=None, excluded_members=None, static_members=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a RDS Aurora Cluster Endpoint.
        You can refer to the [User Guide][1].
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_endpoint_identifier: The identifier to use for the new endpoint. This parameter is stored as a lowercase string.
        :param pulumi.Input[str] cluster_identifier: The cluster identifier.
        :param pulumi.Input[str] custom_endpoint_type: The type of the endpoint. One of: READER , ANY .
        :param pulumi.Input[list] excluded_members: List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with `static_members`.
        :param pulumi.Input[list] static_members: List of DB instance identifiers that are part of the custom endpoint group. Conflicts with `excluded_members`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/rds_cluster_endpoint.html.markdown.
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

            if cluster_endpoint_identifier is None:
                raise TypeError("Missing required property 'cluster_endpoint_identifier'")
            __props__['cluster_endpoint_identifier'] = cluster_endpoint_identifier
            if cluster_identifier is None:
                raise TypeError("Missing required property 'cluster_identifier'")
            __props__['cluster_identifier'] = cluster_identifier
            if custom_endpoint_type is None:
                raise TypeError("Missing required property 'custom_endpoint_type'")
            __props__['custom_endpoint_type'] = custom_endpoint_type
            __props__['excluded_members'] = excluded_members
            __props__['static_members'] = static_members
            __props__['arn'] = None
            __props__['endpoint'] = None
        super(ClusterEndpoint, __self__).__init__(
            'aws:rds/clusterEndpoint:ClusterEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, cluster_endpoint_identifier=None, cluster_identifier=None, custom_endpoint_type=None, endpoint=None, excluded_members=None, static_members=None):
        """
        Get an existing ClusterEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of cluster
        :param pulumi.Input[str] cluster_endpoint_identifier: The identifier to use for the new endpoint. This parameter is stored as a lowercase string.
        :param pulumi.Input[str] cluster_identifier: The cluster identifier.
        :param pulumi.Input[str] custom_endpoint_type: The type of the endpoint. One of: READER , ANY .
        :param pulumi.Input[str] endpoint: A custom endpoint for the Aurora cluster
        :param pulumi.Input[list] excluded_members: List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with `static_members`.
        :param pulumi.Input[list] static_members: List of DB instance identifiers that are part of the custom endpoint group. Conflicts with `excluded_members`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/rds_cluster_endpoint.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["arn"] = arn
        __props__["cluster_endpoint_identifier"] = cluster_endpoint_identifier
        __props__["cluster_identifier"] = cluster_identifier
        __props__["custom_endpoint_type"] = custom_endpoint_type
        __props__["endpoint"] = endpoint
        __props__["excluded_members"] = excluded_members
        __props__["static_members"] = static_members
        return ClusterEndpoint(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

