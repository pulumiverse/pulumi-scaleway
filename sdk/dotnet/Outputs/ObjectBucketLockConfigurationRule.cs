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
    public sealed class ObjectBucketLockConfigurationRule
    {
        /// <summary>
        /// The default retention for the lock.
        /// </summary>
        public readonly Outputs.ObjectBucketLockConfigurationRuleDefaultRetention DefaultRetention;

        [OutputConstructor]
        private ObjectBucketLockConfigurationRule(Outputs.ObjectBucketLockConfigurationRuleDefaultRetention defaultRetention)
        {
            DefaultRetention = defaultRetention;
        }
    }
}
