// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Inputs
{

    public sealed class VpcGatewayNetworkIpamConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Use this IPAM-booked IP ID as the Gateway's IP in this Private Network.
        /// </summary>
        [Input("ipamIpId")]
        public Input<string>? IpamIpId { get; set; }

        /// <summary>
        /// Defines whether to enable the default route on the GatewayNetwork.
        /// </summary>
        [Input("pushDefaultRoute")]
        public Input<bool>? PushDefaultRoute { get; set; }

        public VpcGatewayNetworkIpamConfigGetArgs()
        {
        }
        public static new VpcGatewayNetworkIpamConfigGetArgs Empty => new VpcGatewayNetworkIpamConfigGetArgs();
    }
}