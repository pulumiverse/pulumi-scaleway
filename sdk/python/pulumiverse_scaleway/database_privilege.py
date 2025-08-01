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

__all__ = ['DatabasePrivilegeArgs', 'DatabasePrivilege']

@pulumi.input_type
class DatabasePrivilegeArgs:
    def __init__(__self__, *,
                 database_name: pulumi.Input[builtins.str],
                 instance_id: pulumi.Input[builtins.str],
                 permission: pulumi.Input[builtins.str],
                 user_name: pulumi.Input[builtins.str],
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a DatabasePrivilege resource.
        :param pulumi.Input[builtins.str] database_name: Name of the database (e.g. `my-db-name`).
        :param pulumi.Input[builtins.str] instance_id: UUID of the Database Instance.
        :param pulumi.Input[builtins.str] permission: Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        :param pulumi.Input[builtins.str] user_name: Name of the user (e.g. `my-db-user`).
        :param pulumi.Input[builtins.str] region: `region`) The region in which the resource exists.
        """
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "permission", permission)
        pulumi.set(__self__, "user_name", user_name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the database (e.g. `my-db-name`).
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[builtins.str]:
        """
        UUID of the Database Instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def permission(self) -> pulumi.Input[builtins.str]:
        """
        Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        """
        return pulumi.get(self, "permission")

    @permission.setter
    def permission(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "permission", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the user (e.g. `my-db-user`).
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        `region`) The region in which the resource exists.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _DatabasePrivilegeState:
    def __init__(__self__, *,
                 database_name: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 permission: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 user_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering DatabasePrivilege resources.
        :param pulumi.Input[builtins.str] database_name: Name of the database (e.g. `my-db-name`).
        :param pulumi.Input[builtins.str] instance_id: UUID of the Database Instance.
        :param pulumi.Input[builtins.str] permission: Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        :param pulumi.Input[builtins.str] region: `region`) The region in which the resource exists.
        :param pulumi.Input[builtins.str] user_name: Name of the user (e.g. `my-db-user`).
        """
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if permission is not None:
            pulumi.set(__self__, "permission", permission)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the database (e.g. `my-db-name`).
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        UUID of the Database Instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def permission(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        """
        return pulumi.get(self, "permission")

    @permission.setter
    def permission(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "permission", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        `region`) The region in which the resource exists.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the user (e.g. `my-db-user`).
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "user_name", value)


warnings.warn("""scaleway.index/databaseprivilege.DatabasePrivilege has been deprecated in favor of scaleway.databases/privilege.Privilege""", DeprecationWarning)


@pulumi.type_token("scaleway:index/databasePrivilege:DatabasePrivilege")
class DatabasePrivilege(pulumi.CustomResource):
    warnings.warn("""scaleway.index/databaseprivilege.DatabasePrivilege has been deprecated in favor of scaleway.databases/privilege.Privilege""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 permission: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 user_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create and manage Scaleway database privileges.
        For more information refer to the [API documentation](https://www.scaleway.com/en/developers/api/managed-database-postgre-mysql/#user-and-permissions).

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.databases.Instance("main",
            name="rdb",
            node_type="DB-DEV-S",
            engine="PostgreSQL-11",
            is_ha_cluster=True,
            disable_backup=True,
            user_name="my_initial_user",
            password="thiZ_is_v&ry_s3cret")
        main_database = scaleway.databases.Database("main",
            instance_id=main.id,
            name="database")
        main_user = scaleway.databases.User("main",
            instance_id=main.id,
            name="my-db-user",
            password="thiZ_is_v&ry_s3cret",
            is_admin=False)
        main_privilege = scaleway.databases.Privilege("main",
            instance_id=main.id,
            user_name=main_user.name,
            database_name=main_database.name,
            permission="all")
        ```

        ## Import

        The user privileges can be imported using the `{region}/{instance_id}/{database_name}/{user_name}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/databasePrivilege:DatabasePrivilege o fr-par/11111111-1111-1111-1111-111111111111/database_name/foo
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] database_name: Name of the database (e.g. `my-db-name`).
        :param pulumi.Input[builtins.str] instance_id: UUID of the Database Instance.
        :param pulumi.Input[builtins.str] permission: Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        :param pulumi.Input[builtins.str] region: `region`) The region in which the resource exists.
        :param pulumi.Input[builtins.str] user_name: Name of the user (e.g. `my-db-user`).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabasePrivilegeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create and manage Scaleway database privileges.
        For more information refer to the [API documentation](https://www.scaleway.com/en/developers/api/managed-database-postgre-mysql/#user-and-permissions).

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.databases.Instance("main",
            name="rdb",
            node_type="DB-DEV-S",
            engine="PostgreSQL-11",
            is_ha_cluster=True,
            disable_backup=True,
            user_name="my_initial_user",
            password="thiZ_is_v&ry_s3cret")
        main_database = scaleway.databases.Database("main",
            instance_id=main.id,
            name="database")
        main_user = scaleway.databases.User("main",
            instance_id=main.id,
            name="my-db-user",
            password="thiZ_is_v&ry_s3cret",
            is_admin=False)
        main_privilege = scaleway.databases.Privilege("main",
            instance_id=main.id,
            user_name=main_user.name,
            database_name=main_database.name,
            permission="all")
        ```

        ## Import

        The user privileges can be imported using the `{region}/{instance_id}/{database_name}/{user_name}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/databasePrivilege:DatabasePrivilege o fr-par/11111111-1111-1111-1111-111111111111/database_name/foo
        ```

        :param str resource_name: The name of the resource.
        :param DatabasePrivilegeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabasePrivilegeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 permission: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 user_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        pulumi.log.warn("""DatabasePrivilege is deprecated: scaleway.index/databaseprivilege.DatabasePrivilege has been deprecated in favor of scaleway.databases/privilege.Privilege""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabasePrivilegeArgs.__new__(DatabasePrivilegeArgs)

            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__.__dict__["database_name"] = database_name
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if permission is None and not opts.urn:
                raise TypeError("Missing required property 'permission'")
            __props__.__dict__["permission"] = permission
            __props__.__dict__["region"] = region
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
        super(DatabasePrivilege, __self__).__init__(
            'scaleway:index/databasePrivilege:DatabasePrivilege',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            database_name: Optional[pulumi.Input[builtins.str]] = None,
            instance_id: Optional[pulumi.Input[builtins.str]] = None,
            permission: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            user_name: Optional[pulumi.Input[builtins.str]] = None) -> 'DatabasePrivilege':
        """
        Get an existing DatabasePrivilege resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] database_name: Name of the database (e.g. `my-db-name`).
        :param pulumi.Input[builtins.str] instance_id: UUID of the Database Instance.
        :param pulumi.Input[builtins.str] permission: Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        :param pulumi.Input[builtins.str] region: `region`) The region in which the resource exists.
        :param pulumi.Input[builtins.str] user_name: Name of the user (e.g. `my-db-user`).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatabasePrivilegeState.__new__(_DatabasePrivilegeState)

        __props__.__dict__["database_name"] = database_name
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["permission"] = permission
        __props__.__dict__["region"] = region
        __props__.__dict__["user_name"] = user_name
        return DatabasePrivilege(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the database (e.g. `my-db-name`).
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[builtins.str]:
        """
        UUID of the Database Instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def permission(self) -> pulumi.Output[builtins.str]:
        """
        Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        """
        return pulumi.get(self, "permission")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        `region`) The region in which the resource exists.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the user (e.g. `my-db-user`).
        """
        return pulumi.get(self, "user_name")

