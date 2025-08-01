// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `scaleway.object.BucketPolicy` data source is used to retrieve information about the bucket policy of an Object Storage bucket.
 *
 * Refer to the Object Storage [documentation](https://www.scaleway.com/en/docs/object-storage/api-cli/bucket-policy/) for more information.
 *
 * ## Retrieve the bucket policy of a bucket
 *
 * The following command allows you to retrieve a bucket policy by its bucket.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const main = scaleway.object.getBucketPolicy({
 *     bucket: "bucket.test.com",
 * });
 * ```
 */
/** @deprecated scaleway.index/getobjectbucketpolicy.getObjectBucketPolicy has been deprecated in favor of scaleway.object/getbucketpolicy.getBucketPolicy */
export function getObjectBucketPolicy(args: GetObjectBucketPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetObjectBucketPolicyResult> {
    pulumi.log.warn("getObjectBucketPolicy is deprecated: scaleway.index/getobjectbucketpolicy.getObjectBucketPolicy has been deprecated in favor of scaleway.object/getbucketpolicy.getBucketPolicy")
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getObjectBucketPolicy:getObjectBucketPolicy", {
        "bucket": args.bucket,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getObjectBucketPolicy.
 */
export interface GetObjectBucketPolicyArgs {
    /**
     * The name of the bucket.
     */
    bucket: string;
    projectId?: string;
    /**
     * `region`) The region in which the Object Storage exists.
     */
    region?: string;
}

/**
 * A collection of values returned by getObjectBucketPolicy.
 */
export interface GetObjectBucketPolicyResult {
    readonly bucket: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The content of the bucket policy in JSON format.
     */
    readonly policy: string;
    readonly projectId?: string;
    readonly region?: string;
}
/**
 * The `scaleway.object.BucketPolicy` data source is used to retrieve information about the bucket policy of an Object Storage bucket.
 *
 * Refer to the Object Storage [documentation](https://www.scaleway.com/en/docs/object-storage/api-cli/bucket-policy/) for more information.
 *
 * ## Retrieve the bucket policy of a bucket
 *
 * The following command allows you to retrieve a bucket policy by its bucket.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const main = scaleway.object.getBucketPolicy({
 *     bucket: "bucket.test.com",
 * });
 * ```
 */
/** @deprecated scaleway.index/getobjectbucketpolicy.getObjectBucketPolicy has been deprecated in favor of scaleway.object/getbucketpolicy.getBucketPolicy */
export function getObjectBucketPolicyOutput(args: GetObjectBucketPolicyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetObjectBucketPolicyResult> {
    pulumi.log.warn("getObjectBucketPolicy is deprecated: scaleway.index/getobjectbucketpolicy.getObjectBucketPolicy has been deprecated in favor of scaleway.object/getbucketpolicy.getBucketPolicy")
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getObjectBucketPolicy:getObjectBucketPolicy", {
        "bucket": args.bucket,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getObjectBucketPolicy.
 */
export interface GetObjectBucketPolicyOutputArgs {
    /**
     * The name of the bucket.
     */
    bucket: pulumi.Input<string>;
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the Object Storage exists.
     */
    region?: pulumi.Input<string>;
}
