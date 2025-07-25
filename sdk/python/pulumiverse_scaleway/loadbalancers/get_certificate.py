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
    'GetCertificateResult',
    'AwaitableGetCertificateResult',
    'get_certificate',
    'get_certificate_output',
]

@pulumi.output_type
class GetCertificateResult:
    """
    A collection of values returned by getCertificate.
    """
    def __init__(__self__, certificate_id=None, common_name=None, custom_certificates=None, fingerprint=None, id=None, lb_id=None, letsencrypts=None, name=None, not_valid_after=None, not_valid_before=None, status=None, subject_alternative_names=None):
        if certificate_id and not isinstance(certificate_id, str):
            raise TypeError("Expected argument 'certificate_id' to be a str")
        pulumi.set(__self__, "certificate_id", certificate_id)
        if common_name and not isinstance(common_name, str):
            raise TypeError("Expected argument 'common_name' to be a str")
        pulumi.set(__self__, "common_name", common_name)
        if custom_certificates and not isinstance(custom_certificates, list):
            raise TypeError("Expected argument 'custom_certificates' to be a list")
        pulumi.set(__self__, "custom_certificates", custom_certificates)
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError("Expected argument 'fingerprint' to be a str")
        pulumi.set(__self__, "fingerprint", fingerprint)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lb_id and not isinstance(lb_id, str):
            raise TypeError("Expected argument 'lb_id' to be a str")
        pulumi.set(__self__, "lb_id", lb_id)
        if letsencrypts and not isinstance(letsencrypts, list):
            raise TypeError("Expected argument 'letsencrypts' to be a list")
        pulumi.set(__self__, "letsencrypts", letsencrypts)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if not_valid_after and not isinstance(not_valid_after, str):
            raise TypeError("Expected argument 'not_valid_after' to be a str")
        pulumi.set(__self__, "not_valid_after", not_valid_after)
        if not_valid_before and not isinstance(not_valid_before, str):
            raise TypeError("Expected argument 'not_valid_before' to be a str")
        pulumi.set(__self__, "not_valid_before", not_valid_before)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if subject_alternative_names and not isinstance(subject_alternative_names, list):
            raise TypeError("Expected argument 'subject_alternative_names' to be a list")
        pulumi.set(__self__, "subject_alternative_names", subject_alternative_names)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> builtins.str:
        return pulumi.get(self, "common_name")

    @property
    @pulumi.getter(name="customCertificates")
    def custom_certificates(self) -> Sequence['outputs.GetCertificateCustomCertificateResult']:
        return pulumi.get(self, "custom_certificates")

    @property
    @pulumi.getter
    def fingerprint(self) -> builtins.str:
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "lb_id")

    @property
    @pulumi.getter
    def letsencrypts(self) -> Sequence['outputs.GetCertificateLetsencryptResult']:
        return pulumi.get(self, "letsencrypts")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notValidAfter")
    def not_valid_after(self) -> builtins.str:
        return pulumi.get(self, "not_valid_after")

    @property
    @pulumi.getter(name="notValidBefore")
    def not_valid_before(self) -> builtins.str:
        return pulumi.get(self, "not_valid_before")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subjectAlternativeNames")
    def subject_alternative_names(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "subject_alternative_names")


class AwaitableGetCertificateResult(GetCertificateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCertificateResult(
            certificate_id=self.certificate_id,
            common_name=self.common_name,
            custom_certificates=self.custom_certificates,
            fingerprint=self.fingerprint,
            id=self.id,
            lb_id=self.lb_id,
            letsencrypts=self.letsencrypts,
            name=self.name,
            not_valid_after=self.not_valid_after,
            not_valid_before=self.not_valid_before,
            status=self.status,
            subject_alternative_names=self.subject_alternative_names)


def get_certificate(certificate_id: Optional[builtins.str] = None,
                    lb_id: Optional[builtins.str] = None,
                    name: Optional[builtins.str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCertificateResult:
    """
    Get information about Scaleway Load Balancer certificates.

    This data source can prove useful when a module accepts a Load Balancer certificate as an input variable and needs to, for example, determine the security of a certificate for the frontend associated with your domain.

    For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/how-to/add-certificate/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-certificate).

    ## Examples


    :param builtins.str certificate_id: The certificate ID.
           - Only one of `name` and `certificate_id` should be specified.
    :param builtins.str lb_id: The Load Balancer ID this certificate is attached to.
    :param builtins.str name: The name of the Load Balancer certificate.
           - When using a certificate `name` you should specify the `lb-id`
    """
    __args__ = dict()
    __args__['certificateId'] = certificate_id
    __args__['lbId'] = lb_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:loadbalancers/getCertificate:getCertificate', __args__, opts=opts, typ=GetCertificateResult).value

    return AwaitableGetCertificateResult(
        certificate_id=pulumi.get(__ret__, 'certificate_id'),
        common_name=pulumi.get(__ret__, 'common_name'),
        custom_certificates=pulumi.get(__ret__, 'custom_certificates'),
        fingerprint=pulumi.get(__ret__, 'fingerprint'),
        id=pulumi.get(__ret__, 'id'),
        lb_id=pulumi.get(__ret__, 'lb_id'),
        letsencrypts=pulumi.get(__ret__, 'letsencrypts'),
        name=pulumi.get(__ret__, 'name'),
        not_valid_after=pulumi.get(__ret__, 'not_valid_after'),
        not_valid_before=pulumi.get(__ret__, 'not_valid_before'),
        status=pulumi.get(__ret__, 'status'),
        subject_alternative_names=pulumi.get(__ret__, 'subject_alternative_names'))
def get_certificate_output(certificate_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           lb_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCertificateResult]:
    """
    Get information about Scaleway Load Balancer certificates.

    This data source can prove useful when a module accepts a Load Balancer certificate as an input variable and needs to, for example, determine the security of a certificate for the frontend associated with your domain.

    For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/how-to/add-certificate/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-certificate).

    ## Examples


    :param builtins.str certificate_id: The certificate ID.
           - Only one of `name` and `certificate_id` should be specified.
    :param builtins.str lb_id: The Load Balancer ID this certificate is attached to.
    :param builtins.str name: The name of the Load Balancer certificate.
           - When using a certificate `name` you should specify the `lb-id`
    """
    __args__ = dict()
    __args__['certificateId'] = certificate_id
    __args__['lbId'] = lb_id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:loadbalancers/getCertificate:getCertificate', __args__, opts=opts, typ=GetCertificateResult)
    return __ret__.apply(lambda __response__: GetCertificateResult(
        certificate_id=pulumi.get(__response__, 'certificate_id'),
        common_name=pulumi.get(__response__, 'common_name'),
        custom_certificates=pulumi.get(__response__, 'custom_certificates'),
        fingerprint=pulumi.get(__response__, 'fingerprint'),
        id=pulumi.get(__response__, 'id'),
        lb_id=pulumi.get(__response__, 'lb_id'),
        letsencrypts=pulumi.get(__response__, 'letsencrypts'),
        name=pulumi.get(__response__, 'name'),
        not_valid_after=pulumi.get(__response__, 'not_valid_after'),
        not_valid_before=pulumi.get(__response__, 'not_valid_before'),
        status=pulumi.get(__response__, 'status'),
        subject_alternative_names=pulumi.get(__response__, 'subject_alternative_names')))
