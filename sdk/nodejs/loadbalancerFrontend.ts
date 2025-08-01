// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Load Balancer frontends.
 *
 * For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/reference-content/configuring-frontends/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const frontend01 = new scaleway.loadbalancers.Frontend("frontend01", {
 *     lbId: lb01.id,
 *     backendId: backend01.id,
 *     name: "frontend01",
 *     inboundPort: 80,
 * });
 * ```
 *
 * ## With ACLs
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const frontend01 = new scaleway.loadbalancers.Frontend("frontend01", {
 *     lbId: lb01.id,
 *     backendId: backend01.id,
 *     name: "frontend01",
 *     inboundPort: 80,
 *     acls: [
 *         {
 *             name: "blacklist wellknwon IPs",
 *             action: {
 *                 type: "allow",
 *             },
 *             match: {
 *                 ipSubnets: [
 *                     "192.168.0.1",
 *                     "192.168.0.2",
 *                     "192.168.10.0/24",
 *                 ],
 *             },
 *         },
 *         {
 *             action: {
 *                 type: "deny",
 *             },
 *             match: {
 *                 ipSubnets: ["51.51.51.51"],
 *                 httpFilter: "regex",
 *                 httpFilterValues: ["^foo*bar$"],
 *             },
 *         },
 *         {
 *             action: {
 *                 type: "allow",
 *             },
 *             match: {
 *                 httpFilter: "path_begin",
 *                 httpFilterValues: [
 *                     "foo",
 *                     "bar",
 *                 ],
 *             },
 *         },
 *         {
 *             action: {
 *                 type: "allow",
 *             },
 *             match: {
 *                 httpFilter: "path_begin",
 *                 httpFilterValues: ["hi"],
 *                 invert: true,
 *             },
 *         },
 *         {
 *             action: {
 *                 type: "allow",
 *             },
 *             match: {
 *                 httpFilter: "http_header_match",
 *                 httpFilterValues: "foo",
 *                 httpFilterOption: "bar",
 *             },
 *         },
 *         {
 *             action: {
 *                 type: "redirect",
 *                 redirects: [{
 *                     type: "location",
 *                     target: "https://example.com",
 *                     code: 307,
 *                 }],
 *             },
 *             match: {
 *                 ipSubnets: ["10.0.0.10"],
 *                 httpFilter: "path_begin",
 *                 httpFilterValues: [
 *                     "foo",
 *                     "bar",
 *                 ],
 *             },
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Load Balancer frontends can be imported using `{zone}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/loadbalancerFrontend:LoadbalancerFrontend frontend01 fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 *
 * @deprecated scaleway.index/loadbalancerfrontend.LoadbalancerFrontend has been deprecated in favor of scaleway.loadbalancers/frontend.Frontend
 */
export class LoadbalancerFrontend extends pulumi.CustomResource {
    /**
     * Get an existing LoadbalancerFrontend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadbalancerFrontendState, opts?: pulumi.CustomResourceOptions): LoadbalancerFrontend {
        pulumi.log.warn("LoadbalancerFrontend is deprecated: scaleway.index/loadbalancerfrontend.LoadbalancerFrontend has been deprecated in favor of scaleway.loadbalancers/frontend.Frontend")
        return new LoadbalancerFrontend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/loadbalancerFrontend:LoadbalancerFrontend';

    /**
     * Returns true if the given object is an instance of LoadbalancerFrontend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadbalancerFrontend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadbalancerFrontend.__pulumiType;
    }

    /**
     * A list of ACL rules to apply to the Load Balancer frontend.  Defined below.
     */
    public readonly acls!: pulumi.Output<outputs.LoadbalancerFrontendAcl[] | undefined>;
    /**
     * The ID of the Load Balancer backend this frontend is attached to.
     *
     * > **Important:** Updates to `lbId` or `backendId` will recreate the frontend.
     */
    public readonly backendId!: pulumi.Output<string>;
    /**
     * (Deprecated, use `certificateIds` instead) First certificate ID used by the frontend.
     *
     * @deprecated Please use certificate_ids
     */
    public /*out*/ readonly certificateId!: pulumi.Output<string>;
    /**
     * List of certificate IDs that should be used by the frontend.
     *
     * > **Important:** Certificates are not allowed on port 80.
     */
    public readonly certificateIds!: pulumi.Output<string[] | undefined>;
    /**
     * The rate limit for new connections established on this frontend. Use 0 value to disable, else value is connections per second.
     */
    public readonly connectionRateLimit!: pulumi.Output<number | undefined>;
    /**
     * The date and time the frontend was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Defines whether to enable access logs on the frontend.
     */
    public readonly enableAccessLogs!: pulumi.Output<boolean | undefined>;
    /**
     * Activates HTTP/3 protocol.
     */
    public readonly enableHttp3!: pulumi.Output<boolean | undefined>;
    /**
     * A boolean to specify whether to use lb_acl.
     * If `externalAcls` is set to `true`, `acl` can not be set directly in the Load Balancer frontend.
     */
    public readonly externalAcls!: pulumi.Output<boolean | undefined>;
    /**
     * TCP port to listen to on the front side.
     */
    public readonly inboundPort!: pulumi.Output<number>;
    /**
     * The ID of the Load Balancer this frontend is attached to.
     */
    public readonly lbId!: pulumi.Output<string>;
    /**
     * The ACL name. If not provided it will be randomly generated.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Maximum inactivity time on the client side. (e.g. `1s`)
     */
    public readonly timeoutClient!: pulumi.Output<string | undefined>;
    /**
     * The date and time the frontend resource was updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a LoadbalancerFrontend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated scaleway.index/loadbalancerfrontend.LoadbalancerFrontend has been deprecated in favor of scaleway.loadbalancers/frontend.Frontend */
    constructor(name: string, args: LoadbalancerFrontendArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated scaleway.index/loadbalancerfrontend.LoadbalancerFrontend has been deprecated in favor of scaleway.loadbalancers/frontend.Frontend */
    constructor(name: string, argsOrState?: LoadbalancerFrontendArgs | LoadbalancerFrontendState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("LoadbalancerFrontend is deprecated: scaleway.index/loadbalancerfrontend.LoadbalancerFrontend has been deprecated in favor of scaleway.loadbalancers/frontend.Frontend")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadbalancerFrontendState | undefined;
            resourceInputs["acls"] = state ? state.acls : undefined;
            resourceInputs["backendId"] = state ? state.backendId : undefined;
            resourceInputs["certificateId"] = state ? state.certificateId : undefined;
            resourceInputs["certificateIds"] = state ? state.certificateIds : undefined;
            resourceInputs["connectionRateLimit"] = state ? state.connectionRateLimit : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["enableAccessLogs"] = state ? state.enableAccessLogs : undefined;
            resourceInputs["enableHttp3"] = state ? state.enableHttp3 : undefined;
            resourceInputs["externalAcls"] = state ? state.externalAcls : undefined;
            resourceInputs["inboundPort"] = state ? state.inboundPort : undefined;
            resourceInputs["lbId"] = state ? state.lbId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["timeoutClient"] = state ? state.timeoutClient : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as LoadbalancerFrontendArgs | undefined;
            if ((!args || args.backendId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backendId'");
            }
            if ((!args || args.inboundPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'inboundPort'");
            }
            if ((!args || args.lbId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lbId'");
            }
            resourceInputs["acls"] = args ? args.acls : undefined;
            resourceInputs["backendId"] = args ? args.backendId : undefined;
            resourceInputs["certificateIds"] = args ? args.certificateIds : undefined;
            resourceInputs["connectionRateLimit"] = args ? args.connectionRateLimit : undefined;
            resourceInputs["enableAccessLogs"] = args ? args.enableAccessLogs : undefined;
            resourceInputs["enableHttp3"] = args ? args.enableHttp3 : undefined;
            resourceInputs["externalAcls"] = args ? args.externalAcls : undefined;
            resourceInputs["inboundPort"] = args ? args.inboundPort : undefined;
            resourceInputs["lbId"] = args ? args.lbId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["timeoutClient"] = args ? args.timeoutClient : undefined;
            resourceInputs["certificateId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadbalancerFrontend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadbalancerFrontend resources.
 */
export interface LoadbalancerFrontendState {
    /**
     * A list of ACL rules to apply to the Load Balancer frontend.  Defined below.
     */
    acls?: pulumi.Input<pulumi.Input<inputs.LoadbalancerFrontendAcl>[]>;
    /**
     * The ID of the Load Balancer backend this frontend is attached to.
     *
     * > **Important:** Updates to `lbId` or `backendId` will recreate the frontend.
     */
    backendId?: pulumi.Input<string>;
    /**
     * (Deprecated, use `certificateIds` instead) First certificate ID used by the frontend.
     *
     * @deprecated Please use certificate_ids
     */
    certificateId?: pulumi.Input<string>;
    /**
     * List of certificate IDs that should be used by the frontend.
     *
     * > **Important:** Certificates are not allowed on port 80.
     */
    certificateIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The rate limit for new connections established on this frontend. Use 0 value to disable, else value is connections per second.
     */
    connectionRateLimit?: pulumi.Input<number>;
    /**
     * The date and time the frontend was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Defines whether to enable access logs on the frontend.
     */
    enableAccessLogs?: pulumi.Input<boolean>;
    /**
     * Activates HTTP/3 protocol.
     */
    enableHttp3?: pulumi.Input<boolean>;
    /**
     * A boolean to specify whether to use lb_acl.
     * If `externalAcls` is set to `true`, `acl` can not be set directly in the Load Balancer frontend.
     */
    externalAcls?: pulumi.Input<boolean>;
    /**
     * TCP port to listen to on the front side.
     */
    inboundPort?: pulumi.Input<number>;
    /**
     * The ID of the Load Balancer this frontend is attached to.
     */
    lbId?: pulumi.Input<string>;
    /**
     * The ACL name. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * Maximum inactivity time on the client side. (e.g. `1s`)
     */
    timeoutClient?: pulumi.Input<string>;
    /**
     * The date and time the frontend resource was updated.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoadbalancerFrontend resource.
 */
export interface LoadbalancerFrontendArgs {
    /**
     * A list of ACL rules to apply to the Load Balancer frontend.  Defined below.
     */
    acls?: pulumi.Input<pulumi.Input<inputs.LoadbalancerFrontendAcl>[]>;
    /**
     * The ID of the Load Balancer backend this frontend is attached to.
     *
     * > **Important:** Updates to `lbId` or `backendId` will recreate the frontend.
     */
    backendId: pulumi.Input<string>;
    /**
     * List of certificate IDs that should be used by the frontend.
     *
     * > **Important:** Certificates are not allowed on port 80.
     */
    certificateIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The rate limit for new connections established on this frontend. Use 0 value to disable, else value is connections per second.
     */
    connectionRateLimit?: pulumi.Input<number>;
    /**
     * Defines whether to enable access logs on the frontend.
     */
    enableAccessLogs?: pulumi.Input<boolean>;
    /**
     * Activates HTTP/3 protocol.
     */
    enableHttp3?: pulumi.Input<boolean>;
    /**
     * A boolean to specify whether to use lb_acl.
     * If `externalAcls` is set to `true`, `acl` can not be set directly in the Load Balancer frontend.
     */
    externalAcls?: pulumi.Input<boolean>;
    /**
     * TCP port to listen to on the front side.
     */
    inboundPort: pulumi.Input<number>;
    /**
     * The ID of the Load Balancer this frontend is attached to.
     */
    lbId: pulumi.Input<string>;
    /**
     * The ACL name. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * Maximum inactivity time on the client side. (e.g. `1s`)
     */
    timeoutClient?: pulumi.Input<string>;
}
