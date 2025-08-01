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
    'GetLbFrontendsResult',
    'AwaitableGetLbFrontendsResult',
    'get_lb_frontends',
    'get_lb_frontends_output',
]

warnings.warn("""scaleway.index/getlbfrontends.getLbFrontends has been deprecated in favor of scaleway.loadbalancers/getfrontends.getFrontends""", DeprecationWarning)

@pulumi.output_type
class GetLbFrontendsResult:
    """
    A collection of values returned by getLbFrontends.
    """
    def __init__(__self__, frontends=None, id=None, lb_id=None, name=None, organization_id=None, project_id=None, zone=None):
        if frontends and not isinstance(frontends, list):
            raise TypeError("Expected argument 'frontends' to be a list")
        pulumi.set(__self__, "frontends", frontends)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lb_id and not isinstance(lb_id, str):
            raise TypeError("Expected argument 'lb_id' to be a str")
        pulumi.set(__self__, "lb_id", lb_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def frontends(self) -> Sequence['outputs.GetLbFrontendsFrontendResult']:
        """
        List of retrieved frontends
        """
        return pulumi.get(self, "frontends")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> builtins.str:
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
    @pulumi.getter(name="projectId")
    def project_id(self) -> builtins.str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def zone(self) -> builtins.str:
        return pulumi.get(self, "zone")


class AwaitableGetLbFrontendsResult(GetLbFrontendsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLbFrontendsResult(
            frontends=self.frontends,
            id=self.id,
            lb_id=self.lb_id,
            name=self.name,
            organization_id=self.organization_id,
            project_id=self.project_id,
            zone=self.zone)


def get_lb_frontends(lb_id: Optional[builtins.str] = None,
                     name: Optional[builtins.str] = None,
                     project_id: Optional[builtins.str] = None,
                     zone: Optional[builtins.str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLbFrontendsResult:
    """
    Gets information about multiple Load Balancer frontends.

    For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/reference-content/configuring-frontends/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    # Find frontends that share the same LB ID
    by_lbid = scaleway.loadbalancers.get_frontends(lb_id=lb01["id"])
    # Find frontends by LB ID and name
    by_lbid_and_name = scaleway.loadbalancers.get_frontends(lb_id=lb01["id"],
        name="tf-frontend-datasource")
    ```


    :param builtins.str lb_id: The Load Balancer ID this frontend is attached to. Frontends with a matching ID are listed.
    :param builtins.str name: The frontend name to filter for. Frontends with a matching name are listed.
    :param builtins.str zone: `zone`) The zone in which the frontends exist.
    """
    pulumi.log.warn("""get_lb_frontends is deprecated: scaleway.index/getlbfrontends.getLbFrontends has been deprecated in favor of scaleway.loadbalancers/getfrontends.getFrontends""")
    __args__ = dict()
    __args__['lbId'] = lb_id
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getLbFrontends:getLbFrontends', __args__, opts=opts, typ=GetLbFrontendsResult).value

    return AwaitableGetLbFrontendsResult(
        frontends=pulumi.get(__ret__, 'frontends'),
        id=pulumi.get(__ret__, 'id'),
        lb_id=pulumi.get(__ret__, 'lb_id'),
        name=pulumi.get(__ret__, 'name'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        zone=pulumi.get(__ret__, 'zone'))
def get_lb_frontends_output(lb_id: Optional[pulumi.Input[builtins.str]] = None,
                            name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            project_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            zone: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetLbFrontendsResult]:
    """
    Gets information about multiple Load Balancer frontends.

    For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/reference-content/configuring-frontends/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    # Find frontends that share the same LB ID
    by_lbid = scaleway.loadbalancers.get_frontends(lb_id=lb01["id"])
    # Find frontends by LB ID and name
    by_lbid_and_name = scaleway.loadbalancers.get_frontends(lb_id=lb01["id"],
        name="tf-frontend-datasource")
    ```


    :param builtins.str lb_id: The Load Balancer ID this frontend is attached to. Frontends with a matching ID are listed.
    :param builtins.str name: The frontend name to filter for. Frontends with a matching name are listed.
    :param builtins.str zone: `zone`) The zone in which the frontends exist.
    """
    pulumi.log.warn("""get_lb_frontends is deprecated: scaleway.index/getlbfrontends.getLbFrontends has been deprecated in favor of scaleway.loadbalancers/getfrontends.getFrontends""")
    __args__ = dict()
    __args__['lbId'] = lb_id
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getLbFrontends:getLbFrontends', __args__, opts=opts, typ=GetLbFrontendsResult)
    return __ret__.apply(lambda __response__: GetLbFrontendsResult(
        frontends=pulumi.get(__response__, 'frontends'),
        id=pulumi.get(__response__, 'id'),
        lb_id=pulumi.get(__response__, 'lb_id'),
        name=pulumi.get(__response__, 'name'),
        organization_id=pulumi.get(__response__, 'organization_id'),
        project_id=pulumi.get(__response__, 'project_id'),
        zone=pulumi.get(__response__, 'zone')))
