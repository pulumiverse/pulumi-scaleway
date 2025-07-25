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
    'SecretEphemeralPolicy',
    'SecretVersion',
    'GetSecretEphemeralPolicyResult',
    'GetSecretVersionResult',
]

@pulumi.output_type
class SecretEphemeralPolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "expiresOnceAccessed":
            suggest = "expires_once_accessed"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SecretEphemeralPolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SecretEphemeralPolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SecretEphemeralPolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: builtins.str,
                 expires_once_accessed: Optional[builtins.bool] = None,
                 ttl: Optional[builtins.str] = None):
        """
        :param builtins.str action: Action to perform when the version of a secret expires. Available values can be found in [SDK constants](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/secret/v1beta1#pkg-constants).
        :param builtins.bool expires_once_accessed: True if the secret version expires after a single user access.
        :param builtins.str ttl: Time frame, from one second and up to one year, during which the secret's versions are valid. Has to be specified in [Go Duration format](https://pkg.go.dev/time#ParseDuration) (ex: "30m", "24h").
        """
        pulumi.set(__self__, "action", action)
        if expires_once_accessed is not None:
            pulumi.set(__self__, "expires_once_accessed", expires_once_accessed)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter
    def action(self) -> builtins.str:
        """
        Action to perform when the version of a secret expires. Available values can be found in [SDK constants](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/secret/v1beta1#pkg-constants).
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="expiresOnceAccessed")
    def expires_once_accessed(self) -> Optional[builtins.bool]:
        """
        True if the secret version expires after a single user access.
        """
        return pulumi.get(self, "expires_once_accessed")

    @property
    @pulumi.getter
    def ttl(self) -> Optional[builtins.str]:
        """
        Time frame, from one second and up to one year, during which the secret's versions are valid. Has to be specified in [Go Duration format](https://pkg.go.dev/time#ParseDuration) (ex: "30m", "24h").
        """
        return pulumi.get(self, "ttl")


@pulumi.output_type
class SecretVersion(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "secretId":
            suggest = "secret_id"
        elif key == "updatedAt":
            suggest = "updated_at"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SecretVersion. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SecretVersion.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SecretVersion.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: Optional[builtins.str] = None,
                 description: Optional[builtins.str] = None,
                 latest: Optional[builtins.bool] = None,
                 revision: Optional[builtins.str] = None,
                 secret_id: Optional[builtins.str] = None,
                 status: Optional[builtins.str] = None,
                 updated_at: Optional[builtins.str] = None):
        """
        :param builtins.str created_at: Date and time of the secret's creation (in RFC 3339 format).
        :param builtins.str description: Description of the secret (e.g. `my-new-description`).
        :param builtins.bool latest: Returns true if the version is the latest.
        :param builtins.str revision: The revision of secret version
        :param builtins.str secret_id: The secret ID associated with this version
        :param builtins.str status: The status of the secret.
        :param builtins.str updated_at: Date and time of the secret's last update (in RFC 3339 format).
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if latest is not None:
            pulumi.set(__self__, "latest", latest)
        if revision is not None:
            pulumi.set(__self__, "revision", revision)
        if secret_id is not None:
            pulumi.set(__self__, "secret_id", secret_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[builtins.str]:
        """
        Date and time of the secret's creation (in RFC 3339 format).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        Description of the secret (e.g. `my-new-description`).
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def latest(self) -> Optional[builtins.bool]:
        """
        Returns true if the version is the latest.
        """
        return pulumi.get(self, "latest")

    @property
    @pulumi.getter
    def revision(self) -> Optional[builtins.str]:
        """
        The revision of secret version
        """
        return pulumi.get(self, "revision")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[builtins.str]:
        """
        The secret ID associated with this version
        """
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter
    def status(self) -> Optional[builtins.str]:
        """
        The status of the secret.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[builtins.str]:
        """
        Date and time of the secret's last update (in RFC 3339 format).
        """
        return pulumi.get(self, "updated_at")


@pulumi.output_type
class GetSecretEphemeralPolicyResult(dict):
    def __init__(__self__, *,
                 action: builtins.str,
                 expires_once_accessed: builtins.bool,
                 ttl: builtins.str):
        """
        :param builtins.str action: Action to perform when the version of a secret expires.
        :param builtins.bool expires_once_accessed: True if the secret version expires after a single user access.
        :param builtins.str ttl: Time frame, from one second and up to one year, during which the secret's versions are valid. Has to be specified in Go Duration format
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "expires_once_accessed", expires_once_accessed)
        pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter
    def action(self) -> builtins.str:
        """
        Action to perform when the version of a secret expires.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="expiresOnceAccessed")
    def expires_once_accessed(self) -> builtins.bool:
        """
        True if the secret version expires after a single user access.
        """
        return pulumi.get(self, "expires_once_accessed")

    @property
    @pulumi.getter
    def ttl(self) -> builtins.str:
        """
        Time frame, from one second and up to one year, during which the secret's versions are valid. Has to be specified in Go Duration format
        """
        return pulumi.get(self, "ttl")


@pulumi.output_type
class GetSecretVersionResult(dict):
    def __init__(__self__, *,
                 created_at: builtins.str,
                 description: builtins.str,
                 latest: builtins.bool,
                 revision: builtins.str,
                 secret_id: builtins.str,
                 status: builtins.str,
                 updated_at: builtins.str):
        """
        :param builtins.str created_at: Date and time of secret version's creation (RFC 3339 format)
        :param builtins.str description: Description of the secret version
        :param builtins.bool latest: Returns true if the version is the latest.
        :param builtins.str revision: The revision of secret version
        :param builtins.str secret_id: The ID of the secret.
               Only one of `name` and `secret_id` should be specified.
        :param builtins.str status: Status of the secret version
        :param builtins.str updated_at: Date and time of secret version's creation (RFC 3339 format)
        """
        pulumi.set(__self__, "created_at", created_at)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "latest", latest)
        pulumi.set(__self__, "revision", revision)
        pulumi.set(__self__, "secret_id", secret_id)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> builtins.str:
        """
        Date and time of secret version's creation (RFC 3339 format)
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        Description of the secret version
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def latest(self) -> builtins.bool:
        """
        Returns true if the version is the latest.
        """
        return pulumi.get(self, "latest")

    @property
    @pulumi.getter
    def revision(self) -> builtins.str:
        """
        The revision of secret version
        """
        return pulumi.get(self, "revision")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> builtins.str:
        """
        The ID of the secret.
        Only one of `name` and `secret_id` should be specified.
        """
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        Status of the secret version
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> builtins.str:
        """
        Date and time of secret version's creation (RFC 3339 format)
        """
        return pulumi.get(self, "updated_at")


