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
    [Obsolete(@"scaleway.index/getlbfrontend.getLbFrontend has been deprecated in favor of scaleway.loadbalancers/getfrontend.getFrontend")]
    public static class GetLbFrontend
    {
        /// <summary>
        /// Get information about Scaleway Load Balancer frontends.
        /// 
        /// For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/reference-content/configuring-frontends/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = Pulumiverse.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ip01 = new Scaleway.Loadbalancers.Ip("ip01");
        /// 
        ///     var lb01 = new Scaleway.Loadbalancers.LoadBalancer("lb01", new()
        ///     {
        ///         IpId = ip01.Id,
        ///         Name = "test-lb",
        ///         Type = "lb-s",
        ///     });
        /// 
        ///     var bkd01 = new Scaleway.Loadbalancers.Backend("bkd01", new()
        ///     {
        ///         LbId = lb01.Id,
        ///         ForwardProtocol = "tcp",
        ///         ForwardPort = 80,
        ///         ProxyProtocol = "none",
        ///     });
        /// 
        ///     var frt01 = new Scaleway.Loadbalancers.Frontend("frt01", new()
        ///     {
        ///         LbId = lb01.Id,
        ///         BackendId = bkd01.Id,
        ///         InboundPort = 80,
        ///     });
        /// 
        ///     var byID = Scaleway.Loadbalancers.GetFrontend.Invoke(new()
        ///     {
        ///         FrontendId = frt01.Id,
        ///     });
        /// 
        ///     var byName = Scaleway.Loadbalancers.GetFrontend.Invoke(new()
        ///     {
        ///         Name = frt01.Name,
        ///         LbId = lb01.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetLbFrontendResult> InvokeAsync(GetLbFrontendArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLbFrontendResult>("scaleway:index/getLbFrontend:getLbFrontend", args ?? new GetLbFrontendArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about Scaleway Load Balancer frontends.
        /// 
        /// For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/reference-content/configuring-frontends/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = Pulumiverse.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ip01 = new Scaleway.Loadbalancers.Ip("ip01");
        /// 
        ///     var lb01 = new Scaleway.Loadbalancers.LoadBalancer("lb01", new()
        ///     {
        ///         IpId = ip01.Id,
        ///         Name = "test-lb",
        ///         Type = "lb-s",
        ///     });
        /// 
        ///     var bkd01 = new Scaleway.Loadbalancers.Backend("bkd01", new()
        ///     {
        ///         LbId = lb01.Id,
        ///         ForwardProtocol = "tcp",
        ///         ForwardPort = 80,
        ///         ProxyProtocol = "none",
        ///     });
        /// 
        ///     var frt01 = new Scaleway.Loadbalancers.Frontend("frt01", new()
        ///     {
        ///         LbId = lb01.Id,
        ///         BackendId = bkd01.Id,
        ///         InboundPort = 80,
        ///     });
        /// 
        ///     var byID = Scaleway.Loadbalancers.GetFrontend.Invoke(new()
        ///     {
        ///         FrontendId = frt01.Id,
        ///     });
        /// 
        ///     var byName = Scaleway.Loadbalancers.GetFrontend.Invoke(new()
        ///     {
        ///         Name = frt01.Name,
        ///         LbId = lb01.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetLbFrontendResult> Invoke(GetLbFrontendInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLbFrontendResult>("scaleway:index/getLbFrontend:getLbFrontend", args ?? new GetLbFrontendInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about Scaleway Load Balancer frontends.
        /// 
        /// For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/reference-content/configuring-frontends/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = Pulumiverse.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ip01 = new Scaleway.Loadbalancers.Ip("ip01");
        /// 
        ///     var lb01 = new Scaleway.Loadbalancers.LoadBalancer("lb01", new()
        ///     {
        ///         IpId = ip01.Id,
        ///         Name = "test-lb",
        ///         Type = "lb-s",
        ///     });
        /// 
        ///     var bkd01 = new Scaleway.Loadbalancers.Backend("bkd01", new()
        ///     {
        ///         LbId = lb01.Id,
        ///         ForwardProtocol = "tcp",
        ///         ForwardPort = 80,
        ///         ProxyProtocol = "none",
        ///     });
        /// 
        ///     var frt01 = new Scaleway.Loadbalancers.Frontend("frt01", new()
        ///     {
        ///         LbId = lb01.Id,
        ///         BackendId = bkd01.Id,
        ///         InboundPort = 80,
        ///     });
        /// 
        ///     var byID = Scaleway.Loadbalancers.GetFrontend.Invoke(new()
        ///     {
        ///         FrontendId = frt01.Id,
        ///     });
        /// 
        ///     var byName = Scaleway.Loadbalancers.GetFrontend.Invoke(new()
        ///     {
        ///         Name = frt01.Name,
        ///         LbId = lb01.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetLbFrontendResult> Invoke(GetLbFrontendInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLbFrontendResult>("scaleway:index/getLbFrontend:getLbFrontend", args ?? new GetLbFrontendInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLbFrontendArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The frontend ID.
        /// - Only one of `name` and `frontend_id` should be specified.
        /// </summary>
        [Input("frontendId")]
        public string? FrontendId { get; set; }

        /// <summary>
        /// The Load Balancer ID this frontend is attached to.
        /// </summary>
        [Input("lbId")]
        public string? LbId { get; set; }

        /// <summary>
        /// The name of the frontend.
        /// - When using the `name` you should specify the `lb-id`
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetLbFrontendArgs()
        {
        }
        public static new GetLbFrontendArgs Empty => new GetLbFrontendArgs();
    }

    public sealed class GetLbFrontendInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The frontend ID.
        /// - Only one of `name` and `frontend_id` should be specified.
        /// </summary>
        [Input("frontendId")]
        public Input<string>? FrontendId { get; set; }

        /// <summary>
        /// The Load Balancer ID this frontend is attached to.
        /// </summary>
        [Input("lbId")]
        public Input<string>? LbId { get; set; }

        /// <summary>
        /// The name of the frontend.
        /// - When using the `name` you should specify the `lb-id`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetLbFrontendInvokeArgs()
        {
        }
        public static new GetLbFrontendInvokeArgs Empty => new GetLbFrontendInvokeArgs();
    }


    [OutputType]
    public sealed class GetLbFrontendResult
    {
        public readonly ImmutableArray<Outputs.GetLbFrontendAclResult> Acls;
        public readonly string BackendId;
        public readonly string CertificateId;
        public readonly ImmutableArray<string> CertificateIds;
        public readonly int ConnectionRateLimit;
        public readonly string CreatedAt;
        public readonly bool EnableAccessLogs;
        public readonly bool EnableHttp3;
        public readonly bool ExternalAcls;
        public readonly string? FrontendId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int InboundPort;
        public readonly string? LbId;
        public readonly string? Name;
        public readonly string TimeoutClient;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetLbFrontendResult(
            ImmutableArray<Outputs.GetLbFrontendAclResult> acls,

            string backendId,

            string certificateId,

            ImmutableArray<string> certificateIds,

            int connectionRateLimit,

            string createdAt,

            bool enableAccessLogs,

            bool enableHttp3,

            bool externalAcls,

            string? frontendId,

            string id,

            int inboundPort,

            string? lbId,

            string? name,

            string timeoutClient,

            string updatedAt)
        {
            Acls = acls;
            BackendId = backendId;
            CertificateId = certificateId;
            CertificateIds = certificateIds;
            ConnectionRateLimit = connectionRateLimit;
            CreatedAt = createdAt;
            EnableAccessLogs = enableAccessLogs;
            EnableHttp3 = enableHttp3;
            ExternalAcls = externalAcls;
            FrontendId = frontendId;
            Id = id;
            InboundPort = inboundPort;
            LbId = lbId;
            Name = name;
            TimeoutClient = timeoutClient;
            UpdatedAt = updatedAt;
        }
    }
}
