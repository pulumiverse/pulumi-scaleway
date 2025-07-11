// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Gets information about a Public Gateway public flexible IP address.
 *
 * For further information, please see the [API documentation](https://www.scaleway.com/en/developers/api/public-gateway/#path-ips-list-ips).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const main = new scaleway.network.PublicGatewayIp("main", {});
 * const ipById = scaleway.network.getPublicGatewayIpOutput({
 *     ipId: main.id,
 * });
 * ```
 */
export function getPublicGatewayIp(args?: GetPublicGatewayIpArgs, opts?: pulumi.InvokeOptions): Promise<GetPublicGatewayIpResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:network/getPublicGatewayIp:getPublicGatewayIp", {
        "ipId": args.ipId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPublicGatewayIp.
 */
export interface GetPublicGatewayIpArgs {
    ipId?: string;
}

/**
 * A collection of values returned by getPublicGatewayIp.
 */
export interface GetPublicGatewayIpResult {
    readonly address: string;
    readonly createdAt: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipId?: string;
    readonly organizationId: string;
    readonly projectId: string;
    readonly reverse: string;
    readonly tags: string[];
    readonly updatedAt: string;
    readonly zone: string;
}
/**
 * Gets information about a Public Gateway public flexible IP address.
 *
 * For further information, please see the [API documentation](https://www.scaleway.com/en/developers/api/public-gateway/#path-ips-list-ips).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const main = new scaleway.network.PublicGatewayIp("main", {});
 * const ipById = scaleway.network.getPublicGatewayIpOutput({
 *     ipId: main.id,
 * });
 * ```
 */
export function getPublicGatewayIpOutput(args?: GetPublicGatewayIpOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPublicGatewayIpResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:network/getPublicGatewayIp:getPublicGatewayIp", {
        "ipId": args.ipId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPublicGatewayIp.
 */
export interface GetPublicGatewayIpOutputArgs {
    ipId?: pulumi.Input<string>;
}
