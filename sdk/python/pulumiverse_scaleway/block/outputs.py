# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    'SnapshotImport',
    'GetSnapshotImportResult',
]

@pulumi.output_type
class SnapshotImport(dict):
    def __init__(__self__, *,
                 bucket: str,
                 key: str):
        """
        :param str bucket: Bucket containing qcow
        :param str key: Key of the qcow file in the specified bucket
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "key", key)

    @property
    @pulumi.getter
    def bucket(self) -> str:
        """
        Bucket containing qcow
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Key of the qcow file in the specified bucket
        """
        return pulumi.get(self, "key")


@pulumi.output_type
class GetSnapshotImportResult(dict):
    def __init__(__self__, *,
                 bucket: str,
                 key: str):
        """
        :param str bucket: Bucket containing qcow
        :param str key: Key of the qcow file in the specified bucket
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "key", key)

    @property
    @pulumi.getter
    def bucket(self) -> str:
        """
        Bucket containing qcow
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Key of the qcow file in the specified bucket
        """
        return pulumi.get(self, "key")


