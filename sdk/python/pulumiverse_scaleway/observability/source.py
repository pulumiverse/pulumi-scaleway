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

__all__ = ['SourceArgs', 'Source']

@pulumi.input_type
class SourceArgs:
    def __init__(__self__, *,
                 retention_days: pulumi.Input[builtins.int],
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Source resource.
        :param pulumi.Input[builtins.int] retention_days: The number of days to retain data in the data source. Must be a value between 1 and 365. For more details on retention policies, please refer to the [Scaleway Retention Documentation](https://www.scaleway.com/en/docs/cockpit/concepts/#retention). Note: Changes to this field will force the creation of a new resource.
        :param pulumi.Input[builtins.str] name: The name of the data source.
        :param pulumi.Input[builtins.str] project_id: ) The ID of the Project the data source is associated with.
        :param pulumi.Input[builtins.str] region: ) The region where the data source is located.
        :param pulumi.Input[builtins.str] type: The [type](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-types) of data source. Possible values are: `metrics`, `logs`, or `traces`.
        """
        pulumi.set(__self__, "retention_days", retention_days)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> pulumi.Input[builtins.int]:
        """
        The number of days to retain data in the data source. Must be a value between 1 and 365. For more details on retention policies, please refer to the [Scaleway Retention Documentation](https://www.scaleway.com/en/docs/cockpit/concepts/#retention). Note: Changes to this field will force the creation of a new resource.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "retention_days", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the data source.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ) The ID of the Project the data source is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ) The region where the data source is located.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The [type](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-types) of data source. Possible values are: `metrics`, `logs`, or `traces`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _SourceState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 origin: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 push_url: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 retention_days: Optional[pulumi.Input[builtins.int]] = None,
                 synchronized_with_grafana: Optional[pulumi.Input[builtins.bool]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 updated_at: Optional[pulumi.Input[builtins.str]] = None,
                 url: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering Source resources.
        :param pulumi.Input[builtins.str] created_at: The date and time the data source was created (in RFC 3339 format).
        :param pulumi.Input[builtins.str] name: The name of the data source.
        :param pulumi.Input[builtins.str] origin: The origin of the Cockpit data source.
        :param pulumi.Input[builtins.str] project_id: ) The ID of the Project the data source is associated with.
        :param pulumi.Input[builtins.str] push_url: The URL endpoint used for pushing data to the Cockpit data source.
        :param pulumi.Input[builtins.str] region: ) The region where the data source is located.
        :param pulumi.Input[builtins.int] retention_days: The number of days to retain data in the data source. Must be a value between 1 and 365. For more details on retention policies, please refer to the [Scaleway Retention Documentation](https://www.scaleway.com/en/docs/cockpit/concepts/#retention). Note: Changes to this field will force the creation of a new resource.
        :param pulumi.Input[builtins.bool] synchronized_with_grafana: Indicates whether the data source is synchronized with Grafana.
        :param pulumi.Input[builtins.str] type: The [type](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-types) of data source. Possible values are: `metrics`, `logs`, or `traces`.
        :param pulumi.Input[builtins.str] updated_at: The date and time the data source was last updated (in RFC 3339 format).
        :param pulumi.Input[builtins.str] url: The URL of the Cockpit data source.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if origin is not None:
            pulumi.set(__self__, "origin", origin)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if push_url is not None:
            pulumi.set(__self__, "push_url", push_url)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)
        if synchronized_with_grafana is not None:
            pulumi.set(__self__, "synchronized_with_grafana", synchronized_with_grafana)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The date and time the data source was created (in RFC 3339 format).
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the data source.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def origin(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The origin of the Cockpit data source.
        """
        return pulumi.get(self, "origin")

    @origin.setter
    def origin(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "origin", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ) The ID of the Project the data source is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="pushUrl")
    def push_url(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The URL endpoint used for pushing data to the Cockpit data source.
        """
        return pulumi.get(self, "push_url")

    @push_url.setter
    def push_url(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "push_url", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ) The region where the data source is located.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The number of days to retain data in the data source. Must be a value between 1 and 365. For more details on retention policies, please refer to the [Scaleway Retention Documentation](https://www.scaleway.com/en/docs/cockpit/concepts/#retention). Note: Changes to this field will force the creation of a new resource.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "retention_days", value)

    @property
    @pulumi.getter(name="synchronizedWithGrafana")
    def synchronized_with_grafana(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Indicates whether the data source is synchronized with Grafana.
        """
        return pulumi.get(self, "synchronized_with_grafana")

    @synchronized_with_grafana.setter
    def synchronized_with_grafana(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "synchronized_with_grafana", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The [type](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-types) of data source. Possible values are: `metrics`, `logs`, or `traces`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The date and time the data source was last updated (in RFC 3339 format).
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The URL of the Cockpit data source.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "url", value)


@pulumi.type_token("scaleway:observability/source:Source")
class Source(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 retention_days: Optional[pulumi.Input[builtins.int]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The `observability.Source` resource allows you to create and manage [data sources](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-sources) in Scaleway's Cockpit.

        Refer to Cockpit's [product documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/) and [API documentation](https://www.scaleway.com/en/developers/api/cockpit/regional-api) for more information.

        ## Example Usage

        ### Create a data source

        The following command allows you to create a [metrics](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#metric) data source named `my-data-source` in a given Project.

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        project = scaleway.account.Project("project", name="test project data source")
        main = scaleway.observability.Source("main",
            project_id=project.id,
            name="my-data-source",
            type="metrics",
            retention_days=6)
        ```

        ## Import

        This section explains how to import a data source using the ID of the region it is located in, in the `{region}/{id}` format.

        bash

        ```sh
        $ pulumi import scaleway:observability/source:Source main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] name: The name of the data source.
        :param pulumi.Input[builtins.str] project_id: ) The ID of the Project the data source is associated with.
        :param pulumi.Input[builtins.str] region: ) The region where the data source is located.
        :param pulumi.Input[builtins.int] retention_days: The number of days to retain data in the data source. Must be a value between 1 and 365. For more details on retention policies, please refer to the [Scaleway Retention Documentation](https://www.scaleway.com/en/docs/cockpit/concepts/#retention). Note: Changes to this field will force the creation of a new resource.
        :param pulumi.Input[builtins.str] type: The [type](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-types) of data source. Possible values are: `metrics`, `logs`, or `traces`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `observability.Source` resource allows you to create and manage [data sources](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-sources) in Scaleway's Cockpit.

        Refer to Cockpit's [product documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/) and [API documentation](https://www.scaleway.com/en/developers/api/cockpit/regional-api) for more information.

        ## Example Usage

        ### Create a data source

        The following command allows you to create a [metrics](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#metric) data source named `my-data-source` in a given Project.

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        project = scaleway.account.Project("project", name="test project data source")
        main = scaleway.observability.Source("main",
            project_id=project.id,
            name="my-data-source",
            type="metrics",
            retention_days=6)
        ```

        ## Import

        This section explains how to import a data source using the ID of the region it is located in, in the `{region}/{id}` format.

        bash

        ```sh
        $ pulumi import scaleway:observability/source:Source main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param SourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 retention_days: Optional[pulumi.Input[builtins.int]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SourceArgs.__new__(SourceArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            if retention_days is None and not opts.urn:
                raise TypeError("Missing required property 'retention_days'")
            __props__.__dict__["retention_days"] = retention_days
            __props__.__dict__["type"] = type
            __props__.__dict__["created_at"] = None
            __props__.__dict__["origin"] = None
            __props__.__dict__["push_url"] = None
            __props__.__dict__["synchronized_with_grafana"] = None
            __props__.__dict__["updated_at"] = None
            __props__.__dict__["url"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="scaleway:index/cockpitSource:CockpitSource")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Source, __self__).__init__(
            'scaleway:observability/source:Source',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            origin: Optional[pulumi.Input[builtins.str]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None,
            push_url: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            retention_days: Optional[pulumi.Input[builtins.int]] = None,
            synchronized_with_grafana: Optional[pulumi.Input[builtins.bool]] = None,
            type: Optional[pulumi.Input[builtins.str]] = None,
            updated_at: Optional[pulumi.Input[builtins.str]] = None,
            url: Optional[pulumi.Input[builtins.str]] = None) -> 'Source':
        """
        Get an existing Source resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] created_at: The date and time the data source was created (in RFC 3339 format).
        :param pulumi.Input[builtins.str] name: The name of the data source.
        :param pulumi.Input[builtins.str] origin: The origin of the Cockpit data source.
        :param pulumi.Input[builtins.str] project_id: ) The ID of the Project the data source is associated with.
        :param pulumi.Input[builtins.str] push_url: The URL endpoint used for pushing data to the Cockpit data source.
        :param pulumi.Input[builtins.str] region: ) The region where the data source is located.
        :param pulumi.Input[builtins.int] retention_days: The number of days to retain data in the data source. Must be a value between 1 and 365. For more details on retention policies, please refer to the [Scaleway Retention Documentation](https://www.scaleway.com/en/docs/cockpit/concepts/#retention). Note: Changes to this field will force the creation of a new resource.
        :param pulumi.Input[builtins.bool] synchronized_with_grafana: Indicates whether the data source is synchronized with Grafana.
        :param pulumi.Input[builtins.str] type: The [type](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-types) of data source. Possible values are: `metrics`, `logs`, or `traces`.
        :param pulumi.Input[builtins.str] updated_at: The date and time the data source was last updated (in RFC 3339 format).
        :param pulumi.Input[builtins.str] url: The URL of the Cockpit data source.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SourceState.__new__(_SourceState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["name"] = name
        __props__.__dict__["origin"] = origin
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["push_url"] = push_url
        __props__.__dict__["region"] = region
        __props__.__dict__["retention_days"] = retention_days
        __props__.__dict__["synchronized_with_grafana"] = synchronized_with_grafana
        __props__.__dict__["type"] = type
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["url"] = url
        return Source(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        The date and time the data source was created (in RFC 3339 format).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the data source.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def origin(self) -> pulumi.Output[builtins.str]:
        """
        The origin of the Cockpit data source.
        """
        return pulumi.get(self, "origin")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        """
        ) The ID of the Project the data source is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="pushUrl")
    def push_url(self) -> pulumi.Output[builtins.str]:
        """
        The URL endpoint used for pushing data to the Cockpit data source.
        """
        return pulumi.get(self, "push_url")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        ) The region where the data source is located.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> pulumi.Output[builtins.int]:
        """
        The number of days to retain data in the data source. Must be a value between 1 and 365. For more details on retention policies, please refer to the [Scaleway Retention Documentation](https://www.scaleway.com/en/docs/cockpit/concepts/#retention). Note: Changes to this field will force the creation of a new resource.
        """
        return pulumi.get(self, "retention_days")

    @property
    @pulumi.getter(name="synchronizedWithGrafana")
    def synchronized_with_grafana(self) -> pulumi.Output[builtins.bool]:
        """
        Indicates whether the data source is synchronized with Grafana.
        """
        return pulumi.get(self, "synchronized_with_grafana")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The [type](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-types) of data source. Possible values are: `metrics`, `logs`, or `traces`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[builtins.str]:
        """
        The date and time the data source was last updated (in RFC 3339 format).
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[builtins.str]:
        """
        The URL of the Cockpit data source.
        """
        return pulumi.get(self, "url")

