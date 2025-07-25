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

__all__ = ['BucketPolicyArgs', 'BucketPolicy']

@pulumi.input_type
class BucketPolicyArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[builtins.str],
                 policy: pulumi.Input[builtins.str],
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a BucketPolicy resource.
        :param pulumi.Input[builtins.str] bucket: The bucket's name or regional ID.
        :param pulumi.Input[builtins.str] policy: The text of the policy.
        :param pulumi.Input[builtins.str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[builtins.str] region: The Scaleway region this bucket resides in.
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "policy", policy)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[builtins.str]:
        """
        The bucket's name or regional ID.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Input[builtins.str]:
        """
        The text of the policy.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Scaleway region this bucket resides in.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _BucketPolicyState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[builtins.str]] = None,
                 policy: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering BucketPolicy resources.
        :param pulumi.Input[builtins.str] bucket: The bucket's name or regional ID.
        :param pulumi.Input[builtins.str] policy: The text of the policy.
        :param pulumi.Input[builtins.str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[builtins.str] region: The Scaleway region this bucket resides in.
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The bucket's name or regional ID.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The text of the policy.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Scaleway region this bucket resides in.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.type_token("scaleway:object/bucketPolicy:BucketPolicy")
class BucketPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[builtins.str]] = None,
                 policy: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The `object.BucketPolicy` resource allows you to create and manage bucket policies for [Scaleway Object storage](https://www.scaleway.com/en/docs/object-storage/).

        Refer to the [dedicated documentation](https://www.scaleway.com/en/docs/object-storage/api-cli/bucket-policy/) for more information on Object Storage bucket policies.

        ## Example Usage

        ### Example Usage with an IAM user

        ```python
        import pulumi
        import json
        import pulumi_scaleway as scaleway
        import pulumiverse_scaleway as scaleway

        # Project ID
        default = scaleway.account.get_project(name="default")
        # IAM configuration
        user = scaleway.iam.get_user(email="user@scaleway.com")
        policy = scaleway.iam.Policy("policy",
            name="object-storage-policy",
            user_id=user.id,
            rules=[{
                "project_ids": [default.id],
                "permission_set_names": ["ObjectStorageFullAccess"],
            }])
        # Object storage configuration
        bucket = scaleway.object.Bucket("bucket", name="some-unique-name")
        policy_bucket_policy = scaleway.object.BucketPolicy("policy",
            bucket=bucket.name,
            policy=pulumi.Output.json_dumps({
                "Version": "2023-04-17",
                "Id": "MyBucketPolicy",
                "Statement": [{
                    "Effect": "Allow",
                    "Action": ["s3:*"],
                    "Principal": {
                        "SCW": f"user_id:{user.id}",
                    },
                    "Resource": [
                        bucket.name,
                        bucket.name.apply(lambda name: f"{name}/*"),
                    ],
                }],
            }))
        ```

        ### Example with an IAM application

        ### Creating a bucket and delegating read access to an application

        ```python
        import pulumi
        import json
        import pulumi_scaleway as scaleway
        import pulumiverse_scaleway as scaleway

        # Project ID
        default = scaleway.account.get_project(name="default")
        # IAM configuration
        reading_app = scaleway.iam.Application("reading-app", name="reading-app")
        policy = scaleway.iam.Policy("policy",
            name="object-storage-policy",
            application_id=reading_app.id,
            rules=[{
                "project_ids": [default.id],
                "permission_set_names": ["ObjectStorageBucketsRead"],
            }])
        # Object storage configuration
        bucket = scaleway.object.Bucket("bucket", name="some-unique-name")
        policy_bucket_policy = scaleway.object.BucketPolicy("policy",
            bucket=bucket.id,
            policy=pulumi.Output.json_dumps({
                "Version": "2023-04-17",
                "Statement": [{
                    "Sid": "Delegate read access",
                    "Effect": "Allow",
                    "Principal": {
                        "SCW": reading_app.id.apply(lambda id: f"application_id:{id}"),
                    },
                    "Action": [
                        "s3:ListBucket",
                        "s3:GetObject",
                    ],
                    "Resource": [
                        bucket.name,
                        bucket.name.apply(lambda name: f"{name}/*"),
                    ],
                }],
            }))
        ```

        ### Reading the bucket with the application

        ```python
        import pulumi
        import pulumi_scaleway as scaleway
        import pulumiverse_scaleway as scaleway

        reading_app = scaleway.iam.get_application(name="reading-app")
        reading_api_key = scaleway.iam.ApiKey("reading-api-key", application_id=reading_app.id)
        bucket = scaleway.object.get_bucket(name="some-unique-name")
        ```

        ### Example with AWS provider

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_scaleway as scaleway
        import pulumiverse_scaleway as scaleway

        # Scaleway project ID
        default = scaleway.account.get_project(name="default")
        # Object storage configuration
        bucket = scaleway.object.Bucket("bucket", name="some-unique-name")
        # AWS data source
        policy = aws.iam.get_policy_document_output(version="2012-10-17",
            statements=[{
                "sid": "Delegate access",
                "effect": "Allow",
                "principals": [{
                    "type": "SCW",
                    "identifiers": [f"project_id:{default.id}"],
                }],
                "actions": ["s3:ListBucket"],
                "resources": [
                    bucket.name,
                    bucket.name.apply(lambda name: f"{name}/*"),
                ],
            }])
        main = scaleway.object.BucketPolicy("main",
            bucket=bucket.id,
            policy=policy.json)
        ```

        ### Example with deprecated version 2012-10-17

        ```python
        import pulumi
        import json
        import pulumi_scaleway as scaleway
        import pulumiverse_scaleway as scaleway

        # Project ID
        default = scaleway.account.get_project(name="default")
        # Object storage configuration
        bucket = scaleway.object.Bucket("bucket",
            name="mia-cross-crash-tests",
            region="fr-par")
        policy = scaleway.object.BucketPolicy("policy",
            bucket=bucket.name,
            policy=pulumi.Output.json_dumps({
                "Version": "2012-10-17",
                "Statement": [{
                    "Effect": "Allow",
                    "Action": [
                        "s3:ListBucket",
                        "s3:GetObjectTagging",
                    ],
                    "Principal": {
                        "SCW": f"project_id:{default.id}",
                    },
                    "Resource": [
                        bucket.name,
                        bucket.name.apply(lambda name: f"{name}/*"),
                    ],
                }],
            }))
        ```

        **NB:** To configure the AWS provider with Scaleway credentials, refer to the [dedicated documentation](https://www.scaleway.com/en/docs/object-storage/api-cli/object-storage-aws-cli/).

        ## Import

        Bucket policies can be imported using the `{region}/{bucketName}` identifier, as shown below:

        bash

        ```sh
        $ pulumi import scaleway:object/bucketPolicy:BucketPolicy some_bucket fr-par/some-bucket
        ```

        ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project.

        If you are using a project different from the default one, you have to specify the project ID at the end of the import command.

        bash

        ```sh
        $ pulumi import scaleway:object/bucketPolicy:BucketPolicy some_bucket fr-par/some-bucket@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] bucket: The bucket's name or regional ID.
        :param pulumi.Input[builtins.str] policy: The text of the policy.
        :param pulumi.Input[builtins.str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[builtins.str] region: The Scaleway region this bucket resides in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BucketPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `object.BucketPolicy` resource allows you to create and manage bucket policies for [Scaleway Object storage](https://www.scaleway.com/en/docs/object-storage/).

        Refer to the [dedicated documentation](https://www.scaleway.com/en/docs/object-storage/api-cli/bucket-policy/) for more information on Object Storage bucket policies.

        ## Example Usage

        ### Example Usage with an IAM user

        ```python
        import pulumi
        import json
        import pulumi_scaleway as scaleway
        import pulumiverse_scaleway as scaleway

        # Project ID
        default = scaleway.account.get_project(name="default")
        # IAM configuration
        user = scaleway.iam.get_user(email="user@scaleway.com")
        policy = scaleway.iam.Policy("policy",
            name="object-storage-policy",
            user_id=user.id,
            rules=[{
                "project_ids": [default.id],
                "permission_set_names": ["ObjectStorageFullAccess"],
            }])
        # Object storage configuration
        bucket = scaleway.object.Bucket("bucket", name="some-unique-name")
        policy_bucket_policy = scaleway.object.BucketPolicy("policy",
            bucket=bucket.name,
            policy=pulumi.Output.json_dumps({
                "Version": "2023-04-17",
                "Id": "MyBucketPolicy",
                "Statement": [{
                    "Effect": "Allow",
                    "Action": ["s3:*"],
                    "Principal": {
                        "SCW": f"user_id:{user.id}",
                    },
                    "Resource": [
                        bucket.name,
                        bucket.name.apply(lambda name: f"{name}/*"),
                    ],
                }],
            }))
        ```

        ### Example with an IAM application

        ### Creating a bucket and delegating read access to an application

        ```python
        import pulumi
        import json
        import pulumi_scaleway as scaleway
        import pulumiverse_scaleway as scaleway

        # Project ID
        default = scaleway.account.get_project(name="default")
        # IAM configuration
        reading_app = scaleway.iam.Application("reading-app", name="reading-app")
        policy = scaleway.iam.Policy("policy",
            name="object-storage-policy",
            application_id=reading_app.id,
            rules=[{
                "project_ids": [default.id],
                "permission_set_names": ["ObjectStorageBucketsRead"],
            }])
        # Object storage configuration
        bucket = scaleway.object.Bucket("bucket", name="some-unique-name")
        policy_bucket_policy = scaleway.object.BucketPolicy("policy",
            bucket=bucket.id,
            policy=pulumi.Output.json_dumps({
                "Version": "2023-04-17",
                "Statement": [{
                    "Sid": "Delegate read access",
                    "Effect": "Allow",
                    "Principal": {
                        "SCW": reading_app.id.apply(lambda id: f"application_id:{id}"),
                    },
                    "Action": [
                        "s3:ListBucket",
                        "s3:GetObject",
                    ],
                    "Resource": [
                        bucket.name,
                        bucket.name.apply(lambda name: f"{name}/*"),
                    ],
                }],
            }))
        ```

        ### Reading the bucket with the application

        ```python
        import pulumi
        import pulumi_scaleway as scaleway
        import pulumiverse_scaleway as scaleway

        reading_app = scaleway.iam.get_application(name="reading-app")
        reading_api_key = scaleway.iam.ApiKey("reading-api-key", application_id=reading_app.id)
        bucket = scaleway.object.get_bucket(name="some-unique-name")
        ```

        ### Example with AWS provider

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_scaleway as scaleway
        import pulumiverse_scaleway as scaleway

        # Scaleway project ID
        default = scaleway.account.get_project(name="default")
        # Object storage configuration
        bucket = scaleway.object.Bucket("bucket", name="some-unique-name")
        # AWS data source
        policy = aws.iam.get_policy_document_output(version="2012-10-17",
            statements=[{
                "sid": "Delegate access",
                "effect": "Allow",
                "principals": [{
                    "type": "SCW",
                    "identifiers": [f"project_id:{default.id}"],
                }],
                "actions": ["s3:ListBucket"],
                "resources": [
                    bucket.name,
                    bucket.name.apply(lambda name: f"{name}/*"),
                ],
            }])
        main = scaleway.object.BucketPolicy("main",
            bucket=bucket.id,
            policy=policy.json)
        ```

        ### Example with deprecated version 2012-10-17

        ```python
        import pulumi
        import json
        import pulumi_scaleway as scaleway
        import pulumiverse_scaleway as scaleway

        # Project ID
        default = scaleway.account.get_project(name="default")
        # Object storage configuration
        bucket = scaleway.object.Bucket("bucket",
            name="mia-cross-crash-tests",
            region="fr-par")
        policy = scaleway.object.BucketPolicy("policy",
            bucket=bucket.name,
            policy=pulumi.Output.json_dumps({
                "Version": "2012-10-17",
                "Statement": [{
                    "Effect": "Allow",
                    "Action": [
                        "s3:ListBucket",
                        "s3:GetObjectTagging",
                    ],
                    "Principal": {
                        "SCW": f"project_id:{default.id}",
                    },
                    "Resource": [
                        bucket.name,
                        bucket.name.apply(lambda name: f"{name}/*"),
                    ],
                }],
            }))
        ```

        **NB:** To configure the AWS provider with Scaleway credentials, refer to the [dedicated documentation](https://www.scaleway.com/en/docs/object-storage/api-cli/object-storage-aws-cli/).

        ## Import

        Bucket policies can be imported using the `{region}/{bucketName}` identifier, as shown below:

        bash

        ```sh
        $ pulumi import scaleway:object/bucketPolicy:BucketPolicy some_bucket fr-par/some-bucket
        ```

        ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project.

        If you are using a project different from the default one, you have to specify the project ID at the end of the import command.

        bash

        ```sh
        $ pulumi import scaleway:object/bucketPolicy:BucketPolicy some_bucket fr-par/some-bucket@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param BucketPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BucketPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[builtins.str]] = None,
                 policy: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BucketPolicyArgs.__new__(BucketPolicyArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__.__dict__["policy"] = policy
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="scaleway:index/objectBucketPolicy:ObjectBucketPolicy")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(BucketPolicy, __self__).__init__(
            'scaleway:object/bucketPolicy:BucketPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[builtins.str]] = None,
            policy: Optional[pulumi.Input[builtins.str]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None) -> 'BucketPolicy':
        """
        Get an existing BucketPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] bucket: The bucket's name or regional ID.
        :param pulumi.Input[builtins.str] policy: The text of the policy.
        :param pulumi.Input[builtins.str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[builtins.str] region: The Scaleway region this bucket resides in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BucketPolicyState.__new__(_BucketPolicyState)

        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["policy"] = policy
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        return BucketPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[builtins.str]:
        """
        The bucket's name or regional ID.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[builtins.str]:
        """
        The text of the policy.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        The Scaleway region this bucket resides in.
        """
        return pulumi.get(self, "region")

