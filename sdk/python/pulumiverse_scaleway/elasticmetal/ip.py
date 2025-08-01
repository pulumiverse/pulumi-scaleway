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

__all__ = ['IpArgs', 'Ip']

@pulumi.input_type
class IpArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 is_ipv6: Optional[pulumi.Input[builtins.bool]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 reverse: Optional[pulumi.Input[builtins.str]] = None,
                 server_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Ip resource.
        :param pulumi.Input[builtins.str] description: A description of the flexible IP.
        :param pulumi.Input[builtins.bool] is_ipv6: Defines whether the flexible IP has an IPv6 address.
        :param pulumi.Input[builtins.str] project_id: `project_id`) The ID of the Project the Flexible IP is associated with.
        :param pulumi.Input[builtins.str] reverse: The reverse domain associated with this flexible IP.
        :param pulumi.Input[builtins.str] server_id: The ID of the associated server.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: A list of tags to apply to the flexible IP.
        :param pulumi.Input[builtins.str] zone: `zone`) The zone of the Flexible IP.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_ipv6 is not None:
            pulumi.set(__self__, "is_ipv6", is_ipv6)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if reverse is not None:
            pulumi.set(__self__, "reverse", reverse)
        if server_id is not None:
            pulumi.set(__self__, "server_id", server_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A description of the flexible IP.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="isIpv6")
    def is_ipv6(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Defines whether the flexible IP has an IPv6 address.
        """
        return pulumi.get(self, "is_ipv6")

    @is_ipv6.setter
    def is_ipv6(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "is_ipv6", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        `project_id`) The ID of the Project the Flexible IP is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def reverse(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The reverse domain associated with this flexible IP.
        """
        return pulumi.get(self, "reverse")

    @reverse.setter
    def reverse(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "reverse", value)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the associated server.
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "server_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A list of tags to apply to the flexible IP.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        `zone`) The zone of the Flexible IP.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _IpState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 ip_address: Optional[pulumi.Input[builtins.str]] = None,
                 is_ipv6: Optional[pulumi.Input[builtins.bool]] = None,
                 organization_id: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 reverse: Optional[pulumi.Input[builtins.str]] = None,
                 server_id: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 updated_at: Optional[pulumi.Input[builtins.str]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering Ip resources.
        :param pulumi.Input[builtins.str] created_at: The date and time of the creation of the Flexible IP (Format ISO 8601).
        :param pulumi.Input[builtins.str] description: A description of the flexible IP.
        :param pulumi.Input[builtins.str] ip_address: The IP address of the Flexible IP.
        :param pulumi.Input[builtins.bool] is_ipv6: Defines whether the flexible IP has an IPv6 address.
        :param pulumi.Input[builtins.str] organization_id: The organization of the Flexible IP.
        :param pulumi.Input[builtins.str] project_id: `project_id`) The ID of the Project the Flexible IP is associated with.
        :param pulumi.Input[builtins.str] reverse: The reverse domain associated with this flexible IP.
        :param pulumi.Input[builtins.str] server_id: The ID of the associated server.
        :param pulumi.Input[builtins.str] status: The status of the flexible IP.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: A list of tags to apply to the flexible IP.
        :param pulumi.Input[builtins.str] updated_at: The date and time of the last update of the Flexible IP (Format ISO 8601).
        :param pulumi.Input[builtins.str] zone: `zone`) The zone of the Flexible IP.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if is_ipv6 is not None:
            pulumi.set(__self__, "is_ipv6", is_ipv6)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if reverse is not None:
            pulumi.set(__self__, "reverse", reverse)
        if server_id is not None:
            pulumi.set(__self__, "server_id", server_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The date and time of the creation of the Flexible IP (Format ISO 8601).
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A description of the flexible IP.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The IP address of the Flexible IP.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="isIpv6")
    def is_ipv6(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Defines whether the flexible IP has an IPv6 address.
        """
        return pulumi.get(self, "is_ipv6")

    @is_ipv6.setter
    def is_ipv6(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "is_ipv6", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The organization of the Flexible IP.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        `project_id`) The ID of the Project the Flexible IP is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def reverse(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The reverse domain associated with this flexible IP.
        """
        return pulumi.get(self, "reverse")

    @reverse.setter
    def reverse(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "reverse", value)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the associated server.
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "server_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The status of the flexible IP.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A list of tags to apply to the flexible IP.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The date and time of the last update of the Flexible IP (Format ISO 8601).
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        `zone`) The zone of the Flexible IP.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "zone", value)


@pulumi.type_token("scaleway:elasticmetal/ip:Ip")
class Ip(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 is_ipv6: Optional[pulumi.Input[builtins.bool]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 reverse: Optional[pulumi.Input[builtins.str]] = None,
                 server_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway flexible IPs.
        For more information, see the [API documentation](https://www.scaleway.com/en/developers/api/elastic-metal-flexible-ip).

        > **Note:**
        Flexible IPs are exclusively available for Elastic Metal (bare metal) servers. They are not compatible with other Scaleway products.

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.elasticmetal.Ip("main", reverse="my-reverse.com")
        ```

        ### With zone

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.elasticmetal.Ip("main", zone="fr-par-2")
        ```

        ### With IPv6

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.elasticmetal.Ip("main", is_ipv6=True)
        ```

        ### With baremetal server

        ```python
        import pulumi
        import pulumi_scaleway as scaleway
        import pulumiverse_scaleway as scaleway

        main = scaleway.account.SshKey("main",
            name="main",
            public_key="ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAILHy/M5FVm5ydLGcal3e5LNcfTalbeN7QL/ZGCvDEdqJ foobar@example.com")
        by_id = scaleway.elasticmetal.get_os(zone="fr-par-2",
            name="Ubuntu",
            version="20.04 LTS (Focal Fossa)")
        my_offer = scaleway.elasticmetal.get_offer(zone="fr-par-2",
            name="EM-A210R-HDD")
        base = scaleway.elasticmetal.Server("base",
            zone="fr-par-2",
            offer=my_offer.offer_id,
            os=by_id.os_id,
            ssh_key_ids=main.id)
        main_ip = scaleway.elasticmetal.Ip("main",
            server_id=base.id,
            zone="fr-par-2")
        ```

        ## Import

        Flexible IPs can be imported using the `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:elasticmetal/ip:Ip main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: A description of the flexible IP.
        :param pulumi.Input[builtins.bool] is_ipv6: Defines whether the flexible IP has an IPv6 address.
        :param pulumi.Input[builtins.str] project_id: `project_id`) The ID of the Project the Flexible IP is associated with.
        :param pulumi.Input[builtins.str] reverse: The reverse domain associated with this flexible IP.
        :param pulumi.Input[builtins.str] server_id: The ID of the associated server.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: A list of tags to apply to the flexible IP.
        :param pulumi.Input[builtins.str] zone: `zone`) The zone of the Flexible IP.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[IpArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway flexible IPs.
        For more information, see the [API documentation](https://www.scaleway.com/en/developers/api/elastic-metal-flexible-ip).

        > **Note:**
        Flexible IPs are exclusively available for Elastic Metal (bare metal) servers. They are not compatible with other Scaleway products.

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.elasticmetal.Ip("main", reverse="my-reverse.com")
        ```

        ### With zone

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.elasticmetal.Ip("main", zone="fr-par-2")
        ```

        ### With IPv6

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.elasticmetal.Ip("main", is_ipv6=True)
        ```

        ### With baremetal server

        ```python
        import pulumi
        import pulumi_scaleway as scaleway
        import pulumiverse_scaleway as scaleway

        main = scaleway.account.SshKey("main",
            name="main",
            public_key="ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAILHy/M5FVm5ydLGcal3e5LNcfTalbeN7QL/ZGCvDEdqJ foobar@example.com")
        by_id = scaleway.elasticmetal.get_os(zone="fr-par-2",
            name="Ubuntu",
            version="20.04 LTS (Focal Fossa)")
        my_offer = scaleway.elasticmetal.get_offer(zone="fr-par-2",
            name="EM-A210R-HDD")
        base = scaleway.elasticmetal.Server("base",
            zone="fr-par-2",
            offer=my_offer.offer_id,
            os=by_id.os_id,
            ssh_key_ids=main.id)
        main_ip = scaleway.elasticmetal.Ip("main",
            server_id=base.id,
            zone="fr-par-2")
        ```

        ## Import

        Flexible IPs can be imported using the `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:elasticmetal/ip:Ip main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param IpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 is_ipv6: Optional[pulumi.Input[builtins.bool]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 reverse: Optional[pulumi.Input[builtins.str]] = None,
                 server_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpArgs.__new__(IpArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["is_ipv6"] = is_ipv6
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["reverse"] = reverse
            __props__.__dict__["server_id"] = server_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["zone"] = zone
            __props__.__dict__["created_at"] = None
            __props__.__dict__["ip_address"] = None
            __props__.__dict__["organization_id"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["updated_at"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="scaleway:index/flexibleIp:FlexibleIp")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Ip, __self__).__init__(
            'scaleway:elasticmetal/ip:Ip',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            ip_address: Optional[pulumi.Input[builtins.str]] = None,
            is_ipv6: Optional[pulumi.Input[builtins.bool]] = None,
            organization_id: Optional[pulumi.Input[builtins.str]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None,
            reverse: Optional[pulumi.Input[builtins.str]] = None,
            server_id: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            updated_at: Optional[pulumi.Input[builtins.str]] = None,
            zone: Optional[pulumi.Input[builtins.str]] = None) -> 'Ip':
        """
        Get an existing Ip resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] created_at: The date and time of the creation of the Flexible IP (Format ISO 8601).
        :param pulumi.Input[builtins.str] description: A description of the flexible IP.
        :param pulumi.Input[builtins.str] ip_address: The IP address of the Flexible IP.
        :param pulumi.Input[builtins.bool] is_ipv6: Defines whether the flexible IP has an IPv6 address.
        :param pulumi.Input[builtins.str] organization_id: The organization of the Flexible IP.
        :param pulumi.Input[builtins.str] project_id: `project_id`) The ID of the Project the Flexible IP is associated with.
        :param pulumi.Input[builtins.str] reverse: The reverse domain associated with this flexible IP.
        :param pulumi.Input[builtins.str] server_id: The ID of the associated server.
        :param pulumi.Input[builtins.str] status: The status of the flexible IP.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: A list of tags to apply to the flexible IP.
        :param pulumi.Input[builtins.str] updated_at: The date and time of the last update of the Flexible IP (Format ISO 8601).
        :param pulumi.Input[builtins.str] zone: `zone`) The zone of the Flexible IP.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpState.__new__(_IpState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["ip_address"] = ip_address
        __props__.__dict__["is_ipv6"] = is_ipv6
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["reverse"] = reverse
        __props__.__dict__["server_id"] = server_id
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["zone"] = zone
        return Ip(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        The date and time of the creation of the Flexible IP (Format ISO 8601).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        A description of the flexible IP.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[builtins.str]:
        """
        The IP address of the Flexible IP.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="isIpv6")
    def is_ipv6(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Defines whether the flexible IP has an IPv6 address.
        """
        return pulumi.get(self, "is_ipv6")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[builtins.str]:
        """
        The organization of the Flexible IP.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        """
        `project_id`) The ID of the Project the Flexible IP is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def reverse(self) -> pulumi.Output[builtins.str]:
        """
        The reverse domain associated with this flexible IP.
        """
        return pulumi.get(self, "reverse")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The ID of the associated server.
        """
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        The status of the flexible IP.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        A list of tags to apply to the flexible IP.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[builtins.str]:
        """
        The date and time of the last update of the Flexible IP (Format ISO 8601).
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[builtins.str]:
        """
        `zone`) The zone of the Flexible IP.
        """
        return pulumi.get(self, "zone")

