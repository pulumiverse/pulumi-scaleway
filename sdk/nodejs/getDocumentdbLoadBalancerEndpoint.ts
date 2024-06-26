// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about an DocumentDB load balancer endpoint.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myEndpoint = scaleway.getDocumentdbLoadBalancerEndpoint({
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getDocumentdbLoadBalancerEndpoint(args?: GetDocumentdbLoadBalancerEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetDocumentdbLoadBalancerEndpointResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getDocumentdbLoadBalancerEndpoint:getDocumentdbLoadBalancerEndpoint", {
        "instanceId": args.instanceId,
        "instanceName": args.instanceName,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getDocumentdbLoadBalancerEndpoint.
 */
export interface GetDocumentdbLoadBalancerEndpointArgs {
    /**
     * The DocumentDB Instance on which the endpoint is attached. Only one of `instanceName` and `instanceId` should be specified.
     */
    instanceId?: string;
    /**
     * The DocumentDB Instance Name on which the endpoint is attached. Only one of `instanceName` and `instanceId` should be specified.
     */
    instanceName?: string;
    /**
     * The ID of the project the DocumentDB endpoint is associated with.
     */
    projectId?: string;
    /**
     * `region`) The region in which the DocumentDB endpoint exists.
     */
    region?: string;
}

/**
 * A collection of values returned by getDocumentdbLoadBalancerEndpoint.
 */
export interface GetDocumentdbLoadBalancerEndpointResult {
    /**
     * The hostname of your endpoint.
     */
    readonly hostname: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly instanceName: string;
    /**
     * The IP of your load balancer service.
     */
    readonly ip: string;
    /**
     * The name of your load balancer service.
     */
    readonly name: string;
    /**
     * The port of your load balancer service.
     */
    readonly port: number;
    readonly projectId: string;
    readonly region: string;
}
/**
 * Gets information about an DocumentDB load balancer endpoint.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myEndpoint = scaleway.getDocumentdbLoadBalancerEndpoint({
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getDocumentdbLoadBalancerEndpointOutput(args?: GetDocumentdbLoadBalancerEndpointOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDocumentdbLoadBalancerEndpointResult> {
    return pulumi.output(args).apply((a: any) => getDocumentdbLoadBalancerEndpoint(a, opts))
}

/**
 * A collection of arguments for invoking getDocumentdbLoadBalancerEndpoint.
 */
export interface GetDocumentdbLoadBalancerEndpointOutputArgs {
    /**
     * The DocumentDB Instance on which the endpoint is attached. Only one of `instanceName` and `instanceId` should be specified.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The DocumentDB Instance Name on which the endpoint is attached. Only one of `instanceName` and `instanceId` should be specified.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The ID of the project the DocumentDB endpoint is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the DocumentDB endpoint exists.
     */
    region?: pulumi.Input<string>;
}
