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
    'SnapshotExportArgs',
    'SnapshotExportArgsDict',
    'SnapshotImportArgs',
    'SnapshotImportArgsDict',
]

MYPY = False

if not MYPY:
    class SnapshotExportArgsDict(TypedDict):
        bucket: pulumi.Input[builtins.str]
        """
        The name of the bucket where the QCOW file will be saved.
        """
        key: pulumi.Input[builtins.str]
        """
        The desired key (path) for the QCOW file within the bucket.
        """
elif False:
    SnapshotExportArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SnapshotExportArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[builtins.str],
                 key: pulumi.Input[builtins.str]):
        """
        :param pulumi.Input[builtins.str] bucket: The name of the bucket where the QCOW file will be saved.
        :param pulumi.Input[builtins.str] key: The desired key (path) for the QCOW file within the bucket.
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "key", key)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[builtins.str]:
        """
        The name of the bucket where the QCOW file will be saved.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[builtins.str]:
        """
        The desired key (path) for the QCOW file within the bucket.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "key", value)


if not MYPY:
    class SnapshotImportArgsDict(TypedDict):
        bucket: pulumi.Input[builtins.str]
        """
        The name of the bucket containing the QCOW file.
        """
        key: pulumi.Input[builtins.str]
        """
        The key of the QCOW file within the bucket.
        """
elif False:
    SnapshotImportArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SnapshotImportArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[builtins.str],
                 key: pulumi.Input[builtins.str]):
        """
        :param pulumi.Input[builtins.str] bucket: The name of the bucket containing the QCOW file.
        :param pulumi.Input[builtins.str] key: The key of the QCOW file within the bucket.
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "key", key)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[builtins.str]:
        """
        The name of the bucket containing the QCOW file.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[builtins.str]:
        """
        The key of the QCOW file within the bucket.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "key", value)


