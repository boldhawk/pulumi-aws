# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class TaskDefinition(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Full ARN of the Task Definition (including both `family` and `revision`).
    """
    container_definitions: pulumi.Output[str]
    """
    A list of valid [container definitions]
    (http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) provided as a
    single valid JSON document. Please note that you should only provide values that are part of the container
    definition document. For a detailed description of what parameters are available, see the [Task Definition Parameters]
    (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) section from the
    official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
    """
    cpu: pulumi.Output[str]
    """
    The number of cpu units used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
    """
    execution_role_arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
    """
    family: pulumi.Output[str]
    """
    A unique name for your task definition.
    """
    ipc_mode: pulumi.Output[str]
    """
    The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
    """
    memory: pulumi.Output[str]
    """
    The amount (in MiB) of memory used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
    """
    network_mode: pulumi.Output[str]
    """
    The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
    """
    pid_mode: pulumi.Output[str]
    """
    The process namespace to use for the containers in the task. The valid values are `host` and `task`.
    """
    placement_constraints: pulumi.Output[list]
    """
    A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placement_constraints` is `10`.
    
      * `expression` (`str`) - Cluster Query Language expression to apply to the constraint.
        For more information, see [Cluster Query Language in the Amazon EC2 Container
        Service Developer
        Guide](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html).
      * `type` (`str`) - The proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.
    """
    proxy_configuration: pulumi.Output[dict]
    """
    The proxy configuration details for the App Mesh proxy.
    
      * `containerName` (`str`) - The name of the container that will serve as the App Mesh proxy.
      * `properties` (`dict`) - The set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified a key-value mapping.
      * `type` (`str`) - The proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.
    """
    requires_compatibilities: pulumi.Output[list]
    """
    A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
    """
    revision: pulumi.Output[float]
    """
    The revision of the task in a particular family.
    """
    tags: pulumi.Output[dict]
    """
    Key-value mapping of resource tags
    """
    task_role_arn: pulumi.Output[str]
    """
    The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
    """
    volumes: pulumi.Output[list]
    """
    A set of volume blocks that containers in your task may use.
    
      * `dockerVolumeConfiguration` (`dict`) - Used to configure a docker volume
    
        * `autoprovision` (`bool`) - If this value is `true`, the Docker volume is created if it does not already exist. *Note*: This field is only used if the scope is `shared`.
        * `driver` (`str`) - The Docker volume driver to use. The driver value must match the driver name provided by Docker because it is used for task placement.
        * `driverOpts` (`dict`) - A map of Docker driver specific options.
        * `labels` (`dict`) - A map of custom metadata to add to your Docker volume.
        * `scope` (`str`) - The scope for the Docker volume, which determines its lifecycle, either `task` or `shared`.  Docker volumes that are scoped to a `task` are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are `scoped` as shared persist after the task stops.
    
      * `hostPath` (`str`) - The path on the host container instance that is presented to the container. If not set, ECS will create a nonpersistent data volume that starts empty and is deleted after the task has finished.
      * `name` (`str`) - The name of the volume. This name is referenced in the `sourceVolume`
        parameter of container definition in the `mountPoints` section.
    """
    def __init__(__self__, resource_name, opts=None, container_definitions=None, cpu=None, execution_role_arn=None, family=None, ipc_mode=None, memory=None, network_mode=None, pid_mode=None, placement_constraints=None, proxy_configuration=None, requires_compatibilities=None, tags=None, task_role_arn=None, volumes=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a revision of an ECS task definition to be used in `ecs.Service`.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_definitions: A list of valid [container definitions]
               (http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) provided as a
               single valid JSON document. Please note that you should only provide values that are part of the container
               definition document. For a detailed description of what parameters are available, see the [Task Definition Parameters]
               (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) section from the
               official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
        :param pulumi.Input[str] cpu: The number of cpu units used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
        :param pulumi.Input[str] execution_role_arn: The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
        :param pulumi.Input[str] family: A unique name for your task definition.
        :param pulumi.Input[str] ipc_mode: The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
        :param pulumi.Input[str] memory: The amount (in MiB) of memory used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
        :param pulumi.Input[str] network_mode: The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
        :param pulumi.Input[str] pid_mode: The process namespace to use for the containers in the task. The valid values are `host` and `task`.
        :param pulumi.Input[list] placement_constraints: A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placement_constraints` is `10`.
        :param pulumi.Input[dict] proxy_configuration: The proxy configuration details for the App Mesh proxy.
        :param pulumi.Input[list] requires_compatibilities: A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags
        :param pulumi.Input[str] task_role_arn: The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
        :param pulumi.Input[list] volumes: A set of volume blocks that containers in your task may use.
        
        The **placement_constraints** object supports the following:
        
          * `expression` (`pulumi.Input[str]`) - Cluster Query Language expression to apply to the constraint.
            For more information, see [Cluster Query Language in the Amazon EC2 Container
            Service Developer
            Guide](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html).
          * `type` (`pulumi.Input[str]`) - The proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.
        
        The **proxy_configuration** object supports the following:
        
          * `containerName` (`pulumi.Input[str]`) - The name of the container that will serve as the App Mesh proxy.
          * `properties` (`pulumi.Input[dict]`) - The set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified a key-value mapping.
          * `type` (`pulumi.Input[str]`) - The proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.
        
        The **volumes** object supports the following:
        
          * `dockerVolumeConfiguration` (`pulumi.Input[dict]`) - Used to configure a docker volume
        
            * `autoprovision` (`pulumi.Input[bool]`) - If this value is `true`, the Docker volume is created if it does not already exist. *Note*: This field is only used if the scope is `shared`.
            * `driver` (`pulumi.Input[str]`) - The Docker volume driver to use. The driver value must match the driver name provided by Docker because it is used for task placement.
            * `driverOpts` (`pulumi.Input[dict]`) - A map of Docker driver specific options.
            * `labels` (`pulumi.Input[dict]`) - A map of custom metadata to add to your Docker volume.
            * `scope` (`pulumi.Input[str]`) - The scope for the Docker volume, which determines its lifecycle, either `task` or `shared`.  Docker volumes that are scoped to a `task` are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are `scoped` as shared persist after the task stops.
        
          * `hostPath` (`pulumi.Input[str]`) - The path on the host container instance that is presented to the container. If not set, ECS will create a nonpersistent data volume that starts empty and is deleted after the task has finished.
          * `name` (`pulumi.Input[str]`) - The name of the volume. This name is referenced in the `sourceVolume`
            parameter of container definition in the `mountPoints` section.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ecs_task_definition.html.markdown.
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

            if container_definitions is None:
                raise TypeError("Missing required property 'container_definitions'")
            __props__['container_definitions'] = container_definitions
            __props__['cpu'] = cpu
            __props__['execution_role_arn'] = execution_role_arn
            if family is None:
                raise TypeError("Missing required property 'family'")
            __props__['family'] = family
            __props__['ipc_mode'] = ipc_mode
            __props__['memory'] = memory
            __props__['network_mode'] = network_mode
            __props__['pid_mode'] = pid_mode
            __props__['placement_constraints'] = placement_constraints
            __props__['proxy_configuration'] = proxy_configuration
            __props__['requires_compatibilities'] = requires_compatibilities
            __props__['tags'] = tags
            __props__['task_role_arn'] = task_role_arn
            __props__['volumes'] = volumes
            __props__['arn'] = None
            __props__['revision'] = None
        super(TaskDefinition, __self__).__init__(
            'aws:ecs/taskDefinition:TaskDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, container_definitions=None, cpu=None, execution_role_arn=None, family=None, ipc_mode=None, memory=None, network_mode=None, pid_mode=None, placement_constraints=None, proxy_configuration=None, requires_compatibilities=None, revision=None, tags=None, task_role_arn=None, volumes=None):
        """
        Get an existing TaskDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Full ARN of the Task Definition (including both `family` and `revision`).
        :param pulumi.Input[str] container_definitions: A list of valid [container definitions]
               (http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) provided as a
               single valid JSON document. Please note that you should only provide values that are part of the container
               definition document. For a detailed description of what parameters are available, see the [Task Definition Parameters]
               (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) section from the
               official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
        :param pulumi.Input[str] cpu: The number of cpu units used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
        :param pulumi.Input[str] execution_role_arn: The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
        :param pulumi.Input[str] family: A unique name for your task definition.
        :param pulumi.Input[str] ipc_mode: The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
        :param pulumi.Input[str] memory: The amount (in MiB) of memory used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
        :param pulumi.Input[str] network_mode: The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
        :param pulumi.Input[str] pid_mode: The process namespace to use for the containers in the task. The valid values are `host` and `task`.
        :param pulumi.Input[list] placement_constraints: A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placement_constraints` is `10`.
        :param pulumi.Input[dict] proxy_configuration: The proxy configuration details for the App Mesh proxy.
        :param pulumi.Input[list] requires_compatibilities: A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
        :param pulumi.Input[float] revision: The revision of the task in a particular family.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags
        :param pulumi.Input[str] task_role_arn: The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
        :param pulumi.Input[list] volumes: A set of volume blocks that containers in your task may use.
        
        The **placement_constraints** object supports the following:
        
          * `expression` (`pulumi.Input[str]`) - Cluster Query Language expression to apply to the constraint.
            For more information, see [Cluster Query Language in the Amazon EC2 Container
            Service Developer
            Guide](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html).
          * `type` (`pulumi.Input[str]`) - The proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.
        
        The **proxy_configuration** object supports the following:
        
          * `containerName` (`pulumi.Input[str]`) - The name of the container that will serve as the App Mesh proxy.
          * `properties` (`pulumi.Input[dict]`) - The set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified a key-value mapping.
          * `type` (`pulumi.Input[str]`) - The proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.
        
        The **volumes** object supports the following:
        
          * `dockerVolumeConfiguration` (`pulumi.Input[dict]`) - Used to configure a docker volume
        
            * `autoprovision` (`pulumi.Input[bool]`) - If this value is `true`, the Docker volume is created if it does not already exist. *Note*: This field is only used if the scope is `shared`.
            * `driver` (`pulumi.Input[str]`) - The Docker volume driver to use. The driver value must match the driver name provided by Docker because it is used for task placement.
            * `driverOpts` (`pulumi.Input[dict]`) - A map of Docker driver specific options.
            * `labels` (`pulumi.Input[dict]`) - A map of custom metadata to add to your Docker volume.
            * `scope` (`pulumi.Input[str]`) - The scope for the Docker volume, which determines its lifecycle, either `task` or `shared`.  Docker volumes that are scoped to a `task` are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are `scoped` as shared persist after the task stops.
        
          * `hostPath` (`pulumi.Input[str]`) - The path on the host container instance that is presented to the container. If not set, ECS will create a nonpersistent data volume that starts empty and is deleted after the task has finished.
          * `name` (`pulumi.Input[str]`) - The name of the volume. This name is referenced in the `sourceVolume`
            parameter of container definition in the `mountPoints` section.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ecs_task_definition.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["arn"] = arn
        __props__["container_definitions"] = container_definitions
        __props__["cpu"] = cpu
        __props__["execution_role_arn"] = execution_role_arn
        __props__["family"] = family
        __props__["ipc_mode"] = ipc_mode
        __props__["memory"] = memory
        __props__["network_mode"] = network_mode
        __props__["pid_mode"] = pid_mode
        __props__["placement_constraints"] = placement_constraints
        __props__["proxy_configuration"] = proxy_configuration
        __props__["requires_compatibilities"] = requires_compatibilities
        __props__["revision"] = revision
        __props__["tags"] = tags
        __props__["task_role_arn"] = task_role_arn
        __props__["volumes"] = volumes
        return TaskDefinition(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

