// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about an instance volume.
 */
/** @deprecated scaleway.index/getinstancevolume.getInstanceVolume has been deprecated in favor of scaleway.instance/getvolume.getVolume */
export function getInstanceVolume(args?: GetInstanceVolumeArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceVolumeResult> {
    pulumi.log.warn("getInstanceVolume is deprecated: scaleway.index/getinstancevolume.getInstanceVolume has been deprecated in favor of scaleway.instance/getvolume.getVolume")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getInstanceVolume:getInstanceVolume", {
        "name": args.name,
        "projectId": args.projectId,
        "volumeId": args.volumeId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceVolume.
 */
export interface GetInstanceVolumeArgs {
    /**
     * The volume name.
     * Only one of `name` and `volumeId` should be specified.
     */
    name?: string;
    /**
     * The ID of the project the volume is associated with.
     */
    projectId?: string;
    /**
     * The volume id.
     * Only one of `name` and `volumeId` should be specified.
     */
    volumeId?: string;
    /**
     * `zone`) The zone in which the volume exists.
     */
    zone?: string;
}

/**
 * A collection of values returned by getInstanceVolume.
 */
export interface GetInstanceVolumeResult {
    readonly fromSnapshotId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly migrateToSbs: boolean;
    readonly name?: string;
    /**
     * The ID of the organization the volume is associated with.
     */
    readonly organizationId: string;
    readonly projectId?: string;
    readonly serverId: string;
    readonly sizeInGb: number;
    readonly tags: string[];
    readonly type: string;
    readonly volumeId?: string;
    readonly zone?: string;
}
/**
 * Gets information about an instance volume.
 */
/** @deprecated scaleway.index/getinstancevolume.getInstanceVolume has been deprecated in favor of scaleway.instance/getvolume.getVolume */
export function getInstanceVolumeOutput(args?: GetInstanceVolumeOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetInstanceVolumeResult> {
    pulumi.log.warn("getInstanceVolume is deprecated: scaleway.index/getinstancevolume.getInstanceVolume has been deprecated in favor of scaleway.instance/getvolume.getVolume")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getInstanceVolume:getInstanceVolume", {
        "name": args.name,
        "projectId": args.projectId,
        "volumeId": args.volumeId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceVolume.
 */
export interface GetInstanceVolumeOutputArgs {
    /**
     * The volume name.
     * Only one of `name` and `volumeId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project the volume is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The volume id.
     * Only one of `name` and `volumeId` should be specified.
     */
    volumeId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the volume exists.
     */
    zone?: pulumi.Input<string>;
}
