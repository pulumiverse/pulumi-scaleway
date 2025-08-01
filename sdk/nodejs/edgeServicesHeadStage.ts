// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Sets the Scaleway Edge Services head stage of your pipeline.
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const main = new scaleway.EdgeServicesPipeline("main", {
 *     name: "my-edge_services-pipeline",
 *     description: "pipeline description",
 * });
 * const mainEdgeServicesDnsStage = new scaleway.EdgeServicesDnsStage("main", {
 *     pipelineId: main.id,
 *     tlsStageId: mainScalewayEdgeServicesTlsStage.id,
 *     fqdns: ["subdomain.example.com"],
 * });
 * const mainEdgeServicesHeadStage = new scaleway.EdgeServicesHeadStage("main", {
 *     pipelineId: main.id,
 *     headStageId: mainEdgeServicesDnsStage.id,
 * });
 * ```
 *
 * ## Import
 *
 * Head stages can be imported using the `{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/edgeServicesHeadStage:EdgeServicesHeadStage main 11111111-1111-1111-1111-111111111111
 * ```
 */
export class EdgeServicesHeadStage extends pulumi.CustomResource {
    /**
     * Get an existing EdgeServicesHeadStage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EdgeServicesHeadStageState, opts?: pulumi.CustomResourceOptions): EdgeServicesHeadStage {
        return new EdgeServicesHeadStage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/edgeServicesHeadStage:EdgeServicesHeadStage';

    /**
     * Returns true if the given object is an instance of EdgeServicesHeadStage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EdgeServicesHeadStage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EdgeServicesHeadStage.__pulumiType;
    }

    /**
     * The ID of head stage of the pipeline.
     */
    public readonly headStageId!: pulumi.Output<string>;
    /**
     * The ID of the pipeline.
     */
    public readonly pipelineId!: pulumi.Output<string>;

    /**
     * Create a EdgeServicesHeadStage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EdgeServicesHeadStageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EdgeServicesHeadStageArgs | EdgeServicesHeadStageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EdgeServicesHeadStageState | undefined;
            resourceInputs["headStageId"] = state ? state.headStageId : undefined;
            resourceInputs["pipelineId"] = state ? state.pipelineId : undefined;
        } else {
            const args = argsOrState as EdgeServicesHeadStageArgs | undefined;
            if ((!args || args.pipelineId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pipelineId'");
            }
            resourceInputs["headStageId"] = args ? args.headStageId : undefined;
            resourceInputs["pipelineId"] = args ? args.pipelineId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EdgeServicesHeadStage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EdgeServicesHeadStage resources.
 */
export interface EdgeServicesHeadStageState {
    /**
     * The ID of head stage of the pipeline.
     */
    headStageId?: pulumi.Input<string>;
    /**
     * The ID of the pipeline.
     */
    pipelineId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EdgeServicesHeadStage resource.
 */
export interface EdgeServicesHeadStageArgs {
    /**
     * The ID of head stage of the pipeline.
     */
    headStageId?: pulumi.Input<string>;
    /**
     * The ID of the pipeline.
     */
    pipelineId: pulumi.Input<string>;
}
