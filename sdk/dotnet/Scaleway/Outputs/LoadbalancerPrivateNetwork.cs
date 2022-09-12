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
    public sealed class LoadbalancerPrivateNetwork
    {
        /// <summary>
        /// (Optional) Set to true if you want to let DHCP assign IP addresses. See below.
        /// </summary>
        public readonly bool? DhcpConfig;
        /// <summary>
        /// (Required) The ID of the Private Network to associate.
        /// </summary>
        public readonly string PrivateNetworkId;
        /// <summary>
        /// (Optional) Define two local ip address of your choice for each load balancer instance. See below.
        /// </summary>
        public readonly ImmutableArray<string> StaticConfigs;
        public readonly string? Status;
        /// <summary>
        /// `zone`) The zone in which the IP should be reserved.
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private LoadbalancerPrivateNetwork(
            bool? dhcpConfig,

            string privateNetworkId,

            ImmutableArray<string> staticConfigs,

            string? status,

            string? zone)
        {
            DhcpConfig = dhcpConfig;
            PrivateNetworkId = privateNetworkId;
            StaticConfigs = staticConfigs;
            Status = status;
            Zone = zone;
        }
    }
}