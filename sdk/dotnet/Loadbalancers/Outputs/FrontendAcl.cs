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
    public sealed class FrontendAcl
    {
        /// <summary>
        /// Action to undertake when an ACL filter matches.
        /// </summary>
        public readonly Outputs.FrontendAclAction Action;
        /// <summary>
        /// The date and time the frontend was created.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// Description of the ACL
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The ACL match rule. At least `ip_subnet` or `ips_edge_services` or `http_filter` and `http_filter_value` are required.
        /// </summary>
        public readonly Outputs.FrontendAclMatch Match;
        /// <summary>
        /// The ACL name. If not provided it will be randomly generated.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The date and time the frontend resource was updated.
        /// </summary>
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private FrontendAcl(
            Outputs.FrontendAclAction action,

            string? createdAt,

            string? description,

            Outputs.FrontendAclMatch match,

            string? name,

            string? updatedAt)
        {
            Action = action;
            CreatedAt = createdAt;
            Description = description;
            Match = match;
            Name = name;
            UpdatedAt = updatedAt;
        }
    }
}
