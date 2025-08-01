// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about multiple Virtual Private Clouds.
 */
/** @deprecated scaleway.index/getvpcs.getVpcs has been deprecated in favor of scaleway.network/getvpcs.getVpcs */
export function getVpcs(args?: GetVpcsArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcsResult> {
    pulumi.log.warn("getVpcs is deprecated: scaleway.index/getvpcs.getVpcs has been deprecated in favor of scaleway.network/getvpcs.getVpcs")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getVpcs:getVpcs", {
        "name": args.name,
        "projectId": args.projectId,
        "region": args.region,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcs.
 */
export interface GetVpcsArgs {
    /**
     * The VPC name to filter for. VPCs with a similar name are listed.
     */
    name?: string;
    /**
     * The ID of the Project the VPC is associated with.
     */
    projectId?: string;
    /**
     * `region`). The region in which the VPCs exist.
     */
    region?: string;
    /**
     * List of tags to filter for. VPCs with these exact tags are listed.
     */
    tags?: string[];
}

/**
 * A collection of values returned by getVpcs.
 */
export interface GetVpcsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    /**
     * The Organization ID the VPC is associated with.
     */
    readonly organizationId: string;
    /**
     * The ID of the Project the VPC is associated with.
     */
    readonly projectId: string;
    readonly region: string;
    readonly tags?: string[];
    /**
     * List of retrieved VPCs
     */
    readonly vpcs: outputs.GetVpcsVpc[];
}
/**
 * Gets information about multiple Virtual Private Clouds.
 */
/** @deprecated scaleway.index/getvpcs.getVpcs has been deprecated in favor of scaleway.network/getvpcs.getVpcs */
export function getVpcsOutput(args?: GetVpcsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVpcsResult> {
    pulumi.log.warn("getVpcs is deprecated: scaleway.index/getvpcs.getVpcs has been deprecated in favor of scaleway.network/getvpcs.getVpcs")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getVpcs:getVpcs", {
        "name": args.name,
        "projectId": args.projectId,
        "region": args.region,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcs.
 */
export interface GetVpcsOutputArgs {
    /**
     * The VPC name to filter for. VPCs with a similar name are listed.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the Project the VPC is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region in which the VPCs exist.
     */
    region?: pulumi.Input<string>;
    /**
     * List of tags to filter for. VPCs with these exact tags are listed.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
