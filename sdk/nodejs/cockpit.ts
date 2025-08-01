// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * This section explains how to import a Cockpit using its `{project_id}`.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/cockpit:Cockpit main 11111111-1111-1111-1111-111111111111
 * ```
 *
 * @deprecated scaleway.index/cockpit.Cockpit has been deprecated in favor of scaleway.observability/cockpit.Cockpit
 */
export class Cockpit extends pulumi.CustomResource {
    /**
     * Get an existing Cockpit resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CockpitState, opts?: pulumi.CustomResourceOptions): Cockpit {
        pulumi.log.warn("Cockpit is deprecated: scaleway.index/cockpit.Cockpit has been deprecated in favor of scaleway.observability/cockpit.Cockpit")
        return new Cockpit(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/cockpit:Cockpit';

    /**
     * Returns true if the given object is an instance of Cockpit.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cockpit {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cockpit.__pulumiType;
    }

    /**
     * (Deprecated) A list of [endpoints](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#endpoints) related to Cockpit, each with specific URLs:
     *
     * @deprecated Use 'scaleway_cockpit_source' instead of 'endpoints'. This field will be removed in future releases.
     */
    public /*out*/ readonly endpoints!: pulumi.Output<outputs.CockpitEndpoint[]>;
    /**
     * Name of the plan to use. Available plans are: free, premium, and custom.
     * > **Important:** The plan field is deprecated. Any modification or selection will have no effect.
     *
     * @deprecated The 'plan' attribute is deprecated and no longer has any effect. Future updates will remove this attribute entirely.
     */
    public readonly plan!: pulumi.Output<string | undefined>;
    /**
     * (Deprecated) The ID of the current pricing plan.
     *
     * @deprecated The 'plan_id' attribute is deprecated and will be removed in a future release.
     */
    public /*out*/ readonly planId!: pulumi.Output<string>;
    /**
     * ) The ID of the Project the Cockpit is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * [DEPRECATED] Push_url
     *
     * @deprecated Please use `scaleway.observability.Source` instead
     */
    public /*out*/ readonly pushUrls!: pulumi.Output<outputs.CockpitPushUrl[]>;

    /**
     * Create a Cockpit resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated scaleway.index/cockpit.Cockpit has been deprecated in favor of scaleway.observability/cockpit.Cockpit */
    constructor(name: string, args?: CockpitArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated scaleway.index/cockpit.Cockpit has been deprecated in favor of scaleway.observability/cockpit.Cockpit */
    constructor(name: string, argsOrState?: CockpitArgs | CockpitState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Cockpit is deprecated: scaleway.index/cockpit.Cockpit has been deprecated in favor of scaleway.observability/cockpit.Cockpit")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CockpitState | undefined;
            resourceInputs["endpoints"] = state ? state.endpoints : undefined;
            resourceInputs["plan"] = state ? state.plan : undefined;
            resourceInputs["planId"] = state ? state.planId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["pushUrls"] = state ? state.pushUrls : undefined;
        } else {
            const args = argsOrState as CockpitArgs | undefined;
            resourceInputs["plan"] = args ? args.plan : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["endpoints"] = undefined /*out*/;
            resourceInputs["planId"] = undefined /*out*/;
            resourceInputs["pushUrls"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cockpit.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cockpit resources.
 */
export interface CockpitState {
    /**
     * (Deprecated) A list of [endpoints](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#endpoints) related to Cockpit, each with specific URLs:
     *
     * @deprecated Use 'scaleway_cockpit_source' instead of 'endpoints'. This field will be removed in future releases.
     */
    endpoints?: pulumi.Input<pulumi.Input<inputs.CockpitEndpoint>[]>;
    /**
     * Name of the plan to use. Available plans are: free, premium, and custom.
     * > **Important:** The plan field is deprecated. Any modification or selection will have no effect.
     *
     * @deprecated The 'plan' attribute is deprecated and no longer has any effect. Future updates will remove this attribute entirely.
     */
    plan?: pulumi.Input<string>;
    /**
     * (Deprecated) The ID of the current pricing plan.
     *
     * @deprecated The 'plan_id' attribute is deprecated and will be removed in a future release.
     */
    planId?: pulumi.Input<string>;
    /**
     * ) The ID of the Project the Cockpit is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * [DEPRECATED] Push_url
     *
     * @deprecated Please use `scaleway.observability.Source` instead
     */
    pushUrls?: pulumi.Input<pulumi.Input<inputs.CockpitPushUrl>[]>;
}

/**
 * The set of arguments for constructing a Cockpit resource.
 */
export interface CockpitArgs {
    /**
     * Name of the plan to use. Available plans are: free, premium, and custom.
     * > **Important:** The plan field is deprecated. Any modification or selection will have no effect.
     *
     * @deprecated The 'plan' attribute is deprecated and no longer has any effect. Future updates will remove this attribute entirely.
     */
    plan?: pulumi.Input<string>;
    /**
     * ) The ID of the Project the Cockpit is associated with.
     */
    projectId?: pulumi.Input<string>;
}
