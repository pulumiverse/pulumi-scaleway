// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Loadbalancers.Inputs
{

    public sealed class CertificateCustomCertificateArgs : global::Pulumi.ResourceArgs
    {
        [Input("certificateChain", required: true)]
        private Input<string>? _certificateChain;

        /// <summary>
        /// The full PEM-formatted certificate chain
        /// </summary>
        public Input<string>? CertificateChain
        {
            get => _certificateChain;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _certificateChain = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public CertificateCustomCertificateArgs()
        {
        }
        public static new CertificateCustomCertificateArgs Empty => new CertificateCustomCertificateArgs();
    }
}
