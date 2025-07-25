// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Hosting.Inputs
{

    public sealed class HostingCpanelUrlGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL of the Dashboard.
        /// </summary>
        [Input("dashboard")]
        public Input<string>? Dashboard { get; set; }

        /// <summary>
        /// The URL of the Webmail interface.
        /// </summary>
        [Input("webmail")]
        public Input<string>? Webmail { get; set; }

        public HostingCpanelUrlGetArgs()
        {
        }
        public static new HostingCpanelUrlGetArgs Empty => new HostingCpanelUrlGetArgs();
    }
}
