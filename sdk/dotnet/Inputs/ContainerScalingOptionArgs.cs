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

    public sealed class ContainerScalingOptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Scale depending on the number of concurrent requests being processed per container instance.
        /// </summary>
        [Input("concurrentRequestsThreshold")]
        public Input<int>? ConcurrentRequestsThreshold { get; set; }

        /// <summary>
        /// Scale depending on the CPU usage of a container instance.
        /// </summary>
        [Input("cpuUsageThreshold")]
        public Input<int>? CpuUsageThreshold { get; set; }

        /// <summary>
        /// Scale depending on the memory usage of a container instance.
        /// </summary>
        [Input("memoryUsageThreshold")]
        public Input<int>? MemoryUsageThreshold { get; set; }

        public ContainerScalingOptionArgs()
        {
        }
        public static new ContainerScalingOptionArgs Empty => new ContainerScalingOptionArgs();
    }
}
