# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
from . import _utilities
from . import outputs

__all__ = [
    'GetLoadbalancerResult',
    'AwaitableGetLoadbalancerResult',
    'get_loadbalancer',
    'get_loadbalancer_output',
]

warnings.warn("""scaleway.index/getloadbalancer.getLoadbalancer has been deprecated in favor of scaleway.loadbalancers/getloadbalancer.getLoadBalancer""", DeprecationWarning)

@pulumi.output_type
class GetLoadbalancerResult:
    """
    A collection of values returned by getLoadbalancer.
    """
    def __init__(__self__, assign_flexible_ip=None, assign_flexible_ipv6=None, description=None, id=None, ip_address=None, ip_id=None, ip_ids=None, ipv6_address=None, lb_id=None, name=None, organization_id=None, private_ips=None, private_networks=None, project_id=None, region=None, release_ip=None, ssl_compatibility_level=None, tags=None, type=None, zone=None):
        if assign_flexible_ip and not isinstance(assign_flexible_ip, bool):
            raise TypeError("Expected argument 'assign_flexible_ip' to be a bool")
        pulumi.set(__self__, "assign_flexible_ip", assign_flexible_ip)
        if assign_flexible_ipv6 and not isinstance(assign_flexible_ipv6, bool):
            raise TypeError("Expected argument 'assign_flexible_ipv6' to be a bool")
        pulumi.set(__self__, "assign_flexible_ipv6", assign_flexible_ipv6)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_address and not isinstance(ip_address, str):
            raise TypeError("Expected argument 'ip_address' to be a str")
        pulumi.set(__self__, "ip_address", ip_address)
        if ip_id and not isinstance(ip_id, str):
            raise TypeError("Expected argument 'ip_id' to be a str")
        pulumi.set(__self__, "ip_id", ip_id)
        if ip_ids and not isinstance(ip_ids, list):
            raise TypeError("Expected argument 'ip_ids' to be a list")
        pulumi.set(__self__, "ip_ids", ip_ids)
        if ipv6_address and not isinstance(ipv6_address, str):
            raise TypeError("Expected argument 'ipv6_address' to be a str")
        pulumi.set(__self__, "ipv6_address", ipv6_address)
        if lb_id and not isinstance(lb_id, str):
            raise TypeError("Expected argument 'lb_id' to be a str")
        pulumi.set(__self__, "lb_id", lb_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if private_ips and not isinstance(private_ips, list):
            raise TypeError("Expected argument 'private_ips' to be a list")
        pulumi.set(__self__, "private_ips", private_ips)
        if private_networks and not isinstance(private_networks, list):
            raise TypeError("Expected argument 'private_networks' to be a list")
        pulumi.set(__self__, "private_networks", private_networks)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if release_ip and not isinstance(release_ip, bool):
            raise TypeError("Expected argument 'release_ip' to be a bool")
        pulumi.set(__self__, "release_ip", release_ip)
        if ssl_compatibility_level and not isinstance(ssl_compatibility_level, str):
            raise TypeError("Expected argument 'ssl_compatibility_level' to be a str")
        pulumi.set(__self__, "ssl_compatibility_level", ssl_compatibility_level)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="assignFlexibleIp")
    def assign_flexible_ip(self) -> builtins.bool:
        return pulumi.get(self, "assign_flexible_ip")

    @property
    @pulumi.getter(name="assignFlexibleIpv6")
    def assign_flexible_ipv6(self) -> builtins.bool:
        return pulumi.get(self, "assign_flexible_ipv6")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> builtins.str:
        """
        The Load Balancer public IP address.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="ipId")
    def ip_id(self) -> builtins.str:
        return pulumi.get(self, "ip_id")

    @property
    @pulumi.getter(name="ipIds")
    def ip_ids(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "ip_ids")

    @property
    @pulumi.getter(name="ipv6Address")
    def ipv6_address(self) -> builtins.str:
        return pulumi.get(self, "ipv6_address")

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "lb_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> builtins.str:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="privateIps")
    def private_ips(self) -> Sequence['outputs.GetLoadbalancerPrivateIpResult']:
        return pulumi.get(self, "private_ips")

    @property
    @pulumi.getter(name="privateNetworks")
    def private_networks(self) -> Sequence['outputs.GetLoadbalancerPrivateNetworkResult']:
        return pulumi.get(self, "private_networks")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="releaseIp")
    def release_ip(self) -> Optional[builtins.bool]:
        return pulumi.get(self, "release_ip")

    @property
    @pulumi.getter(name="sslCompatibilityLevel")
    def ssl_compatibility_level(self) -> builtins.str:
        return pulumi.get(self, "ssl_compatibility_level")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[builtins.str]:
        """
        The tags associated with the Load Balancer.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The Load Balancer type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def zone(self) -> Optional[builtins.str]:
        """
        (Defaults to provider `zone`) The zone in which the Load Balancer exists.
        """
        return pulumi.get(self, "zone")


class AwaitableGetLoadbalancerResult(GetLoadbalancerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLoadbalancerResult(
            assign_flexible_ip=self.assign_flexible_ip,
            assign_flexible_ipv6=self.assign_flexible_ipv6,
            description=self.description,
            id=self.id,
            ip_address=self.ip_address,
            ip_id=self.ip_id,
            ip_ids=self.ip_ids,
            ipv6_address=self.ipv6_address,
            lb_id=self.lb_id,
            name=self.name,
            organization_id=self.organization_id,
            private_ips=self.private_ips,
            private_networks=self.private_networks,
            project_id=self.project_id,
            region=self.region,
            release_ip=self.release_ip,
            ssl_compatibility_level=self.ssl_compatibility_level,
            tags=self.tags,
            type=self.type,
            zone=self.zone)


def get_loadbalancer(lb_id: Optional[builtins.str] = None,
                     name: Optional[builtins.str] = None,
                     project_id: Optional[builtins.str] = None,
                     release_ip: Optional[builtins.bool] = None,
                     zone: Optional[builtins.str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLoadbalancerResult:
    """
    Gets information about a Load Balancer.

    For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/concepts/#load-balancers) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-list-load-balancers).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    # Get info by name
    by_name = scaleway.loadbalancers.get_load_balancer(name="foobar")
    # Get info by ID
    by_id = scaleway.loadbalancers.get_load_balancer(lb_id="11111111-1111-1111-1111-111111111111")
    ```


    :param builtins.str name: The Load Balancer name.
    :param builtins.str project_id: The ID of the Project the Load Balancer is associated with.
    :param builtins.str zone: (Defaults to provider `zone`) The zone in which the Load Balancer exists.
    """
    pulumi.log.warn("""get_loadbalancer is deprecated: scaleway.index/getloadbalancer.getLoadbalancer has been deprecated in favor of scaleway.loadbalancers/getloadbalancer.getLoadBalancer""")
    __args__ = dict()
    __args__['lbId'] = lb_id
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['releaseIp'] = release_ip
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getLoadbalancer:getLoadbalancer', __args__, opts=opts, typ=GetLoadbalancerResult).value

    return AwaitableGetLoadbalancerResult(
        assign_flexible_ip=pulumi.get(__ret__, 'assign_flexible_ip'),
        assign_flexible_ipv6=pulumi.get(__ret__, 'assign_flexible_ipv6'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        ip_address=pulumi.get(__ret__, 'ip_address'),
        ip_id=pulumi.get(__ret__, 'ip_id'),
        ip_ids=pulumi.get(__ret__, 'ip_ids'),
        ipv6_address=pulumi.get(__ret__, 'ipv6_address'),
        lb_id=pulumi.get(__ret__, 'lb_id'),
        name=pulumi.get(__ret__, 'name'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        private_ips=pulumi.get(__ret__, 'private_ips'),
        private_networks=pulumi.get(__ret__, 'private_networks'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        release_ip=pulumi.get(__ret__, 'release_ip'),
        ssl_compatibility_level=pulumi.get(__ret__, 'ssl_compatibility_level'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        zone=pulumi.get(__ret__, 'zone'))
def get_loadbalancer_output(lb_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            project_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            release_ip: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                            zone: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetLoadbalancerResult]:
    """
    Gets information about a Load Balancer.

    For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/concepts/#load-balancers) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-list-load-balancers).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    # Get info by name
    by_name = scaleway.loadbalancers.get_load_balancer(name="foobar")
    # Get info by ID
    by_id = scaleway.loadbalancers.get_load_balancer(lb_id="11111111-1111-1111-1111-111111111111")
    ```


    :param builtins.str name: The Load Balancer name.
    :param builtins.str project_id: The ID of the Project the Load Balancer is associated with.
    :param builtins.str zone: (Defaults to provider `zone`) The zone in which the Load Balancer exists.
    """
    pulumi.log.warn("""get_loadbalancer is deprecated: scaleway.index/getloadbalancer.getLoadbalancer has been deprecated in favor of scaleway.loadbalancers/getloadbalancer.getLoadBalancer""")
    __args__ = dict()
    __args__['lbId'] = lb_id
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['releaseIp'] = release_ip
    __args__['zone'] = zone
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getLoadbalancer:getLoadbalancer', __args__, opts=opts, typ=GetLoadbalancerResult)
    return __ret__.apply(lambda __response__: GetLoadbalancerResult(
        assign_flexible_ip=pulumi.get(__response__, 'assign_flexible_ip'),
        assign_flexible_ipv6=pulumi.get(__response__, 'assign_flexible_ipv6'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        ip_address=pulumi.get(__response__, 'ip_address'),
        ip_id=pulumi.get(__response__, 'ip_id'),
        ip_ids=pulumi.get(__response__, 'ip_ids'),
        ipv6_address=pulumi.get(__response__, 'ipv6_address'),
        lb_id=pulumi.get(__response__, 'lb_id'),
        name=pulumi.get(__response__, 'name'),
        organization_id=pulumi.get(__response__, 'organization_id'),
        private_ips=pulumi.get(__response__, 'private_ips'),
        private_networks=pulumi.get(__response__, 'private_networks'),
        project_id=pulumi.get(__response__, 'project_id'),
        region=pulumi.get(__response__, 'region'),
        release_ip=pulumi.get(__response__, 'release_ip'),
        ssl_compatibility_level=pulumi.get(__response__, 'ssl_compatibility_level'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        zone=pulumi.get(__response__, 'zone')))
