// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Outputs
{

    [OutputType]
    public sealed class DocumentdbPrivateNetworkEndpointPrivateNetwork
    {
        /// <summary>
        /// The hostname of your endpoint.
        /// </summary>
        public readonly string? Hostname;
        /// <summary>
        /// The private network ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The IP of your private network service.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// The IP network address within the private subnet. This must be an IPv4 address with a CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM) service if not set.
        /// </summary>
        public readonly string? IpNet;
        /// <summary>
        /// The name of your private service.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The port of your private service.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// The zone of your endpoint.
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private DocumentdbPrivateNetworkEndpointPrivateNetwork(
            string? hostname,

            string id,

            string? ip,

            string? ipNet,

            string? name,

            int? port,

            string? zone)
        {
            Hostname = hostname;
            Id = id;
            Ip = ip;
            IpNet = ipNet;
            Name = name;
            Port = port;
            Zone = zone;
        }
    }
}