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

__all__ = ['FunctionDomainArgs', 'FunctionDomain']

@pulumi.input_type
class FunctionDomainArgs:
    def __init__(__self__, *,
                 function_id: pulumi.Input[builtins.str],
                 hostname: pulumi.Input[builtins.str],
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a FunctionDomain resource.
        :param pulumi.Input[builtins.str] function_id: The unique identifier of the function.
        :param pulumi.Input[builtins.str] hostname: The hostname with a CNAME record.
               
               We recommend you use a CNAME domain record that point to your native function `domain_name` for it.
               
               > **Important** Updating the `function_id` or `hostname` arguments will recreate the domain.
        :param pulumi.Input[builtins.str] region: (Defaults to provider `region`) The region in which the domain was created.
        """
        pulumi.set(__self__, "function_id", function_id)
        pulumi.set(__self__, "hostname", hostname)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> pulumi.Input[builtins.str]:
        """
        The unique identifier of the function.
        """
        return pulumi.get(self, "function_id")

    @function_id.setter
    def function_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "function_id", value)

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Input[builtins.str]:
        """
        The hostname with a CNAME record.

        We recommend you use a CNAME domain record that point to your native function `domain_name` for it.

        > **Important** Updating the `function_id` or `hostname` arguments will recreate the domain.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        (Defaults to provider `region`) The region in which the domain was created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _FunctionDomainState:
    def __init__(__self__, *,
                 function_id: Optional[pulumi.Input[builtins.str]] = None,
                 hostname: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 url: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering FunctionDomain resources.
        :param pulumi.Input[builtins.str] function_id: The unique identifier of the function.
        :param pulumi.Input[builtins.str] hostname: The hostname with a CNAME record.
               
               We recommend you use a CNAME domain record that point to your native function `domain_name` for it.
               
               > **Important** Updating the `function_id` or `hostname` arguments will recreate the domain.
        :param pulumi.Input[builtins.str] region: (Defaults to provider `region`) The region in which the domain was created.
        :param pulumi.Input[builtins.str] url: The URL used to query the function.
        """
        if function_id is not None:
            pulumi.set(__self__, "function_id", function_id)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The unique identifier of the function.
        """
        return pulumi.get(self, "function_id")

    @function_id.setter
    def function_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "function_id", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The hostname with a CNAME record.

        We recommend you use a CNAME domain record that point to your native function `domain_name` for it.

        > **Important** Updating the `function_id` or `hostname` arguments will recreate the domain.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        (Defaults to provider `region`) The region in which the domain was created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The URL used to query the function.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "url", value)


warnings.warn("""scaleway.index/functiondomain.FunctionDomain has been deprecated in favor of scaleway.functions/domain.Domain""", DeprecationWarning)


@pulumi.type_token("scaleway:index/functionDomain:FunctionDomain")
class FunctionDomain(pulumi.CustomResource):
    warnings.warn("""scaleway.index/functiondomain.FunctionDomain has been deprecated in favor of scaleway.functions/domain.Domain""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 function_id: Optional[pulumi.Input[builtins.str]] = None,
                 hostname: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The `functions.Domain` resource allows you to create and manage domain name bindings for Scaleway [Serverless Functions](https://www.scaleway.com/en/docs/serverless/functions/).

        Refer to the Functions domain [documentation](https://www.scaleway.com/en/docs/serverless/functions/how-to/add-a-custom-domain-name-to-a-function/) and the [API documentation](https://www.scaleway.com/en/developers/api/serverless-functions/#path-domains-list-all-domain-name-bindings) for more information.

        ## Example Usage

        This command allows to bind a custom domain name to a function.

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main_namespace = scaleway.functions.Namespace("main")
        main_function = scaleway.functions.Function("main",
            namespace_id=main_namespace.id,
            runtime="go118",
            privacy="private",
            handler="Handle",
            zip_file="testfixture/gofunction.zip",
            deploy=True)
        main = scaleway.functions.Domain("main",
            function_id=main_function.id,
            hostname="example.com",
            opts = pulumi.ResourceOptions(depends_on=[main_function]))
        ```

        ## Import

        Function domain binding can be imported using `{region}/{id}`, as shown below:

        bash

        ```sh
        $ pulumi import scaleway:index/functionDomain:FunctionDomain main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] function_id: The unique identifier of the function.
        :param pulumi.Input[builtins.str] hostname: The hostname with a CNAME record.
               
               We recommend you use a CNAME domain record that point to your native function `domain_name` for it.
               
               > **Important** Updating the `function_id` or `hostname` arguments will recreate the domain.
        :param pulumi.Input[builtins.str] region: (Defaults to provider `region`) The region in which the domain was created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FunctionDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `functions.Domain` resource allows you to create and manage domain name bindings for Scaleway [Serverless Functions](https://www.scaleway.com/en/docs/serverless/functions/).

        Refer to the Functions domain [documentation](https://www.scaleway.com/en/docs/serverless/functions/how-to/add-a-custom-domain-name-to-a-function/) and the [API documentation](https://www.scaleway.com/en/developers/api/serverless-functions/#path-domains-list-all-domain-name-bindings) for more information.

        ## Example Usage

        This command allows to bind a custom domain name to a function.

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main_namespace = scaleway.functions.Namespace("main")
        main_function = scaleway.functions.Function("main",
            namespace_id=main_namespace.id,
            runtime="go118",
            privacy="private",
            handler="Handle",
            zip_file="testfixture/gofunction.zip",
            deploy=True)
        main = scaleway.functions.Domain("main",
            function_id=main_function.id,
            hostname="example.com",
            opts = pulumi.ResourceOptions(depends_on=[main_function]))
        ```

        ## Import

        Function domain binding can be imported using `{region}/{id}`, as shown below:

        bash

        ```sh
        $ pulumi import scaleway:index/functionDomain:FunctionDomain main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param FunctionDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FunctionDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 function_id: Optional[pulumi.Input[builtins.str]] = None,
                 hostname: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        pulumi.log.warn("""FunctionDomain is deprecated: scaleway.index/functiondomain.FunctionDomain has been deprecated in favor of scaleway.functions/domain.Domain""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FunctionDomainArgs.__new__(FunctionDomainArgs)

            if function_id is None and not opts.urn:
                raise TypeError("Missing required property 'function_id'")
            __props__.__dict__["function_id"] = function_id
            if hostname is None and not opts.urn:
                raise TypeError("Missing required property 'hostname'")
            __props__.__dict__["hostname"] = hostname
            __props__.__dict__["region"] = region
            __props__.__dict__["url"] = None
        super(FunctionDomain, __self__).__init__(
            'scaleway:index/functionDomain:FunctionDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            function_id: Optional[pulumi.Input[builtins.str]] = None,
            hostname: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            url: Optional[pulumi.Input[builtins.str]] = None) -> 'FunctionDomain':
        """
        Get an existing FunctionDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] function_id: The unique identifier of the function.
        :param pulumi.Input[builtins.str] hostname: The hostname with a CNAME record.
               
               We recommend you use a CNAME domain record that point to your native function `domain_name` for it.
               
               > **Important** Updating the `function_id` or `hostname` arguments will recreate the domain.
        :param pulumi.Input[builtins.str] region: (Defaults to provider `region`) The region in which the domain was created.
        :param pulumi.Input[builtins.str] url: The URL used to query the function.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FunctionDomainState.__new__(_FunctionDomainState)

        __props__.__dict__["function_id"] = function_id
        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["region"] = region
        __props__.__dict__["url"] = url
        return FunctionDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> pulumi.Output[builtins.str]:
        """
        The unique identifier of the function.
        """
        return pulumi.get(self, "function_id")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[builtins.str]:
        """
        The hostname with a CNAME record.

        We recommend you use a CNAME domain record that point to your native function `domain_name` for it.

        > **Important** Updating the `function_id` or `hostname` arguments will recreate the domain.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        (Defaults to provider `region`) The region in which the domain was created.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[builtins.str]:
        """
        The URL used to query the function.
        """
        return pulumi.get(self, "url")

