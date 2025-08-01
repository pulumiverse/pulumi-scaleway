// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    public sealed class DatabaseReadReplicaDirectAccess
    {
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

        [OutputConstructor]
        private DatabaseReadReplicaDirectAccess(
            string? endpointId,

            string? hostname,

            string? ip,

            string? name,

            int? port)
        {
            EndpointId = endpointId;
            Hostname = hostname;
            Ip = ip;
            Name = name;
            Port = port;
        }
    }
}
