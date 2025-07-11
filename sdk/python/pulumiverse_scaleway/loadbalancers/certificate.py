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
from ._inputs import *

__all__ = ['CertificateArgs', 'Certificate']

@pulumi.input_type
class CertificateArgs:
    def __init__(__self__, *,
                 lb_id: pulumi.Input[builtins.str],
                 custom_certificate: Optional[pulumi.Input['CertificateCustomCertificateArgs']] = None,
                 letsencrypt: Optional[pulumi.Input['CertificateLetsencryptArgs']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Certificate resource.
        :param pulumi.Input[builtins.str] lb_id: The load-balancer ID
        :param pulumi.Input['CertificateCustomCertificateArgs'] custom_certificate: The custom type certificate type configuration
        :param pulumi.Input['CertificateLetsencryptArgs'] letsencrypt: The Let's Encrypt type certificate configuration
        :param pulumi.Input[builtins.str] name: The name of the load-balancer certificate
        """
        pulumi.set(__self__, "lb_id", lb_id)
        if custom_certificate is not None:
            pulumi.set(__self__, "custom_certificate", custom_certificate)
        if letsencrypt is not None:
            pulumi.set(__self__, "letsencrypt", letsencrypt)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> pulumi.Input[builtins.str]:
        """
        The load-balancer ID
        """
        return pulumi.get(self, "lb_id")

    @lb_id.setter
    def lb_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "lb_id", value)

    @property
    @pulumi.getter(name="customCertificate")
    def custom_certificate(self) -> Optional[pulumi.Input['CertificateCustomCertificateArgs']]:
        """
        The custom type certificate type configuration
        """
        return pulumi.get(self, "custom_certificate")

    @custom_certificate.setter
    def custom_certificate(self, value: Optional[pulumi.Input['CertificateCustomCertificateArgs']]):
        pulumi.set(self, "custom_certificate", value)

    @property
    @pulumi.getter
    def letsencrypt(self) -> Optional[pulumi.Input['CertificateLetsencryptArgs']]:
        """
        The Let's Encrypt type certificate configuration
        """
        return pulumi.get(self, "letsencrypt")

    @letsencrypt.setter
    def letsencrypt(self, value: Optional[pulumi.Input['CertificateLetsencryptArgs']]):
        pulumi.set(self, "letsencrypt", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the load-balancer certificate
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _CertificateState:
    def __init__(__self__, *,
                 common_name: Optional[pulumi.Input[builtins.str]] = None,
                 custom_certificate: Optional[pulumi.Input['CertificateCustomCertificateArgs']] = None,
                 fingerprint: Optional[pulumi.Input[builtins.str]] = None,
                 lb_id: Optional[pulumi.Input[builtins.str]] = None,
                 letsencrypt: Optional[pulumi.Input['CertificateLetsencryptArgs']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 not_valid_after: Optional[pulumi.Input[builtins.str]] = None,
                 not_valid_before: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 subject_alternative_names: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering Certificate resources.
        :param pulumi.Input[builtins.str] common_name: Main domain of the certificate
        :param pulumi.Input['CertificateCustomCertificateArgs'] custom_certificate: The custom type certificate type configuration
        :param pulumi.Input[builtins.str] fingerprint: The identifier (SHA-1) of the certificate
        :param pulumi.Input[builtins.str] lb_id: The load-balancer ID
        :param pulumi.Input['CertificateLetsencryptArgs'] letsencrypt: The Let's Encrypt type certificate configuration
        :param pulumi.Input[builtins.str] name: The name of the load-balancer certificate
        :param pulumi.Input[builtins.str] not_valid_after: The not valid after validity bound timestamp
        :param pulumi.Input[builtins.str] not_valid_before: The not valid before validity bound timestamp
        :param pulumi.Input[builtins.str] status: Certificate status
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] subject_alternative_names: The alternative domain names of the certificate
        """
        if common_name is not None:
            pulumi.set(__self__, "common_name", common_name)
        if custom_certificate is not None:
            pulumi.set(__self__, "custom_certificate", custom_certificate)
        if fingerprint is not None:
            pulumi.set(__self__, "fingerprint", fingerprint)
        if lb_id is not None:
            pulumi.set(__self__, "lb_id", lb_id)
        if letsencrypt is not None:
            pulumi.set(__self__, "letsencrypt", letsencrypt)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if not_valid_after is not None:
            pulumi.set(__self__, "not_valid_after", not_valid_after)
        if not_valid_before is not None:
            pulumi.set(__self__, "not_valid_before", not_valid_before)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if subject_alternative_names is not None:
            pulumi.set(__self__, "subject_alternative_names", subject_alternative_names)

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Main domain of the certificate
        """
        return pulumi.get(self, "common_name")

    @common_name.setter
    def common_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "common_name", value)

    @property
    @pulumi.getter(name="customCertificate")
    def custom_certificate(self) -> Optional[pulumi.Input['CertificateCustomCertificateArgs']]:
        """
        The custom type certificate type configuration
        """
        return pulumi.get(self, "custom_certificate")

    @custom_certificate.setter
    def custom_certificate(self, value: Optional[pulumi.Input['CertificateCustomCertificateArgs']]):
        pulumi.set(self, "custom_certificate", value)

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The identifier (SHA-1) of the certificate
        """
        return pulumi.get(self, "fingerprint")

    @fingerprint.setter
    def fingerprint(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "fingerprint", value)

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The load-balancer ID
        """
        return pulumi.get(self, "lb_id")

    @lb_id.setter
    def lb_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "lb_id", value)

    @property
    @pulumi.getter
    def letsencrypt(self) -> Optional[pulumi.Input['CertificateLetsencryptArgs']]:
        """
        The Let's Encrypt type certificate configuration
        """
        return pulumi.get(self, "letsencrypt")

    @letsencrypt.setter
    def letsencrypt(self, value: Optional[pulumi.Input['CertificateLetsencryptArgs']]):
        pulumi.set(self, "letsencrypt", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the load-balancer certificate
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notValidAfter")
    def not_valid_after(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The not valid after validity bound timestamp
        """
        return pulumi.get(self, "not_valid_after")

    @not_valid_after.setter
    def not_valid_after(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "not_valid_after", value)

    @property
    @pulumi.getter(name="notValidBefore")
    def not_valid_before(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The not valid before validity bound timestamp
        """
        return pulumi.get(self, "not_valid_before")

    @not_valid_before.setter
    def not_valid_before(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "not_valid_before", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Certificate status
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="subjectAlternativeNames")
    def subject_alternative_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The alternative domain names of the certificate
        """
        return pulumi.get(self, "subject_alternative_names")

    @subject_alternative_names.setter
    def subject_alternative_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "subject_alternative_names", value)


@pulumi.type_token("scaleway:loadbalancers/certificate:Certificate")
class Certificate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_certificate: Optional[pulumi.Input[Union['CertificateCustomCertificateArgs', 'CertificateCustomCertificateArgsDict']]] = None,
                 lb_id: Optional[pulumi.Input[builtins.str]] = None,
                 letsencrypt: Optional[pulumi.Input[Union['CertificateLetsencryptArgs', 'CertificateLetsencryptArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a Certificate resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['CertificateCustomCertificateArgs', 'CertificateCustomCertificateArgsDict']] custom_certificate: The custom type certificate type configuration
        :param pulumi.Input[builtins.str] lb_id: The load-balancer ID
        :param pulumi.Input[Union['CertificateLetsencryptArgs', 'CertificateLetsencryptArgsDict']] letsencrypt: The Let's Encrypt type certificate configuration
        :param pulumi.Input[builtins.str] name: The name of the load-balancer certificate
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CertificateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Certificate resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CertificateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CertificateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_certificate: Optional[pulumi.Input[Union['CertificateCustomCertificateArgs', 'CertificateCustomCertificateArgsDict']]] = None,
                 lb_id: Optional[pulumi.Input[builtins.str]] = None,
                 letsencrypt: Optional[pulumi.Input[Union['CertificateLetsencryptArgs', 'CertificateLetsencryptArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CertificateArgs.__new__(CertificateArgs)

            __props__.__dict__["custom_certificate"] = custom_certificate
            if lb_id is None and not opts.urn:
                raise TypeError("Missing required property 'lb_id'")
            __props__.__dict__["lb_id"] = lb_id
            __props__.__dict__["letsencrypt"] = letsencrypt
            __props__.__dict__["name"] = name
            __props__.__dict__["common_name"] = None
            __props__.__dict__["fingerprint"] = None
            __props__.__dict__["not_valid_after"] = None
            __props__.__dict__["not_valid_before"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["subject_alternative_names"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="scaleway:index/loadbalancerCertificate:LoadbalancerCertificate")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Certificate, __self__).__init__(
            'scaleway:loadbalancers/certificate:Certificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            common_name: Optional[pulumi.Input[builtins.str]] = None,
            custom_certificate: Optional[pulumi.Input[Union['CertificateCustomCertificateArgs', 'CertificateCustomCertificateArgsDict']]] = None,
            fingerprint: Optional[pulumi.Input[builtins.str]] = None,
            lb_id: Optional[pulumi.Input[builtins.str]] = None,
            letsencrypt: Optional[pulumi.Input[Union['CertificateLetsencryptArgs', 'CertificateLetsencryptArgsDict']]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            not_valid_after: Optional[pulumi.Input[builtins.str]] = None,
            not_valid_before: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None,
            subject_alternative_names: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None) -> 'Certificate':
        """
        Get an existing Certificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] common_name: Main domain of the certificate
        :param pulumi.Input[Union['CertificateCustomCertificateArgs', 'CertificateCustomCertificateArgsDict']] custom_certificate: The custom type certificate type configuration
        :param pulumi.Input[builtins.str] fingerprint: The identifier (SHA-1) of the certificate
        :param pulumi.Input[builtins.str] lb_id: The load-balancer ID
        :param pulumi.Input[Union['CertificateLetsencryptArgs', 'CertificateLetsencryptArgsDict']] letsencrypt: The Let's Encrypt type certificate configuration
        :param pulumi.Input[builtins.str] name: The name of the load-balancer certificate
        :param pulumi.Input[builtins.str] not_valid_after: The not valid after validity bound timestamp
        :param pulumi.Input[builtins.str] not_valid_before: The not valid before validity bound timestamp
        :param pulumi.Input[builtins.str] status: Certificate status
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] subject_alternative_names: The alternative domain names of the certificate
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CertificateState.__new__(_CertificateState)

        __props__.__dict__["common_name"] = common_name
        __props__.__dict__["custom_certificate"] = custom_certificate
        __props__.__dict__["fingerprint"] = fingerprint
        __props__.__dict__["lb_id"] = lb_id
        __props__.__dict__["letsencrypt"] = letsencrypt
        __props__.__dict__["name"] = name
        __props__.__dict__["not_valid_after"] = not_valid_after
        __props__.__dict__["not_valid_before"] = not_valid_before
        __props__.__dict__["status"] = status
        __props__.__dict__["subject_alternative_names"] = subject_alternative_names
        return Certificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> pulumi.Output[builtins.str]:
        """
        Main domain of the certificate
        """
        return pulumi.get(self, "common_name")

    @property
    @pulumi.getter(name="customCertificate")
    def custom_certificate(self) -> pulumi.Output[Optional['outputs.CertificateCustomCertificate']]:
        """
        The custom type certificate type configuration
        """
        return pulumi.get(self, "custom_certificate")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[builtins.str]:
        """
        The identifier (SHA-1) of the certificate
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> pulumi.Output[builtins.str]:
        """
        The load-balancer ID
        """
        return pulumi.get(self, "lb_id")

    @property
    @pulumi.getter
    def letsencrypt(self) -> pulumi.Output[Optional['outputs.CertificateLetsencrypt']]:
        """
        The Let's Encrypt type certificate configuration
        """
        return pulumi.get(self, "letsencrypt")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the load-balancer certificate
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notValidAfter")
    def not_valid_after(self) -> pulumi.Output[builtins.str]:
        """
        The not valid after validity bound timestamp
        """
        return pulumi.get(self, "not_valid_after")

    @property
    @pulumi.getter(name="notValidBefore")
    def not_valid_before(self) -> pulumi.Output[builtins.str]:
        """
        The not valid before validity bound timestamp
        """
        return pulumi.get(self, "not_valid_before")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        Certificate status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subjectAlternativeNames")
    def subject_alternative_names(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        The alternative domain names of the certificate
        """
        return pulumi.get(self, "subject_alternative_names")

