// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Outputs
{

    [OutputType]
    public sealed class DomainRecordGeoIp
    {
        /// <summary>
        /// The list of matches
        /// </summary>
        public readonly ImmutableArray<Outputs.DomainRecordGeoIpMatch> Matches;

        [OutputConstructor]
        private DomainRecordGeoIp(ImmutableArray<Outputs.DomainRecordGeoIpMatch> matches)
        {
            Matches = matches;
        }
    }
}
