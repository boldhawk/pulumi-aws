# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class GetInstanceProfileResult(object):
    """
    A collection of values returned by getInstanceProfile.
    """
    def __init__(__self__, arn=None, create_date=None, path=None, role_arn=None, role_id=None, role_name=None):
        if arn and not isinstance(arn, basestring):
            raise TypeError('Expected argument arn to be a basestring')
        __self__.arn = arn
        """
        The Amazon Resource Name (ARN) specifying the instance profile.
        """
        if create_date and not isinstance(create_date, basestring):
            raise TypeError('Expected argument create_date to be a basestring')
        __self__.create_date = create_date
        """
        The string representation of the date the instance profile
        was created.
        """
        if path and not isinstance(path, basestring):
            raise TypeError('Expected argument path to be a basestring')
        __self__.path = path
        """
        The path to the instance profile.
        """
        if role_arn and not isinstance(role_arn, basestring):
            raise TypeError('Expected argument role_arn to be a basestring')
        __self__.role_arn = role_arn
        """
        The role arn associated with this instance profile.
        """
        if role_id and not isinstance(role_id, basestring):
            raise TypeError('Expected argument role_id to be a basestring')
        __self__.role_id = role_id
        """
        The role id associated with this instance profile.
        """
        if role_name and not isinstance(role_name, basestring):
            raise TypeError('Expected argument role_name to be a basestring')
        __self__.role_name = role_name
        """
        The role name associated with this instance profile.
        """

def get_instance_profile(name=None):
    """
    This data source can be used to fetch information about a specific
    IAM instance profile. By using this data source, you can reference IAM
    instance profile properties without having to hard code ARNs as input.
    """
    __args__ = dict()

    __args__['name'] = name
    __ret__ = pulumi.runtime.invoke('aws:iam/getInstanceProfile:getInstanceProfile', __args__)

    return GetInstanceProfileResult(
        arn=__ret__.get('arn'),
        create_date=__ret__.get('createDate'),
        path=__ret__.get('path'),
        role_arn=__ret__.get('roleArn'),
        role_id=__ret__.get('roleId'),
        role_name=__ret__.get('roleName'))
