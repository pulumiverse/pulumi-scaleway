// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `scaleway.object.BucketWebsiteConfiguration` resource allows you to deploy and manage a bucket website with [Scaleway Object storage](https://www.scaleway.com/en/docs/object-storage/).
 *
 * Refer to the [dedicated documentation](https://www.scaleway.com/en/docs/object-storage/how-to/use-bucket-website/) for more information on bucket websites.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const test = new scaleway.object.Bucket("test", {
 *     name: "my-bucket",
 *     acl: "public-read",
 * });
 * const someFile = new scaleway.object.Item("some_file", {
 *     bucket: test.name,
 *     key: "index.html",
 *     file: "index.html",
 *     visibility: "public-read",
 *     contentType: "text/html",
 * });
 * const testBucketWebsiteConfiguration = new scaleway.object.BucketWebsiteConfiguration("test", {
 *     bucket: test.name,
 *     indexDocument: {
 *         suffix: "index.html",
 *     },
 *     errorDocument: {
 *         key: "error.html",
 *     },
 * });
 * ```
 *
 * ### With A Bucket Policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const main = new scaleway.object.Bucket("main", {
 *     name: "MyBucket",
 *     acl: "public-read",
 * });
 * const mainBucketPolicy = new scaleway.object.BucketPolicy("main", {
 *     bucket: main.id,
 *     policy: JSON.stringify({
 *         Version: "2012-10-17",
 *         Id: "MyPolicy",
 *         Statement: [{
 *             Sid: "GrantToEveryone",
 *             Effect: "Allow",
 *             Principal: "*",
 *             Action: ["s3:GetObject"],
 *             Resource: ["<bucket-name>/*"],
 *         }],
 *     }),
 * });
 * const mainBucketWebsiteConfiguration = new scaleway.object.BucketWebsiteConfiguration("main", {
 *     bucket: main.id,
 *     indexDocument: {
 *         suffix: "index.html",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Bucket website configurations can be imported using the `{region}/{bucketName}` identifier, as shown below:
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration some_bucket fr-par/some-bucket
 * ```
 *
 * ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project.
 *
 * If you are using a project different from the default one, you have to specify the project ID at the end of the import command.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration some_bucket fr-par/some-bucket@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
 * ```
 *
 * @deprecated scaleway.index/objectbucketwebsiteconfiguration.ObjectBucketWebsiteConfiguration has been deprecated in favor of scaleway.object/bucketwebsiteconfiguration.BucketWebsiteConfiguration
 */
export class ObjectBucketWebsiteConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing ObjectBucketWebsiteConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ObjectBucketWebsiteConfigurationState, opts?: pulumi.CustomResourceOptions): ObjectBucketWebsiteConfiguration {
        pulumi.log.warn("ObjectBucketWebsiteConfiguration is deprecated: scaleway.index/objectbucketwebsiteconfiguration.ObjectBucketWebsiteConfiguration has been deprecated in favor of scaleway.object/bucketwebsiteconfiguration.BucketWebsiteConfiguration")
        return new ObjectBucketWebsiteConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration';

    /**
     * Returns true if the given object is an instance of ObjectBucketWebsiteConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ObjectBucketWebsiteConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ObjectBucketWebsiteConfiguration.__pulumiType;
    }

    /**
     * The name of the bucket.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * The name of the error file for the website detailed below.
     */
    public readonly errorDocument!: pulumi.Output<outputs.ObjectBucketWebsiteConfigurationErrorDocument | undefined>;
    /**
     * The name of the index file for the website detailed below.
     */
    public readonly indexDocument!: pulumi.Output<outputs.ObjectBucketWebsiteConfigurationIndexDocument>;
    /**
     * The projectId you want to attach the resource to
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The region you want to attach the resource to
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The domain of the website endpoint. This is used to create DNS alias [records](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/).
     */
    public /*out*/ readonly websiteDomain!: pulumi.Output<string>;
    /**
     * The website endpoint.
     */
    public /*out*/ readonly websiteEndpoint!: pulumi.Output<string>;

    /**
     * Create a ObjectBucketWebsiteConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated scaleway.index/objectbucketwebsiteconfiguration.ObjectBucketWebsiteConfiguration has been deprecated in favor of scaleway.object/bucketwebsiteconfiguration.BucketWebsiteConfiguration */
    constructor(name: string, args: ObjectBucketWebsiteConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated scaleway.index/objectbucketwebsiteconfiguration.ObjectBucketWebsiteConfiguration has been deprecated in favor of scaleway.object/bucketwebsiteconfiguration.BucketWebsiteConfiguration */
    constructor(name: string, argsOrState?: ObjectBucketWebsiteConfigurationArgs | ObjectBucketWebsiteConfigurationState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ObjectBucketWebsiteConfiguration is deprecated: scaleway.index/objectbucketwebsiteconfiguration.ObjectBucketWebsiteConfiguration has been deprecated in favor of scaleway.object/bucketwebsiteconfiguration.BucketWebsiteConfiguration")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ObjectBucketWebsiteConfigurationState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["errorDocument"] = state ? state.errorDocument : undefined;
            resourceInputs["indexDocument"] = state ? state.indexDocument : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["websiteDomain"] = state ? state.websiteDomain : undefined;
            resourceInputs["websiteEndpoint"] = state ? state.websiteEndpoint : undefined;
        } else {
            const args = argsOrState as ObjectBucketWebsiteConfigurationArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.indexDocument === undefined) && !opts.urn) {
                throw new Error("Missing required property 'indexDocument'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["errorDocument"] = args ? args.errorDocument : undefined;
            resourceInputs["indexDocument"] = args ? args.indexDocument : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["websiteDomain"] = undefined /*out*/;
            resourceInputs["websiteEndpoint"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ObjectBucketWebsiteConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ObjectBucketWebsiteConfiguration resources.
 */
export interface ObjectBucketWebsiteConfigurationState {
    /**
     * The name of the bucket.
     */
    bucket?: pulumi.Input<string>;
    /**
     * The name of the error file for the website detailed below.
     */
    errorDocument?: pulumi.Input<inputs.ObjectBucketWebsiteConfigurationErrorDocument>;
    /**
     * The name of the index file for the website detailed below.
     */
    indexDocument?: pulumi.Input<inputs.ObjectBucketWebsiteConfigurationIndexDocument>;
    /**
     * The projectId you want to attach the resource to
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region you want to attach the resource to
     */
    region?: pulumi.Input<string>;
    /**
     * The domain of the website endpoint. This is used to create DNS alias [records](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/).
     */
    websiteDomain?: pulumi.Input<string>;
    /**
     * The website endpoint.
     */
    websiteEndpoint?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ObjectBucketWebsiteConfiguration resource.
 */
export interface ObjectBucketWebsiteConfigurationArgs {
    /**
     * The name of the bucket.
     */
    bucket: pulumi.Input<string>;
    /**
     * The name of the error file for the website detailed below.
     */
    errorDocument?: pulumi.Input<inputs.ObjectBucketWebsiteConfigurationErrorDocument>;
    /**
     * The name of the index file for the website detailed below.
     */
    indexDocument: pulumi.Input<inputs.ObjectBucketWebsiteConfigurationIndexDocument>;
    /**
     * The projectId you want to attach the resource to
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region you want to attach the resource to
     */
    region?: pulumi.Input<string>;
}
