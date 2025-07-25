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
    'SnsCredentialsPermissionsArgs',
    'SnsCredentialsPermissionsArgsDict',
    'SqsCredentialsPermissionsArgs',
    'SqsCredentialsPermissionsArgsDict',
]

MYPY = False

if not MYPY:
    class SnsCredentialsPermissionsArgsDict(TypedDict):
        can_manage: NotRequired[pulumi.Input[builtins.bool]]
        """
        . Defines whether the user can manage the associated resource(s).
        """
        can_publish: NotRequired[pulumi.Input[builtins.bool]]
        """
        . Defines whether the user can publish messages to the service.
        """
        can_receive: NotRequired[pulumi.Input[builtins.bool]]
        """
        . Defines whether the user can receive messages from the service.
        """
elif False:
    SnsCredentialsPermissionsArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SnsCredentialsPermissionsArgs:
    def __init__(__self__, *,
                 can_manage: Optional[pulumi.Input[builtins.bool]] = None,
                 can_publish: Optional[pulumi.Input[builtins.bool]] = None,
                 can_receive: Optional[pulumi.Input[builtins.bool]] = None):
        """
        :param pulumi.Input[builtins.bool] can_manage: . Defines whether the user can manage the associated resource(s).
        :param pulumi.Input[builtins.bool] can_publish: . Defines whether the user can publish messages to the service.
        :param pulumi.Input[builtins.bool] can_receive: . Defines whether the user can receive messages from the service.
        """
        if can_manage is not None:
            pulumi.set(__self__, "can_manage", can_manage)
        if can_publish is not None:
            pulumi.set(__self__, "can_publish", can_publish)
        if can_receive is not None:
            pulumi.set(__self__, "can_receive", can_receive)

    @property
    @pulumi.getter(name="canManage")
    def can_manage(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        . Defines whether the user can manage the associated resource(s).
        """
        return pulumi.get(self, "can_manage")

    @can_manage.setter
    def can_manage(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "can_manage", value)

    @property
    @pulumi.getter(name="canPublish")
    def can_publish(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        . Defines whether the user can publish messages to the service.
        """
        return pulumi.get(self, "can_publish")

    @can_publish.setter
    def can_publish(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "can_publish", value)

    @property
    @pulumi.getter(name="canReceive")
    def can_receive(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        . Defines whether the user can receive messages from the service.
        """
        return pulumi.get(self, "can_receive")

    @can_receive.setter
    def can_receive(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "can_receive", value)


if not MYPY:
    class SqsCredentialsPermissionsArgsDict(TypedDict):
        can_manage: NotRequired[pulumi.Input[builtins.bool]]
        """
        . Defines whether the user can manage the associated resource(s).
        """
        can_publish: NotRequired[pulumi.Input[builtins.bool]]
        """
        . Defines whether the user can publish messages to the service.
        """
        can_receive: NotRequired[pulumi.Input[builtins.bool]]
        """
        . Defines whether the user can receive messages from the service.
        """
elif False:
    SqsCredentialsPermissionsArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SqsCredentialsPermissionsArgs:
    def __init__(__self__, *,
                 can_manage: Optional[pulumi.Input[builtins.bool]] = None,
                 can_publish: Optional[pulumi.Input[builtins.bool]] = None,
                 can_receive: Optional[pulumi.Input[builtins.bool]] = None):
        """
        :param pulumi.Input[builtins.bool] can_manage: . Defines whether the user can manage the associated resource(s).
        :param pulumi.Input[builtins.bool] can_publish: . Defines whether the user can publish messages to the service.
        :param pulumi.Input[builtins.bool] can_receive: . Defines whether the user can receive messages from the service.
        """
        if can_manage is not None:
            pulumi.set(__self__, "can_manage", can_manage)
        if can_publish is not None:
            pulumi.set(__self__, "can_publish", can_publish)
        if can_receive is not None:
            pulumi.set(__self__, "can_receive", can_receive)

    @property
    @pulumi.getter(name="canManage")
    def can_manage(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        . Defines whether the user can manage the associated resource(s).
        """
        return pulumi.get(self, "can_manage")

    @can_manage.setter
    def can_manage(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "can_manage", value)

    @property
    @pulumi.getter(name="canPublish")
    def can_publish(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        . Defines whether the user can publish messages to the service.
        """
        return pulumi.get(self, "can_publish")

    @can_publish.setter
    def can_publish(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "can_publish", value)

    @property
    @pulumi.getter(name="canReceive")
    def can_receive(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        . Defines whether the user can receive messages from the service.
        """
        return pulumi.get(self, "can_receive")

    @can_receive.setter
    def can_receive(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "can_receive", value)


