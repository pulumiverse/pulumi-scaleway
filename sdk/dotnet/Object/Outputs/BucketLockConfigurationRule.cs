// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Object.Outputs
{

    [OutputType]
    public sealed class BucketLockConfigurationRule
    {
        /// <summary>
        /// The default retention for the lock.
        /// </summary>
        public readonly Outputs.BucketLockConfigurationRuleDefaultRetention DefaultRetention;

        [OutputConstructor]
        private BucketLockConfigurationRule(Outputs.BucketLockConfigurationRuleDefaultRetention defaultRetention)
        {
            DefaultRetention = defaultRetention;
        }
    }
}
