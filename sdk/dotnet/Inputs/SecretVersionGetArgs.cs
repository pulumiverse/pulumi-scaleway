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

    public sealed class SecretVersionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Date and time of the secret's creation (in RFC 3339 format).
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Description of the secret (e.g. `my-new-description`).
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Returns true if the version is the latest.
        /// </summary>
        [Input("latest")]
        public Input<bool>? Latest { get; set; }

        /// <summary>
        /// The revision of secret version
        /// </summary>
        [Input("revision")]
        public Input<string>? Revision { get; set; }

        /// <summary>
        /// The secret ID associated with this version
        /// </summary>
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        /// <summary>
        /// The status of the secret.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Date and time of the secret's last update (in RFC 3339 format).
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public SecretVersionGetArgs()
        {
        }
        public static new SecretVersionGetArgs Empty => new SecretVersionGetArgs();
    }
}
