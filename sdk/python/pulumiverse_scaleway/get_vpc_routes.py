# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetVpcRoutesResult',
    'AwaitableGetVpcRoutesResult',
    'get_vpc_routes',
    'get_vpc_routes_output',
]

@pulumi.output_type
class GetVpcRoutesResult:
    """
    A collection of values returned by getVpcRoutes.
    """
    def __init__(__self__, id=None, is_ipv6=None, nexthop_private_network_id=None, nexthop_resource_id=None, nexthop_resource_type=None, region=None, routes=None, tags=None, vpc_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_ipv6 and not isinstance(is_ipv6, bool):
            raise TypeError("Expected argument 'is_ipv6' to be a bool")
        pulumi.set(__self__, "is_ipv6", is_ipv6)
        if nexthop_private_network_id and not isinstance(nexthop_private_network_id, str):
            raise TypeError("Expected argument 'nexthop_private_network_id' to be a str")
        pulumi.set(__self__, "nexthop_private_network_id", nexthop_private_network_id)
        if nexthop_resource_id and not isinstance(nexthop_resource_id, str):
            raise TypeError("Expected argument 'nexthop_resource_id' to be a str")
        pulumi.set(__self__, "nexthop_resource_id", nexthop_resource_id)
        if nexthop_resource_type and not isinstance(nexthop_resource_type, str):
            raise TypeError("Expected argument 'nexthop_resource_type' to be a str")
        pulumi.set(__self__, "nexthop_resource_type", nexthop_resource_type)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if routes and not isinstance(routes, list):
            raise TypeError("Expected argument 'routes' to be a list")
        pulumi.set(__self__, "routes", routes)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isIpv6")
    def is_ipv6(self) -> Optional[bool]:
        return pulumi.get(self, "is_ipv6")

    @property
    @pulumi.getter(name="nexthopPrivateNetworkId")
    def nexthop_private_network_id(self) -> Optional[str]:
        return pulumi.get(self, "nexthop_private_network_id")

    @property
    @pulumi.getter(name="nexthopResourceId")
    def nexthop_resource_id(self) -> Optional[str]:
        return pulumi.get(self, "nexthop_resource_id")

    @property
    @pulumi.getter(name="nexthopResourceType")
    def nexthop_resource_type(self) -> Optional[str]:
        return pulumi.get(self, "nexthop_resource_type")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def routes(self) -> Sequence['outputs.GetVpcRoutesRouteResult']:
        """
        List of retrieved routes
        """
        return pulumi.get(self, "routes")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        return pulumi.get(self, "vpc_id")


class AwaitableGetVpcRoutesResult(GetVpcRoutesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcRoutesResult(
            id=self.id,
            is_ipv6=self.is_ipv6,
            nexthop_private_network_id=self.nexthop_private_network_id,
            nexthop_resource_id=self.nexthop_resource_id,
            nexthop_resource_type=self.nexthop_resource_type,
            region=self.region,
            routes=self.routes,
            tags=self.tags,
            vpc_id=self.vpc_id)


def get_vpc_routes(is_ipv6: Optional[bool] = None,
                   nexthop_private_network_id: Optional[str] = None,
                   nexthop_resource_id: Optional[str] = None,
                   nexthop_resource_type: Optional[str] = None,
                   region: Optional[str] = None,
                   tags: Optional[Sequence[str]] = None,
                   vpc_id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcRoutesResult:
    """
    Gets information about multiple VPC routes.


    :param bool is_ipv6: Routes with an IPv6 destination will be listed.
    :param str nexthop_private_network_id: The next hop private network ID to filter for. routes with a similar next hop private network ID are listed.
    :param str nexthop_resource_id: The next hop resource ID to filter for. routes with a similar next hop resource ID are listed.
    :param str nexthop_resource_type: The next hop resource type to filter for. routes with a similar next hop resource type are listed.
    :param str region: `region`). The region in which the routes exist.
    :param Sequence[str] tags: List of tags to filter for. routes with these exact tags are listed.
    :param str vpc_id: The VPC ID to filter for. routes with a similar VPC ID are listed.
    """
    __args__ = dict()
    __args__['isIpv6'] = is_ipv6
    __args__['nexthopPrivateNetworkId'] = nexthop_private_network_id
    __args__['nexthopResourceId'] = nexthop_resource_id
    __args__['nexthopResourceType'] = nexthop_resource_type
    __args__['region'] = region
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getVpcRoutes:getVpcRoutes', __args__, opts=opts, typ=GetVpcRoutesResult).value

    return AwaitableGetVpcRoutesResult(
        id=pulumi.get(__ret__, 'id'),
        is_ipv6=pulumi.get(__ret__, 'is_ipv6'),
        nexthop_private_network_id=pulumi.get(__ret__, 'nexthop_private_network_id'),
        nexthop_resource_id=pulumi.get(__ret__, 'nexthop_resource_id'),
        nexthop_resource_type=pulumi.get(__ret__, 'nexthop_resource_type'),
        region=pulumi.get(__ret__, 'region'),
        routes=pulumi.get(__ret__, 'routes'),
        tags=pulumi.get(__ret__, 'tags'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'))


@_utilities.lift_output_func(get_vpc_routes)
def get_vpc_routes_output(is_ipv6: Optional[pulumi.Input[Optional[bool]]] = None,
                          nexthop_private_network_id: Optional[pulumi.Input[Optional[str]]] = None,
                          nexthop_resource_id: Optional[pulumi.Input[Optional[str]]] = None,
                          nexthop_resource_type: Optional[pulumi.Input[Optional[str]]] = None,
                          region: Optional[pulumi.Input[Optional[str]]] = None,
                          tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                          vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcRoutesResult]:
    """
    Gets information about multiple VPC routes.


    :param bool is_ipv6: Routes with an IPv6 destination will be listed.
    :param str nexthop_private_network_id: The next hop private network ID to filter for. routes with a similar next hop private network ID are listed.
    :param str nexthop_resource_id: The next hop resource ID to filter for. routes with a similar next hop resource ID are listed.
    :param str nexthop_resource_type: The next hop resource type to filter for. routes with a similar next hop resource type are listed.
    :param str region: `region`). The region in which the routes exist.
    :param Sequence[str] tags: List of tags to filter for. routes with these exact tags are listed.
    :param str vpc_id: The VPC ID to filter for. routes with a similar VPC ID are listed.
    """
    ...