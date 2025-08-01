// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Instance
{
    public static class GetPlacementGroup
    {
        /// <summary>
        /// Gets information about a Security Group.
        /// </summary>
        public static Task<GetPlacementGroupResult> InvokeAsync(GetPlacementGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPlacementGroupResult>("scaleway:instance/getPlacementGroup:getPlacementGroup", args ?? new GetPlacementGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Security Group.
        /// </summary>
        public static Output<GetPlacementGroupResult> Invoke(GetPlacementGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPlacementGroupResult>("scaleway:instance/getPlacementGroup:getPlacementGroup", args ?? new GetPlacementGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Security Group.
        /// </summary>
        public static Output<GetPlacementGroupResult> Invoke(GetPlacementGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPlacementGroupResult>("scaleway:instance/getPlacementGroup:getPlacementGroup", args ?? new GetPlacementGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPlacementGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The placement group name. Only one of `name` and `placement_group_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The placement group id. Only one of `name` and `placement_group_id` should be specified.
        /// </summary>
        [Input("placementGroupId")]
        public string? PlacementGroupId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the placement group is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the placement group exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetPlacementGroupArgs()
        {
        }
        public static new GetPlacementGroupArgs Empty => new GetPlacementGroupArgs();
    }

    public sealed class GetPlacementGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The placement group name. Only one of `name` and `placement_group_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The placement group id. Only one of `name` and `placement_group_id` should be specified.
        /// </summary>
        [Input("placementGroupId")]
        public Input<string>? PlacementGroupId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the placement group is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the placement group exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetPlacementGroupInvokeArgs()
        {
        }
        public static new GetPlacementGroupInvokeArgs Empty => new GetPlacementGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetPlacementGroupResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        /// <summary>
        /// The organization ID the placement group is associated with.
        /// </summary>
        public readonly string OrganizationId;
        public readonly string? PlacementGroupId;
        /// <summary>
        /// The [policy mode](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) of the placement group.
        /// </summary>
        public readonly string PolicyMode;
        /// <summary>
        /// Is true when the policy is respected.
        /// </summary>
        public readonly bool PolicyRespected;
        /// <summary>
        /// The [policy type](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) of the placement group.
        /// </summary>
        public readonly string PolicyType;
        public readonly string ProjectId;
        /// <summary>
        /// A list of tags to apply to the placement group.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        public readonly string? Zone;

        [OutputConstructor]
        private GetPlacementGroupResult(
            string id,

            string? name,

            string organizationId,

            string? placementGroupId,

            string policyMode,

            bool policyRespected,

            string policyType,

            string projectId,

            ImmutableArray<string> tags,

            string? zone)
        {
            Id = id;
            Name = name;
            OrganizationId = organizationId;
            PlacementGroupId = placementGroupId;
            PolicyMode = policyMode;
            PolicyRespected = policyRespected;
            PolicyType = policyType;
            ProjectId = projectId;
            Tags = tags;
            Zone = zone;
        }
    }
}
