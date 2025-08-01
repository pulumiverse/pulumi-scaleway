// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `scaleway.object.Bucket` resource allows you to create and manage buckets for [Scaleway Object storage](https://www.scaleway.com/en/docs/object-storage/).
 *
 * Refer to the [dedicated documentation](https://www.scaleway.com/en/docs/object-storage/how-to/create-a-bucket/) for more information on Object Storage buckets.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const someBucket = new scaleway.object.Bucket("some_bucket", {
 *     name: "some-unique-name",
 *     tags: {
 *         key: "value",
 *     },
 * });
 * ```
 *
 * ### Creating the bucket in a specific project
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const someBucket = new scaleway.object.Bucket("some_bucket", {
 *     name: "some-unique-name",
 *     projectId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 *
 * ### Using object lifecycle
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const main = new scaleway.object.Bucket("main", {
 *     name: "mybuckectid",
 *     region: "fr-par",
 *     lifecycleRules: [
 *         {
 *             id: "id1",
 *             prefix: "path1/",
 *             enabled: true,
 *             expiration: {
 *                 days: 365,
 *             },
 *             transitions: [{
 *                 days: 120,
 *                 storageClass: "GLACIER",
 *             }],
 *         },
 *         {
 *             id: "id2",
 *             prefix: "path2/",
 *             enabled: true,
 *             expiration: {
 *                 days: 50,
 *             },
 *         },
 *         {
 *             id: "id3",
 *             prefix: "path3/",
 *             enabled: false,
 *             tags: {
 *                 tagKey: "tagValue",
 *                 terraform: "hashicorp",
 *             },
 *             expiration: {
 *                 days: 1,
 *             },
 *         },
 *         {
 *             id: "id4",
 *             enabled: true,
 *             tags: {
 *                 tag1: "value1",
 *             },
 *             transitions: [{
 *                 days: 1,
 *                 storageClass: "GLACIER",
 *             }],
 *         },
 *         {
 *             enabled: true,
 *             abortIncompleteMultipartUploadDays: 30,
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Buckets can be imported using the `{region}/{bucketName}` identifier, as shown below:
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/objectBucket:ObjectBucket some_bucket fr-par/some-bucket
 * ```
 *
 * ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project.
 *
 * If you are using a project different from the default one, you have to specify the project ID at the end of the import command.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/objectBucket:ObjectBucket some_bucket fr-par/some-bucket@11111111-1111-1111-1111-111111111111
 * ```
 *
 * @deprecated scaleway.index/objectbucket.ObjectBucket has been deprecated in favor of scaleway.object/bucket.Bucket
 */
export class ObjectBucket extends pulumi.CustomResource {
    /**
     * Get an existing ObjectBucket resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ObjectBucketState, opts?: pulumi.CustomResourceOptions): ObjectBucket {
        pulumi.log.warn("ObjectBucket is deprecated: scaleway.index/objectbucket.ObjectBucket has been deprecated in favor of scaleway.object/bucket.Bucket")
        return new ObjectBucket(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/objectBucket:ObjectBucket';

    /**
     * Returns true if the given object is an instance of ObjectBucket.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ObjectBucket {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ObjectBucket.__pulumiType;
    }

    /**
     * (Deprecated) The canned ACL you want to apply to the bucket.
     *
     * > **Note:** The `acl` attribute is deprecated. See scaleway.object.BucketAcl resource documentation. Refer to the [official canned ACL documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) for more information on the different roles.
     *
     * @deprecated ACL attribute is deprecated. Please use the resource scaleway.object.BucketAcl instead.
     */
    public readonly acl!: pulumi.Output<string | undefined>;
    /**
     * API URL of the bucket
     */
    public /*out*/ readonly apiEndpoint!: pulumi.Output<string>;
    public readonly corsRules!: pulumi.Output<outputs.ObjectBucketCorsRule[]>;
    /**
     * The endpoint URL of the bucket.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * Boolean that, when set to true, allows the deletion of all objects (including locked objects) when the bucket is destroyed. This operation is irreversible, and the objects cannot be recovered. The default is false.
     */
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
     */
    public readonly lifecycleRules!: pulumi.Output<outputs.ObjectBucketLifecycleRule[] | undefined>;
    /**
     * The name of the bucket.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable object lock
     */
    public readonly objectLockEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * `projectId`) The ID of the project the bucket is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The [region](https://www.scaleway.com/en/developers/api/#region-definition) in which the bucket will be created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * A list of tags (key/value) for the bucket.
     *
     * * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
     * If you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Allow multiple versions of an object in the same bucket
     */
    public readonly versioning!: pulumi.Output<outputs.ObjectBucketVersioning>;

    /**
     * Create a ObjectBucket resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated scaleway.index/objectbucket.ObjectBucket has been deprecated in favor of scaleway.object/bucket.Bucket */
    constructor(name: string, args?: ObjectBucketArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated scaleway.index/objectbucket.ObjectBucket has been deprecated in favor of scaleway.object/bucket.Bucket */
    constructor(name: string, argsOrState?: ObjectBucketArgs | ObjectBucketState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ObjectBucket is deprecated: scaleway.index/objectbucket.ObjectBucket has been deprecated in favor of scaleway.object/bucket.Bucket")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ObjectBucketState | undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["apiEndpoint"] = state ? state.apiEndpoint : undefined;
            resourceInputs["corsRules"] = state ? state.corsRules : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            resourceInputs["lifecycleRules"] = state ? state.lifecycleRules : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["objectLockEnabled"] = state ? state.objectLockEnabled : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["versioning"] = state ? state.versioning : undefined;
        } else {
            const args = argsOrState as ObjectBucketArgs | undefined;
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["corsRules"] = args ? args.corsRules : undefined;
            resourceInputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            resourceInputs["lifecycleRules"] = args ? args.lifecycleRules : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["objectLockEnabled"] = args ? args.objectLockEnabled : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["versioning"] = args ? args.versioning : undefined;
            resourceInputs["apiEndpoint"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ObjectBucket.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ObjectBucket resources.
 */
export interface ObjectBucketState {
    /**
     * (Deprecated) The canned ACL you want to apply to the bucket.
     *
     * > **Note:** The `acl` attribute is deprecated. See scaleway.object.BucketAcl resource documentation. Refer to the [official canned ACL documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) for more information on the different roles.
     *
     * @deprecated ACL attribute is deprecated. Please use the resource scaleway.object.BucketAcl instead.
     */
    acl?: pulumi.Input<string>;
    /**
     * API URL of the bucket
     */
    apiEndpoint?: pulumi.Input<string>;
    corsRules?: pulumi.Input<pulumi.Input<inputs.ObjectBucketCorsRule>[]>;
    /**
     * The endpoint URL of the bucket.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * Boolean that, when set to true, allows the deletion of all objects (including locked objects) when the bucket is destroyed. This operation is irreversible, and the objects cannot be recovered. The default is false.
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
     */
    lifecycleRules?: pulumi.Input<pulumi.Input<inputs.ObjectBucketLifecycleRule>[]>;
    /**
     * The name of the bucket.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable object lock
     */
    objectLockEnabled?: pulumi.Input<boolean>;
    /**
     * `projectId`) The ID of the project the bucket is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The [region](https://www.scaleway.com/en/developers/api/#region-definition) in which the bucket will be created.
     */
    region?: pulumi.Input<string>;
    /**
     * A list of tags (key/value) for the bucket.
     *
     * * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
     * If you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Allow multiple versions of an object in the same bucket
     */
    versioning?: pulumi.Input<inputs.ObjectBucketVersioning>;
}

/**
 * The set of arguments for constructing a ObjectBucket resource.
 */
export interface ObjectBucketArgs {
    /**
     * (Deprecated) The canned ACL you want to apply to the bucket.
     *
     * > **Note:** The `acl` attribute is deprecated. See scaleway.object.BucketAcl resource documentation. Refer to the [official canned ACL documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) for more information on the different roles.
     *
     * @deprecated ACL attribute is deprecated. Please use the resource scaleway.object.BucketAcl instead.
     */
    acl?: pulumi.Input<string>;
    corsRules?: pulumi.Input<pulumi.Input<inputs.ObjectBucketCorsRule>[]>;
    /**
     * Boolean that, when set to true, allows the deletion of all objects (including locked objects) when the bucket is destroyed. This operation is irreversible, and the objects cannot be recovered. The default is false.
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
     */
    lifecycleRules?: pulumi.Input<pulumi.Input<inputs.ObjectBucketLifecycleRule>[]>;
    /**
     * The name of the bucket.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable object lock
     */
    objectLockEnabled?: pulumi.Input<boolean>;
    /**
     * `projectId`) The ID of the project the bucket is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The [region](https://www.scaleway.com/en/developers/api/#region-definition) in which the bucket will be created.
     */
    region?: pulumi.Input<string>;
    /**
     * A list of tags (key/value) for the bucket.
     *
     * * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
     * If you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Allow multiple versions of an object in the same bucket
     */
    versioning?: pulumi.Input<inputs.ObjectBucketVersioning>;
}
