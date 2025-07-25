// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Databases.Inputs
{

    public sealed class AclAclRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A text describing this rule. Default description: `IP allowed`
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The IP range to whitelist in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        public AclAclRuleGetArgs()
        {
        }
        public static new AclAclRuleGetArgs Empty => new AclAclRuleGetArgs();
    }
}
