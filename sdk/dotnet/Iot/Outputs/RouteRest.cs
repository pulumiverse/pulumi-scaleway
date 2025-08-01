// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Iot.Outputs
{

    [OutputType]
    public sealed class RouteRest
    {
        /// <summary>
        /// a map of the extra headers to send with the HTTP call (e.g. `X-Header = Value`).
        /// </summary>
        public readonly ImmutableDictionary<string, string> Headers;
        /// <summary>
        /// The URI of the Rest endpoint (e.g. `https://internal.mycompany.com/ingest/mqttdata`).
        /// </summary>
        public readonly string Uri;
        /// <summary>
        /// The HTTP Verb used to call Rest URI (e.g. `post`).
        /// </summary>
        public readonly string Verb;

        [OutputConstructor]
        private RouteRest(
            ImmutableDictionary<string, string> headers,

            string uri,

            string verb)
        {
            Headers = headers;
            Uri = uri;
            Verb = verb;
        }
    }
}
