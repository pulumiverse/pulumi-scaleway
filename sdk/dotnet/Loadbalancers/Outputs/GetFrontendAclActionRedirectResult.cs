// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Loadbalancers.Outputs
{

    [OutputType]
    public sealed class GetFrontendAclActionRedirectResult
    {
        /// <summary>
        /// The HTTP redirect code to use
        /// </summary>
        public readonly int Code;
        /// <summary>
        /// An URL can be used in case of a location redirect
        /// </summary>
        public readonly string Target;
        /// <summary>
        /// The redirect type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetFrontendAclActionRedirectResult(
            int code,

            string target,

            string type)
        {
            Code = code;
            Target = target;
            Type = type;
        }
    }
}
