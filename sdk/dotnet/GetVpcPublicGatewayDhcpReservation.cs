// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway
{
    [Obsolete(@"scaleway.index/getvpcpublicgatewaydhcpreservation.getVpcPublicGatewayDhcpReservation has been deprecated in favor of scaleway.network/getpublicgatewaydhcpreservation.getPublicGatewayDhcpReservation")]
    public static class GetVpcPublicGatewayDhcpReservation
    {
        /// <summary>
        /// &gt; **Important:**  The data source `scaleway.network.PublicGatewayDhcpReservation` has been deprecated and will no longer be supported.
        /// In 2023, DHCP functionality was moved from Public Gateways to Private Networks, DHCP resources are now no longer needed.
        /// You can use IPAM to manage your IPs. For more information, please refer to the dedicated guide.
        /// 
        /// Gets information about a DHCP entry. For further information, please see the
        /// [API documentation](https://www.scaleway.com/en/developers/api/public-gateway/#path-dhcp-entries-list-dhcp-entries).
        /// 
        /// ## Example Dynamic
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = Pulumiverse.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = new Scaleway.Network.PrivateNetwork("main");
        /// 
        ///     var mainServer = new Scaleway.Instance.Server("main", new()
        ///     {
        ///         Image = "ubuntu_jammy",
        ///         Type = "DEV1-S",
        ///         Zone = "fr-par-1",
        ///     });
        /// 
        ///     var mainPrivateNic = new Scaleway.Instance.PrivateNic("main", new()
        ///     {
        ///         ServerId = mainServer.Id,
        ///         PrivateNetworkId = main.Id,
        ///     });
        /// 
        ///     var mainPublicGatewayIp = new Scaleway.Network.PublicGatewayIp("main");
        /// 
        ///     var mainPublicGatewayDhcp = new Scaleway.Network.PublicGatewayDhcp("main", new()
        ///     {
        ///         Subnet = "192.168.1.0/24",
        ///     });
        /// 
        ///     var mainPublicGateway = new Scaleway.Network.PublicGateway("main", new()
        ///     {
        ///         Name = "foobar",
        ///         Type = "VPC-GW-S",
        ///         IpId = mainPublicGatewayIp.Id,
        ///     });
        /// 
        ///     var mainGatewayNetwork = new Scaleway.Network.GatewayNetwork("main", new()
        ///     {
        ///         GatewayId = mainPublicGateway.Id,
        ///         PrivateNetworkId = main.Id,
        ///         DhcpId = mainPublicGatewayDhcp.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     //# Retrieve the dynamic entries generated by mac address &amp; gateway network
        ///     var byMacAddressAndGwNetwork = Scaleway.Network.GetPublicGatewayDhcpReservation.Invoke(new()
        ///     {
        ///         MacAddress = mainPrivateNic.MacAddress,
        ///         GatewayNetworkId = mainGatewayNetwork.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Example Static and PAT Rule
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = Pulumiverse.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = new Scaleway.Network.PrivateNetwork("main");
        /// 
        ///     var mainSecurityGroup = new Scaleway.Instance.SecurityGroup("main", new()
        ///     {
        ///         InboundDefaultPolicy = "drop",
        ///         OutboundDefaultPolicy = "accept",
        ///         InboundRules = new[]
        ///         {
        ///             new Scaleway.Instance.Inputs.SecurityGroupInboundRuleArgs
        ///             {
        ///                 Action = "accept",
        ///                 Port = 22,
        ///             },
        ///         },
        ///     });
        /// 
        ///     var mainServer = new Scaleway.Instance.Server("main", new()
        ///     {
        ///         Image = "ubuntu_jammy",
        ///         Type = "DEV1-S",
        ///         Zone = "fr-par-1",
        ///         SecurityGroupId = mainSecurityGroup.Id,
        ///     });
        /// 
        ///     var mainPrivateNic = new Scaleway.Instance.PrivateNic("main", new()
        ///     {
        ///         ServerId = mainServer.Id,
        ///         PrivateNetworkId = main.Id,
        ///     });
        /// 
        ///     var mainPublicGatewayIp = new Scaleway.Network.PublicGatewayIp("main");
        /// 
        ///     var mainPublicGatewayDhcp = new Scaleway.Network.PublicGatewayDhcp("main", new()
        ///     {
        ///         Subnet = "192.168.1.0/24",
        ///     });
        /// 
        ///     var mainPublicGateway = new Scaleway.Network.PublicGateway("main", new()
        ///     {
        ///         Name = "foobar",
        ///         Type = "VPC-GW-S",
        ///         IpId = mainPublicGatewayIp.Id,
        ///     });
        /// 
        ///     var mainGatewayNetwork = new Scaleway.Network.GatewayNetwork("main", new()
        ///     {
        ///         GatewayId = mainPublicGateway.Id,
        ///         PrivateNetworkId = main.Id,
        ///         DhcpId = mainPublicGatewayDhcp.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     var mainPublicGatewayDhcpReservation = new Scaleway.Network.PublicGatewayDhcpReservation("main", new()
        ///     {
        ///         GatewayNetworkId = mainGatewayNetwork.Id,
        ///         MacAddress = mainPrivateNic.MacAddress,
        ///         IpAddress = "192.168.1.4",
        ///     });
        /// 
        ///     //## VPC PAT RULE
        ///     var mainPublicGatewayPatRule = new Scaleway.Network.PublicGatewayPatRule("main", new()
        ///     {
        ///         GatewayId = mainPublicGateway.Id,
        ///         PrivateIp = mainPublicGatewayDhcpReservation.IpAddress,
        ///         PrivatePort = 22,
        ///         PublicPort = 2222,
        ///         Protocol = "tcp",
        ///     });
        /// 
        ///     var byId = Scaleway.Network.GetPublicGatewayDhcpReservation.Invoke(new()
        ///     {
        ///         ReservationId = mainPublicGatewayDhcpReservation.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetVpcPublicGatewayDhcpReservationResult> InvokeAsync(GetVpcPublicGatewayDhcpReservationArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcPublicGatewayDhcpReservationResult>("scaleway:index/getVpcPublicGatewayDhcpReservation:getVpcPublicGatewayDhcpReservation", args ?? new GetVpcPublicGatewayDhcpReservationArgs(), options.WithDefaults());

        /// <summary>
        /// &gt; **Important:**  The data source `scaleway.network.PublicGatewayDhcpReservation` has been deprecated and will no longer be supported.
        /// In 2023, DHCP functionality was moved from Public Gateways to Private Networks, DHCP resources are now no longer needed.
        /// You can use IPAM to manage your IPs. For more information, please refer to the dedicated guide.
        /// 
        /// Gets information about a DHCP entry. For further information, please see the
        /// [API documentation](https://www.scaleway.com/en/developers/api/public-gateway/#path-dhcp-entries-list-dhcp-entries).
        /// 
        /// ## Example Dynamic
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = Pulumiverse.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = new Scaleway.Network.PrivateNetwork("main");
        /// 
        ///     var mainServer = new Scaleway.Instance.Server("main", new()
        ///     {
        ///         Image = "ubuntu_jammy",
        ///         Type = "DEV1-S",
        ///         Zone = "fr-par-1",
        ///     });
        /// 
        ///     var mainPrivateNic = new Scaleway.Instance.PrivateNic("main", new()
        ///     {
        ///         ServerId = mainServer.Id,
        ///         PrivateNetworkId = main.Id,
        ///     });
        /// 
        ///     var mainPublicGatewayIp = new Scaleway.Network.PublicGatewayIp("main");
        /// 
        ///     var mainPublicGatewayDhcp = new Scaleway.Network.PublicGatewayDhcp("main", new()
        ///     {
        ///         Subnet = "192.168.1.0/24",
        ///     });
        /// 
        ///     var mainPublicGateway = new Scaleway.Network.PublicGateway("main", new()
        ///     {
        ///         Name = "foobar",
        ///         Type = "VPC-GW-S",
        ///         IpId = mainPublicGatewayIp.Id,
        ///     });
        /// 
        ///     var mainGatewayNetwork = new Scaleway.Network.GatewayNetwork("main", new()
        ///     {
        ///         GatewayId = mainPublicGateway.Id,
        ///         PrivateNetworkId = main.Id,
        ///         DhcpId = mainPublicGatewayDhcp.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     //# Retrieve the dynamic entries generated by mac address &amp; gateway network
        ///     var byMacAddressAndGwNetwork = Scaleway.Network.GetPublicGatewayDhcpReservation.Invoke(new()
        ///     {
        ///         MacAddress = mainPrivateNic.MacAddress,
        ///         GatewayNetworkId = mainGatewayNetwork.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Example Static and PAT Rule
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = Pulumiverse.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = new Scaleway.Network.PrivateNetwork("main");
        /// 
        ///     var mainSecurityGroup = new Scaleway.Instance.SecurityGroup("main", new()
        ///     {
        ///         InboundDefaultPolicy = "drop",
        ///         OutboundDefaultPolicy = "accept",
        ///         InboundRules = new[]
        ///         {
        ///             new Scaleway.Instance.Inputs.SecurityGroupInboundRuleArgs
        ///             {
        ///                 Action = "accept",
        ///                 Port = 22,
        ///             },
        ///         },
        ///     });
        /// 
        ///     var mainServer = new Scaleway.Instance.Server("main", new()
        ///     {
        ///         Image = "ubuntu_jammy",
        ///         Type = "DEV1-S",
        ///         Zone = "fr-par-1",
        ///         SecurityGroupId = mainSecurityGroup.Id,
        ///     });
        /// 
        ///     var mainPrivateNic = new Scaleway.Instance.PrivateNic("main", new()
        ///     {
        ///         ServerId = mainServer.Id,
        ///         PrivateNetworkId = main.Id,
        ///     });
        /// 
        ///     var mainPublicGatewayIp = new Scaleway.Network.PublicGatewayIp("main");
        /// 
        ///     var mainPublicGatewayDhcp = new Scaleway.Network.PublicGatewayDhcp("main", new()
        ///     {
        ///         Subnet = "192.168.1.0/24",
        ///     });
        /// 
        ///     var mainPublicGateway = new Scaleway.Network.PublicGateway("main", new()
        ///     {
        ///         Name = "foobar",
        ///         Type = "VPC-GW-S",
        ///         IpId = mainPublicGatewayIp.Id,
        ///     });
        /// 
        ///     var mainGatewayNetwork = new Scaleway.Network.GatewayNetwork("main", new()
        ///     {
        ///         GatewayId = mainPublicGateway.Id,
        ///         PrivateNetworkId = main.Id,
        ///         DhcpId = mainPublicGatewayDhcp.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     var mainPublicGatewayDhcpReservation = new Scaleway.Network.PublicGatewayDhcpReservation("main", new()
        ///     {
        ///         GatewayNetworkId = mainGatewayNetwork.Id,
        ///         MacAddress = mainPrivateNic.MacAddress,
        ///         IpAddress = "192.168.1.4",
        ///     });
        /// 
        ///     //## VPC PAT RULE
        ///     var mainPublicGatewayPatRule = new Scaleway.Network.PublicGatewayPatRule("main", new()
        ///     {
        ///         GatewayId = mainPublicGateway.Id,
        ///         PrivateIp = mainPublicGatewayDhcpReservation.IpAddress,
        ///         PrivatePort = 22,
        ///         PublicPort = 2222,
        ///         Protocol = "tcp",
        ///     });
        /// 
        ///     var byId = Scaleway.Network.GetPublicGatewayDhcpReservation.Invoke(new()
        ///     {
        ///         ReservationId = mainPublicGatewayDhcpReservation.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetVpcPublicGatewayDhcpReservationResult> Invoke(GetVpcPublicGatewayDhcpReservationInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcPublicGatewayDhcpReservationResult>("scaleway:index/getVpcPublicGatewayDhcpReservation:getVpcPublicGatewayDhcpReservation", args ?? new GetVpcPublicGatewayDhcpReservationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// &gt; **Important:**  The data source `scaleway.network.PublicGatewayDhcpReservation` has been deprecated and will no longer be supported.
        /// In 2023, DHCP functionality was moved from Public Gateways to Private Networks, DHCP resources are now no longer needed.
        /// You can use IPAM to manage your IPs. For more information, please refer to the dedicated guide.
        /// 
        /// Gets information about a DHCP entry. For further information, please see the
        /// [API documentation](https://www.scaleway.com/en/developers/api/public-gateway/#path-dhcp-entries-list-dhcp-entries).
        /// 
        /// ## Example Dynamic
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = Pulumiverse.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = new Scaleway.Network.PrivateNetwork("main");
        /// 
        ///     var mainServer = new Scaleway.Instance.Server("main", new()
        ///     {
        ///         Image = "ubuntu_jammy",
        ///         Type = "DEV1-S",
        ///         Zone = "fr-par-1",
        ///     });
        /// 
        ///     var mainPrivateNic = new Scaleway.Instance.PrivateNic("main", new()
        ///     {
        ///         ServerId = mainServer.Id,
        ///         PrivateNetworkId = main.Id,
        ///     });
        /// 
        ///     var mainPublicGatewayIp = new Scaleway.Network.PublicGatewayIp("main");
        /// 
        ///     var mainPublicGatewayDhcp = new Scaleway.Network.PublicGatewayDhcp("main", new()
        ///     {
        ///         Subnet = "192.168.1.0/24",
        ///     });
        /// 
        ///     var mainPublicGateway = new Scaleway.Network.PublicGateway("main", new()
        ///     {
        ///         Name = "foobar",
        ///         Type = "VPC-GW-S",
        ///         IpId = mainPublicGatewayIp.Id,
        ///     });
        /// 
        ///     var mainGatewayNetwork = new Scaleway.Network.GatewayNetwork("main", new()
        ///     {
        ///         GatewayId = mainPublicGateway.Id,
        ///         PrivateNetworkId = main.Id,
        ///         DhcpId = mainPublicGatewayDhcp.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     //# Retrieve the dynamic entries generated by mac address &amp; gateway network
        ///     var byMacAddressAndGwNetwork = Scaleway.Network.GetPublicGatewayDhcpReservation.Invoke(new()
        ///     {
        ///         MacAddress = mainPrivateNic.MacAddress,
        ///         GatewayNetworkId = mainGatewayNetwork.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Example Static and PAT Rule
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = Pulumiverse.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = new Scaleway.Network.PrivateNetwork("main");
        /// 
        ///     var mainSecurityGroup = new Scaleway.Instance.SecurityGroup("main", new()
        ///     {
        ///         InboundDefaultPolicy = "drop",
        ///         OutboundDefaultPolicy = "accept",
        ///         InboundRules = new[]
        ///         {
        ///             new Scaleway.Instance.Inputs.SecurityGroupInboundRuleArgs
        ///             {
        ///                 Action = "accept",
        ///                 Port = 22,
        ///             },
        ///         },
        ///     });
        /// 
        ///     var mainServer = new Scaleway.Instance.Server("main", new()
        ///     {
        ///         Image = "ubuntu_jammy",
        ///         Type = "DEV1-S",
        ///         Zone = "fr-par-1",
        ///         SecurityGroupId = mainSecurityGroup.Id,
        ///     });
        /// 
        ///     var mainPrivateNic = new Scaleway.Instance.PrivateNic("main", new()
        ///     {
        ///         ServerId = mainServer.Id,
        ///         PrivateNetworkId = main.Id,
        ///     });
        /// 
        ///     var mainPublicGatewayIp = new Scaleway.Network.PublicGatewayIp("main");
        /// 
        ///     var mainPublicGatewayDhcp = new Scaleway.Network.PublicGatewayDhcp("main", new()
        ///     {
        ///         Subnet = "192.168.1.0/24",
        ///     });
        /// 
        ///     var mainPublicGateway = new Scaleway.Network.PublicGateway("main", new()
        ///     {
        ///         Name = "foobar",
        ///         Type = "VPC-GW-S",
        ///         IpId = mainPublicGatewayIp.Id,
        ///     });
        /// 
        ///     var mainGatewayNetwork = new Scaleway.Network.GatewayNetwork("main", new()
        ///     {
        ///         GatewayId = mainPublicGateway.Id,
        ///         PrivateNetworkId = main.Id,
        ///         DhcpId = mainPublicGatewayDhcp.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     var mainPublicGatewayDhcpReservation = new Scaleway.Network.PublicGatewayDhcpReservation("main", new()
        ///     {
        ///         GatewayNetworkId = mainGatewayNetwork.Id,
        ///         MacAddress = mainPrivateNic.MacAddress,
        ///         IpAddress = "192.168.1.4",
        ///     });
        /// 
        ///     //## VPC PAT RULE
        ///     var mainPublicGatewayPatRule = new Scaleway.Network.PublicGatewayPatRule("main", new()
        ///     {
        ///         GatewayId = mainPublicGateway.Id,
        ///         PrivateIp = mainPublicGatewayDhcpReservation.IpAddress,
        ///         PrivatePort = 22,
        ///         PublicPort = 2222,
        ///         Protocol = "tcp",
        ///     });
        /// 
        ///     var byId = Scaleway.Network.GetPublicGatewayDhcpReservation.Invoke(new()
        ///     {
        ///         ReservationId = mainPublicGatewayDhcpReservation.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetVpcPublicGatewayDhcpReservationResult> Invoke(GetVpcPublicGatewayDhcpReservationInvokeArgs args, InvokeOutputOptions options)
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
        /// The ID of the reservation (DHCP entry) to retrieve.
        /// </summary>
        [Input("reservationId")]
        public string? ReservationId { get; set; }

        /// <summary>
        /// Whether to wait for `mac_address` to exist in DHCP.
        /// </summary>
        [Input("waitForDhcp")]
        public bool? WaitForDhcp { get; set; }

        /// <summary>
        /// `zone`). The zone in which the reservation exists.
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
        /// The ID of the reservation (DHCP entry) to retrieve.
        /// </summary>
        [Input("reservationId")]
        public Input<string>? ReservationId { get; set; }

        /// <summary>
        /// Whether to wait for `mac_address` to exist in DHCP.
        /// </summary>
        [Input("waitForDhcp")]
        public Input<bool>? WaitForDhcp { get; set; }

        /// <summary>
        /// `zone`). The zone in which the reservation exists.
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
