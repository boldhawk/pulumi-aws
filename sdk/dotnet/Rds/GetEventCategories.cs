// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds
{
    public static partial class Invokes
    {
        public static Task<GetEventCategoriesResult> GetEventCategories(GetEventCategoriesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEventCategoriesResult>("aws:rds/getEventCategories:getEventCategories", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetEventCategoriesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of source that will be generating the events. Valid options are db-instance, db-security-group, db-parameter-group, db-snapshot, db-cluster or db-cluster-snapshot.
        /// </summary>
        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        public GetEventCategoriesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetEventCategoriesResult
    {
        /// <summary>
        /// A list of the event categories.
        /// </summary>
        public readonly ImmutableArray<string> EventCategories;
        public readonly string? SourceType;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetEventCategoriesResult(
            ImmutableArray<string> eventCategories,
            string? sourceType,
            string id)
        {
            EventCategories = eventCategories;
            SourceType = sourceType;
            Id = id;
        }
    }
}
