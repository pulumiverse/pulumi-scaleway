// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Inputs
{

    public sealed class GetIpamIpsResourceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the attached resource.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the attached resource.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The type of the attached resource. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public GetIpamIpsResourceArgs()
        {
        }
        public static new GetIpamIpsResourceArgs Empty => new GetIpamIpsResourceArgs();
    }
}
