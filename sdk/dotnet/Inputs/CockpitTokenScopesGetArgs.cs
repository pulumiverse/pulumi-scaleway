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

    public sealed class CockpitTokenScopesGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Permission to query logs.
        /// </summary>
        [Input("queryLogs")]
        public Input<bool>? QueryLogs { get; set; }

        /// <summary>
        /// Permission to query metrics.
        /// </summary>
        [Input("queryMetrics")]
        public Input<bool>? QueryMetrics { get; set; }

        /// <summary>
        /// Permission to query traces.
        /// </summary>
        [Input("queryTraces")]
        public Input<bool>? QueryTraces { get; set; }

        /// <summary>
        /// Permission to set up alerts.
        /// </summary>
        [Input("setupAlerts")]
        public Input<bool>? SetupAlerts { get; set; }

        /// <summary>
        /// Permission to set up logs rules.
        /// </summary>
        [Input("setupLogsRules")]
        public Input<bool>? SetupLogsRules { get; set; }

        /// <summary>
        /// Permission to set up metrics rules.
        /// </summary>
        [Input("setupMetricsRules")]
        public Input<bool>? SetupMetricsRules { get; set; }

        /// <summary>
        /// Permission to write logs.
        /// </summary>
        [Input("writeLogs")]
        public Input<bool>? WriteLogs { get; set; }

        /// <summary>
        /// Permission to write metrics.
        /// </summary>
        [Input("writeMetrics")]
        public Input<bool>? WriteMetrics { get; set; }

        /// <summary>
        /// Permission to write traces.
        /// </summary>
        [Input("writeTraces")]
        public Input<bool>? WriteTraces { get; set; }

        public CockpitTokenScopesGetArgs()
        {
        }
        public static new CockpitTokenScopesGetArgs Empty => new CockpitTokenScopesGetArgs();
    }
}