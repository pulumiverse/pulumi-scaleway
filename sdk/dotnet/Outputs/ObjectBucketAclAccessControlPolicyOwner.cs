// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
    public sealed class ObjectBucketAclAccessControlPolicyOwner
    {
        /// <summary>
        /// The project ID of the grantee.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The `region`, `bucket` and `acl` separated by (`/`).
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private ObjectBucketAclAccessControlPolicyOwner(
            string? displayName,

            string id)
        {
            DisplayName = displayName;
            Id = id;
        }
    }
}