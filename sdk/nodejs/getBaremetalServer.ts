// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about a baremetal server.
 * For more information, see the [API documentation](https://developers.scaleway.com/en/products/baremetal/api).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * // Get info by server name
 * const byName = scaleway.elasticmetal.getServer({
 *     name: "foobar",
 *     zone: "fr-par-2",
 * });
 * // Get info by server id
 * const byId = scaleway.elasticmetal.getServer({
 *     serverId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
/** @deprecated scaleway.index/getbaremetalserver.getBaremetalServer has been deprecated in favor of scaleway.elasticmetal/getserver.getServer */
export function getBaremetalServer(args?: GetBaremetalServerArgs, opts?: pulumi.InvokeOptions): Promise<GetBaremetalServerResult> {
    pulumi.log.warn("getBaremetalServer is deprecated: scaleway.index/getbaremetalserver.getBaremetalServer has been deprecated in favor of scaleway.elasticmetal/getserver.getServer")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getBaremetalServer:getBaremetalServer", {
        "name": args.name,
        "projectId": args.projectId,
        "serverId": args.serverId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getBaremetalServer.
 */
export interface GetBaremetalServerArgs {
    /**
     * The server name. Only one of `name` and `serverId` should be specified.
     */
    name?: string;
    /**
     * The ID of the project the baremetal server is associated with.
     */
    projectId?: string;
    serverId?: string;
    /**
     * `zone`) The zone in which the server exists.
     */
    zone?: string;
}

/**
 * A collection of values returned by getBaremetalServer.
 */
export interface GetBaremetalServerResult {
    readonly description: string;
    readonly domain: string;
    readonly hostname: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly installConfigAfterward: boolean;
    readonly ips: outputs.GetBaremetalServerIp[];
    readonly ipv4s: outputs.GetBaremetalServerIpv4[];
    readonly ipv6s: outputs.GetBaremetalServerIpv6[];
    readonly name?: string;
    readonly offer: string;
    readonly offerId: string;
    readonly offerName: string;
    readonly options: outputs.GetBaremetalServerOption[];
    readonly organizationId: string;
    readonly os: string;
    readonly osName: string;
    readonly partitioning: string;
    readonly password: string;
    readonly privateIps: outputs.GetBaremetalServerPrivateIp[];
    readonly privateNetworks: outputs.GetBaremetalServerPrivateNetwork[];
    readonly projectId?: string;
    readonly reinstallOnConfigChanges: boolean;
    readonly serverId?: string;
    readonly servicePassword: string;
    readonly serviceUser: string;
    readonly sshKeyIds: string[];
    readonly tags: string[];
    readonly user: string;
    readonly zone?: string;
}
/**
 * Gets information about a baremetal server.
 * For more information, see the [API documentation](https://developers.scaleway.com/en/products/baremetal/api).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * // Get info by server name
 * const byName = scaleway.elasticmetal.getServer({
 *     name: "foobar",
 *     zone: "fr-par-2",
 * });
 * // Get info by server id
 * const byId = scaleway.elasticmetal.getServer({
 *     serverId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
/** @deprecated scaleway.index/getbaremetalserver.getBaremetalServer has been deprecated in favor of scaleway.elasticmetal/getserver.getServer */
export function getBaremetalServerOutput(args?: GetBaremetalServerOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetBaremetalServerResult> {
    pulumi.log.warn("getBaremetalServer is deprecated: scaleway.index/getbaremetalserver.getBaremetalServer has been deprecated in favor of scaleway.elasticmetal/getserver.getServer")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getBaremetalServer:getBaremetalServer", {
        "name": args.name,
        "projectId": args.projectId,
        "serverId": args.serverId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getBaremetalServer.
 */
export interface GetBaremetalServerOutputArgs {
    /**
     * The server name. Only one of `name` and `serverId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project the baremetal server is associated with.
     */
    projectId?: pulumi.Input<string>;
    serverId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the server exists.
     */
    zone?: pulumi.Input<string>;
}
