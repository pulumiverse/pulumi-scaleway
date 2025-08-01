// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    public sealed class RegistrationTechnicalContactExtensionFrAssociationInfo
    {
        /// <summary>
        /// Publication date in the Official Journal (RFC3339 format) for association information.
        /// </summary>
        public readonly string? PublicationJo;
        /// <summary>
        /// Page number of the publication in the Official Journal for association information.
        /// </summary>
        public readonly int? PublicationJoPage;

        [OutputConstructor]
        private RegistrationTechnicalContactExtensionFrAssociationInfo(
            string? publicationJo,

            int? publicationJoPage)
        {
            PublicationJo = publicationJo;
            PublicationJoPage = publicationJoPage;
        }
    }
}
