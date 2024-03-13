// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway
{
    public static class GetVpcPublicGatewayDhcpReservation
    {
        /// <summary>
        /// Gets information about a dhcp entries. For further information please check the
        /// API [documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#dhcp-entries-e40fb6)
        /// 
        /// ## Example Dynamic
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = Pulumiverse.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mainVpcPrivateNetwork = new Scaleway.VpcPrivateNetwork("mainVpcPrivateNetwork");
        /// 
        ///     var mainInstanceServer = new Scaleway.InstanceServer("mainInstanceServer", new()
        ///     {
        ///         Image = "ubuntu_jammy",
        ///         Type = "DEV1-S",
        ///         Zone = "fr-par-1",
        ///     });
        /// 
        ///     var mainInstancePrivateNic = new Scaleway.InstancePrivateNic("mainInstancePrivateNic", new()
        ///     {
        ///         ServerId = mainInstanceServer.Id,
        ///         PrivateNetworkId = mainVpcPrivateNetwork.Id,
        ///     });
        /// 
        ///     var mainVpcPublicGatewayIp = new Scaleway.VpcPublicGatewayIp("mainVpcPublicGatewayIp");
        /// 
        ///     var mainVpcPublicGatewayDhcp = new Scaleway.VpcPublicGatewayDhcp("mainVpcPublicGatewayDhcp", new()
        ///     {
        ///         Subnet = "192.168.1.0/24",
        ///     });
        /// 
        ///     var mainVpcPublicGateway = new Scaleway.VpcPublicGateway("mainVpcPublicGateway", new()
        ///     {
        ///         Type = "VPC-GW-S",
        ///         IpId = mainVpcPublicGatewayIp.Id,
        ///     });
        /// 
        ///     var mainVpcGatewayNetwork = new Scaleway.VpcGatewayNetwork("mainVpcGatewayNetwork", new()
        ///     {
        ///         GatewayId = mainVpcPublicGateway.Id,
        ///         PrivateNetworkId = mainVpcPrivateNetwork.Id,
        ///         DhcpId = mainVpcPublicGatewayDhcp.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     //# Retrieve the dynamic entries generated by mac address &amp; gateway network
        ///     var byMacAddressAndGwNetwork = Scaleway.GetVpcPublicGatewayDhcpReservation.Invoke(new()
        ///     {
        ///         MacAddress = mainInstancePrivateNic.MacAddress,
        ///         GatewayNetworkId = mainVpcGatewayNetwork.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// ## Example Static and PAT rule
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = Pulumiverse.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mainVpcPrivateNetwork = new Scaleway.VpcPrivateNetwork("mainVpcPrivateNetwork");
        /// 
        ///     var mainInstanceSecurityGroup = new Scaleway.InstanceSecurityGroup("mainInstanceSecurityGroup", new()
        ///     {
        ///         InboundDefaultPolicy = "drop",
        ///         OutboundDefaultPolicy = "accept",
        ///         InboundRules = new[]
        ///         {
        ///             new Scaleway.Inputs.InstanceSecurityGroupInboundRuleArgs
        ///             {
        ///                 Action = "accept",
        ///                 Port = 22,
        ///             },
        ///         },
        ///     });
        /// 
        ///     var mainInstanceServer = new Scaleway.InstanceServer("mainInstanceServer", new()
        ///     {
        ///         Image = "ubuntu_jammy",
        ///         Type = "DEV1-S",
        ///         Zone = "fr-par-1",
        ///         SecurityGroupId = mainInstanceSecurityGroup.Id,
        ///     });
        /// 
        ///     var mainInstancePrivateNic = new Scaleway.InstancePrivateNic("mainInstancePrivateNic", new()
        ///     {
        ///         ServerId = mainInstanceServer.Id,
        ///         PrivateNetworkId = mainVpcPrivateNetwork.Id,
        ///     });
        /// 
        ///     var mainVpcPublicGatewayIp = new Scaleway.VpcPublicGatewayIp("mainVpcPublicGatewayIp");
        /// 
        ///     var mainVpcPublicGatewayDhcp = new Scaleway.VpcPublicGatewayDhcp("mainVpcPublicGatewayDhcp", new()
        ///     {
        ///         Subnet = "192.168.1.0/24",
        ///     });
        /// 
        ///     var mainVpcPublicGateway = new Scaleway.VpcPublicGateway("mainVpcPublicGateway", new()
        ///     {
        ///         Type = "VPC-GW-S",
        ///         IpId = mainVpcPublicGatewayIp.Id,
        ///     });
        /// 
        ///     var mainVpcGatewayNetwork = new Scaleway.VpcGatewayNetwork("mainVpcGatewayNetwork", new()
        ///     {
        ///         GatewayId = mainVpcPublicGateway.Id,
        ///         PrivateNetworkId = mainVpcPrivateNetwork.Id,
        ///         DhcpId = mainVpcPublicGatewayDhcp.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     var mainVpcPublicGatewayDhcpReservation = new Scaleway.VpcPublicGatewayDhcpReservation("mainVpcPublicGatewayDhcpReservation", new()
        ///     {
        ///         GatewayNetworkId = mainVpcGatewayNetwork.Id,
        ///         MacAddress = mainInstancePrivateNic.MacAddress,
        ///         IpAddress = "192.168.1.4",
        ///     });
        /// 
        ///     //## VPC PAT RULE
        ///     var mainVpcPublicGatewayPatRule = new Scaleway.VpcPublicGatewayPatRule("mainVpcPublicGatewayPatRule", new()
        ///     {
        ///         GatewayId = mainVpcPublicGateway.Id,
        ///         PrivateIp = mainVpcPublicGatewayDhcpReservation.IpAddress,
        ///         PrivatePort = 22,
        ///         PublicPort = 2222,
        ///         Protocol = "tcp",
        ///     });
        /// 
        ///     var byId = Scaleway.GetVpcPublicGatewayDhcpReservation.Invoke(new()
        ///     {
        ///         ReservationId = mainVpcPublicGatewayDhcpReservation.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetVpcPublicGatewayDhcpReservationResult> InvokeAsync(GetVpcPublicGatewayDhcpReservationArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcPublicGatewayDhcpReservationResult>("scaleway:index/getVpcPublicGatewayDhcpReservation:getVpcPublicGatewayDhcpReservation", args ?? new GetVpcPublicGatewayDhcpReservationArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a dhcp entries. For further information please check the
        /// API [documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#dhcp-entries-e40fb6)
        /// 
        /// ## Example Dynamic
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = Pulumiverse.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mainVpcPrivateNetwork = new Scaleway.VpcPrivateNetwork("mainVpcPrivateNetwork");
        /// 
        ///     var mainInstanceServer = new Scaleway.InstanceServer("mainInstanceServer", new()
        ///     {
        ///         Image = "ubuntu_jammy",
        ///         Type = "DEV1-S",
        ///         Zone = "fr-par-1",
        ///     });
        /// 
        ///     var mainInstancePrivateNic = new Scaleway.InstancePrivateNic("mainInstancePrivateNic", new()
        ///     {
        ///         ServerId = mainInstanceServer.Id,
        ///         PrivateNetworkId = mainVpcPrivateNetwork.Id,
        ///     });
        /// 
        ///     var mainVpcPublicGatewayIp = new Scaleway.VpcPublicGatewayIp("mainVpcPublicGatewayIp");
        /// 
        ///     var mainVpcPublicGatewayDhcp = new Scaleway.VpcPublicGatewayDhcp("mainVpcPublicGatewayDhcp", new()
        ///     {
        ///         Subnet = "192.168.1.0/24",
        ///     });
        /// 
        ///     var mainVpcPublicGateway = new Scaleway.VpcPublicGateway("mainVpcPublicGateway", new()
        ///     {
        ///         Type = "VPC-GW-S",
        ///         IpId = mainVpcPublicGatewayIp.Id,
        ///     });
        /// 
        ///     var mainVpcGatewayNetwork = new Scaleway.VpcGatewayNetwork("mainVpcGatewayNetwork", new()
        ///     {
        ///         GatewayId = mainVpcPublicGateway.Id,
        ///         PrivateNetworkId = mainVpcPrivateNetwork.Id,
        ///         DhcpId = mainVpcPublicGatewayDhcp.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     //# Retrieve the dynamic entries generated by mac address &amp; gateway network
        ///     var byMacAddressAndGwNetwork = Scaleway.GetVpcPublicGatewayDhcpReservation.Invoke(new()
        ///     {
        ///         MacAddress = mainInstancePrivateNic.MacAddress,
        ///         GatewayNetworkId = mainVpcGatewayNetwork.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// ## Example Static and PAT rule
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = Pulumiverse.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mainVpcPrivateNetwork = new Scaleway.VpcPrivateNetwork("mainVpcPrivateNetwork");
        /// 
        ///     var mainInstanceSecurityGroup = new Scaleway.InstanceSecurityGroup("mainInstanceSecurityGroup", new()
        ///     {
        ///         InboundDefaultPolicy = "drop",
        ///         OutboundDefaultPolicy = "accept",
        ///         InboundRules = new[]
        ///         {
        ///             new Scaleway.Inputs.InstanceSecurityGroupInboundRuleArgs
        ///             {
        ///                 Action = "accept",
        ///                 Port = 22,
        ///             },
        ///         },
        ///     });
        /// 
        ///     var mainInstanceServer = new Scaleway.InstanceServer("mainInstanceServer", new()
        ///     {
        ///         Image = "ubuntu_jammy",
        ///         Type = "DEV1-S",
        ///         Zone = "fr-par-1",
        ///         SecurityGroupId = mainInstanceSecurityGroup.Id,
        ///     });
        /// 
        ///     var mainInstancePrivateNic = new Scaleway.InstancePrivateNic("mainInstancePrivateNic", new()
        ///     {
        ///         ServerId = mainInstanceServer.Id,
        ///         PrivateNetworkId = mainVpcPrivateNetwork.Id,
        ///     });
        /// 
        ///     var mainVpcPublicGatewayIp = new Scaleway.VpcPublicGatewayIp("mainVpcPublicGatewayIp");
        /// 
        ///     var mainVpcPublicGatewayDhcp = new Scaleway.VpcPublicGatewayDhcp("mainVpcPublicGatewayDhcp", new()
        ///     {
        ///         Subnet = "192.168.1.0/24",
        ///     });
        /// 
        ///     var mainVpcPublicGateway = new Scaleway.VpcPublicGateway("mainVpcPublicGateway", new()
        ///     {
        ///         Type = "VPC-GW-S",
        ///         IpId = mainVpcPublicGatewayIp.Id,
        ///     });
        /// 
        ///     var mainVpcGatewayNetwork = new Scaleway.VpcGatewayNetwork("mainVpcGatewayNetwork", new()
        ///     {
        ///         GatewayId = mainVpcPublicGateway.Id,
        ///         PrivateNetworkId = mainVpcPrivateNetwork.Id,
        ///         DhcpId = mainVpcPublicGatewayDhcp.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     var mainVpcPublicGatewayDhcpReservation = new Scaleway.VpcPublicGatewayDhcpReservation("mainVpcPublicGatewayDhcpReservation", new()
        ///     {
        ///         GatewayNetworkId = mainVpcGatewayNetwork.Id,
        ///         MacAddress = mainInstancePrivateNic.MacAddress,
        ///         IpAddress = "192.168.1.4",
        ///     });
        /// 
        ///     //## VPC PAT RULE
        ///     var mainVpcPublicGatewayPatRule = new Scaleway.VpcPublicGatewayPatRule("mainVpcPublicGatewayPatRule", new()
        ///     {
        ///         GatewayId = mainVpcPublicGateway.Id,
        ///         PrivateIp = mainVpcPublicGatewayDhcpReservation.IpAddress,
        ///         PrivatePort = 22,
        ///         PublicPort = 2222,
        ///         Protocol = "tcp",
        ///     });
        /// 
        ///     var byId = Scaleway.GetVpcPublicGatewayDhcpReservation.Invoke(new()
        ///     {
        ///         ReservationId = mainVpcPublicGatewayDhcpReservation.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetVpcPublicGatewayDhcpReservationResult> Invoke(GetVpcPublicGatewayDhcpReservationInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcPublicGatewayDhcpReservationResult>("scaleway:index/getVpcPublicGatewayDhcpReservation:getVpcPublicGatewayDhcpReservation", args ?? new GetVpcPublicGatewayDhcpReservationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcPublicGatewayDhcpReservationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the owning GatewayNetwork.
        /// 
        /// &gt; Only one of `reservation_id` or `mac_address` with `gateway_network_id` should be specified.
        /// </summary>
        [Input("gatewayNetworkId")]
        public string? GatewayNetworkId { get; set; }

        /// <summary>
        /// The MAC address of the reservation to retrieve.
        /// </summary>
        [Input("macAddress")]
        public string? MacAddress { get; set; }

        /// <summary>
        /// The ID of the Reservation to retrieve.
        /// </summary>
        [Input("reservationId")]
        public string? ReservationId { get; set; }

        /// <summary>
        /// Boolean to wait for mac_address to exist in dhcp.
        /// </summary>
        [Input("waitForDhcp")]
        public bool? WaitForDhcp { get; set; }

        /// <summary>
        /// `zone`) The zone in which the image exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetVpcPublicGatewayDhcpReservationArgs()
        {
        }
        public static new GetVpcPublicGatewayDhcpReservationArgs Empty => new GetVpcPublicGatewayDhcpReservationArgs();
    }

    public sealed class GetVpcPublicGatewayDhcpReservationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the owning GatewayNetwork.
        /// 
        /// &gt; Only one of `reservation_id` or `mac_address` with `gateway_network_id` should be specified.
        /// </summary>
        [Input("gatewayNetworkId")]
        public Input<string>? GatewayNetworkId { get; set; }

        /// <summary>
        /// The MAC address of the reservation to retrieve.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// The ID of the Reservation to retrieve.
        /// </summary>
        [Input("reservationId")]
        public Input<string>? ReservationId { get; set; }

        /// <summary>
        /// Boolean to wait for mac_address to exist in dhcp.
        /// </summary>
        [Input("waitForDhcp")]
        public Input<bool>? WaitForDhcp { get; set; }

        /// <summary>
        /// `zone`) The zone in which the image exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetVpcPublicGatewayDhcpReservationInvokeArgs()
        {
        }
        public static new GetVpcPublicGatewayDhcpReservationInvokeArgs Empty => new GetVpcPublicGatewayDhcpReservationInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcPublicGatewayDhcpReservationResult
    {
        public readonly string CreatedAt;
        public readonly string? GatewayNetworkId;
        public readonly string Hostname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string IpAddress;
        public readonly string? MacAddress;
        public readonly string? ReservationId;
        public readonly string Type;
        public readonly string UpdatedAt;
        public readonly bool? WaitForDhcp;
        public readonly string? Zone;

        [OutputConstructor]
        private GetVpcPublicGatewayDhcpReservationResult(
            string createdAt,

            string? gatewayNetworkId,

            string hostname,

            string id,

            string ipAddress,

            string? macAddress,

            string? reservationId,

            string type,

            string updatedAt,

            bool? waitForDhcp,

            string? zone)
        {
            CreatedAt = createdAt;
            GatewayNetworkId = gatewayNetworkId;
            Hostname = hostname;
            Id = id;
            IpAddress = ipAddress;
            MacAddress = macAddress;
            ReservationId = reservationId;
            Type = type;
            UpdatedAt = updatedAt;
            WaitForDhcp = waitForDhcp;
            Zone = zone;
        }
    }
}