// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manage Scaleway Messaging and queuing SNS topics.
 * For further information, see
 * our [main documentation](https://www.scaleway.com/en/docs/messaging/how-to/create-manage-topics/).
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const main = new scaleway.mnq.Sns("main", {});
 * const mainSnsCredentials = new scaleway.mnq.SnsCredentials("main", {
 *     projectId: main.projectId,
 *     permissions: {
 *         canManage: true,
 *     },
 * });
 * const topic = new scaleway.mnq.SnsTopic("topic", {
 *     projectId: main.projectId,
 *     name: "my-topic",
 *     accessKey: mainSnsCredentials.accessKey,
 *     secretKey: mainSnsCredentials.secretKey,
 * });
 * ```
 *
 * ## Import
 *
 * SNS topics can be imported using `{region}/{project-id}/{topic-name}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:mnq/snsTopic:SnsTopic main fr-par/11111111111111111111111111111111/my-topic
 * ```
 */
export class SnsTopic extends pulumi.CustomResource {
    /**
     * Get an existing SnsTopic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnsTopicState, opts?: pulumi.CustomResourceOptions): SnsTopic {
        return new SnsTopic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:mnq/snsTopic:SnsTopic';

    /**
     * Returns true if the given object is an instance of SnsTopic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SnsTopic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SnsTopic.__pulumiType;
    }

    /**
     * The access key of the SNS credentials.
     */
    public readonly accessKey!: pulumi.Output<string>;
    /**
     * The ARN of the topic
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Specifies whether to enable content-based deduplication.
     */
    public readonly contentBasedDeduplication!: pulumi.Output<boolean>;
    /**
     * Whether the topic is a FIFO topic. If true, the topic name must end with .fifo.
     */
    public readonly fifoTopic!: pulumi.Output<boolean>;
    /**
     * The unique name of the SNS topic. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * Owner of the SNS topic, should have format 'project-${project_id}'
     */
    public /*out*/ readonly owner!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the Project in which SNS is enabled.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`). The region
     * in which SNS is enabled.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The secret key of the SNS credentials.
     */
    public readonly secretKey!: pulumi.Output<string>;
    /**
     * The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
     */
    public readonly snsEndpoint!: pulumi.Output<string | undefined>;

    /**
     * Create a SnsTopic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnsTopicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnsTopicArgs | SnsTopicState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnsTopicState | undefined;
            resourceInputs["accessKey"] = state ? state.accessKey : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["contentBasedDeduplication"] = state ? state.contentBasedDeduplication : undefined;
            resourceInputs["fifoTopic"] = state ? state.fifoTopic : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["secretKey"] = state ? state.secretKey : undefined;
            resourceInputs["snsEndpoint"] = state ? state.snsEndpoint : undefined;
        } else {
            const args = argsOrState as SnsTopicArgs | undefined;
            if ((!args || args.accessKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessKey'");
            }
            if ((!args || args.secretKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretKey'");
            }
            resourceInputs["accessKey"] = args?.accessKey ? pulumi.secret(args.accessKey) : undefined;
            resourceInputs["contentBasedDeduplication"] = args ? args.contentBasedDeduplication : undefined;
            resourceInputs["fifoTopic"] = args ? args.fifoTopic : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["secretKey"] = args?.secretKey ? pulumi.secret(args.secretKey) : undefined;
            resourceInputs["snsEndpoint"] = args ? args.snsEndpoint : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "scaleway:index/mnqSnsTopic:MnqSnsTopic" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        const secretOpts = { additionalSecretOutputs: ["accessKey", "secretKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SnsTopic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SnsTopic resources.
 */
export interface SnsTopicState {
    /**
     * The access key of the SNS credentials.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * The ARN of the topic
     */
    arn?: pulumi.Input<string>;
    /**
     * Specifies whether to enable content-based deduplication.
     */
    contentBasedDeduplication?: pulumi.Input<boolean>;
    /**
     * Whether the topic is a FIFO topic. If true, the topic name must end with .fifo.
     */
    fifoTopic?: pulumi.Input<boolean>;
    /**
     * The unique name of the SNS topic. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * Owner of the SNS topic, should have format 'project-${project_id}'
     */
    owner?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the Project in which SNS is enabled.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region
     * in which SNS is enabled.
     */
    region?: pulumi.Input<string>;
    /**
     * The secret key of the SNS credentials.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
     */
    snsEndpoint?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SnsTopic resource.
 */
export interface SnsTopicArgs {
    /**
     * The access key of the SNS credentials.
     */
    accessKey: pulumi.Input<string>;
    /**
     * Specifies whether to enable content-based deduplication.
     */
    contentBasedDeduplication?: pulumi.Input<boolean>;
    /**
     * Whether the topic is a FIFO topic. If true, the topic name must end with .fifo.
     */
    fifoTopic?: pulumi.Input<boolean>;
    /**
     * The unique name of the SNS topic. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the Project in which SNS is enabled.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region
     * in which SNS is enabled.
     */
    region?: pulumi.Input<string>;
    /**
     * The secret key of the SNS credentials.
     */
    secretKey: pulumi.Input<string>;
    /**
     * The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
     */
    snsEndpoint?: pulumi.Input<string>;
}
