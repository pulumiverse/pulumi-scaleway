// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Ipam.Outputs
{

    [OutputType]
    public sealed class GetIpResourceResult
    {
        /// <summary>
        /// The ID of the resource that the IP is attached to.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the resource the IP is attached to.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The type of the resource the IP is attached to. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetIpResourceResult(
            string? id,

            string? name,

            string type)
        {
            Id = id;
            Name = name;
            Type = type;
        }
    }
}
