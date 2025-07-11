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
from . import outputs

__all__ = [
    'GetSecretResult',
    'AwaitableGetSecretResult',
    'get_secret',
    'get_secret_output',
]

@pulumi.output_type
class GetSecretResult:
    """
    A collection of values returned by getSecret.
    """
    def __init__(__self__, created_at=None, description=None, ephemeral_policies=None, id=None, name=None, organization_id=None, path=None, project_id=None, protected=None, region=None, secret_id=None, status=None, tags=None, type=None, updated_at=None, version_count=None, versions=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if ephemeral_policies and not isinstance(ephemeral_policies, list):
            raise TypeError("Expected argument 'ephemeral_policies' to be a list")
        pulumi.set(__self__, "ephemeral_policies", ephemeral_policies)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if protected and not isinstance(protected, bool):
            raise TypeError("Expected argument 'protected' to be a bool")
        pulumi.set(__self__, "protected", protected)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if secret_id and not isinstance(secret_id, str):
            raise TypeError("Expected argument 'secret_id' to be a str")
        pulumi.set(__self__, "secret_id", secret_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if version_count and not isinstance(version_count, int):
            raise TypeError("Expected argument 'version_count' to be a int")
        pulumi.set(__self__, "version_count", version_count)
        if versions and not isinstance(versions, list):
            raise TypeError("Expected argument 'versions' to be a list")
        pulumi.set(__self__, "versions", versions)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> builtins.str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ephemeralPolicies")
    def ephemeral_policies(self) -> Sequence['outputs.GetSecretEphemeralPolicyResult']:
        return pulumi.get(self, "ephemeral_policies")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> builtins.str:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter
    def path(self) -> Optional[builtins.str]:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def protected(self) -> builtins.bool:
        return pulumi.get(self, "protected")

    @property
    @pulumi.getter
    def region(self) -> Optional[builtins.str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> builtins.str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="versionCount")
    def version_count(self) -> builtins.int:
        return pulumi.get(self, "version_count")

    @property
    @pulumi.getter
    def versions(self) -> Sequence['outputs.GetSecretVersionResult']:
        return pulumi.get(self, "versions")


class AwaitableGetSecretResult(GetSecretResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecretResult(
            created_at=self.created_at,
            description=self.description,
            ephemeral_policies=self.ephemeral_policies,
            id=self.id,
            name=self.name,
            organization_id=self.organization_id,
            path=self.path,
            project_id=self.project_id,
            protected=self.protected,
            region=self.region,
            secret_id=self.secret_id,
            status=self.status,
            tags=self.tags,
            type=self.type,
            updated_at=self.updated_at,
            version_count=self.version_count,
            versions=self.versions)


def get_secret(name: Optional[builtins.str] = None,
               organization_id: Optional[builtins.str] = None,
               path: Optional[builtins.str] = None,
               project_id: Optional[builtins.str] = None,
               region: Optional[builtins.str] = None,
               secret_id: Optional[builtins.str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecretResult:
    """
    The `secrets.Secret` data source is used to get information about a specific secret in Scaleway's Secret Manager.

    Refer to the Secret Manager [product documentation](https://www.scaleway.com/en/docs/identity-and-access-management/secret-manager/) and [API documentation](https://www.scaleway.com/en/developers/api/secret-manager/) for more information.

    ## Example Usage

    ### Create a secret and get its information

    The following commands allow you to:

    - create a secret named `foo` with the description `barr`
    - retrieve the secret's information using the secret's ID
    - retrieve the secret's information using the secret's name

    ```python
    import pulumi
    import pulumi_scaleway as scaleway
    import pulumiverse_scaleway as scaleway

    # Create a secret
    main = scaleway.secrets.Secret("main",
        name="foo",
        description="barr")
    # Get the secret information specified by the secret ID
    my_secret = scaleway.secrets.get_secret(secret_id="11111111-1111-1111-1111-111111111111")
    # Get the secret information specified by the secret name
    by_name = scaleway.secrets.get_secret(name="your_secret_name")
    ```


    :param builtins.str name: The name of the secret.
           Only one of `name` and `secret_id` should be specified.
    :param builtins.str organization_id: The ID of the Scaleway Organization the Project is associated with. If no default `organization_id` is set, it must be set explicitly in this data source.
    :param builtins.str path: The path of the secret.
           Conflicts with `secret_id`.
    :param builtins.str project_id: ). The ID of the
           Project the secret is associated with.
    :param builtins.str region: ). The region in which the secret exists.
    :param builtins.str secret_id: The ID of the secret.
           Only one of `name` and `secret_id` should be specified.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['organizationId'] = organization_id
    __args__['path'] = path
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['secretId'] = secret_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:secrets/getSecret:getSecret', __args__, opts=opts, typ=GetSecretResult).value

    return AwaitableGetSecretResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        description=pulumi.get(__ret__, 'description'),
        ephemeral_policies=pulumi.get(__ret__, 'ephemeral_policies'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        path=pulumi.get(__ret__, 'path'),
        project_id=pulumi.get(__ret__, 'project_id'),
        protected=pulumi.get(__ret__, 'protected'),
        region=pulumi.get(__ret__, 'region'),
        secret_id=pulumi.get(__ret__, 'secret_id'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        version_count=pulumi.get(__ret__, 'version_count'),
        versions=pulumi.get(__ret__, 'versions'))
def get_secret_output(name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      organization_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      path: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      project_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      secret_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSecretResult]:
    """
    The `secrets.Secret` data source is used to get information about a specific secret in Scaleway's Secret Manager.

    Refer to the Secret Manager [product documentation](https://www.scaleway.com/en/docs/identity-and-access-management/secret-manager/) and [API documentation](https://www.scaleway.com/en/developers/api/secret-manager/) for more information.

    ## Example Usage

    ### Create a secret and get its information

    The following commands allow you to:

    - create a secret named `foo` with the description `barr`
    - retrieve the secret's information using the secret's ID
    - retrieve the secret's information using the secret's name

    ```python
    import pulumi
    import pulumi_scaleway as scaleway
    import pulumiverse_scaleway as scaleway

    # Create a secret
    main = scaleway.secrets.Secret("main",
        name="foo",
        description="barr")
    # Get the secret information specified by the secret ID
    my_secret = scaleway.secrets.get_secret(secret_id="11111111-1111-1111-1111-111111111111")
    # Get the secret information specified by the secret name
    by_name = scaleway.secrets.get_secret(name="your_secret_name")
    ```


    :param builtins.str name: The name of the secret.
           Only one of `name` and `secret_id` should be specified.
    :param builtins.str organization_id: The ID of the Scaleway Organization the Project is associated with. If no default `organization_id` is set, it must be set explicitly in this data source.
    :param builtins.str path: The path of the secret.
           Conflicts with `secret_id`.
    :param builtins.str project_id: ). The ID of the
           Project the secret is associated with.
    :param builtins.str region: ). The region in which the secret exists.
    :param builtins.str secret_id: The ID of the secret.
           Only one of `name` and `secret_id` should be specified.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['organizationId'] = organization_id
    __args__['path'] = path
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['secretId'] = secret_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:secrets/getSecret:getSecret', __args__, opts=opts, typ=GetSecretResult)
    return __ret__.apply(lambda __response__: GetSecretResult(
        created_at=pulumi.get(__response__, 'created_at'),
        description=pulumi.get(__response__, 'description'),
        ephemeral_policies=pulumi.get(__response__, 'ephemeral_policies'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        organization_id=pulumi.get(__response__, 'organization_id'),
        path=pulumi.get(__response__, 'path'),
        project_id=pulumi.get(__response__, 'project_id'),
        protected=pulumi.get(__response__, 'protected'),
        region=pulumi.get(__response__, 'region'),
        secret_id=pulumi.get(__response__, 'secret_id'),
        status=pulumi.get(__response__, 'status'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        updated_at=pulumi.get(__response__, 'updated_at'),
        version_count=pulumi.get(__response__, 'version_count'),
        versions=pulumi.get(__response__, 'versions')))
