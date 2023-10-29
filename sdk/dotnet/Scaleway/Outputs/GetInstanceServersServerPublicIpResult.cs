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
    public sealed class GetInstanceServersServerPublicIpResult
    {
        /// <summary>
        /// The address of the IP
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// The ID of the IP
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetInstanceServersServerPublicIpResult(
            string address,

            string id)
        {
            Address = address;
            Id = id;
        }
    }
}