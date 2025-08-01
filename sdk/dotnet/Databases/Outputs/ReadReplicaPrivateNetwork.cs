// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Databases.Outputs
{

    [OutputType]
    public sealed class ReadReplicaPrivateNetwork
    {
        /// <summary>
        /// If true, the IP network address within the private subnet is determined by the IP Address Management (IPAM) service.
        /// 
        /// &gt; **Important:** One of `service_ip` or `enable_ipam=true` must be set.
        /// </summary>
        public readonly bool? EnableIpam;
        /// <summary>
        /// The ID of the endpoint of the Read Replica.
        /// </summary>
        public readonly string? EndpointId;
        /// <summary>
        /// Hostname of the endpoint. Only one of IP and hostname may be set.
        /// </summary>
        public readonly string? Hostname;
        /// <summary>
        /// IPv4 address of the endpoint (IP address). Only one of IP and hostname may be set.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// Name of the endpoint.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// TCP port of the endpoint.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// UUID of the Private Netork to be connected to the Read Replica.
        /// </summary>
        public readonly string PrivateNetworkId;
        /// <summary>
        /// The IP network address within the private subnet. This must be an IPv4 address with a CIDR notation. If not set, The IP network address within the private subnet is determined by the IP Address Management (IPAM) service.
        /// </summary>
        public readonly string? ServiceIp;
        /// <summary>
        /// Private network zone
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private ReadReplicaPrivateNetwork(
            bool? enableIpam,

            string? endpointId,

            string? hostname,

            string? ip,

            string? name,

            int? port,

            string privateNetworkId,

            string? serviceIp,

            string? zone)
        {
            EnableIpam = enableIpam;
            EndpointId = endpointId;
            Hostname = hostname;
            Ip = ip;
            Name = name;
            Port = port;
            PrivateNetworkId = privateNetworkId;
            ServiceIp = serviceIp;
            Zone = zone;
        }
    }
}
