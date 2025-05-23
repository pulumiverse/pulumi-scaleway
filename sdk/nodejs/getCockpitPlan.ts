// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/** @deprecated scaleway.index/getcockpitplan.getCockpitPlan has been deprecated in favor of scaleway.observability/getplan.getPlan */
export function getCockpitPlan(args: GetCockpitPlanArgs, opts?: pulumi.InvokeOptions): Promise<GetCockpitPlanResult> {
    pulumi.log.warn("getCockpitPlan is deprecated: scaleway.index/getcockpitplan.getCockpitPlan has been deprecated in favor of scaleway.observability/getplan.getPlan")
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getCockpitPlan:getCockpitPlan", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCockpitPlan.
 */
export interface GetCockpitPlanArgs {
    /**
     * Name of the pricing plan you want to retrieve information about.
     *
     * @deprecated The 'plan' attribute is deprecated and no longer has any effect. Future updates will remove this attribute entirely.
     */
    name: string;
}

/**
 * A collection of values returned by getCockpitPlan.
 */
export interface GetCockpitPlanResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * @deprecated The 'plan' attribute is deprecated and no longer has any effect. Future updates will remove this attribute entirely.
     */
    readonly name: string;
}
/** @deprecated scaleway.index/getcockpitplan.getCockpitPlan has been deprecated in favor of scaleway.observability/getplan.getPlan */
export function getCockpitPlanOutput(args: GetCockpitPlanOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCockpitPlanResult> {
    pulumi.log.warn("getCockpitPlan is deprecated: scaleway.index/getcockpitplan.getCockpitPlan has been deprecated in favor of scaleway.observability/getplan.getPlan")
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getCockpitPlan:getCockpitPlan", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCockpitPlan.
 */
export interface GetCockpitPlanOutputArgs {
    /**
     * Name of the pricing plan you want to retrieve information about.
     *
     * @deprecated The 'plan' attribute is deprecated and no longer has any effect. Future updates will remove this attribute entirely.
     */
    name: pulumi.Input<string>;
}
