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
    public sealed class GetLbAclsAclActionResult
    {
        /// <summary>
        /// Redirect parameters when using an ACL with `redirect` action.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLbAclsAclActionRedirectResult> Redirects;
        /// <summary>
        /// The redirect type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetLbAclsAclActionResult(
            ImmutableArray<Outputs.GetLbAclsAclActionRedirectResult> redirects,

            string type)
        {
            Redirects = redirects;
            Type = type;
        }
    }
}
