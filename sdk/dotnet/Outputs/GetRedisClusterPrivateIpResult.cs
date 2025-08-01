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
    public sealed class GetRedisClusterPrivateIpResult
    {
        /// <summary>
        /// The private IPv4 address
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// The ID of the Redis cluster.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetRedisClusterPrivateIpResult(
            string address,

            string id)
        {
            Address = address;
            Id = id;
        }
    }
}
