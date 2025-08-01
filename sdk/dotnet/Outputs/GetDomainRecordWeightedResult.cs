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
    public sealed class GetDomainRecordWeightedResult
    {
        /// <summary>
        /// The weighted IP
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// The weight of the IP
        /// </summary>
        public readonly int Weight;

        [OutputConstructor]
        private GetDomainRecordWeightedResult(
            string ip,

            int weight)
        {
            Ip = ip;
            Weight = weight;
        }
    }
}
