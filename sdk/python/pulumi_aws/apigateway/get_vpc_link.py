# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetVpcLinkResult:
    """
    A collection of values returned by getVpcLink.
    """
    def __init__(__self__, id=None, name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        Set to the ID of the found API Gateway VPC Link.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
class AwaitableGetVpcLinkResult(GetVpcLinkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcLinkResult(
            id=self.id,
            name=self.name)

def get_vpc_link(name=None,opts=None):
    """
    Use this data source to get the id of a VPC Link in
    API Gateway. To fetch the VPC Link you must provide a name to match against. 
    As there is no unique name constraint on API Gateway VPC Links this data source will 
    error if there is more than one match.
    
    :param str name: The name of the API Gateway VPC Link to look up. If no API Gateway VPC Link is found with this name, an error will be returned. 
           If multiple API Gateway VPC Links are found with this name, an error will be returned.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/api_gateway_vpc_link.html.markdown.
    """
    __args__ = dict()

    __args__['name'] = name
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:apigateway/getVpcLink:getVpcLink', __args__, opts=opts).value

    return AwaitableGetVpcLinkResult(
        id=__ret__.get('id'),
        name=__ret__.get('name'))
