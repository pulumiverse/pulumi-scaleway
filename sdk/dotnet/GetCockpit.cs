// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway
{
    public static class GetCockpit
    {
        /// <summary>
        /// &gt; **Important:**  The data source `scaleway.Cockpit` has been deprecated and will no longer be supported. Instead, use resource `scaleway.Cockpit`.
        /// 
        /// &gt; **Note:**
        /// As of April 2024, Cockpit has introduced [regionalization](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#region) to offer more flexibility and resilience.
        /// If you have created customized dashboards with data for your Scaleway resources before April 2024, you will need to update your queries in Grafana, with the new regionalized data sources.
        /// 
        /// The `scaleway.Cockpit` data source is used to retrieve information about a Scaleway Cockpit associated with a given Project. This can be the default Project or a specific Project identified by its ID.
        /// 
        /// Refer to Cockpit's [product documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/) and [API documentation](https://www.scaleway.com/en/developers/api/cockpit/regional-api) for more information.
        /// 
        /// ## Retrieve a Cockpit
        /// 
        /// The following commands allow you to:
        /// 
        /// - get information on the Cockpit associated with your Scaleway default Project
        /// - get information on the Cockpit associated with a specific Scaleway Project
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Get the default Project's Cockpit
        ///     var main = Scaleway.GetCockpit.Invoke();
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Get a specific Project's Cockpit
        ///     var main = Scaleway.GetCockpit.Invoke(new()
        ///     {
        ///         ProjectId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetCockpitResult> InvokeAsync(GetCockpitArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCockpitResult>("scaleway:index/getCockpit:getCockpit", args ?? new GetCockpitArgs(), options.WithDefaults());

        /// <summary>
        /// &gt; **Important:**  The data source `scaleway.Cockpit` has been deprecated and will no longer be supported. Instead, use resource `scaleway.Cockpit`.
        /// 
        /// &gt; **Note:**
        /// As of April 2024, Cockpit has introduced [regionalization](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#region) to offer more flexibility and resilience.
        /// If you have created customized dashboards with data for your Scaleway resources before April 2024, you will need to update your queries in Grafana, with the new regionalized data sources.
        /// 
        /// The `scaleway.Cockpit` data source is used to retrieve information about a Scaleway Cockpit associated with a given Project. This can be the default Project or a specific Project identified by its ID.
        /// 
        /// Refer to Cockpit's [product documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/) and [API documentation](https://www.scaleway.com/en/developers/api/cockpit/regional-api) for more information.
        /// 
        /// ## Retrieve a Cockpit
        /// 
        /// The following commands allow you to:
        /// 
        /// - get information on the Cockpit associated with your Scaleway default Project
        /// - get information on the Cockpit associated with a specific Scaleway Project
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Get the default Project's Cockpit
        ///     var main = Scaleway.GetCockpit.Invoke();
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Get a specific Project's Cockpit
        ///     var main = Scaleway.GetCockpit.Invoke(new()
        ///     {
        ///         ProjectId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCockpitResult> Invoke(GetCockpitInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCockpitResult>("scaleway:index/getCockpit:getCockpit", args ?? new GetCockpitInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCockpitArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the ID of the Scaleway Project that the Cockpit is associated with. If not specified, it defaults to the Project ID specified in the provider configuration.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        public GetCockpitArgs()
        {
        }
        public static new GetCockpitArgs Empty => new GetCockpitArgs();
    }

    public sealed class GetCockpitInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the ID of the Scaleway Project that the Cockpit is associated with. If not specified, it defaults to the Project ID specified in the provider configuration.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public GetCockpitInvokeArgs()
        {
        }
        public static new GetCockpitInvokeArgs Empty => new GetCockpitInvokeArgs();
    }


    [OutputType]
    public sealed class GetCockpitResult
    {
        /// <summary>
        /// (Deprecated) A list of [endpoints](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#endpoints) related to Cockpit, each with specific URLs:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCockpitEndpointResult> Endpoints;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Plan;
        /// <summary>
        /// (Deprecated) ID of the current pricing plan
        /// </summary>
        public readonly string PlanId;
        public readonly string? ProjectId;
        public readonly ImmutableArray<Outputs.GetCockpitPushUrlResult> PushUrls;

        [OutputConstructor]
        private GetCockpitResult(
            ImmutableArray<Outputs.GetCockpitEndpointResult> endpoints,

            string id,

            string plan,

            string planId,

            string? projectId,

            ImmutableArray<Outputs.GetCockpitPushUrlResult> pushUrls)
        {
            Endpoints = endpoints;
            Id = id;
            Plan = plan;
            PlanId = planId;
            ProjectId = projectId;
            PushUrls = pushUrls;
        }
    }
}