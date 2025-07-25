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

__all__ = ['ContainerCronArgs', 'ContainerCron']

@pulumi.input_type
class ContainerCronArgs:
    def __init__(__self__, *,
                 args: pulumi.Input[builtins.str],
                 container_id: pulumi.Input[builtins.str],
                 schedule: pulumi.Input[builtins.str],
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ContainerCron resource.
        :param pulumi.Input[builtins.str] args: The key-value mapping to define arguments that will be passed to your container’s event object
        :param pulumi.Input[builtins.str] container_id: The unique identifier of the container to link to your CRON trigger.
        :param pulumi.Input[builtins.str] schedule: CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/containers/reference-content/cron-schedules/) for more information).
        :param pulumi.Input[builtins.str] name: The name of the container CRON trigger. If not provided, a random name is generated.
        :param pulumi.Input[builtins.str] region: (Defaults to provider `region`) The region
               in which the CRON trigger is created.
        """
        pulumi.set(__self__, "args", args)
        pulumi.set(__self__, "container_id", container_id)
        pulumi.set(__self__, "schedule", schedule)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def args(self) -> pulumi.Input[builtins.str]:
        """
        The key-value mapping to define arguments that will be passed to your container’s event object
        """
        return pulumi.get(self, "args")

    @args.setter
    def args(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "args", value)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> pulumi.Input[builtins.str]:
        """
        The unique identifier of the container to link to your CRON trigger.
        """
        return pulumi.get(self, "container_id")

    @container_id.setter
    def container_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "container_id", value)

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Input[builtins.str]:
        """
        CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/containers/reference-content/cron-schedules/) for more information).
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the container CRON trigger. If not provided, a random name is generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        (Defaults to provider `region`) The region
        in which the CRON trigger is created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _ContainerCronState:
    def __init__(__self__, *,
                 args: Optional[pulumi.Input[builtins.str]] = None,
                 container_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 schedule: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering ContainerCron resources.
        :param pulumi.Input[builtins.str] args: The key-value mapping to define arguments that will be passed to your container’s event object
        :param pulumi.Input[builtins.str] container_id: The unique identifier of the container to link to your CRON trigger.
        :param pulumi.Input[builtins.str] name: The name of the container CRON trigger. If not provided, a random name is generated.
        :param pulumi.Input[builtins.str] region: (Defaults to provider `region`) The region
               in which the CRON trigger is created.
        :param pulumi.Input[builtins.str] schedule: CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/containers/reference-content/cron-schedules/) for more information).
        :param pulumi.Input[builtins.str] status: The CRON status.
        """
        if args is not None:
            pulumi.set(__self__, "args", args)
        if container_id is not None:
            pulumi.set(__self__, "container_id", container_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def args(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The key-value mapping to define arguments that will be passed to your container’s event object
        """
        return pulumi.get(self, "args")

    @args.setter
    def args(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "args", value)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The unique identifier of the container to link to your CRON trigger.
        """
        return pulumi.get(self, "container_id")

    @container_id.setter
    def container_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "container_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the container CRON trigger. If not provided, a random name is generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        (Defaults to provider `region`) The region
        in which the CRON trigger is created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/containers/reference-content/cron-schedules/) for more information).
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The CRON status.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)


warnings.warn("""scaleway.index/containercron.ContainerCron has been deprecated in favor of scaleway.containers/cron.Cron""", DeprecationWarning)


@pulumi.type_token("scaleway:index/containerCron:ContainerCron")
class ContainerCron(pulumi.CustomResource):
    warnings.warn("""scaleway.index/containercron.ContainerCron has been deprecated in favor of scaleway.containers/cron.Cron""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 args: Optional[pulumi.Input[builtins.str]] = None,
                 container_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 schedule: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The `containers.Cron` resource allows you to create and manage CRON triggers for Scaleway [Serverless Containers](https://www.scaleway.com/en/docs/serverless/containers/).

        Refer to the Containers CRON triggers [documentation](https://www.scaleway.com/en/docs/serverless/containers/how-to/add-trigger-to-a-container/) and [API documentation](https://www.scaleway.com/en/developers/api/serverless-containers/#path-triggers-list-all-triggers) for more information.

        ## Example Usage

        The following command allows you to add a CRON trigger to a Serverless Container.

        ```python
        import pulumi
        import json
        import pulumiverse_scaleway as scaleway

        main = scaleway.containers.Namespace("main")
        main_container = scaleway.containers.Container("main",
            name="my-container-with-cron-tf",
            namespace_id=main.id)
        main_cron = scaleway.containers.Cron("main",
            container_id=main_container.id,
            name="my-cron-name",
            schedule="5 4 1 * *",
            args=json.dumps({
                "address": {
                    "city": "Paris",
                    "country": "FR",
                },
                "age": 23,
                "firstName": "John",
                "isAlive": True,
                "lastName": "Smith",
            }))
        ```

        ## Import

        Container Cron can be imported using `{region}/{id}`, as shown below:

        bash

        ```sh
        $ pulumi import scaleway:index/containerCron:ContainerCron main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] args: The key-value mapping to define arguments that will be passed to your container’s event object
        :param pulumi.Input[builtins.str] container_id: The unique identifier of the container to link to your CRON trigger.
        :param pulumi.Input[builtins.str] name: The name of the container CRON trigger. If not provided, a random name is generated.
        :param pulumi.Input[builtins.str] region: (Defaults to provider `region`) The region
               in which the CRON trigger is created.
        :param pulumi.Input[builtins.str] schedule: CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/containers/reference-content/cron-schedules/) for more information).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ContainerCronArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `containers.Cron` resource allows you to create and manage CRON triggers for Scaleway [Serverless Containers](https://www.scaleway.com/en/docs/serverless/containers/).

        Refer to the Containers CRON triggers [documentation](https://www.scaleway.com/en/docs/serverless/containers/how-to/add-trigger-to-a-container/) and [API documentation](https://www.scaleway.com/en/developers/api/serverless-containers/#path-triggers-list-all-triggers) for more information.

        ## Example Usage

        The following command allows you to add a CRON trigger to a Serverless Container.

        ```python
        import pulumi
        import json
        import pulumiverse_scaleway as scaleway

        main = scaleway.containers.Namespace("main")
        main_container = scaleway.containers.Container("main",
            name="my-container-with-cron-tf",
            namespace_id=main.id)
        main_cron = scaleway.containers.Cron("main",
            container_id=main_container.id,
            name="my-cron-name",
            schedule="5 4 1 * *",
            args=json.dumps({
                "address": {
                    "city": "Paris",
                    "country": "FR",
                },
                "age": 23,
                "firstName": "John",
                "isAlive": True,
                "lastName": "Smith",
            }))
        ```

        ## Import

        Container Cron can be imported using `{region}/{id}`, as shown below:

        bash

        ```sh
        $ pulumi import scaleway:index/containerCron:ContainerCron main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param ContainerCronArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContainerCronArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 args: Optional[pulumi.Input[builtins.str]] = None,
                 container_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 schedule: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        pulumi.log.warn("""ContainerCron is deprecated: scaleway.index/containercron.ContainerCron has been deprecated in favor of scaleway.containers/cron.Cron""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ContainerCronArgs.__new__(ContainerCronArgs)

            if args is None and not opts.urn:
                raise TypeError("Missing required property 'args'")
            __props__.__dict__["args"] = args
            if container_id is None and not opts.urn:
                raise TypeError("Missing required property 'container_id'")
            __props__.__dict__["container_id"] = container_id
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            if schedule is None and not opts.urn:
                raise TypeError("Missing required property 'schedule'")
            __props__.__dict__["schedule"] = schedule
            __props__.__dict__["status"] = None
        super(ContainerCron, __self__).__init__(
            'scaleway:index/containerCron:ContainerCron',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            args: Optional[pulumi.Input[builtins.str]] = None,
            container_id: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            schedule: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None) -> 'ContainerCron':
        """
        Get an existing ContainerCron resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] args: The key-value mapping to define arguments that will be passed to your container’s event object
        :param pulumi.Input[builtins.str] container_id: The unique identifier of the container to link to your CRON trigger.
        :param pulumi.Input[builtins.str] name: The name of the container CRON trigger. If not provided, a random name is generated.
        :param pulumi.Input[builtins.str] region: (Defaults to provider `region`) The region
               in which the CRON trigger is created.
        :param pulumi.Input[builtins.str] schedule: CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/containers/reference-content/cron-schedules/) for more information).
        :param pulumi.Input[builtins.str] status: The CRON status.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ContainerCronState.__new__(_ContainerCronState)

        __props__.__dict__["args"] = args
        __props__.__dict__["container_id"] = container_id
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["schedule"] = schedule
        __props__.__dict__["status"] = status
        return ContainerCron(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def args(self) -> pulumi.Output[builtins.str]:
        """
        The key-value mapping to define arguments that will be passed to your container’s event object
        """
        return pulumi.get(self, "args")

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> pulumi.Output[builtins.str]:
        """
        The unique identifier of the container to link to your CRON trigger.
        """
        return pulumi.get(self, "container_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the container CRON trigger. If not provided, a random name is generated.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        (Defaults to provider `region`) The region
        in which the CRON trigger is created.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[builtins.str]:
        """
        CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/containers/reference-content/cron-schedules/) for more information).
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        The CRON status.
        """
        return pulumi.get(self, "status")

