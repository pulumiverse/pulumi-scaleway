// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Loadbalancers
{
    public static class GetRoutes
    {
        /// <summary>
        /// Gets information about multiple Load Balancer routes.
        /// 
        /// For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/how-to/create-manage-routes/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-route).
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Find routes that share the same frontend ID
        ///     var byFrontendID = Scaleway.Loadbalancers.GetRoutes.Invoke(new()
        ///     {
        ///         FrontendId = frt01.Id,
        ///     });
        /// 
        ///     // Find routes by frontend ID and zone
        ///     var myKey = Scaleway.Loadbalancers.GetRoutes.Invoke(new()
        ///     {
        ///         FrontendId = "11111111-1111-1111-1111-111111111111",
        ///         Zone = "fr-par-2",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRoutesResult> InvokeAsync(GetRoutesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRoutesResult>("scaleway:loadbalancers/getRoutes:getRoutes", args ?? new GetRoutesArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about multiple Load Balancer routes.
        /// 
        /// For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/how-to/create-manage-routes/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-route).
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Find routes that share the same frontend ID
        ///     var byFrontendID = Scaleway.Loadbalancers.GetRoutes.Invoke(new()
        ///     {
        ///         FrontendId = frt01.Id,
        ///     });
        /// 
        ///     // Find routes by frontend ID and zone
        ///     var myKey = Scaleway.Loadbalancers.GetRoutes.Invoke(new()
        ///     {
        ///         FrontendId = "11111111-1111-1111-1111-111111111111",
        ///         Zone = "fr-par-2",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRoutesResult> Invoke(GetRoutesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRoutesResult>("scaleway:loadbalancers/getRoutes:getRoutes", args ?? new GetRoutesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about multiple Load Balancer routes.
        /// 
        /// For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/how-to/create-manage-routes/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-route).
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Find routes that share the same frontend ID
        ///     var byFrontendID = Scaleway.Loadbalancers.GetRoutes.Invoke(new()
        ///     {
        ///         FrontendId = frt01.Id,
        ///     });
        /// 
        ///     // Find routes by frontend ID and zone
        ///     var myKey = Scaleway.Loadbalancers.GetRoutes.Invoke(new()
        ///     {
        ///         FrontendId = "11111111-1111-1111-1111-111111111111",
        ///         Zone = "fr-par-2",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRoutesResult> Invoke(GetRoutesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRoutesResult>("scaleway:loadbalancers/getRoutes:getRoutes", args ?? new GetRoutesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRoutesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The frontend ID (the origin of the redirection), to filter for. Routes with a matching frontend ID are listed.
        /// </summary>
        [Input("frontendId")]
        public string? FrontendId { get; set; }

        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the routes exist.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetRoutesArgs()
        {
        }
        public static new GetRoutesArgs Empty => new GetRoutesArgs();
    }

    public sealed class GetRoutesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The frontend ID (the origin of the redirection), to filter for. Routes with a matching frontend ID are listed.
        /// </summary>
        [Input("frontendId")]
        public Input<string>? FrontendId { get; set; }

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the routes exist.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetRoutesInvokeArgs()
        {
        }
        public static new GetRoutesInvokeArgs Empty => new GetRoutesInvokeArgs();
    }


    [OutputType]
    public sealed class GetRoutesResult
    {
        public readonly string? FrontendId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OrganizationId;
        public readonly string ProjectId;
        /// <summary>
        /// List of retrieved routes
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRoutesRouteResult> Routes;
        public readonly string Zone;

        [OutputConstructor]
        private GetRoutesResult(
            string? frontendId,

            string id,

            string organizationId,

            string projectId,

            ImmutableArray<Outputs.GetRoutesRouteResult> routes,

            string zone)
        {
            FrontendId = frontendId;
            Id = id;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Routes = routes;
            Zone = zone;
        }
    }
}
