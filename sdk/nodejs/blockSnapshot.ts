// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `scaleway.block.Snapshot` resource is used to create and manage snapshots of Block Storage volumes.
 *
 * Refer to the Block Storage [product documentation](https://www.scaleway.com/en/docs/block-storage/) and [API documentation](https://www.scaleway.com/en/developers/api/block/) for more information.
 *
 * ## Example Usage
 *
 * ### Create a snapshot of a Block Storage volume
 *
 * The following command allows you to create a snapshot (`some-snapshot-name`) from a Block Storage volume specified by its ID.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const blockVolume = new scaleway.block.Volume("block_volume", {
 *     iops: 5000,
 *     name: "some-volume-name",
 *     sizeInGb: 20,
 * });
 * const blockSnapshot = new scaleway.block.Snapshot("block_snapshot", {
 *     name: "some-snapshot-name",
 *     volumeId: blockVolume.id,
 * });
 * ```
 *
 * ## Import
 *
 * This section explains how to import the snapshot of a Block Storage volume using the zoned ID format (`{zone}/{id}`).
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/blockSnapshot:BlockSnapshot main fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 *
 * @deprecated scaleway.index/blocksnapshot.BlockSnapshot has been deprecated in favor of scaleway.block/snapshot.Snapshot
 */
export class BlockSnapshot extends pulumi.CustomResource {
    /**
     * Get an existing BlockSnapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BlockSnapshotState, opts?: pulumi.CustomResourceOptions): BlockSnapshot {
        pulumi.log.warn("BlockSnapshot is deprecated: scaleway.index/blocksnapshot.BlockSnapshot has been deprecated in favor of scaleway.block/snapshot.Snapshot")
        return new BlockSnapshot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/blockSnapshot:BlockSnapshot';

    /**
     * Returns true if the given object is an instance of BlockSnapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BlockSnapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BlockSnapshot.__pulumiType;
    }

    /**
     * Use this block to export the volume as a QCOW file to Object Storage.
     */
    public readonly export!: pulumi.Output<outputs.BlockSnapshotExport | undefined>;
    /**
     * Use this block to import a QCOW image from Object Storage to create a volume.
     */
    public readonly import!: pulumi.Output<outputs.BlockSnapshotImport | undefined>;
    /**
     * The name of the snapshot. If not provided, a name will be randomly generated.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ). The ID of the Scaleway Project the snapshot is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * A list of tags to apply to the snapshot.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The ID of the volume to take a snapshot from.
     */
    public readonly volumeId!: pulumi.Output<string | undefined>;
    /**
     * ). The zone in which the snapshot should be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a BlockSnapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated scaleway.index/blocksnapshot.BlockSnapshot has been deprecated in favor of scaleway.block/snapshot.Snapshot */
    constructor(name: string, args?: BlockSnapshotArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated scaleway.index/blocksnapshot.BlockSnapshot has been deprecated in favor of scaleway.block/snapshot.Snapshot */
    constructor(name: string, argsOrState?: BlockSnapshotArgs | BlockSnapshotState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("BlockSnapshot is deprecated: scaleway.index/blocksnapshot.BlockSnapshot has been deprecated in favor of scaleway.block/snapshot.Snapshot")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BlockSnapshotState | undefined;
            resourceInputs["export"] = state ? state.export : undefined;
            resourceInputs["import"] = state ? state.import : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["volumeId"] = state ? state.volumeId : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as BlockSnapshotArgs | undefined;
            resourceInputs["export"] = args ? args.export : undefined;
            resourceInputs["import"] = args ? args.import : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["volumeId"] = args ? args.volumeId : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BlockSnapshot.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BlockSnapshot resources.
 */
export interface BlockSnapshotState {
    /**
     * Use this block to export the volume as a QCOW file to Object Storage.
     */
    export?: pulumi.Input<inputs.BlockSnapshotExport>;
    /**
     * Use this block to import a QCOW image from Object Storage to create a volume.
     */
    import?: pulumi.Input<inputs.BlockSnapshotImport>;
    /**
     * The name of the snapshot. If not provided, a name will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * ). The ID of the Scaleway Project the snapshot is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * A list of tags to apply to the snapshot.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the volume to take a snapshot from.
     */
    volumeId?: pulumi.Input<string>;
    /**
     * ). The zone in which the snapshot should be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BlockSnapshot resource.
 */
export interface BlockSnapshotArgs {
    /**
     * Use this block to export the volume as a QCOW file to Object Storage.
     */
    export?: pulumi.Input<inputs.BlockSnapshotExport>;
    /**
     * Use this block to import a QCOW image from Object Storage to create a volume.
     */
    import?: pulumi.Input<inputs.BlockSnapshotImport>;
    /**
     * The name of the snapshot. If not provided, a name will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * ). The ID of the Scaleway Project the snapshot is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * A list of tags to apply to the snapshot.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the volume to take a snapshot from.
     */
    volumeId?: pulumi.Input<string>;
    /**
     * ). The zone in which the snapshot should be created.
     */
    zone?: pulumi.Input<string>;
}
