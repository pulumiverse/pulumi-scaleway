// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `scaleway.BlockSnapshot` data source is used to retrieve information about a Block Storage volume snapshot.
 *
 * Refer to the Block Storage [product documentation](https://www.scaleway.com/en/docs/storage/block/) and [API documentation](https://www.scaleway.com/en/developers/api/block/) for more information.
 */
export function getBlockSnapshot(args?: GetBlockSnapshotArgs, opts?: pulumi.InvokeOptions): Promise<GetBlockSnapshotResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getBlockSnapshot:getBlockSnapshot", {
        "name": args.name,
        "projectId": args.projectId,
        "snapshotId": args.snapshotId,
        "volumeId": args.volumeId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getBlockSnapshot.
 */
export interface GetBlockSnapshotArgs {
    /**
     * The name of the snapshot. Only one of name or snapshotId should be specified.
     */
    name?: string;
    /**
     * The unique identifier of the Project to which the snapshot is associated.
     */
    projectId?: string;
    /**
     * The unique identifier of the snapshot. Only one of `name` and `snapshotId` should be specified.
     */
    snapshotId?: string;
    /**
     * The unique identifier of the volume from which the snapshot was created.
     */
    volumeId?: string;
    /**
     * ) The zone in which the snapshot exists.
     */
    zone?: string;
}

/**
 * A collection of values returned by getBlockSnapshot.
 */
export interface GetBlockSnapshotResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    readonly projectId?: string;
    readonly snapshotId?: string;
    readonly tags: string[];
    readonly volumeId?: string;
    readonly zone?: string;
}
/**
 * The `scaleway.BlockSnapshot` data source is used to retrieve information about a Block Storage volume snapshot.
 *
 * Refer to the Block Storage [product documentation](https://www.scaleway.com/en/docs/storage/block/) and [API documentation](https://www.scaleway.com/en/developers/api/block/) for more information.
 */
export function getBlockSnapshotOutput(args?: GetBlockSnapshotOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBlockSnapshotResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getBlockSnapshot:getBlockSnapshot", {
        "name": args.name,
        "projectId": args.projectId,
        "snapshotId": args.snapshotId,
        "volumeId": args.volumeId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getBlockSnapshot.
 */
export interface GetBlockSnapshotOutputArgs {
    /**
     * The name of the snapshot. Only one of name or snapshotId should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The unique identifier of the Project to which the snapshot is associated.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The unique identifier of the snapshot. Only one of `name` and `snapshotId` should be specified.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * The unique identifier of the volume from which the snapshot was created.
     */
    volumeId?: pulumi.Input<string>;
    /**
     * ) The zone in which the snapshot exists.
     */
    zone?: pulumi.Input<string>;
}