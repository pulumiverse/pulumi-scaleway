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

    public sealed class WebhostingOptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The option ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The option name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public WebhostingOptionArgs()
        {
        }
        public static new WebhostingOptionArgs Empty => new WebhostingOptionArgs();
    }
}
