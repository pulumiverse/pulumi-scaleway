// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Inputs
{

    public sealed class InstanceSnapshotImportArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bucket name containing [qcow2](https://en.wikipedia.org/wiki/Qcow) to import
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Key of the object to import
        /// 
        /// &gt; **Note:** The type `unified` could be instantiated on both `l_ssd` and `b_ssd` volumes.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        public InstanceSnapshotImportArgs()
        {
        }
        public static new InstanceSnapshotImportArgs Empty => new InstanceSnapshotImportArgs();
    }
}