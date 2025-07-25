// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Add members to an IAM group.
 * For more information refer to the [IAM API documentation](https://www.scaleway.com/en/developers/api/iam/#groups-f592eb).
 *
 * ## Example Usage
 *
 * ### Application Membership
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const group = new scaleway.iam.Group("group", {
 *     name: "my_group",
 *     externalMembership: true,
 * });
 * const app = new scaleway.iam.Application("app", {name: "my_app"});
 * const member = new scaleway.iam.GroupMembership("member", {
 *     groupId: group.id,
 *     applicationId: app.id,
 * });
 * ```
 *
 * ## Import
 *
 * IAM group memberships can be imported using two format:
 *
 * - For user: `{group_id}/user/{user_id}`
 *
 * - For application: `{group_id}/app/{application_id}`
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:iam/groupMembership:GroupMembership app 11111111-1111-1111-1111-111111111111/app/11111111-1111-1111-1111-111111111111
 * ```
 */
export class GroupMembership extends pulumi.CustomResource {
    /**
     * Get an existing GroupMembership resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupMembershipState, opts?: pulumi.CustomResourceOptions): GroupMembership {
        return new GroupMembership(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:iam/groupMembership:GroupMembership';

    /**
     * Returns true if the given object is an instance of GroupMembership.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupMembership {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupMembership.__pulumiType;
    }

    /**
     * The ID of the application that will be added to the group.
     */
    public readonly applicationId!: pulumi.Output<string | undefined>;
    /**
     * ID of the group to add members to.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * The ID of the user that will be added to the group
     *
     * > **Note** You must specify at least one: `applicationId` and/or `userId`.
     */
    public readonly userId!: pulumi.Output<string | undefined>;

    /**
     * Create a GroupMembership resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupMembershipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupMembershipArgs | GroupMembershipState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupMembershipState | undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as GroupMembershipArgs | undefined;
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "scaleway:index/iamGroupMembership:IamGroupMembership" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(GroupMembership.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupMembership resources.
 */
export interface GroupMembershipState {
    /**
     * The ID of the application that will be added to the group.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * ID of the group to add members to.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The ID of the user that will be added to the group
     *
     * > **Note** You must specify at least one: `applicationId` and/or `userId`.
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupMembership resource.
 */
export interface GroupMembershipArgs {
    /**
     * The ID of the application that will be added to the group.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * ID of the group to add members to.
     */
    groupId: pulumi.Input<string>;
    /**
     * The ID of the user that will be added to the group
     *
     * > **Note** You must specify at least one: `applicationId` and/or `userId`.
     */
    userId?: pulumi.Input<string>;
}
