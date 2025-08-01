// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about multiple IP addresses managed by Scaleway's IP Address Management (IPAM) service.
 *
 * For more information about IPAM, see the main [documentation](https://www.scaleway.com/en/docs/vpc/concepts/#ipam).
 *
 * ## Examples
 *
 * ### By tag
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const byTag = scaleway.ipam.getIps({
 *     tags: ["tag"],
 * });
 * ```
 *
 * ### By type and resource
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const vpc01 = new scaleway.network.Vpc("vpc01", {name: "my vpc"});
 * const pn01 = new scaleway.network.PrivateNetwork("pn01", {
 *     vpcId: vpc01.id,
 *     ipv4Subnet: {
 *         subnet: "172.16.32.0/22",
 *     },
 * });
 * const redis01 = new scaleway.redis.Cluster("redis01", {
 *     name: "my_redis_cluster",
 *     version: "7.0.5",
 *     nodeType: "RED1-XS",
 *     userName: "my_initial_user",
 *     password: "thiZ_is_v&ry_s3cret",
 *     clusterSize: 3,
 *     privateNetworks: [{
 *         id: pn01.id,
 *     }],
 * });
 * const byTypeAndResource = scaleway.ipam.getIpsOutput({
 *     type: "ipv4",
 *     resource: {
 *         id: redis01.id,
 *         type: "redis_cluster",
 *     },
 * });
 * ```
 */
/** @deprecated scaleway.index/getipamips.getIpamIps has been deprecated in favor of scaleway.ipam/getips.getIps */
export function getIpamIps(args?: GetIpamIpsArgs, opts?: pulumi.InvokeOptions): Promise<GetIpamIpsResult> {
    pulumi.log.warn("getIpamIps is deprecated: scaleway.index/getipamips.getIpamIps has been deprecated in favor of scaleway.ipam/getips.getIps")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getIpamIps:getIpamIps", {
        "attached": args.attached,
        "macAddress": args.macAddress,
        "privateNetworkId": args.privateNetworkId,
        "projectId": args.projectId,
        "region": args.region,
        "resource": args.resource,
        "tags": args.tags,
        "type": args.type,
        "zonal": args.zonal,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpamIps.
 */
export interface GetIpamIpsArgs {
    /**
     * Defines whether to filter only for IPs which are attached to a resource.
     */
    attached?: boolean;
    /**
     * The linked MAC address to filter for.
     */
    macAddress?: string;
    /**
     * The ID of the Private Network to filter for.
     */
    privateNetworkId?: string;
    /**
     * The ID of the Project to filter for.
     */
    projectId?: string;
    /**
     * The region to filter for.
     */
    region?: string;
    /**
     * Filter for a resource attached to the IP, using resource ID, type or name.
     */
    resource?: inputs.GetIpamIpsResource;
    /**
     * The IP tags to filter for.
     */
    tags?: string[];
    /**
     * The type of IP to filter for (`ipv4` or `ipv6`).
     */
    type?: string;
    /**
     * Only IPs that are zonal, and in this zone, will be returned.
     */
    zonal?: string;
}

/**
 * A collection of values returned by getIpamIps.
 */
export interface GetIpamIpsResult {
    readonly attached?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of found IPs.
     */
    readonly ips: outputs.GetIpamIpsIp[];
    /**
     * The associated MAC address.
     */
    readonly macAddress?: string;
    readonly organizationId: string;
    readonly privateNetworkId?: string;
    /**
     * The ID of the Project the resource is associated with.
     */
    readonly projectId: string;
    /**
     * The region of the IP.
     */
    readonly region: string;
    /**
     * The list of public IPs attached to the resource.
     */
    readonly resource?: outputs.GetIpamIpsResource;
    /**
     * The tags associated with the IP.
     */
    readonly tags?: string[];
    /**
     * The type of resource.
     */
    readonly type?: string;
    readonly zonal: string;
}
/**
 * Gets information about multiple IP addresses managed by Scaleway's IP Address Management (IPAM) service.
 *
 * For more information about IPAM, see the main [documentation](https://www.scaleway.com/en/docs/vpc/concepts/#ipam).
 *
 * ## Examples
 *
 * ### By tag
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const byTag = scaleway.ipam.getIps({
 *     tags: ["tag"],
 * });
 * ```
 *
 * ### By type and resource
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const vpc01 = new scaleway.network.Vpc("vpc01", {name: "my vpc"});
 * const pn01 = new scaleway.network.PrivateNetwork("pn01", {
 *     vpcId: vpc01.id,
 *     ipv4Subnet: {
 *         subnet: "172.16.32.0/22",
 *     },
 * });
 * const redis01 = new scaleway.redis.Cluster("redis01", {
 *     name: "my_redis_cluster",
 *     version: "7.0.5",
 *     nodeType: "RED1-XS",
 *     userName: "my_initial_user",
 *     password: "thiZ_is_v&ry_s3cret",
 *     clusterSize: 3,
 *     privateNetworks: [{
 *         id: pn01.id,
 *     }],
 * });
 * const byTypeAndResource = scaleway.ipam.getIpsOutput({
 *     type: "ipv4",
 *     resource: {
 *         id: redis01.id,
 *         type: "redis_cluster",
 *     },
 * });
 * ```
 */
/** @deprecated scaleway.index/getipamips.getIpamIps has been deprecated in favor of scaleway.ipam/getips.getIps */
export function getIpamIpsOutput(args?: GetIpamIpsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetIpamIpsResult> {
    pulumi.log.warn("getIpamIps is deprecated: scaleway.index/getipamips.getIpamIps has been deprecated in favor of scaleway.ipam/getips.getIps")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getIpamIps:getIpamIps", {
        "attached": args.attached,
        "macAddress": args.macAddress,
        "privateNetworkId": args.privateNetworkId,
        "projectId": args.projectId,
        "region": args.region,
        "resource": args.resource,
        "tags": args.tags,
        "type": args.type,
        "zonal": args.zonal,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpamIps.
 */
export interface GetIpamIpsOutputArgs {
    /**
     * Defines whether to filter only for IPs which are attached to a resource.
     */
    attached?: pulumi.Input<boolean>;
    /**
     * The linked MAC address to filter for.
     */
    macAddress?: pulumi.Input<string>;
    /**
     * The ID of the Private Network to filter for.
     */
    privateNetworkId?: pulumi.Input<string>;
    /**
     * The ID of the Project to filter for.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region to filter for.
     */
    region?: pulumi.Input<string>;
    /**
     * Filter for a resource attached to the IP, using resource ID, type or name.
     */
    resource?: pulumi.Input<inputs.GetIpamIpsResourceArgs>;
    /**
     * The IP tags to filter for.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of IP to filter for (`ipv4` or `ipv6`).
     */
    type?: pulumi.Input<string>;
    /**
     * Only IPs that are zonal, and in this zone, will be returned.
     */
    zonal?: pulumi.Input<string>;
}
