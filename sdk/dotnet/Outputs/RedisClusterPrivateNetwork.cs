// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace lbrlabs.Scaleway.Outputs
{

    [OutputType]
    public sealed class RedisClusterPrivateNetwork
    {
        public readonly string? EndpointId;
        /// <summary>
        /// The UUID of the private network resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Endpoint IPv4 addresses in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation). You must provide at least one IP per node.
        /// </summary>
        public readonly ImmutableArray<string> ServiceIps;
        /// <summary>
        /// `zone`) The zone in which the Redis Cluster should be created.
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private RedisClusterPrivateNetwork(
            string? endpointId,

            string id,

            ImmutableArray<string> serviceIps,

            string? zone)
        {
            EndpointId = endpointId;
            Id = id;
            ServiceIps = serviceIps;
            Zone = zone;
        }
    }
}