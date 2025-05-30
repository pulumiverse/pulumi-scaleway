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
    'InstancePrivateIp',
    'InstancePrivateNetwork',
    'InstancePublicNetwork',
    'GetInstancePrivateIpResult',
    'GetInstancePrivateNetworkResult',
    'GetInstancePublicNetworkResult',
]

@pulumi.output_type
class InstancePrivateIp(dict):
    def __init__(__self__, *,
                 address: Optional[str] = None,
                 id: Optional[str] = None):
        """
        :param str address: The private IPv4 address.
        :param str id: The ID of the endpoint.
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
        The ID of the endpoint.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class InstancePrivateNetwork(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "pnId":
            suggest = "pn_id"
        elif key == "dnsRecords":
            suggest = "dns_records"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstancePrivateNetwork. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstancePrivateNetwork.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstancePrivateNetwork.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 pn_id: str,
                 dns_records: Optional[Sequence[str]] = None,
                 id: Optional[str] = None,
                 ips: Optional[Sequence[str]] = None,
                 port: Optional[int] = None):
        """
        :param str pn_id: The ID of the Private Network.
        :param Sequence[str] dns_records: List of DNS records for your endpoint.
        :param str id: The ID of the endpoint.
        :param Sequence[str] ips: List of IP addresses for your endpoint.
        :param int port: TCP port of the endpoint.
        """
        pulumi.set(__self__, "pn_id", pn_id)
        if dns_records is not None:
            pulumi.set(__self__, "dns_records", dns_records)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if ips is not None:
            pulumi.set(__self__, "ips", ips)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter(name="pnId")
    def pn_id(self) -> str:
        """
        The ID of the Private Network.
        """
        return pulumi.get(self, "pn_id")

    @property
    @pulumi.getter(name="dnsRecords")
    def dns_records(self) -> Optional[Sequence[str]]:
        """
        List of DNS records for your endpoint.
        """
        return pulumi.get(self, "dns_records")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of the endpoint.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ips(self) -> Optional[Sequence[str]]:
        """
        List of IP addresses for your endpoint.
        """
        return pulumi.get(self, "ips")

    @property
    @pulumi.getter
    def port(self) -> Optional[int]:
        """
        TCP port of the endpoint.
        """
        return pulumi.get(self, "port")


@pulumi.output_type
class InstancePublicNetwork(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dnsRecord":
            suggest = "dns_record"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstancePublicNetwork. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstancePublicNetwork.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstancePublicNetwork.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 dns_record: Optional[str] = None,
                 id: Optional[str] = None,
                 port: Optional[int] = None):
        """
        :param str dns_record: The DNS record of your endpoint
        :param str id: The ID of the endpoint.
        :param int port: TCP port of the endpoint.
        """
        if dns_record is not None:
            pulumi.set(__self__, "dns_record", dns_record)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter(name="dnsRecord")
    def dns_record(self) -> Optional[str]:
        """
        The DNS record of your endpoint
        """
        return pulumi.get(self, "dns_record")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of the endpoint.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def port(self) -> Optional[int]:
        """
        TCP port of the endpoint.
        """
        return pulumi.get(self, "port")


@pulumi.output_type
class GetInstancePrivateIpResult(dict):
    def __init__(__self__, *,
                 address: str,
                 id: str):
        """
        :param str address: The private IPv4 address
        :param str id: The ID of the MongoDB® Instance.
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def address(self) -> str:
        """
        The private IPv4 address
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the MongoDB® Instance.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class GetInstancePrivateNetworkResult(dict):
    def __init__(__self__, *,
                 dns_records: Sequence[str],
                 id: str,
                 ips: Sequence[str],
                 pn_id: str,
                 port: int):
        """
        :param Sequence[str] dns_records: List of DNS records for your endpoint
        :param str id: The ID of the MongoDB® Instance.
        :param Sequence[str] ips: List of IP addresses for your endpoint
        :param str pn_id: The private network ID
        :param int port: TCP port of the endpoint
        """
        pulumi.set(__self__, "dns_records", dns_records)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ips", ips)
        pulumi.set(__self__, "pn_id", pn_id)
        pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter(name="dnsRecords")
    def dns_records(self) -> Sequence[str]:
        """
        List of DNS records for your endpoint
        """
        return pulumi.get(self, "dns_records")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the MongoDB® Instance.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ips(self) -> Sequence[str]:
        """
        List of IP addresses for your endpoint
        """
        return pulumi.get(self, "ips")

    @property
    @pulumi.getter(name="pnId")
    def pn_id(self) -> str:
        """
        The private network ID
        """
        return pulumi.get(self, "pn_id")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        TCP port of the endpoint
        """
        return pulumi.get(self, "port")


@pulumi.output_type
class GetInstancePublicNetworkResult(dict):
    def __init__(__self__, *,
                 dns_record: str,
                 id: str,
                 port: int):
        """
        :param str dns_record: The DNS record of your endpoint
        :param str id: The ID of the MongoDB® Instance.
        :param int port: TCP port of the endpoint
        """
        pulumi.set(__self__, "dns_record", dns_record)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter(name="dnsRecord")
    def dns_record(self) -> str:
        """
        The DNS record of your endpoint
        """
        return pulumi.get(self, "dns_record")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the MongoDB® Instance.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        TCP port of the endpoint
        """
        return pulumi.get(self, "port")


