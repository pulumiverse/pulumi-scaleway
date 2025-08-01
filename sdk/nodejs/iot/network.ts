// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const mainHub = new scaleway.iot.Hub("main", {
 *     name: "main",
 *     productPlan: "plan_shared",
 * });
 * const main = new scaleway.iot.Network("main", {
 *     name: "main",
 *     hubId: mainHub.id,
 *     type: "sigfox",
 * });
 * ```
 *
 * ## Import
 *
 * IoT Networks can be imported using the `{region}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:iot/network:Network net01 fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class Network extends pulumi.CustomResource {
    /**
     * Get an existing Network resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkState, opts?: pulumi.CustomResourceOptions): Network {
        return new Network(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:iot/network:Network';

    /**
     * Returns true if the given object is an instance of Network.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Network {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Network.__pulumiType;
    }

    /**
     * The date and time the Network was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The endpoint to use when interacting with the network.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * The hub ID to which the Network will be attached to.
     */
    public readonly hubId!: pulumi.Output<string>;
    /**
     * The name of the IoT Network you want to create (e.g. `my-net`).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * (Defaults to provider `region`) The region in which the Network is attached to.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The endpoint key to keep secret.
     */
    public /*out*/ readonly secret!: pulumi.Output<string>;
    /**
     * The prefix that will be prepended to all topics for this Network.
     */
    public readonly topicPrefix!: pulumi.Output<string | undefined>;
    /**
     * The network type to create (e.g. `sigfox`).
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Network resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkArgs | NetworkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["hubId"] = state ? state.hubId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["secret"] = state ? state.secret : undefined;
            resourceInputs["topicPrefix"] = state ? state.topicPrefix : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as NetworkArgs | undefined;
            if ((!args || args.hubId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hubId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["hubId"] = args ? args.hubId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["topicPrefix"] = args ? args.topicPrefix : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["secret"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "scaleway:index/iotNetwork:IotNetwork" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        const secretOpts = { additionalSecretOutputs: ["secret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Network.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Network resources.
 */
export interface NetworkState {
    /**
     * The date and time the Network was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The endpoint to use when interacting with the network.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * The hub ID to which the Network will be attached to.
     */
    hubId?: pulumi.Input<string>;
    /**
     * The name of the IoT Network you want to create (e.g. `my-net`).
     */
    name?: pulumi.Input<string>;
    /**
     * (Defaults to provider `region`) The region in which the Network is attached to.
     */
    region?: pulumi.Input<string>;
    /**
     * The endpoint key to keep secret.
     */
    secret?: pulumi.Input<string>;
    /**
     * The prefix that will be prepended to all topics for this Network.
     */
    topicPrefix?: pulumi.Input<string>;
    /**
     * The network type to create (e.g. `sigfox`).
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Network resource.
 */
export interface NetworkArgs {
    /**
     * The hub ID to which the Network will be attached to.
     */
    hubId: pulumi.Input<string>;
    /**
     * The name of the IoT Network you want to create (e.g. `my-net`).
     */
    name?: pulumi.Input<string>;
    /**
     * (Defaults to provider `region`) The region in which the Network is attached to.
     */
    region?: pulumi.Input<string>;
    /**
     * The prefix that will be prepended to all topics for this Network.
     */
    topicPrefix?: pulumi.Input<string>;
    /**
     * The network type to create (e.g. `sigfox`).
     */
    type: pulumi.Input<string>;
}
