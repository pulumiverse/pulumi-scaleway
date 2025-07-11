// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Autoscaling.Outputs
{

    [OutputType]
    public sealed class InstanceTemplateVolumeFromEmpty
    {
        /// <summary>
        /// Size in GB of the new empty volume
        /// </summary>
        public readonly int Size;

        [OutputConstructor]
        private InstanceTemplateVolumeFromEmpty(int size)
        {
            Size = size;
        }
    }
}
