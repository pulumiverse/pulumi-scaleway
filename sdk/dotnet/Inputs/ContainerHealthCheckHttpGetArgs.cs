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

    public sealed class ContainerHealthCheckHttpGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path to use for the HTTP health check.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public ContainerHealthCheckHttpGetArgs()
        {
        }
        public static new ContainerHealthCheckHttpGetArgs Empty => new ContainerHealthCheckHttpGetArgs();
    }
}
