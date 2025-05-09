// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Domain.Outputs
{

    [OutputType]
    public sealed class RegistrationTechnicalContactExtensionEu
    {
        /// <summary>
        /// Indicates the European citizenship of the contact.
        /// </summary>
        public readonly string? EuropeanCitizenship;

        [OutputConstructor]
        private RegistrationTechnicalContactExtensionEu(string? europeanCitizenship)
        {
            EuropeanCitizenship = europeanCitizenship;
        }
    }
}
