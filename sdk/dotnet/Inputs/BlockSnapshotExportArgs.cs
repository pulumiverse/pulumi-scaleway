// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Inputs
{

    public sealed class BlockSnapshotExportArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket where the QCOW file will be saved.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The desired key (path) for the QCOW file within the bucket.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        public BlockSnapshotExportArgs()
        {
        }
        public static new BlockSnapshotExportArgs Empty => new BlockSnapshotExportArgs();
    }
}
