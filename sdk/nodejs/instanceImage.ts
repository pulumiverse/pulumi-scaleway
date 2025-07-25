// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Compute Images.
 * For more information, see the [API documentation](https://www.scaleway.com/en/developers/api/instance/#path-images-list-instance-images).
 *
 * ## Example Usage
 *
 * ### From a volume
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const volume = new scaleway.instance.Volume("volume", {
 *     type: "b_ssd",
 *     sizeInGb: 20,
 * });
 * const volumeSnapshot = new scaleway.instance.Snapshot("volume_snapshot", {volumeId: volume.id});
 * const volumeImage = new scaleway.instance.Image("volume_image", {
 *     name: "image_from_volume",
 *     rootVolumeId: volumeSnapshot.id,
 * });
 * ```
 *
 * ### From a server
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const server = new scaleway.instance.Server("server", {
 *     image: "ubuntu_jammy",
 *     type: "DEV1-S",
 * });
 * const serverSnapshot = new scaleway.instance.Snapshot("server_snapshot", {volumeId: main.rootVolume[0].volumeId});
 * const serverImage = new scaleway.instance.Image("server_image", {
 *     name: "image_from_server",
 *     rootVolumeId: serverSnapshot.id,
 * });
 * ```
 *
 * ## Import
 *
 * Images can be imported using the `{zone}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/instanceImage:InstanceImage main fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 *
 * @deprecated scaleway.index/instanceimage.InstanceImage has been deprecated in favor of scaleway.instance/image.Image
 */
export class InstanceImage extends pulumi.CustomResource {
    /**
     * Get an existing InstanceImage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceImageState, opts?: pulumi.CustomResourceOptions): InstanceImage {
        pulumi.log.warn("InstanceImage is deprecated: scaleway.index/instanceimage.InstanceImage has been deprecated in favor of scaleway.instance/image.Image")
        return new InstanceImage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/instanceImage:InstanceImage';

    /**
     * Returns true if the given object is an instance of InstanceImage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceImage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceImage.__pulumiType;
    }

    /**
     * List of IDs of the snapshots of the additional volumes to be attached to the image.
     *
     * > **Important:** For now it is only possible to have 1 additional_volume.
     */
    public readonly additionalVolumeIds!: pulumi.Output<string | undefined>;
    /**
     * The description of the extra volumes attached to the image.
     */
    public /*out*/ readonly additionalVolumes!: pulumi.Output<outputs.InstanceImageAdditionalVolume[]>;
    /**
     * The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
     */
    public readonly architecture!: pulumi.Output<string | undefined>;
    /**
     * Date of the volume creation.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * ID of the server the image is based on (in case it is a backup).
     */
    public /*out*/ readonly fromServerId!: pulumi.Output<string>;
    /**
     * Date of volume latest update.
     */
    public /*out*/ readonly modificationDate!: pulumi.Output<string>;
    /**
     * The name of the image. If not provided it will be randomly generated.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The organization ID the image is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * The ID of the project the image is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Set to `true` if the image is public.
     */
    public readonly public!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the snapshot of the volume to be used as root in the image.
     */
    public readonly rootVolumeId!: pulumi.Output<string>;
    /**
     * State of the volume.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * A list of tags to apply to the image.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The zone in which the image should be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstanceImage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated scaleway.index/instanceimage.InstanceImage has been deprecated in favor of scaleway.instance/image.Image */
    constructor(name: string, args: InstanceImageArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated scaleway.index/instanceimage.InstanceImage has been deprecated in favor of scaleway.instance/image.Image */
    constructor(name: string, argsOrState?: InstanceImageArgs | InstanceImageState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("InstanceImage is deprecated: scaleway.index/instanceimage.InstanceImage has been deprecated in favor of scaleway.instance/image.Image")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceImageState | undefined;
            resourceInputs["additionalVolumeIds"] = state ? state.additionalVolumeIds : undefined;
            resourceInputs["additionalVolumes"] = state ? state.additionalVolumes : undefined;
            resourceInputs["architecture"] = state ? state.architecture : undefined;
            resourceInputs["creationDate"] = state ? state.creationDate : undefined;
            resourceInputs["fromServerId"] = state ? state.fromServerId : undefined;
            resourceInputs["modificationDate"] = state ? state.modificationDate : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["public"] = state ? state.public : undefined;
            resourceInputs["rootVolumeId"] = state ? state.rootVolumeId : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceImageArgs | undefined;
            if ((!args || args.rootVolumeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rootVolumeId'");
            }
            resourceInputs["additionalVolumeIds"] = args ? args.additionalVolumeIds : undefined;
            resourceInputs["architecture"] = args ? args.architecture : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["public"] = args ? args.public : undefined;
            resourceInputs["rootVolumeId"] = args ? args.rootVolumeId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["additionalVolumes"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["fromServerId"] = undefined /*out*/;
            resourceInputs["modificationDate"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceImage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceImage resources.
 */
export interface InstanceImageState {
    /**
     * List of IDs of the snapshots of the additional volumes to be attached to the image.
     *
     * > **Important:** For now it is only possible to have 1 additional_volume.
     */
    additionalVolumeIds?: pulumi.Input<string>;
    /**
     * The description of the extra volumes attached to the image.
     */
    additionalVolumes?: pulumi.Input<pulumi.Input<inputs.InstanceImageAdditionalVolume>[]>;
    /**
     * The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
     */
    architecture?: pulumi.Input<string>;
    /**
     * Date of the volume creation.
     */
    creationDate?: pulumi.Input<string>;
    /**
     * ID of the server the image is based on (in case it is a backup).
     */
    fromServerId?: pulumi.Input<string>;
    /**
     * Date of volume latest update.
     */
    modificationDate?: pulumi.Input<string>;
    /**
     * The name of the image. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * The organization ID the image is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * The ID of the project the image is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Set to `true` if the image is public.
     */
    public?: pulumi.Input<boolean>;
    /**
     * The ID of the snapshot of the volume to be used as root in the image.
     */
    rootVolumeId?: pulumi.Input<string>;
    /**
     * State of the volume.
     */
    state?: pulumi.Input<string>;
    /**
     * A list of tags to apply to the image.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The zone in which the image should be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceImage resource.
 */
export interface InstanceImageArgs {
    /**
     * List of IDs of the snapshots of the additional volumes to be attached to the image.
     *
     * > **Important:** For now it is only possible to have 1 additional_volume.
     */
    additionalVolumeIds?: pulumi.Input<string>;
    /**
     * The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
     */
    architecture?: pulumi.Input<string>;
    /**
     * The name of the image. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project the image is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Set to `true` if the image is public.
     */
    public?: pulumi.Input<boolean>;
    /**
     * The ID of the snapshot of the volume to be used as root in the image.
     */
    rootVolumeId: pulumi.Input<string>;
    /**
     * A list of tags to apply to the image.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The zone in which the image should be created.
     */
    zone?: pulumi.Input<string>;
}
