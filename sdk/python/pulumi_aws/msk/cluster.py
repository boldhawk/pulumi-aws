# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Cluster(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
    """
    bootstrap_brokers: pulumi.Output[str]
    """
    A comma separated list of one or more hostname:port pairs of kafka brokers suitable to boostrap connectivity to the kafka cluster. Only contains value if `client_broker` encryption in transit is set to `PLAINTEXT` or `TLS_PLAINTEXT`.
    """
    bootstrap_brokers_tls: pulumi.Output[str]
    """
    A comma separated list of one or more DNS names (or IPs) and TLS port pairs kafka brokers suitable to boostrap connectivity to the kafka cluster. Only contains value if `client_broker` encryption in transit is set to `TLS_PLAINTEXT` or `TLS`.
    """
    broker_node_group_info: pulumi.Output[dict]
    """
    Configuration block for the broker nodes of the Kafka cluster.
    
      * `az_distribution` (`str`) - The distribution of broker nodes across availability zones ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-model-brokerazdistribution)). Currently the only valid value is `DEFAULT`.
      * `client_subnets` (`list`) - A list of subnets to connect to in client VPC ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-prop-brokernodegroupinfo-clientsubnets)).
      * `ebs_volume_size` (`float`) - The size in GiB of the EBS volume for the data drive on each broker node.
      * `instance_type` (`str`) - Specify the instance type to use for the kafka brokers. e.g. kafka.m5.large. ([Pricing info](https://aws.amazon.com/msk/pricing/))
      * `security_groups` (`list`) - A list of the security groups to associate with the elastic network interfaces to control who can communicate with the cluster.
    """
    client_authentication: pulumi.Output[dict]
    """
    Configuration block for specifying a client authentication. See below.
    
      * `tls` (`dict`) - Configuration block for specifying TLS client authentication. See below.
    
        * `certificate_authority_arns` (`list`) - List of ACM Certificate Authority Amazon Resource Names (ARNs).
    """
    cluster_name: pulumi.Output[str]
    """
    Name of the MSK cluster.
    """
    configuration_info: pulumi.Output[dict]
    """
    Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
    
      * `arn` (`str`) - Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
      * `revision` (`float`) - Revision of the MSK Configuration to use in the cluster.
    """
    current_version: pulumi.Output[str]
    """
    Current version of the MSK Cluster used for updates, e.g. `K13V1IB3VIYZZH`
    * `encryption_info.0.encryption_at_rest_kms_key_arn` - The ARN of the KMS key used for encryption at rest of the broker data volumes.
    """
    encryption_info: pulumi.Output[dict]
    """
    Configuration block for specifying encryption. See below.
    
      * `encryption_at_rest_kms_key_arn` (`str`) - You may specify a KMS key short ID or ARN (it will always output an ARN) to use for encrypting your data at rest.  If no key is specified, an AWS managed KMS ('aws/msk' managed service) key will be used for encrypting the data at rest.
      * `encryption_in_transit` (`dict`) - Configuration block to specify encryption in transit. See below.
    
        * `client_broker` (`str`) - Encryption setting for data in transit between clients and brokers. Valid values: `TLS`, `TLS_PLAINTEXT`, and `PLAINTEXT`. Default value: `TLS_PLAINTEXT`.
        * `in_cluster` (`bool`) - Whether data communication among broker nodes is encrypted. Default value: `true`.
    """
    enhanced_monitoring: pulumi.Output[str]
    """
    Specify the desired enhanced MSK CloudWatch monitoring level.  See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
    """
    kafka_version: pulumi.Output[str]
    """
    Specify the desired Kafka software version.
    """
    number_of_broker_nodes: pulumi.Output[float]
    """
    The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource
    """
    zookeeper_connect_string: pulumi.Output[str]
    """
    A comma separated list of one or more IP:port pairs to use to connect to the Apache Zookeeper cluster.
    """
    def __init__(__self__, resource_name, opts=None, broker_node_group_info=None, client_authentication=None, cluster_name=None, configuration_info=None, encryption_info=None, enhanced_monitoring=None, kafka_version=None, number_of_broker_nodes=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages AWS Managed Streaming for Kafka cluster
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] broker_node_group_info: Configuration block for the broker nodes of the Kafka cluster.
        :param pulumi.Input[dict] client_authentication: Configuration block for specifying a client authentication. See below.
        :param pulumi.Input[str] cluster_name: Name of the MSK cluster.
        :param pulumi.Input[dict] configuration_info: Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
        :param pulumi.Input[dict] encryption_info: Configuration block for specifying encryption. See below.
        :param pulumi.Input[str] enhanced_monitoring: Specify the desired enhanced MSK CloudWatch monitoring level.  See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
        :param pulumi.Input[str] kafka_version: Specify the desired Kafka software version.
        :param pulumi.Input[float] number_of_broker_nodes: The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource
        
        The **client_authentication** object supports the following:
        
          * `tls` (`pulumi.Input[dict]`) - Configuration block for specifying TLS client authentication. See below.
        
            * `certificate_authority_arns` (`pulumi.Input[list]`) - List of ACM Certificate Authority Amazon Resource Names (ARNs).
        
        The **configuration_info** object supports the following:
        
          * `arn` (`pulumi.Input[str]`) - Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
          * `revision` (`pulumi.Input[float]`) - Revision of the MSK Configuration to use in the cluster.
        
        The **encryption_info** object supports the following:
        
          * `encryption_at_rest_kms_key_arn` (`pulumi.Input[str]`) - You may specify a KMS key short ID or ARN (it will always output an ARN) to use for encrypting your data at rest.  If no key is specified, an AWS managed KMS ('aws/msk' managed service) key will be used for encrypting the data at rest.
          * `encryption_in_transit` (`pulumi.Input[dict]`) - Configuration block to specify encryption in transit. See below.
        
            * `client_broker` (`pulumi.Input[str]`) - Encryption setting for data in transit between clients and brokers. Valid values: `TLS`, `TLS_PLAINTEXT`, and `PLAINTEXT`. Default value: `TLS_PLAINTEXT`.
            * `in_cluster` (`pulumi.Input[bool]`) - Whether data communication among broker nodes is encrypted. Default value: `true`.
        
        The **broker_node_group_info** object supports the following:
        
          * `az_distribution` (`pulumi.Input[str]`) - The distribution of broker nodes across availability zones ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-model-brokerazdistribution)). Currently the only valid value is `DEFAULT`.
          * `client_subnets` (`pulumi.Input[list]`) - A list of subnets to connect to in client VPC ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-prop-brokernodegroupinfo-clientsubnets)).
          * `ebs_volume_size` (`pulumi.Input[float]`) - The size in GiB of the EBS volume for the data drive on each broker node.
          * `instance_type` (`pulumi.Input[str]`) - Specify the instance type to use for the kafka brokers. e.g. kafka.m5.large. ([Pricing info](https://aws.amazon.com/msk/pricing/))
          * `security_groups` (`pulumi.Input[list]`) - A list of the security groups to associate with the elastic network interfaces to control who can communicate with the cluster.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/msk_cluster.html.markdown.
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

            if broker_node_group_info is None:
                raise TypeError("Missing required property 'broker_node_group_info'")
            __props__['broker_node_group_info'] = broker_node_group_info
            __props__['client_authentication'] = client_authentication
            if cluster_name is None:
                raise TypeError("Missing required property 'cluster_name'")
            __props__['cluster_name'] = cluster_name
            __props__['configuration_info'] = configuration_info
            __props__['encryption_info'] = encryption_info
            __props__['enhanced_monitoring'] = enhanced_monitoring
            if kafka_version is None:
                raise TypeError("Missing required property 'kafka_version'")
            __props__['kafka_version'] = kafka_version
            if number_of_broker_nodes is None:
                raise TypeError("Missing required property 'number_of_broker_nodes'")
            __props__['number_of_broker_nodes'] = number_of_broker_nodes
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['bootstrap_brokers'] = None
            __props__['bootstrap_brokers_tls'] = None
            __props__['current_version'] = None
            __props__['zookeeper_connect_string'] = None
        super(Cluster, __self__).__init__(
            'aws:msk/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, bootstrap_brokers=None, bootstrap_brokers_tls=None, broker_node_group_info=None, client_authentication=None, cluster_name=None, configuration_info=None, current_version=None, encryption_info=None, enhanced_monitoring=None, kafka_version=None, number_of_broker_nodes=None, tags=None, zookeeper_connect_string=None):
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
        :param pulumi.Input[str] bootstrap_brokers: A comma separated list of one or more hostname:port pairs of kafka brokers suitable to boostrap connectivity to the kafka cluster. Only contains value if `client_broker` encryption in transit is set to `PLAINTEXT` or `TLS_PLAINTEXT`.
        :param pulumi.Input[str] bootstrap_brokers_tls: A comma separated list of one or more DNS names (or IPs) and TLS port pairs kafka brokers suitable to boostrap connectivity to the kafka cluster. Only contains value if `client_broker` encryption in transit is set to `TLS_PLAINTEXT` or `TLS`.
        :param pulumi.Input[dict] broker_node_group_info: Configuration block for the broker nodes of the Kafka cluster.
        :param pulumi.Input[dict] client_authentication: Configuration block for specifying a client authentication. See below.
        :param pulumi.Input[str] cluster_name: Name of the MSK cluster.
        :param pulumi.Input[dict] configuration_info: Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
        :param pulumi.Input[str] current_version: Current version of the MSK Cluster used for updates, e.g. `K13V1IB3VIYZZH`
               * `encryption_info.0.encryption_at_rest_kms_key_arn` - The ARN of the KMS key used for encryption at rest of the broker data volumes.
        :param pulumi.Input[dict] encryption_info: Configuration block for specifying encryption. See below.
        :param pulumi.Input[str] enhanced_monitoring: Specify the desired enhanced MSK CloudWatch monitoring level.  See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
        :param pulumi.Input[str] kafka_version: Specify the desired Kafka software version.
        :param pulumi.Input[float] number_of_broker_nodes: The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource
        :param pulumi.Input[str] zookeeper_connect_string: A comma separated list of one or more IP:port pairs to use to connect to the Apache Zookeeper cluster.
        
        The **broker_node_group_info** object supports the following:
        
          * `az_distribution` (`pulumi.Input[str]`) - The distribution of broker nodes across availability zones ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-model-brokerazdistribution)). Currently the only valid value is `DEFAULT`.
          * `client_subnets` (`pulumi.Input[list]`) - A list of subnets to connect to in client VPC ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-prop-brokernodegroupinfo-clientsubnets)).
          * `ebs_volume_size` (`pulumi.Input[float]`) - The size in GiB of the EBS volume for the data drive on each broker node.
          * `instance_type` (`pulumi.Input[str]`) - Specify the instance type to use for the kafka brokers. e.g. kafka.m5.large. ([Pricing info](https://aws.amazon.com/msk/pricing/))
          * `security_groups` (`pulumi.Input[list]`) - A list of the security groups to associate with the elastic network interfaces to control who can communicate with the cluster.
        
        The **client_authentication** object supports the following:
        
          * `tls` (`pulumi.Input[dict]`) - Configuration block for specifying TLS client authentication. See below.
        
            * `certificate_authority_arns` (`pulumi.Input[list]`) - List of ACM Certificate Authority Amazon Resource Names (ARNs).
        
        The **configuration_info** object supports the following:
        
          * `arn` (`pulumi.Input[str]`) - Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
          * `revision` (`pulumi.Input[float]`) - Revision of the MSK Configuration to use in the cluster.
        
        The **encryption_info** object supports the following:
        
          * `encryption_at_rest_kms_key_arn` (`pulumi.Input[str]`) - You may specify a KMS key short ID or ARN (it will always output an ARN) to use for encrypting your data at rest.  If no key is specified, an AWS managed KMS ('aws/msk' managed service) key will be used for encrypting the data at rest.
          * `encryption_in_transit` (`pulumi.Input[dict]`) - Configuration block to specify encryption in transit. See below.
        
            * `client_broker` (`pulumi.Input[str]`) - Encryption setting for data in transit between clients and brokers. Valid values: `TLS`, `TLS_PLAINTEXT`, and `PLAINTEXT`. Default value: `TLS_PLAINTEXT`.
            * `in_cluster` (`pulumi.Input[bool]`) - Whether data communication among broker nodes is encrypted. Default value: `true`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/msk_cluster.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["arn"] = arn
        __props__["bootstrap_brokers"] = bootstrap_brokers
        __props__["bootstrap_brokers_tls"] = bootstrap_brokers_tls
        __props__["broker_node_group_info"] = broker_node_group_info
        __props__["client_authentication"] = client_authentication
        __props__["cluster_name"] = cluster_name
        __props__["configuration_info"] = configuration_info
        __props__["current_version"] = current_version
        __props__["encryption_info"] = encryption_info
        __props__["enhanced_monitoring"] = enhanced_monitoring
        __props__["kafka_version"] = kafka_version
        __props__["number_of_broker_nodes"] = number_of_broker_nodes
        __props__["tags"] = tags
        __props__["zookeeper_connect_string"] = zookeeper_connect_string
        return Cluster(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

