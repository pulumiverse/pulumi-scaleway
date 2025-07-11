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
    public static class GetIp
    {
        /// <summary>
        /// Gets information about a Load Balancer IP address.
        /// 
        /// For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/how-to/create-manage-flex-ips/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-ip-addresses-list-ip-addresses).
        /// </summary>
        public static Task<GetIpResult> InvokeAsync(GetIpArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpResult>("scaleway:loadbalancers/getIp:getIp", args ?? new GetIpArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Load Balancer IP address.
        /// 
        /// For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/how-to/create-manage-flex-ips/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-ip-addresses-list-ip-addresses).
        /// </summary>
        public static Output<GetIpResult> Invoke(GetIpInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpResult>("scaleway:loadbalancers/getIp:getIp", args ?? new GetIpInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Load Balancer IP address.
        /// 
        /// For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/how-to/create-manage-flex-ips/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-ip-addresses-list-ip-addresses).
        /// </summary>
        public static Output<GetIpResult> Invoke(GetIpInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpResult>("scaleway:loadbalancers/getIp:getIp", args ?? new GetIpInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IP address.
        /// Only one of `ip_address` and `ip_id` should be specified.
        /// </summary>
        [Input("ipAddress")]
        public string? IpAddress { get; set; }

        /// <summary>
        /// The IP ID.
        /// Only one of `ip_address` and `ip_id` should be specified.
        /// </summary>
        [Input("ipId")]
        public string? IpId { get; set; }

        /// <summary>
        /// The ID of the Project the Load Balancer IP is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the IP was reserved.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetIpArgs()
        {
        }
        public static new GetIpArgs Empty => new GetIpArgs();
    }

    public sealed class GetIpInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IP address.
        /// Only one of `ip_address` and `ip_id` should be specified.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The IP ID.
        /// Only one of `ip_address` and `ip_id` should be specified.
        /// </summary>
        [Input("ipId")]
        public Input<string>? IpId { get; set; }

        /// <summary>
        /// The ID of the Project the Load Balancer IP is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the IP was reserved.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetIpInvokeArgs()
        {
        }
        public static new GetIpInvokeArgs Empty => new GetIpInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IpAddress;
        public readonly string? IpId;
        public readonly bool IsIpv6;
        /// <summary>
        /// The ID of the associated Load Balancer, if any
        /// </summary>
        public readonly string LbId;
        /// <summary>
        /// (Defaults to provider `organization_id`) The ID of the Organization the Load Balancer IP is associated with.
        /// </summary>
        public readonly string OrganizationId;
        public readonly string? ProjectId;
        public readonly string Region;
        /// <summary>
        /// The reverse domain associated with this IP.
        /// </summary>
        public readonly string Reverse;
        /// <summary>
        /// The tags associated with this IP.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        public readonly string? Zone;

        [OutputConstructor]
        private GetIpResult(
            string id,

            string? ipAddress,

            string? ipId,

            bool isIpv6,

            string lbId,

            string organizationId,

            string? projectId,

            string region,

            string reverse,

            ImmutableArray<string> tags,

            string? zone)
        {
            Id = id;
            IpAddress = ipAddress;
            IpId = ipId;
            IsIpv6 = isIpv6;
            LbId = lbId;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Region = region;
            Reverse = reverse;
            Tags = tags;
            Zone = zone;
        }
    }
}
