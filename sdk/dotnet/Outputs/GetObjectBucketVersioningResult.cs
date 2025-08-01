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
    public sealed class GetObjectBucketVersioningResult
    {
        /// <summary>
        /// Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private GetObjectBucketVersioningResult(bool enabled)
        {
            Enabled = enabled;
        }
    }
}
