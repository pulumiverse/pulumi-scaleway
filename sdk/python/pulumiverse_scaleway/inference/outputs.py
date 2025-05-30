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
from . import outputs

__all__ = [
    'DeploymentPrivateEndpoint',
    'DeploymentPublicEndpoint',
    'ModelNodesSupport',
    'ModelNodesSupportQuantization',
    'GetModelNodesSupportResult',
    'GetModelNodesSupportQuantizationResult',
]

@pulumi.output_type
class DeploymentPrivateEndpoint(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "disableAuth":
            suggest = "disable_auth"
        elif key == "privateNetworkId":
            suggest = "private_network_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DeploymentPrivateEndpoint. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DeploymentPrivateEndpoint.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DeploymentPrivateEndpoint.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 disable_auth: Optional[bool] = None,
                 id: Optional[str] = None,
                 private_network_id: Optional[str] = None,
                 url: Optional[str] = None):
        """
        :param bool disable_auth: Disable the authentication on the endpoint.
        :param str id: (Optional) The id of the public endpoint.
        :param str private_network_id: The ID of the private network to use.
        :param str url: (Optional) The URL of the endpoint.
        """
        if disable_auth is not None:
            pulumi.set(__self__, "disable_auth", disable_auth)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if private_network_id is not None:
            pulumi.set(__self__, "private_network_id", private_network_id)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="disableAuth")
    def disable_auth(self) -> Optional[bool]:
        """
        Disable the authentication on the endpoint.
        """
        return pulumi.get(self, "disable_auth")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        (Optional) The id of the public endpoint.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> Optional[str]:
        """
        The ID of the private network to use.
        """
        return pulumi.get(self, "private_network_id")

    @property
    @pulumi.getter
    def url(self) -> Optional[str]:
        """
        (Optional) The URL of the endpoint.
        """
        return pulumi.get(self, "url")


@pulumi.output_type
class DeploymentPublicEndpoint(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "disableAuth":
            suggest = "disable_auth"
        elif key == "isEnabled":
            suggest = "is_enabled"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DeploymentPublicEndpoint. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DeploymentPublicEndpoint.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DeploymentPublicEndpoint.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 disable_auth: Optional[bool] = None,
                 id: Optional[str] = None,
                 is_enabled: Optional[bool] = None,
                 url: Optional[str] = None):
        """
        :param bool disable_auth: Disable the authentication on the endpoint.
        :param str id: (Optional) The id of the public endpoint.
        :param bool is_enabled: Enable or disable public endpoint.
        :param str url: (Optional) The URL of the endpoint.
        """
        if disable_auth is not None:
            pulumi.set(__self__, "disable_auth", disable_auth)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="disableAuth")
    def disable_auth(self) -> Optional[bool]:
        """
        Disable the authentication on the endpoint.
        """
        return pulumi.get(self, "disable_auth")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        (Optional) The id of the public endpoint.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[bool]:
        """
        Enable or disable public endpoint.
        """
        return pulumi.get(self, "is_enabled")

    @property
    @pulumi.getter
    def url(self) -> Optional[str]:
        """
        (Optional) The URL of the endpoint.
        """
        return pulumi.get(self, "url")


@pulumi.output_type
class ModelNodesSupport(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "nodeTypeName":
            suggest = "node_type_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ModelNodesSupport. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ModelNodesSupport.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ModelNodesSupport.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 node_type_name: Optional[str] = None,
                 quantizations: Optional[Sequence['outputs.ModelNodesSupportQuantization']] = None):
        """
        :param str node_type_name: The type of node supported.
        :param Sequence['ModelNodesSupportQuantizationArgs'] quantizations: A list of supported quantization options, including:
        """
        if node_type_name is not None:
            pulumi.set(__self__, "node_type_name", node_type_name)
        if quantizations is not None:
            pulumi.set(__self__, "quantizations", quantizations)

    @property
    @pulumi.getter(name="nodeTypeName")
    def node_type_name(self) -> Optional[str]:
        """
        The type of node supported.
        """
        return pulumi.get(self, "node_type_name")

    @property
    @pulumi.getter
    def quantizations(self) -> Optional[Sequence['outputs.ModelNodesSupportQuantization']]:
        """
        A list of supported quantization options, including:
        """
        return pulumi.get(self, "quantizations")


@pulumi.output_type
class ModelNodesSupportQuantization(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maxContextSize":
            suggest = "max_context_size"
        elif key == "quantizationBits":
            suggest = "quantization_bits"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ModelNodesSupportQuantization. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ModelNodesSupportQuantization.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ModelNodesSupportQuantization.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 allowed: Optional[bool] = None,
                 max_context_size: Optional[int] = None,
                 quantization_bits: Optional[int] = None):
        """
        :param bool allowed: Whether this quantization is allowed.
        :param int max_context_size: Maximum context length supported by this quantization.
        :param int quantization_bits: Number of bits used for quantization (e.g., 8, 16).
        """
        if allowed is not None:
            pulumi.set(__self__, "allowed", allowed)
        if max_context_size is not None:
            pulumi.set(__self__, "max_context_size", max_context_size)
        if quantization_bits is not None:
            pulumi.set(__self__, "quantization_bits", quantization_bits)

    @property
    @pulumi.getter
    def allowed(self) -> Optional[bool]:
        """
        Whether this quantization is allowed.
        """
        return pulumi.get(self, "allowed")

    @property
    @pulumi.getter(name="maxContextSize")
    def max_context_size(self) -> Optional[int]:
        """
        Maximum context length supported by this quantization.
        """
        return pulumi.get(self, "max_context_size")

    @property
    @pulumi.getter(name="quantizationBits")
    def quantization_bits(self) -> Optional[int]:
        """
        Number of bits used for quantization (e.g., 8, 16).
        """
        return pulumi.get(self, "quantization_bits")


@pulumi.output_type
class GetModelNodesSupportResult(dict):
    def __init__(__self__, *,
                 node_type_name: str,
                 quantizations: Sequence['outputs.GetModelNodesSupportQuantizationResult']):
        """
        :param str node_type_name: The type of node supported.
        :param Sequence['GetModelNodesSupportQuantizationArgs'] quantizations: A list of supported quantization options, including:
        """
        pulumi.set(__self__, "node_type_name", node_type_name)
        pulumi.set(__self__, "quantizations", quantizations)

    @property
    @pulumi.getter(name="nodeTypeName")
    def node_type_name(self) -> str:
        """
        The type of node supported.
        """
        return pulumi.get(self, "node_type_name")

    @property
    @pulumi.getter
    def quantizations(self) -> Sequence['outputs.GetModelNodesSupportQuantizationResult']:
        """
        A list of supported quantization options, including:
        """
        return pulumi.get(self, "quantizations")


@pulumi.output_type
class GetModelNodesSupportQuantizationResult(dict):
    def __init__(__self__, *,
                 allowed: bool,
                 max_context_size: int,
                 quantization_bits: int):
        """
        :param bool allowed: Whether this quantization is allowed.
        :param int max_context_size: Maximum context length supported by this quantization.
        :param int quantization_bits: Number of bits used for quantization (e.g., 8, 16).
        """
        pulumi.set(__self__, "allowed", allowed)
        pulumi.set(__self__, "max_context_size", max_context_size)
        pulumi.set(__self__, "quantization_bits", quantization_bits)

    @property
    @pulumi.getter
    def allowed(self) -> bool:
        """
        Whether this quantization is allowed.
        """
        return pulumi.get(self, "allowed")

    @property
    @pulumi.getter(name="maxContextSize")
    def max_context_size(self) -> int:
        """
        Maximum context length supported by this quantization.
        """
        return pulumi.get(self, "max_context_size")

    @property
    @pulumi.getter(name="quantizationBits")
    def quantization_bits(self) -> int:
        """
        Number of bits used for quantization (e.g., 8, 16).
        """
        return pulumi.get(self, "quantization_bits")


