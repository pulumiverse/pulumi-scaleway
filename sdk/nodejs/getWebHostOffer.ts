// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about a webhosting offer.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * // Get info by offer name
 * const byName = scaleway.hosting.getOffer({
 *     name: "performance",
 *     controlPanel: "Cpanel",
 * });
 * // Get info by offer id
 * const byId = scaleway.hosting.getOffer({
 *     offerId: "de2426b4-a9e9-11ec-b909-0242ac120002",
 * });
 * ```
 */
/** @deprecated scaleway.index/getwebhostoffer.getWebHostOffer has been deprecated in favor of scaleway.hosting/getoffer.getOffer */
export function getWebHostOffer(args?: GetWebHostOfferArgs, opts?: pulumi.InvokeOptions): Promise<GetWebHostOfferResult> {
    pulumi.log.warn("getWebHostOffer is deprecated: scaleway.index/getwebhostoffer.getWebHostOffer has been deprecated in favor of scaleway.hosting/getoffer.getOffer")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getWebHostOffer:getWebHostOffer", {
        "controlPanel": args.controlPanel,
        "name": args.name,
        "offerId": args.offerId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getWebHostOffer.
 */
export interface GetWebHostOfferArgs {
    /**
     * Name of the control panel (Cpanel or Plesk). This argument is only used when `offerId` is not specified.
     */
    controlPanel?: string;
    /**
     * The offer name. Only one of `name` and `offerId` should be specified.
     */
    name?: string;
    /**
     * The offer id. Only one of `name` and `offerId` should be specified.
     */
    offerId?: string;
    /**
     * `region`) The region in which offer exists.
     */
    region?: string;
}

/**
 * A collection of values returned by getWebHostOffer.
 */
export interface GetWebHostOfferResult {
    /**
     * The billing operation identifier for the option.
     */
    readonly billingOperationPath: string;
    readonly controlPanel?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the option.
     */
    readonly name?: string;
    readonly offerId?: string;
    /**
     * The detailed offer of the hosting.
     */
    readonly offers: outputs.GetWebHostOfferOffer[];
    /**
     * The offer price.
     */
    readonly price: string;
    /**
     * (deprecated) The offer product.
     *
     * @deprecated The product field is deprecated. Please use the offer field instead.
     */
    readonly products: outputs.GetWebHostOfferProduct[];
    readonly region: string;
}
/**
 * Gets information about a webhosting offer.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * // Get info by offer name
 * const byName = scaleway.hosting.getOffer({
 *     name: "performance",
 *     controlPanel: "Cpanel",
 * });
 * // Get info by offer id
 * const byId = scaleway.hosting.getOffer({
 *     offerId: "de2426b4-a9e9-11ec-b909-0242ac120002",
 * });
 * ```
 */
/** @deprecated scaleway.index/getwebhostoffer.getWebHostOffer has been deprecated in favor of scaleway.hosting/getoffer.getOffer */
export function getWebHostOfferOutput(args?: GetWebHostOfferOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetWebHostOfferResult> {
    pulumi.log.warn("getWebHostOffer is deprecated: scaleway.index/getwebhostoffer.getWebHostOffer has been deprecated in favor of scaleway.hosting/getoffer.getOffer")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getWebHostOffer:getWebHostOffer", {
        "controlPanel": args.controlPanel,
        "name": args.name,
        "offerId": args.offerId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getWebHostOffer.
 */
export interface GetWebHostOfferOutputArgs {
    /**
     * Name of the control panel (Cpanel or Plesk). This argument is only used when `offerId` is not specified.
     */
    controlPanel?: pulumi.Input<string>;
    /**
     * The offer name. Only one of `name` and `offerId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The offer id. Only one of `name` and `offerId` should be specified.
     */
    offerId?: pulumi.Input<string>;
    /**
     * `region`) The region in which offer exists.
     */
    region?: pulumi.Input<string>;
}
