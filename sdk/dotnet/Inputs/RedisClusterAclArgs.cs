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

    public sealed class RedisClusterAclArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A text describing this rule. Default description: `Allow IP`
        /// 
        /// &gt; The `acl` conflict with `private_network`. Only one should be specified.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the IPv4 address resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The IP range to whitelist
        /// in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        public RedisClusterAclArgs()
        {
        }
        public static new RedisClusterAclArgs Empty => new RedisClusterAclArgs();
    }
}
