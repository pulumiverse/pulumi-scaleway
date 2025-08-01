// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * This section explains how to import a zone using the `{subdomain}.{domain}` format.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/domainZone:DomainZone test test.scaleway-terraform.com
 * ```
 *
 * @deprecated scaleway.index/domainzone.DomainZone has been deprecated in favor of scaleway.domain/zone.Zone
 */
export class DomainZone extends pulumi.CustomResource {
    /**
     * Get an existing DomainZone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainZoneState, opts?: pulumi.CustomResourceOptions): DomainZone {
        pulumi.log.warn("DomainZone is deprecated: scaleway.index/domainzone.DomainZone has been deprecated in favor of scaleway.domain/zone.Zone")
        return new DomainZone(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/domainZone:DomainZone';

    /**
     * Returns true if the given object is an instance of DomainZone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainZone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainZone.__pulumiType;
    }

    /**
     * The main domain where the DNS zone will be created.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Message.
     */
    public /*out*/ readonly message!: pulumi.Output<string>;
    /**
     * The list of same servers for the zone.
     */
    public /*out*/ readonly ns!: pulumi.Output<string[]>;
    /**
     * The default list of same servers for the zone.
     */
    public /*out*/ readonly nsDefaults!: pulumi.Output<string[]>;
    /**
     * The master list of same servers for the zone.
     */
    public /*out*/ readonly nsMasters!: pulumi.Output<string[]>;
    /**
     * `projectId`) The ID of the Project associated with the domain.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The status of the domain zone.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The name of the subdomain (zone name) to create within the domain.
     */
    public readonly subdomain!: pulumi.Output<string>;
    /**
     * The date and time at which the DNS zone was last updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a DomainZone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated scaleway.index/domainzone.DomainZone has been deprecated in favor of scaleway.domain/zone.Zone */
    constructor(name: string, args: DomainZoneArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated scaleway.index/domainzone.DomainZone has been deprecated in favor of scaleway.domain/zone.Zone */
    constructor(name: string, argsOrState?: DomainZoneArgs | DomainZoneState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DomainZone is deprecated: scaleway.index/domainzone.DomainZone has been deprecated in favor of scaleway.domain/zone.Zone")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainZoneState | undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["message"] = state ? state.message : undefined;
            resourceInputs["ns"] = state ? state.ns : undefined;
            resourceInputs["nsDefaults"] = state ? state.nsDefaults : undefined;
            resourceInputs["nsMasters"] = state ? state.nsMasters : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subdomain"] = state ? state.subdomain : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as DomainZoneArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.subdomain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subdomain'");
            }
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["subdomain"] = args ? args.subdomain : undefined;
            resourceInputs["message"] = undefined /*out*/;
            resourceInputs["ns"] = undefined /*out*/;
            resourceInputs["nsDefaults"] = undefined /*out*/;
            resourceInputs["nsMasters"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DomainZone.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainZone resources.
 */
export interface DomainZoneState {
    /**
     * The main domain where the DNS zone will be created.
     */
    domain?: pulumi.Input<string>;
    /**
     * Message.
     */
    message?: pulumi.Input<string>;
    /**
     * The list of same servers for the zone.
     */
    ns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The default list of same servers for the zone.
     */
    nsDefaults?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The master list of same servers for the zone.
     */
    nsMasters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `projectId`) The ID of the Project associated with the domain.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The status of the domain zone.
     */
    status?: pulumi.Input<string>;
    /**
     * The name of the subdomain (zone name) to create within the domain.
     */
    subdomain?: pulumi.Input<string>;
    /**
     * The date and time at which the DNS zone was last updated.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DomainZone resource.
 */
export interface DomainZoneArgs {
    /**
     * The main domain where the DNS zone will be created.
     */
    domain: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the Project associated with the domain.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The name of the subdomain (zone name) to create within the domain.
     */
    subdomain: pulumi.Input<string>;
}
