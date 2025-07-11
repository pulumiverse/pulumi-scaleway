// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Autoscaling.Outputs
{

    [OutputType]
    public sealed class InstanceGroupLoadBalancer
    {
        /// <summary>
        /// The Load Balancer backend IDs.
        /// </summary>
        public readonly ImmutableArray<string> BackendIds;
        /// <summary>
        /// The ID of the Load Balancer.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The ID of the Private Network attached to the Load Balancer.
        /// </summary>
        public readonly string? PrivateNetworkId;

        [OutputConstructor]
        private InstanceGroupLoadBalancer(
            ImmutableArray<string> backendIds,

            string? id,

            string? privateNetworkId)
        {
            BackendIds = backendIds;
            Id = id;
            PrivateNetworkId = privateNetworkId;
        }
    }
}
