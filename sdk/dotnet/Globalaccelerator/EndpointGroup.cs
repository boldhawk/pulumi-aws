// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GlobalAccelerator
{
    /// <summary>
    /// Provides a Global Accelerator endpoint group.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/globalaccelerator_endpoint_group.html.markdown.
    /// </summary>
    public partial class EndpointGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The list of endpoint objects. Fields documented below.
        /// </summary>
        [Output("endpointConfigurations")]
        public Output<ImmutableArray<Outputs.EndpointGroupEndpointConfigurations>> EndpointConfigurations { get; private set; } = null!;

        [Output("endpointGroupRegion")]
        public Output<string> EndpointGroupRegion { get; private set; } = null!;

        /// <summary>
        /// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
        /// </summary>
        [Output("healthCheckIntervalSeconds")]
        public Output<int?> HealthCheckIntervalSeconds { get; private set; } = null!;

        /// <summary>
        /// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (/).
        /// </summary>
        [Output("healthCheckPath")]
        public Output<string?> HealthCheckPath { get; private set; } = null!;

        /// <summary>
        /// The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
        /// </summary>
        [Output("healthCheckPort")]
        public Output<int?> HealthCheckPort { get; private set; } = null!;

        /// <summary>
        /// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
        /// </summary>
        [Output("healthCheckProtocol")]
        public Output<string?> HealthCheckProtocol { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the listener.
        /// </summary>
        [Output("listenerArn")]
        public Output<string> ListenerArn { get; private set; } = null!;

        /// <summary>
        /// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
        /// </summary>
        [Output("thresholdCount")]
        public Output<int?> ThresholdCount { get; private set; } = null!;

        /// <summary>
        /// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
        /// </summary>
        [Output("trafficDialPercentage")]
        public Output<double?> TrafficDialPercentage { get; private set; } = null!;


        /// <summary>
        /// Create a EndpointGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EndpointGroup(string name, EndpointGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:globalaccelerator/endpointGroup:EndpointGroup", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private EndpointGroup(string name, Input<string> id, EndpointGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:globalaccelerator/endpointGroup:EndpointGroup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EndpointGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EndpointGroup Get(string name, Input<string> id, EndpointGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new EndpointGroup(name, id, state, options);
        }
    }

    public sealed class EndpointGroupArgs : Pulumi.ResourceArgs
    {
        [Input("endpointConfigurations")]
        private InputList<Inputs.EndpointGroupEndpointConfigurationsArgs>? _endpointConfigurations;

        /// <summary>
        /// The list of endpoint objects. Fields documented below.
        /// </summary>
        public InputList<Inputs.EndpointGroupEndpointConfigurationsArgs> EndpointConfigurations
        {
            get => _endpointConfigurations ?? (_endpointConfigurations = new InputList<Inputs.EndpointGroupEndpointConfigurationsArgs>());
            set => _endpointConfigurations = value;
        }

        [Input("endpointGroupRegion")]
        public Input<string>? EndpointGroupRegion { get; set; }

        /// <summary>
        /// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
        /// </summary>
        [Input("healthCheckIntervalSeconds")]
        public Input<int>? HealthCheckIntervalSeconds { get; set; }

        /// <summary>
        /// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (/).
        /// </summary>
        [Input("healthCheckPath")]
        public Input<string>? HealthCheckPath { get; set; }

        /// <summary>
        /// The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
        /// </summary>
        [Input("healthCheckPort")]
        public Input<int>? HealthCheckPort { get; set; }

        /// <summary>
        /// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
        /// </summary>
        [Input("healthCheckProtocol")]
        public Input<string>? HealthCheckProtocol { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the listener.
        /// </summary>
        [Input("listenerArn", required: true)]
        public Input<string> ListenerArn { get; set; } = null!;

        /// <summary>
        /// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
        /// </summary>
        [Input("thresholdCount")]
        public Input<int>? ThresholdCount { get; set; }

        /// <summary>
        /// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
        /// </summary>
        [Input("trafficDialPercentage")]
        public Input<double>? TrafficDialPercentage { get; set; }

        public EndpointGroupArgs()
        {
        }
    }

    public sealed class EndpointGroupState : Pulumi.ResourceArgs
    {
        [Input("endpointConfigurations")]
        private InputList<Inputs.EndpointGroupEndpointConfigurationsGetArgs>? _endpointConfigurations;

        /// <summary>
        /// The list of endpoint objects. Fields documented below.
        /// </summary>
        public InputList<Inputs.EndpointGroupEndpointConfigurationsGetArgs> EndpointConfigurations
        {
            get => _endpointConfigurations ?? (_endpointConfigurations = new InputList<Inputs.EndpointGroupEndpointConfigurationsGetArgs>());
            set => _endpointConfigurations = value;
        }

        [Input("endpointGroupRegion")]
        public Input<string>? EndpointGroupRegion { get; set; }

        /// <summary>
        /// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
        /// </summary>
        [Input("healthCheckIntervalSeconds")]
        public Input<int>? HealthCheckIntervalSeconds { get; set; }

        /// <summary>
        /// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (/).
        /// </summary>
        [Input("healthCheckPath")]
        public Input<string>? HealthCheckPath { get; set; }

        /// <summary>
        /// The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
        /// </summary>
        [Input("healthCheckPort")]
        public Input<int>? HealthCheckPort { get; set; }

        /// <summary>
        /// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
        /// </summary>
        [Input("healthCheckProtocol")]
        public Input<string>? HealthCheckProtocol { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the listener.
        /// </summary>
        [Input("listenerArn")]
        public Input<string>? ListenerArn { get; set; }

        /// <summary>
        /// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
        /// </summary>
        [Input("thresholdCount")]
        public Input<int>? ThresholdCount { get; set; }

        /// <summary>
        /// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
        /// </summary>
        [Input("trafficDialPercentage")]
        public Input<double>? TrafficDialPercentage { get; set; }

        public EndpointGroupState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class EndpointGroupEndpointConfigurationsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An ID for the endpoint. If the endpoint is a Network Load Balancer or Application Load Balancer, this is the Amazon Resource Name (ARN) of the resource. If the endpoint is an Elastic IP address, this is the Elastic IP address allocation ID.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// The weight associated with the endpoint. When you add weights to endpoints, you configure AWS Global Accelerator to route traffic based on proportions that you specify. 
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public EndpointGroupEndpointConfigurationsArgs()
        {
        }
    }

    public sealed class EndpointGroupEndpointConfigurationsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An ID for the endpoint. If the endpoint is a Network Load Balancer or Application Load Balancer, this is the Amazon Resource Name (ARN) of the resource. If the endpoint is an Elastic IP address, this is the Elastic IP address allocation ID.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// The weight associated with the endpoint. When you add weights to endpoints, you configure AWS Global Accelerator to route traffic based on proportions that you specify. 
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public EndpointGroupEndpointConfigurationsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class EndpointGroupEndpointConfigurations
    {
        /// <summary>
        /// An ID for the endpoint. If the endpoint is a Network Load Balancer or Application Load Balancer, this is the Amazon Resource Name (ARN) of the resource. If the endpoint is an Elastic IP address, this is the Elastic IP address allocation ID.
        /// </summary>
        public readonly string? EndpointId;
        /// <summary>
        /// The weight associated with the endpoint. When you add weights to endpoints, you configure AWS Global Accelerator to route traffic based on proportions that you specify. 
        /// </summary>
        public readonly int? Weight;

        [OutputConstructor]
        private EndpointGroupEndpointConfigurations(
            string? endpointId,
            int? weight)
        {
            EndpointId = endpointId;
            Weight = weight;
        }
    }
    }
}
