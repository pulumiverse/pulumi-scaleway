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
from .. import _utilities

__all__ = [
    'GetAvailabilityZonesResult',
    'AwaitableGetAvailabilityZonesResult',
    'get_availability_zones',
    'get_availability_zones_output',
]

@pulumi.output_type
class GetAvailabilityZonesResult:
    """
    A collection of values returned by getAvailabilityZones.
    """
    def __init__(__self__, id=None, region=None, zones=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def region(self) -> Optional[builtins.str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def zones(self) -> Sequence[builtins.str]:
        """
        The list of availability zones in each Region
        """
        return pulumi.get(self, "zones")


class AwaitableGetAvailabilityZonesResult(GetAvailabilityZonesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAvailabilityZonesResult(
            id=self.id,
            region=self.region,
            zones=self.zones)


def get_availability_zones(region: Optional[builtins.str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAvailabilityZonesResult:
    """
    The `account_get_availability_zones` data source is used to retrieve information about the available zones based on its Region.

    For technical and legal reasons, some products are split by Region or by Availability Zones. When using such product,
    you can choose the location that better fits your need (country, latency, etc.).

    Refer to the Account [documentation](https://www.scaleway.com/en/docs/console/account/reference-content/products-availability/) for more information.

    ## Retrieve the Availability Zones of a Region

    The following command allow you to retrieve a the AZs of a Region.

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    # Get info by Region key
    main = scaleway.account.get_availability_zones(region="nl-ams")
    ```


    :param builtins.str region: Region is represented as a Geographical area, such as France. Defaults to `fr-par`.
    """
    __args__ = dict()
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:account/getAvailabilityZones:getAvailabilityZones', __args__, opts=opts, typ=GetAvailabilityZonesResult).value

    return AwaitableGetAvailabilityZonesResult(
        id=pulumi.get(__ret__, 'id'),
        region=pulumi.get(__ret__, 'region'),
        zones=pulumi.get(__ret__, 'zones'))
def get_availability_zones_output(region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAvailabilityZonesResult]:
    """
    The `account_get_availability_zones` data source is used to retrieve information about the available zones based on its Region.

    For technical and legal reasons, some products are split by Region or by Availability Zones. When using such product,
    you can choose the location that better fits your need (country, latency, etc.).

    Refer to the Account [documentation](https://www.scaleway.com/en/docs/console/account/reference-content/products-availability/) for more information.

    ## Retrieve the Availability Zones of a Region

    The following command allow you to retrieve a the AZs of a Region.

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    # Get info by Region key
    main = scaleway.account.get_availability_zones(region="nl-ams")
    ```


    :param builtins.str region: Region is represented as a Geographical area, such as France. Defaults to `fr-par`.
    """
    __args__ = dict()
    __args__['region'] = region
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:account/getAvailabilityZones:getAvailabilityZones', __args__, opts=opts, typ=GetAvailabilityZonesResult)
    return __ret__.apply(lambda __response__: GetAvailabilityZonesResult(
        id=pulumi.get(__response__, 'id'),
        region=pulumi.get(__response__, 'region'),
        zones=pulumi.get(__response__, 'zones')))
