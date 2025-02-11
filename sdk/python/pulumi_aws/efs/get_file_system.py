# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetFileSystemResult:
    """
    A collection of values returned by getFileSystem.
    """
    def __init__(__self__, arn=None, creation_token=None, dns_name=None, encrypted=None, file_system_id=None, kms_key_id=None, performance_mode=None, tags=None, id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        Amazon Resource Name of the file system.
        """
        if creation_token and not isinstance(creation_token, str):
            raise TypeError("Expected argument 'creation_token' to be a str")
        __self__.creation_token = creation_token
        if dns_name and not isinstance(dns_name, str):
            raise TypeError("Expected argument 'dns_name' to be a str")
        __self__.dns_name = dns_name
        """
        The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
        """
        if encrypted and not isinstance(encrypted, bool):
            raise TypeError("Expected argument 'encrypted' to be a bool")
        __self__.encrypted = encrypted
        """
        Whether EFS is encrypted.
        """
        if file_system_id and not isinstance(file_system_id, str):
            raise TypeError("Expected argument 'file_system_id' to be a str")
        __self__.file_system_id = file_system_id
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError("Expected argument 'kms_key_id' to be a str")
        __self__.kms_key_id = kms_key_id
        """
        The ARN for the KMS encryption key.
        """
        if performance_mode and not isinstance(performance_mode, str):
            raise TypeError("Expected argument 'performance_mode' to be a str")
        __self__.performance_mode = performance_mode
        """
        The PerformanceMode of the file system.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        The list of tags assigned to the file system.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetFileSystemResult(GetFileSystemResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFileSystemResult(
            arn=self.arn,
            creation_token=self.creation_token,
            dns_name=self.dns_name,
            encrypted=self.encrypted,
            file_system_id=self.file_system_id,
            kms_key_id=self.kms_key_id,
            performance_mode=self.performance_mode,
            tags=self.tags,
            id=self.id)

def get_file_system(creation_token=None,file_system_id=None,tags=None,opts=None):
    """
    Provides information about an Elastic File System (EFS).
    
    :param str creation_token: Restricts the list to the file system with this creation token.
    :param str file_system_id: The ID that identifies the file system (e.g. fs-ccfc0d65).

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/efs_file_system.html.markdown.
    """
    __args__ = dict()

    __args__['creationToken'] = creation_token
    __args__['fileSystemId'] = file_system_id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:efs/getFileSystem:getFileSystem', __args__, opts=opts).value

    return AwaitableGetFileSystemResult(
        arn=__ret__.get('arn'),
        creation_token=__ret__.get('creationToken'),
        dns_name=__ret__.get('dnsName'),
        encrypted=__ret__.get('encrypted'),
        file_system_id=__ret__.get('fileSystemId'),
        kms_key_id=__ret__.get('kmsKeyId'),
        performance_mode=__ret__.get('performanceMode'),
        tags=__ret__.get('tags'),
        id=__ret__.get('id'))
