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

__all__ = [
    'GetLbRouteResult',
    'AwaitableGetLbRouteResult',
    'get_lb_route',
    'get_lb_route_output',
]

warnings.warn("""scaleway.index/getlbroute.getLbRoute has been deprecated in favor of scaleway.loadbalancers/getroute.getRoute""", DeprecationWarning)

@pulumi.output_type
class GetLbRouteResult:
    """
    A collection of values returned by getLbRoute.
    """
    def __init__(__self__, backend_id=None, created_at=None, frontend_id=None, id=None, match_host_header=None, match_path_begin=None, match_sni=None, match_subdomains=None, route_id=None, updated_at=None):
        if backend_id and not isinstance(backend_id, str):
            raise TypeError("Expected argument 'backend_id' to be a str")
        pulumi.set(__self__, "backend_id", backend_id)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if frontend_id and not isinstance(frontend_id, str):
            raise TypeError("Expected argument 'frontend_id' to be a str")
        pulumi.set(__self__, "frontend_id", frontend_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if match_host_header and not isinstance(match_host_header, str):
            raise TypeError("Expected argument 'match_host_header' to be a str")
        pulumi.set(__self__, "match_host_header", match_host_header)
        if match_path_begin and not isinstance(match_path_begin, str):
            raise TypeError("Expected argument 'match_path_begin' to be a str")
        pulumi.set(__self__, "match_path_begin", match_path_begin)
        if match_sni and not isinstance(match_sni, str):
            raise TypeError("Expected argument 'match_sni' to be a str")
        pulumi.set(__self__, "match_sni", match_sni)
        if match_subdomains and not isinstance(match_subdomains, bool):
            raise TypeError("Expected argument 'match_subdomains' to be a bool")
        pulumi.set(__self__, "match_subdomains", match_subdomains)
        if route_id and not isinstance(route_id, str):
            raise TypeError("Expected argument 'route_id' to be a str")
        pulumi.set(__self__, "route_id", route_id)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="backendId")
    def backend_id(self) -> builtins.str:
        return pulumi.get(self, "backend_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> builtins.str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="frontendId")
    def frontend_id(self) -> builtins.str:
        return pulumi.get(self, "frontend_id")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="matchHostHeader")
    def match_host_header(self) -> builtins.str:
        return pulumi.get(self, "match_host_header")

    @property
    @pulumi.getter(name="matchPathBegin")
    def match_path_begin(self) -> builtins.str:
        return pulumi.get(self, "match_path_begin")

    @property
    @pulumi.getter(name="matchSni")
    def match_sni(self) -> builtins.str:
        return pulumi.get(self, "match_sni")

    @property
    @pulumi.getter(name="matchSubdomains")
    def match_subdomains(self) -> builtins.bool:
        return pulumi.get(self, "match_subdomains")

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> builtins.str:
        return pulumi.get(self, "route_id")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> builtins.str:
        return pulumi.get(self, "updated_at")


class AwaitableGetLbRouteResult(GetLbRouteResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLbRouteResult(
            backend_id=self.backend_id,
            created_at=self.created_at,
            frontend_id=self.frontend_id,
            id=self.id,
            match_host_header=self.match_host_header,
            match_path_begin=self.match_path_begin,
            match_sni=self.match_sni,
            match_subdomains=self.match_subdomains,
            route_id=self.route_id,
            updated_at=self.updated_at)


def get_lb_route(route_id: Optional[builtins.str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLbRouteResult:
    """
    Get information about Scaleway Load Balancer routes.

    For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/how-to/create-manage-routes/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-route).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway
    import pulumiverse_scaleway as scaleway

    ip01 = scaleway.loadbalancers.Ip("ip01")
    lb01 = scaleway.loadbalancers.LoadBalancer("lb01",
        ip_id=ip01.id,
        name="test-lb",
        type="lb-s")
    bkd01 = scaleway.loadbalancers.Backend("bkd01",
        lb_id=lb01.id,
        forward_protocol="tcp",
        forward_port=80,
        proxy_protocol="none")
    frt01 = scaleway.loadbalancers.Frontend("frt01",
        lb_id=lb01.id,
        backend_id=bkd01.id,
        inbound_port=80)
    rt01 = scaleway.loadbalancers.Route("rt01",
        frontend_id=frt01.id,
        backend_id=bkd01.id,
        match_sni="sni.scaleway.com")
    by_id = scaleway.loadbalancers.get_route_output(route_id=rt01.id)
    ```


    :param builtins.str route_id: The route ID.
    """
    pulumi.log.warn("""get_lb_route is deprecated: scaleway.index/getlbroute.getLbRoute has been deprecated in favor of scaleway.loadbalancers/getroute.getRoute""")
    __args__ = dict()
    __args__['routeId'] = route_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getLbRoute:getLbRoute', __args__, opts=opts, typ=GetLbRouteResult).value

    return AwaitableGetLbRouteResult(
        backend_id=pulumi.get(__ret__, 'backend_id'),
        created_at=pulumi.get(__ret__, 'created_at'),
        frontend_id=pulumi.get(__ret__, 'frontend_id'),
        id=pulumi.get(__ret__, 'id'),
        match_host_header=pulumi.get(__ret__, 'match_host_header'),
        match_path_begin=pulumi.get(__ret__, 'match_path_begin'),
        match_sni=pulumi.get(__ret__, 'match_sni'),
        match_subdomains=pulumi.get(__ret__, 'match_subdomains'),
        route_id=pulumi.get(__ret__, 'route_id'),
        updated_at=pulumi.get(__ret__, 'updated_at'))
def get_lb_route_output(route_id: Optional[pulumi.Input[builtins.str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetLbRouteResult]:
    """
    Get information about Scaleway Load Balancer routes.

    For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/how-to/create-manage-routes/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-route).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway
    import pulumiverse_scaleway as scaleway

    ip01 = scaleway.loadbalancers.Ip("ip01")
    lb01 = scaleway.loadbalancers.LoadBalancer("lb01",
        ip_id=ip01.id,
        name="test-lb",
        type="lb-s")
    bkd01 = scaleway.loadbalancers.Backend("bkd01",
        lb_id=lb01.id,
        forward_protocol="tcp",
        forward_port=80,
        proxy_protocol="none")
    frt01 = scaleway.loadbalancers.Frontend("frt01",
        lb_id=lb01.id,
        backend_id=bkd01.id,
        inbound_port=80)
    rt01 = scaleway.loadbalancers.Route("rt01",
        frontend_id=frt01.id,
        backend_id=bkd01.id,
        match_sni="sni.scaleway.com")
    by_id = scaleway.loadbalancers.get_route_output(route_id=rt01.id)
    ```


    :param builtins.str route_id: The route ID.
    """
    pulumi.log.warn("""get_lb_route is deprecated: scaleway.index/getlbroute.getLbRoute has been deprecated in favor of scaleway.loadbalancers/getroute.getRoute""")
    __args__ = dict()
    __args__['routeId'] = route_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getLbRoute:getLbRoute', __args__, opts=opts, typ=GetLbRouteResult)
    return __ret__.apply(lambda __response__: GetLbRouteResult(
        backend_id=pulumi.get(__response__, 'backend_id'),
        created_at=pulumi.get(__response__, 'created_at'),
        frontend_id=pulumi.get(__response__, 'frontend_id'),
        id=pulumi.get(__response__, 'id'),
        match_host_header=pulumi.get(__response__, 'match_host_header'),
        match_path_begin=pulumi.get(__response__, 'match_path_begin'),
        match_sni=pulumi.get(__response__, 'match_sni'),
        match_subdomains=pulumi.get(__response__, 'match_subdomains'),
        route_id=pulumi.get(__response__, 'route_id'),
        updated_at=pulumi.get(__response__, 'updated_at')))
