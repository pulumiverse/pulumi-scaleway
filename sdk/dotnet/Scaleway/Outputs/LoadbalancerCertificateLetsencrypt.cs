// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Scaleway.Outputs
{

    [OutputType]
    public sealed class LoadbalancerCertificateLetsencrypt
    {
        /// <summary>
        /// Main domain of the certificate. A new certificate will be created if this field is changed.
        /// </summary>
        public readonly string CommonName;
        /// <summary>
        /// Array of alternative domain names.  A new certificate will be created if this field is changed.
        /// </summary>
        public readonly ImmutableArray<string> SubjectAlternativeNames;

        [OutputConstructor]
        private LoadbalancerCertificateLetsencrypt(
            string commonName,

            ImmutableArray<string> subjectAlternativeNames)
        {
            CommonName = commonName;
            SubjectAlternativeNames = subjectAlternativeNames;
        }
    }
}