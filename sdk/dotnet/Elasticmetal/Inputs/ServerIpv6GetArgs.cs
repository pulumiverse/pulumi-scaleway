// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Elasticmetal.Inputs
{

    public sealed class ServerIpv6GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The address of the IPv6.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The ID of the IPv6.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The reverse of the IPv6.
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        /// <summary>
        /// The type of the IPv6.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ServerIpv6GetArgs()
        {
        }
        public static new ServerIpv6GetArgs Empty => new ServerIpv6GetArgs();
    }
}
