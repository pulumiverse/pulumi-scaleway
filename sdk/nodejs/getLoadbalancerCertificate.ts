// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get information about Scaleway Load Balancer certificates.
 *
 * This data source can prove useful when a module accepts a Load Balancer certificate as an input variable and needs to, for example, determine the security of a certificate for the frontend associated with your domain.
 *
 * For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/how-to/add-certificate/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-certificate).
 *
 * ## Examples
 */
/** @deprecated scaleway.index/getloadbalancercertificate.getLoadbalancerCertificate has been deprecated in favor of scaleway.loadbalancers/getcertificate.getCertificate */
export function getLoadbalancerCertificate(args?: GetLoadbalancerCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetLoadbalancerCertificateResult> {
    pulumi.log.warn("getLoadbalancerCertificate is deprecated: scaleway.index/getloadbalancercertificate.getLoadbalancerCertificate has been deprecated in favor of scaleway.loadbalancers/getcertificate.getCertificate")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getLoadbalancerCertificate:getLoadbalancerCertificate", {
        "certificateId": args.certificateId,
        "lbId": args.lbId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getLoadbalancerCertificate.
 */
export interface GetLoadbalancerCertificateArgs {
    /**
     * The certificate ID.
     * - Only one of `name` and `certificateId` should be specified.
     */
    certificateId?: string;
    /**
     * The Load Balancer ID this certificate is attached to.
     */
    lbId?: string;
    /**
     * The name of the Load Balancer certificate.
     * - When using a certificate `name` you should specify the `lb-id`
     */
    name?: string;
}

/**
 * A collection of values returned by getLoadbalancerCertificate.
 */
export interface GetLoadbalancerCertificateResult {
    readonly certificateId?: string;
    readonly commonName: string;
    readonly customCertificates: outputs.GetLoadbalancerCertificateCustomCertificate[];
    readonly fingerprint: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lbId?: string;
    readonly letsencrypts: outputs.GetLoadbalancerCertificateLetsencrypt[];
    readonly name?: string;
    readonly notValidAfter: string;
    readonly notValidBefore: string;
    readonly status: string;
    readonly subjectAlternativeNames: string[];
}
/**
 * Get information about Scaleway Load Balancer certificates.
 *
 * This data source can prove useful when a module accepts a Load Balancer certificate as an input variable and needs to, for example, determine the security of a certificate for the frontend associated with your domain.
 *
 * For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/how-to/add-certificate/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-certificate).
 *
 * ## Examples
 */
/** @deprecated scaleway.index/getloadbalancercertificate.getLoadbalancerCertificate has been deprecated in favor of scaleway.loadbalancers/getcertificate.getCertificate */
export function getLoadbalancerCertificateOutput(args?: GetLoadbalancerCertificateOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLoadbalancerCertificateResult> {
    pulumi.log.warn("getLoadbalancerCertificate is deprecated: scaleway.index/getloadbalancercertificate.getLoadbalancerCertificate has been deprecated in favor of scaleway.loadbalancers/getcertificate.getCertificate")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getLoadbalancerCertificate:getLoadbalancerCertificate", {
        "certificateId": args.certificateId,
        "lbId": args.lbId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getLoadbalancerCertificate.
 */
export interface GetLoadbalancerCertificateOutputArgs {
    /**
     * The certificate ID.
     * - Only one of `name` and `certificateId` should be specified.
     */
    certificateId?: pulumi.Input<string>;
    /**
     * The Load Balancer ID this certificate is attached to.
     */
    lbId?: pulumi.Input<string>;
    /**
     * The name of the Load Balancer certificate.
     * - When using a certificate `name` you should specify the `lb-id`
     */
    name?: pulumi.Input<string>;
}
