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

    public sealed class EdgeServicesBackendStageLbBackendConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Load Balancer config.
        /// </summary>
        [Input("lbConfig")]
        public Input<Inputs.EdgeServicesBackendStageLbBackendConfigLbConfigGetArgs>? LbConfig { get; set; }

        public EdgeServicesBackendStageLbBackendConfigGetArgs()
        {
        }
        public static new EdgeServicesBackendStageLbBackendConfigGetArgs Empty => new EdgeServicesBackendStageLbBackendConfigGetArgs();
    }
}
