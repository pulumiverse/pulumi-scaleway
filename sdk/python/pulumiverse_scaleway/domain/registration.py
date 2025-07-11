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

__all__ = ['RegistrationArgs', 'Registration']

@pulumi.input_type
class RegistrationArgs:
    def __init__(__self__, *,
                 domain_names: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 auto_renew: Optional[pulumi.Input[builtins.bool]] = None,
                 dnssec: Optional[pulumi.Input[builtins.bool]] = None,
                 duration_in_years: Optional[pulumi.Input[builtins.int]] = None,
                 owner_contact: Optional[pulumi.Input['RegistrationOwnerContactArgs']] = None,
                 owner_contact_id: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Registration resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] domain_names: : A list of domain names to be registered.
        :param pulumi.Input[builtins.bool] auto_renew: : Enables or disables auto-renewal.
        :param pulumi.Input[builtins.bool] dnssec: : Enables or disables DNSSEC.
        :param pulumi.Input[builtins.int] duration_in_years: : The registration period in years.
        :param pulumi.Input['RegistrationOwnerContactArgs'] owner_contact: : Details of the owner contact.
        :param pulumi.Input[builtins.str] owner_contact_id: : The ID of an existing owner contact.
        :param pulumi.Input[builtins.str] project_id: : The Scaleway project ID.
        """
        pulumi.set(__self__, "domain_names", domain_names)
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if dnssec is not None:
            pulumi.set(__self__, "dnssec", dnssec)
        if duration_in_years is not None:
            pulumi.set(__self__, "duration_in_years", duration_in_years)
        if owner_contact is not None:
            pulumi.set(__self__, "owner_contact", owner_contact)
        if owner_contact_id is not None:
            pulumi.set(__self__, "owner_contact_id", owner_contact_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        : A list of domain names to be registered.
        """
        return pulumi.get(self, "domain_names")

    @domain_names.setter
    def domain_names(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "domain_names", value)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        : Enables or disables auto-renewal.
        """
        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter
    def dnssec(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        : Enables or disables DNSSEC.
        """
        return pulumi.get(self, "dnssec")

    @dnssec.setter
    def dnssec(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "dnssec", value)

    @property
    @pulumi.getter(name="durationInYears")
    def duration_in_years(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        : The registration period in years.
        """
        return pulumi.get(self, "duration_in_years")

    @duration_in_years.setter
    def duration_in_years(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "duration_in_years", value)

    @property
    @pulumi.getter(name="ownerContact")
    def owner_contact(self) -> Optional[pulumi.Input['RegistrationOwnerContactArgs']]:
        """
        : Details of the owner contact.
        """
        return pulumi.get(self, "owner_contact")

    @owner_contact.setter
    def owner_contact(self, value: Optional[pulumi.Input['RegistrationOwnerContactArgs']]):
        pulumi.set(self, "owner_contact", value)

    @property
    @pulumi.getter(name="ownerContactId")
    def owner_contact_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        : The ID of an existing owner contact.
        """
        return pulumi.get(self, "owner_contact_id")

    @owner_contact_id.setter
    def owner_contact_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "owner_contact_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        : The Scaleway project ID.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)


@pulumi.input_type
class _RegistrationState:
    def __init__(__self__, *,
                 administrative_contacts: Optional[pulumi.Input[Sequence[pulumi.Input['RegistrationAdministrativeContactArgs']]]] = None,
                 auto_renew: Optional[pulumi.Input[builtins.bool]] = None,
                 dnssec: Optional[pulumi.Input[builtins.bool]] = None,
                 domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 ds_records: Optional[pulumi.Input[Sequence[pulumi.Input['RegistrationDsRecordArgs']]]] = None,
                 duration_in_years: Optional[pulumi.Input[builtins.int]] = None,
                 owner_contact: Optional[pulumi.Input['RegistrationOwnerContactArgs']] = None,
                 owner_contact_id: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 task_id: Optional[pulumi.Input[builtins.str]] = None,
                 technical_contacts: Optional[pulumi.Input[Sequence[pulumi.Input['RegistrationTechnicalContactArgs']]]] = None):
        """
        Input properties used for looking up and filtering Registration resources.
        :param pulumi.Input[Sequence[pulumi.Input['RegistrationAdministrativeContactArgs']]] administrative_contacts: : Administrative contact information.
        :param pulumi.Input[builtins.bool] auto_renew: : Enables or disables auto-renewal.
        :param pulumi.Input[builtins.bool] dnssec: : Enables or disables DNSSEC.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] domain_names: : A list of domain names to be registered.
        :param pulumi.Input[Sequence[pulumi.Input['RegistrationDsRecordArgs']]] ds_records: DNSSEC DS record configuration.
        :param pulumi.Input[builtins.int] duration_in_years: : The registration period in years.
        :param pulumi.Input['RegistrationOwnerContactArgs'] owner_contact: : Details of the owner contact.
        :param pulumi.Input[builtins.str] owner_contact_id: : The ID of an existing owner contact.
        :param pulumi.Input[builtins.str] project_id: : The Scaleway project ID.
        :param pulumi.Input[builtins.str] task_id: ID of the task that created the domain.
        :param pulumi.Input[Sequence[pulumi.Input['RegistrationTechnicalContactArgs']]] technical_contacts: : Technical contact information.
        """
        if administrative_contacts is not None:
            pulumi.set(__self__, "administrative_contacts", administrative_contacts)
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if dnssec is not None:
            pulumi.set(__self__, "dnssec", dnssec)
        if domain_names is not None:
            pulumi.set(__self__, "domain_names", domain_names)
        if ds_records is not None:
            pulumi.set(__self__, "ds_records", ds_records)
        if duration_in_years is not None:
            pulumi.set(__self__, "duration_in_years", duration_in_years)
        if owner_contact is not None:
            pulumi.set(__self__, "owner_contact", owner_contact)
        if owner_contact_id is not None:
            pulumi.set(__self__, "owner_contact_id", owner_contact_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if task_id is not None:
            pulumi.set(__self__, "task_id", task_id)
        if technical_contacts is not None:
            pulumi.set(__self__, "technical_contacts", technical_contacts)

    @property
    @pulumi.getter(name="administrativeContacts")
    def administrative_contacts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RegistrationAdministrativeContactArgs']]]]:
        """
        : Administrative contact information.
        """
        return pulumi.get(self, "administrative_contacts")

    @administrative_contacts.setter
    def administrative_contacts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RegistrationAdministrativeContactArgs']]]]):
        pulumi.set(self, "administrative_contacts", value)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        : Enables or disables auto-renewal.
        """
        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter
    def dnssec(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        : Enables or disables DNSSEC.
        """
        return pulumi.get(self, "dnssec")

    @dnssec.setter
    def dnssec(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "dnssec", value)

    @property
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        : A list of domain names to be registered.
        """
        return pulumi.get(self, "domain_names")

    @domain_names.setter
    def domain_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "domain_names", value)

    @property
    @pulumi.getter(name="dsRecords")
    def ds_records(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RegistrationDsRecordArgs']]]]:
        """
        DNSSEC DS record configuration.
        """
        return pulumi.get(self, "ds_records")

    @ds_records.setter
    def ds_records(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RegistrationDsRecordArgs']]]]):
        pulumi.set(self, "ds_records", value)

    @property
    @pulumi.getter(name="durationInYears")
    def duration_in_years(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        : The registration period in years.
        """
        return pulumi.get(self, "duration_in_years")

    @duration_in_years.setter
    def duration_in_years(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "duration_in_years", value)

    @property
    @pulumi.getter(name="ownerContact")
    def owner_contact(self) -> Optional[pulumi.Input['RegistrationOwnerContactArgs']]:
        """
        : Details of the owner contact.
        """
        return pulumi.get(self, "owner_contact")

    @owner_contact.setter
    def owner_contact(self, value: Optional[pulumi.Input['RegistrationOwnerContactArgs']]):
        pulumi.set(self, "owner_contact", value)

    @property
    @pulumi.getter(name="ownerContactId")
    def owner_contact_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        : The ID of an existing owner contact.
        """
        return pulumi.get(self, "owner_contact_id")

    @owner_contact_id.setter
    def owner_contact_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "owner_contact_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        : The Scaleway project ID.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="taskId")
    def task_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the task that created the domain.
        """
        return pulumi.get(self, "task_id")

    @task_id.setter
    def task_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "task_id", value)

    @property
    @pulumi.getter(name="technicalContacts")
    def technical_contacts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RegistrationTechnicalContactArgs']]]]:
        """
        : Technical contact information.
        """
        return pulumi.get(self, "technical_contacts")

    @technical_contacts.setter
    def technical_contacts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RegistrationTechnicalContactArgs']]]]):
        pulumi.set(self, "technical_contacts", value)


@pulumi.type_token("scaleway:domain/registration:Registration")
class Registration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew: Optional[pulumi.Input[builtins.bool]] = None,
                 dnssec: Optional[pulumi.Input[builtins.bool]] = None,
                 domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 duration_in_years: Optional[pulumi.Input[builtins.int]] = None,
                 owner_contact: Optional[pulumi.Input[Union['RegistrationOwnerContactArgs', 'RegistrationOwnerContactArgsDict']]] = None,
                 owner_contact_id: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The `domain.Registration` resource allows you to purchase and manage domain registrations with Scaleway. Using this resource you can register one or more domains for a specified duration, configure auto-renewal and DNSSEC options, and set contact information. You can supply an owner contact either by providing an existing contact ID or by specifying the complete contact details. The resource automatically returns additional contact information (administrative and technical) as provided by the Scaleway API.

        Refer to the [Domains and DNS documentation](https://www.scaleway.com/en/docs/network/domains-and-dns/) and the [API documentation](https://developers.scaleway.com/) for more details.

        ## Example Usage

        ### Purchase a Single Domain

        The following example purchases a domain with a one-year registration period and specifies the owner contact details. Administrative and technical contacts are returned by the API.

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        test = scaleway.domain.Registration("test",
            domain_names=["example.com"],
            duration_in_years=1,
            owner_contact={
                "legal_form": "individual",
                "firstname": "John",
                "lastname": "DOE",
                "email": "john.doe@example.com",
                "phone_number": "+1.23456789",
                "address_line1": "123 Main Street",
                "city": "Paris",
                "zip": "75001",
                "country": "FR",
                "vat_identification_code": "FR12345678901",
                "company_identification_code": "123456789",
            })
        ```

        ### Update Domain Settings

        You can update the resource to enable auto-renewal and DNSSEC:

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        test = scaleway.domain.Registration("test",
            domain_names=["example.com"],
            duration_in_years=1,
            owner_contact={
                "legal_form": "individual",
                "firstname": "John",
                "lastname": "DOE",
                "email": "john.doe@example.com",
                "phone_number": "+1.23456789",
                "address_line1": "123 Main Street",
                "city": "Paris",
                "zip": "75001",
                "country": "FR",
                "vat_identification_code": "FR12345678901",
                "company_identification_code": "123456789",
            },
            auto_renew=True,
            dnssec=True)
        ```

        ### Purchase Multiple Domains

        The following example registers several domains in one go:

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        multi = scaleway.domain.Registration("multi",
            domain_names=[
                "domain1.com",
                "domain2.com",
                "domain3.com",
            ],
            duration_in_years=1,
            owner_contact={
                "legal_form": "individual",
                "firstname": "John",
                "lastname": "DOE",
                "email": "john.doe@example.com",
                "phone_number": "+1.23456789",
                "address_line1": "123 Main Street",
                "city": "Paris",
                "zip": "75001",
                "country": "FR",
                "vat_identification_code": "FR12345678901",
                "company_identification_code": "123456789",
            })
        ```

        ## Contact Blocks

        Each contact block supports the following attributes:

        - `legal_form` (Required, String): Legal form of the contact.
        - `firstname` (Required, String): First name.
        - `lastname` (Required, String): Last name.
        - `company_name` (Optional, String): Company name.
        - `email` (Required, String): Primary email.
        - `phone_number` (Required, String): Primary phone number.
        - `address_line_1` (Required, String): Primary address.
        - `zip` (Required, String): Postal code.
        - `city` (Required, String): City.
        - `country` (Required, String): Country code (ISO format).
        - `vat_identification_code` (Required, String): VAT identification code.
        - `company_identification_code` (Required, String): Company identification code.

        ## Import

        To import an existing domain registration, use:

        bash

        ```sh
        $ pulumi import scaleway:domain/registration:Registration test <project_id>/<task_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] auto_renew: : Enables or disables auto-renewal.
        :param pulumi.Input[builtins.bool] dnssec: : Enables or disables DNSSEC.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] domain_names: : A list of domain names to be registered.
        :param pulumi.Input[builtins.int] duration_in_years: : The registration period in years.
        :param pulumi.Input[Union['RegistrationOwnerContactArgs', 'RegistrationOwnerContactArgsDict']] owner_contact: : Details of the owner contact.
        :param pulumi.Input[builtins.str] owner_contact_id: : The ID of an existing owner contact.
        :param pulumi.Input[builtins.str] project_id: : The Scaleway project ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegistrationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `domain.Registration` resource allows you to purchase and manage domain registrations with Scaleway. Using this resource you can register one or more domains for a specified duration, configure auto-renewal and DNSSEC options, and set contact information. You can supply an owner contact either by providing an existing contact ID or by specifying the complete contact details. The resource automatically returns additional contact information (administrative and technical) as provided by the Scaleway API.

        Refer to the [Domains and DNS documentation](https://www.scaleway.com/en/docs/network/domains-and-dns/) and the [API documentation](https://developers.scaleway.com/) for more details.

        ## Example Usage

        ### Purchase a Single Domain

        The following example purchases a domain with a one-year registration period and specifies the owner contact details. Administrative and technical contacts are returned by the API.

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        test = scaleway.domain.Registration("test",
            domain_names=["example.com"],
            duration_in_years=1,
            owner_contact={
                "legal_form": "individual",
                "firstname": "John",
                "lastname": "DOE",
                "email": "john.doe@example.com",
                "phone_number": "+1.23456789",
                "address_line1": "123 Main Street",
                "city": "Paris",
                "zip": "75001",
                "country": "FR",
                "vat_identification_code": "FR12345678901",
                "company_identification_code": "123456789",
            })
        ```

        ### Update Domain Settings

        You can update the resource to enable auto-renewal and DNSSEC:

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        test = scaleway.domain.Registration("test",
            domain_names=["example.com"],
            duration_in_years=1,
            owner_contact={
                "legal_form": "individual",
                "firstname": "John",
                "lastname": "DOE",
                "email": "john.doe@example.com",
                "phone_number": "+1.23456789",
                "address_line1": "123 Main Street",
                "city": "Paris",
                "zip": "75001",
                "country": "FR",
                "vat_identification_code": "FR12345678901",
                "company_identification_code": "123456789",
            },
            auto_renew=True,
            dnssec=True)
        ```

        ### Purchase Multiple Domains

        The following example registers several domains in one go:

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        multi = scaleway.domain.Registration("multi",
            domain_names=[
                "domain1.com",
                "domain2.com",
                "domain3.com",
            ],
            duration_in_years=1,
            owner_contact={
                "legal_form": "individual",
                "firstname": "John",
                "lastname": "DOE",
                "email": "john.doe@example.com",
                "phone_number": "+1.23456789",
                "address_line1": "123 Main Street",
                "city": "Paris",
                "zip": "75001",
                "country": "FR",
                "vat_identification_code": "FR12345678901",
                "company_identification_code": "123456789",
            })
        ```

        ## Contact Blocks

        Each contact block supports the following attributes:

        - `legal_form` (Required, String): Legal form of the contact.
        - `firstname` (Required, String): First name.
        - `lastname` (Required, String): Last name.
        - `company_name` (Optional, String): Company name.
        - `email` (Required, String): Primary email.
        - `phone_number` (Required, String): Primary phone number.
        - `address_line_1` (Required, String): Primary address.
        - `zip` (Required, String): Postal code.
        - `city` (Required, String): City.
        - `country` (Required, String): Country code (ISO format).
        - `vat_identification_code` (Required, String): VAT identification code.
        - `company_identification_code` (Required, String): Company identification code.

        ## Import

        To import an existing domain registration, use:

        bash

        ```sh
        $ pulumi import scaleway:domain/registration:Registration test <project_id>/<task_id>
        ```

        :param str resource_name: The name of the resource.
        :param RegistrationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegistrationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew: Optional[pulumi.Input[builtins.bool]] = None,
                 dnssec: Optional[pulumi.Input[builtins.bool]] = None,
                 domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 duration_in_years: Optional[pulumi.Input[builtins.int]] = None,
                 owner_contact: Optional[pulumi.Input[Union['RegistrationOwnerContactArgs', 'RegistrationOwnerContactArgsDict']]] = None,
                 owner_contact_id: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RegistrationArgs.__new__(RegistrationArgs)

            __props__.__dict__["auto_renew"] = auto_renew
            __props__.__dict__["dnssec"] = dnssec
            if domain_names is None and not opts.urn:
                raise TypeError("Missing required property 'domain_names'")
            __props__.__dict__["domain_names"] = domain_names
            __props__.__dict__["duration_in_years"] = duration_in_years
            __props__.__dict__["owner_contact"] = owner_contact
            __props__.__dict__["owner_contact_id"] = owner_contact_id
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["administrative_contacts"] = None
            __props__.__dict__["ds_records"] = None
            __props__.__dict__["task_id"] = None
            __props__.__dict__["technical_contacts"] = None
        super(Registration, __self__).__init__(
            'scaleway:domain/registration:Registration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            administrative_contacts: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RegistrationAdministrativeContactArgs', 'RegistrationAdministrativeContactArgsDict']]]]] = None,
            auto_renew: Optional[pulumi.Input[builtins.bool]] = None,
            dnssec: Optional[pulumi.Input[builtins.bool]] = None,
            domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            ds_records: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RegistrationDsRecordArgs', 'RegistrationDsRecordArgsDict']]]]] = None,
            duration_in_years: Optional[pulumi.Input[builtins.int]] = None,
            owner_contact: Optional[pulumi.Input[Union['RegistrationOwnerContactArgs', 'RegistrationOwnerContactArgsDict']]] = None,
            owner_contact_id: Optional[pulumi.Input[builtins.str]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None,
            task_id: Optional[pulumi.Input[builtins.str]] = None,
            technical_contacts: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RegistrationTechnicalContactArgs', 'RegistrationTechnicalContactArgsDict']]]]] = None) -> 'Registration':
        """
        Get an existing Registration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RegistrationAdministrativeContactArgs', 'RegistrationAdministrativeContactArgsDict']]]] administrative_contacts: : Administrative contact information.
        :param pulumi.Input[builtins.bool] auto_renew: : Enables or disables auto-renewal.
        :param pulumi.Input[builtins.bool] dnssec: : Enables or disables DNSSEC.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] domain_names: : A list of domain names to be registered.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RegistrationDsRecordArgs', 'RegistrationDsRecordArgsDict']]]] ds_records: DNSSEC DS record configuration.
        :param pulumi.Input[builtins.int] duration_in_years: : The registration period in years.
        :param pulumi.Input[Union['RegistrationOwnerContactArgs', 'RegistrationOwnerContactArgsDict']] owner_contact: : Details of the owner contact.
        :param pulumi.Input[builtins.str] owner_contact_id: : The ID of an existing owner contact.
        :param pulumi.Input[builtins.str] project_id: : The Scaleway project ID.
        :param pulumi.Input[builtins.str] task_id: ID of the task that created the domain.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RegistrationTechnicalContactArgs', 'RegistrationTechnicalContactArgsDict']]]] technical_contacts: : Technical contact information.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RegistrationState.__new__(_RegistrationState)

        __props__.__dict__["administrative_contacts"] = administrative_contacts
        __props__.__dict__["auto_renew"] = auto_renew
        __props__.__dict__["dnssec"] = dnssec
        __props__.__dict__["domain_names"] = domain_names
        __props__.__dict__["ds_records"] = ds_records
        __props__.__dict__["duration_in_years"] = duration_in_years
        __props__.__dict__["owner_contact"] = owner_contact
        __props__.__dict__["owner_contact_id"] = owner_contact_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["task_id"] = task_id
        __props__.__dict__["technical_contacts"] = technical_contacts
        return Registration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="administrativeContacts")
    def administrative_contacts(self) -> pulumi.Output[Sequence['outputs.RegistrationAdministrativeContact']]:
        """
        : Administrative contact information.
        """
        return pulumi.get(self, "administrative_contacts")

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        : Enables or disables auto-renewal.
        """
        return pulumi.get(self, "auto_renew")

    @property
    @pulumi.getter
    def dnssec(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        : Enables or disables DNSSEC.
        """
        return pulumi.get(self, "dnssec")

    @property
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        : A list of domain names to be registered.
        """
        return pulumi.get(self, "domain_names")

    @property
    @pulumi.getter(name="dsRecords")
    def ds_records(self) -> pulumi.Output[Sequence['outputs.RegistrationDsRecord']]:
        """
        DNSSEC DS record configuration.
        """
        return pulumi.get(self, "ds_records")

    @property
    @pulumi.getter(name="durationInYears")
    def duration_in_years(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        : The registration period in years.
        """
        return pulumi.get(self, "duration_in_years")

    @property
    @pulumi.getter(name="ownerContact")
    def owner_contact(self) -> pulumi.Output['outputs.RegistrationOwnerContact']:
        """
        : Details of the owner contact.
        """
        return pulumi.get(self, "owner_contact")

    @property
    @pulumi.getter(name="ownerContactId")
    def owner_contact_id(self) -> pulumi.Output[builtins.str]:
        """
        : The ID of an existing owner contact.
        """
        return pulumi.get(self, "owner_contact_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        """
        : The Scaleway project ID.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="taskId")
    def task_id(self) -> pulumi.Output[builtins.str]:
        """
        ID of the task that created the domain.
        """
        return pulumi.get(self, "task_id")

    @property
    @pulumi.getter(name="technicalContacts")
    def technical_contacts(self) -> pulumi.Output[Sequence['outputs.RegistrationTechnicalContact']]:
        """
        : Technical contact information.
        """
        return pulumi.get(self, "technical_contacts")

