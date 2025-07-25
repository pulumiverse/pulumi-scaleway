// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > **Important:**  The data source `scaleway.observability.Cockpit` has been deprecated and will no longer be supported. Instead, use resource `scaleway.observability.Cockpit`.
 *
 * > **Note:**
 * As of April 2024, Cockpit has introduced [regionalization](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#region) to offer more flexibility and resilience.
 * If you have created customized dashboards with data for your Scaleway resources before April 2024, you will need to update your queries in Grafana, with the new regionalized data sources.
 *
 * The `scaleway.observability.Cockpit` data source is used to retrieve information about a Scaleway Cockpit associated with a given Project. This can be the default Project or a specific Project identified by its ID.
 *
 * Refer to Cockpit's [product documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/) and [API documentation](https://www.scaleway.com/en/developers/api/cockpit/regional-api) for more information.
 *
 * ## Retrieve a Cockpit
 *
 * The following commands allow you to:
 *
 * - get information on the Cockpit associated with your Scaleway default Project
 * - get information on the Cockpit associated with a specific Scaleway Project
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * // Get the default Project's Cockpit
 * const main = scaleway.observability.getInstance({});
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * // Get a specific Project's Cockpit
 * const main = scaleway.observability.getInstance({
 *     projectId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getInstance(args?: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:observability/getInstance:getInstance", {
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceArgs {
    /**
     * Specifies the ID of the Scaleway Project that the Cockpit is associated with. If not specified, it defaults to the Project ID specified in the provider configuration.
     */
    projectId?: string;
}

/**
 * A collection of values returned by getInstance.
 */
export interface GetInstanceResult {
    /**
     * (Deprecated) A list of [endpoints](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#endpoints) related to Cockpit, each with specific URLs:
     */
    readonly endpoints: outputs.observability.GetInstanceEndpoint[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * @deprecated The 'plan' attribute is deprecated and will be removed in a future version. Any changes to this attribute will have no effect.
     */
    readonly plan: string;
    /**
     * (Deprecated) ID of the current pricing plan
     */
    readonly planId: string;
    readonly projectId?: string;
    readonly pushUrls: outputs.observability.GetInstancePushUrl[];
}
/**
 * > **Important:**  The data source `scaleway.observability.Cockpit` has been deprecated and will no longer be supported. Instead, use resource `scaleway.observability.Cockpit`.
 *
 * > **Note:**
 * As of April 2024, Cockpit has introduced [regionalization](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#region) to offer more flexibility and resilience.
 * If you have created customized dashboards with data for your Scaleway resources before April 2024, you will need to update your queries in Grafana, with the new regionalized data sources.
 *
 * The `scaleway.observability.Cockpit` data source is used to retrieve information about a Scaleway Cockpit associated with a given Project. This can be the default Project or a specific Project identified by its ID.
 *
 * Refer to Cockpit's [product documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/) and [API documentation](https://www.scaleway.com/en/developers/api/cockpit/regional-api) for more information.
 *
 * ## Retrieve a Cockpit
 *
 * The following commands allow you to:
 *
 * - get information on the Cockpit associated with your Scaleway default Project
 * - get information on the Cockpit associated with a specific Scaleway Project
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * // Get the default Project's Cockpit
 * const main = scaleway.observability.getInstance({});
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * // Get a specific Project's Cockpit
 * const main = scaleway.observability.getInstance({
 *     projectId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getInstanceOutput(args?: GetInstanceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetInstanceResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:observability/getInstance:getInstance", {
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceOutputArgs {
    /**
     * Specifies the ID of the Scaleway Project that the Cockpit is associated with. If not specified, it defaults to the Project ID specified in the provider configuration.
     */
    projectId?: pulumi.Input<string>;
}
