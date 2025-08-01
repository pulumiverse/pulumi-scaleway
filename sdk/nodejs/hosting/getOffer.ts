// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

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
export function getOffer(args?: GetOfferArgs, opts?: pulumi.InvokeOptions): Promise<GetOfferResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:hosting/getOffer:getOffer", {
        "controlPanel": args.controlPanel,
        "name": args.name,
        "offerId": args.offerId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getOffer.
 */
export interface GetOfferArgs {
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
 * A collection of values returned by getOffer.
 */
export interface GetOfferResult {
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
    readonly offers: outputs.hosting.GetOfferOffer[];
    /**
     * The offer price.
     */
    readonly price: string;
    /**
     * (deprecated) The offer product.
     *
     * @deprecated The product field is deprecated. Please use the offer field instead.
     */
    readonly products: outputs.hosting.GetOfferProduct[];
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
export function getOfferOutput(args?: GetOfferOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetOfferResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:hosting/getOffer:getOffer", {
        "controlPanel": args.controlPanel,
        "name": args.name,
        "offerId": args.offerId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getOffer.
 */
export interface GetOfferOutputArgs {
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
