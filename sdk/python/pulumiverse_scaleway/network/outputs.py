# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = [
    'AclRule',
    'GatewayNetworkIpamConfig',
    'GatewayNetworkPrivateIp',
    'PrivateNetworkIpv4Subnet',
    'PrivateNetworkIpv6Subnet',
    'GetGatewayNetworkIpamConfigResult',
    'GetGatewayNetworkPrivateIpResult',
    'GetPrivateNetworkIpv4SubnetResult',
    'GetPrivateNetworkIpv6SubnetResult',
    'GetRoutesRouteResult',
    'GetVpcsVpcResult',
]

@pulumi.output_type
class AclRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dstPortHigh":
            suggest = "dst_port_high"
        elif key == "dstPortLow":
            suggest = "dst_port_low"
        elif key == "srcPortHigh":
            suggest = "src_port_high"
        elif key == "srcPortLow":
            suggest = "src_port_low"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AclRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AclRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AclRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: Optional[str] = None,
                 description: Optional[str] = None,
                 destination: Optional[str] = None,
                 dst_port_high: Optional[int] = None,
                 dst_port_low: Optional[int] = None,
                 protocol: Optional[str] = None,
                 source: Optional[str] = None,
                 src_port_high: Optional[int] = None,
                 src_port_low: Optional[int] = None):
        """
        :param str action: The policy to apply to the packet.
        :param str description: The rule description.
        :param str destination: The destination IP range to which this rule applies (CIDR notation with subnet mask).
        :param int dst_port_high: The ending port of the destination port range to which this rule applies (inclusive).
        :param int dst_port_low: The starting port of the destination port range to which this rule applies (inclusive).
        :param str protocol: The protocol to which this rule applies. Default value: ANY.
        :param str source: The Source IP range to which this rule applies (CIDR notation with subnet mask).
        :param int src_port_high: The ending port of the source port range to which this rule applies (inclusive).
        :param int src_port_low: The starting port of the source port range to which this rule applies (inclusive).
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destination is not None:
            pulumi.set(__self__, "destination", destination)
        if dst_port_high is not None:
            pulumi.set(__self__, "dst_port_high", dst_port_high)
        if dst_port_low is not None:
            pulumi.set(__self__, "dst_port_low", dst_port_low)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if src_port_high is not None:
            pulumi.set(__self__, "src_port_high", src_port_high)
        if src_port_low is not None:
            pulumi.set(__self__, "src_port_low", src_port_low)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        The policy to apply to the packet.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The rule description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def destination(self) -> Optional[str]:
        """
        The destination IP range to which this rule applies (CIDR notation with subnet mask).
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter(name="dstPortHigh")
    def dst_port_high(self) -> Optional[int]:
        """
        The ending port of the destination port range to which this rule applies (inclusive).
        """
        return pulumi.get(self, "dst_port_high")

    @property
    @pulumi.getter(name="dstPortLow")
    def dst_port_low(self) -> Optional[int]:
        """
        The starting port of the destination port range to which this rule applies (inclusive).
        """
        return pulumi.get(self, "dst_port_low")

    @property
    @pulumi.getter
    def protocol(self) -> Optional[str]:
        """
        The protocol to which this rule applies. Default value: ANY.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def source(self) -> Optional[str]:
        """
        The Source IP range to which this rule applies (CIDR notation with subnet mask).
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="srcPortHigh")
    def src_port_high(self) -> Optional[int]:
        """
        The ending port of the source port range to which this rule applies (inclusive).
        """
        return pulumi.get(self, "src_port_high")

    @property
    @pulumi.getter(name="srcPortLow")
    def src_port_low(self) -> Optional[int]:
        """
        The starting port of the source port range to which this rule applies (inclusive).
        """
        return pulumi.get(self, "src_port_low")


@pulumi.output_type
class GatewayNetworkIpamConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ipamIpId":
            suggest = "ipam_ip_id"
        elif key == "pushDefaultRoute":
            suggest = "push_default_route"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GatewayNetworkIpamConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GatewayNetworkIpamConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GatewayNetworkIpamConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ipam_ip_id: Optional[str] = None,
                 push_default_route: Optional[bool] = None):
        """
        :param str ipam_ip_id: Use this IPAM-booked IP ID as the Gateway's IP in this Private Network.
        :param bool push_default_route: Defines whether to enable the default route on the GatewayNetwork.
        """
        if ipam_ip_id is not None:
            pulumi.set(__self__, "ipam_ip_id", ipam_ip_id)
        if push_default_route is not None:
            pulumi.set(__self__, "push_default_route", push_default_route)

    @property
    @pulumi.getter(name="ipamIpId")
    def ipam_ip_id(self) -> Optional[str]:
        """
        Use this IPAM-booked IP ID as the Gateway's IP in this Private Network.
        """
        return pulumi.get(self, "ipam_ip_id")

    @property
    @pulumi.getter(name="pushDefaultRoute")
    def push_default_route(self) -> Optional[bool]:
        """
        Defines whether to enable the default route on the GatewayNetwork.
        """
        return pulumi.get(self, "push_default_route")


@pulumi.output_type
class GatewayNetworkPrivateIp(dict):
    def __init__(__self__, *,
                 address: Optional[str] = None,
                 id: Optional[str] = None):
        """
        :param str address: The private IPv4 address.
        :param str id: The ID of the IPv4 address resource.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def address(self) -> Optional[str]:
        """
        The private IPv4 address.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of the IPv4 address resource.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class PrivateNetworkIpv4Subnet(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "prefixLength":
            suggest = "prefix_length"
        elif key == "subnetMask":
            suggest = "subnet_mask"
        elif key == "updatedAt":
            suggest = "updated_at"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PrivateNetworkIpv4Subnet. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PrivateNetworkIpv4Subnet.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PrivateNetworkIpv4Subnet.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 address: Optional[str] = None,
                 created_at: Optional[str] = None,
                 id: Optional[str] = None,
                 prefix_length: Optional[int] = None,
                 subnet: Optional[str] = None,
                 subnet_mask: Optional[str] = None,
                 updated_at: Optional[str] = None):
        """
        :param str address: The network address of the subnet in hexadecimal notation, e.g., '2001:db8::' for a '2001:db8::/64' subnet.
        :param str created_at: The date and time of the creation of the subnet.
        :param str id: The subnet ID.
        :param int prefix_length: The length of the network prefix, e.g., 64 for a 'ffff:ffff:ffff:ffff::' mask.
        :param str subnet: The subnet CIDR.
        :param str subnet_mask: The subnet mask expressed in dotted decimal notation, e.g., '255.255.255.0' for a /24 subnet
        :param str updated_at: The date and time of the last update of the subnet.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if prefix_length is not None:
            pulumi.set(__self__, "prefix_length", prefix_length)
        if subnet is not None:
            pulumi.set(__self__, "subnet", subnet)
        if subnet_mask is not None:
            pulumi.set(__self__, "subnet_mask", subnet_mask)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def address(self) -> Optional[str]:
        """
        The network address of the subnet in hexadecimal notation, e.g., '2001:db8::' for a '2001:db8::/64' subnet.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        The date and time of the creation of the subnet.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The subnet ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> Optional[int]:
        """
        The length of the network prefix, e.g., 64 for a 'ffff:ffff:ffff:ffff::' mask.
        """
        return pulumi.get(self, "prefix_length")

    @property
    @pulumi.getter
    def subnet(self) -> Optional[str]:
        """
        The subnet CIDR.
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter(name="subnetMask")
    def subnet_mask(self) -> Optional[str]:
        """
        The subnet mask expressed in dotted decimal notation, e.g., '255.255.255.0' for a /24 subnet
        """
        return pulumi.get(self, "subnet_mask")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[str]:
        """
        The date and time of the last update of the subnet.
        """
        return pulumi.get(self, "updated_at")


@pulumi.output_type
class PrivateNetworkIpv6Subnet(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "prefixLength":
            suggest = "prefix_length"
        elif key == "subnetMask":
            suggest = "subnet_mask"
        elif key == "updatedAt":
            suggest = "updated_at"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PrivateNetworkIpv6Subnet. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PrivateNetworkIpv6Subnet.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PrivateNetworkIpv6Subnet.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 address: Optional[str] = None,
                 created_at: Optional[str] = None,
                 id: Optional[str] = None,
                 prefix_length: Optional[int] = None,
                 subnet: Optional[str] = None,
                 subnet_mask: Optional[str] = None,
                 updated_at: Optional[str] = None):
        """
        :param str address: The network address of the subnet in hexadecimal notation, e.g., '2001:db8::' for a '2001:db8::/64' subnet.
        :param str created_at: The date and time of the creation of the subnet.
        :param str id: The subnet ID.
        :param int prefix_length: The length of the network prefix, e.g., 64 for a 'ffff:ffff:ffff:ffff::' mask.
        :param str subnet: The subnet CIDR.
        :param str subnet_mask: The subnet mask expressed in dotted decimal notation, e.g., '255.255.255.0' for a /24 subnet
        :param str updated_at: The date and time of the last update of the subnet.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if prefix_length is not None:
            pulumi.set(__self__, "prefix_length", prefix_length)
        if subnet is not None:
            pulumi.set(__self__, "subnet", subnet)
        if subnet_mask is not None:
            pulumi.set(__self__, "subnet_mask", subnet_mask)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def address(self) -> Optional[str]:
        """
        The network address of the subnet in hexadecimal notation, e.g., '2001:db8::' for a '2001:db8::/64' subnet.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        The date and time of the creation of the subnet.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The subnet ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> Optional[int]:
        """
        The length of the network prefix, e.g., 64 for a 'ffff:ffff:ffff:ffff::' mask.
        """
        return pulumi.get(self, "prefix_length")

    @property
    @pulumi.getter
    def subnet(self) -> Optional[str]:
        """
        The subnet CIDR.
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter(name="subnetMask")
    def subnet_mask(self) -> Optional[str]:
        """
        The subnet mask expressed in dotted decimal notation, e.g., '255.255.255.0' for a /24 subnet
        """
        return pulumi.get(self, "subnet_mask")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[str]:
        """
        The date and time of the last update of the subnet.
        """
        return pulumi.get(self, "updated_at")


@pulumi.output_type
class GetGatewayNetworkIpamConfigResult(dict):
    def __init__(__self__, *,
                 ipam_ip_id: str,
                 push_default_route: bool):
        """
        :param str ipam_ip_id: Use this IPAM-booked IP ID as the Gateway's IP in this Private Network
        :param bool push_default_route: Defines whether the default route is enabled on that Gateway Network
        """
        pulumi.set(__self__, "ipam_ip_id", ipam_ip_id)
        pulumi.set(__self__, "push_default_route", push_default_route)

    @property
    @pulumi.getter(name="ipamIpId")
    def ipam_ip_id(self) -> str:
        """
        Use this IPAM-booked IP ID as the Gateway's IP in this Private Network
        """
        return pulumi.get(self, "ipam_ip_id")

    @property
    @pulumi.getter(name="pushDefaultRoute")
    def push_default_route(self) -> bool:
        """
        Defines whether the default route is enabled on that Gateway Network
        """
        return pulumi.get(self, "push_default_route")


@pulumi.output_type
class GetGatewayNetworkPrivateIpResult(dict):
    def __init__(__self__, *,
                 address: str,
                 id: str):
        """
        :param str address: The private IPv4 address.
        :param str id: The ID of the IPv4 address resource.
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def address(self) -> str:
        """
        The private IPv4 address.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the IPv4 address resource.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class GetPrivateNetworkIpv4SubnetResult(dict):
    def __init__(__self__, *,
                 address: str,
                 created_at: str,
                 id: str,
                 prefix_length: int,
                 subnet: str,
                 subnet_mask: str,
                 updated_at: str):
        """
        :param str address: The network address of the subnet in dotted decimal notation, e.g., '192.168.0.0' for a '192.168.0.0/24' subnet
        :param str created_at: The date and time of the creation of the subnet
        :param str id: The ID of the Private Network.
        :param int prefix_length: The length of the network prefix, e.g., 24 for a 255.255.255.0 mask
        :param str subnet: The subnet CIDR
        :param str subnet_mask: The subnet mask expressed in dotted decimal notation, e.g., '255.255.255.0' for a /24 subnet
        :param str updated_at: The date and time of the last update of the subnet
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "created_at", created_at)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "prefix_length", prefix_length)
        pulumi.set(__self__, "subnet", subnet)
        pulumi.set(__self__, "subnet_mask", subnet_mask)
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def address(self) -> str:
        """
        The network address of the subnet in dotted decimal notation, e.g., '192.168.0.0' for a '192.168.0.0/24' subnet
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The date and time of the creation of the subnet
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Private Network.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> int:
        """
        The length of the network prefix, e.g., 24 for a 255.255.255.0 mask
        """
        return pulumi.get(self, "prefix_length")

    @property
    @pulumi.getter
    def subnet(self) -> str:
        """
        The subnet CIDR
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter(name="subnetMask")
    def subnet_mask(self) -> str:
        """
        The subnet mask expressed in dotted decimal notation, e.g., '255.255.255.0' for a /24 subnet
        """
        return pulumi.get(self, "subnet_mask")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        The date and time of the last update of the subnet
        """
        return pulumi.get(self, "updated_at")


@pulumi.output_type
class GetPrivateNetworkIpv6SubnetResult(dict):
    def __init__(__self__, *,
                 address: str,
                 created_at: str,
                 id: str,
                 prefix_length: int,
                 subnet: str,
                 subnet_mask: str,
                 updated_at: str):
        """
        :param str address: The network address of the subnet in dotted decimal notation, e.g., '192.168.0.0' for a '192.168.0.0/24' subnet
        :param str created_at: The date and time of the creation of the subnet
        :param str id: The ID of the Private Network.
        :param int prefix_length: The length of the network prefix, e.g., 24 for a 255.255.255.0 mask
        :param str subnet: The subnet CIDR
        :param str subnet_mask: The subnet mask expressed in dotted decimal notation, e.g., '255.255.255.0' for a /24 subnet
        :param str updated_at: The date and time of the last update of the subnet
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "created_at", created_at)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "prefix_length", prefix_length)
        pulumi.set(__self__, "subnet", subnet)
        pulumi.set(__self__, "subnet_mask", subnet_mask)
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def address(self) -> str:
        """
        The network address of the subnet in dotted decimal notation, e.g., '192.168.0.0' for a '192.168.0.0/24' subnet
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The date and time of the creation of the subnet
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Private Network.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> int:
        """
        The length of the network prefix, e.g., 24 for a 255.255.255.0 mask
        """
        return pulumi.get(self, "prefix_length")

    @property
    @pulumi.getter
    def subnet(self) -> str:
        """
        The subnet CIDR
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter(name="subnetMask")
    def subnet_mask(self) -> str:
        """
        The subnet mask expressed in dotted decimal notation, e.g., '255.255.255.0' for a /24 subnet
        """
        return pulumi.get(self, "subnet_mask")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        The date and time of the last update of the subnet
        """
        return pulumi.get(self, "updated_at")


@pulumi.output_type
class GetRoutesRouteResult(dict):
    def __init__(__self__, *,
                 created_at: str,
                 description: str,
                 destination: str,
                 id: str,
                 nexthop_ip: str,
                 nexthop_name: str,
                 nexthop_private_network_id: str,
                 nexthop_resource_id: str,
                 nexthop_resource_type: str,
                 region: str,
                 tags: Sequence[str],
                 vpc_id: str):
        """
        :param str created_at: The date on which the route was created (RFC 3339 format).
        :param str description: The description of the route.
        :param str destination: The destination IP or IP range of the route.
        :param str id: The ID of the route.
               > **Important:** route IDs are regional, which means they are of the form `{region}/{id}`, e.g. `fr-par/11111111-1111-1111-1111-111111111111
        :param str nexthop_ip: The IP of the route's next hop.
        :param str nexthop_name: The name of the route's next hop.
        :param str nexthop_private_network_id: The next hop private network ID to filter for. routes with a similar next hop private network ID are listed.
        :param str nexthop_resource_id: The next hop resource ID to filter for. routes with a similar next hop resource ID are listed.
        :param str nexthop_resource_type: The next hop resource type to filter for. routes with a similar next hop resource type are listed.
        :param str region: `region`). The region in which the routes exist.
        :param Sequence[str] tags: List of tags to filter for. routes with these exact tags are listed.
        :param str vpc_id: The VPC ID to filter for. routes with a similar VPC ID are listed.
        """
        pulumi.set(__self__, "created_at", created_at)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "destination", destination)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "nexthop_ip", nexthop_ip)
        pulumi.set(__self__, "nexthop_name", nexthop_name)
        pulumi.set(__self__, "nexthop_private_network_id", nexthop_private_network_id)
        pulumi.set(__self__, "nexthop_resource_id", nexthop_resource_id)
        pulumi.set(__self__, "nexthop_resource_type", nexthop_resource_type)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The date on which the route was created (RFC 3339 format).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the route.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def destination(self) -> str:
        """
        The destination IP or IP range of the route.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the route.
        > **Important:** route IDs are regional, which means they are of the form `{region}/{id}`, e.g. `fr-par/11111111-1111-1111-1111-111111111111
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="nexthopIp")
    def nexthop_ip(self) -> str:
        """
        The IP of the route's next hop.
        """
        return pulumi.get(self, "nexthop_ip")

    @property
    @pulumi.getter(name="nexthopName")
    def nexthop_name(self) -> str:
        """
        The name of the route's next hop.
        """
        return pulumi.get(self, "nexthop_name")

    @property
    @pulumi.getter(name="nexthopPrivateNetworkId")
    def nexthop_private_network_id(self) -> str:
        """
        The next hop private network ID to filter for. routes with a similar next hop private network ID are listed.
        """
        return pulumi.get(self, "nexthop_private_network_id")

    @property
    @pulumi.getter(name="nexthopResourceId")
    def nexthop_resource_id(self) -> str:
        """
        The next hop resource ID to filter for. routes with a similar next hop resource ID are listed.
        """
        return pulumi.get(self, "nexthop_resource_id")

    @property
    @pulumi.getter(name="nexthopResourceType")
    def nexthop_resource_type(self) -> str:
        """
        The next hop resource type to filter for. routes with a similar next hop resource type are listed.
        """
        return pulumi.get(self, "nexthop_resource_type")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        `region`). The region in which the routes exist.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        List of tags to filter for. routes with these exact tags are listed.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The VPC ID to filter for. routes with a similar VPC ID are listed.
        """
        return pulumi.get(self, "vpc_id")


@pulumi.output_type
class GetVpcsVpcResult(dict):
    def __init__(__self__, *,
                 created_at: str,
                 id: str,
                 is_default: bool,
                 name: str,
                 organization_id: str,
                 project_id: str,
                 region: str,
                 tags: Sequence[str],
                 update_at: str):
        """
        :param str created_at: Date and time of VPC's creation (RFC 3339 format).
        :param str id: The associated VPC ID.
               > **Important:** VPC IDs are regional, which means they are of the form `{region}/{id}`, e.g. `fr-par/11111111-1111-1111-1111-111111111111
        :param bool is_default: Defines whether the VPC is the default one for its Project.
        :param str name: The VPC name to filter for. VPCs with a similar name are listed.
        :param str organization_id: The Organization ID the VPC is associated with.
        :param str project_id: The ID of the Project the VPC is associated with.
        :param str region: `region`). The region in which the VPCs exist.
        :param Sequence[str] tags: List of tags to filter for. VPCs with these exact tags are listed.
        """
        pulumi.set(__self__, "created_at", created_at)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "is_default", is_default)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "organization_id", organization_id)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "update_at", update_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Date and time of VPC's creation (RFC 3339 format).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The associated VPC ID.
        > **Important:** VPC IDs are regional, which means they are of the form `{region}/{id}`, e.g. `fr-par/11111111-1111-1111-1111-111111111111
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> bool:
        """
        Defines whether the VPC is the default one for its Project.
        """
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The VPC name to filter for. VPCs with a similar name are listed.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        """
        The Organization ID the VPC is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The ID of the Project the VPC is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        `region`). The region in which the VPCs exist.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        List of tags to filter for. VPCs with these exact tags are listed.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updateAt")
    def update_at(self) -> str:
        return pulumi.get(self, "update_at")


