# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ListenerRule(pulumi.CustomResource):
    actions: pulumi.Output[list]
    """
    An Action block. Action blocks are documented below.
    
      * `authenticate_cognito` (`dict`) - Information for creating an authenticate action using Cognito. Required if `type` is `authenticate-cognito`.
    
        * `authentication_request_extra_params` (`dict`) - The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
        * `on_unauthenticated_request` (`str`) - The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
        * `scope` (`str`) - The set of user claims to be requested from the IdP.
        * `session_cookie_name` (`str`) - The name of the cookie used to maintain session information.
        * `session_timeout` (`float`) - The maximum duration of the authentication session, in seconds.
        * `user_pool_arn` (`str`) - The ARN of the Cognito user pool.
        * `user_pool_client_id` (`str`) - The ID of the Cognito user pool client.
        * `user_pool_domain` (`str`) - The domain prefix or fully-qualified domain name of the Cognito user pool.
    
      * `authenticate_oidc` (`dict`) - Information for creating an authenticate action using OIDC. Required if `type` is `authenticate-oidc`.
    
        * `authentication_request_extra_params` (`dict`) - The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
        * `authorization_endpoint` (`str`) - The authorization endpoint of the IdP.
        * `client_id` (`str`) - The OAuth 2.0 client identifier.
        * `client_secret` (`str`) - The OAuth 2.0 client secret.
        * `issuer` (`str`) - The OIDC issuer identifier of the IdP.
        * `on_unauthenticated_request` (`str`) - The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
        * `scope` (`str`) - The set of user claims to be requested from the IdP.
        * `session_cookie_name` (`str`) - The name of the cookie used to maintain session information.
        * `session_timeout` (`float`) - The maximum duration of the authentication session, in seconds.
        * `token_endpoint` (`str`) - The token endpoint of the IdP.
        * `user_info_endpoint` (`str`) - The user info endpoint of the IdP.
    
      * `fixed_response` (`dict`) - Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
    
        * `content_type` (`str`) - The content type. Valid values are `text/plain`, `text/css`, `text/html`, `application/javascript` and `application/json`.
        * `message_body` (`str`) - The message body.
        * `status_code` (`str`) - The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
    
      * `order` (`float`)
      * `redirect` (`dict`) - Information for creating a redirect action. Required if `type` is `redirect`.
    
        * `host` (`str`) - The hostname. This component is not percent-encoded. The hostname can contain `#{host}`. Defaults to `#{host}`.
        * `path` (`str`) - The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to `/#{path}`.
        * `port` (`str`) - The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
        * `protocol` (`str`) - The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
        * `query` (`str`) - The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to `#{query}`.
        * `status_code` (`str`) - The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
    
      * `target_group_arn` (`str`) - The ARN of the Target Group to which to route traffic. Required if `type` is `forward`.
      * `type` (`str`) - The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.
    """
    arn: pulumi.Output[str]
    """
    The ARN of the rule (matches `id`)
    """
    conditions: pulumi.Output[list]
    """
    A Condition block. Condition blocks are documented below.
    
      * `field` (`str`) - The name of the field. Must be one of `path-pattern` for path based routing or `host-header` for host based routing.
      * `values` (`str`) - The path patterns to match. A maximum of 1 can be defined.
    """
    listener_arn: pulumi.Output[str]
    """
    The ARN of the listener to which to attach the rule.
    """
    priority: pulumi.Output[float]
    """
    The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
    """
    def __init__(__self__, resource_name, opts=None, actions=None, conditions=None, listener_arn=None, priority=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Load Balancer Listener Rule resource.
        
        > **Note:** `alb.ListenerRule` is known as `lb.ListenerRule`. The functionality is identical.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] actions: An Action block. Action blocks are documented below.
        :param pulumi.Input[list] conditions: A Condition block. Condition blocks are documented below.
        :param pulumi.Input[str] listener_arn: The ARN of the listener to which to attach the rule.
        :param pulumi.Input[float] priority: The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
        
        The **actions** object supports the following:
        
          * `authenticate_cognito` (`pulumi.Input[dict]`) - Information for creating an authenticate action using Cognito. Required if `type` is `authenticate-cognito`.
        
            * `authentication_request_extra_params` (`pulumi.Input[dict]`) - The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
            * `on_unauthenticated_request` (`pulumi.Input[str]`) - The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
            * `scope` (`pulumi.Input[str]`) - The set of user claims to be requested from the IdP.
            * `session_cookie_name` (`pulumi.Input[str]`) - The name of the cookie used to maintain session information.
            * `session_timeout` (`pulumi.Input[float]`) - The maximum duration of the authentication session, in seconds.
            * `user_pool_arn` (`pulumi.Input[str]`) - The ARN of the Cognito user pool.
            * `user_pool_client_id` (`pulumi.Input[str]`) - The ID of the Cognito user pool client.
            * `user_pool_domain` (`pulumi.Input[str]`) - The domain prefix or fully-qualified domain name of the Cognito user pool.
        
          * `authenticate_oidc` (`pulumi.Input[dict]`) - Information for creating an authenticate action using OIDC. Required if `type` is `authenticate-oidc`.
        
            * `authentication_request_extra_params` (`pulumi.Input[dict]`) - The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
            * `authorization_endpoint` (`pulumi.Input[str]`) - The authorization endpoint of the IdP.
            * `client_id` (`pulumi.Input[str]`) - The OAuth 2.0 client identifier.
            * `client_secret` (`pulumi.Input[str]`) - The OAuth 2.0 client secret.
            * `issuer` (`pulumi.Input[str]`) - The OIDC issuer identifier of the IdP.
            * `on_unauthenticated_request` (`pulumi.Input[str]`) - The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
            * `scope` (`pulumi.Input[str]`) - The set of user claims to be requested from the IdP.
            * `session_cookie_name` (`pulumi.Input[str]`) - The name of the cookie used to maintain session information.
            * `session_timeout` (`pulumi.Input[float]`) - The maximum duration of the authentication session, in seconds.
            * `token_endpoint` (`pulumi.Input[str]`) - The token endpoint of the IdP.
            * `user_info_endpoint` (`pulumi.Input[str]`) - The user info endpoint of the IdP.
        
          * `fixed_response` (`pulumi.Input[dict]`) - Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
        
            * `content_type` (`pulumi.Input[str]`) - The content type. Valid values are `text/plain`, `text/css`, `text/html`, `application/javascript` and `application/json`.
            * `message_body` (`pulumi.Input[str]`) - The message body.
            * `status_code` (`pulumi.Input[str]`) - The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
        
          * `order` (`pulumi.Input[float]`)
          * `redirect` (`pulumi.Input[dict]`) - Information for creating a redirect action. Required if `type` is `redirect`.
        
            * `host` (`pulumi.Input[str]`) - The hostname. This component is not percent-encoded. The hostname can contain `#{host}`. Defaults to `#{host}`.
            * `path` (`pulumi.Input[str]`) - The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to `/#{path}`.
            * `port` (`pulumi.Input[str]`) - The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
            * `protocol` (`pulumi.Input[str]`) - The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
            * `query` (`pulumi.Input[str]`) - The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to `#{query}`.
            * `status_code` (`pulumi.Input[str]`) - The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
        
          * `target_group_arn` (`pulumi.Input[str]`) - The ARN of the Target Group to which to route traffic. Required if `type` is `forward`.
          * `type` (`pulumi.Input[str]`) - The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.
        
        The **conditions** object supports the following:
        
          * `field` (`pulumi.Input[str]`) - The name of the field. Must be one of `path-pattern` for path based routing or `host-header` for host based routing.
          * `values` (`pulumi.Input[str]`) - The path patterns to match. A maximum of 1 can be defined.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_listener_rule_legacy.html.markdown.
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

            if actions is None:
                raise TypeError("Missing required property 'actions'")
            __props__['actions'] = actions
            if conditions is None:
                raise TypeError("Missing required property 'conditions'")
            __props__['conditions'] = conditions
            if listener_arn is None:
                raise TypeError("Missing required property 'listener_arn'")
            __props__['listener_arn'] = listener_arn
            __props__['priority'] = priority
            __props__['arn'] = None
        super(ListenerRule, __self__).__init__(
            'aws:elasticloadbalancingv2/listenerRule:ListenerRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, actions=None, arn=None, conditions=None, listener_arn=None, priority=None):
        """
        Get an existing ListenerRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] actions: An Action block. Action blocks are documented below.
        :param pulumi.Input[str] arn: The ARN of the rule (matches `id`)
        :param pulumi.Input[list] conditions: A Condition block. Condition blocks are documented below.
        :param pulumi.Input[str] listener_arn: The ARN of the listener to which to attach the rule.
        :param pulumi.Input[float] priority: The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
        
        The **actions** object supports the following:
        
          * `authenticate_cognito` (`pulumi.Input[dict]`) - Information for creating an authenticate action using Cognito. Required if `type` is `authenticate-cognito`.
        
            * `authentication_request_extra_params` (`pulumi.Input[dict]`) - The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
            * `on_unauthenticated_request` (`pulumi.Input[str]`) - The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
            * `scope` (`pulumi.Input[str]`) - The set of user claims to be requested from the IdP.
            * `session_cookie_name` (`pulumi.Input[str]`) - The name of the cookie used to maintain session information.
            * `session_timeout` (`pulumi.Input[float]`) - The maximum duration of the authentication session, in seconds.
            * `user_pool_arn` (`pulumi.Input[str]`) - The ARN of the Cognito user pool.
            * `user_pool_client_id` (`pulumi.Input[str]`) - The ID of the Cognito user pool client.
            * `user_pool_domain` (`pulumi.Input[str]`) - The domain prefix or fully-qualified domain name of the Cognito user pool.
        
          * `authenticate_oidc` (`pulumi.Input[dict]`) - Information for creating an authenticate action using OIDC. Required if `type` is `authenticate-oidc`.
        
            * `authentication_request_extra_params` (`pulumi.Input[dict]`) - The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
            * `authorization_endpoint` (`pulumi.Input[str]`) - The authorization endpoint of the IdP.
            * `client_id` (`pulumi.Input[str]`) - The OAuth 2.0 client identifier.
            * `client_secret` (`pulumi.Input[str]`) - The OAuth 2.0 client secret.
            * `issuer` (`pulumi.Input[str]`) - The OIDC issuer identifier of the IdP.
            * `on_unauthenticated_request` (`pulumi.Input[str]`) - The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
            * `scope` (`pulumi.Input[str]`) - The set of user claims to be requested from the IdP.
            * `session_cookie_name` (`pulumi.Input[str]`) - The name of the cookie used to maintain session information.
            * `session_timeout` (`pulumi.Input[float]`) - The maximum duration of the authentication session, in seconds.
            * `token_endpoint` (`pulumi.Input[str]`) - The token endpoint of the IdP.
            * `user_info_endpoint` (`pulumi.Input[str]`) - The user info endpoint of the IdP.
        
          * `fixed_response` (`pulumi.Input[dict]`) - Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
        
            * `content_type` (`pulumi.Input[str]`) - The content type. Valid values are `text/plain`, `text/css`, `text/html`, `application/javascript` and `application/json`.
            * `message_body` (`pulumi.Input[str]`) - The message body.
            * `status_code` (`pulumi.Input[str]`) - The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
        
          * `order` (`pulumi.Input[float]`)
          * `redirect` (`pulumi.Input[dict]`) - Information for creating a redirect action. Required if `type` is `redirect`.
        
            * `host` (`pulumi.Input[str]`) - The hostname. This component is not percent-encoded. The hostname can contain `#{host}`. Defaults to `#{host}`.
            * `path` (`pulumi.Input[str]`) - The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to `/#{path}`.
            * `port` (`pulumi.Input[str]`) - The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
            * `protocol` (`pulumi.Input[str]`) - The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
            * `query` (`pulumi.Input[str]`) - The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to `#{query}`.
            * `status_code` (`pulumi.Input[str]`) - The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
        
          * `target_group_arn` (`pulumi.Input[str]`) - The ARN of the Target Group to which to route traffic. Required if `type` is `forward`.
          * `type` (`pulumi.Input[str]`) - The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.
        
        The **conditions** object supports the following:
        
          * `field` (`pulumi.Input[str]`) - The name of the field. Must be one of `path-pattern` for path based routing or `host-header` for host based routing.
          * `values` (`pulumi.Input[str]`) - The path patterns to match. A maximum of 1 can be defined.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_listener_rule_legacy.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["actions"] = actions
        __props__["arn"] = arn
        __props__["conditions"] = conditions
        __props__["listener_arn"] = listener_arn
        __props__["priority"] = priority
        return ListenerRule(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

