// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Hosting.Outputs
{

    [OutputType]
    public sealed class HostingCpanelUrl
    {
        /// <summary>
        /// The URL of the Dashboard.
        /// </summary>
        public readonly string? Dashboard;
        /// <summary>
        /// The URL of the Webmail interface.
        /// </summary>
        public readonly string? Webmail;

        [OutputConstructor]
        private HostingCpanelUrl(
            string? dashboard,

            string? webmail)
        {
            Dashboard = dashboard;
            Webmail = webmail;
        }
    }
}
