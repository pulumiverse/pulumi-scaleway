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
    public sealed class GetIotDeviceMessageFilterSubscribeResult
    {
        /// <summary>
        /// Subscribe message filter policy
        /// </summary>
        public readonly string Policy;
        /// <summary>
        /// List of topics in the set
        /// </summary>
        public readonly ImmutableArray<string> Topics;

        [OutputConstructor]
        private GetIotDeviceMessageFilterSubscribeResult(
            string policy,

            ImmutableArray<string> topics)
        {
            Policy = policy;
            Topics = topics;
        }
    }
}