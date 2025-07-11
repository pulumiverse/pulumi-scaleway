// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway
{
    [Obsolete(@"scaleway.index/getbillingconsumptions.getBillingConsumptions has been deprecated in favor of scaleway.billing/getconsumptions.getConsumptions")]
    public static class GetBillingConsumptions
    {
        /// <summary>
        /// Gets information about your Consumptions.
        /// </summary>
        public static Task<GetBillingConsumptionsResult> InvokeAsync(GetBillingConsumptionsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBillingConsumptionsResult>("scaleway:index/getBillingConsumptions:getBillingConsumptions", args ?? new GetBillingConsumptionsArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about your Consumptions.
        /// </summary>
        public static Output<GetBillingConsumptionsResult> Invoke(GetBillingConsumptionsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBillingConsumptionsResult>("scaleway:index/getBillingConsumptions:getBillingConsumptions", args ?? new GetBillingConsumptionsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about your Consumptions.
        /// </summary>
        public static Output<GetBillingConsumptionsResult> Invoke(GetBillingConsumptionsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetBillingConsumptionsResult>("scaleway:index/getBillingConsumptions:getBillingConsumptions", args ?? new GetBillingConsumptionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBillingConsumptionsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// `project_id`) The ID of the project the consumption list is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        public GetBillingConsumptionsArgs()
        {
        }
        public static new GetBillingConsumptionsArgs Empty => new GetBillingConsumptionsArgs();
    }

    public sealed class GetBillingConsumptionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// `project_id`) The ID of the project the consumption list is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public GetBillingConsumptionsInvokeArgs()
        {
        }
        public static new GetBillingConsumptionsInvokeArgs Empty => new GetBillingConsumptionsInvokeArgs();
    }


    [OutputType]
    public sealed class GetBillingConsumptionsResult
    {
        /// <summary>
        /// List of found consumptions
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBillingConsumptionsConsumptionResult> Consumptions;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OrganizationId;
        /// <summary>
        /// The project ID of the consumption.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The last consumption update date.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetBillingConsumptionsResult(
            ImmutableArray<Outputs.GetBillingConsumptionsConsumptionResult> consumptions,

            string id,

            string organizationId,

            string projectId,

            string updatedAt)
        {
            Consumptions = consumptions;
            Id = id;
            OrganizationId = organizationId;
            ProjectId = projectId;
            UpdatedAt = updatedAt;
        }
    }
}
