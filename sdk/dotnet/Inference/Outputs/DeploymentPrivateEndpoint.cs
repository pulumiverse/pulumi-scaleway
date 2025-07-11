// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Inference.Outputs
{

    [OutputType]
    public sealed class DeploymentPrivateEndpoint
    {
        /// <summary>
        /// Disable the authentication on the endpoint.
        /// </summary>
        public readonly bool? DisableAuth;
        /// <summary>
        /// (Optional) The id of the public endpoint.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The ID of the private network to use.
        /// </summary>
        public readonly string? PrivateNetworkId;
        /// <summary>
        /// (Optional) The URL of the endpoint.
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private DeploymentPrivateEndpoint(
            bool? disableAuth,

            string? id,

            string? privateNetworkId,

            string? url)
        {
            DisableAuth = disableAuth;
            Id = id;
            PrivateNetworkId = privateNetworkId;
            Url = url;
        }
    }
}
