// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Job.Inputs
{

    public sealed class DefinitionSecretReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An environment variable containing the secret value. Must be specified if `file` is not specified.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// The absolute file path where the secret will be mounted. Must be specified if `environment` is not specified.
        /// </summary>
        [Input("file")]
        public Input<string>? File { get; set; }

        /// <summary>
        /// The secret unique identifier, it could be formatted as region/UUID or UUID. In case the region is passed, it must be the same as the job definition. You could reference the same secret multiple times in the same job definition.
        /// </summary>
        [Input("secretId", required: true)]
        public Input<string> SecretId { get; set; } = null!;

        /// <summary>
        /// The secret reference UUID that is automatically generated by the provider.
        /// </summary>
        [Input("secretReferenceId")]
        public Input<string>? SecretReferenceId { get; set; }

        /// <summary>
        /// The secret version.
        /// </summary>
        [Input("secretVersion")]
        public Input<string>? SecretVersion { get; set; }

        public DefinitionSecretReferenceArgs()
        {
        }
        public static new DefinitionSecretReferenceArgs Empty => new DefinitionSecretReferenceArgs();
    }
}
