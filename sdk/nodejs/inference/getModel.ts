// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The `scaleway.inference.Model` data source allows you to retrieve information about an inference model available in the Scaleway Inference API, either by providing the model's `name` or its `modelId`.
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const myModel = scaleway.inference.getModel({
 *     name: "meta/llama-3.1-8b-instruct:fp8",
 * });
 * ```
 */
export function getModel(args?: GetModelArgs, opts?: pulumi.InvokeOptions): Promise<GetModelResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:inference/getModel:getModel", {
        "modelId": args.modelId,
        "name": args.name,
        "url": args.url,
    }, opts);
}

/**
 * A collection of arguments for invoking getModel.
 */
export interface GetModelArgs {
    /**
     * The ID of the model to retrieve. Must be a valid UUID with locality (i.e., Scaleway's zoned UUID format).
     */
    modelId?: string;
    /**
     * The fully qualified name of the model to look up (e.g., "meta/llama-3.1-8b-instruct:fp8"). The provider will search for a model with an exact name match in the selected region and project.
     */
    name?: string;
    url?: string;
}

/**
 * A collection of values returned by getModel.
 */
export interface GetModelResult {
    readonly createdAt: string;
    /**
     * A textual description of the model (if available).
     */
    readonly description: string;
    /**
     * Whether the model requires end-user license agreement acceptance before use.
     */
    readonly hasEula: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly modelId?: string;
    readonly name?: string;
    /**
     * List of supported node types and their quantization options. Each entry contains:
     */
    readonly nodesSupports: outputs.inference.GetModelNodesSupport[];
    /**
     * Size, in bits, of the model parameters.
     */
    readonly parameterSizeBits: number;
    readonly projectId: string;
    readonly region: string;
    readonly secret: string;
    /**
     * Total size, in bytes, of the model archive.
     */
    readonly sizeBytes: number;
    /**
     * The current status of the model (e.g., ready, error, etc.).
     */
    readonly status: string;
    /**
     * Tags associated with the model.
     */
    readonly tags: string[];
    readonly updatedAt: string;
    readonly url?: string;
}
/**
 * The `scaleway.inference.Model` data source allows you to retrieve information about an inference model available in the Scaleway Inference API, either by providing the model's `name` or its `modelId`.
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const myModel = scaleway.inference.getModel({
 *     name: "meta/llama-3.1-8b-instruct:fp8",
 * });
 * ```
 */
export function getModelOutput(args?: GetModelOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetModelResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:inference/getModel:getModel", {
        "modelId": args.modelId,
        "name": args.name,
        "url": args.url,
    }, opts);
}

/**
 * A collection of arguments for invoking getModel.
 */
export interface GetModelOutputArgs {
    /**
     * The ID of the model to retrieve. Must be a valid UUID with locality (i.e., Scaleway's zoned UUID format).
     */
    modelId?: pulumi.Input<string>;
    /**
     * The fully qualified name of the model to look up (e.g., "meta/llama-3.1-8b-instruct:fp8"). The provider will search for a model with an exact name match in the selected region and project.
     */
    name?: pulumi.Input<string>;
    url?: pulumi.Input<string>;
}
