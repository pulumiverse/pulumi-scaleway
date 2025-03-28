// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Object.Inputs
{

    public sealed class BucketLockConfigurationRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default retention for the lock.
        /// </summary>
        [Input("defaultRetention", required: true)]
        public Input<Inputs.BucketLockConfigurationRuleDefaultRetentionGetArgs> DefaultRetention { get; set; } = null!;

        public BucketLockConfigurationRuleGetArgs()
        {
        }
        public static new BucketLockConfigurationRuleGetArgs Empty => new BucketLockConfigurationRuleGetArgs();
    }
}
