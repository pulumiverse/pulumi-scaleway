// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about a baremetal option.
 * For more information, see the [API documentation](https://developers.scaleway.com/en/products/baremetal/api).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * // Get info by option name
 * const byName = scaleway.elasticmetal.getOption({
 *     name: "Remote Access",
 * });
 * // Get info by option id
 * const byId = scaleway.elasticmetal.getOption({
 *     optionId: "931df052-d713-4674-8b58-96a63244c8e2",
 * });
 * ```
 */
/** @deprecated scaleway.index/getbaremetaloption.getBaremetalOption has been deprecated in favor of scaleway.elasticmetal/getoption.getOption */
export function getBaremetalOption(args?: GetBaremetalOptionArgs, opts?: pulumi.InvokeOptions): Promise<GetBaremetalOptionResult> {
    pulumi.log.warn("getBaremetalOption is deprecated: scaleway.index/getbaremetaloption.getBaremetalOption has been deprecated in favor of scaleway.elasticmetal/getoption.getOption")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getBaremetalOption:getBaremetalOption", {
        "name": args.name,
        "optionId": args.optionId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getBaremetalOption.
 */
export interface GetBaremetalOptionArgs {
    /**
     * The option name. Only one of `name` and `optionId` should be specified.
     */
    name?: string;
    /**
     * The option id. Only one of `name` and `optionId` should be specified.
     */
    optionId?: string;
    /**
     * `zone`) The zone in which the option exists.
     */
    zone?: string;
}

/**
 * A collection of values returned by getBaremetalOption.
 */
export interface GetBaremetalOptionResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Is false if the option could not be added or removed.
     */
    readonly manageable: boolean;
    /**
     * The name of the option.
     */
    readonly name?: string;
    readonly optionId?: string;
    readonly zone: string;
}
/**
 * Gets information about a baremetal option.
 * For more information, see the [API documentation](https://developers.scaleway.com/en/products/baremetal/api).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * // Get info by option name
 * const byName = scaleway.elasticmetal.getOption({
 *     name: "Remote Access",
 * });
 * // Get info by option id
 * const byId = scaleway.elasticmetal.getOption({
 *     optionId: "931df052-d713-4674-8b58-96a63244c8e2",
 * });
 * ```
 */
/** @deprecated scaleway.index/getbaremetaloption.getBaremetalOption has been deprecated in favor of scaleway.elasticmetal/getoption.getOption */
export function getBaremetalOptionOutput(args?: GetBaremetalOptionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetBaremetalOptionResult> {
    pulumi.log.warn("getBaremetalOption is deprecated: scaleway.index/getbaremetaloption.getBaremetalOption has been deprecated in favor of scaleway.elasticmetal/getoption.getOption")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getBaremetalOption:getBaremetalOption", {
        "name": args.name,
        "optionId": args.optionId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getBaremetalOption.
 */
export interface GetBaremetalOptionOutputArgs {
    /**
     * The option name. Only one of `name` and `optionId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The option id. Only one of `name` and `optionId` should be specified.
     */
    optionId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the option exists.
     */
    zone?: pulumi.Input<string>;
}
