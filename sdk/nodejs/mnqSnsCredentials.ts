// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Messaging and Queuing SNS credentials.
 * For further information, see
 * our [main documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sns-overview/)
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const main = new scaleway.MnqSns("main", {});
 * const mainMnqSnsCredentials = new scaleway.MnqSnsCredentials("main", {
 *     projectId: main.projectId,
 *     name: "sns-credentials",
 *     permissions: {
 *         canManage: false,
 *         canReceive: true,
 *         canPublish: false,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * SNS credentials can be imported using `{region}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/mnqSnsCredentials:MnqSnsCredentials main fr-par/11111111111111111111111111111111
 * ```
 */
export class MnqSnsCredentials extends pulumi.CustomResource {
    /**
     * Get an existing MnqSnsCredentials resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MnqSnsCredentialsState, opts?: pulumi.CustomResourceOptions): MnqSnsCredentials {
        return new MnqSnsCredentials(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/mnqSnsCredentials:MnqSnsCredentials';

    /**
     * Returns true if the given object is an instance of MnqSnsCredentials.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MnqSnsCredentials {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MnqSnsCredentials.__pulumiType;
    }

    /**
     * The ID of the key.
     */
    public /*out*/ readonly accessKey!: pulumi.Output<string>;
    /**
     * The unique name of the SNS credentials.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * . List of permissions associated with these credentials. Only one of the following permissions may be set:
     */
    public readonly permissions!: pulumi.Output<outputs.MnqSnsCredentialsPermissions>;
    /**
     * `projectId`) The ID of the Project in which SNS is enabled.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`). The region in which SNS is enabled.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The secret value of the key.
     */
    public /*out*/ readonly secretKey!: pulumi.Output<string>;

    /**
     * Create a MnqSnsCredentials resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: MnqSnsCredentialsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MnqSnsCredentialsArgs | MnqSnsCredentialsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MnqSnsCredentialsState | undefined;
            resourceInputs["accessKey"] = state ? state.accessKey : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["secretKey"] = state ? state.secretKey : undefined;
        } else {
            const args = argsOrState as MnqSnsCredentialsArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["accessKey"] = undefined /*out*/;
            resourceInputs["secretKey"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accessKey", "secretKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(MnqSnsCredentials.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MnqSnsCredentials resources.
 */
export interface MnqSnsCredentialsState {
    /**
     * The ID of the key.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * The unique name of the SNS credentials.
     */
    name?: pulumi.Input<string>;
    /**
     * . List of permissions associated with these credentials. Only one of the following permissions may be set:
     */
    permissions?: pulumi.Input<inputs.MnqSnsCredentialsPermissions>;
    /**
     * `projectId`) The ID of the Project in which SNS is enabled.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region in which SNS is enabled.
     */
    region?: pulumi.Input<string>;
    /**
     * The secret value of the key.
     */
    secretKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MnqSnsCredentials resource.
 */
export interface MnqSnsCredentialsArgs {
    /**
     * The unique name of the SNS credentials.
     */
    name?: pulumi.Input<string>;
    /**
     * . List of permissions associated with these credentials. Only one of the following permissions may be set:
     */
    permissions?: pulumi.Input<inputs.MnqSnsCredentialsPermissions>;
    /**
     * `projectId`) The ID of the Project in which SNS is enabled.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region in which SNS is enabled.
     */
    region?: pulumi.Input<string>;
}